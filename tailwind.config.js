/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.templ"],
  theme: {
    fontFamily: {
      sans: ['Thasadith', 'system-ui'],
      serif: ['Prata', 'ui-sans-serif']
    },
    colors: {
      'primary-light': "#EFEBE5",
      'primary-md': "#CBC7B7",
      'primary-dark': "#848175",
      warn: "#991b1b"
    },
    extend: {},
  },
  plugins: [],
}

