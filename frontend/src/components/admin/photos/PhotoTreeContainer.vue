<script setup>
const props = defineProps(['label'])
</script>

<template>
<details class="w-3/4 collapse bg-base-100 border-y border-r border-base-300 mx-0 rounded-none p-0" name="photo-tree-container">
  <summary class="collapse-title font-semibold text-white/75">{{ props.label }}</summary>
  <div class="collapse-content text-sm pl-0">
    <slot />
  </div>
</details>
</template>