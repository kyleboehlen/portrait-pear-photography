<script setup>
import {useAdminStore} from "@/stores/useAdminStore.js";

const props = defineProps(['url'])
const adminStore = useAdminStore();

const handleMouseEnter = () => {
  adminStore.previewPhotoUrl = props.url;
}

const handleMouseLeave = () => {
  adminStore.previewPhotoUrl = '';
}
</script>

<template>
  <div class="h-10 w-full flex flex-row justify-start items-center" @mouseenter="handleMouseEnter"
       @mouseleave="handleMouseLeave">
    <div class="border-b-2 border-white w-1/10"></div>
    <p class="text-white ml-2">Photo Tree Entry</p>
  </div>
</template>

<style scoped>

</style>