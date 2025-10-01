import { ref } from 'vue';
export const useFridayApi = () => {
    // Set base URL for Friday API
    // Base POST and GET functions
    // Base function checks errors, error code, and logs to console
    // Base function parses and returns JSON

    // Health check on first reference

    // Should we expose a ref here for when an API call is in progress?
    const apiCallInProgress = ref(false);

    // Return
    return { apiCallInProgress };
}