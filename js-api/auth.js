const axios = require("axios")
const opts = {
  withCredentials: true
}

class AuthAPI {
  constructor(backend) {
    this.backend = backend
  }

  async login(creds) {
    return await axios.post(`${this.backend}/login`, creds, opts)
  }

  async register(creds) {
    return await axios.post(`${this.backend}/register`, creds)
  }

  async logout() {
    return await axios.get(`${this.backend}/logout`, opts)
  }

  async me(cookie) {
    let opts = {}

    if (cookie) {
      opts = { headers: { cookie } }
    }

    const auth = await axios.get(`${this.backend}/me`, {
      withCredentials: true,
      ...opts
    }).data

    return auth
  }
}

module.exports = AuthAPI
