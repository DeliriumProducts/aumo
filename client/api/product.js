import axios from "axios"
import { BACKEND_URL } from "../config"

export class ProductAPI {
  static opts = {
    withCredentials: true
  }

  static async getAll() {
    return await axios.get(`${BACKEND_URL}/products`, this.opts)
  }
}
