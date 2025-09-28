import { createRouter, createWebHistory } from "vue-router"
import HomeView from "@/views/HomeView.vue"
import ShootView from "@/views/ShootView.vue"
import ContactView from "@/views/ContactView.vue"
import FavoritesView from "@/views/FavoritesView.vue"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/contact",
      name: "contact",
      component: ContactView,
    },
    {
      path: "/favorites",
      name: "favorites",
      component: FavoritesView,
    },
    {
      path: "/:shoot_slug",
      name: "shoot",
      component: ShootView,
      props: true,
    },
  ],
})

export default router
