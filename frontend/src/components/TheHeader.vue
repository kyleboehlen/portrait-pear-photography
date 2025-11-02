<script setup>
import logo from "@/assets/imgs/green-camera-logo.png?url"
import {onMounted, ref} from "vue"
import {useRouter} from "vue-router"
import {addIcon, Icon} from "@iconify/vue/offline"
import hamburger from "@iconify-icons/pajamas/hamburger"

addIcon("hamburger", hamburger)

// Shrink on scroll
const shrink = ref(false)
const shrinkOnScroll = () => {
  if (window.scrollY > 50) {
    shrink.value = true
  } else {
    shrink.value = false
  }
}

// Mounted
onMounted(() => {
  // Mount on scroll for shrink
  window.addEventListener("scroll", shrinkOnScroll)
})

// Home link
const router = useRouter()
const navigateHome = () => {
  router.push({name: "home"})
}
</script>

<template>
  <!-- Header -->
  <header class="w-full sticky top-0">
    <!-- Header proper -->
    <div
        class="w-full px-2 bg-secondary flex flex-row flex-nowrap justify-between items-center rounded-b-xl"
        :class="{ 'h-28': !shrink, 'md:h-24': !shrink, 'h-12': shrink }">
      <img :src="logo" class="h-full xs:h-full cursor-pointer" @click="navigateHome"/>
      <h1
          class="md:whitespace-nowrap text-center uppercase text-white/75"
          :class="{ 'text-4xl': !shrink, 'sm:text-5xl': !shrink, 'md:text-6xl': !shrink, 'text-3xl': shrink }">
        Portrait Pear
      </h1>
      <slot name="action">
        <div></div>
      </slot>
    </div>
  </header>
</template>

<style scoped>
h1 {
  font-family: Kodchasan, sans-serif;
  transition: 0.2s;
}

header {
  transition: 0.2s;
}
</style>
