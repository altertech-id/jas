/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			colors: {
				navy: { DEFAULT: '#0e304f', deep: '#081d31' },
				brand: {
					blue: '#2284c7',
					'blue-hover': '#196fa8',
					'blue-tint': '#f0f7fc',
					red: '#d91b1b',
					'red-hover': '#b81414',
					'red-active': '#940f0f'
				},
				cta: { DEFAULT: '#d91b1b', hover: '#b81414', active: '#940f0f' },
				amber: { DEFAULT: '#f59e0b', 500: '#f59e0b' },
				surface: { light: '#f9fbfd', mist: '#f1f5f9' },
				ink: { 900: '#0f172a', 700: '#334155', 500: '#64748b', 300: '#94a3b8' },
				line: { soft: '#f1f5f9', strong: '#e2e8f0' },
				wa: { DEFAULT: '#25d366', hover: '#1ebd5a' }
			},
			fontFamily: {
				heading: ['"Plus Jakarta Sans"', '-apple-system', 'BlinkMacSystemFont', '"Segoe UI"', 'sans-serif'],
				body: ['Roboto', '-apple-system', 'BlinkMacSystemFont', '"Segoe UI"', 'sans-serif']
			},
			maxWidth: { container: '1180px', article: '720px' },
			boxShadow: {
				card: '0 4px 20px -2px rgba(15, 23, 42, 0.05)',
				float: '0 10px 30px -4px rgba(14, 48, 79, 0.1)'
			}
		}
	},
	plugins: []
};
