/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        p1: '#0B2D3C',
        p2: '#286e6b'
      }
    },
  },
  plugins: [],
}

