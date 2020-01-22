const axios = require("axios")

class AuthAPI {
  static opts = {
    withCredentials: true
  }

  constructor(backend) {
    this.backend = backend
  }

  async login(creds) {
    return await axios.post(`${this.backend}/login`, creds, this.opts)
  }

  async register(creds) {
    return await axios.post(`${this.backend}/register`, creds)
  }

  async logout() {
    return await axios.get(`${this.backend}/logout`, this.opts)
  }

  async me() {
    return await axios.get(`${this.backend}/me`, this.opts)
  }
}

export default AuthAPI
