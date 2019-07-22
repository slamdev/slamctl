import path from 'path';
import webpack from 'webpack';
import HtmlWebpackPlugin from 'html-webpack-plugin';

const config: webpack.Configuration = {
    stats: "minimal",

    entry: './src/index',

    output: {
        path: path.join(__dirname, '/dist'),
        filename: 'bundle.js'
    },

    resolve: {
        extensions: ['.ts', '.tsx', '.js']
    },

    module: {
        rules: [
            {
                test: /\.(ts|js)x?$/,
                exclude: /node_modules/,
                use: {
                    loader: 'babel-loader'
                }
            },

            {
                test: /\.css$/,
                use: ['style-loader', 'css-loader']
            }
        ]
    },

    plugins: [
        new HtmlWebpackPlugin({
            template: './src/index.html'
        })
    ]
};

// noinspection JSUnusedGlobalSymbols
export default config;
