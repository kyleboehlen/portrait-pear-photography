<script setup>
import {onMounted} from "vue";
import {useAdminStore} from "@/stores/useAdminStore.js";
// TODO: scroll down to the selected shoot, and open it by default, if there is a selected shoot.
// TODO: then clear it?

const adminStore = useAdminStore();
onMounted(() => {
  adminStore.loadShootsFromApi()
  adminStore.refreshPreviewPhotosFromApi()
})

// TODO: computed to link photos to shoots
// TODO: computed function to return linked photos array by shoot name order
// TODO: same, but for date desc

</script>

<template>

</template>

<style scoped>

</style>