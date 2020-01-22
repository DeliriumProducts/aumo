const withLess = require("@zeit/next-less")
const withSass = require("@zeit/next-sass")
const withOffline = require("next-offline")
const webpack = require("webpack")
const TerserPlugin = require("terser-webpack-plugin")
const OptimizeCSSAssetsPlugin = require("optimize-css-assets-webpack-plugin")
const fs = require("fs")
const path = require("path")
const lessToJS = require("less-vars-to-js")
const Dotenv = require("dotenv-webpack")

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
      webpack(config) {
        // Fixes npm packages that depend on `fs` module
        config.node = {
          fs: "empty"
        }

        config.plugins.push(
          new Dotenv({
            path: path.join(__dirname, "/../.env"),
            systemvars: true
          })
        )

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

          config.plugins.push(new OptimizeCSSAssetsPlugin())
        }

        return config
      },
      lessLoaderOptions: {
        javascriptEnabled: true,
        modifyVars: themeVariables
      },
      env: {
        THEME_VARIABLES: themeVariables
      }
    })
  )
)
