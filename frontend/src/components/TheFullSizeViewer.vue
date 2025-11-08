<script setup>
import {computed, nextTick, onMounted, ref} from "vue"
import pearLoader from "@/assets/imgs/metronome-pear-white.apng?url"
import {addIcon, Icon} from "@iconify/vue/offline"
import close from "@iconify-icons/pajamas/close"
import download from "@iconify-icons/pajamas/download"
import left from "@iconify-icons/pajamas/chevron-lg-left"
import right from "@iconify-icons/pajamas/chevron-lg-right"
import {usePhotosStore} from "@/stores/usePhotosStore.js";
import {usePhotoUtils} from "@/composables/usePhotoUtils.js";

addIcon("close", close)
addIcon("download", download)
addIcon("left", left)
addIcon("right", right)

const isLoading = ref(true)

const photosStore = usePhotosStore()
const {fullResUrl, previewUrl} = usePhotoUtils()
const downloadLink = computed(() => {
  return fullResUrl(photosStore.selectedPhoto.guid)
})

const displaySrc = computed(() => {
  return previewUrl(photosStore.selectedPhoto.guid)
})

onMounted(() => {
  isLoading.value = false
})

const numDisplayPhotos = computed(() => {
  return photosStore.displayPhotos.length
})

const goLeft = () => {
  isLoading.value = true

  nextTick(() => {
    let nextIndex = photosStore.selectedPhotoIndex - 1
    if (nextIndex < 0) {
      nextIndex = numDisplayPhotos.value - 1
    }

    photosStore.selectedPhotoId = photosStore.displayPhotos[nextIndex].id

    isLoading.value = false
  })
}
const goRight = () => {
  isLoading.value = true

  nextTick(() => {
    let nextIndex = photosStore.selectedPhotoIndex + 1
    if (nextIndex >= numDisplayPhotos.value) {
      nextIndex = 0
    }

    photosStore.selectedPhotoId = photosStore.displayPhotos[nextIndex].id

    isLoading.value = false
  })
}

// Swipe
let touchstartX = 0
let touchendX = 0

function checkSwipe() {
  if (Math.abs(touchstartX - touchendX) > window.innerWidth / 4) {
    if (touchendX < touchstartX) {
      goLeft()
    }
    if (touchendX > touchstartX) {
      goRight()
    }
  }
}

const touchStart = (e) => {
  touchstartX = e.changedTouches[0].screenX
}
const touchEnd = (e) => {
  touchendX = e.changedTouches[0].screenX
  checkSwipe()
}
</script>

<template>
  <div class="modal w-screen h-screen">
    <div
        class="bg-base-100 grow w-screen h-screen flex justify-between sm:p-2 flex-nowrap flex-col">
      <!-- Actions bar -->
      <div
          class="flex justify-between items-center w-auto py-2 px-4 sm:px-6">
        <!-- Download -->
        <a class="btn btn-sm btn-primary xs:btn-md sm:btn-lg" :href="downloadLink" target="_blank">
          <Icon icon="download" class="h-4/6 w-auto px-0 sm:px-1"/>
        </a>

        <!-- Toggle back/forth -->
        <div class="flex justify-center items-center">
          <button class="btn btn-sm btn-primary xs:btn-md sm:btn-lg" @click="goLeft">
            <Icon icon="left" class="h-4/6 w-auto px-0 sm:px-1"/>
          </button>

          <h2 class="text-white/70 text-xl xs:text-2xl sm:text-3xl pb-1 sm:pb-2 px-4 sm:px-6">
            {{ photosStore.selectedPhotoIndex + 1 }}/{{ numDisplayPhotos }}</h2>

          <button class="btn btn-sm btn-primary xs:btn-md sm:btn-lg" @click="goRight">
            <Icon icon="right" class="h-4/6 w-auto px-0 sm:px-1"/>
          </button>
        </div>

        <!-- Close -->
        <button class="btn btn-sm btn-primary xs:btn-md sm:btn-lg" @click="$emit('close')">
          <Icon icon="close" class="h-4/6 w-auto"/>
        </button>
      </div>

      <!-- Image container -->
      <div
          v-if="!isLoading"
          class="w-auto h-5/6 flex justify-center items-center shrink m-1 grow"
          @touchstart="touchStart($event)"
          @touchend="touchEnd($event)">
        <img
            v-if="!isLoading"
            v-lazy="{ src: displaySrc, loading: pearLoader, delay: 1000 }"
            class="object-contain max-h-full rounded-lg"/>
      </div>
    </div>
  </div>
</template>
