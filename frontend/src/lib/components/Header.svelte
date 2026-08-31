<script lang="ts">
	import { page } from '$app/state';
	import { base } from '$app/paths';

	const links = [
		{ href: `${base}/`, label: 'Inicio' },
		{ href: `${base}/#proyectos`, label: 'Proyectos' },
		{ href: `${base}/#nosotros`, label: 'Nosotros' },
		{ href: `${base}/#contacto`, label: 'Contacto' }
	];

	let menuOpen = $state(false);
	let currentPath = $derived(page.url.pathname);

	// La cabecera se compacta en cuanto la página deja de estar arriba del todo. El umbral
	// no es cero para que un rebote de scroll de unos pocos píxeles no la haga parpadear.
	let scrollY = $state(0);
	let scrolled = $derived(scrollY > 24);

	function isActive(href: string) {
		if (href.includes('#')) return false;
		return href === `${base}/` ? currentPath === `${base}/` : currentPath.startsWith(href);
	}
</script>

<svelte:window bind:scrollY />

<header class:menu-open={menuOpen} class:scrolled>
	<a
		class="brand"
		href={`${base}/`}
		aria-label="Unicore, ir al inicio"
		onclick={() => (menuOpen = false)}
	>
		<img src={`${base}/svg/logo_unicore_horizontal_light_mini.svg`} alt="Unicore Labs" />
	</a>

	<nav aria-label="Navegación principal">
		{#each links as link}
			<a
				href={link.href}
				class:active={isActive(link.href)}
				aria-current={isActive(link.href) ? 'page' : undefined}
				onclick={() => (menuOpen = false)}
			>
				{link.label}
			</a>
		{/each}
	</nav>

	<button
		class="menu-toggle"
		type="button"
		aria-label={menuOpen ? 'Cerrar menú' : 'Abrir menú'}
		aria-expanded={menuOpen}
		onclick={() => (menuOpen = !menuOpen)}
	>
		<span></span><span></span>
	</button>
</header>

<style>
	header {
		position: fixed;
		z-index: 100;
		top: 0;
		left: 0;
		display: flex;
		width: 100%;
		height: var(--header-height);
		align-items: center;
		justify-content: space-between;
		padding: 0 clamp(1.25rem, 5vw, 5rem);
		border-bottom: 1px solid rgba(255, 255, 255, 0.12);
		background: rgba(8, 9, 28, 0.82);
		color: white;
		backdrop-filter: blur(18px);
		transition: height 220ms ease;
	}

	header.scrolled {
		height: 54px;
	}

	.brand {
		display: flex;
		align-items: center;
		gap: 0.65rem;
		color: white;
		text-decoration: none;
	}

	.brand img {
		width: auto;
		height: 44px;
		object-fit: contain;
		filter: drop-shadow(0 5px 14px rgba(0, 0, 0, 0.3));
		transition: height 220ms ease;
	}

	/* 2.6rem no cabe en 54px: el logo se encoge con la barra. */
	.scrolled .brand img {
		height: 38px;
	}

	@media (max-width: 400px) {
		.brand img {
			height: 2.1rem;
		}
	}

	nav {
		display: flex;
		align-items: center;
		gap: clamp(1.5rem, 3vw, 3rem);
	}

	nav a {
		position: relative;
		padding: 0.65rem 0;
		color: rgba(255, 255, 255, 0.72);
		font-size: 0.72rem;
		font-weight: 700;
		letter-spacing: 0.16em;
		text-decoration: none;
		text-transform: uppercase;
		transition:
			color 180ms ease,
			transform 180ms ease;
	}

	nav a::after {
		position: absolute;
		right: 0;
		bottom: 0.2rem;
		left: 0;
		height: 2px;
		background: var(--aqua);
		content: '';
		transform: scaleX(0);
		transition: transform 180ms ease;
	}

	nav a:hover,
	nav a.active {
		color: white;
		transform: translateY(-1px);
	}

	nav a.active::after {
		transform: scaleX(1);
	}

	.menu-toggle {
		display: none;
		width: 2.75rem;
		height: 2.75rem;
		border: 1px solid rgba(255, 255, 255, 0.18);
		border-radius: 999px;
		background: transparent;
	}

	.menu-toggle span {
		display: block;
		width: 1.05rem;
		height: 1px;
		margin: 0.28rem auto;
		background: white;
		transition: transform 180ms ease;
	}

	@media (max-width: 720px) {
		header {
			padding-inline: 1rem;
		}

		nav {
			position: absolute;
			top: calc(100% + 0.5rem);
			right: 1rem;
			left: 1rem;
			display: grid;
			padding: 1rem;
			border: 1px solid rgba(255, 255, 255, 0.13);
			border-radius: 1.25rem;
			background: rgba(8, 9, 28, 0.98);
			box-shadow: 0 24px 60px rgba(0, 0, 0, 0.32);
			opacity: 0;
			pointer-events: none;
			transform: translateY(-0.75rem);
			transition:
				opacity 180ms ease,
				transform 180ms ease;
		}

		.menu-open nav {
			opacity: 1;
			pointer-events: auto;
			transform: translateY(0);
		}

		nav a {
			padding: 0.85rem 0.65rem;
		}

		nav a::after {
			display: none;
		}

		.menu-toggle {
			display: block;
		}

		.menu-open .menu-toggle span:first-child {
			transform: translateY(0.19rem) rotate(45deg);
		}

		.menu-open .menu-toggle span:last-child {
			transform: translateY(-0.19rem) rotate(-45deg);
		}
	}

	@media (prefers-reduced-motion: reduce) {
		header,
		.brand img {
			transition: none;
		}
	}
</style>
