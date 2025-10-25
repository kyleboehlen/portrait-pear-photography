<script setup>
import {addIcon, Icon} from "@iconify/vue/offline";

const props = defineProps(['icon'])
addIcon("fab", props.icon)

const emit = defineEmits(['fab-clicked']);
const handleClick = () => {
  emit('fab-clicked');
}
</script>

<template>
  <span class="fixed right-0 bottom-0 pb-3 pr-3">
    <button
        class="w-18 h-18 bg-neutral btn-circle flex flex-row justify-center items-center p-2 hover:text-primary text-white/75 hover:cursor-pointer"
        @click="handleClick">
      <Icon icon="fab" class=" w-full h-auto"/>
    </button>
  </span>
</template>

<style scoped>

</style>