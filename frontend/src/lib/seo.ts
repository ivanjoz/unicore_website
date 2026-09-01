/*
 * Un único sitio en dos idiomas, cada uno con su URL propia: `/` en español y
 * `/en/` en inglés. Antes el idioma vivía sólo en el cliente, así que Google
 * únicamente veía la versión española; con dos URL prerenderizadas y `hreflang`
 * recíproco cada buscador indexa la que le corresponde.
 *
 * Aquí vive todo lo que necesita saber la URL absoluta del sitio, porque los
 * crawlers de Open Graph no resuelven rutas relativas: una `og:image` como
 * `./images/foo.jpg` deja la previsualización de WhatsApp sin imagen.
 */
import { base } from '$app/paths';
import type { Lang } from './i18n.svelte';

/*
 * El origen no se puede deducir en tiempo de build (el sitio es estático y se
 * prerenderiza sin saber dónde va a servirse), así que viaja como variable de
 * build igual que `VITE_API_URL`. El valor por defecto es el dominio de
 * producción, el mismo que el backend acepta en `allowed_origins`.
 */
const DEFAULT_SITE_URL = 'https://un.pe';

export const SITE_URL = (import.meta.env.VITE_SITE_URL || DEFAULT_SITE_URL).replace(/\/+$/, '');

export const SITE_NAME = 'Unicore Labs';

/** Ruta interna de cada idioma, con el `base` del despliegue ya aplicado. */
export function pathFor(lang: Lang): string {
	return lang === 'es' ? `${base}/` : `${base}/en/`;
}

/** URL absoluta a partir de una ruta que ya incluye `base`. */
export function absolute(path: string): string {
	return `${SITE_URL}${path.startsWith('/') ? path : `/${path}`}`;
}

/**
 * El idioma se deduce del segmento `/en` de la ruta. Es la única fuente de
 * verdad: así el HTML prerenderizado, la hidratación y la navegación entre
 * idiomas coinciden siempre, sin depender de `localStorage`.
 */
export function langFromPath(pathname: string): Lang {
	return /(^|\/)en(\/|$)/.test(pathname.slice(base.length)) ? 'en' : 'es';
}

type Meta = {
	title: string;
	description: string;
	ogDescription: string;
	ogImage: string;
	ogImageAlt: string;
	locale: string;
};

/*
 * Los textos de `<head>` no usan el separador `es|en` de i18n.svelte: aquí no
 * se elige el idioma en runtime, cada URL emite su bloque completo y fijo.
 */
export const meta: Record<Lang, Meta> = {
	es: {
		title: 'Unicore Labs | Código abierto de alto impacto',
		description:
			'Iniciativas de código abierto que democratizan el acceso a tecnología de alto impacto. Sistemas y herramientas que incentivan la participación comunitaria y la innovación tecnológica.',
		ogDescription:
			'Desarrollamos sistemas y herramientas de código abierto que incentivan la participación comunitaria y la innovación tecnológica.',
		ogImage: 'og-card-es.jpg',
		ogImageAlt: 'Unicore Labs — Código abierto de alto impacto',
		locale: 'es_PE'
	},
	en: {
		title: 'Unicore Labs | High-impact open source',
		description:
			'Open-source initiatives that democratize access to high-impact technology. Systems and tools that encourage community participation and technological innovation.',
		ogDescription:
			'We build open-source systems and tools that encourage community participation and technological innovation.',
		ogImage: 'og-card-en.jpg',
		ogImageAlt: 'Unicore Labs — High-impact open source',
		locale: 'en_US'
	}
};

/** Las tarjetas de compartir se generan a 1200x630 con scripts/og-image.sh. */
export const OG_IMAGE_WIDTH = 1200;
export const OG_IMAGE_HEIGHT = 630;
