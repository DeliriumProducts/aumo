import axios from 'axios';
import { User } from './aumo';
import { axiosRequest } from './axios';
import { options } from './config';

export async function login(creds: LoginRequest): Promise<LoginResponse> {
  return (await axios.post(`${options.Backend}/login`, creds, axiosRequest))
    .data;
}

export async function register(
  creds: RegisterRequest
): Promise<RegisterResponse> {
  return (await axios.post(`${options.Backend}/register`, creds, axiosRequest))
    .data;
}

interface LoginRequest {
  email: string;
  password: string;
}

interface LoginResponse extends User {}

interface RegisterRequest {
  name: string;
  email: string;
  password: string;
  avatar: string;
}

interface RegisterResponse extends User {}

export default {
  login,
  register
};
