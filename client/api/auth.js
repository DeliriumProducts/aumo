import axios from "axios"
import { BACKEND_URL } from "../config"

export class AuthAPI {
  static opts = {
    withCredentials: true
  }

  static async login(creds) {
    return await axios.post(`${BACKEND_URL}/login`, creds, this.opts)
  }

  static async register(creds) {
    return await axios.post(`${BACKEND_URL}/register`, creds)
  }

  static async logout() {
    return await axios.get(`${BACKEND_URL}/logout`, this.opts)
  }

  static async me() {
    return await axios.get(`${BACKEND_URL}/me`, this.opts)
  }
}
