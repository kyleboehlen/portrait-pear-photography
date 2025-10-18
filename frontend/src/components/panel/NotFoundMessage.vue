<template>
  <div class="flex flex-col justify-center items-center">
    <Icon icon="shutter" class="h-48 w-auto text-white/75"/>
    <p class="text-3xl mt-4 text-white/75">No photos found :(</p>
  </div>
</template>

<script setup>
// Icons
import {Icon, addIcon} from "@iconify/vue/offline"
import shutter from "@iconify-icons/material-symbols/camera"

addIcon("shutter", shutter)
</script>
