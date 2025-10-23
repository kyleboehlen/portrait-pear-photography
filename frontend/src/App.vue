<template>
  <router-view v-slot="{ Component }">
    <component :is="Component" class="grow"/>
  </router-view>
</template>

<script setup>
// Vue
import {RouterView} from "vue-router"
// Components
import TheHeader from "@/components/TheHeader.vue"
import TheFooter from "@/components/TheFooter.vue"
import TheInstagramButton from "@/components/TheInstagramButton.vue";
</script>

<style scoped></style>
