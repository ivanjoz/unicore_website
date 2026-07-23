<script lang="ts">
	import { base } from '$app/paths';
	import Contact from '$lib/components/Contact.svelte';
	import Reveal from '$lib/components/Reveal.svelte';
	import { gerpFeatures, services } from '$lib/data/site';
</script>

<svelte:head>
	<title>Unicore Perú | Tecnología que impulsa empresas</title>
	<meta
		name="description"
		content="Desarrollamos páginas web, comercio electrónico y software empresarial para que las empresas peruanas lleguen más lejos."
	/>
	<meta property="og:title" content="Unicore Perú | Tecnología que impulsa empresas" />
	<meta
		property="og:description"
		content="Tecnología y gestión para transformar los procesos de tu empresa."
	/>
</svelte:head>

<main>
	<section class="hero" style:--hero-image={`url("${base}/images/fondo5.jpg")`}>
		<div class="hero-overlay"></div>
		<div class="hero-content">
			<p class="hero-kicker">SOLUCIONES INFORMÁTICAS Y EMPRESARIALES</p>
			<img src={`${base}/svg/logounicore3.svg`} alt="Unicore" class="hero-logo" />
			<h1>Tecnología para<br /><em>llegar más lejos.</em></h1>
			<p class="hero-lead">
				Conectamos estrategia, diseño y desarrollo para transformar la forma en que tu
				empresa trabaja, decide y crece.
			</p>
			<div class="hero-actions">
				<a class="button-primary" href="#servicios">Explorar soluciones <span>↓</span></a>
				<a class="button-secondary" href={`${base}/portafolio/`}>Ver portafolio <span>↗</span></a>
			</div>
		</div>
		<div class="hero-side">TRUJILLO · PERÚ <span></span> DESDE 2018</div>
	</section>

	<section class="intro">
		<Reveal>
			<p class="eyebrow">Quiénes somos</p>
			<h2 class="section-heading">Pensamos en tu negocio antes de escribir código.</h2>
		</Reveal>
		<Reveal delay={120}>
			<div class="intro-copy">
				<p>
					Unicore es un equipo profesional que crea herramientas informáticas y
					empresariales para pequeñas y medianas empresas.
				</p>
				<p>
					Las ayudamos a insertarse en el mundo digital, conectar con sus clientes y
					organizar mejor su gestión con soluciones comprensibles, escalables y hechas para
					su realidad.
				</p>
				<a href={`${base}/nosotros/`}>Conoce nuestra forma de pensar <span>→</span></a>
			</div>
		</Reveal>
	</section>

	<section class="services" id="servicios">
		<div class="section-top">
			<div>
				<p class="eyebrow">Qué hacemos</p>
				<h2 class="section-heading">Un ecosistema digital para avanzar.</h2>
			</div>
			<p>
				Desde tu primera presencia en internet hasta el sistema que articula toda tu
				operación.
			</p>
		</div>

		<div class="service-grid">
			{#each services as service, index}
				<Reveal delay={index * 100}>
					<article class="service-card">
						<div class="service-head">
							<p>{service.tag}</p>
							<span>↗</span>
						</div>
						<img src={`${base}${service.icon}`} alt="" />
						<h3>{service.title}</h3>
						<p>{service.description}</p>
					</article>
				</Reveal>
			{/each}
		</div>
	</section>

	<section class="web-story">
		<div
			class="web-photo"
			role="img"
			aria-label="Equipo de desarrollo trabajando"
			style:--web-image={`url("${base}/images/office.jpg")`}
		></div>
		<div class="web-copy">
			<p class="eyebrow light">Tecnología basada en la web</p>
			<h2>Tu empresa disponible donde estés.</h2>
			<p>
				El software ya no necesita vivir en una sola computadora. La web permite que tu
				equipo trabaje simultáneamente desde cualquier lugar y dispositivo.
			</p>
			<div class="capabilities">
				<div><strong>01</strong><span>Acceso desde cualquier dispositivo</span></div>
				<div><strong>02</strong><span>Información centralizada y segura</span></div>
				<div><strong>03</strong><span>Actualizaciones sin interrupciones</span></div>
				<div><strong>04</strong><span>Capacidad de crecer contigo</span></div>
			</div>
		</div>
	</section>

	<section class="business">
		<Reveal>
			<div class="business-heading">
				<p class="eyebrow">Servicios empresariales</p>
				<h2 class="section-heading">Datos claros.<br />Decisiones mejores.</h2>
			</div>
		</Reveal>
		<Reveal delay={100}>
			<div class="business-copy">
				<p>
					Toda empresa es una red de procesos. Te ayudamos a organizar operaciones,
					finanzas, costos e información para que tengas mayor control de cada decisión.
				</p>
				<div class="metrics">
					<div><strong>1</strong><span>visión integrada de tu empresa</span></div>
					<div><strong>∞</strong><span>posibilidades de automatización</span></div>
				</div>
			</div>
		</Reveal>
	</section>

	<section class="gerp">
		<div class="gerp-intro">
			<div>
				<p class="eyebrow light">Producto en desarrollo</p>
				<img src={`${base}/svg/gerp_logo.svg`} alt="GERP" />
			</div>
			<div>
				<h2>Gestor Empresarial<br />de Recursos y Procesos.</h2>
				<p>
					Una plataforma web intuitiva y accesible para gestionar los procesos
					comerciales, logísticos, financieros y contables de las MYPE.
				</p>
			</div>
		</div>

		<div class="feature-grid">
			{#each gerpFeatures as feature, index}
				<Reveal delay={(index % 4) * 70}>
					<article>
						<span>{String(index + 1).padStart(2, '0')}</span>
						<img src={`${base}${feature.icon}`} alt="" />
						<h3>{feature.title}</h3>
						<p>{feature.description}</p>
					</article>
				</Reveal>
			{/each}
		</div>
	</section>

	<section
		class="about-callout"
		style:--about-image={`url("${base}/images/fondounicore.jpg")`}
	>
		<div>
			<p class="eyebrow light">Sobre Unicore</p>
			<h2>Una marca que nace de unir ideas, personas y propósito.</h2>
			<a class="button-primary" href={`${base}/nosotros/`}>Conócenos <span>→</span></a>
		</div>
	</section>

	<Contact />
</main>

<style>
	.hero {
		position: relative;
		display: flex;
		min-height: max(46rem, 100svh);
		align-items: center;
		overflow: hidden;
		padding: calc(var(--header-height) + 4rem) var(--page-pad) 5rem;
		background: var(--hero-image) center 62% / cover no-repeat;
		color: white;
	}

	.hero::before {
		position: absolute;
		inset: 0;
		background:
			linear-gradient(90deg, rgba(3, 18, 28, 0.9) 0%, rgba(3, 18, 28, 0.55) 48%, transparent 80%),
			linear-gradient(0deg, rgba(4, 19, 28, 0.72), transparent 42%);
		content: '';
	}

	.hero-overlay {
		position: absolute;
		inset: 0;
		background-image: linear-gradient(rgba(255, 255, 255, 0.045) 1px, transparent 1px);
		background-size: 100% 7rem;
		mask-image: linear-gradient(to right, black, transparent 72%);
	}

	.hero-content {
		position: relative;
		z-index: 1;
		width: min(54rem, 82%);
	}

	.hero-kicker {
		margin: 0 0 1.5rem;
		color: var(--aqua);
		font-size: 0.67rem;
		font-weight: 700;
		letter-spacing: 0.2em;
	}

	.hero-logo {
		display: none;
	}

	.hero h1 {
		max-width: 11ch;
		margin: 0;
		font-family: var(--font-display);
		font-size: clamp(3.6rem, 8.4vw, 8.2rem);
		font-weight: 500;
		line-height: 0.88;
		letter-spacing: -0.055em;
	}

	.hero h1 em {
		color: var(--aqua);
		font-style: normal;
	}

	.hero-lead {
		max-width: 38rem;
		margin: 2rem 0 0;
		color: rgba(255, 255, 255, 0.72);
		font-size: clamp(1rem, 1.5vw, 1.15rem);
		line-height: 1.7;
	}

	.hero-actions {
		display: flex;
		flex-wrap: wrap;
		gap: 0.8rem;
		margin-top: 2.3rem;
	}

	.hero-side {
		position: absolute;
		z-index: 1;
		right: clamp(1rem, 3vw, 3rem);
		bottom: 4rem;
		display: flex;
		align-items: center;
		gap: 1rem;
		color: rgba(255, 255, 255, 0.5);
		font-size: 0.58rem;
		letter-spacing: 0.15em;
		writing-mode: vertical-rl;
	}

	.hero-side span {
		width: 1px;
		height: 3.5rem;
		background: rgba(255, 255, 255, 0.28);
	}

	.intro,
	.business {
		display: grid;
		grid-template-columns: minmax(0, 1.05fr) minmax(20rem, 0.95fr);
		gap: clamp(3rem, 10vw, 10rem);
		align-items: start;
		padding: clamp(6rem, 12vw, 11rem) var(--page-pad);
	}

	.intro-copy {
		padding-top: 1.5rem;
	}

	.intro-copy p {
		margin: 0 0 1.2rem;
		color: var(--muted);
		font-size: clamp(1rem, 1.5vw, 1.16rem);
		line-height: 1.8;
	}

	.intro-copy a {
		display: inline-flex;
		gap: 1rem;
		align-items: center;
		margin-top: 1rem;
		color: var(--teal);
		font-size: 0.78rem;
		font-weight: 800;
		letter-spacing: 0.08em;
		text-decoration: none;
		text-transform: uppercase;
	}

	.services {
		padding: clamp(5rem, 10vw, 9rem) var(--page-pad);
		background: var(--mist);
	}

	.section-top {
		display: flex;
		align-items: end;
		justify-content: space-between;
		gap: 3rem;
		margin-bottom: clamp(3rem, 6vw, 5.5rem);
	}

	.section-top > p {
		max-width: 30rem;
		margin: 0;
		color: var(--muted);
		line-height: 1.7;
	}

	.service-grid {
		display: grid;
		grid-template-columns: repeat(3, 1fr);
		gap: 1rem;
	}

	.service-card {
		display: flex;
		min-height: 34rem;
		flex-direction: column;
		padding: clamp(1.4rem, 3vw, 2.2rem);
		border: 1px solid var(--line);
		border-radius: var(--radius-lg);
		background: rgba(255, 255, 255, 0.72);
		transition:
			transform 250ms ease,
			background 250ms ease,
			box-shadow 250ms ease;
	}

	.service-card:hover {
		background: white;
		box-shadow: var(--shadow);
		transform: translateY(-0.4rem);
	}

	.service-head {
		display: flex;
		align-items: center;
		justify-content: space-between;
	}

	.service-head p {
		margin: 0;
		color: var(--teal);
		font-size: 0.61rem;
		font-weight: 800;
		letter-spacing: 0.14em;
	}

	.service-head span {
		font-size: 1.1rem;
	}

	.service-card img {
		width: 76%;
		height: 12rem;
		margin: auto;
		object-fit: contain;
	}

	.service-card h3 {
		max-width: 15ch;
		margin: 0 0 1rem;
		font-family: var(--font-display);
		font-size: clamp(1.6rem, 2.5vw, 2.4rem);
		font-weight: 500;
		line-height: 1.08;
	}

	.service-card > p {
		margin: 0;
		color: var(--muted);
		font-size: 0.9rem;
		line-height: 1.7;
	}

	.web-story {
		display: grid;
		grid-template-columns: 1fr 1fr;
		min-height: 49rem;
		background: #071b2b;
		color: white;
	}

	.web-photo {
		background:
			linear-gradient(rgba(1, 17, 27, 0.12), rgba(1, 17, 27, 0.3)),
			var(--web-image) center / cover;
	}

	.web-copy {
		display: flex;
		flex-direction: column;
		justify-content: center;
		padding: clamp(3rem, 8vw, 8rem);
	}

	.web-copy h2 {
		max-width: 12ch;
		margin: 1rem 0 1.5rem;
		font-family: var(--font-display);
		font-size: clamp(2.6rem, 5vw, 5rem);
		font-weight: 500;
		line-height: 1;
		letter-spacing: -0.04em;
	}

	.web-copy > p:not(.eyebrow) {
		max-width: 37rem;
		color: rgba(255, 255, 255, 0.62);
		line-height: 1.75;
	}

	.capabilities {
		display: grid;
		grid-template-columns: 1fr 1fr;
		gap: 1.5rem;
		margin-top: 3rem;
	}

	.capabilities div {
		display: grid;
		gap: 0.45rem;
		padding-top: 1rem;
		border-top: 1px solid rgba(255, 255, 255, 0.18);
	}

	.capabilities strong {
		color: var(--aqua);
		font-size: 0.68rem;
		letter-spacing: 0.1em;
	}

	.capabilities span {
		color: rgba(255, 255, 255, 0.78);
		font-size: 0.84rem;
		line-height: 1.4;
	}

	.business {
		background: var(--sand);
	}

	.business-copy > p {
		margin: 0;
		color: var(--muted);
		font-size: clamp(1rem, 1.5vw, 1.2rem);
		line-height: 1.8;
	}

	.metrics {
		display: grid;
		grid-template-columns: 1fr 1fr;
		gap: 1rem;
		margin-top: 3rem;
	}

	.metrics div {
		display: grid;
		gap: 0.4rem;
		padding-top: 1.3rem;
		border-top: 1px solid #cfc7b9;
	}

	.metrics strong {
		color: var(--teal);
		font-family: var(--font-display);
		font-size: 2.7rem;
		font-weight: 500;
	}

	.metrics span {
		color: var(--muted);
		font-size: 0.78rem;
		line-height: 1.45;
	}

	.gerp {
		padding: clamp(5rem, 10vw, 9rem) var(--page-pad);
		background:
			radial-gradient(circle at 85% 15%, rgba(93, 72, 151, 0.22), transparent 34rem),
			#101523;
		color: white;
	}

	.gerp-intro {
		display: grid;
		grid-template-columns: minmax(17rem, 0.7fr) minmax(0, 1.3fr);
		gap: clamp(3rem, 9vw, 9rem);
		align-items: center;
	}

	.gerp-intro img {
		width: min(18rem, 80%);
		margin-top: 2rem;
		filter: drop-shadow(0 18px 32px rgba(0, 0, 0, 0.28));
	}

	.gerp-intro h2 {
		margin: 0;
		font-family: var(--font-display);
		font-size: clamp(2.5rem, 5vw, 5rem);
		font-weight: 500;
		line-height: 1;
		letter-spacing: -0.04em;
	}

	.gerp-intro > div:last-child p {
		max-width: 40rem;
		margin: 1.5rem 0 0;
		color: rgba(255, 255, 255, 0.6);
		line-height: 1.75;
	}

	.feature-grid {
		display: grid;
		grid-template-columns: repeat(4, 1fr);
		gap: 1px;
		margin-top: clamp(4rem, 8vw, 7rem);
		background: rgba(255, 255, 255, 0.12);
	}

	.feature-grid article {
		min-height: 20rem;
		padding: 1.7rem;
		background: #101523;
	}

	.feature-grid article > span {
		color: rgba(255, 255, 255, 0.35);
		font-size: 0.65rem;
		letter-spacing: 0.1em;
	}

	.feature-grid img {
		width: 5.5rem;
		height: 5.5rem;
		margin: 2.4rem 0;
		object-fit: contain;
	}

	.feature-grid h3 {
		margin: 0 0 0.7rem;
		font-family: var(--font-display);
		font-size: 1.2rem;
	}

	.feature-grid p {
		margin: 0;
		color: rgba(255, 255, 255, 0.5);
		font-size: 0.8rem;
		line-height: 1.55;
	}

	.about-callout {
		display: flex;
		min-height: 42rem;
		align-items: center;
		padding: 5rem var(--page-pad);
		background:
			linear-gradient(90deg, rgba(3, 22, 33, 0.92), rgba(3, 22, 33, 0.28)),
			var(--about-image) center / cover;
		color: white;
	}

	.about-callout > div {
		max-width: 53rem;
	}

	.about-callout h2 {
		margin: 1rem 0 2rem;
		font-family: var(--font-display);
		font-size: clamp(2.5rem, 5.7vw, 5.8rem);
		font-weight: 500;
		line-height: 1;
		letter-spacing: -0.04em;
	}

	@media (max-width: 980px) {
		.service-grid {
			grid-template-columns: 1fr;
		}

		.service-card {
			display: grid;
			grid-template-columns: 1fr 0.65fr;
			column-gap: 2rem;
			min-height: 22rem;
		}

		.service-head {
			grid-column: 1 / -1;
		}

		.service-card img {
			grid-column: 2;
			grid-row: 2 / 4;
			width: 100%;
		}

		.service-card h3,
		.service-card > p {
			grid-column: 1;
		}

		.feature-grid {
			grid-template-columns: repeat(2, 1fr);
		}
	}

	@media (max-width: 760px) {
		.hero {
			min-height: max(50rem, 100svh);
			background-position: 58% center;
		}

		.hero::before {
			background: linear-gradient(90deg, rgba(3, 18, 28, 0.88), rgba(3, 18, 28, 0.3));
		}

		.hero-content {
			width: 100%;
		}

		.hero-side {
			display: none;
		}

		.intro,
		.business {
			grid-template-columns: 1fr;
		}

		.section-top {
			display: grid;
			align-items: start;
		}

		.web-story {
			grid-template-columns: 1fr;
		}

		.web-photo {
			min-height: 25rem;
		}

		.gerp-intro {
			grid-template-columns: 1fr;
		}

		.service-card {
			display: flex;
			min-height: 31rem;
		}

		.service-card img {
			width: 78%;
		}
	}

	@media (max-width: 520px) {
		.capabilities,
		.metrics,
		.feature-grid {
			grid-template-columns: 1fr;
		}

		.feature-grid article {
			min-height: 16rem;
		}
	}
</style>
