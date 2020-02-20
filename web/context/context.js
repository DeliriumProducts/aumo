import React from "react"

export const initialState = {
  user: null,
  products: []
}

export const Context = React.createContext(initialState)
