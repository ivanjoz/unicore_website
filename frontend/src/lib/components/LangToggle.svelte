<script lang="ts">
	import { lang, t, toggleLang } from '$lib/i18n.svelte';
</script>

<!-- El idioma activo es el que queda resaltado; el otro es el destino del click. -->
<button
	class="lang-toggle"
	type="button"
	aria-label={t('Cambiar a inglés|Switch to Spanish')}
	onclick={toggleLang}
>
	<span class:active={lang.value === 'es'}>ES</span>
	<span class="lang-sep" aria-hidden="true"></span>
	<span class:active={lang.value === 'en'}>EN</span>
</button>

<style>
	.lang-toggle {
		display: flex;
		align-items: center;
		gap: 0.45rem;
		padding: 0.34rem 0.7rem;
		border: 1px solid rgba(255, 255, 255, 0.18);
		border-radius: 999px;
		background: rgba(255, 255, 255, 0.05);
		color: rgba(255, 255, 255, 0.55);
		font: inherit;
		font-size: calc(0.66rem + var(--fs-bump));
		font-weight: 700;
		letter-spacing: 0.12em;
		cursor: pointer;
		transition:
			border-color 180ms ease,
			background 180ms ease;
	}

	.lang-toggle:hover {
		border-color: rgba(61, 220, 201, 0.55);
		background: rgba(255, 255, 255, 0.1);
	}

	.lang-toggle span.active {
		color: var(--aqua);
	}

	.lang-sep {
		width: 1px;
		height: 0.75rem;
		background: rgba(255, 255, 255, 0.24);
	}

	/*
	 * En móvil el selector deja la cabecera y se queda solo sobre la foto del
	 * hero, sin nada al lado que le dé escala: al tamaño de la barra se leía como
	 * un adorno. Aquí crece hasta ser un control de pleno derecho, con un área de
	 * toque cómoda.
	 */
	@media (max-width: 720px) {
		.lang-toggle {
			gap: 0.55rem;
			padding: 0.5rem 0.9rem;
			font-size: calc(0.82rem + var(--fs-bump));
		}

		.lang-sep {
			height: 0.95rem;
		}
	}
</style>
