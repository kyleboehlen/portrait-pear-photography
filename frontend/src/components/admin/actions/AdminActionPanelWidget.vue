<script setup>
import AdminSaveButton from "@/components/admin/actions/AdminSaveButton.vue";
import {useAdminStore} from "@/stores/useAdminStore";
import AdminActionFlowWidget from "@/components/admin/actions/AdminActionFlowWidget.vue";
import PhotoPreviewPanel from "@/components/admin/photos/PhotoPreviewPanel.vue";
import UpdateShootWidget from "@/components/admin/shoots/UpdateShootWidget.vue";
import {computed} from "vue";
import {useAdminRoutes} from "@/composables/useAdminRoutes";

const adminStore = useAdminStore();
const adminRoutes = useAdminRoutes()

const showUpdateShootWidget = computed(() => {
  return adminRoutes.entity.value === 'shoots' && adminRoutes.action.value === 'update' && adminStore.selectedShoot.value !== null
})
</script>

<template>
  <!--Container to split the panel space in half -->
  <div class="w-full flex flex-row flex-nowrap h-full">
<!--    If show edit shoot (entity shoots, update, shoot selected -->
    <UpdateShootWidget v-if="showUpdateShootWidget" class="flex-1"/>
    <!--    TODO: Favorites photo tree-->
    <!--    TODO: Delete photo tree-->
    <!--    TODO: Categorize photo tree-->
    <PhotoPreviewPanel class="flex-1" />
    <AdminSaveButton v-if="adminStore.isDirty" />
  </div>
</template>

<style scoped>

</style>