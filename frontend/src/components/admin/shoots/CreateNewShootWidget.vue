<script setup lang="ts">
import {ref} from "vue";
import {useFridayApi} from '@/composables/useFridayApi'
import {useAdminStore} from '@/stores/useAdminStore'
import {useAdminRoutes} from '@/composables/useAdminRoutes'
import TheLoader from "@/components/TheLoader.vue";

const {apiCallInProgress} = useFridayApi()
const adminStore = useAdminStore()
const {setAction} = useAdminRoutes()
const shootName = ref("");
const handleCreateShoot = () => {
  adminStore.createShoot(shootName.value).then(() => {
    setAction('update')
  });
};
</script>

<template>
  <label class="input input-ghost pr-0 focus-within:ring-2 focus-within:ring-primary ring-2 ring-neutral">
    <input class="text-sm text-white" v-model="shootName" @keyup.enter="handleCreateShoot" type="text"
           placeholder="Shoot name...." required/>
    <button class="btn btn-neutral mr-0 rounded-r-10 flex flex-col justify-center items-center"
            @click="handleCreateShoot">
      <TheLoader v-if="apiCallInProgress" class="w-6 h-6"/>
      <span v-else class="w-auto">Do It</span>
    </button>
  </label>
</template>

<style scoped>

</style>