<script setup>
import {addIcon, Icon} from "@iconify/vue/offline";
import share from "@iconify-icons/pajamas/share"
import {ref} from "vue";

addIcon("share", share)

const showTooltip = ref(false)
const copyLink = () => {
  navigator.clipboard.writeText(window.location.href)
  showTooltip.value = true
  setTimeout(() => {
    showTooltip.value = false
  }, 1000)
}
</script>

<template>
  <div
      class=" bg-neutral flex w-auto p-0 aspect-square rounded-xl flex-row justify-center items-center text-white/70 h-3/5 sm:h-3/4"
      :class="{ 'tooltip tooltip-left tooltip-open': showTooltip}"
      data-tip="Copied!" @click="copyLink">
    <Icon icon="share" class="h-2/3 w-auto hover:text-primary hover:cursor-pointer"/>
  </div>
</template>

<style scoped>

</style>