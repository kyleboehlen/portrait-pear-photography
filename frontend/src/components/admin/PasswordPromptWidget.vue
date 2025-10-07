<script setup lang="js">
import { ref } from 'vue'
import { useAdminStore } from '@/stores/useAdminStore'
import { useFridayApi } from '@/composables/useFridayApi'
import TheLoader from '@/components/TheLoader.vue'

const { apiCallInProgress } = useFridayApi()

const password = ref('')
const adminStore = useAdminStore()
const handleAuthenticate = () => {
  adminStore.authenticate(password.value)
};
</script>

<template>
<div class="w-full flex flex-row justify-center">
  <div class="join w-auto h-auto mt-64">
    <input type="password" class="input input-bordered join-item w-96 focus:outline-none border-primary border-10" placeholder="Password" required v-model="password" />
    <button class="btn btn-primary join-item rounded-r-full border-none w-32" @click="handleAuthenticate">
      <TheLoader v-if="apiCallInProgress" class="h-6 w-6" />
      <span v-else class="text-white font-thin text-lg">
        Authenticate
      </span>
    </button>
  </div>
</div>
</template>

<style scoped>

</style>