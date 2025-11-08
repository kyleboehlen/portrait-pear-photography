<script setup>
import {ref, useTemplateRef} from "vue";
import ConfirmationModal from "@/components/admin/ConfirmationModal.vue";
import {useFridayApi} from "@/composables/useFridayApi";
import {useAdminStore} from "@/stores/useAdminStore";

const input = useTemplateRef('photo-upload')
const clearFiles = () => {
  input.value.value = null
}

const numFiles = ref(0)
const updateNumFiles = () => {
  numFiles.value = input.value?.files.length ?? 0;
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
    uploadApi('/admin/upload-photo', files[i], extra, adminStore.bearerToken).then((photo) => {
      if (photo !== false) {
        adminStore.previewPhotos.push(photo);
      }
    });
  }

  clearFiles();
}
</script>

<template>
<span>
<!--  Don't allow uploading files if the shoot is in a dirty state -->
    <input ref="photo-upload" type="file" class="file-input file-input-primary text-white w-full" accept=".jpg, .jpeg"
           multiple @change="updateNumFiles" onchange="are_you_sure.showModal()" :disabled="adminStore.isDirty"/>

  <ConfirmationModal @cancel="clearFiles" @confirm="handlePhotoUpload">
    Are you sure you want to upload {{ numFiles }} photos??
  </ConfirmationModal>
</span>
</template>

<style scoped>

</style>