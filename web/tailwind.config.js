const colors = require("tailwindcss/colors");
const daisyBase = require("daisyui/dist/base");

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
  daisyui: {
    themes: [
      {
        light: {
          ...require("daisyui/src/colors/themes")["[data-theme=light]"],
          primary: colors.purple[500],
          "primary-focus": colors.purple[600],
        },
      },
      {
        dark: {
          ...require("daisyui/src/colors/themes")["[data-theme=dark]"],
          primary: colors.purple[300],
          "primary-focus": colors.purple[400],
          "primary-content": "rgb(31, 36, 45)",
        },
      },
    ],
  },
};
