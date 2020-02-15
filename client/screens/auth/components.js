import { Input, Text } from "@ui-kitten/components"
import { Image, KeyboardAvoidingView, View } from "react-native"
import styled from "styled-components/native"
import theme from "../../theme"

export const Container = styled(KeyboardAvoidingView)`
  background-color: ${theme["color-background-main"]};
  flex: 1;
  height: 100%;
  justify-content: space-between;
`

export const MainContainer = styled(View)`
  align-items: center;
  margin-top: 10px;
  margin-bottom: 20px;
`

export const Form = styled(View)`
  padding-horizontal: 32px;
`

export const Aumo = styled(Image)`
  width: 220px;
  resize-mode: contain;
  margin-bottom: -20px;
`

export const Subheading = styled(Text)`
  font-size: 17px;
  color: ${theme["color-primary-500"]};
  margin-bottom: 20px;
  text-align: center;
`

export const FormInput = styled(Input)`
  border-radius: 10px;
`
