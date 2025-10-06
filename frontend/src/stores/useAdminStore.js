import {defineStore} from "pinia";

export const useAdminStore = defineStore("admin",  {
    state: () => ({
        bearerToken: '',
    }),
    actions: {
        // Somehow we need to catch forbidden errors and clear the token
        authenticate(password) {
            // Call the API to get the auth token
            // If successful, set the bearerToken property
            //
        }
    }
})