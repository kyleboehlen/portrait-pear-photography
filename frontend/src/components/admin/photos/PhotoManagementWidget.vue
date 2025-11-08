<script setup>
import MutatePhotosButton from "@/components/admin/photos/MutatePhotosButton.vue";
import {useAdminRoutes} from "@/composables/useAdminRoutes";
import GroupByShootsPhotoTree from "@/components/admin/photos/GroupByShootsPhotoTree.vue";
import GroupByCategoryPhotoTree from "@/components/admin/photos/GroupByCategoryPhotoTree.vue";
import GroupByFavoritesPhotoTree from "@/components/admin/photos/GroupByFavoritesPhotoTree.vue";
import {computed, ref} from "vue";

const {selector} = useAdminRoutes()

const keyIncrementor = ref(0)
const shootsKey = computed(() => `shoots-${keyIncrementor.value}`)
const categoryKey = computed(() => `category-${keyIncrementor.value}`)
const favoritesKey = computed(() => `favorites-${keyIncrementor.value}`)
</script>

<template>
<span>
  <GroupByShootsPhotoTree v-if="selector === 'shoots' || selector === 'shoots-by-date'" :key="shootsKey"/>
  <GroupByCategoryPhotoTree v-else-if="selector === 'category'" :key="categoryKey"/>
  <GroupByFavoritesPhotoTree v-else-if="selector === 'favorites'" :key="favoritesKey"/>

  <MutatePhotosButton @mutation-complete="keyIncrementor++"/>
</span>
</template>

<style scoped>

</style>