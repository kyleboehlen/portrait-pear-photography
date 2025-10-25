<template>
  <router-view v-slot="{ Component }">
    <component :is="Component" class="grow"/>
  </router-view>
</template>

<script setup>
// Vue
import {RouterView} from "vue-router"
// Components
</script>

<style scoped></style>
