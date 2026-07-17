import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import mkcert from 'vite-plugin-mkcert'

// https://vite.dev/config/
export default defineConfig({
  publicDir: 'public',
  plugins: [react(), mkcert({
    savePath: 'public/certificates'
  })],
  server: {
    port: 8000,
    host: true,
    open: true,
    proxy: {
      '/api': {
        target: 'https://localhost:8080',
        changeOrigin: true,
        secure: false,
      },
    },
  },
  build: {
    outDir: 'dist',
  },
  resolve: {
    alias: {
      // '@': path.resolve(__dirname, 'src'),
    },
  },
});
