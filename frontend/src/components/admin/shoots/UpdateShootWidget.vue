<script setup>
import {computed, onMounted, onUnmounted, ref, watch} from 'vue';
import {useAdminStore} from '@/stores/useAdminStore';
import {useAdminRoutes} from '@/composables/useAdminRoutes';
import UploadPhotoWidget from "@/components/admin/photos/UploadPhotoWidget.vue";
import ShootPhotoTree from "@/components/admin/photos/ShootPhotoTree.vue";
import save from "@iconify-icons/material-symbols/save.js";
import TheFabButton from "@/components/TheFabButton.vue";

const adminStore = useAdminStore();

// The only component that needs access to a mutated model of a shoot is this one, therefore we will set the object
// when the component is mounted and clear it when unmounted.
onMounted(() => {
  // Make sure you're mutating a copy - including nested arrays
  adminStore.updateShoot = {
    ...adminStore.selectedShoot,
    default_categories: [...(adminStore.selectedShoot.default_categories || [])]
  }
})
watch(() => adminStore.selectedShoot, (newShoot) => {
  // Deep copy including nested arrays - gotta mutate that copy
  adminStore.updateShoot = {
    ...newShoot,
    default_categories: [...(newShoot.default_categories)]
  }
})
onUnmounted(() => {
  adminStore.updateShoot = {default_categories: []}
})

const displayUrl = computed(() => {
  const baseUrl = window.location.origin;
  return `${baseUrl}/`;
});

const copyText = ref('Copy')
const copyLink = () => {
  navigator.clipboard.writeText(`${displayUrl.value}${adminStore.updateShoot.slug}`);
  copyText.value = 'Copied!'
  setTimeout(() => {
    copyText.value = 'Copy'
  }, 2000);
};

const categories = [
  {id: 1, name: 'Portrait', colorClass: 'badge-primary'},
  {id: 2, name: 'Automotive', colorClass: 'badge-accent'},
  {id: 3, name: 'Street', colorClass: 'badge-warning'},
  {id: 4, name: 'B&W', colorClass: 'badge-neutral'},]

const toggleCategory = (categoryId) => {
  const index = adminStore.updateShoot.default_categories.indexOf(categoryId);
  if (index === -1) {
    adminStore.updateShoot.default_categories.push(categoryId);
  } else {
    adminStore.updateShoot.default_categories.splice(index, 1);
  }
}

const {setEntity, setAction, setSelector} = useAdminRoutes()
const navigateToDeletePhotos = () => {
  setEntity('photos').then(() => {
    setAction('delete').then(() => {
      setSelector('shoots-by-date')
    })
  })
}
const navigateToFavoritePhotos = () => {
  setEntity('photos').then(() => {
    setAction('favorite').then(() => {
      setSelector('shoots-by-date')
    })
  })
}
const navigateToCategorizePhotos = () => {
  setEntity('photos').then(() => {
    setAction('categorize').then(() => {
      setSelector('shoots-by-date')
    })
  })
}

const handleSave = () => {
  adminStore.persistUpdateShoot()
}
</script>

<template>
  <!-- The v-if is only here to prevent the v-model from binding to updateShoot.name before it exists -->
  <div v-if="adminStore.updateShoot" class="flex flex-col items-center">
    <!--  Title (name), and Date -->
    <div class="w-3/4 flex flex-row justify-center items-center mt-6 border-b-2 border-neutral">
      <input type="text"
             class="input input-lg input-ghost input-neutral !focus:outline-none !text-white text-right !p-0 text-3xl"
             v-model="adminStore.updateShoot.name"/>
      <p class="text-white text-2xl mx-6">
        On
      </p>
      <input type="date" class="input input-lg input-ghost input-primary !text-white text-center !p-0 text-3xl"
             v-model="adminStore.updateShoot.date"/>
    </div>

    <!-- Slug and link copy -->
    <label class="input input-ghost w-3/4 pr-0 mt-4">
    <span class="text-neutral text-lg">
      {{ displayUrl }}
    </span>
      <input class="w-full text-lg text-white ml-0 pl-0" type="text" v-model="adminStore.updateShoot.slug"/>
      <button class="btn btn-neutral mr-0 rounded-r-10" @click="copyLink">
        <span class="w-12 h-6">{{ copyText }}</span>
      </button>
    </label>

    <!--  Default category badges -->
    <div class="w-3/4 flex flex-row flex-nowrap mt-4 gap-6">
      <span v-for="category in categories" :key="category.id" class="badge badge-lg hover:cursor-pointer flex-1"
            :class="{
        'badge-dash': adminStore.updateShoot.default_categories.indexOf(category.id) === -1,
        [category.colorClass]: true
      }" @click="toggleCategory(category.id)">
        {{ category.name }}
      </span>
    </div>

    <!--  Manage photos links buttons -->
    <div class="w-3/4 flex flex-row flex-nowrap mt-16 gap-4">
      <button class="btn btn-error btn-outline flex-grow" @click="navigateToDeletePhotos">Delete Photos</button>
      <button class="btn btn-warning btn-outline flex-grow" @click="navigateToFavoritePhotos">Favorite Photos</button>
      <button class="btn btn-accent btn-outline flex-grow" @click="navigateToCategorizePhotos">Categorize Photos
      </button>
    </div>

    <!--  Upload photos control -->
    <UploadPhotoWidget class="w-3/4 mt-4"/>

    <!--  Photo tree for preview -->
    <div class="max-h-[40vh] w-3/4 mt-2 pl-4 pr-16">
      <!--  Let the photo tree re-mount after API calls to refresh state    -->
      <ShootPhotoTree class="h-full w-full"/>
    </div>

    <!--    Save button -->
    <TheFabButton v-if="adminStore.isDirty" class="hover:text-primary" :icon="save" @fab-clicked="handleSave"/>
  </div>
</template>

<style scoped>
/* Override focus styles from daisyUI */
.input:focus, label:focus-within {
  outline: none !important;
  box-shadow: none !important;
  border-color: transparent !important;
}

.badge:hover {
  border: 1px solid !important;
}
</style>