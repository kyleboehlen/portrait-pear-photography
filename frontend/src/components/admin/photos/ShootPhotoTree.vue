<script setup>
import PhotoTree from "@/components/admin/photos/PhotoTree.vue";
import {onMounted, watch} from "vue";
import {useAdminStore} from "@/stores/useAdminStore.js";
import PhotoTreeEntry from "@/components/admin/photos/PhotoTreeEntry.vue";

const adminStore = useAdminStore();
onMounted(() => {
  adminStore.refreshPreviewPhotosFromApi({shoot_id: adminStore.selectedShootId});
})
watch(() => adminStore.selectedShootId, (newShootId) => {
  adminStore.refreshPreviewPhotosFromApi({shoot_id: newShootId});
})
</script>

<template>
  <PhotoTree>
    <PhotoTreeEntry v-for="photo in adminStore.previewPhotos" :key="photo.id" :photo="photo"/>
  </PhotoTree>
</template>