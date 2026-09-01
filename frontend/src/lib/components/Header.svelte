<script lang="ts">
	import { page } from '$app/state';
	import { base } from '$app/paths';
	import T from '$lib/components/T.svelte';
	import LangToggle from '$lib/components/LangToggle.svelte';
	import { t } from '$lib/i18n.svelte';
	import { langFromPath, pathFor } from '$lib/seo';

	// La navegación tiene que quedarse dentro del idioma actual: desde `/en/` un
	// enlace a `${base}/` devolvería al visitante al español sin avisar.
	let home = $derived(pathFor(langFromPath(page.url.pathname)));
	let links = $derived([
		{ href: home, label: 'Inicio|Home' },
		{ href: `${home}#proyectos`, label: 'Proyectos|Projects' },
		{ href: `${home}#nosotros`, label: 'Nosotros|About us' },
		{ href: `${home}#contacto`, label: 'Contacto|Contact' }
	]);

	let menuOpen = $state(false);
	let currentPath = $derived(page.url.pathname);

	// La cabecera se compacta en cuanto la página deja de estar arriba del todo. El umbral
	// no es cero para que un rebote de scroll de unos pocos píxeles no la haga parpadear.
	let scrollY = $state(0);
	let scrolled = $derived(scrollY > 24);

	function isActive(href: string) {
		if (href.includes('#')) return false;
		return href === home ? currentPath === home : currentPath.startsWith(href);
	}
</script>

<svelte:window bind:scrollY />

<header class:menu-open={menuOpen} class:scrolled>
	<a
		class="brand"
		href={home}
		aria-label={t('Unicore, ir al inicio|Unicore, go to the home page')}
		onclick={() => (menuOpen = false)}
	>
		<img src={`${base}/svg/logo_unicore_horizontal_light_mini.svg`} alt="Unicore Labs" />
	</a>

	<nav aria-label={t('Navegación principal|Main navigation')}>
		{#each links as link}
			<a
				href={link.href}
				class:active={isActive(link.href)}
				aria-current={isActive(link.href) ? 'page' : undefined}
				onclick={() => (menuOpen = false)}
			>
				<T text={link.label} />
			</a>
		{/each}
	</nav>

	<div class="header-actions">
		<div class="header-lang"><LangToggle /></div>

		<button
			class="menu-toggle"
			type="button"
			aria-label={menuOpen ? t('Cerrar menú|Close menu') : t('Abrir menú|Open menu')}
			aria-expanded={menuOpen}
			onclick={() => (menuOpen = !menuOpen)}
		>
			<!--
				Los dos glifos comparten celda y se cruzan al abrir: así el botón no
				cambia de tamaño ni la barra da un salto. Vienen de Iconify (Lucide),
				que el plugin incrusta como máscara CSS, sin petición de red.
			-->
			<span class="icon-[lucide--menu] menu-icon" aria-hidden="true"></span>
			<span class="icon-[lucide--x] menu-icon menu-icon-close" aria-hidden="true"></span>
		</button>
	</div>
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

	/* Ancho fijo: el logo se encoge con la barra al hacer scroll y, sin este contenedor,
	   el menú se desplazaría lateralmente al cambiar el ancho de la imagen. */
	.brand {
		display: flex;
		width: 122px;
		flex: none;
		align-items: center;
		justify-content: flex-start;
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
		.brand {
			width: 94px;
		}

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
		font-size: calc(0.72rem + var(--fs-bump));
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

	.header-actions {
		display: flex;
		align-items: center;
		gap: 0.75rem;
	}

	.menu-toggle {
		display: none;
		width: 2.75rem;
		height: 2.75rem;
		border: 1px solid rgba(255, 255, 255, 0.18);
		border-radius: 999px;
		background: transparent;
		color: white;
		place-items: center;
	}

	.menu-icon {
		/* Misma celda para los dos: se superponen en vez de apilarse. */
		width: 1.5rem;
		height: 1.5rem;
		grid-area: 1 / 1;
		transition:
			opacity 160ms ease,
			transform 220ms ease;
	}

	.menu-icon-close,
	.menu-open .menu-icon {
		opacity: 0;
		transform: rotate(-90deg) scale(0.7);
	}

	.menu-open .menu-icon-close {
		opacity: 1;
		transform: none;
	}

	/*
	 * En móvil no hay cabecera: la barra desaparece por completo y sólo queda la
	 * pestaña del menú anclada en la esquina superior derecha. El <header>
	 * sigue existiendo como ancla del desplegable, así que se vuelve invisible y
	 * deja pasar los clicks (`pointer-events: none`) para no tapar el hero.
	 */
	@media (max-width: 720px) {
		header,
		header.scrolled {
			/* Sin el logo, `space-between` dejaría el botón pegado a la izquierda. */
			height: auto;
			justify-content: flex-end;
			padding: 0;
			border-bottom: none;
			background: transparent;
			backdrop-filter: none;
			pointer-events: none;
		}

		.brand {
			display: none;
		}

		/* En móvil el selector de idioma vive en la primera sección, no aquí. */
		.header-lang {
			display: none;
		}

		.header-actions {
			gap: 0;
			pointer-events: auto;
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

		/*
		 * Pestaña, no botón: dos lados van a ras del borde de la pantalla y el
		 * único vértice que queda dentro se redondea, así que el panel parece
		 * salir de la esquina. Sin borde —a ras del canto no habría nada al otro
		 * lado que separar— y con el sangrado sólo a la izquierda, que es por
		 * donde el redondeo come sitio.
		 */
		.menu-toggle {
			display: grid;
			width: 4rem;
			height: 3rem;
			padding: 0 0 0 8px;
			border: none;
			border-radius: 0 0 0 32px;
			background: rgba(8, 9, 28, 0.72);
			backdrop-filter: blur(14px);
		}
	}

	@media (prefers-reduced-motion: reduce) {
		header,
		.brand img,
		.menu-icon {
			transition: none;
		}
	}
</style>
