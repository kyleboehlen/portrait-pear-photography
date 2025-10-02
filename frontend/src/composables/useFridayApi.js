import axios from 'axios';
import { ref } from 'vue';

const baseUrl = import.meta.env.VITE_API_URL;
const apiCallInProgress = ref(false);

async function baseApiCall(method, path, body = null) {
    apiCallInProgress.value = true
    try {
        const response = await axios({
            method,
            url: `${baseUrl}${path}`,
            data: body,
        });

        const apiRes = response.data;
        apiCallInProgress.value = false

        if (!apiRes.success) {
            console.error(`API Error [${apiRes.error_code}]: ${apiRes.error_message}`);
            return false;
        }

        return apiRes.content;
    } catch (err) {
        console.error('API error:', err);

        apiCallInProgress.value = true
        return false;
    }
}

function getApi(path) {
    return baseApiCall('get', path);
}

function postApi(path, body) {
    return baseApiCall('post', path, body);
}

async function healthCheck() {
    // Because the API is serverless, we hit the health endpoint first to warm it up
    console.log('Warming up the API...')

    const result = await getApi('/health');
    if (result !== false) {
        console.log('Health check success:', result); // This should write "heartbeat"
    } else {
        console.log('Health check failed.');
    }
}

export const useFridayApi = () => {
    // Perform a health check when the composable is first used
    healthCheck();

    // Expose the ref for logic regarding waiting for API calls, and expose methods to make GET and POST requests
    return { apiCallInProgress, getApi, postApi };
};