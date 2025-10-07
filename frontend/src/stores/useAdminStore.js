import { defineStore } from "pinia";
import { useFridayApi } from "@/composables/useFridayApi";

export const useAdminStore = defineStore("admin",  {
    state: () => ({
        bearerToken: '',
    }),
    actions: {
        // Somehow we need to catch forbidden errors and clear the token
        async authenticate(password) {
            const { postApi } = useFridayApi();
            const res = await postApi('/authenticate', { password: password });
            if (res !== false) {
                this.bearerToken = res.token;
            }
        }
    }
})