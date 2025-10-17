import { defineStore } from "pinia";
import { useFridayApi } from "@/composables/useFridayApi";

const { postApi, getApi } = useFridayApi();

export const useAdminStore = defineStore("admin",  {
    state: () => ({
        bearerToken: '',
        shoots: [],
        selectedShootId: 0,
        updateShoot: {},
    }),
    // TODO: isDirty computed property - for when a shoot or a list of photos is changed w/o being persisted yet
    actions: {
        // Somehow we need to catch forbidden errors and clear the token
        async authenticate(password) {
            const res = await postApi('/authenticate', { password: password });
            if (res !== false) {
                this.bearerToken = res.token;
            }
        },
        async testToken() {
            await getApi('/admin/test', this.bearerToken);
        },
        async createShoot(shootName) {
            const res = await postApi('/admin/shoot', { name: shootName }, this.bearerToken);
            if (res !== false) {
                this.selectedShootId = res.id;
            }
        },
        async loadShootsFromApi(){
            const res = await getApi('/admin/shoots', this.bearerToken);
            if (res !== false) {
                this.shoots = res;
            }
        }
    },
    persist: true,
})