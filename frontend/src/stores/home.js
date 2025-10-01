import { defineStore } from "pinia"

/**
 * @deprecated use usePhotosStore instead
 */
export const useHomeStore = defineStore("home",{
  state: () => ({
    lastUpdated: 0, // Last time we've gotten the home view from the API
    photos: [],
  }),
  actions: {
    setLastUpdated() {
      this.lastUpdated = Date.now()
    },
  },
  persist: true,
})
