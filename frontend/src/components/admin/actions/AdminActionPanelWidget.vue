<script setup>
import AdminSaveButton from "@/components/admin/actions/AdminSaveButton.vue";
import {useAdminStore} from "@/stores/useAdminStore";
import AdminActionFlowWidget from "@/components/admin/actions/AdminActionFlowWidget.vue";
import PhotoPreviewPanel from "@/components/admin/photos/PhotoPreviewPanel.vue";
import UpdateShootWidget from "@/components/admin/shoots/UpdateShootWidget.vue";
import {computed} from "vue";

const adminStore = useAdminStore();

const showUpdateShootWidget = computed(() => {
  return true
})
</script>

<template>
  <!--Container to split the panel space in half -->
  <div class="w-full flex flex-row flex-nowrap h-full">
<!--    If show edit shoot (entity shoots, update, shoot selected -->
    <UpdateShootWidget v-if="showUpdateShootWidget" class="flex-1"/>
    <PhotoPreviewPanel class="flex-1" />
  </div>
</template>

<style scoped>

</style>