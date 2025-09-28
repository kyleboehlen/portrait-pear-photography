/** @type {import('tailwindcss').Config} */
// eslint-disable-next-line no-undef
const defaultTheme = require("tailwindcss/defaultTheme")
// eslint-disable-next-line no-undef
module.exports = {
  purge: ["./index.html", "./src/**/*.{vue,js}"],
  darkMode: false, // or 'media' or 'class'
  content: [],
  theme: {
    extend: {},
    screens: {
      xs: "450px",
      ...defaultTheme.screens,
    },
  },
  daisyui: {
    themes: [
      {
        pear: {
          primary: "#3db084",
          secondary: "#012217",
          accent: "#4278A9",
          neutral: "#4f5256",
          "base-100": "#222325",
          info: "#4278A9",
          success: "#3eb389",
          warning: "#FFBA58",
          error: "#DB2D5C",
        },
      },
    ],
  },
  // eslint-disable-next-line no-undef
  plugins: [require("daisyui")],
}
