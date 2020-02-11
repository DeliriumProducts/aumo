import React from "react"
import { Context, initialState } from "../context.js"

export const actions = {
  SET_USER: "SET_USER"
}

const handlers = {
  [actions.SET_USER]: (state, user) => {
    return {
      ...state,
      user
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
