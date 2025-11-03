<script setup>
import insta from "@iconify-icons/mdi/instagram"
import TheFabButton from "@/components/TheFabButton.vue";

const redirectToInstagram = () => {
  // Portrait Pear Instagram link - hard coded as it's not used anywhere else
  window.location.href = "https://www.instagram.com/portraitpear/";
}
</script>


<template>
  <TheFabButton class="hover:text-primary" :icon="insta" @fab-clicked="redirectToInstagram"/>
</template>
