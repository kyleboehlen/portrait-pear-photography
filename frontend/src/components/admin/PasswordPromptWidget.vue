<script setup lang="js">
import { ref } from 'vue'
const password = ref('')
const handleAuthenticate = () => {
  // Handle authentication logic here
  console.log(password.value)
};
</script>

<template>
<div class="w-full flex flex-row justify-center">
  <div class="join w-auto h-auto mt-64">
    <input type="password" class="input input-bordered join-item w-96 focus:outline-none border-primary border-10" placeholder="Password" required v-model="password" />
    <button class="btn btn-primary join-item rounded-r-full border-none" @click="handleAuthenticate">Authenticate</button>
  </div>
</div>
</template>

<style scoped>

</style>