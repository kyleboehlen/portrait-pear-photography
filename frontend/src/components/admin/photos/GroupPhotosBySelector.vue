<script setup lang="ts">
import {onUnmounted, watch} from 'vue'
import {useAdminRoutes} from '@/composables/useAdminRoutes'

const {selector, setSelector} = useAdminRoutes()

// TODO: default to selected shoot scroll to anchor and expanded section
watch(selector, (_, oldValue) => {
  if (oldValue === null || oldValue === '') {
    console.log(selector.value)
  }
})

onUnmounted(() => {
setSelector(null)
})
</script>

<template>

</template>

<style scoped>

</style>