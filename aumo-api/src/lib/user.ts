import axios from 'axios';
import { MessageResponse, Role, User } from './aumo';
import { withAuth } from './axios';
import { options } from './config';

export async function getAllUsers(cookie?: string): Promise<User[]> {
  return (await axios.get(`${options.Backend}/users`, withAuth(cookie))).data;
}

export async function getUser(id: string, cookie?: string): Promise<User> {
  return (await axios.get(`${options.Backend}/users/${id}`, withAuth(cookie)))
    .data;
}

export async function setRole(
  id: string,
  role: Role,
  cookie?: string
): Promise<MessageResponse> {
  return (
    await axios.put(
      `${options.Backend}/users/${id}/set-role`,
      { role: role },
      withAuth(cookie)
    )
  ).data;
}

export async function addPoints(
  id: string,
  points: number,
  cookie?: string
): Promise<MessageResponse> {
  return (
    await axios.put(
      `${options.Backend}/users/${id}/add-points`,
      { points: points },
      withAuth(cookie)
    )
  ).data;
}

export async function subPoints(
  id: string,
  points: number,
  cookie?: string
): Promise<MessageResponse> {
  return (
    await axios.put(
      `${options.Backend}/users/${id}/sub-points`,
      { points: points },
      withAuth(cookie)
    )
  ).data;
}

export async function deleteUser(
  id: string,
  cookie?: string
): Promise<MessageResponse> {
  return (
    await axios.delete(`${options.Backend}/users/${id}`, withAuth(cookie))
  ).data;
}

export default {
  getAllUsers,
  getUser,
  deleteUser,
  subPoints,
  addPoints,
  setRole
};
