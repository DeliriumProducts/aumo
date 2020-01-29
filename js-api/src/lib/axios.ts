import { AxiosRequestConfig } from 'axios';

export const axiosRequest: AxiosRequestConfig = {
  withCredentials: true
};

export function withAuth(cookie?: string) {
  let opts = {};

  if (cookie) {
    opts = { headers: { cookie } };
  }

  return opts;
}
