const axios = require("axios")
const opts = {
  withCredentials: true
}

class UserAPI {
  constructor(backend) {
    this.backend = backend
  }

  async getAll() {
    return (await axios.get(`${this.backend}/users`, opts)).data
  }

  async get(id) {
    return (await axios.get(`${this.backend}/users/${id}`, opts)).data
  }

  async setRole(id, role) {
    return (
      await axios.put(
        `${this.backend}/users/${id}/set-role`,
        { role: role },
        opts
      )
    ).data
  }

  async addPoints(id, points) {
    return (
      await axios.put(
        `${this.backend}/users/${id}/add-points`,
        { points: points },
        opts
      )
    ).data
  }

  async subPoints(id, points) {
    return (
      await axios.put(
        `${this.backend}/users/${id}/sub-points`,
        { points: points },
        opts
      )
    ).data
  }

  async delete(id) {
    return (await axios.delete(`${this.backend}/users/${id}`, opts)).data
  }
}

module.exports = UserAPI
