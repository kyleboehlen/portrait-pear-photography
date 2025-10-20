import {defineStore} from "pinia";
import {useFridayApi} from "@/composables/useFridayApi";

const {postApi, getApi, deleteApi} = useFridayApi();

export const useAdminStore = defineStore("admin", {
    state: () => ({
        bearerToken: '',
        shoots: [],
        selectedShootId: 0,
        updateShoot: {},
        previewPhotoUrl: '',
    }),
    getters: {
        selectedShoot(state) {
            return state.shoots.find(shoot => shoot.id === state.selectedShootId) || {};
        },
        isDirty(state) {
            return state.updateShoot && JSON.stringify(state.updateShoot) !== JSON.stringify(state.selectedShoot);
        },
    },
    actions: {
        // Somehow we need to catch forbidden errors and clear the token
        async authenticate(password) {
            const res = await postApi('/authenticate', {password: password});
            if (res !== false) {
                this.bearerToken = res.token;
            }
        },
        async testToken() {
            await getApi('/admin/test', this.bearerToken);
        },
        async createShoot(shootName) {
            const res = await postApi('/admin/shoot', {name: shootName}, this.bearerToken);
            if (res !== false) {
                this.selectedShootId = res.id;
            }
        },
        async loadShootsFromApi() {
            const res = await getApi('/admin/shoots', this.bearerToken);
            if (res !== false) {
                this.shoots = res;
            }
        },
        async deleteSelectedShoot() {
            const res = await deleteApi(`/admin/unshoot`, {id: this.selectedShootId}, this.bearerToken);
            if (res !== false) {
                this.shoots = this.shoots.filter(shoot => shoot.id !== this.selectedShootId);
                this.selectedShootId = 0;
            }
        }
    },
    persist: true,
})