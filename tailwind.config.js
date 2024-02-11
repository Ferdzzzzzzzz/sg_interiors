/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.templ"],
  theme: {
    fontFamily: {
      sans: ['Thasadith', 'system-ui'],
      serif: ['Prata', 'ui-sans-serif']
    },
    colors: {
      white: "#ffffff",
      'primary-light': "#EFEBE5",
      'primary-md': "#CBC7B7",
      'primary-dark': "#57534e",
      'primary-400': "#a8a29e",
      'primary-500': "#78716c",
      warn: "#991b1b"
    },
    extend: {},
  },
  plugins: [],
}

