<template>
  <!-- Favorite heart button -->
  <span class="fixed right-0 bottom-0 pb-4 pr-4">
    <button class="btn btn-lg btn-neutral btn-circle" @click="redirectToInstagram">
      <Icon icon="insta" class="text-primary w-2/3 h-auto" />
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
