<script setup>
import {useAdminStore} from "@/stores/useAdminStore.js";
import {onMounted, ref} from "vue";
import {filterCategories} from "@/constants/categories.js";
import PhotoTree from "@/components/admin/photos/PhotoTree.vue";
import PhotoTreeContainer from "@/components/admin/photos/PhotoTreeContainer.vue";
import PhotoTreeEntry from "@/components/admin/photos/PhotoTreeEntry.vue";

const adminStore = useAdminStore();
const categoriesWithPhotos = ref([])
onMounted(() => {
  adminStore.refreshPreviewPhotosFromApi().then(() => {
    categoriesWithPhotos.value = filterCategories.map(category => {
      return {
        id: category.id,
        name: category.name,
        photos: adminStore.previewPhotos.filter(photo => photo.categories.includes(category.id))
      }
    }).filter(category => category.photos.length > 0)
  })
})
</script>

<template>
  <div class="w-full h-full">
    <PhotoTree v-if="categoriesWithPhotos.length > 0" class="h-full">
      <PhotoTreeContainer v-for="category in categoriesWithPhotos" :key="category.id"
                          :label="category.name">
        <PhotoTreeEntry v-for="photo in category.photos" :key="photo.id" :photo="photo"/>
      </PhotoTreeContainer>
    </PhotoTree>
  </div>
</template>