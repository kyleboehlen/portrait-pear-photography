<script setup lang="ts">
import {useTemplateRef} from "vue";
import ConfirmationModal from "@/components/admin/ConfirmationModal.vue";
import {useFridayApi} from "@/composables/useFridayApi";
import {useAdminStore} from "@/stores/useAdminStore";

const input = useTemplateRef('database-upload')
const clearFile = () => {
  input.value.value = null
}

const adminStore = useAdminStore();
const {uploadApi} = useFridayApi();
const handleDbUpload = () => {
  const file = input.value?.files?.[0];
  if (!file) {
    return;
  }

  uploadApi('/admin/import-database', file, adminStore.bearerToken);
  clearFile();
}
</script>

<template>
  <input ref="database-upload" type="file" class="file-input file-input-primary w-full text-white" accept=".db"
         onchange="are_you_sure.showModal()"/>

  <ConfirmationModal @cancel="clearFile" @confirm="handleDbUpload">
    Are you sure you want to upload a new database file??
  </ConfirmationModal>
</template>

<style scoped>

</style>