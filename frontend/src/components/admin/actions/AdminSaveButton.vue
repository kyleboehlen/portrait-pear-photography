<script setup>
import save from "@iconify-icons/material-symbols/save"
import TheFabButton from "@/components/TheFabButton.vue";
import {useAdminStore} from "@/stores/useAdminStore.js";

const adminStore = useAdminStore();
const handleSave = () => {
  adminStore.persistUpdateShoot()
}
</script>


<template>
  <TheFabButton :icon="save" @fab-clicked="handleSave"/>
</template>
