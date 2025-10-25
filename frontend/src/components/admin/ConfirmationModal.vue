<script setup lang="ts">
const emit = defineEmits(['cancel', 'confirm'])
</script>

<template>
  <Teleport to="body">
    <dialog id="are_you_sure" class="modal">
      <div class="modal-box">
        <h3 class="text-lg font-bold text-white">Oi</h3>
        <p class="py-4 text-white">
          <slot/>
        </p>
        <form method="dialog" class="flex flex-row w-full justify-center items-center gap-8">
          <button class="btn btn-primary btn-outline" @click="emit('cancel')">Abort mission</button>
          <button class="btn btn-error btn-outline" @click="emit('confirm')">Yeah, do it</button>
        </form>
      </div>
      <!--  This handles closing the modal by clicking out of the modal -->
      <form method="dialog" class="modal-backdrop" @submit="emit('cancel')">
        <button>close</button>
      </form>
    </dialog>
  </Teleport>
</template>

<style scoped>

</style>