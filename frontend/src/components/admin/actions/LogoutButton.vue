<script setup>
import {useRouter} from "vue-router";
import {useAdminStore} from "@/stores/useAdminStore";

const router = useRouter();
const adminStore = useAdminStore();

// No need to call the back end, just clear the state which includes the bearer token
const handleLogout = () => {
  adminStore.$reset()
  router.push('/admin').then(() => {
    // The application state doesn't know it's unauthenticated until we test an endpoint that requires a JWT
    adminStore.testToken()
  })
}
</script>

<template>
  <button class="btn btn-outline btn-error" @click="handleLogout">Log out</button>
</template>

<style scoped>

</style>