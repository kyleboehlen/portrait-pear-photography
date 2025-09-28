<template>
  <div
    class="flex justify-center bg-secondary/50 max-xs:w-full max-sm:w-3/4 sm:h-96 border-1 shadow-md shadow-secondary rounded-lg border-secondary m-4">
    <!-- If not cached show lazy load, otherwise paint it from cache -->
    <img
      v-if="crossorigin"
      ref="img"
      v-lazy="{ src: props.photo.compressed_asset_url, loading: logo, delay: 500, lifecycle: lazyLifecycle }"
      :class="imgClasses"
      crossorigin="anonymous" />
    <img
      v-else-if="!useCached"
      v-lazy="{ src: props.photo.compressed_asset_url, loading: logo, delay: 500 }"
      :class="imgClasses" />
    <img v-else :src="cachedSource" :class="imgClasses" />
  </div>
</template>

<script setup>
// Vue
import { computed, ref, onMounted } from "vue"
// Images
import logo from "@/assets/imgs/green-camera-logo.png?url"
// Store
import { useImagesStore } from "@/stores/images.js"
// Capacitor
import { Capacitor } from "@capacitor/core"
// Composable
import { toDataURL } from "@/composables/todataurl.js"

const props = defineProps(["photo", "cache", "cachedPhotos"])

const images = useImagesStore()

const img = ref()
const isOnline = ref(false)

const isCached = computed(() => {
  if (props.cachedPhotos !== undefined) {
    return props.cachedPhotos.includes(props.photo.compressed_asset_url)
  }

  return false
})

const useCached = computed(() => {
  return isCached.value && Capacitor.isNativePlatform() && !isOnline.value
})

const crossorigin = computed(() => {
  if (Capacitor.isNativePlatform() && !isCached.value) {
    return true
  }
  return false
})

const cachedSource = ref(null)
onMounted(() => {
  // Handle online state
  isOnline.value = navigator.onLine
  window.addEventListener("online", setOnline)
  window.addEventListener("offline", setOffline)

  if (isCached.value) {
    images.getBase64(props.photo.compressed_asset_url).then((src) => {
      cachedSource.value = src
    })
  }
})

const setOnline = () => {
  isOnline.value = true
}

const setOffline = () => {
  isOnline.value = false
}

// const toDataURL = useToDataURL()
const onLoad = () => {
  if (props.cache && !isCached.value && Capacitor.isNativePlatform() && navigator.onLine) {
    toDataURL(props.photo.full_res_asset_url, async (dataUrl) => {
      images.store(props.photo.compressed_asset_url, dataUrl)
    })
  }
}

const lazyLifecycle = {
  loaded: onLoad,
}

const imgClasses = "max-w-full max-h-full object-contain rounded-lg"
</script>
