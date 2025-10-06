import { createRouter, createWebHistory } from "vue-router"
import PhotosView from "@/views/PhotosView.vue"
import AdminView from "@/views/AdminView.vue"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: PhotosView,
    },
    {
      path: "/admin",
      name: "admin",
      component: AdminView,
    },
    {
      path: "/:shoot_slug",
      name: "shoot",
      component: PhotosView,
      props: true,
    },
  ],
})

export default router
