import type { Handle } from '@sveltejs/kit';
import { langFromPath } from '$lib/seo';

/*
 * El atributo `lang` de <html> no puede fijarse desde <svelte:head>, así que
 * app.html lleva un marcador `%lang%` y se sustituye aquí, al prerenderizar
 * cada ruta. Antes estaba escrito a mano como "es" y la versión en inglés se
 * anunciaba como española a buscadores y lectores de pantalla.
 */
export const handle: Handle = ({ event, resolve }) =>
	resolve(event, {
		transformPageChunk: ({ html }) => html.replaceAll('%lang%', langFromPath(event.url.pathname))
	});
