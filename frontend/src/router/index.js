import { createRouter, createWebHistory } from "vue-router"
import PhotosView from "@/views/PhotosView.vue"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: PhotosView,
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
