import React from "react"
import { Context, initialState } from "../context.js"

export const actions = {
  SET_USER: "SET_USER",
  SET_PRODUCTS: "SET_PRODUCTS"
}

const handlers = {
  [actions.SET_USER]: (state, user) => {
    return {
      ...state,
      user
    }
  },
  [actions.SET_PRODUCTS]: (state, products) => {
    return {
      ...state,
      products
    }
  }
}

const contextReducer = (state, { type, payload }) => {
  if (handlers[type]) {
    return handlers[type](state, payload)
  } else {
    return state
  }
}

const ContextProvider = ({ children }) => {
  const [state, dispatch] = React.useReducer(contextReducer, initialState)
  return (
    <Context.Provider value={{ state, dispatch }}>{children}</Context.Provider>
  )
}

export default ContextProvider
