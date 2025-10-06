import axios from 'axios';
import { ref } from 'vue';

const baseUrl = import.meta.env.VITE_API_URL;
const apiCallInProgress = ref(false);
const adminApiIsAuthenticated = ref(false);

async function baseApiCall(method, path, body = null, token = null) {
    apiCallInProgress.value = true
    try {
        const config = {
            method,
            url: `${baseUrl}${path}`,
            data: body,
        };

        if (token) {
            config.headers = {
                'Authorization': `Bearer ${token}`
            };
        }

        const response = await axios(config);
        const apiRes = response.data;

        apiCallInProgress.value = false

        if (!apiRes.success) {
            // TODO: If token was provided and we got a forbidden response set authenticated to false
            console.error(`API Error [${apiRes.error_code}]: ${apiRes.error_message}`);
            return false;
        }

        if (token) {
            adminApiIsAuthenticated.value = true;
        }

        return apiRes.content;
    } catch (err) {
        console.error('API error:', err);

        apiCallInProgress.value = false
        return false;
    }
}

function deleteApi(path, token = null) {
    return baseApiCall('delete', path, token);
}

function getApi(path, token = null) {
    return baseApiCall('get', path, null, token);
}

function postApi(path, body, token = null) {
    return baseApiCall('post', path, body, token);
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
    return { apiCallInProgress, adminApiIsAuthenticated, deleteApi, getApi, postApi };
};