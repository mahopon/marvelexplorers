import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  build: {
    rollupOptions: {
      output: {
        assetFileNames: (assetInfo) => {
          let extType = assetInfo.name!.split('.').at(1);
          if (/png|jpe?g|svg|gif|tiff|bmp|ico/i.test(extType!)) {
            extType = 'img';
          }
          return `marvelexplorers/assets/${extType}/[name]-[hash][extname]`;
        },
        chunkFileNames: 'marvelexplorers/assets/js/[name]-[hash].js',
        entryFileNames: 'marvelexplorers/assets/js/[name]-[hash].js',
      }
    }
  }
})
