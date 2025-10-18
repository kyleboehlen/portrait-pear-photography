import {createApp} from "vue"
import {createPinia} from "pinia"
import piniaPluginPersistedstate from "pinia-plugin-persistedstate"

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

app.mount("#app")