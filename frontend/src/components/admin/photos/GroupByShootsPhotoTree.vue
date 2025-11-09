<script setup>
import {onMounted, ref, watch} from "vue";
import {useAdminStore} from "@/stores/useAdminStore.js";
import PhotoTreeContainer from "@/components/admin/photos/PhotoTreeContainer.vue";
import PhotoTreeEntry from "@/components/admin/photos/PhotoTreeEntry.vue";
import PhotoTree from "@/components/admin/photos/PhotoTree.vue";
import {useAdminRoutes} from "@/composables/useAdminRoutes.js";

const adminStore = useAdminStore();
const shootsWithPhotos = ref([])
const {selector} = useAdminRoutes()
onMounted(() => {
  let loadPromises = []
  loadPromises.push(adminStore.loadShootsFromApi())
  loadPromises.push(adminStore.refreshPreviewPhotosFromApi())
  Promise.all(loadPromises).then(() => {
    mapAndSortShoots()
  })
})
watch(selector, () => {
  mapAndSortShoots()
})

const mapAndSortShoots = () => {
  shootsWithPhotos.value = adminStore.shoots.map(shoot => {
    return {
      id: shoot.id,
      name: shoot.name,
      date: shoot.date,
      photos: adminStore.previewPhotos.filter(photo => photo.shoot_id === shoot.id)
    }
  }).sort((a, b) => {
    if (selector.value === 'shoots-by-date') {
      // Sort by shoot date descending
      return a.date.localeCompare(b.date);
    } else {
      // Sort by shoot name ascending
      return a.name.localeCompare(b.name)
    }
  })
}
</script>

<template>
  <div class="w-full h-full">
    <PhotoTree v-if="shootsWithPhotos.length > 0" class="h-full">
      <PhotoTreeContainer v-for="shoot in shootsWithPhotos" :key="shoot.id"
                          :label="shoot.name">
        <PhotoTreeEntry v-for="photo in shoot.photos" :key="photo.id" :photo="photo"/>
      </PhotoTreeContainer>
    </PhotoTree>
  </div>
</template>