import axios from "axios"
import { BACKEND_URL } from "../config"

export class OrderAPI {
  static opts = {
    withCredentials: true
  }

  static async placeOrder(productID) {
    return await axios.post(
      `${BACKEND_URL}/orders`,
      { product_id: productID },
      this.opts
    )
  }
}
