<script lang="ts">
	import { t } from '$lib/i18n.svelte';

	interface TProps {
		/** Texto bilingüe: 'español|english'. Sin "|" se muestra igual en ambos idiomas. */
		text: string;
		css?: string;
		/**
		 * El texto trae etiquetas propias (<em>, <strong>, <a>…) y se inyecta como HTML.
		 * Solo para textos escritos en el código: nunca para datos de terceros.
		 * El CSS que apunte a esas etiquetas debe envolverse en :global().
		 */
		html?: boolean;
	}

	let { text = '', css, html = false }: TProps = $props();

	const content = $derived(t(text));
</script>

{#if html}
	<span class={css}>{@html content}</span>
{:else}
	<span class={css}>{content}</span>
{/if}
