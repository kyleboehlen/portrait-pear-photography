<script setup lang="js">
import { useFridayApi } from "@/composables/useFridayApi";
import PasswordPromptWidget from "@/components/admin/PasswordPromptWidget.vue"
import AdminActionFlowWidget from "@/components/admin/AdminActionFlowWidget.vue";

const { adminApiIsAuthenticated } = useFridayApi();
</script>

<template>
  <main class="flex">
    <!-- Show password widget if we don't have an admin token -->
    <PasswordPromptWidget v-if="!adminApiIsAuthenticated" />
    <AdminActionFlowWidget v-else />
  </main>
<!-- If we do have an authZ token, show the admin action widget -->
<!-- Show the shoot edit form if shoot id is selected -->
<!-- Show the full photo tree, or show a mini photo tree based on the shoot photos -->
<!-- Show the hover preview panel, or show the upload widget -->
<!-- Possibly disable this view on a mobile screen -->
</template>

<style scoped>

</style>