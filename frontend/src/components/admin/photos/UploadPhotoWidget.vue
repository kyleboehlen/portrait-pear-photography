<script setup>
import {computed, useTemplateRef} from "vue";
import ConfirmationModal from "@/components/admin/ConfirmationModal.vue";
import {useFridayApi} from "@/composables/useFridayApi";
import {useAdminStore} from "@/stores/useAdminStore";

const input = useTemplateRef('photo-upload')
const numFiles = computed(() => {
  return input.value?.files?.length || 0;
})
const clearFiles = () => {
  input.value.value = null
}

const adminStore = useAdminStore();
const {uploadApi} = useFridayApi();
const handlePhotoUpload = () => {
  const files = input.value?.files;
  if (!files || files.length === 0) {
    return;
  }

  const extra = {
    shoot_id: adminStore.selectedShootId,
    // API endpoint requires categories as integers, not strings
    categories: adminStore.selectedShoot.default_categories.map(cat => parseInt(cat, 10))
  }

  for (let i = 0; i < files.length; i++) {
    uploadApi('/admin/upload-photo', files[i], extra, adminStore.bearerToken);
  }

  clearFiles();
}
</script>

<template>
<span>
<!--  Don't allow uploading files if the shoot is in a dirty state -->
    <input ref="photo-upload" type="file" class="file-input file-input-primary text-white w-full" accept=".jpg, .jpeg"
           multiple onchange="are_you_sure.showModal()" :disabled="adminStore.isDirty"/>

  <ConfirmationModal @cancel="clearFiles" @confirm="handlePhotoUpload">
    Are you sure you want to upload the {{ numFiles }} photos??
  </ConfirmationModal>
</span>
</template>

<style scoped>

</style>