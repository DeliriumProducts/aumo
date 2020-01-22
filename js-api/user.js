const axios = require("axios")

class UserAPI {
  static opts = {
    withCredentials: true
  }

  constructor(backend) {
    this.backend = backend
  }

  async getAll() {
    return await axios.get(`${this.backend}/users`, this.opts)
  }

  async get(id) {
    return await axios.get(`${this.backend}/users/${id}`, this.opts)
  }

  async setRole(id, role) {
    return await axios.put(
      `${this.backend}/users/${id}/set-role`,
      { role: role },
      this.opts
    )
  }

  async addPoints(id, points) {
    return await axios.put(
      `${this.backend}/users/${id}/add-points`,
      { points: points },
      this.opts
    )
  }

  async subPoints(id, points) {
    return await axios.put(
      `${this.backend}/users/${id}/sub-points`,
      { points: points },
      this.opts
    )
  }

  async delete(id) {
    return await axios.delete(`${this.backend}/users/${id}`, this.opts)
  }
}

export default UserAPI
