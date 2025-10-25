<script setup>
import {useFridayApi} from "@/composables/useFridayApi";
import {usePhotosStore} from "@/stores/usePhotosStore"
import NotFoundMessage from "@/components/panel/NotFoundMessage.vue";
import TheInstagramButton from "@/components/TheInstagramButton.vue";
import TheHeader from "@/components/TheHeader.vue";
import TheFooter from "@/components/TheFooter.vue";

const props = defineProps(["shoot_slug"])

const {apiCallInProgress} = useFridayApi()
const photos = usePhotosStore()
// Health check on mounted
// Set shoot ID if it exists in the route
//
</script>

<template>
  <div class="app-wrapper flex flex-col min-h-screen justify-start">
    <TheHeader/>

    <main class="flex grow">
      <!-- No images found message -->
      <Transition name="fade">
        <NotFoundMessage v-if="!apiCallInProgress && photos.displayPhotos.length === 0" class="w-full"/>
      </Transition>

      <!-- Actually show images! -->
      <!--    <Transition name="fade">-->
      <!--      <div-->
      <!--          v-if="apiCallFinished && filteredPhotos.length > 0"-->
      <!--          class="w-full flex flex-wrap items-center justify-around">-->
      <!--        <PhotoCard-->
      <!--            v-for="photo in filteredPhotos"-->
      <!--            class="hover:cursor-zoom-in"-->
      <!--            :key="photo.id"-->
      <!--            :photo="photo"-->
      <!--            :cachedPhotos="cachedPhotos"-->
      <!--            cache="true"-->
      <!--            @click="fullsizeImage(photo)" />-->
      <!--      </div>-->
      <!--    </Transition>-->

      <TheInstagramButton/>
      <!-- Full size viewer -->
      <!--    <TheFullsizeViewer-->
      <!--        v-if="showFullsize"-->
      <!--        class="modal-open"-->
      <!--        :photo="fullsizePhoto"-->
      <!--        :photos="filteredPhotos"-->
      <!--        :cachedPhotos="cachedPhotos"-->
      <!--        @close="showFullsize = false" />-->
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