<script setup>
import {computed, onMounted, ref} from "vue"
import logo from "@/assets/imgs/green-camera-logo.png?url"
import {usePhotoUtils} from "@/composables/usePhotoUtils.js"

const props = defineProps(["photo"])

const img = ref()

// const lazyLifecycle = {
//   loaded: onLoad,
// }

const {previewUrl} = usePhotoUtils()
const photoUrl = previewUrl(props.photo.guid)
</script>

<template>
  <div
      class="flex justify-center bg-secondary/50 max-xs:w-full max-sm:w-3/4 sm:h-96 border shadow-md shadow-secondary rounded-lg border-secondary mt-4 p-1">
    <img :src="photoUrl" class="max-w-full max-h-full object-contain rounded-lg" />
  </div>
</template>