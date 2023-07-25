/** @type {import('tailwindcss').Config} */
export default {
	mode: 'jit',
	darkMode: 'class',
	content: ['./src/**/*.{html,js,svelte,ts}', './node_modules/flowbite-svelte/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			colors: {
				// flowbite-svelte
				primary: {
					50: '#FFF5F2',
					100: '#FFF1EE',
					200: '#FFE4DE',
					300: '#FFD5CC',
					400: '#FFBCAD',
					500: '#FE795D',
					600: '#EF562F',
					700: '#EB4F27',
					800: '#CC4522',
					900: '#A5371B'
				},
				dark: {
					50: '#C1C2C5',
					100: '#A6A7AB',
					200: '#909296',
					300: '#5C5F66',
					400: '#373A40',
					500: '#2C2E33',
					600: '#25262B',
					700: '#1A1B1E',
					800: '#141517',
					900: '#101113'
				},
				theme: {
					50: '#E3FBFF',
					100: '#D2F2FC',
					200: '#A6E3F5',
					300: '#78D3EE',
					400: '#53C4E8',
					500: '#3CBDE5',
					600: '#2BB9E4',
					700: '#18A2CB',
					800: '#0090B6',
					900: '#007DA1'
				},
				'text-theme': '#e9ecef'
			},
			textShadow: {
				sm: '0 1px 2px rgba(0, 0, 0, 0.1)',
				default: '0 2px 4px rgba(0, 0, 0, 0.1)',
				lg: '0 8px 16px rgba(0, 0, 0, 0.1)'
			},
			keyframes: {
				shine: {
					'0%': { transform: 'translateY(120%)' },
					'33%, 100%': { transform: 'translateY(-120%)' }
				}
			}
		}
	},
	plugins: [require('flowbite/plugin')]
};