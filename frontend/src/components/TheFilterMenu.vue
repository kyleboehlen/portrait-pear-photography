<script setup>
import {addIcon, Icon} from "@iconify/vue/offline";
import {usePhotosStore} from "@/stores/usePhotosStore.js";
import filter from "@iconify-icons/pajamas/filter"
import {ref} from "vue";
import {filterCategories} from "@/constants/categories.js";

addIcon("filter", filter)

const photosStore = usePhotosStore()

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
</script>

<template>
  <details
      ref="dropdownRef"
      @focusout="closeDropdown"
      class="dropdown dropdown-bottom dropdown-end bg-neutral flex w-auto p-0 aspect-square rounded-xl flex-row justify-center items-center text-white/70 h-3/5 sm:h-3/4">
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
</template>

<style scoped>
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