<script lang="ts">
	import Header from '$lib/components/Header.svelte';
	import Footer from '$lib/components/Footer.svelte';
	import { base } from '$app/paths';
	import { page } from '$app/state';
	import { lang } from '$lib/i18n.svelte';
	import {
		OG_IMAGE_HEIGHT,
		OG_IMAGE_WIDTH,
		SITE_NAME,
		SITE_URL,
		absolute,
		langFromPath,
		meta,
		pathFor
	} from '$lib/seo';
	import './global.css';

	let { children } = $props();

	const current = $derived(langFromPath(page.url.pathname));
	const other = $derived(current === 'es' ? 'en' : 'es');
	const head = $derived(meta[current]);
	const canonical = $derived(absolute(pathFor(current)));

	/*
	 * El idioma se fija de forma síncrona, en el cuerpo del script, porque el
	 * layout se renderiza antes que la página y `t()` tiene que devolver ya el
	 * idioma correcto. Un `$effect` no serviría: no corre en SSR y en el cliente
	 * se ejecutaría después del primer render, rompiendo la hidratación.
	 */
	lang.value = langFromPath(page.url.pathname);

	// Y de nuevo al navegar entre `/` y `/en/`, cuando el layout ya está montado
	// y su script no vuelve a ejecutarse.
	$effect.pre(() => {
		lang.value = current;
	});

	/* Datos estructurados para el panel de conocimiento de Google. */
	const organizationLd = $derived({
		'@context': 'https://schema.org',
		'@type': 'Organization',
		name: SITE_NAME,
		url: SITE_URL,
		logo: absolute(`${base}/images/unicore_500.png`),
		description: head.description,
		sameAs: ['https://github.com/ivanjoz']
	});
</script>

<svelte:head>
	<title>{head.title}</title>
	<meta name="description" content={head.description} />
	<link rel="canonical" href={canonical} />

	<!--
		Cada idioma se anuncia a sí mismo y al otro. `x-default` marca a dónde
		mandar al visitante cuyo idioma no coincide con ninguno de los dos.
	-->
	<link rel="alternate" hreflang="es" href={absolute(pathFor('es'))} />
	<link rel="alternate" hreflang="en" href={absolute(pathFor('en'))} />
	<link rel="alternate" hreflang="x-default" href={absolute(pathFor('es'))} />

	<link rel="icon" type="image/png" href={`${base}/images/unicore_32.png`} />
	<link rel="apple-touch-icon" href={`${base}/images/unicore_500.png`} />

	<!--
		Las URL de Open Graph son absolutas a propósito: los crawlers de WhatsApp y
		Facebook no tienen contexto de página para resolver una ruta relativa, así
		que con `./images/...` la previsualización salía sin imagen.
	-->
	<meta property="og:site_name" content={SITE_NAME} />
	<meta property="og:type" content="website" />
	<meta property="og:locale" content={head.locale} />
	<meta property="og:locale:alternate" content={meta[other].locale} />
	<meta property="og:url" content={canonical} />
	<meta property="og:title" content={head.title} />
	<meta property="og:description" content={head.ogDescription} />
	<meta property="og:image" content={absolute(`${base}/images/${head.ogImage}`)} />
	<!-- Sin las dimensiones, Facebook a veces omite la imagen en el primer scrape. -->
	<meta property="og:image:width" content={String(OG_IMAGE_WIDTH)} />
	<meta property="og:image:height" content={String(OG_IMAGE_HEIGHT)} />
	<meta property="og:image:type" content="image/jpeg" />
	<meta property="og:image:alt" content={head.ogImageAlt} />

	<meta name="twitter:card" content="summary_large_image" />
	<meta name="twitter:title" content={head.title} />
	<meta name="twitter:description" content={head.ogDescription} />
	<meta name="twitter:image" content={absolute(`${base}/images/${head.ogImage}`)} />
	<meta name="twitter:image:alt" content={head.ogImageAlt} />

	{@html `<script type="application/ld+json">${JSON.stringify(organizationLd)}<\/script>`}
</svelte:head>

<Header />
{@render children()}
<Footer />
