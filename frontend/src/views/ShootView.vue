<template>
  <main class="flex">
    <!-- Shoot not found :( -->
    <Transition name="fade">
      <NotFoundMessage v-if="apiCallFinished && !shootFound" :msg="notFoundMsg" class="w-full" />
    </Transition>

    <!-- The photos, yay -->
    <Transition name="fade">
      <div v-if="isShoot" class="w-full flex flex-wrap items-center justify-around">
        <PhotoCard
          v-for="photo in filteredPhotos"
          :key="photo.id"
          :photo="photo"
          :cachedPhotos="cachedPhotos"
          class="hover:cursor-zoom-in"
          @click="fullsizeImage(photo)" />
      </div>
    </Transition>

    <!-- Favorites button -->
    <TheFavoritesButton
      v-if="showFavoritesButton"
      :shootSlug="shootSlug"
      :photos="photos"
      @refreshCache="loadCachedImages" />

    <!-- Full size viewer -->
    <TheFullsizeViewer
      v-if="showFullsize"
      class="modal-open"
      :photo="fullsizePhoto"
      :photos="filteredPhotos"
      :cachedPhotos="cachedPhotos"
      @close="showFullsize = false" />
  </main>
</template>

<script setup>
// Axios
import axios from "axios"
// Capacitor
import { Capacitor } from "@capacitor/core"
// Vue
import { onMounted, computed, ref, watch } from "vue"
// Components
import TheFavoritesButton from "@/components/TheFavoritesButton.vue"
import NotFoundMessage from "@/components/panel/NotFoundMessage.vue"
import PhotoCard from "@/components/panel/PhotoCard.vue"
import TheFullsizeViewer from "@/components/TheFullsizeViewer.vue"
// Store
import { useFilterStore } from "@/stores/filter.js"
import { useFavoritesStore } from "@/stores/favorites.js"
import { useImagesStore } from "@/stores/images.js"

const props = defineProps(["shoot_slug"])

// Photo refs
const photos = ref(null)
const cachedPhotos = ref([])

// Stores
const favorites = useFavoritesStore()
const images = useImagesStore()

const shootSlug = computed(() => {
  if (props.shoot_slug === undefined || props.shoot_slug === null) {
    return null
  }
  return props.shoot_slug.toUpperCase()
})

const shootFound = ref(false)
const apiCallFinished = ref(false)
const isShoot = computed(() => {
  return shootSlug.value !== null && apiCallFinished.value && shootFound.value
})

const showFavoritesButton = computed(() => {
  return isShoot.value && Capacitor.isNativePlatform()
})

const filter = useFilterStore()
const filteredPhotos = computed(() => {
  if (filter.category > 0) {
    return photos.value.filter((p) => p.categories_array.includes(filter.category))
  }
  return photos.value
})

const apiUrl = import.meta.env.VITE_API_URL
const notFoundMsg = ref("")
const getShootFromAPI = async () => {
  apiCallFinished.value = false
  shootFound.value = false
  notFoundMsg.value = "No shoot found :("

  // Load cache if native
  if (Capacitor.isNativePlatform()) {
    loadCachedImages()
  }

  if (Capacitor.isNativePlatform() && !navigator.onLine) {
    // Load favorites from cache
    await favorites.loadFromUserDefaults()
    const shootIndex = favorites.shoots.findIndex((s) => s.slug === shootSlug.value)
    if (shootIndex < 0) {
      apiCallFinished.value = true
      shootFound.value = false
      notFoundMsg.value = "Go online to load shoot :)"
    } else {
      apiCallFinished.value = true
      shootFound.value = true
      photos.value = favorites.shoots[shootIndex].photos
    }
  } else if (!navigator.onLine) {
    apiCallFinished.value = true
    shootFound.value = false
    notFoundMsg.value = "Go online to load shoot :)"
  } else {
    axios({
      method: "get",
      url: `${apiUrl}/api/pear/shoot/${shootSlug.value}`,
    })
      .then(function (response) {
        apiCallFinished.value = true
        shootFound.value = true
        photos.value = response.data.photos
        // Update cache if favorited
        updateFavoritedShoot(photos.value)
      })
      .catch(function () {
        apiCallFinished.value = true
        shootFound.value = false
      })
  }
}

const updateFavoritedShoot = async (photos) => {
  await favorites.loadFromUserDefaults()
  const shootIndex = favorites.shoots.findIndex((s) => s.slug === shootSlug.value)
  if (shootIndex >= 0) {
    favorites.shoots[shootIndex].photos = photos
    favorites.persistToUserDefaults()
  }
}

onMounted(() => {
  // Default filters to show all
  filter.setCategory(0)

  // Set online listener - regardless if online
  if (Capacitor.isNativePlatform()) {
    window.addEventListener("online", getShootFromAPI) // Bind to online event
  }

  // Load shoot
  getShootFromAPI()
})

// Update when we get a new shoot slug
watch(props, async () => {
  window.scrollTo(0, 0)
  getShootFromAPI()
})

const loadCachedImages = () => {
  // Load cache if native
  images.getAll().then((values) => {
    cachedPhotos.value = values
  })
}

// Fullsize viewer
const showFullsize = ref(false)
const fullsizePhoto = ref(null)
const fullsizeImage = (photo) => {
  fullsizePhoto.value = photo
  showFullsize.value = true
}
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease;
}

.fade-enter-active {
  transition-delay: 0.5s;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
