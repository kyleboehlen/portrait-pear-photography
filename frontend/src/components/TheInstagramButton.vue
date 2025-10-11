<template>
  <!-- Favorite heart button -->
  <span class="fixed right-0 bottom-0 pb-3 pr-3">
    <button class="w-18 h-18 bg-neutral btn-circle flex flex-row justify-center items-center p-2 hover:text-primary text-white/75 hover:cursor-pointer" @click="redirectToInstagram">
      <Icon icon="insta" class=" w-full h-auto" />
    </button>
  </span>
</template>

<script setup>
// Icons
import { Icon, addIcon } from "@iconify/vue/offline"
import insta from "@iconify-icons/mdi/instagram"

addIcon("insta", insta)

const redirectToInstagram = () => {
  // Portrait Pear Instagram link - hard coded as it's not used anywhere else
  window.location.href = "https://www.instagram.com/portraitpear/";
}
</script>
