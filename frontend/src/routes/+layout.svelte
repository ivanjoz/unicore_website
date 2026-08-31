<script lang="ts">
	import { onMount } from 'svelte';
	import Header from '$lib/components/Header.svelte';
	import Footer from '$lib/components/Footer.svelte';
	import { base } from '$app/paths';
	import { lang, restoreLang } from '$lib/i18n.svelte';
	import './global.css';

	let { children } = $props();

	// El HTML prerenderizado sale siempre en español: la preferencia guardada (o la
	// del navegador) solo puede aplicarse una vez montada la página en el cliente.
	onMount(restoreLang);
</script>

<svelte:head>
	<link rel="icon" type="image/png" href={`${base}/images/unicore_32.png`} />
	<meta property="og:locale" content={lang.value === 'es' ? 'es_PE' : 'en_US'} />
	<meta property="og:type" content="website" />
	<meta property="og:image" content={`${base}/images/unicore_logo.png`} />
	<meta name="twitter:card" content="summary_large_image" />
</svelte:head>

<Header />
{@render children()}
<Footer />
