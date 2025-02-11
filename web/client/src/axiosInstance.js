import axios from "axios";

const axiosInstance = axios.create({
    baseURL: "http://localhost/auth/",
    withCredentials: true,
});

axiosInstance.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem("accessToken");
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    (error) => Promise.reject(error)
);

axiosInstance.interceptors.response.use(
    (response) => response,
    async (error) => {
        const originalRequest = error.config;

        if (error.response && error.response.status === 401 && !originalRequest._retry) {
            originalRequest._retry = true;

            try {
                const refreshResponse = await axios.post("http://localhost/auth/users/refresh", {}, { withCredentials: true });

                const newAccessToken = refreshResponse.data.accessToken;
                localStorage.setItem("accessToken", newAccessToken);

                originalRequest.headers.Authorization = `Bearer ${newAccessToken}`;
                return axiosInstance(originalRequest);
            } catch (refreshError) {
                console.error("Ошибка обновления токена:", refreshError);
                localStorage.removeItem("accessToken");
            }
        }

        return Promise.reject(error);
    }
);

export default axiosInstance;
