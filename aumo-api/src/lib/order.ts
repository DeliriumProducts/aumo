import axios from 'axios';
import { Order } from './aumo';
import { withAuth } from './axios';
import { options } from './config';

export async function placeOrder(
  order: PlaceRequest,
  cookie?: string
): Promise<Order> {
  return (
    await axios.post(`${options.Backend}/orders`, order, withAuth(cookie))
  ).data;
}

export async function getAllOrders(cookie?: string): Promise<Order[]> {
  return (await axios.get(`${options.Backend}/orders`, withAuth(cookie))).data;
}

export async function getOrder(id: string, cookie?: string): Promise<Order> {
  return (await axios.get(`${options.Backend}/orders/${id}`, withAuth(cookie)))
    .data;
}

interface PlaceRequest {
  product_id: number;
}
