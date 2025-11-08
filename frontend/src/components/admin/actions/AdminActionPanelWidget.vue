<script setup>
import {useAdminStore} from "@/stores/useAdminStore";
import PhotoPreviewPanel from "@/components/admin/photos/PhotoPreviewPanel.vue";
import UpdateShootWidget from "@/components/admin/shoots/UpdateShootWidget.vue";
import {computed} from "vue";
import {useAdminRoutes} from "@/composables/useAdminRoutes";
import PhotoManagementWidget from "@/components/admin/photos/PhotoManagementWidget.vue";

const adminStore = useAdminStore();
const {entity, action, selector} = useAdminRoutes()

const showUpdateShootWidget = computed(() => {
  return entity.value === 'shoots' && action.value === 'update' && adminStore.selectedShootId !== 0
})

const showPhotoManagementWidget = computed(() => {
  return entity.value === 'photos' && !!selector.value
})
</script>

<template>
  <!--Container to split the panel space in half -->
  <div class="w-full flex flex-row flex-nowrap h-full mt-4">
    <!--    If show edit shoot (entity shoots, update, shoot selected -->
    <UpdateShootWidget v-if="showUpdateShootWidget" class="flex-1"/>

    <PhotoManagementWidget v-else-if="showPhotoManagementWidget" class="max-h-[75vh] w-3/4 mt-2 pl-4 pr-16 flex-1" />

    <PhotoPreviewPanel v-if="showUpdateShootWidget || showPhotoManagementWidget" class="p-4 flex-1"/>
  </div>
</template>

<style scoped>

</style>