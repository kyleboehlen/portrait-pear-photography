<template>
  <main class="flex">
    <!-- No images found message -->
    <Transition name="fade">
      <NotFoundMessage v-if="apiCallFinished && filteredPhotos.length == 0" :msg="notFoundMsg" class="w-full" />
    </Transition>

    <!-- Actually show images! -->
    <Transition name="fade">
      <div
        v-if="apiCallFinished && filteredPhotos.length > 0"
        class="w-full flex flex-wrap items-center justify-around">
        <PhotoCard
          v-for="photo in filteredPhotos"
          class="hover:cursor-zoom-in"
          :key="photo.id"
          :photo="photo"
          :cachedPhotos="cachedPhotos"
          cache="true"
          @click="fullsizeImage(photo)" />
      </div>
    </Transition>

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
// Vue
import { onMounted, ref, computed } from "vue"
// Components
import NotFoundMessage from "@/components/panel/NotFoundMessage.vue"
import PhotoCard from "@/components/panel/PhotoCard.vue"
import TheFullsizeViewer from "@/components/TheFullsizeViewer.vue"
// Store
import { useHomeStore } from "@/stores/home.js"
import { useImagesStore } from "@/stores/images.js"
import { useFilterStore } from "@/stores/filter.js"
// Capacitor
import { Capacitor } from "@capacitor/core"

const home = useHomeStore()
const images = useImagesStore()

// Show loader logic
const apiCallFinished = ref(false)
const isOnline = ref(false)

// Not found message
const notFoundMsg = computed(() => {
  if (isOnline.value) {
    return "No photos found :("
  }

  return "Go online to load photos :)"
})

// Photo refs
const photos = ref([])
const cachedPhotos = ref([])

// Filtering
const filter = useFilterStore()
const filteredPhotos = computed(() => {
  let returnPhotos = []

  // Filter to cached
  if (Capacitor.isNativePlatform() && !isOnline.value) {
    // Check cache
    returnPhotos = photos.value.filter((p) => cachedPhotos.value.includes(p.compressed_asset_url))
  } else {
    returnPhotos = photos.value
  }

  // Filter by category
  if (filter.category > 0) {
    returnPhotos = returnPhotos.filter((p) => p.categories_array.includes(filter.category))
  }

  return returnPhotos
})

onMounted(() => {
  // Set cached images if we're native
  if (Capacitor.isNativePlatform()) {
    setCachedImages().then(() => {
      setHomePhotos()
    })
  } else {
    setHomePhotos()
  }
})

const setHomePhotos = () => {
  // If online we'll refresh the content from the API
  if (navigator.onLine) {
    refreshHomePhotos()
  } else {
    // We'll refresh content when we come online
    window.addEventListener("online", refreshHomePhotos) // Bind to online event

    // Get from cache if we're native
    if (Capacitor.isNativePlatform()) {
      // Get image info from cache
      photos.value = home.photos
    }

    apiCallFinished.value = true
  }
}

const apiUrl = import.meta.env.VITE_API_URL
const refreshHomePhotos = () => {
  isOnline.value = true
  // Get from cache if we have updated in the past 30 minutes
  if (Date.now() - home.lastUpdated < 1800000) {
    // Get from cache
    photos.value = home.photos
    apiCallFinished.value = true
  } else {
    axios({
      method: "get",
      url: `${apiUrl}/api/pear/home`,
    })
      .then(function (response) {
        // Cache
        home.photos = response.data.photos
        home.setLastUpdated()

        // Set photos
        photos.value = response.data.photos
        apiCallFinished.value = true
      })
      .catch(function () {
        apiCallFinished.value = true
      })
  }

  // Set cached
  if (Capacitor.isNativePlatform()) {
    setCachedImages()
  }
}

async function setCachedImages() {
  // Get cached images
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
