import { Button, Icon, Layout, Spinner } from "@ui-kitten/components"
import aumo from "aumo"
import React from "react"
import { useForm } from "react-hook-form"
import { View } from "react-native"
import styled from "styled-components/native"
import Avatar from "../../../components/Avatar"
import ErrorContainer from "../../../components/ErrorContainer"
import FormInput from "../../../components/FormInput"
import { Context } from "../../../context/context"
import theme from "../../../theme"

export default ({ navigation }) => {
  const ctx = React.useContext(Context)
  const { register, handleSubmit, setValue, errors } = useForm({
    defaultValues: {
      name: ctx.state.user?.name
    }
  })
  const [loading, setLoading] = React.useState(false)
  const [err, setErr] = React.useState("")

  const handleEdit = async data => {
    try {
      setLoading(true)
      await aumo.user.edit(data)
      navigation.goBack()
    } catch (error) {
      switch (error.response.status) {
        case 400:
          setErr("Bad Request")
          break
        case 500:
          setErr("Internal server error")
          break
      }
    } finally {
      setLoading(false)
    }
  }

  return (
    <Container>
      <AvatarContainer>
        <Avatar
          size="giant"
          source={{ uri: ctx.state.user?.avatar }}
          fallbackSource={require("../../../assets/Avatar.png")}
          style={{ height: 128, aspectRatio: 1.0, alignSelf: "center" }}
        />
      </AvatarContainer>
      <Layout level="1" style={{ marginTop: 25, padding: 25 }}>
        <FormInput
          status={errors.name ? "danger" : "basic"}
          placeholder="Name"
          size="medium"
          disabled={loading}
          defaultValue={ctx.state.user?.name}
          icon={style => <Icon {...style} name="person-outline" />}
          ref={register("name", { required: "Required" })}
          onChangeText={val => setValue("name", val)}
          caption={errors.name?.message}
        />
        {err != "" && <ErrorContainer error={err} />}
      </Layout>
      <View
        style={{
          paddingHorizontal: 25,
          justifyContent: "center",
          alignItems: "center"
        }}
      >
        <Button
          disabled={loading}
          icon={style => <Icon name="edit-outline" {...style} />}
          size="large"
          style={{ width: "100%", marginTop: 10, borderRadius: 10 }}
          onPress={handleSubmit(handleEdit)}
        >
          DONE
        </Button>
        <View style={{ marginTop: 10 }}>
          {loading && <Spinner size="giant" />}
        </View>
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
