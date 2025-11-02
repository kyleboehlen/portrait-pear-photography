<script setup>
import logo from "@/assets/imgs/green-camera-logo.png?url"
import {onMounted, ref} from "vue"
import {useRouter} from "vue-router"
import {addIcon, Icon} from "@iconify/vue/offline"
import hamburger from "@iconify-icons/pajamas/hamburger"
import funnel from "@iconify-icons/pajamas/filter"
import {usePhotosStore} from "@/stores/usePhotosStore.js"

addIcon("hamburger", hamburger)
addIcon("filter", funnel)

const photosStore = usePhotosStore()

// Shrink on scroll
const shrink = ref(false)
const shrinkOnScroll = () => {
  if (window.scrollY > 50) {
    shrink.value = true
  } else {
    shrink.value = false
  }
}

// Mounted
onMounted(() => {
  // Mount on scroll for shrink
  window.addEventListener("scroll", shrinkOnScroll)
})

// Home link
const router = useRouter()
const navigateHome = () => {
  router.push({name: "home"})
}

// Filter dropdown
const dropdownRef = ref(null)
const setCategoryTo = (categoryId) => {
  photosStore.setFilterCategory(categoryId)
}
const closeDropdown = () => {
  if (dropdownRef.value) {
    dropdownRef.value.removeAttribute('open')
  }
}

const filterCategories = [
  {
    id: 0,
    name: "Show All"
  },
  {
    id: 1,
    name: "Portrait"
  },
  {
    id: 2,
    name: "Automotive"
  },
  {
    id: 3,
    name: "Street"
  },
  {
    id: 4,
    name: "B&W"
  },
]
</script>

<template>
  <!-- Header -->
  <header class="w-full sticky top-0">
    <!-- Header proper -->
    <div
        class="w-full px-2 bg-secondary flex flex-row flex-nowrap justify-between items-center rounded-b-xl"
        :class="{ 'h-28': !shrink, 'md:h-24': !shrink, 'h-12': shrink }">
      <img :src="logo" class="h-full xs:h-full cursor-pointer" @click="navigateHome"/>
      <h1
          class="md:whitespace-nowrap text-center uppercase text-white/75"
          :class="{ 'text-4xl': !shrink, 'sm:text-5xl': !shrink, 'md:text-6xl': !shrink, 'text-3xl': shrink }">
        Portrait Pear
      </h1>
      <slot name="action">
        <details
            ref="dropdownRef"
            @focusout="closeDropdown"
            class="dropdown dropdown-bottom dropdown-end bg-neutral flex w-auto p-0 aspect-square rounded-xl flex-row justify-center items-center text-white/70"
            :class="{ 'h-3/5': !shrink, 'sm:h-3/4': !shrink, 'h-6/7': shrink }">
          <summary
              class="h-full w-auto hover:text-primary hover:cursor-pointer flex flex-col justify-center items-center">
            <Icon icon="filter" class="h-2/3 w-auto"></Icon>
          </summary>
          <ul class="dropdown-content menu bg-neutral rounded-box z-1 w-36 p-2 mt-1">
            <li v-for="category in filterCategories" :id="category.id"
                tabindex="0"
                class="hover:text-primary hover:cursor-pointer flex flex-row justify-end"
                @focusin="setCategoryTo(category.id)">
              <a
                  class="text-right text-lg">{{ category.name }}</a></li>
          </ul>
        </details>
      </slot>
    </div>
  </header>
</template>

<style scoped>
h1 {
  font-family: Kodchasan, sans-serif;
  transition: 0.2s;
}

header {
  transition: 0.2s;
}

details summary {
  list-style: none;
}

details summary::-webkit-details-marker {
  display: none;
}

.dropdown-content li a:hover, .dropdown-content li:hover a {
  background-color: hsl(var(--color-neutral)) !important;
  color: inherit !important;
}
</style>
