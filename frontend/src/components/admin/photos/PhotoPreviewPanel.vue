<script setup>
import {computed} from 'vue'
import {useAdminStore} from "@/stores/useAdminStore";
import logo from "@/assets/imgs/green-camera-logo.png?url"

const adminStore = useAdminStore();
const previewSrc = computed(() => {
  if (adminStore.previewPhotoUrl === null || adminStore.previewPhotoUrl === "") {
    return logo
  }
  return adminStore.previewPhotoUrl
})
</script>

<template>
  <div class="flex flex-col justify-center items-center max-h-[75vh]">
    <img :src="previewSrc" class="max-h-[75vh] max-w-full object-contain border-dashed border-4 p-1 border-secondary"
         alt="this literally changes all the time"/>
  </div>
</template>

<style scoped>

</style>