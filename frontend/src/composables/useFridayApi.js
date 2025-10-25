import axios from 'axios';
import {ref} from 'vue';

const baseUrl = import.meta.env.VITE_API_URL;
const apiCallInProgress = ref(false);
const adminApiIsAuthenticated = ref(false);

async function baseApiCall(method, path, body = null, token = null, blob = false) {
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

        if (body instanceof FormData) {
            config.headers = {
                ...config.headers,
                'Content-Type': 'multipart/form-data',
            };
        }

        if (blob) {
            config.responseType = 'blob';
        }

        const response = await axios(config);
        const apiRes = response.data;
        /* Expected API response structure:
            {
                "success": bool,
                "error_code": number, // 0 if success is true
                "error_message": "string", // "" if success is true
                "content": {...}, // JSON content based on response struct
             }
         */

        apiCallInProgress.value = false

        if (apiRes instanceof Blob) {
            return new Blob([apiRes]);
        }

        if (!apiRes.success) {
            const errorCode = apiRes.error_code;

            // These are specifically the error codes for admin auth related issues, if we get one of these then the
            // token is no longer valid, or we don't have a token.
            if (errorCode >= 20 && errorCode < 30) {
                adminApiIsAuthenticated.value = false;
            }

            console.error(`API Error [${errorCode}]: ${apiRes.error_message}`);
            return false;
        }

        if (token || 'token' in apiRes.content) {
            adminApiIsAuthenticated.value = true;
        }

        return apiRes.content;
    } catch (err) {
        console.error('API error:', err);

        apiCallInProgress.value = false
        return false;
    }
}

function deleteApi(path, body = null, token = null) {
    return baseApiCall('delete', path, body, token);
}

function getApi(path, token = null) {
    return baseApiCall('get', path, null, token);
}

function postApi(path, body, token = null) {
    return baseApiCall('post', path, body, token);
}

async function downloadApi(path, filename, token = null) {
    const blob = await baseApiCall('get', path, null, token, true);
    if (blob === false) {
        return false;
    }

    const url = window.URL.createObjectURL(blob);
    const link = document.createElement('a');
    link.href = url;
    link.download = filename;
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
    window.URL.revokeObjectURL(url);

    return true;
}

function uploadApi(path, file, token = null) {
    const formData = new FormData();
    formData.append('file', file);
    return baseApiCall('post', path, formData, token);
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
    // healthCheck();

    // Expose the ref for logic regarding waiting for API calls, and expose methods to make GET and POST requests
    return {apiCallInProgress, adminApiIsAuthenticated, deleteApi, getApi, postApi, downloadApi, uploadApi};
};