<script setup>
import save from "@iconify-icons/material-symbols/save"
import TheFabButton from "@/components/TheFabButton.vue";

const handleSave = () => {
  // TODO
  console.log("Save mentioned lets goooooo")
}
</script>


<template>
  <TheFabButton :icon="save" @fab-clicked="handleSave"/>
</template>
