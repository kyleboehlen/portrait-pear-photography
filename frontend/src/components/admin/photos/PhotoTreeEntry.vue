<script setup>
import {useAdminStore} from "@/stores/useAdminStore.js";
import {usePhotoUtils} from "@/composables/usePhotoUtils.js";
import {useAdminRoutes} from "@/composables/useAdminRoutes.js";
import {computed} from "vue";

const props = defineProps(['photo'])

const adminStore = useAdminStore();
const isSelectedForDelete = computed(() => {
  return adminStore.photosToMutate.includes(props.photo.id);
})

const {action} = useAdminRoutes();

const {gridUrl} = usePhotoUtils();
const handleMouseEnter = () => {
  adminStore.previewPhotoUrl = gridUrl(props.photo.guid);
}

const handleMouseLeave = () => {
  adminStore.previewPhotoUrl = '';
}

const handleEntryClick = () => {
  if (action.value === 'delete') {
    if (isSelectedForDelete.value) {
      adminStore.photosToMutate.splice(adminStore.photosToMutate.findIndex(id => id === props.photo.id), 1);
    } else {
      adminStore.photosToMutate.push(props.photo.id);
    }
  }
}
</script>

<template>
  <div class="h-10 w-full flex flex-row justify-start items-center hover:bg-neutral" @mouseenter="handleMouseEnter"
       @mouseleave="handleMouseLeave" @click="handleEntryClick">
    <div class="border-b-2 border-white w-1/10"></div>
    <p class="text-white ml-2">{{ props.photo.id }} - </p>
    <slot>
      <p class="pl-1"
         :class="{
        'text-white': !isSelectedForDelete,
        'text-red-600 line-through': isSelectedForDelete
      }">{{ photo.guid }}</p>
<!--      Show faves heart if action is favorite - mutate the image in preview images -->
<!--      Show categories to toggle on hover if action is categorize - mutate the image in preview images -->
    </slot>
  </div>
</template>

<style scoped>

</style>