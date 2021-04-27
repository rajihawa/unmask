import axios, { AxiosInstance } from "axios";

export function useFetch(): AxiosInstance {
  const fetch = axios.create({
    baseURL:
      process.env.NODE_ENV === "development"
        ? "http://localhost:4000"
        : undefined,
    withCredentials: true,
    timeout: 3000,
  });
  return fetch;
}
