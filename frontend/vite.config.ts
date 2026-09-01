import { sveltekit } from '@sveltejs/kit/vite';
import tailwindcss from '@tailwindcss/vite';
import { defineConfig } from 'vite';
import { svgIcons } from './scripts/svg-icons';

export default defineConfig({
	plugins: [svgIcons(), tailwindcss(), sveltekit()]
});
