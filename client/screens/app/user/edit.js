import { Button, Icon, Layout } from "@ui-kitten/components"
import React from "react"
import { useForm } from "react-hook-form"
import { View } from "react-native"
import styled from "styled-components/native"
import Avatar from "../../../components/Avatar"
import ErrorContainer from "../../../components/ErrorContainer"
import FormInput from "../../../components/FormInput"
import { Context } from "../../../context/context"
import theme from "../../../theme"

export default () => {
  const ctx = React.useContext(Context)
  const { register, handleSubmit, setValue, errors } = useForm()
  const [loading, setLoading] = React.useState(false)
  const [passwordVisible, setPasswordVisible] = React.useState(false)
  const [err, setErr] = React.useState("")

  const onPasswordIconPress = () => {
    setPasswordVisible(!passwordVisible)
  }

  const handleEdit = async () => {}

  return (
    <Container>
      <AvatarContainer>
        <Avatar
          size="giant"
          source={{ uri: ctx?.state?.user?.avatar }}
          fallbackSource={require("../../../assets/Avatar.png")}
          style={{ height: 128, aspectRatio: 1.0, alignSelf: "center" }}
        />
      </AvatarContainer>
      <Layout level="1" style={{ marginTop: 25, padding: 25 }}>
        <FormInput
          status={errors.name ? "danger" : "basic"}
          placeholder="Name"
          size="medium"
          defaultValue={ctx.state.user?.name}
          icon={style => <Icon {...style} name="person-outline" />}
          ref={register("name", { required: "Required" })}
          onChangeText={val => setValue("name", val)}
          caption={errors.name?.message}
        />
        {err != "" && <ErrorContainer error={err} />}
      </Layout>
      <View style={{ paddingHorizontal: 25 }}>
        <Button
          disabled={loading}
          icon={style => <Icon name="edit-outline" {...style} />}
          size="large"
          style={{ width: "100%", marginTop: 10, borderRadius: 10 }}
          onPress={handleSubmit(handleEdit)}
        >
          DONE
        </Button>
      </View>
    </Container>
  )
}

const Container = styled(View)`
  background-color: ${theme["color-background-main"]};
  flex: 1;
  height: 100%;
`

const AvatarContainer = styled(View)`
  aspect-ratio: 1;
  height: 128px;
  margin-top: 25px;
  align-self: center;
`
