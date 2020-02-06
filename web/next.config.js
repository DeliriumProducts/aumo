const withLess = require("@zeit/next-less")
const withSass = require("@zeit/next-sass")
const withOffline = require("next-offline")
const webpack = require("webpack")
const TerserPlugin = require("terser-webpack-plugin")
const OptimizeCSSAssetsPlugin = require("optimize-css-assets-webpack-plugin")
const fs = require("fs")
const path = require("path")
const lessToJS = require("less-vars-to-js")

require("dotenv").config({
  path: path.join(__dirname, "/../.env")
})

require("aumo").config({
  Backend: process.env.BACKEND_URL
})

const themeVariables = lessToJS(
  fs.readFileSync(path.join(__dirname, "./assets/antd-custom.less"), "utf8")
)

// fix: prevents error when .css or .less files are required by node
if (typeof require !== "undefined") {
  require.extensions[".css"] = file => {}
  require.extensions[".less"] = file => {}
}

module.exports = withOffline(
  withLess(
    withSass({
      webpack(config, { isServer }) {
        config.node = {
          fs: "empty"
        }

        if (isServer) {
          const antStyles = /antd\/.*?\/style.*?/
          const origExternals = [...config.externals]
          config.externals = [
            (context, request, callback) => {
              if (request.match(antStyles)) return callback()
              if (typeof origExternals[0] === "function") {
                origExternals[0](context, request, callback)
              } else {
                callback()
              }
            },
            ...(typeof origExternals[0] === "function" ? [] : origExternals)
          ]

          config.module.rules.unshift({
            test: antStyles,
            use: "null-loader"
          })
        }

        if (process.env.NODE_ENV === "production") {
          config.plugins = config.plugins.filter(
            plugin => plugin.constructor.name !== "UglifyJsPlugin"
          )

          config.plugins.push(
            new TerserPlugin({
              parallel: true,
              terserOptions: {
                ecma: 6
              }
            })
          )
        }

        config.plugins.push(new OptimizeCSSAssetsPlugin())

        return config
      },
      lessLoaderOptions: {
        javascriptEnabled: true,
        modifyVars: themeVariables
      },
      env: {
        THEME_VARIABLES: themeVariables,
        BACKEND_URL: process.env.BACKEND_URL
      }
    })
  )
)
