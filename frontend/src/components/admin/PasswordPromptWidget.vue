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
    <div>
      <label class="input input-primary input-ghost validator join-item text-white bg-neutral">
        <input class="w-72 focus:ring-offset-0 focus:ring-0 focus:bg-neutral focus:ring-0 focus:outline-none focus:border-none p-0" type="password" placeholder="64 Characters" required />
      </label>
      <div class="validator-hint hidden">Enter valid email address</div>
    </div>
    <button class="btn btn-primary join-item border-none shadow-none text-white">Authenticate</button>
  </div>
<!--  <div class="join w-auto h-auto mt-64">-->
<!--    <input type="password" class="input input-bordered join-item w-96 focus:outline-none border-primary border-10" placeholder="Password" required v-model="password" />-->
<!--    <button class="btn btn-primary join-item rounded-r-full border-none w-32" @click="handleAuthenticate">-->
<!--      <TheLoader v-if="apiCallInProgress" class="h-6 w-6" />-->
<!--      <span v-else class="text-white font-thin text-lg">-->
<!--        Authenticate-->
<!--      </span>-->
<!--    </button>-->
<!--  </div>-->
</div>
</template>

<style scoped>

</style>