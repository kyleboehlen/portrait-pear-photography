import {defineStore} from "pinia";
import {useFridayApi} from "@/composables/useFridayApi";

const {postApi, getApi, deleteApi} = useFridayApi();

export const useAdminStore = defineStore("admin", {
    state: () => ({
        bearerToken: '',
        shoots: [],
        selectedShootId: 0,
        updateShoot: {default_categories: []},
        previewPhotoUrl: '',
        previewPhotos: [],
        photosToMutate: [],
    }),
    getters: {
        selectedShoot(state) {
            return state.shoots.find(shoot => shoot.id === state.selectedShootId) || {default_categories: []};
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
        async persistUpdateShoot() {
            // Do not persist if nothing has changed, or we don't have shoot information to update
            if (!this.updateShoot || !this.selectedShootId || !this.isDirty) {
                return;
            }

            // Convert default_categories to array of integers to keep the API happy
            let shootToUpdate = {...this.updateShoot};
            if (shootToUpdate.default_categories) {
                shootToUpdate.default_categories = shootToUpdate.default_categories.map(cat => parseInt(cat, 10));
            }

            const res = await postApi('/admin/shoot', shootToUpdate, this.bearerToken);

            // On a successful update we need to update the shoot array with the persisted data
            if (res !== false) {
                const index = this.shoots.findIndex(shoot => shoot.id === this.selectedShootId);
                if (index !== -1) {
                    this.shoots[index] = this.updateShoot
                }
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
        },
        async refreshPreviewPhotosFromApi(filterParams = {}) {
            this.previewPhotos = await postApi('/admin/photos', {filter_parameters: filterParams}, this.bearerToken);
        },
        async deletePhotos() {
            const {deleteApi} = useFridayApi();
            this.photosToMutate.forEach(photo => {
                let photoId = parseInt(photo);
                deleteApi('/admin/delete-photo', {id: photoId}, this.bearerToken);
            })
            this.photosToMutate = []
            // TODO: Refresh the preview photos after deletion
        }
    },
    persist: true,
})