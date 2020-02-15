import { Text } from "@ui-kitten/components"
import React from "react"
import { View } from "react-native"
import styled from "styled-components/native"
import theme from "../theme"

export default ({ error }) => (
  <ErrorContainer>
    <Text style={{ color: "white" }}>{error}</Text>
  </ErrorContainer>
)

const ErrorContainer = styled(View)`
  border-radius: 8px;
  padding: 15px;
  width: 100%;
  background-color: ${theme["color-danger-500"]};
`
