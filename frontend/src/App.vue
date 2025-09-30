<template>
  <div class="app-wrapper flex flex-col min-h-screen justify-start">
    <TheHeader />

    <router-view v-slot="{ Component }">
      <transition name="fade">
        <component :is="Component" class="grow" />
      </transition>
    </router-view>

    <TheFooter class="justify-self-end" />
  </div>
</template>

<script setup>
// Vue
import { RouterView } from "vue-router"
// Components
import TheHeader from "@/components/TheHeader.vue"
import TheFooter from "@/components/TheFooter.vue"
</script>

<style scoped>
/* .fade classes might show as unused by your linter, but Vue's named transitions expect to see these classes */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease;
}

.fade-enter-active {
  transition-delay: 0.5s;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
