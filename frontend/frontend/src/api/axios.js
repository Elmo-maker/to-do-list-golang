import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:8080", // backend Go
  withCredentials: true,
});

export default api;
