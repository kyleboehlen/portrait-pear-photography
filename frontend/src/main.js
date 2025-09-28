import { createApp } from "vue"
import { createPinia } from "pinia"
import piniaPluginPersistedstate from "pinia-plugin-persistedstate"
import { Capacitor } from "@capacitor/core"
import { App as Capp } from "@capacitor/app"
import { CapacitorSQLite, SQLiteConnection } from "@capacitor-community/sqlite"

import App from "./App.vue"
import router from "./router"
import VueClickAway from "vue3-click-away"
import VueLazyLoad from "vue3-lazyload"

import "./assets/styles/app.css"

const app = createApp(App)
const pinia = createPinia()

pinia.use(piniaPluginPersistedstate)
app.use(pinia)
app.use(router)
app.use(VueClickAway) // Makes 'v-click-away' directive usable in every component
app.use(VueLazyLoad)

// Deep links
Capp.addListener("appUrlOpen", function (event) {
  // Example url: https://www.portraitpear.photography/HEYYOU
  // slug = /HEYYOU
  const slug = event.url.split(".photography").pop()

  // We only push to the route if there is a slug present
  if (slug) {
    router.push({
      path: slug,
    })
  }
})

async function setUpDatabase() {
  const sqlite = new SQLiteConnection(CapacitorSQLite)
  const ret = await sqlite.checkConnectionsConsistency()
  const isConn = (await sqlite.isConnection("PortraitPearDatabase")).result
  let db = null
  if (ret.result && isConn) {
    db = await sqlite.retrieveConnection("PortraitPearDatabase")
  } else {
    db = await sqlite.createConnection("PortraitPearDatabase", false, "no-encryption", 1)
  }
  await db.open()
  // DROP TABLE IF EXISTS images;
  const createImagesTable = `
    CREATE TABLE IF NOT EXISTS images (
    id INTEGER PRIMARY KEY NOT NULL,
    asset_url STRING NOT NULL,
    base_sixty_four TEXT NOT NULL
    );
    CREATE INDEX IF NOT EXISTS images_index_asset_rl ON images (asset_url);
    `
  await db.execute(createImagesTable, false)
  return db
}

// Set up sqlite DB
if (Capacitor.isNativePlatform()) {
  setUpDatabase().then((db) => {
    app.provide("db", db)
    app.mount("#app")
  })
} else {
  app.provide("db", null)
  app.mount("#app")
}
