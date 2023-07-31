import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig, loadEnv } from 'vite';

export default ({mode}: {mode: any}) => {
	const env = loadEnv(mode, process.cwd(), '')
	return defineConfig({
		plugins: [sveltekit()],
		server: {
			port: parseInt(env.PORT) || 3000
		}
	});
}