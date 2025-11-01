import {defineStore} from "pinia";
import {useFridayApi} from "@/composables/useFridayApi";

export const usePhotosStore = defineStore("photos", {
    state: () => ({
        filterCategory: 0,
        homeLastLoaded: 0,
        homePhotos: [],
        shootPhotos: [],
        shootSlug: '',
        selectedPhotoId: 0,
    }),
    getters: {
        displayPhotos: (state) => {
            // TODO: this can't be exclusive
            if (state.shootSlug !== '') {
                return state.shootPhotos
            } else if (state.filterCategory === 0) {
                return state.homePhotos
            } else {
                return state.homePhotos.filter(photo => photo.category_id === state.filterCategory)
            }
        },
        selectedPhoto(state) {
            return state.displayPhotos.find(photo => photo.id === state.selectedPhotoId) || null
        },
        selectedPhotoIndex(state) {
            return state.displayPhotos.findIndex(photo => photo.id === state.selectedPhotoId)
        }
    },
    actions: {
        setFilterCategory(categoryId) {
            this.category = parseInt(categoryId)
        },
        setShootId(shootSlug) {
            this.shootSlug = shootSlug
            // TODO: API call?
            this.filterCategory = 0
        },
        async loadHomePhotos() {
            const {getApi} = useFridayApi()
            const photos = await getApi('/photos/favorites')
            if (photos) {
                this.homePhotos = photos
            }
            // TODO: cache the last time we loaded the home photos
            // TODO: randomize the order of the home photos
            // TODO: check the last time we loaded the home photos and ignore if we already have them
        }
    },
    persist: true,
})