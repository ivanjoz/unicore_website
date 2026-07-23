<script lang="ts">
	import { base } from '$app/paths';
	import Contact from '$lib/components/Contact.svelte';
	import Reveal from '$lib/components/Reveal.svelte';
	import { projects } from '$lib/data/site';
</script>

<svelte:head>
	<title>Portafolio | Unicore Perú</title>
	<meta
		name="description"
		content="Conoce una selección de páginas web y sistemas desarrollados por Unicore Perú."
	/>
</svelte:head>

<main>
	<section
		class="portfolio-hero"
		style:--portfolio-image={`url("${base}/images/reunion4.jpg")`}
	>
		<div>
			<p class="eyebrow light">Trabajo seleccionado</p>
			<h1>Proyectos que llevan<br /><em>ideas a la realidad.</em></h1>
			<p>
				Diseñamos productos digitales administrables y entregamos el código fuente para
				que cada cliente conserve el control de su tecnología.
			</p>
		</div>
		<span class="project-count">{String(projects.length).padStart(2, '0')} PROYECTOS</span>
	</section>

	<section class="portfolio-list">
		<div class="list-heading">
			<p class="eyebrow">Portafolio</p>
			<h2 class="section-heading">Construido en colaboración.</h2>
			<p>
				Una selección de páginas, plataformas y experiencias desarrolladas junto a
				organizaciones de diferentes sectores.
			</p>
		</div>

		<div class="project-grid">
			{#each projects as project, index}
				<Reveal delay={(index % 2) * 100}>
					<article class:featured={index === 0 || index === 5}>
						<a href={project.url} target="_blank" rel="noreferrer">
							<div class="project-image">
								<img src={`${base}${project.image}`} alt={`Vista de ${project.title}`} />
								<span>Visitar proyecto ↗</span>
							</div>
							<div class="project-meta">
								<div>
									<p>{project.type}</p>
									<h3>{project.title}</h3>
								</div>
								<span>{String(index + 1).padStart(2, '0')}</span>
							</div>
							<p class="description">{project.description}</p>
						</a>
					</article>
				</Reveal>
			{/each}
		</div>
	</section>

	<section class="principles">
		<div>
			<span>01</span>
			<h3>Contenido administrable</h3>
			<p>Tu equipo puede actualizar textos, galerías y productos sin depender de terceros.</p>
		</div>
		<div>
			<span>02</span>
			<h3>Diseño adaptable</h3>
			<p>Cada experiencia está preparada para pantallas grandes, tablets y celulares.</p>
		</div>
		<div>
			<span>03</span>
			<h3>Tecnología transferible</h3>
			<p>Entregamos el código fuente para conservar la autonomía sobre cada solución.</p>
		</div>
	</section>

	<Contact />
</main>

<style>
	.portfolio-hero {
		position: relative;
		display: flex;
		min-height: 42rem;
		align-items: end;
		justify-content: space-between;
		gap: 3rem;
		padding: calc(var(--header-height) + 5rem) var(--page-pad) 6rem;
		background:
			linear-gradient(90deg, rgba(3, 20, 31, 0.95), rgba(3, 20, 31, 0.3)),
			var(--portfolio-image) center / cover;
		color: white;
	}

	.portfolio-hero > div {
		max-width: 65rem;
	}

	.portfolio-hero h1 {
		margin: 1rem 0 1.6rem;
		font-family: var(--font-display);
		font-size: clamp(3.4rem, 7.5vw, 7.5rem);
		font-weight: 500;
		line-height: 0.91;
		letter-spacing: -0.052em;
	}

	.portfolio-hero h1 em {
		color: var(--aqua);
		font-style: normal;
	}

	.portfolio-hero div > p:last-child {
		max-width: 38rem;
		color: rgba(255, 255, 255, 0.68);
		line-height: 1.7;
	}

	.project-count {
		flex: none;
		color: rgba(255, 255, 255, 0.62);
		font-size: 0.65rem;
		letter-spacing: 0.16em;
		writing-mode: vertical-rl;
	}

	.portfolio-list {
		padding: clamp(5rem, 10vw, 9rem) var(--page-pad);
		background: #f8fbfa;
	}

	.list-heading {
		display: grid;
		grid-template-columns: minmax(0, 1fr) minmax(19rem, 0.6fr);
		gap: 1rem 5rem;
		align-items: end;
		margin-bottom: clamp(4rem, 8vw, 7rem);
	}

	.list-heading .eyebrow {
		grid-column: 1 / -1;
	}

	.list-heading > p:last-child {
		margin: 0;
		color: var(--muted);
		line-height: 1.7;
	}

	.project-grid {
		display: grid;
		grid-template-columns: 1fr 1fr;
		gap: clamp(3rem, 7vw, 7rem) clamp(1rem, 3vw, 3rem);
		align-items: start;
	}

	.project-grid article:nth-child(even) {
		margin-top: 7rem;
	}

	.project-grid a {
		color: inherit;
		text-decoration: none;
	}

	.project-image {
		position: relative;
		aspect-ratio: 1.32 / 1;
		overflow: hidden;
		border-radius: var(--radius-lg);
		background: #dce7e6;
	}

	.project-image::after {
		position: absolute;
		inset: 0;
		background: rgba(5, 33, 43, 0.35);
		content: '';
		opacity: 0;
		transition: opacity 280ms ease;
	}

	.project-image img {
		width: 100%;
		height: 100%;
		object-fit: cover;
		transition: transform 550ms cubic-bezier(0.22, 1, 0.36, 1);
	}

	.project-image > span {
		position: absolute;
		z-index: 1;
		right: 1.5rem;
		bottom: 1.5rem;
		padding: 0.85rem 1rem;
		border-radius: 999px;
		background: white;
		color: var(--ink);
		font-size: 0.67rem;
		font-weight: 800;
		letter-spacing: 0.08em;
		opacity: 0;
		transform: translateY(0.75rem);
		transition:
			opacity 250ms ease,
			transform 250ms ease;
	}

	.project-grid a:hover .project-image img {
		transform: scale(1.045);
	}

	.project-grid a:hover .project-image::after,
	.project-grid a:hover .project-image > span {
		opacity: 1;
	}

	.project-grid a:hover .project-image > span {
		transform: translateY(0);
	}

	.project-meta {
		display: flex;
		align-items: start;
		justify-content: space-between;
		gap: 1rem;
		margin-top: 1.5rem;
	}

	.project-meta p {
		margin: 0 0 0.45rem;
		color: var(--teal);
		font-size: 0.62rem;
		font-weight: 800;
		letter-spacing: 0.12em;
		text-transform: uppercase;
	}

	.project-meta h3 {
		margin: 0;
		font-family: var(--font-display);
		font-size: clamp(1.5rem, 2.7vw, 2.35rem);
		font-weight: 500;
		line-height: 1.1;
	}

	.project-meta > span {
		color: #9aabad;
		font-size: 0.68rem;
	}

	.description {
		max-width: 34rem;
		margin: 0.9rem 0 0;
		color: var(--muted);
		font-size: 0.86rem;
		line-height: 1.65;
	}

	.principles {
		display: grid;
		grid-template-columns: repeat(3, 1fr);
		gap: 1px;
		padding: 1px 0;
		background: #24404c;
	}

	.principles div {
		min-height: 22rem;
		padding: clamp(2rem, 5vw, 5rem);
		background: #071b2b;
		color: white;
	}

	.principles span {
		color: var(--aqua);
		font-size: 0.68rem;
	}

	.principles h3 {
		margin: 5rem 0 1rem;
		font-family: var(--font-display);
		font-size: clamp(1.4rem, 2.3vw, 2rem);
		font-weight: 500;
	}

	.principles p {
		margin: 0;
		color: rgba(255, 255, 255, 0.54);
		font-size: 0.85rem;
		line-height: 1.65;
	}

	@media (max-width: 760px) {
		.list-heading,
		.project-grid,
		.principles {
			grid-template-columns: 1fr;
		}

		.project-grid article:nth-child(even) {
			margin-top: 0;
		}

		.project-count {
			display: none;
		}

		.principles div {
			min-height: 18rem;
		}
	}
</style>
