/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./**/*.{html,templ}", "./index.html"],
    theme: {
        extend: {
            colors: {
                primary: '#003452',
                lightPrimary: '#0067a3',
                secondary: '#bccdd6',
                accent: '#fabb00',
                lightAccent: '#f5d26c',
                light: '#ffecb3',
                error: '#ea5353',
                warn: '#fdf7dd',
                warn2: '#ca991c',
                grey: '#f2f2f2',
                greyDeep: '#dedede',
                dark: '#919191',
                darker: '#666565',
                black: '#000000',
                white: '#ffffff',
                header: '#f6f8f9'
            },
        },
        fontFamily: {
            'sans': ['Ubuntu', 'sans-serif'],
            'serif': ['Ubuntu', 'serif'],
            'mono': ['Ubuntu', 'monospace'],
            'poppins': ['Poppins', 'sans-serif'],
        },
    },
    plugins: [],
}


