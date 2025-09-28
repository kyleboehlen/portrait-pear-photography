<template>
  <div class="app-wrapper flex flex-col min-h-screen justify-start">
    <TheHeader />

    <router-view v-slot="{ Component }">
      <transition name="fade">
        <component :is="Component" class="grow" />
      </transition>
    </router-view>

    <!-- <RouterView class="grow" /> -->
    <TheFooter class="justify-self-end" />
  </div>
</template>

<script setup>
// Vue
import { onMounted } from "vue"
import { RouterView } from "vue-router"
// Components
import TheHeader from "@/components/TheHeader.vue"
import TheFooter from "@/components/TheFooter.vue"
// Capacitor
import { SplashScreen } from "@capacitor/splash-screen"

onMounted(() => {
  setTimeout(() => {
    SplashScreen.hide()
  }, 750)
})
</script>

<style scoped>
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
