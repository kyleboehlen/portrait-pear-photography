import {defineStore} from "pinia";
import {useFridayApi} from "@/composables/useFridayApi";

export const usePhotosStore = defineStore("photos", {
    state: () => ({
        filterCategory: 0,
        homeLastLoaded: 0,
        homePhotos: [],
        shootPhotos: [],
        loadedShootSlug: '',
        selectedShootSlug: '',
        selectedPhotoId: 0,
    }),
    getters: {
        displayPhotos: (state) => {
            if (state.selectedShootSlug !== '') {
                return state.shootPhotos
            } else {
                if (state.filterCategory === 0) {
                    return state.homePhotos
                } else {
                    return state.homePhotos.filter(photo => photo.category_id === state.filterCategory)
                }
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
            this.filterCategory = parseInt(categoryId)
        },
        async loadPhotos() {
            if (this.selectedShootSlug !== '') {
                if (this.loadedShootSlug !== this.selectedShootSlug) {
                    this.filterCategory = 0
                    const {postApi} = useFridayApi()
                    this.shootPhotos = await postApi('/photos/shoot', {shoot_slug: this.selectedShootSlug}) || []
                    this.loadedShootSlug = this.selectedShootSlug
                }
            } else {
                const {getApi} = useFridayApi()
                // 30 minutes in milliseconds
                if (this.homeLastLoaded + 1800000 <= Date.now()) {
                    const photos = await getApi('/photos/favorites')
                    if (photos) {
                        // Fisher-Yates (Knuth) shuffle:
                        let currentIndex = photos.length;
                        let randomIndex;

                        // While there remain elements to shuffle
                        while (currentIndex !== 0) {
                            // Pick a remaining element
                            randomIndex = Math.floor(Math.random() * currentIndex);
                            currentIndex--;

                            // And swap it with the current element
                            [photos[currentIndex], photos[randomIndex]] = [
                                photos[randomIndex],
                                photos[currentIndex],
                            ];
                        }

                        this.homePhotos = photos
                        this.homeLastLoaded = Date.now()
                    }
                }
            }
        }
    },
    persist: true,
})