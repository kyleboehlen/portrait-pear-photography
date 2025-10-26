<script setup>
import {useAdminStore} from '@/stores/useAdminStore'
import ConfirmationModal from "@/components/admin/ConfirmationModal.vue";

const adminStore = useAdminStore()

const deleteShoot = () => {
  adminStore.deleteSelectedShoot()
}
</script>

<template>
<span>
  <button class="btn btn-outline btn-error w-full whitespace-nowrap" onclick="are_you_sure.showModal()">
    Delete Shoot
  </button>

  <ConfirmationModal @confirm="deleteShoot">
    Are you sure you want to delete shoot "{{ adminStore.selectedShoot?.name }}"??
  </ConfirmationModal>
</span>
</template>

<style scoped>

</style>