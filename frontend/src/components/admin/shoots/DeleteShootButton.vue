<script setup lang="ts">
import {useAdminStore} from '@/stores/useAdminStore'

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
  <Teleport to="body">
    <dialog id="are_you_sure" class="modal">
      <div class="modal-box">
        <h3 class="text-lg font-bold text-white">Oi</h3>
        <p class="py-4 text-white">Are you sure you want to delete shoot "{{ adminStore.selectedShoot?.name }}"??</p>
        <form method="dialog" class="flex flex-row w-full justify-center items-center gap-8">
          <button class="btn btn-primary btn-outline">Abort mission</button>
          <button class="btn btn-error btn-outline" @click="deleteShoot">Yeah, do it</button>
        </form>
      </div>
      <!--  This handles closing the modal by clicking out of the modal -->
      <form method="dialog" class="modal-backdrop">
        <button>close</button>
      </form>
    </dialog>
  </Teleport>
</span>
</template>

<style scoped>

</style>