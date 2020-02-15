// https://github.com/HandlebarLabs/react-native-examples-and-tutorials/blob/master/tutorials/progressive-image-loading/ProgressiveImage.js
import { Avatar } from "@ui-kitten/components"
import React from "react"
import { Animated, View } from "react-native"

const AnimatedAvatar = Animated.createAnimatedComponent(Avatar)

export default ({ fallbackSource, source, style, ...rest }) => {
  const [fallback] = React.useState(() => new Animated.Value(0))
  const [image] = React.useState(() => new Animated.Value(0))

  const handleFallbackLoad = () => {
    Animated.timing(fallback, {
      toValue: 1
    }).start()
  }

  const handleSourceLoad = () => {
    Animated.timing(image, {
      toValue: 1
    }).start()
  }

  return (
    <Container>
      <AnimatedAvatar
        {...rest}
        source={fallbackSource}
        style={[style, { opacity: fallback }]}
        onLoad={handleFallbackLoad}
        blurRadius={1}
      />
      <AnimatedAvatar
        {...rest}
        source={source}
        style={[
          {
            position: "absolute",
            left: 0,
            right: 0,
            bottom: 0,
            top: 0,
            opacity: image
          },
          style
        ]}
        onLoad={handleSourceLoad}
      />
    </Container>
  )
}

const Container = View
