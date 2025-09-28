import { defineStore } from "pinia"

export const useHomeStore = defineStore({
  id: "home",
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
