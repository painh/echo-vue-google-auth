import {defineConfig} from 'vite'
import mkcert from 'vite-plugin-mkcert'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue(), mkcert()],
    build: {
        // outDir: '../backend/static/',
        emptyOutDir: true,
    }
})
