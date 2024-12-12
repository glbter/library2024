import plugin from 'tailwindcss/plugin';
import forms from '@tailwindcss/forms';
import typography from '@tailwindcss/typography';
import daisyui from 'daisyui';

/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        './go/internal/templates/*.templ',
        './static/image/*.svg',
    ],
    theme: {
        // container: {
        //   center: true,
        //   padding: {
        //     DEFAULT: "1rem",
        //     mobile: "2rem",
        //     tablet: "4rem",
        //     desktop: "5rem",
        //   },
        // },
        extend: {
            // colors: {
            //   primary: colors.blue,
            //   secondary: colors.yellow,
            //   neutral: colors.gray,
            // }
        },
    },
    plugins: [
        forms,
        typography,
        daisyui,
        plugin(({addUtilities}) => {
            addUtilities({
                /* Hide scrollbar for Chrome, Safari and Opera */
                '.no-scrollbar::-webkit-scrollbar': {
                    display: 'none',
                },
                /* Hide scrollbar for IE, Edge and Firefox */
                '.no-scrollbar': {
                    '-ms-overflow-style': 'none', // IE and Edge
                    'scrollbar-width': 'none', // Firefox
                },
            })
        }),
    ]
}