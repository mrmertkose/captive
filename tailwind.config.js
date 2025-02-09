/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "internal/views/*.{html,templ}",
    "internal/views/**/*.{html,templ}",
    "internal/views/**/**/*.{html,templ}",
    "internal/views/**/**/**/*.{html,templ}"
  ]
}

