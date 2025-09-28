<template>
  <!-- If message hasn't been sent -->
  <main v-if="!messageSent" class="flex flex-row justify-center align-center">
    <div class="w-11/12 sm:w-2/3 xl:w-1/2 flex flex-row flex-wrap justify-between content-center">
      <!-- NAME -->
      <div class="form-control w-full md:w-1/2 md:pr-7">
        <label class="label">
          <span class="label-text">Name</span>
        </label>
        <input
          v-model="name"
          type="text"
          placeholder="Annie Leibovitz"
          class="input input-bordered input-primary"
          :class="{ 'input-primary': nameValid, 'input-error': !nameValid }"
          @focus="nameValid = true" />
      </div>

      <!-- INSTAGRAM -->
      <div class="form-control w-full md:w-1/2 md:pl-7 max-md:pt-4">
        <label class="label">
          <span class="label-text">Instagram</span>
        </label>
        <label class="input-group w-full">
          <span class="bg-neutral text-primary text-2xl w-1/6 flex justify-center">@</span>
          <input
            v-model="insta"
            type="text"
            placeholder="annieleibovitz"
            class="input input-bordered input-primary w-5/6" />
        </label>
      </div>

      <!-- EMAIL/PHONE -->
      <div class="flex w-full max-md:flex-col mt-2 sm:mt-4 md:mt-8">
        <div class="grid md:h-20 flex-grow card rounded-box place-items-center">
          <div class="form-control w-full">
            <label class="label">
              <span class="label-text">Email</span>
            </label>
            <input
              v-model="email"
              type="email"
              placeholder="youremail@gmail.com"
              class="input input-bordered"
              :class="{ 'input-primary': emailValid, 'input-error': !emailValid }"
              :disabled="phone?.length > 0"
              @focus="phoneValid = emailValid = true" />
          </div>
        </div>
        <div class="divider md:divider-horizontal md:mx-6 max-md:mb-0">OR</div>
        <div class="grid md:h-20 flex-grow card rounded-box place-items-center">
          <div class="form-control w-full">
            <label class="label max-md:pt-0">
              <span class="label-text">Phone #</span>
            </label>
            <input
              v-model="phone"
              type="tel"
              placeholder="8019991234"
              class="input input-bordered"
              :class="{ 'input-primary': phoneValid, 'input-error': !phoneValid }"
              :disabled="email?.length > 0"
              @focus="phoneValid = emailValid = true" />
          </div>
        </div>
      </div>

      <!-- TEXTAREA -->
      <div class="form-control w-full mt-2 sm:mt-4 md:mt-8 grow">
        <label class="label">
          <span class="label-text">Message</span>
        </label>
        <textarea
          v-model="message"
          placeholder="Hey! Let's shoot."
          class="textarea textarea-bordered h-100 textarea-primary"
          :class="{ 'textarea-primary': messageValid, 'textarea-error': !messageValid }"
          @focus="messageValid = true"></textarea>
      </div>

      <!-- Failed to send -->
      <div v-if="messageFailed" class="alert alert-error shadow-lg mt-8">
        <div>
          <Icon icon="failed" class="font-extrabold mr-2 text-3xl text-center" />
          <span>Failed to send - try again later</span>
        </div>
      </div>

      <!-- BUTTON -->
      <div class="w-full flex justify-center sm:justify-end mt-8 max-md:pb-4">
        <button class="btn btn-outline btn-success px-16 sm:px-8" @click="fullSend" :disabled="messageSending">
          {{ buttonText }}
        </button>
      </div>
    </div>
  </main>

  <!-- Message sent successfully -->
  <main v-else class="flex flex-row justify-center align-center grow">
    <div class="w-96 flex flex-col justify-between content-center pt-16">
      <!-- Success alert -->
      <div class="alert alert-success shadow-lg mt-8">
        <div>
          <Icon icon="success" class="font-extrabold mr-2 text-3xl" />
          <span>Message sent!</span>
        </div>
      </div>
    </div>
    <div></div>
  </main>
</template>

<script setup>
// axios
import axios from "axios"
// Vue
import { ref, computed, onMounted } from "vue"
// Store
import { useContactStore } from "@/stores/contact.js"
// Icons
import { Icon, addIcon } from "@iconify/vue/offline"
import failed from "@iconify-icons/iconoir/bubble-error"
import success from "@iconify-icons/iconoir/chat-bubble-check"

// Icon
addIcon("failed", failed)
addIcon("success", success)

// Store
const contact = useContactStore()

// v-models
const name = ref()
const insta = ref()
const email = ref()
const phone = ref()
const message = ref()

// Validation refs
const nameValid = ref(true)
const emailValid = ref(true)
const phoneValid = ref(true)
const messageValid = ref(true)

// Sending/Display State
const messageSent = ref(false)
const messageSending = ref(false)
const messageFailed = ref(false)

const buttonText = computed(() => {
  if (messageSending.value) {
    if (messageFailed.value) {
      return "Send Failed"
    } else {
      return "Sending..."
    }
  }
  return "Full Send"
})

// Validation functions
const emailIsValid = computed(() => {
  // Email validation regex
  const regex = RegExp(/^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/)

  // Get parts of the email that match the regex
  const match = email.value.match(regex)

  // If we have a match verify it matches the entire email submitted for validation
  return match !== null && match[0] === email.value
})

const phoneIsValid = computed(() => {
  // Phone number validation regex
  const regex = RegExp(/^\(?(\d{3})\)?[- ]?(\d{3})[- ]?(\d{4})$/)

  // Get parts of the phone that match the regex
  const match = phone.value.match(regex)

  // If we have a match verify it matches the entire phone number submitted for validation
  return match !== null && match[0] === phone.value
})

// Full send
const apiUrl = import.meta.env.VITE_API_URL
const fullSend = () => {
  // Show sending
  messageSending.value = true

  // Validate name
  if (name.value == null || name.value.length == 0 || name.value.length > 50) {
    nameValid.value = false
  }

  // Validate email/phone
  let contactInfo = null
  if ((email.value == null || email.value.length == 0) && (phone.value == null || phone.value.length == 0)) {
    emailValid.value = phoneValid.value = false
  } else if (email.value != null && email.value.length > 0) {
    emailValid.value = emailIsValid.value
    contactInfo = email.value
  } else if (phone.value != null && phone.value.length > 0) {
    phoneValid.value = phoneIsValid.value
    contactInfo = phone.value
  }

  // Validate message
  if (message.value == null || message.value.length == 0) {
    messageValid.value = false
  }

  // Make api call
  if (nameValid.value && emailValid.value && phoneValid.value && messageValid.value) {
    axios({
      method: "post",
      url: `${apiUrl}/api/pear/contact`,
      data: {
        name: name.value,
        contact_info: contactInfo,
        reason: "photography",
        instagram: insta?.value,
        message: message.value,
      },
    })
      .then(function () {
        // Mark message as sent
        messageSent.value = true
        contact.setSent()
      })
      .catch(function () {
        // Show/Hide error alert
        messageFailed.value = true
        setTimeout(() => {
          messageFailed.value = false
          messageSending.value = false
        }, 5000)
      })
  } else {
    messageSending.value = false
  }
}

// On mount check if we've already sent a message
onMounted(() => {
  if (contact.sendNext > new Date().getTime()) {
    messageSent.value = true
  }
})
</script>
