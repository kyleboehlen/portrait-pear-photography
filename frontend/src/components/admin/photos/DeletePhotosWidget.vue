<script setup>
import PhotoTree from "@/components/admin/photos/PhotoTree.vue";
import PhotoTreeEntry from "@/components/admin/photos/PhotoTreeEntry.vue";
import {useAdminStore} from "@/stores/useAdminStore.js";
import TheFabButton from "@/components/TheFabButton.vue";
import {addIcon} from "@iconify/vue";
import trash from "@iconify-icons/mdi/trash-can-outline";

addIcon("trash", trash);

const adminStore = useAdminStore();

const isSelectedForDelete = (photoId) => {
  return adminStore.photosToMutate.includes(photoId);
}
const toggleDeleteMutate = (photoId) => {
  if (isSelectedForDelete(photoId)) {
    adminStore.photosToMutate.splice(adminStore.photosToMutate.findIndex(id => id === photoId), 1);
  } else {
    adminStore.photosToMutate.push(photoId);
  }
}
</script>

<template>
  <div class="max-h-[75vh] w-3/4 mt-2 pl-4 pr-16">
    <PhotoTree class="h-full">
      <PhotoTreeEntry v-for="photo in adminStore.previewPhotos" :key="photo.id" :photo="photo"
                      @click="toggleDeleteMutate(photo.id)">
        <p class="pl-1"
           :class="{
        'text-white': !isSelectedForDelete(photo.id),
        'text-red-600 line-through': isSelectedForDelete(photo.id)
      }">{{ photo.guid }}</p>
      </PhotoTreeEntry>
    </PhotoTree>
    <TheFabButton v-if="adminStore.photosToMutate.length > 0" class="hover:text-error" :icon="trash" @fab-clicked="adminStore.deletePhotos()" />
  </div>
</template>

<style scoped>

</style>