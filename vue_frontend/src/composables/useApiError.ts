import { ref } from "vue";

export function useApiError() {
    const errorMessage = ref("");

    function handleApiError(err: any) {
        if (err.response) {
            // Backend returned an error response
            errorMessage.value =
                err.response.data?.message ||
                err.response.data?.error ||
                "An unexpected server error occurred.";
        } else if (err.request) {
            // Request made but no answer (server down, CORS, network)
            errorMessage.value = "Cannot reach server. Please try again later.";
        } else {
            // Any other error
            errorMessage.value = "Unexpected error. Please try again.";
        }
    }

    return { errorMessage, handleApiError };
}
