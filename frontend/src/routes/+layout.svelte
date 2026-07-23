<script lang="ts">
	import Header from '$lib/components/Header.svelte';
	import Footer from '$lib/components/Footer.svelte';
	import { base } from '$app/paths';
	import './global.css';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" type="image/png" href={`${base}/images/unicore_32.png`} />
	<meta property="og:locale" content="es_PE" />
	<meta property="og:type" content="website" />
	<meta property="og:image" content={`${base}/images/unicore_logo.png`} />
	<meta name="twitter:card" content="summary_large_image" />
</svelte:head>

<Header />
{@render children()}
<Footer />
