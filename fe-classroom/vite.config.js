import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@Public': '/public',
      '@PublicAssets': '/public/assets',
      '@Src': '/src',
      '@Components': '/src/components',
      '@Layout': '/src/layout',
    }
  }
})
