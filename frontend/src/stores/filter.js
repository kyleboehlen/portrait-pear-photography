import { defineStore } from "pinia"

export const useFilterStore = defineStore("filter",{
  state: () => ({
    category: 0,
  }),
  actions: {
    setCategory(categoryId) {
      this.category = parseInt(categoryId)
    },
  },
})
