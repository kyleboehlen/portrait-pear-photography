<script setup>
import {useAdminStore} from "@/stores/useAdminStore.js";
import {usePhotoUtils} from "@/composables/usePhotoUtils.js";
import {useAdminRoutes} from "@/composables/useAdminRoutes.js";
import {computed, ref} from "vue";
import {addIcon, Icon} from "@iconify/vue/offline";
import heart from "@iconify-icons/mdi/cards-heart-outline"
import {styleCategories} from "@/constants/categories.js";

addIcon("heart", heart);

const props = defineProps(['photo'])

const adminStore = useAdminStore();
const isSelectedForDelete = computed(() => {
  if (action.value !== 'delete') {
    return false;
  }

  return adminStore.photosToMutate.includes(props.photo.id);
})

const previewPhotoIndex = computed(() => {
  return adminStore.previewPhotos.findIndex(photo => photo.id === props.photo.id);
})
const isFavorite = computed(() => {
  return adminStore.previewPhotos[previewPhotoIndex.value].favorite
})

const {action} = useAdminRoutes();

const {gridUrl} = usePhotoUtils();
const handleMouseEnter = () => {
  adminStore.previewPhotoUrl = gridUrl(props.photo.guid);
}

const handleMouseLeave = () => {
  adminStore.previewPhotoUrl = '';
}

const categories = ref(props.photo.categories);
const toggleCategory = (categoryId) => {
  if (categories.value.includes(categoryId)) {
    categories.value.splice(categories.value.findIndex(id => id === categoryId), 1);
  } else {
    categories.value.push(categoryId);
  }
}

const handleEntryClick = () => {
  if (action.value === 'delete') {
    if (isSelectedForDelete.value) {
      adminStore.photosToMutate.splice(adminStore.photosToMutate.findIndex(id => id === props.photo.id), 1);
    } else {
      adminStore.photosToMutate.push(props.photo.id);
    }
  } else {
    if (action.value === 'favorite') {
      adminStore.previewPhotos[previewPhotoIndex.value].favorite = !isFavorite.value;
    }

    if (action.value === 'categorize') {
      adminStore.previewPhotos[previewPhotoIndex.value].categories = categories.value
    }

    if (!adminStore.photosToMutate.includes(props.photo.id)) {
      adminStore.photosToMutate.push(props.photo.id);
    }
  }
}
</script>

<template>
  <div class="h-10 w-full flex flex-row justify-start items-center hover:bg-neutral" @mouseenter="handleMouseEnter"
       @mouseleave="handleMouseLeave" @click="handleEntryClick">
    <div class="border-b-2 border-white w-1/10"></div>
    <slot>
      <Icon v-if="action === 'favorite'" icon="heart" class="ml-2 text-base-100"
            :class="{ 'text-base-100': !isFavorite, 'text-red-600': isFavorite}"/>
      <p v-else class="text-white ml-2">{{ props.photo.id }} - </p>
      <p class="pl-1"
         :class="{
        'text-white': !isSelectedForDelete,
        'text-red-600 line-through': isSelectedForDelete
      }">{{ photo.guid }}</p>
      <div v-if="action === 'categorize'" class="w-auto flex flex-row flex-nowrap ml-2 gap-2">
        <span v-for="category in styleCategories" :key="category.id" class="badge badge-sm hover:cursor-pointer flex-1"
              :class="{
          'badge-dash': categories.includes(category.id) === false,
          [category.colorClass]: true
        }" @click="toggleCategory(category.id)">
          {{ category.name }}
        </span>
      </div>
      <!--      Show categories to toggle on hover if action is categorize - mutate the image in preview images -->
    </slot>
  </div>
</template>