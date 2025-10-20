<script setup>
import AdminSaveButton from "@/components/admin/actions/AdminSaveButton.vue";
import {useAdminStore} from "@/stores/useAdminStore";
import PhotoPreviewPanel from "@/components/admin/photos/PhotoPreviewPanel.vue";
import UpdateShootWidget from "@/components/admin/shoots/UpdateShootWidget.vue";
import {computed} from "vue";
import {useAdminRoutes} from "@/composables/useAdminRoutes";

const adminStore = useAdminStore();
const adminRoutes = useAdminRoutes()

const showUpdateShootWidget = computed(() => {
  return adminRoutes.entity.value === 'shoots' && adminRoutes.action.value === 'update' && adminStore.selectedShootId !== 0
})
</script>

<template>
  <!--Container to split the panel space in half -->
  <div class="w-full flex flex-row flex-nowrap h-full mt-4">
<!--    If show edit shoot (entity shoots, update, shoot selected -->
    <UpdateShootWidget v-if="showUpdateShootWidget" class="flex-1"/>
    <!--    TODO: Favorites photo tree-->
    <!--    TODO: Delete photo tree-->
    <!--    TODO: Categorize photo tree-->
    <PhotoPreviewPanel v-if="showUpdateShootWidget" class="p-4" />
    <AdminSaveButton v-if="adminStore.isDirty && showUpdateShootWidget" />
  </div>
</template>

<style scoped>

</style>