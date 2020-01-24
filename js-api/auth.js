const axios = require("axios")
const opts = {
  withCredentials: true
}

class AuthAPI {
  constructor(backend) {
    this.backend = backend
  }

  async login(creds) {
    return (await axios.post(`${this.backend}/login`, creds, opts)).data
  }

  async register(creds) {
    return (await axios.post(`${this.backend}/register`, creds)).data
  }

  async logout() {
    return (await axios.get(`${this.backend}/logout`, opts)).data
  }

  async me(cookie) {
    let opts = {}

    if (cookie) {
      opts = { headers: { cookie } }
    }

    return (
      await axios.get(`${this.backend}/me`, {
        withCredentials: true,
        ...opts
      })
    ).data
  }
}

module.exports = AuthAPI
