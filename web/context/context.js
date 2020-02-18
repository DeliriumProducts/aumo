import React from "react"

export const initialState = {
  user: null,
  products: [],
  shops: []
}

export const Context = React.createContext(initialState)
