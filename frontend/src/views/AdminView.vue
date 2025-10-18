<script setup>
import {useFridayApi} from "@/composables/useFridayApi";
import PasswordPromptWidget from "@/components/admin/PasswordPromptWidget.vue"
import AdminActionFlowWidget from "@/components/admin/actions/AdminActionFlowWidget.vue";
import AdminActionPanelWidget from "@/components/admin/actions/AdminActionPanelWidget.vue";

const {adminApiIsAuthenticated} = useFridayApi();
</script>

<template>
  <main class="flex flex-col justify-center items-start">
    <!-- Show password widget if we don't have an admin token -->
    <PasswordPromptWidget v-if="!adminApiIsAuthenticated" />
    <AdminActionFlowWidget v-if="adminApiIsAuthenticated" />
    <AdminActionPanelWidget v-if="adminApiIsAuthenticated" class="flex-1" />
  </main>
</template>

<style scoped>

</style>