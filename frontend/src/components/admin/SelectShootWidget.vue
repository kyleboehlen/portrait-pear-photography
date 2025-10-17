<script setup lang="ts">
import { useAdminStore } from '@/stores/useAdminStore'
// import { useFridayApi } from '@/composables/useFridayApi'; // TODO: see below
import { onMounted } from "vue";
// import TheLoader from "@/components/TheLoader.vue"; // TODO: see below

// const { apiCallInProgress } = useFridayApi() // TODO: see below

const adminStore = useAdminStore()
onMounted(() => {
  adminStore.loadShootsFromApi()
})
</script>

<template>
<!--  TODO: show loader instead of select on api call in progress?-->
  <select
      v-model.number="adminStore.selectedShootId"
      class="select select-ghost select-md focus:outline-none focus:ring-0 !text-white"
      :class="{
          '!border-primary': adminStore.selectedShootId != 0,
          '!border-neutral': adminStore.selectedShootId == 0
        }">
    <option :value="0" disabled>Select a Shoot</option>
    <option v-for="shoot in adminStore.shoots" :key="shoot.id" :value="shoot.id">{{ shoot.name }}</option>
  </select>
</template>

<style scoped>

</style>