<template>
  <!-- Favorite heart button -->
  <span class="fixed right-0 bottom-0 pb-4 pr-4">
    <button class="btn btn-lg btn-circle" @click="handleFavorite">
      <Icon v-if="!isFavorited" icon="heart" class="w-2/3 h-auto" />
      <Icon v-else icon="heartFilled" class="w-2/3 h-auto text-error" />
    </button>
  </span>

  <!-- Add Favorite modal -->
  <input type="checkbox" id="add-modal" class="modal-toggle" />
  <div id="add-modal" class="modal" :class="{ 'modal-open': showAddModal }" @click.self="showAddModal = false">
    <div class="modal-box p-0">
      <div class="card w-full bg-neutral text-neutral-content">
        <div class="card-body items-center text-center">
          <h2 class="card-title">Favorite Shoot</h2>
          <div class="form-control w-full max-w-xs">
            <label class="label">
              <span class="label-text">Name</span>
            </label>
            <input
              v-model="shootName"
              type="text"
              placeholder="My Awesome Shoot"
              class="input input-bordered input-primary focus:outline-none w-full max-w-xs"
              :class="{ 'input-primary': shootNameValid, 'input-error': !shootNameValid }"
              @focus="shootNameValid = true"
              @keyup.enter="saveFavorite" />
          </div>
          <div class="card-actions justify-end w-full max-w-xs pt-4">
            <button class="btn btn-ghost" @click="showAddModal = false">Cancel</button>
            <div class="w-2"></div>
            <button class="btn btn-primary" @click="saveFavorite">Favorite</button>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Percentage loader -->
  <Teleport to="body">
    <div
      v-if="showSavingModal"
      class="radial-progress bg-primary text-primary-content fixed top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2"
      style="--size: 12rem; --thickness: 1rem"
      :style="{ '--value': savePercentage }">
      <span class="!text-4xl">{{ savePercentage }}%</span>
    </div>
  </Teleport>

  <!-- Remove Favorite modal -->
  <input type="checkbox" id="remove-modal" class="modal-toggle" />
  <div id="remove-modal" class="modal" :class="{ 'modal-open': showRemoveModal }" @click.self="showAddModal = false">
    <div class="modal-box p-0">
      <div class="card w-full bg-neutral text-neutral-content">
        <div class="card-body items-center text-center">
          <h2 class="card-title">Remove As Favorite?</h2>
          <div class="card-actions justify-center w-full max-w-xs pt-4">
            <button class="btn btn-ghost" @click="showRemoveModal = false">Cancel</button>
            <div class="w-2"></div>
            <button class="btn btn-primary" @click="removeFavorite">Remove</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// Vue
import { computed, ref, onMounted } from "vue"
// Icons
import { Icon, addIcon } from "@iconify/vue/offline"
import heart from "@iconify-icons/mdi/cards-heart-outline"
import heartFilled from "@iconify-icons/mdi/cards-heart"
// Store
import { useFavoritesStore } from "@/stores/favorites.js"
import { useImagesStore } from "@/stores/images.js"
// Composable
import { toDataURL } from "@/composables/todataurl.js"

const favorites = useFavoritesStore()

addIcon("heart", heart)
addIcon("heartFilled", heartFilled)

const props = defineProps(["shootSlug", "photos"])
const emit = defineEmits(["refreshCache"])

const isFavorited = computed(() => {
  const favorite = favorites.shoots.find((shoot) => {
    return shoot.slug === props.shootSlug
  })

  return favorite !== undefined
})

const showAddModal = ref(false)
const showRemoveModal = ref(false)
const showSavingModal = ref(false)
const shootName = ref()
const shootNameValid = ref(true)

const handleFavorite = () => {
  if (isFavorited.value) {
    // Are you sure modal
    showRemoveModal.value = true
  } else {
    // Add favorites modal
    showAddModal.value = true
  }
}

const saveFavorite = () => {
  if (shootName.value === undefined || shootName.value.length === 0) {
    shootNameValid.value = false
  } else {
    favorites.addShoot(props.shootSlug, shootName.value, props.photos)
    showAddModal.value = false
    saveImages()
  }
}

const removeFavorite = () => {
  favorites.removeShoot(props.shootSlug)
  showRemoveModal.value = false
}

onMounted(() => {
  favorites.loadFromUserDefaults()
})

// Saving
const numToSave = computed(() => {
  if (props.photos === undefined || props.photos === null) {
    return 0
  }
  return props.photos.length * 2
})

const numSaved = ref(0)
const savePercentage = computed(() => {
  if (numToSave.value == 0 || numSaved.value == 0) {
    return 0
  }
  return Math.round((numSaved.value / numToSave.value) * 100)
})

// Caching images on favoriting of shoot
const images = useImagesStore()
const saveImages = async () => {
  showSavingModal.value = true

  const allAssetURLs = await images.getAll()

  await props.photos.forEach(async (photo) => {
    if (!allAssetURLs.includes(photo.compressed_asset_url)) {
      toDataURL(photo.compressed_asset_url, async (dataUrl) => {
        await images.store(photo.compressed_asset_url, dataUrl)
        incrementSave()
      })
    } else {
      incrementSave()
    }
    if (!allAssetURLs.includes(photo.full_res_asset_url)) {
      toDataURL(photo.full_res_asset_url, async (dataUrl) => {
        await images.store(photo.full_res_asset_url, dataUrl)
        incrementSave()
      })
    } else {
      incrementSave()
    }
  })

  if (numSaved.value === numToSave.value) {
    showSavingModal.value = false
    emit("refreshCache")
  }
}

const incrementSave = () => {
  numSaved.value++
  if (numSaved.value === numToSave.value) {
    showSavingModal.value = false
    emit("refreshCache")
  }
}
</script>
