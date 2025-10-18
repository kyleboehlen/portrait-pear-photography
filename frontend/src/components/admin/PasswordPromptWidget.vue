<script setup lang="js">
import {onMounted, ref} from 'vue'
import {useAdminStore} from '@/stores/useAdminStore'
import {useFridayApi} from '@/composables/useFridayApi'
import TheLoader from '@/components/TheLoader.vue'

const {apiCallInProgress} = useFridayApi()

const password = ref('')
const adminStore = useAdminStore()

// It is the job of the password widget to determine whether or not to show the password prompt, therefore it needs to
// be responsible for checking whether the current token is valid. If an authenticated call fails or succeeds the api
// composable will update its state accordingly.
onMounted(() => {
  adminStore.testToken()
});

const handleAuthenticate = () => {
  adminStore.authenticate(password.value)
};
</script>

<template>
  <div class="w-full flex flex-row justify-center items-center">
    <div class="w-1/2 h-auto flex flex-col justify-center items-center">
      <label class="input input-ghost w-1/2 pr-0 focus-within:ring-2 focus-within:ring-primary">
        <input class="w-full text-xs text-white" v-model="password" @keyup.enter="handleAuthenticate" type="password"
               placeholder="x" min="1" max="72" required/>
        <button class="btn btn-neutral mr-0 rounded-r-10" @click="handleAuthenticate">
          <TheLoader v-if="apiCallInProgress" class="w-6 h-6"/>
          <span v-else class="w-6 h-6">Go</span>
        </button>
      </label>
    </div>
  </div>
</template>

<style scoped>

</style>