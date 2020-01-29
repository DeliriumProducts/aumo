import axios from 'axios';
import { BaseResponse, Order } from './aumo';
import { withAuth } from './axios';
import { options } from './config';

export async function place(
  order: PlaceRequest,
  cookie?: string
): Promise<BaseResponse<Order>> {
  return (
    await axios.post(`${options.Backend}/orders`, order, withAuth(cookie))
  ).data;
}

export async function getAll(cookie?: string): Promise<BaseResponse<Order[]>> {
  return (await axios.get(`${options.Backend}/orders`, withAuth(cookie))).data;
}

export async function get(id: number, cookie?: string) {
  return (await axios.get(`${options.Backend}/orders/${id}`, withAuth(cookie)))
    .data;
}

interface PlaceRequest {
  product_id: number;
}

export default {
  place,
  getAll,
  get
};
