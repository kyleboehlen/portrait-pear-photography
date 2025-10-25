<script setup>
import {useFridayApi} from "@/composables/useFridayApi";
import PasswordPromptWidget from "@/components/admin/PasswordPromptWidget.vue"
import AdminActionFlowWidget from "@/components/admin/actions/AdminActionFlowWidget.vue";
import AdminActionPanelWidget from "@/components/admin/actions/AdminActionPanelWidget.vue";
import TheHeader from "@/components/TheHeader.vue";
import TheFooter from "@/components/TheFooter.vue";
import LogoutButton from "@/components/admin/actions/LogoutButton.vue";

const {adminApiIsAuthenticated} = useFridayApi();
</script>

<template>
  <div class="app-wrapper flex flex-col min-h-screen justify-start">
    <TheHeader>
      <template v-if="adminApiIsAuthenticated" v-slot:action>
        <LogoutButton/>
      </template>
    </TheHeader>

    <main class="flex flex-col justify-center items-start grow">
      <!-- Show password widget if we don't have an admin token -->
      <PasswordPromptWidget v-if="!adminApiIsAuthenticated"/>
      <AdminActionFlowWidget v-if="adminApiIsAuthenticated"/>
      <AdminActionPanelWidget v-if="adminApiIsAuthenticated" class="flex-1"/>
    </main>

    <TheFooter class="justify-self-end"/>
  </div>
</template>

<style scoped>

</style>