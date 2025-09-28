<template>
  <input type="checkbox" id="modal-nav" class="modal-toggle" ref="modalToggle" />
  <label for="modal-nav" class="modal cursor-pointer">
    <label class="modal-box relative flex flex-col items-center max-xs:px-0" for="">
      <!-- Filter folders select -->
      <div class="form-control w-11/12 xs:w-5/6 mt-2">
        <label class="label">
          <span class="label-text">Filter Photos</span>
        </label>
        <div class="input-group w-full">
          <select
            v-model="filterSelect"
            id="filter-select"
            class="select select-bordered select-primary w-10/12 focus:outline-none">
            <option value="0">Show All</option>
            <option value="1">Portrait</option>
            <option value="2">Automotive</option>
            <option value="3">Street</option>
            <option value="4">B&W</option>
          </select>
          <button class="btn max-xs:btn-square" @click="filterPhotos">
            <Icon icon="funnel" class="text-primary h-2/3 w-auto"></Icon>
          </button>
        </div>
      </div>

      <!-- Search for shoot -->
      <div class="form-control w-11/12 xs:w-5/6 mt-8">
        <label class="label">
          <span class="label-text">Search Shoots</span>
        </label>
        <div class="input-group w-full" v-click-away="shootIDValid">
          <input
            v-model="shoot"
            type="text"
            placeholder="Shoot ID..."
            maxlength="6"
            class="input input-bordered input-primary focus:outline-none w-10/12"
            :class="{ 'input-primary': !shootIDInvalid, 'input-error': shootIDInvalid }"
            @focus="shootIDInvalid = false" />
          <button class="btn max-xs:btn-square" @click="searchShoot">
            <Icon icon="search" class="text-primary h-2/3 w-auto"></Icon>
          </button>
        </div>
      </div>

      <!-- Spacer -->
      <div class="pt-32"></div>

      <!-- Bottom Nav -->
      <div class="btm-nav mb-2">
        <!-- Home -->
        <BottomButton icon="home" label="Home" @click="navigate('home')" />

        <!-- Contact Me -->
        <BottomButton icon="mail" label="Contact" @click="navigate('contact')" />

        <!-- Instagram -->
        <BottomButton icon="insta" label="Instagram" @click="navigate('instagram')" />

        <!-- If native: Favorited shoots ELSE: app store -->
        <BottomButton v-if="isNative" icon="heart" label="Favorites" @click="navigate('favorites')" />
      </div>
    </label>
  </label>
</template>

<script setup>
// Vue
import { computed, ref, onMounted } from "vue"
import { useRouter } from "vue-router"
// Components
import BottomButton from "@/components/nav/BottomButton.vue"
// Capacitor
import { Capacitor } from "@capacitor/core"
// Store
import { useFilterStore } from "@/stores/filter.js"
// Icons
import { Icon, addIcon } from "@iconify/vue/offline"
import home from "@iconify-icons/pajamas/home"
import search from "@iconify-icons/pajamas/search-sm"
import mail from "@iconify-icons/pajamas/mail"
import insta from "@iconify-icons/mdi/instagram"
import app from "@iconify-icons/ph/app-store-logo-bold"
import funnel from "@iconify-icons/pajamas/filter"
import heart from "@iconify-icons/pajamas/heart"

addIcon("home", home)
addIcon("search", search)
addIcon("mail", mail)
addIcon("insta", insta)
addIcon("app", app)
addIcon("funnel", funnel)
addIcon("heart", heart)

const isNative = computed(() => {
  return Capacitor.isNativePlatform()
})

const filter = useFilterStore()
const filterSelect = ref()
const filterPhotos = () => {
  filter.setCategory(filterSelect.value)
  hideModal()
}

onMounted(() => {
  filterSelect.value = filter.category
})

const modalToggle = ref()

// Navigation
const router = useRouter()
const navigate = (target) => {
  if (target == "instagram") {
    window.open("https://www.instagram.com/portraitpear/", "_blank")
  }

  // Push other routes
  router.push({ name: target })
  hideModal()
}

const shoot = ref()
const searchShoot = () => {
  if (shoot?.value == null || shoot.value.length !== 6) {
    shootIDInvalid.value = true
  } else {
    router.push({ name: "shoot", params: { shoot_slug: shoot.value.toUpperCase() } })
    hideModal()
  }
}

const hideModal = () => {
  modalToggle.value.checked = !modalToggle.value.checked
}

// Highlight the shoot search in red
const shootIDInvalid = ref(false)
const shootIDValid = () => {
  shootIDInvalid.value = false
}
</script>
