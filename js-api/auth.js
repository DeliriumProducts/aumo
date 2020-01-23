const axios = require('axios');
const opts = {
  withCredentials: true
};

class AuthAPI {
  constructor(backend) {
    this.backend = backend;
  }

  async login(creds) {
    return await axios.post(`${this.backend}/login`, creds, opts);
  }

  async register(creds) {
    return await axios.post(`${this.backend}/register`, creds);
  }

  async logout() {
    return await axios.get(`${this.backend}/logout`, opts);
  }

  async me() {
    return await axios.get(`${this.backend}/me`, opts);
  }
}

module.exports = AuthAPI;
