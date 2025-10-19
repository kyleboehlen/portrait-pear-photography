<script setup>
import {onMounted, onUnmounted, watch, ref, computed} from 'vue';
import { useAdminStore } from '@/stores/useAdminStore';
const adminStore = useAdminStore();

// The only component that needs access to a mutated model of a shoot is this one, therefore we will set the object
// when the component is mounted and clear it when unmounted.
onMounted(() => {
  // Make sure you're mutating a copy
  adminStore.updateShoot = {...adminStore.selectedShoot}
})
watch(() => adminStore.selectedShoot, (newShoot) => {
  // Make sure you're mutating a copy
  adminStore.updateShoot = {...newShoot}
})
onUnmounted(() => {
  adminStore.updateShoot = null
})

const displayUrl = computed(() => {
  const baseUrl = window.location.origin;
  return `${baseUrl}/shoots/`;
});

const copyText = ref('Copy')
const copyLink = () => {
  navigator.clipboard.writeText(`${displayUrl.value}${adminStore.updateShoot.slug}`);
  copyText.value = 'Copied!'
  setTimeout(() => {
    copyText.value = 'Copy'
  }, 2000);
};
</script>

<template>
<!-- The v-if is only here to prevent the v-model from binding to updateShoot.name before it exists -->
<div v-if="adminStore.updateShoot" class="border-r-4 border-secondary flex flex-col items-center">
<!--  Title (name), and Date -->
  <div class="w-3/4 flex flex-row justify-center items-center mt-6 border-b-2 border-neutral">
    <input type="text" class="input input-lg input-ghost input-neutral !focus:outline-none !text-white text-right !p-0 text-3xl" v-model="adminStore.updateShoot.name" />
    <p class="text-white text-2xl mx-6">
      On
    </p>
    <input type="date" class="input input-lg input-ghost input-primary !text-white text-center !p-0 text-3xl" v-model="adminStore.updateShoot.date" />
  </div>

<!-- Slug and link copy -->
  <label class="input input-ghost w-3/4 pr-0 mt-4">
    <span class="text-neutral text-lg">
      {{ displayUrl }}
    </span>
    <input class="w-full text-lg text-white ml-0 pl-0" type="text" v-model="adminStore.updateShoot.slug" />
    <button class="btn btn-neutral mr-0 rounded-r-10" @click="copyLink">
      <span class="w-12 h-6">{{ copyText }}</span>
    </button>
  </label>

<!--  Default category badges -->
<!--  Manage photos links buttons -->
<!--  Upload photos control -->
<!--  Photo tree for preview -->
</div>
</template>

<style scoped>
/* Override focus styles from daisyUI */
.input:focus, label:focus-within {
  outline: none !important;
  box-shadow: none !important;
  border-color: transparent !important;
}
</style>