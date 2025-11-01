<script setup>
import {useFridayApi} from "@/composables/useFridayApi";
import {usePhotosStore} from "@/stores/usePhotosStore"
import NotFoundMessage from "@/components/panel/NotFoundMessage.vue";
import TheInstagramButton from "@/components/TheInstagramButton.vue";
import TheHeader from "@/components/TheHeader.vue";
import {onMounted, ref} from "vue";
import TheFooter from "@/components/TheFooter.vue";
import PhotoCard from "@/components/panel/PhotoCard.vue";
import TheFullsizeViewer from "@/components/TheFullsizeViewer.vue";

const props = defineProps(["shoot_slug"])

const {apiCallInProgress} = useFridayApi()
const photosStore = usePhotosStore()

onMounted(async () => {
  await photosStore.loadHomePhotos()
})

const showFullsize = ref(false)
const fullSizeImage = (photo) => {
console.log(photo)
  photosStore.selectedPhotoId = photo.id
  showFullsize.value = true
}
</script>

<template>
  <div class="app-wrapper flex flex-col min-h-screen justify-start">
    <TheHeader/>

    <main class="flex grow">
      <!-- No images found message -->
      <Transition name="fade">
        <NotFoundMessage v-if="!apiCallInProgress && photosStore.displayPhotos.length === 0" class="w-full"/>
      </Transition>

<!--       Actually show images! -->
          <Transition name="fade">
            <div
                v-if="!apiCallInProgress && photosStore.displayPhotos.length > 0"
                class="w-full flex flex-wrap items-center justify-around">
              <PhotoCard
                  v-for="photo in photosStore.displayPhotos"
                  class="hover:cursor-zoom-in"
                  :key="photo.id"
                  :photo="photo"
                  @click="fullSizeImage(photo)" />
            </div>
          </Transition>

      <TheInstagramButton/>

      <TheFullsizeViewer
          v-if="showFullsize"
          class="modal-open"
          @close="showFullsize = false" />
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