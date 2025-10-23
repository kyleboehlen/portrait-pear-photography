<script setup lang="ts">
import {useTemplateRef} from "vue";

const input = useTemplateRef('database-upload')
const clearFile = () => {
  input.value.value = null
}
const handleDbUpload = () => {
  // TODO
  console.log('Upload db here')
}
</script>

<template>
  <input ref="database-upload" type="file" class="file-input file-input-primary w-full text-white" accept=".db" onchange="are_you_sure.showModal()" />

  <Teleport to="body">
    <dialog id="are_you_sure" class="modal">
      <div class="modal-box">
        <h3 class="text-lg font-bold text-white">Oi</h3>
        <p class="py-4 text-white">Are you sure you want to upload a new database file??</p>
        <form method="dialog" class="flex flex-row w-full justify-center items-center gap-8">
          <button class="btn btn-primary btn-outline" @click="clearFile">Abort mission</button>
          <button class="btn btn-error btn-outline" @click="handleDbUpload">Yeah, do it</button>
        </form>
      </div>
      <!--  This handles closing the modal by clicking out of the modal -->
      <form method="dialog" class="modal-backdrop" @submit="clearFile">
        <button>close</button>
      </form>
    </dialog>
  </Teleport>
</template>

<style scoped>

</style>