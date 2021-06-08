import axios from 'axios';

export const getAxiosInstance = () => {
  const instance = axios.create({
    baseURL: process.env.API_ENDPOINT,
    headers: {
      Authorization: `Bearer ${process.env.ACCESS_TOKEN}`
    }
  });
  return instance;
}
