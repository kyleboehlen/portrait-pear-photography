import { defineStore } from "pinia"
import { inject } from "vue"

export const useImagesStore = defineStore("images", () => {
  const db = inject("db")

  async function store(assetUrl, baseSixtyFour) {
    const insertIntoImages = `
      INSERT INTO images (asset_url,base_sixty_four) VALUES ("${assetUrl}","${baseSixtyFour}")
      `
    const res = await db.execute(insertIntoImages, false)
    return res.changes.changes > 0
  }

  async function getAll() {
    const selectAllImages = `
      SELECT * FROM images  
      `
    const res = await db.query(selectAllImages)
    return res.values.map((v) => v.asset_url)
  }

  async function getBase64(assetUrl) {
    const selectImage = `
      SELECT base_sixty_four FROM images WHERE asset_url = "${assetUrl}"
      `
    const res = await db.query(selectImage)
    return res.values[0].base_sixty_four
  }

  // Results in state and actions on a standard Pinia store
  return { store, getAll, getBase64 }
})
