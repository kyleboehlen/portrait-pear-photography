<script setup>
import trash from "@iconify-icons/mdi/trash-can-outline.js";
import heart from "@iconify-icons/mdi/cards-heart-outline"
import save from "@iconify-icons/material-symbols/save.js";
import {addIcon} from "@iconify/vue";
import TheFabButton from "@/components/TheFabButton.vue";
import {computed} from "vue";
import {useAdminRoutes} from "@/composables/useAdminRoutes.js";
import {useAdminStore} from "@/stores/useAdminStore.js";

const {action} = useAdminRoutes();

addIcon("trash", trash);
addIcon("save", save);
addIcon("heart", heart);

const iconToShow = computed(() => {
switch (action.value) {
  case "favorite":
    return heart;
  case "delete":
    return trash;
  default:
    return save;
}
})

const adminStore = useAdminStore();
const emit = defineEmits(['mutation-complete'])
const handleMutatePhotos = () => {
  switch (action.value) {
    case "favorite":
      console.log('todo favorite mutation')
      break;
    case "delete":
      Promise.all(adminStore.deletePhotos()).then(() => {
        emit('mutation-complete')
      });
      break;
    case "categorize":
      console.log('todo categorize mutation')
      break;
  }
}
</script>

<template>
  <TheFabButton v-if="adminStore.photosToMutate.length > 0" class="hover:text-error" :icon="iconToShow" @fab-clicked="handleMutatePhotos" />
</template>