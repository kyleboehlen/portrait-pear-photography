<script setup>
import {useFridayApi} from "@/composables/useFridayApi";
import {usePhotosStore} from "@/stores/usePhotosStore"
import NotFoundMessage from "@/components/panel/NotFoundMessage.vue";
import TheInstagramButton from "@/components/TheInstagramButton.vue";
import TheHeader from "@/components/TheHeader.vue";
import {computed, ref} from "vue";
import TheFooter from "@/components/TheFooter.vue";
import PhotoCard from "@/components/panel/PhotoCard.vue";
import TheFullSizeViewer from "@/components/TheFullSizeViewer.vue";
import TheFilterMenu from "@/components/TheFilterMenu.vue";
import TheShareButton from "@/components/TheShareButton.vue";

const {apiCallInProgress} = useFridayApi()
const photosStore = usePhotosStore()

const isShootView = computed(() => {
  return photosStore.selectedShootSlug !== ''
})

const showFullSize = ref(false)
const fullSizeImage = (photo) => {
  photosStore.selectedPhotoId = photo.id
  showFullSize.value = true
}
</script>

<template>
  <div class="app-wrapper flex flex-col min-h-screen justify-start">
    <TheHeader>
      <template v-slot:action>
        <TheShareButton v-if="isShootView"/>
        <TheFilterMenu v-else/>
      </template>
    </TheHeader>

    <main class="flex grow">
      <!-- No images found message -->
      <Transition name="fade">
        <NotFoundMessage v-if="!apiCallInProgress && photosStore.displayPhotos.length === 0" class="w-full"/>
      </Transition>

      <!-- Actually show images! -->
      <Transition name="fade">
        <div
            v-if="!apiCallInProgress && photosStore.displayPhotos.length > 0"
            class="w-full flex flex-wrap items-center justify-around">
          <PhotoCard
              v-for="photo in photosStore.displayPhotos"
              class="hover:cursor-pointer"
              :key="photo.id"
              :photo="photo"
              @click="fullSizeImage(photo)"/>
        </div>
      </Transition>

      <TheInstagramButton/>

      <TheFullSizeViewer
          v-if="showFullSize"
          class="modal-open"
          @close="showFullSize = false"/>
    </main>

    <TheFooter class="justify-self-end"/>
  </div>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease;
}

.fade-enter-active {
  transition-delay: 0.5s;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>