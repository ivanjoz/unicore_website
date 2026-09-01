import type { Lang } from '$lib/i18n.svelte';
import { absolute, pathFor } from '$lib/seo';

export const prerender = true;
export const trailingSlash = 'never';

const langs: Lang[] = ['es', 'en'];

/*
 * Cada URL declara sus alternativas con xhtml:link, que es lo que Google pide
 * para un sitio bilingüe: duplica lo que ya dicen las etiquetas `hreflang` del
 * <head> y sirve para las dos formas en que puede descubrir las traducciones.
 */
const alternates = langs
	.map((l) => `<xhtml:link rel="alternate" hreflang="${l}" href="${absolute(pathFor(l))}"/>`)
	.join('')
	.concat(`<xhtml:link rel="alternate" hreflang="x-default" href="${absolute(pathFor('es'))}"/>`);

export function GET() {
	const urls = langs
		.map(
			(l) =>
				`<url><loc>${absolute(pathFor(l))}</loc>${alternates}<changefreq>monthly</changefreq><priority>${l === 'es' ? '1.0' : '0.9'}</priority></url>`
		)
		.join('');

	const body = `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:xhtml="http://www.w3.org/1999/xhtml">${urls}</urlset>
`;

	return new Response(body, {
		headers: { 'content-type': 'application/xml; charset=utf-8' }
	});
}
