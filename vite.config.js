import preactPlugin from '@preact/preset-vite'
import { defineConfig } from 'vite'

export default defineConfig({
    root: 'frontend/',
    build: {
        outDir: '../out/frontend/',
    },
    server: {
        port: 3000,
        proxy: {
            '/api': 'http://127.0.0.1:4000/',
        },
    },
    plugins: [preactPlugin()],
})
