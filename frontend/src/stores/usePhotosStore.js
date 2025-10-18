import {defineStore} from "pinia";

export const usePhotosStore = defineStore("photos", {
    // Photos property that can be iterated over in shared photo component
    // Filter id property
    // Action to set filter
    // Cache home photos order
    // Photos property likely needs to include a computed function to filter
    // Action to set a shoot
    // Action to set a shoot clears filter by default
    state: () => ({
        filterCategory: 0,
        homePhotos: [],
        shootPhotos: [],
        shootSlug: '',
    }),
    getters: {
        displayPhotos: (state) => {
            if (state.shootSlug !== '') {
                return state.shootPhotos
            } else if (state.filterCategory === 0) {
                return state.homePhotos
            } else {
                return state.homePhotos.filter(photo => photo.category_id === state.filterCategory)
            }
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
        }
        // TODO: Determine what action needs to take place in order to set home photos the first time
    },
    persist: true,
})