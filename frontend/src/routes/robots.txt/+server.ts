import { SITE_URL, absolute } from '$lib/seo';

export const prerender = true;

/*
 * Va como endpoint y no como archivo en static/ para que el dominio salga de
 * un único sitio (SITE_URL) y no haya que acordarse de tocar dos archivos si
 * cambia. `trailingSlash` se fuerza aquí porque el layout usa 'always' y
 * /robots.txt/ no lo reconoce ningún crawler.
 */
export const trailingSlash = 'never';

export function GET() {
	const body = `User-agent: *
Allow: /

Sitemap: ${absolute('/sitemap.xml')}
Host: ${SITE_URL}
`;

	return new Response(body, {
		headers: { 'content-type': 'text/plain; charset=utf-8' }
	});
}
