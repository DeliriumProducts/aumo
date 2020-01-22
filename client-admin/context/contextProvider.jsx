import { initialState, Context } from "../context.js"
import React from "react"

const handlers = {
  setFoo(state, foo) {
    return {
      ...state,
      foo
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
