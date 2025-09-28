<template>
  <div class="flex flex-col justify-center items-center">
    <Icon icon="shutter" class="h-48 w-auto" />
    <p class="text-3xl mt-4">{{ props.msg }}</p>
  </div>
</template>

<script setup>
// Icons
import { Icon, addIcon } from "@iconify/vue/offline"
import shutter from "@iconify-icons/material-symbols/camera"

addIcon("shutter", shutter)

const props = defineProps(["msg"])
</script>
