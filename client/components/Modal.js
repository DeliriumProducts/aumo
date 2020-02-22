import { Layout, Modal } from "@ui-kitten/components"
import React from "react"
import styled from "styled-components/native"

export default ({ visible, children, onBackdropPress = () => {} }) => (
  <Modal
    visible={visible}
    backdropStyle={{
      backgroundColor: "rgba(0, 0, 0, 0.5)"
    }}
    onBackdropPress={onBackdropPress}
  >
    <ModalContainer level="3">{children}</ModalContainer>
  </Modal>
)

const ModalContainer = styled(Layout)`
  justify-content: center;
  align-items: center;
  border-radius: 8px;
  padding: 16px;
`
