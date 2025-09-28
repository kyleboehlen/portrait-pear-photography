<template>
  <main class="flex flex-col justify-center items-center pb-8">
    <ul v-if="favorites.shoots.length > 0" class="menu bg-secondary w-5/6 xs:w-3/4 sm:w-1/2 p-2 rounded-box text-lg">
      <li class="menu-title">
        <span class="!text-base">Favorite Shoots</span>
      </li>
      <li v-for="shoot in favorites.shoots" :key="shoot.slug">
        <a class="link" @click="navigate(shoot.slug)">
          <Icon icon="shutter" class="inline-block mt-1" />
          {{ shoot.name }}
        </a>
      </li>
    </ul>
    <NotFoundMessage v-if="favorites.shoots.length === 0" msg="No Shoots Favorited :(" />
  </main>
</template>

<script setup>
// Vue
import { onMounted } from "vue"
import { useRouter } from "vue-router"
// Icons
import { Icon, addIcon } from "@iconify/vue/offline"
import shutter from "@iconify-icons/material-symbols/camera"
// Store
import { useFavoritesStore } from "@/stores/favorites.js"
// Components
import NotFoundMessage from "@/components/panel/NotFoundMessage.vue"

const favorites = useFavoritesStore()

addIcon("shutter", shutter)

const router = useRouter()
const navigate = (slug) => {
  router.push({ name: "shoot", params: { shoot_slug: slug.toUpperCase() } })
}

onMounted(() => {
  favorites.loadFromUserDefaults()
})
</script>
