/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      keyframes: {
        bounce: {
          '0%, 100%': { transform: 'translateY(0)' },
          '50%': { transform: 'translateY(-6px)' },
        },
        'bounce2': {
          '0%, 100%': { transform: 'translateY(6)' },
          '50%': { transform: 'translateY(-3px)' },
        },
      },
      animation: {
        'bounce': 'bounce 1s infinite',
        'bounce2': 'bounce2 1s infinite',
      },
    },
  },
  plugins: [],
}

