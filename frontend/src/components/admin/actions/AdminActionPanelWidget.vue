<script setup>
import {useAdminStore} from "@/stores/useAdminStore";
import PhotoPreviewPanel from "@/components/admin/photos/PhotoPreviewPanel.vue";
import UpdateShootWidget from "@/components/admin/shoots/UpdateShootWidget.vue";
import {computed} from "vue";
import {useAdminRoutes} from "@/composables/useAdminRoutes";
import DeletePhotosWidget from "@/components/admin/photos/DeletePhotosWidget.vue";

const adminStore = useAdminStore();
const adminRoutes = useAdminRoutes()

const showUpdateShootWidget = computed(() => {
  return adminRoutes.entity.value === 'shoots' && adminRoutes.action.value === 'update' && adminStore.selectedShootId !== 0
})

const showDeleteShootWidget = computed(() => {
  return adminRoutes.entity.value === 'photos' && adminRoutes.action.value === 'delete'
})
</script>

<template>
  <!--Container to split the panel space in half -->
  <div class="w-full flex flex-row flex-nowrap h-full mt-4">
    <!--    If show edit shoot (entity shoots, update, shoot selected -->
    <UpdateShootWidget v-if="showUpdateShootWidget" class="flex-1"/>
    <!--    TODO: Favorites photo tree-->
    <!--    TODO: Delete photo tree-->
    <DeletePhotosWidget v-if="showDeleteShootWidget" class="flex-1"/>
    <!--    TODO: Categorize photo tree-->
    <PhotoPreviewPanel v-if="showUpdateShootWidget || showDeleteShootWidget" class="p-4 flex-1"/>
  </div>
</template>

<style scoped>

</style>