import { defineStore } from "pinia";
import { useFridayApi } from "@/composables/useFridayApi";

export const useAdminStore = defineStore("admin",  {
    state: () => ({
        bearerToken: '',
        shoots: [], // TODO: Load this... when?
        updateShoot: {},
    }),
    // TODO: isDirty computed property - for when a shoot or a list of photos is changed w/o being persisted yet
    actions: {
        // Somehow we need to catch forbidden errors and clear the token
        async authenticate(password) {
            const { postApi } = useFridayApi();
            const res = await postApi('/authenticate', { password: password });
            if (res !== false) {
                this.bearerToken = res.token;
            }
        },
        async testToken() {
            const { getApi } = useFridayApi();
            const res = await getApi('/admin/test', this.bearerToken);
        },
        async createShoot(shootName) {
            const { postApi } = useFridayApi();
            const res = await postApi('/admin/shoots', { name: shootName }, this.bearerToken);
            if (res !== false) {
                this.shoots.push(res);
                // TODO: Set query params to update new shoot (i.e. action = update and selector = shoot id)
            }
        },
    }
})