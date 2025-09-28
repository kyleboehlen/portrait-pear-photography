import { defineStore } from "pinia"
// Vue
import { ref } from "vue"
// Capacitor
import { Preferences } from "@capacitor/preferences"
import { Capacitor } from "@capacitor/core"

export const useFavoritesStore = defineStore("favorites", () => {
  const shoots = ref([])

  const addShoot = (slug, name, photos) => {
    shoots.value.push({
      slug: slug,
      name: name,
      photos: photos,
    })
    persistToUserDefaults()
  }

  const removeShoot = (slug) => {
    shoots.value = shoots.value.filter((shoot) => {
      return shoot.slug !== slug
    })
    persistToUserDefaults()
  }

  const persistToUserDefaults = async () => {
    if (Capacitor.isNativePlatform()) {
      const shootsJson = JSON.stringify(shoots.value)
      await Preferences.set({
        key: "favoriteShoots",
        value: shootsJson,
      })
    }
  }

  const loadFromUserDefaults = async () => {
    if (Capacitor.isNativePlatform()) {
      const { value } = await Preferences.get({ key: "favoriteShoots" })
      if (value !== null) {
        shoots.value = JSON.parse(value)
      }
    }
  }

  return { shoots, addShoot, removeShoot, loadFromUserDefaults, persistToUserDefaults }
})
