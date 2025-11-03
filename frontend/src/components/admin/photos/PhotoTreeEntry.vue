<script setup>
import {useAdminStore} from "@/stores/useAdminStore.js";
import {usePhotoUtils} from "@/composables/usePhotoUtils.js";

const props = defineProps(['photo'])
const adminStore = useAdminStore();
const {gridUrl} = usePhotoUtils();

const handleMouseEnter = () => {
  adminStore.previewPhotoUrl = gridUrl(props.photo.guid);
}

const handleMouseLeave = () => {
  adminStore.previewPhotoUrl = '';
}
</script>

<template>
  <div class="h-10 w-full flex flex-row justify-start items-center hover:bg-neutral" @mouseenter="handleMouseEnter"
       @mouseleave="handleMouseLeave">
    <div class="border-b-2 border-white w-1/10"></div>
    <p class="text-white ml-2">{{ props.photo.id }} - </p>
    <slot />
  </div>
</template>

<style scoped>

</style>