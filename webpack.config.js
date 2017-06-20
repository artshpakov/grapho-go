const NODE_ENV = process.env.NODE_ENV || 'development';
const webpack = require('webpack');

const ExtractTextPlugin = require("extract-text-webpack-plugin");

module.exports = {
  entry: './src/assets/app',
  output: {
    path: __dirname+'/public',
    filename: 'app.js'
  },

  module: {
    loaders: [{
      test: /\.jsx?$/,
      exclude: /node_modules/,
      loader: 'babel-loader',
      query: {
        presets: ['es2015', 'react']
      }
    }, {
      test: /\.(css|scss|sass)$/,
      loader: ExtractTextPlugin.extract('css-loader!sass-loader')
    }]
  },

  amd: { jQuery: true },

  plugins: [
    new webpack.ProvidePlugin({
      '$': 'jquery',
      'jQuery': 'jquery',
      '_': 'underscore',
      'React': 'react'
    }),
    new ExtractTextPlugin("app.css")
  ],

  watch: NODE_ENV == 'development',
  devtool: 'source-map'
}
