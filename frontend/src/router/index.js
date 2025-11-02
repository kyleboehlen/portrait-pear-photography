import {createRouter, createWebHistory} from "vue-router"
import PhotosView from "@/views/PhotosView.vue"
import AdminView from "@/views/AdminView.vue"
import {usePhotosStore} from "@/stores/usePhotosStore.js"

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: "/",
            name: "home",
            component: PhotosView,
            beforeEnter(to)  {
                const photosStore = usePhotosStore();

                if (to.query?.category) {
                    photosStore.setFilterCategory(to.query?.category)
                }

                photosStore.selectedShootSlug = '';

                photosStore.loadPhotos();
            },
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
            beforeEnter(to)  {
                const photosStore = usePhotosStore();

                if (to.params?.shoot_slug) {
                    photosStore.selectedShootSlug = to.params.shoot_slug;
                } else {
                    photosStore.selectedShootSlug = '';
                }

                photosStore.setFilterCategory(0);

                photosStore.loadPhotos();
            },
        },
    ],
})

export default router
