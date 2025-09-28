<template>
  <button class="text-primary hover:active pt-1 hover:pt-0">
    <Icon :icon="props.icon" class="text-primary h-2/3 w-auto"></Icon>
    <span class="btm-nav-label text-slate-300">{{ props.label }}</span>
  </button>
</template>

<script setup>
import { Icon } from "@iconify/vue/offline"
const props = defineProps(["icon", "label"])
</script>
