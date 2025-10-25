<script setup>
import {useFridayApi} from '@/composables/useFridayApi';
import {useAdminStore} from '@/stores/useAdminStore';

const {downloadApi} = useFridayApi();
const adminStore = useAdminStore();
const handleDownloadDb = () => {
  downloadApi('/admin/export-database', 'friday.db', adminStore.bearerToken);
}
</script>

<template>
  <button class="btn btn-primary m-0" @click="handleDownloadDb">
    Download DB File
  </button>
</template>

<style scoped>

</style>