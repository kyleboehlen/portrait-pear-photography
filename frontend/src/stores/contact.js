import { defineStore } from "pinia"

export const useContactStore = defineStore({
  id: "contact",
  state: () => ({
    sendNext: 0,
  }),
  actions: {
    setSent() {
      this.sendNext = new Date().getTime() + 3600000 // One hour in miliseconds
    },
  },
  persist: true,
})
