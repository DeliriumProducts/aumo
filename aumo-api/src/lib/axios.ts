import { AxiosRequestConfig } from 'axios';

export const axiosRequest: AxiosRequestConfig = {
  withCredentials: true
};

export function withAuth(cookie?: string): AxiosRequestConfig {
  let opts = {};

  if (cookie) {
    opts = { headers: { cookie } };
  }

  return {
    ...axiosRequest,
    ...opts
  };
}
