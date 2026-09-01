import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

const base = process.env.BASE_PATH || '';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	preprocess: vitePreprocess(),
	kit: {
		adapter: adapter({
			pages: 'build',
			assets: 'build',
			precompress: true,
			strict: true
		}),
		paths: {
			base,
			/*
			 * Por defecto SvelteKit resuelve `base` a una ruta relativa ('.', '..'),
			 * lo que hace el sitio portable entre subdirectorios pero deja `base`
			 * inservible para componer URL absolutas: la `og:image` salía como
			 * `./images/...` —que ningún crawler resuelve— y el idioma deducido de
			 * la ruta fallaba al recortar un prefijo de longitud equivocada.
			 * El despliegue conoce su ruta (BASE_PATH), así que no necesita ser
			 * portable.
			 */
			relative: false
		}
	}
};

export default config;
