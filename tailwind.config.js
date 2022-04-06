const colors = require('tailwindcss/colors')
const defaultTheme = require('tailwindcss/defaultTheme')

module.exports = {
  content: [
    './views/*.html'
  ],
  theme: {
    extend: {
      colors: {
        transparent: 'transparent',
        current: 'currentColor',
        black: colors.black,
        white: colors.white,
        'neutral': {
          50: '#fafafa',
          100: '#f5f5f5',
          200: '#e5e5e5',
        },
      },
    },
    fontFamily: {
      'sans': ['Roboto', 'sans-serif'],
    },
    screens: {
      'xxs': '1px',
      'xs': '475px',
      ...defaultTheme.screens,
    },
  },
  plugins: [],
}