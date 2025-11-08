<script setup>
import {onUnmounted, watch} from 'vue'
import {useAdminRoutes} from '@/composables/useAdminRoutes'

const {selector, setSelector} = useAdminRoutes()
onUnmounted(() => {
  setSelector(null)
})
</script>

<template>
  <select
      :value="selector"
      class="select select-ghost select-md focus:outline-none focus:ring-0 !text-white"
      :class="{
          '!border-primary': !!selector,
          '!border-neutral': !selector
        }"
      @change="setSelector($event.target.value)">

    <option value="" disabled>Group By</option>
    <option value="shoots">Shoots</option>
    <option value="shoots-by-date">Date</option>
    <option value="categories">Categories</option>
    <option value="favorites">Favorites</option>
  </select>
</template>

<style scoped>

</style>