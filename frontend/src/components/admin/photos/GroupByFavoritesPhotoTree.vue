<script setup>
import {onMounted} from "vue";
import {useAdminStore} from "@/stores/useAdminStore.js";
import PhotoTree from "@/components/admin/photos/PhotoTree.vue";
import PhotoTreeEntry from "@/components/admin/photos/PhotoTreeEntry.vue";

const adminStore = useAdminStore();
onMounted(() => {
  adminStore.refreshPreviewPhotosFromApi({'favorites': true})
})
</script>

<template>
  <div class="w-full h-full">
    <PhotoTree class="h-full">
      <PhotoTreeEntry v-for="photo in adminStore.previewPhotos" :key="photo.id" :photo="photo">
      </PhotoTreeEntry>
    </PhotoTree>
  </div>
</template>