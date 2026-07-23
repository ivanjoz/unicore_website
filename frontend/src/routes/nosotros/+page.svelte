<script lang="ts">
	import { base } from '$app/paths';
	import Contact from '$lib/components/Contact.svelte';
	import MissionExplorer from '$lib/components/MissionExplorer.svelte';
	import Reveal from '$lib/components/Reveal.svelte';
	import { worldview } from '$lib/data/site';

	const objectives = [
		{
			icon: '/svg/objetivo11.svg',
			title: 'Abrir el horizonte',
			text: 'Poner soluciones tecnológicas al alcance de pequeñas empresas y usuarios para impulsar su actividad.'
		},
		{
			icon: '/svg/cohete1.svg',
			title: 'Crear impacto',
			text: 'Desarrollar productos y servicios accesibles e innovadores que mejoren la vida de las personas.'
		},
		{
			icon: '/svg/people4.svg',
			title: 'Compartir criterio',
			text: 'Fomentar una cultura de libertad, educación e información para la vida personal y profesional.'
		},
		{
			icon: '/svg/idea5.svg',
			title: 'Innovar desde Perú',
			text: 'Construir un equipo multidisciplinario donde nazcan iniciativas capaces de cambiar nuestra realidad.'
		}
	];
</script>

<svelte:head>
	<title>Nosotros | Unicore Perú</title>
	<meta
		name="description"
		content="Conoce la misión, visión, cosmovisión y origen de Unicore Perú."
	/>
</svelte:head>

