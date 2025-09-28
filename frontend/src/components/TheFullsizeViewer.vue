<template>
  <div class="modal w-screen h-screen">
    <div
      class="bg-base-100 grow w-screen h-screen flex justify-between sm:p-2 flex-nowrap"
      :class="{ 'flex-col': !isNative, 'flex-col-reverse': isNative }">
      <!-- Actions bar -->
      <div
        class="flex justify-between items-center w-auto py-2 px-4 sm:px-6"
        :class="{ 'pb-6': isNative, 'px-6': isNative }">
        <!-- Download -->
        <a class="btn btn-sm xs:btn-md sm:btn-lg" :href="downloadLink" target="_blank">
          <Icon icon="download" class="h-4/6 w-auto px-0 sm:px-1" />
        </a>

        <!-- Toggle back/forth -->
        <div class="flex justify-center items-center">
          <button class="btn btn-sm xs:btn-md sm:btn-lg" @click="goLeft">
            <Icon icon="left" class="h-4/6 w-auto px-0 sm:px-1" />
          </button>

          <h2 class="text-xl xs:text-2xl sm:text-3xl pb-1 sm:pb-2 px-4 sm:px-6">{{ photoIndex }}/{{ photoCount }}</h2>

          <button class="btn btn-sm xs:btn-md sm:btn-lg" @click="goRight">
            <Icon icon="right" class="h-4/6 w-auto px-0 sm:px-1" />
          </button>
        </div>

        <!-- Close -->
        <button class="btn btn-sm xs:btn-md sm:btn-lg" @click="$emit('close')">
          <Icon icon="close" class="h-4/6 w-auto" />
        </button>
      </div>

      <!-- Image container -->
      <div
        v-if="isOnline || (useCached && !loading && cachedSource.length > 100)"
        class="w-auto h-5/6 flex justify-center items-center shrink m-1 grow"
        @touchstart="touchStart($event)"
        @touchend="touchEnd($event)">
        <img
          v-if="!useCached && !loading"
          v-lazy="{ src: photo?.full_res_asset_url, loading: pearLoader, delay: 1000 }"
          class="object-contain max-h-full rounded-lg" />
        <img v-else-if="!loading" :src="cachedSource" class="object-contain max-h-full rounded-lg" />
      </div>
      <NotFoundMessage v-else-if="!loading" msg="Go online to load full-res images :)" class="grow" />

      <div class="modal" id="download-modal">
        <div class="modal-box">
          <p class="py-4">Long press the image to save it to your gallery :)</p>
          <div class="modal-action">
            <a href="#" class="btn btn-success">Ok</a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// Vue
import { ref, computed, onMounted, nextTick } from "vue"
// Loader
import pearLoader from "@/assets/imgs/metronome-pear-white.apng?url"
// Icons
import { Icon, addIcon } from "@iconify/vue/offline"
import close from "@iconify-icons/pajamas/close"
import download from "@iconify-icons/pajamas/download"
import left from "@iconify-icons/pajamas/chevron-lg-left"
import right from "@iconify-icons/pajamas/chevron-lg-right"
// Capacitor
import { Capacitor } from "@capacitor/core"
// Components
import NotFoundMessage from "@/components/panel/NotFoundMessage.vue"
// Stores
import { useImagesStore } from "@/stores/images.js"

addIcon("close", close)
addIcon("download", download)
addIcon("left", left)
addIcon("right", right)

const props = defineProps(["photo", "photos", "cachedPhotos"])

const isNative = Capacitor.isNativePlatform()
const isOnline = ref(navigator.onLine)
const loading = ref(true)

const downloadLink = computed(() => {
  return photo?.value?.full_res_asset_url
})

const isCached = computed(() => {
  if (props.cachedPhotos !== undefined && photo.value !== null && isNative) {
    return props.cachedPhotos.includes(photo.value.full_res_asset_url)
  }

  return false
})

const useCached = computed(() => {
  return isCached.value && isNative && !isOnline.value
})

const photo = ref(null)
const cachedSource = ref(null)
const images = useImagesStore()
onMounted(() => {
  photo.value = props.photo
  // Event listener to go online
  window.addEventListener("online", () => {
    isOnline.value = true
  })
  if (isNative && isCached.value) {
    loadCachedSource()
  } else {
    loading.value = false
  }
})
const photoIndex = computed(() => {
  if (photo.value !== null) {
    const index = props.photos.findIndex((p) => p.id == photo.value.id)
    return index + 1
  }
  return null
})
const photoCount = computed(() => {
  if (photo.value !== null) {
    return props.photos.length
  }
  return null
})

const goLeft = () => {
  loading.value = true
  nextTick(() => {
    if (photoIndex.value <= 1) {
      photo.value = props.photos[props.photos.length - 1]
    } else {
      photo.value = props.photos[photoIndex.value - 2]
    }
    if (isNative) {
      loadCachedSource()
    } else {
      loading.value = false
    }
  })
}
const goRight = () => {
  loading.value = true
  nextTick(() => {
    if (photoIndex.value >= props.photos.length) {
      photo.value = props.photos[0]
    } else {
      photo.value = props.photos[photoIndex.value]
    }
    if (isNative) {
      loadCachedSource()
    } else {
      loading.value = false
    }
  })
}

// Swipe
let touchstartX = 0
let touchendX = 0
function checkSwipe() {
  if (Math.abs(touchstartX - touchendX) > window.innerWidth / 4) {
    if (touchendX < touchstartX) {
      console.log(touchendX - touchstartX)
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
  if (isNative) {
    checkSwipe()
  }
}

const loadCachedSource = () => {
  if (isCached.value) {
    images.getBase64(photo.value.full_res_asset_url).then((src) => {
      cachedSource.value = src
      loading.value = false
    })
  } else {
    loading.value = false
  }
}
</script>
