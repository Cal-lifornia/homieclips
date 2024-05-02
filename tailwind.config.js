/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./components/*.templ'],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  corePlugins: {
    preflight: true,
  },
  daisyui: {
    themes: ["light","dark"],
    },
}