<main>
	<section
		class="about-hero"
		style:--about-hero-image={`url("${base}/images/fondonosotros17.jpg")`}
	>
		<div class="about-hero-content">
			<p class="eyebrow light">Nuestra historia</p>
			<h1>Ideas que se unen<br />para <em>crear futuro.</em></h1>
			<p>
				Somos tecnología, gestión y una manera consciente de entender el impacto que
				queremos construir.
			</p>
		</div>
		<img src={`${base}/svg/logounicore9.svg`} alt="" />
	</section>

	<section class="mission">
		<Reveal>
			<div class="mission-heading">
				<p class="eyebrow">Nuestra misión</p>
				<h2 class="section-heading">Cada palabra tiene un propósito.</h2>
				<p>Explora los conceptos que orientan lo que hacemos.</p>
			</div>
		</Reveal>
		<MissionExplorer />
	</section>

	<section class="vision">
		<div
			class="vision-image"
			style:--vision-image={`url("${base}/images/vision1.jpg")`}
		></div>
		<div class="vision-copy">
			<p class="eyebrow light">Nuestra visión</p>
			<h2>Que la tecnología genere bienestar real.</h2>
			<p>
				Crear un impacto significativo en el desarrollo de pequeñas empresas mediante
				métodos y tecnologías que impulsen el bienestar económico y social del país.
			</p>
		</div>
	</section>

	<section class="objectives">
		<div class="section-top">
			<div>
				<p class="eyebrow">Objetivos</p>
				<h2 class="section-heading">La dirección que elegimos.</h2>
			</div>
			<p>Cuatro compromisos que convierten nuestra visión en una forma concreta de actuar.</p>
		</div>

		<div class="objective-grid">
			{#each objectives as objective, index}
				<Reveal delay={index * 80}>
					<article>
						<span>{String(index + 1).padStart(2, '0')}</span>
						<img src={`${base}${objective.icon}`} alt="" />
						<h3>{objective.title}</h3>
						<p>{objective.text}</p>
					</article>
				</Reveal>
			{/each}
		</div>
	</section>

	<section class="worldview" style:--worldview-image={`url("${base}/images/fondo7.jpg")`}>
		<div class="worldview-title">
			<p class="eyebrow light">Cosmovisión</p>
			<h2>Lo que creemos define lo que construimos.</h2>
		</div>
		<div class="beliefs">
			{#each worldview as belief, index}
				<article>
					<span>{String(index + 1).padStart(2, '0')}</span>
					<p>{belief}</p>
				</article>
			{/each}
		</div>
	</section>

	<section class="meaning">
		<div class="meaning-intro">
			<p class="eyebrow">Unicore, significado</p>
			<h2 class="section-heading">Un universo con un núcleo común.</h2>
			<p>
				El significado da coherencia, motivación y dirección a cada proyecto. Es la estrella
				que nos permite recobrar el camino.
			</p>
		</div>

		<div class="brand-story first">
			<div class="brand-art">
				<img src={`${base}/images/primer_logo.png`} alt="Primer logotipo de Unicore" />
			</div>
			<div class="brand-copy">
				<p class="eyebrow">El origen de la marca</p>
				<h3>Naturaleza, conocimiento y evolución.</h3>
				<p>
					El primer logo mostraba un árbol enraizando en un cerebro. La naturaleza,
					representada como sistema y sujeto, entrega al ser humano ideas que se convierten
					en conocimiento e invenciones.
				</p>
				<p>
					El átomo, el sol, el foco y el niño hablaban de materia, energía y de la
					responsabilidad de cambiar la realidad sin olvidar nuestro deber de preservar la
					vida.
				</p>
			</div>
		</div>

		<div class="brand-story current">
			<div class="brand-copy">
				<p class="eyebrow">El logo actual</p>
				<h3>Las mismas ideas, en trazos esenciales.</h3>
				<p>
					La forma inferior es una mano tendida y la verde, una hoja: la sinergia entre
					humanidad y naturaleza. Las partículas y ondas aluden a materia, movimiento y
					energía.
				</p>
				<p>
					El círculo une ese universo —UNI— alrededor de un núcleo común —CORE—. Unicore es
					una aleación de ideas y personas impulsadas a expandir sus fronteras.
				</p>
			</div>
			<div class="brand-art">
				<img src={`${base}/svg/unicore_logo5.svg`} alt="Logotipo actual de Unicore" />
			</div>
		</div>
	</section>

	<Contact />
</main>

<style>
	.about-hero {
		position: relative;
		display: flex;
		min-height: 46rem;
		align-items: center;
		justify-content: space-between;
		gap: 4rem;
		overflow: hidden;
		padding: calc(var(--header-height) + 5rem) var(--page-pad) 6rem;
		background:
			linear-gradient(90deg, rgba(3, 18, 31, 0.96), rgba(3, 18, 31, 0.45)),
			var(--about-hero-image) center / cover;
		color: white;
	}

	.about-hero::after {
		position: absolute;
		right: -10rem;
		bottom: -18rem;
		width: 42rem;
		height: 42rem;
		border: 1px solid rgba(255, 255, 255, 0.1);
		border-radius: 50%;
		content: '';
	}

	.about-hero-content {
		position: relative;
		z-index: 1;
		max-width: 58rem;
	}

	.about-hero h1 {
		margin: 1rem 0 1.8rem;
		font-family: var(--font-display);
		font-size: clamp(3.5rem, 7.8vw, 7.7rem);
		font-weight: 500;
		line-height: 0.91;
		letter-spacing: -0.05em;
	}

	.about-hero h1 em {
		color: var(--aqua);
		font-style: normal;
	}

	.about-hero-content > p:last-child {
		max-width: 36rem;
		color: rgba(255, 255, 255, 0.64);
		font-size: 1.05rem;
		line-height: 1.7;
	}

	.about-hero > img {
		position: relative;
		z-index: 1;
		width: min(24rem, 28vw);
		filter: drop-shadow(0 30px 60px rgba(0, 0, 0, 0.25));
	}

	.mission,
	.objectives,
	.meaning {
		padding: clamp(5rem, 10vw, 9rem) var(--page-pad);
	}

	.mission-heading {
		display: grid;
		grid-template-columns: 1fr auto;
		gap: 1rem 3rem;
		align-items: end;
		margin-bottom: clamp(3rem, 7vw, 6rem);
	}

	.mission-heading > .eyebrow {
		grid-column: 1 / -1;
	}

	.mission-heading > p:last-child {
		max-width: 23rem;
		margin: 0;
		color: var(--muted);
	}

	.vision {
		display: grid;
		grid-template-columns: 1fr 1fr;
		min-height: 45rem;
		background: #071b2b;
		color: white;
	}

	.vision-image {
		background:
			linear-gradient(rgba(4, 25, 36, 0.08), rgba(4, 25, 36, 0.24)),
			var(--vision-image) center / cover;
	}

	.vision-copy {
		display: flex;
		flex-direction: column;
		justify-content: center;
		padding: clamp(3rem, 8vw, 8rem);
	}

	.vision-copy h2,
	.worldview-title h2 {
		margin: 1rem 0 1.5rem;
		font-family: var(--font-display);
		font-size: clamp(2.5rem, 5vw, 5rem);
		font-weight: 500;
		line-height: 1;
		letter-spacing: -0.04em;
	}

	.vision-copy > p:last-child {
		max-width: 38rem;
		color: rgba(255, 255, 255, 0.64);
		line-height: 1.8;
	}

	.objectives {
		background: var(--mist);
	}

	.section-top {
		display: flex;
		align-items: end;
		justify-content: space-between;
		gap: 3rem;
		margin-bottom: clamp(3rem, 7vw, 6rem);
	}

	.section-top > p {
		max-width: 28rem;
		margin: 0;
		color: var(--muted);
		line-height: 1.7;
	}

	.objective-grid {
		display: grid;
		grid-template-columns: repeat(4, 1fr);
		gap: 1rem;
	}

	.objective-grid article {
		min-height: 27rem;
		padding: 1.6rem;
		border: 1px solid var(--line);
		border-radius: var(--radius-lg);
		background: rgba(255, 255, 255, 0.76);
	}

	.objective-grid article > span {
		color: var(--teal);
		font-size: 0.68rem;
	}

	.objective-grid img {
		width: 7rem;
		height: 7rem;
		margin: 3rem auto;
		object-fit: contain;
	}

	.objective-grid h3 {
		margin: 0 0 0.8rem;
		font-family: var(--font-display);
		font-size: 1.35rem;
	}

	.objective-grid p {
		margin: 0;
		color: var(--muted);
		font-size: 0.85rem;
		line-height: 1.65;
	}

	.worldview {
		display: grid;
		grid-template-columns: minmax(18rem, 0.72fr) minmax(0, 1.28fr);
		gap: clamp(3rem, 9vw, 9rem);
		padding: clamp(5rem, 10vw, 9rem) var(--page-pad);
		background:
			linear-gradient(115deg, rgba(3, 19, 29, 0.97), rgba(3, 19, 29, 0.75)),
			var(--worldview-image) center / cover fixed;
		color: white;
	}

	.worldview-title {
		position: sticky;
		top: calc(var(--header-height) + 3rem);
		align-self: start;
	}

	.beliefs {
		border-top: 1px solid rgba(255, 255, 255, 0.15);
	}

	.beliefs article {
		display: grid;
		grid-template-columns: 3rem 1fr;
		gap: 1rem;
		padding: 1.45rem 0;
		border-bottom: 1px solid rgba(255, 255, 255, 0.15);
	}

	.beliefs span {
		color: var(--aqua);
		font-size: 0.68rem;
	}

	.beliefs p {
		margin: 0;
		color: rgba(255, 255, 255, 0.75);
		line-height: 1.6;
	}

	.meaning {
		background: var(--sand);
	}

	.meaning-intro {
		max-width: 63rem;
		margin-bottom: clamp(4rem, 8vw, 8rem);
	}

	.meaning-intro > p:last-child {
		max-width: 40rem;
		margin: 1.6rem 0 0;
		color: var(--muted);
		line-height: 1.75;
	}

	.brand-story {
		display: grid;
		grid-template-columns: 1fr 1fr;
		align-items: stretch;
		overflow: hidden;
		border: 1px solid #d7d0c4;
		border-radius: var(--radius-lg);
		background: #fffdf8;
		box-shadow: 0 30px 80px rgba(65, 54, 32, 0.08);
	}

	.brand-story + .brand-story {
		margin-top: 2rem;
	}

	.brand-art {
		display: grid;
		min-height: 34rem;
		place-items: center;
		padding: 3rem;
		background: #e8e3d8;
	}

	.brand-art img {
		width: min(22rem, 82%);
		max-height: 25rem;
		object-fit: contain;
	}

	.current .brand-art {
		background: #dfeeed;
	}

	.brand-copy {
		display: flex;
		flex-direction: column;
		justify-content: center;
		padding: clamp(2rem, 6vw, 6rem);
	}

	.brand-copy h3 {
		margin: 1rem 0 1.5rem;
		font-family: var(--font-display);
		font-size: clamp(2rem, 4vw, 3.6rem);
		font-weight: 500;
		line-height: 1.04;
		letter-spacing: -0.035em;
	}

	.brand-copy > p:not(.eyebrow) {
		margin: 0 0 1rem;
		color: var(--muted);
		line-height: 1.7;
	}

	@media (max-width: 980px) {
		.objective-grid {
			grid-template-columns: 1fr 1fr;
		}
	}

	@media (max-width: 760px) {
		.about-hero > img {
			position: absolute;
			right: -4rem;
			bottom: -2rem;
			width: 18rem;
			opacity: 0.3;
		}

		.mission-heading,
		.vision,
		.worldview,
		.brand-story {
			grid-template-columns: 1fr;
		}

		.mission-heading > p:last-child {
			margin-top: 0.8rem;
		}

		.vision-image {
			min-height: 26rem;
		}

		.worldview-title {
			position: static;
		}

		.current .brand-copy {
			order: 2;
		}

		.current .brand-art {
			order: 1;
		}

		.section-top {
			display: grid;
			align-items: start;
		}
	}

	@media (max-width: 520px) {
		.objective-grid {
			grid-template-columns: 1fr;
		}
	}
</style>
