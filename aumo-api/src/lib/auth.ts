import axios from 'axios';
import { MessageResponse, User } from './aumo';
import { axiosRequest, withAuth } from './axios';
import { options } from './config';

export async function login(creds: LoginRequest): Promise<User> {
  return (await axios.post(`${options.Backend}/login`, creds, axiosRequest))
    .data;
}

export async function register(
  creds: RegisterRequest
): Promise<MessageResponse> {
  return (await axios.post(`${options.Backend}/register`, creds, axiosRequest))
    .data;
}

export async function logout(cookie?: string): Promise<MessageResponse> {
  return (await axios.get(`${options.Backend}/logout`, withAuth(cookie))).data;
}

export async function me(cookie?: string): Promise<User> {
  return (await axios.get(`${options.Backend}/me`, withAuth(cookie))).data;
}

interface LoginRequest {
  email: string;
  password: string;
}

interface RegisterRequest {
  name: string;
  email: string;
  password: string;
  avatar: string;
}
