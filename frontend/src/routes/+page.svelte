<script lang="ts">
	import { base } from '$app/paths';
	import Contact from '$lib/components/Contact.svelte';
	import Reveal from '$lib/components/Reveal.svelte';
	import SectionCurve from '$lib/components/SectionCurve.svelte';
	import { labProjects } from '$lib/data/site';

	const doubts = [
		{
			text: '¿La empresa proveedora seguirá existiendo en los próximos años?',
			icon: '/svg/icons/icon-continuity.svg'
		},
		{
			text: '¿Puedo mover mis datos a mi propia nube?',
			icon: '/svg/icons/icon-portability.svg'
		},
		{
			text: '¿Por qué mi equipo no puede modificar el sistema para adecuarlo a mis procesos?',
			icon: '/svg/icons/icon-processes.svg'
		},
		{
			text: '¿Dependo de un solo proveedor para cualquier soporte?',
			icon: '/svg/icons/icon-support.svg'
		}
	];

	const legacy = [
		{ place: 'Bell Labs', note: 'El transistor, UNIX y la teoría de la información.' },
		{ place: 'Xerox PARC', note: 'La interfaz gráfica, el ratón y la red local.' },
		{
			place: 'Google Research',
			note: '«Attention is All You Need», el origen de la IA moderna.'
		}
	];
</script>

<svelte:head>
	<title>Unicore Labs | Código abierto de alto impacto</title>
	<meta
		name="description"
		content="Iniciativas de código abierto que democratizan el acceso a tecnología de alto impacto. Sistemas y herramientas que incentivan la participación comunitaria y la exploración científica."
	/>
	<meta property="og:title" content="Unicore Labs | Código abierto de alto impacto" />
	<meta
		property="og:description"
		content="Desarrollamos sistemas y herramientas de código abierto que incentivan la participación comunitaria y la exploración científica."
	/>
	<link rel="preload" as="image" href={`${base}/images/space_earth.webp`} />
</svelte:head>

<main>
	<section
		class="hero"
		style:--hero-webp={`url("${base}/images/space_earth.webp")`}
		style:--hero-jpg={`url("${base}/images/space_earth.jpg")`}
	>
		<div class="hero-grid" aria-hidden="true"></div>

		<div class="hero-inner">
			<div class="hero-copy">
				<p class="hero-kicker"><span></span>CÓDIGO ABIERTO · CIENCIA · COMUNIDAD</p>
				<h1>
					Iniciativas de <em>código abierto</em> que<br />
					democratizan el acceso<br />
					a tecnología de alto impacto
				</h1>
				<p class="hero-lead">
					Desarrollamos sistemas y herramientas que incentivan la participación
					comunitaria y la exploración científica.
				</p>
				<div class="hero-actions">
					<a class="button-primary" href="#proyectos">Ver proyectos <span>↓</span></a>
					<a class="button-secondary" href="#contacto">Colaborar <span>↗</span></a>
				</div>
			</div>

			<div class="hero-mark">
				<span class="orbit orbit-a" aria-hidden="true"></span>
				<span class="orbit orbit-b" aria-hidden="true"></span>
				<img src={`${base}/svg/logo_unicore_vertical_light_2.svg`} alt="Unicore Labs" />
			</div>
		</div>

		<div class="hero-foot">
			<span>GPL · MIT · APACHE 2.0</span>
			<span class="hero-foot-line"></span>
			<span>{labProjects.length} PROYECTOS ABIERTOS</span>
		</div>
	</section>

	<section class="intro" id="nosotros">
		<Reveal>
			<div class="intro-copy">
				<p class="eyebrow">¿Quiénes somos?</p>
				<h2 class="section-heading">Un equipo profesional<br />descentralizado.</h2>
				<p>
					Desarrollamos tecnologías de código abierto y apostamos por iniciativas de
					interés general, con fines comerciales, científicos y/o sociales.
				</p>
				<p>
					Desde sistemas de facturación, logística y finanzas para pequeñas empresas,
					aplicaciones de comercio electrónico y automatización de soporte con IA, hasta
					algoritmos de serialización para reducir el uso de ancho de banda público.
				</p>
				<div class="intro-cta">
					<p>
						¿Tienes un proyecto de código abierto o quieres participar en el desarrollo de
						alguno? Escríbenos: podemos ayudarte y asesorarte.
					</p>
					<a href="#contacto">Conversemos <span>→</span></a>
				</div>
			</div>
		</Reveal>
		<Reveal delay={120}>
			<figure class="intro-art">
				<img
					src={`${base}/svg/coworking.svg`}
					alt="Tres personas trabajando con laptops en un espacio compartido"
					loading="lazy"
				/>
			</figure>
		</Reveal>
	</section>

	<section class="portfolio" id="proyectos">
		<SectionCurve fill="#f8fbfa" />

		<div class="section-top">
			<div>
				<p class="eyebrow">Portafolio</p>
				<h2 class="section-heading">Lo que estamos construyendo.</h2>
			</div>
			<p>
				Cada pieza resuelve un problema concreto y vive en su propio repositorio, con
				licencia abierta y documentación pública.
			</p>
		</div>

		<div class="project-grid">
			{#each labProjects as project, index}
				<Reveal delay={(index % 3) * 90}>
					<article class="project-card">
						<div class="project-logo">
							{#if project.logo}
								<img src={`${base}${project.logo}`} alt={`Logo de ${project.name}`} />
							{:else}
								<span>{project.name}</span>
							{/if}
						</div>

						<div class="project-body">
							<div class="project-meta">
								<p class="project-kind">{project.kind}</p>
								<span
									class="project-status"
									class:live={project.status === 'Activo'}
									class:draft={project.status === 'En diseño'}
								>
									{project.status}
								</span>
							</div>
							<h3>{project.name}</h3>
							<p class="project-text">{project.description}</p>

							<ul class="project-stack">
								{#each project.stack as item}
									<li>{item}</li>
								{/each}
							</ul>

							{#if project.url}
								<a class="project-link" href={project.url} target="_blank" rel="noreferrer">
									Ver repositorio <span>↗</span>
								</a>
							{:else}
								<span class="project-link muted">Próximamente</span>
							{/if}
						</div>
					</article>
				</Reveal>
			{/each}
		</div>
	</section>

	<section class="open-source">
		<SectionCurve fill="var(--mist)" />

		<div class="os-head">
			<p class="eyebrow light">¿Por qué código abierto?</p>
			<h2>Las empresas open source son disruptivas por naturaleza.</h2>
			<p>
				Linux, Red Hat y Odoo demostraron que el código abierto crea disponibilidad de
				mercado, impulsa la adopción, la visibilidad, la construcción de comunidad y la
				ciberseguridad.
			</p>
		</div>

		<div class="doubts">
			{#each doubts as doubt, index}
				<Reveal delay={index * 80}>
					<blockquote>
						<span class="doubt-num">{String(index + 1).padStart(2, '0')}</span>
						<div class="doubt-icon">
							<img src={`${base}${doubt.icon}`} alt="" loading="lazy" />
						</div>
						<p>{doubt.text}</p>
					</blockquote>
				</Reveal>
			{/each}
		</div>

		<div class="horizon" aria-hidden="true">
			<svg viewBox="0 0 1440 80" preserveAspectRatio="none" focusable="false">
				<defs>
					<linearGradient
						id="horizonStroke"
						x1="0"
						y1="0"
						x2="1440"
						y2="0"
						gradientUnits="userSpaceOnUse"
					>
						<stop offset="0" stop-color="#52ded3" stop-opacity="0" />
						<stop offset="0.12" stop-color="#52ded3" stop-opacity="0.55" />
						<stop offset="0.5" stop-color="#dcfaff" stop-opacity="1" />
						<stop offset="0.88" stop-color="#8d7ce6" stop-opacity="0.55" />
						<stop offset="1" stop-color="#8d7ce6" stop-opacity="0" />
					</linearGradient>
				</defs>
				<path
					d="M0 72 C 400 8 1040 8 1440 72"
					fill="none"
					stroke="url(#horizonStroke)"
					stroke-width="2"
					vector-effect="non-scaling-stroke"
				/>
			</svg>
		</div>

		<Reveal>
			<div class="os-answer">
				<p>
					El open source resuelve lo anterior. <strong
						>El proveedor deja de ser el cuello de botella</strong
					>, el soporte se descentraliza, el sistema es más flexible y la IA es más
					autónoma al poder inspeccionar el código. Estos proyectos crecen de forma
					orgánica con la colaboración de estudiantes, profesionales y empresas SaaS que
					los incorporan como base de su oferta de servicios.
				</p>
			</div>
		</Reveal>
	</section>

	<section class="labs">
		<SectionCurve fill="#071b2b" variant="dome" flip />

		<Reveal>
			<div class="labs-head">
				<p class="eyebrow">¿Por qué «labs»?</p>
				<h2 class="section-heading">Un espacio de innovación e investigación</h2>
			</div>
		</Reveal>

		<div class="labs-body">
			<Reveal delay={100}>
				<div class="labs-copy">
					<p>
						Inspirados en Bell Labs y Xerox PARC. Espacios de innovación donde nacieron los
						fundamentos tecnológicos que dieron forma a las décadas siguientes. Y en Google
						Research que produjo «Attention is All You Need», el origen de la IA moderna.
					</p>
					<p>
						Creemos en anteponer el desarrollo de producto por sobre cualquier otro eje empresarial. El producto es la empresa y la investigación da forma al producto. Las regalías comerciales sirven para financiar al producto y su investigación, el cual responde a un impacto esperado. El sufijo «labs» abraza estos conceptos.
					</p>
					<p>
						Si eres una mente curiosa, con habilidades para la programación y quieres aportar al código abierto, escríbemos. Eres libre de hacer un fork de cualquier de nuestros proyectos y usarlos como base de nuevos desarrollos. Háznoslo saber en GitHub.
					</p>
				</div>
			</Reveal>
			<Reveal delay={160}>
				<figure class="labs-art">
					<img
						src={`${base}/svg/connecting-teams.svg`}
						alt="Personas colaborando y compartiendo información entre sí"
						loading="lazy"
					/>
				</figure>
			</Reveal>
		</div>
	</section>

	<section class="ai">
		<SectionCurve fill="var(--sand)" variant="valley" />

		<div class="ai-copy">
			<p class="eyebrow light">Inteligencia artificial</p>
			<h2>IA al alcance de cualquier usuario.</h2>
			<p>
				Enfrentamos el reto de democratizar el acceso al software reduciendo el costo de la
				IA por proceso. Usamos IA de frontera para desarrollar, pero implementamos métodos
				más determinísticos para desplegar servicios eficientes basados en IA.
			</p>
			<p>
				Genix Agentic UI es un ejemplo de cómo modelos pequeños y sin visión pueden servir
				como agentes que ayudan al usuario a navegar el sistema y ejecutar instrucciones.
				Ello permite exprimir cada dólar de inferencia.
			</p>
		</div>

		<div class="ai-stats">
			<div>
				<strong>&lt; 200B</strong>
				<span>modelos sin visión operando como agentes de interfaz</span>
			</div>
			<div>
				<strong>&lt; 32B</strong>
				<span>objetivo de fine-tuning para inferencia asequible</span>
			</div>
			<div>
				<strong>Ad-hoc</strong>
				<span>modelos propios como Classi-Cont para tareas específicas</span>
			</div>
			<div>
				<strong>Local</strong>
				<span>verdadera IA en el dispositivo en los próximos años</span>
			</div>
		</div>
	</section>

	<Contact />
</main>

<style>
	/* ---------------------------------------------------------------- hero */
	.hero {
		position: relative;
		display: flex;
		min-height: max(44rem, 100svh);
		flex-direction: column;
		justify-content: center;
		overflow: hidden;
		padding: calc(var(--header-height) + 4rem) var(--page-pad) 4.5rem;
		background: #02070f var(--hero-jpg) center 32% / cover no-repeat;
		color: white;
	}

	@supports (background-image: image-set(url('a.webp') type('image/webp'))) {
		.hero {
			background-image: image-set(var(--hero-webp) type('image/webp'), var(--hero-jpg) type('image/jpeg'));
		}
	}

	.hero::before {
		position: absolute;
		inset: 0;
		background:
			linear-gradient(100deg, rgba(2, 8, 16, 0.9) 0%, rgba(2, 8, 16, 0.56) 46%, rgba(2, 8, 16, 0.24) 100%),
			linear-gradient(0deg, rgba(2, 8, 16, 0.72), transparent 45%);
		content: '';
	}

	.hero-grid {
		position: absolute;
		inset: 0;
		background-image:
			linear-gradient(rgba(120, 200, 255, 0.05) 1px, transparent 1px),
			linear-gradient(90deg, rgba(120, 200, 255, 0.05) 1px, transparent 1px);
		background-size: 7rem 7rem;
		mask-image: radial-gradient(circle at 20% 40%, black, transparent 70%);
	}

	.hero-inner {
		position: relative;
		z-index: 1;
		display: grid;
		width: 100%;
		grid-template-columns: minmax(0, 1.08fr) minmax(0, 0.92fr);
		gap: clamp(2rem, 6vw, 6rem);
		align-items: center;
	}

	.hero-kicker {
		display: flex;
		align-items: center;
		gap: 0.85rem;
		margin: 0 0 1.6rem;
		color: var(--aqua);
		font-size: 0.68rem;
		font-weight: 700;
		letter-spacing: 0.32em;
	}

	.hero-kicker span {
		width: 2.5rem;
		height: 1px;
		background: linear-gradient(90deg, transparent, var(--aqua));
	}

	.hero h1 {
		margin: 0;
		font-family: var(--font-display);
		font-size: clamp(1.5rem, 2.85vw, 2.75rem);
		font-weight: 500;
		line-height: 1.15;
		letter-spacing: -0.035em;
	}

	.hero h1 em {
		position: relative;
		background: linear-gradient(96deg, #52ded3 0%, #63aeff 52%, #a69aff 100%);
		-webkit-background-clip: text;
		background-clip: text;
		color: transparent;
		font-style: normal;
	}

	.hero-lead {
		max-width: 34rem;
		margin: 1.8rem 0 0;
		color: rgba(226, 240, 248, 0.72);
		font-size: clamp(1rem, 1.35vw, 1.15rem);
		line-height: 1.75;
	}

	.hero-actions {
		display: flex;
		flex-wrap: wrap;
		gap: 0.8rem;
		margin-top: 2.4rem;
	}

	.hero-mark {
		position: relative;
		width: min(20rem, 100%);
		/* Matches the vertical lockup: circular mark stacked over the wordmark. */
		aspect-ratio: 208.38 / 235.73;
		margin-inline: auto;
	}

	.hero-mark img {
		position: relative;
		z-index: 1;
		width: 100%;
		filter: drop-shadow(0 1.5rem 3rem rgba(0, 0, 0, 0.55));
		animation: float 9s ease-in-out infinite;
	}

	/*
	 * The photo is brightest exactly where the lockup sits, so it gets its own
	 * dark backdrop. Everything decorative is centred on the circular mark
	 * (44.2% down the lockup), not on the box, which also holds the wordmark.
	 */
	.hero-mark::before {
		position: absolute;
		top: 44.2%;
		left: 50%;
		width: 190%;
		aspect-ratio: 1;
		border-radius: 50%;
		background: radial-gradient(
			circle,
			rgba(1, 6, 12, 0.78) 0%,
			rgba(1, 6, 12, 0.68) 30%,
			rgba(2, 10, 20, 0.38) 52%,
			transparent 74%
		);
		content: '';
		translate: -50% -50%;
	}

	.orbit {
		position: absolute;
		top: 44.2%;
		left: 50%;
		border: 1px solid rgba(160, 205, 255, 0.18);
		border-radius: 50%;
		aspect-ratio: 1;
		translate: -50% -50%;
	}

	.orbit-a {
		width: 114%;
		border-style: dashed;
		animation: spin 46s linear infinite;
	}

	.orbit-b {
		width: 132%;
		border-color: rgba(160, 205, 255, 0.1);
	}

	@keyframes float {
		0%,
		100% {
			transform: translateY(-0.55rem);
		}
		50% {
			transform: translateY(0.55rem);
		}
	}

	@keyframes spin {
		to {
			transform: rotate(360deg);
		}
	}

	.hero-foot {
		position: relative;
		z-index: 1;
		display: flex;
		align-items: center;
		gap: 1.1rem;
		margin-top: clamp(2.5rem, 6vh, 4.5rem);
		color: rgba(226, 240, 248, 0.45);
		font-size: 0.6rem;
		letter-spacing: 0.24em;
	}

	.hero-foot-line {
		width: clamp(2rem, 12vw, 9rem);
		height: 1px;
		background: rgba(226, 240, 248, 0.22);
	}

	/* --------------------------------------------------------------- intro */
	.intro {
		display: grid;
		grid-template-columns: minmax(0, 1.12fr) minmax(0, 0.88fr);
		gap: clamp(1.75rem, 3.2vw, 3.25rem);
		align-items: center;
		padding: clamp(5rem, 11vw, 10rem) var(--page-pad) clamp(3rem, 6vw, 5.5rem);
	}

	/* Narrower column than the full-bleed headings, so it needs its own scale. */
	.intro-copy .section-heading {
		max-width: none;
		margin: 1rem 0 2rem;
		font-size: clamp(1.9rem, 3.4vw, 3.2rem);
	}

	.intro-copy p:not(.eyebrow) {
		max-width: 44rem;
		margin: 0 0 1.2rem;
		color: var(--muted);
		font-size: clamp(0.98rem, 1.3vw, 1.1rem);
		line-height: 1.8;
	}

	.intro-art {
		margin: 0;
	}

	.intro-art img {
		width: 100%;
		height: auto;
	}

	.intro-cta {
		display: flex;
		flex-wrap: wrap;
		align-items: center;
		justify-content: space-between;
		gap: 1.2rem clamp(1.2rem, 3vw, 2.5rem);
		max-width: 44rem;
		margin-top: 2.25rem;
		padding: clamp(1.4rem, 2.2vw, 1.8rem) clamp(1.4rem, 2.2vw, 1.9rem);
		border: 1px solid var(--line);
		border-left: 3px solid var(--teal);
		border-radius: 0 1rem 1rem 0;
		background: var(--mist);
	}

	.intro-cta p {
		flex: 1 1 20rem;
		margin: 0;
		color: var(--ink);
		font-size: clamp(0.95rem, 1.2vw, 1.05rem);
		line-height: 1.7;
	}

	.intro-cta a {
		display: inline-flex;
		flex: none;
		gap: 0.7rem;
		align-items: center;
		padding: 0.9rem 1.4rem;
		border-radius: 999px;
		background: var(--teal);
		color: white;
		font-size: 0.75rem;
		font-weight: 800;
		letter-spacing: 0.1em;
		text-decoration: none;
		text-transform: uppercase;
		transition:
			gap 180ms ease,
			background 180ms ease;
	}

	.intro-cta a:hover {
		background: var(--teal-dark);
		gap: 1.1rem;
	}

	/* ----------------------------------------------------------- portfolio */
	.portfolio {
		position: relative;
		padding: calc(var(--curve-h) + clamp(2.5rem, 5vw, 4.5rem)) var(--page-pad)
			clamp(3rem, 6vw, 5.5rem);
		background: var(--mist);
	}

	.section-top {
		display: flex;
		align-items: end;
		justify-content: space-between;
		gap: 3rem;
		margin-bottom: clamp(3rem, 6vw, 5rem);
	}

	.section-top > p {
		max-width: 28rem;
		margin: 0;
		color: var(--muted);
		line-height: 1.7;
	}

	.project-grid {
		display: grid;
		grid-template-columns: repeat(3, minmax(0, 1fr));
		gap: 1.1rem;
	}

	/* Reveal wraps each card, so stretch the wrapper to keep rows even. */
	.project-grid :global(.reveal) {
		display: flex;
		height: 100%;
	}

	.project-card {
		display: flex;
		width: 100%;
		flex-direction: column;
		overflow: hidden;
		border: 1px solid var(--line);
		border-radius: var(--radius-lg);
		background: white;
		transition:
			transform 250ms ease,
			box-shadow 250ms ease,
			border-color 250ms ease;
	}

	.project-card:hover {
		border-color: transparent;
		box-shadow: var(--shadow);
		transform: translateY(-0.4rem);
	}

	.project-logo {
		display: grid;
		height: 8.5rem;
		flex: none;
		align-items: center;
		justify-items: center;
		border-bottom: 1px solid var(--line);
		background:
			radial-gradient(circle at 30% 20%, rgba(82, 222, 211, 0.12), transparent 60%),
			linear-gradient(160deg, #f4f9f8, #eaf2f1);
	}

	.project-logo img {
		max-width: 62%;
		max-height: 4.2rem;
		object-fit: contain;
	}

	.project-logo span {
		padding-inline: 1rem;
		color: #8ba3a8;
		font-family: var(--font-display);
		font-size: clamp(1.15rem, 1.8vw, 1.5rem);
		font-weight: 500;
		letter-spacing: 0.02em;
		text-align: center;
	}

	.project-body {
		display: flex;
		flex: 1;
		flex-direction: column;
		padding: clamp(1.3rem, 2.2vw, 1.75rem);
	}

	.project-meta {
		display: flex;
		align-items: center;
		justify-content: space-between;
		gap: 0.75rem;
	}

	.project-kind {
		margin: 0;
		color: var(--teal);
		font-size: 0.6rem;
		font-weight: 800;
		letter-spacing: 0.14em;
		text-transform: uppercase;
	}

	.project-status {
		display: inline-flex;
		align-items: center;
		gap: 0.4rem;
		flex: none;
		padding: 0.3rem 0.6rem;
		border: 1px solid var(--line);
		border-radius: 999px;
		background: #f7fbfa;
		color: var(--muted);
		font-size: 0.58rem;
		font-weight: 700;
		letter-spacing: 0.08em;
		text-transform: uppercase;
	}

	.project-status::before {
		width: 0.4rem;
		height: 0.4rem;
		border-radius: 50%;
		background: #b8c6c8;
		content: '';
	}

	.project-status.live::before {
		background: #16b58a;
	}

	.project-status.draft::before {
		background: #c1a05f;
	}

	.project-card h3 {
		margin: 0.9rem 0 0.7rem;
		font-family: var(--font-display);
		font-size: clamp(1.35rem, 2vw, 1.75rem);
		font-weight: 500;
		line-height: 1.1;
	}

	.project-text {
		margin: 0 0 1.3rem;
		color: var(--muted);
		font-size: 0.86rem;
		line-height: 1.65;
	}

	.project-stack {
		display: flex;
		flex-wrap: wrap;
		gap: 0.4rem;
		margin: auto 0 1.2rem;
		padding: 0;
		list-style: none;
	}

	.project-stack li {
		padding: 0.28rem 0.6rem;
		border-radius: 0.45rem;
		background: var(--mist);
		color: #5f7378;
		font-size: 0.65rem;
		letter-spacing: 0.04em;
	}

	.project-link {
		display: inline-flex;
		align-items: center;
		gap: 0.55rem;
		padding-top: 1rem;
		border-top: 1px solid var(--line);
		color: var(--teal);
		font-size: 0.7rem;
		font-weight: 800;
		letter-spacing: 0.1em;
		text-decoration: none;
		text-transform: uppercase;
		transition: gap 180ms ease;
	}

	.project-link:hover {
		gap: 0.95rem;
	}

	.project-link.muted {
		color: #9fb0b3;
	}

	/* --------------------------------------------------------- open source */
	.open-source {
		position: relative;
		padding: calc(var(--curve-h) + clamp(2.5rem, 5vw, 4.5rem)) var(--page-pad)
			clamp(3rem, 6vw, 5.5rem);
		background:
			radial-gradient(circle at 80% 15%, rgba(93, 122, 220, 0.22), transparent 32rem),
			#071b2b;
		color: white;
	}

	.os-head {
		text-align: center;
	}

	.os-head h2 {
		max-width: 20ch;
		margin: 1rem auto 1.4rem;
		font-family: var(--font-display);
		font-size: clamp(2.2rem, 4.8vw, 4.4rem);
		font-weight: 500;
		line-height: 1.02;
		letter-spacing: -0.04em;
	}

	.os-head > p:last-child {
		max-width: 46rem;
		margin: 0 auto;
		color: rgba(255, 255, 255, 0.6);
		line-height: 1.78;
	}

	.doubts {
		--doubt-lift: clamp(1.5rem, 3.6vw, 3.5rem);

		display: grid;
		grid-template-columns: repeat(4, minmax(0, 1fr));
		gap: clamp(0.75rem, 2vw, 2rem);
		margin: clamp(3rem, 7vw, 5.5rem) 0 0;
		padding-top: var(--doubt-lift);
	}

	.doubts :global(.reveal) {
		height: 100%;
	}

	/*
	 * The two middle cards ride higher so the four of them trace an arc that
	 * follows the horizon drawn underneath. Reveal owns the wrapper's transform,
	 * so the offset goes on the blockquote instead.
	 */
	.doubts > :global(.reveal:nth-child(2) > blockquote),
	.doubts > :global(.reveal:nth-child(3) > blockquote) {
		transform: translateY(calc(var(--doubt-lift) * -1));
	}

	.doubts blockquote {
		display: flex;
		height: 100%;
		flex-direction: column;
		align-items: center;
		margin: 0;
		padding: 0 clamp(0.5rem, 1.5vw, 1.25rem);
		text-align: center;
	}

	.doubt-num {
		color: var(--aqua);
		font-size: 0.65rem;
		letter-spacing: 0.14em;
	}

	.doubt-icon {
		display: grid;
		margin: 1.5rem 0 1.9rem;
		justify-items: center;
		align-items: center;
	}

	/*
	 * Every icon's viewBox is trimmed to its artwork, so locking the height gives
	 * them all the same optical size. A max-width cap would shrink the wider ones.
	 */
	.doubt-icon img {
		width: auto;
		height: 5.6rem;
		filter: drop-shadow(0 0.7rem 1.5rem rgba(82, 222, 211, 0.22));
		transition: transform 320ms cubic-bezier(0.22, 1, 0.36, 1);
	}

	.doubts blockquote:hover .doubt-icon img {
		transform: translateY(-0.3rem) scale(1.05);
	}

	.doubts p {
		margin: auto 0 0;
		color: rgba(255, 255, 255, 0.86);
		font-family: var(--font-display);
		font-size: clamp(1rem, 1.5vw, 1.2rem);
		line-height: 1.42;
	}

	.horizon {
		position: relative;
		margin-top: clamp(1.25rem, 3vw, 2.5rem);
	}

	/*
	 * The glow is centred inside its own box and fades out well before the edges.
	 * Anchoring the ellipse to an edge instead clips half of it into a hard line.
	 */
	.horizon::before {
		position: absolute;
		top: 55%;
		left: 50%;
		width: min(68rem, 92%);
		height: 15rem;
		background: radial-gradient(
			ellipse at 50% 50%,
			rgba(82, 222, 211, 0.13),
			transparent 62%
		);
		content: '';
		pointer-events: none;
		translate: -50% -50%;
	}

	.horizon svg {
		display: block;
		width: 100%;
		height: clamp(2.75rem, 5.5vw, 5.5rem);
		overflow: visible;
		filter: drop-shadow(0 0 1rem rgba(82, 222, 211, 0.35));
		margin-top: -28px;
		margin-bottom: -32px;
	}

	.os-answer {
		max-width: 56rem;
		margin: clamp(1.5rem, 4vw, 3rem) auto 0;
		text-align: center;
	}

	.os-answer p {
		margin: 0;
		color: rgba(255, 255, 255, 0.66);
		font-size: clamp(1rem, 1.35vw, 1.15rem);
		line-height: 1.85;
	}

	.os-answer strong {
		color: white;
		font-weight: 600;
	}

	/* ---------------------------------------------------------------- labs */
	.labs {
		position: relative;
		padding: calc(var(--curve-h) + clamp(2.5rem, 5vw, 4.5rem)) var(--page-pad)
			clamp(3rem, 6vw, 5.5rem);
		background: var(--sand);
	}

	.labs-head {
		text-align: center;
	}

	.labs-head .section-heading {
		max-width: 22ch;
		margin: 1rem auto 0;
	}

	.labs-body {
		display: grid;
		grid-template-columns: minmax(0, 1.08fr) minmax(0, 0.92fr);
		gap: clamp(2rem, 5vw, 4rem);
		align-items: center;
		margin-top: clamp(2.5rem, 6vw, 4.5rem);
	}

	.labs-copy p {
		max-width: 42rem;
		margin: 0 0 1.4rem;
		color: var(--muted);
		font-size: clamp(0.98rem, 1.3vw, 1.1rem);
		line-height: 1.85;
	}

	.labs-copy p:last-child {
		margin-bottom: 0;
	}

	.labs-art {
		max-width: 27rem;
		margin: 0 0 0 auto;
	}

	.labs-art img {
		width: 100%;
		height: auto;
	}

	.legacy {
		display: grid;
		grid-template-columns: repeat(3, minmax(0, 1fr));
		gap: 1rem clamp(1.5rem, 4vw, 3.5rem);
		margin-top: clamp(3rem, 6vw, 5rem);
	}

	.legacy div {
		display: grid;
		gap: 0.5rem;
		padding: 1.4rem 0;
		border-top: 1px solid #d8cfc0;
	}

	.legacy strong {
		font-family: var(--font-display);
		font-size: 1.35rem;
		font-weight: 500;
		letter-spacing: -0.01em;
	}

	.legacy span {
		color: var(--muted);
		font-size: 0.85rem;
		line-height: 1.6;
	}

	/* ------------------------------------------------------------------ ai */
	.ai {
		position: relative;
		display: grid;
		grid-template-columns: minmax(0, 1fr) minmax(19rem, 0.85fr);
		gap: clamp(2.5rem, 8vw, 7rem);
		align-items: center;
		padding: calc(var(--curve-h) + clamp(2.5rem, 5vw, 4.5rem)) var(--page-pad)
			clamp(3rem, 6vw, 5.5rem);
		background:
			radial-gradient(circle at 12% 85%, rgba(125, 98, 217, 0.22), transparent 30rem),
			#101523;
		color: white;
	}

	.ai-copy h2 {
		max-width: 14ch;
		margin: 1rem 0 1.6rem;
		font-family: var(--font-display);
		font-size: clamp(2.2rem, 4.6vw, 4.2rem);
		font-weight: 500;
		line-height: 1.02;
		letter-spacing: -0.04em;
	}

	.ai-copy p:not(.eyebrow) {
		max-width: 40rem;
		margin: 0 0 1.2rem;
		color: rgba(255, 255, 255, 0.6);
		line-height: 1.8;
	}

	.ai-stats {
		display: grid;
		grid-template-columns: 1fr 1fr;
		gap: 1.4rem;
	}

	.ai-stats div {
		display: grid;
		gap: 0.55rem;
		padding-top: 1.1rem;
		border-top: 1px solid rgba(255, 255, 255, 0.18);
	}

	.ai-stats strong {
		color: var(--aqua);
		font-family: var(--font-display);
		font-size: clamp(1.5rem, 2.6vw, 2.2rem);
		font-weight: 500;
	}

	.ai-stats span {
		color: rgba(255, 255, 255, 0.62);
		font-size: 0.8rem;
		line-height: 1.5;
	}

	/* --------------------------------------------------------- responsive */
	@media (max-width: 1080px) {
		.project-grid {
			grid-template-columns: repeat(2, minmax(0, 1fr));
		}

		.doubts {
			/* The arc only reads with the four cards on one row. */
			--doubt-lift: 0rem;

			grid-template-columns: repeat(2, minmax(0, 1fr));
			gap: 2.5rem 1.5rem;
		}
	}

	@media (max-width: 980px) {
		.hero-kicker {
			gap: 0.6rem;
			font-size: 0.6rem;
			letter-spacing: 0.17em;
		}

		.hero-kicker span {
			width: 1.5rem;
		}
	}

	@media (max-width: 860px) {
		.hero {
			padding-top: calc(var(--header-height) + 2.5rem);
			background-position: center 45%;
		}

		.hero::before {
			background: linear-gradient(0deg, rgba(2, 8, 16, 0.88) 24%, rgba(2, 8, 16, 0.42));
		}

		.hero-inner {
			grid-template-columns: 1fr;
			gap: 2.5rem;
		}

		.hero-mark {
			order: -1;
			width: min(12.5rem, 62%);
		}

		.hero h1 {
			max-width: 100%;
		}

		.intro,
		.labs-body,
		.ai {
			grid-template-columns: 1fr;
		}

		.legacy {
			grid-template-columns: 1fr;
		}

		.section-top {
			display: grid;
			align-items: start;
		}
	}

	@media (max-width: 640px) {
		.project-grid,
		.doubts,
		.ai-stats {
			grid-template-columns: 1fr;
		}

		.doubt-icon {
			margin: 1.3rem 0 1.5rem;
		}

		.hero h1 {
			font-size: clamp(1.85rem, 7.6vw, 2.4rem);
			text-wrap: balance;
		}

		.hero h1 br {
			display: none;
		}

		.hero-foot {
			font-size: 0.55rem;
			letter-spacing: 0.16em;
		}

		.hero-foot-line {
			display: none;
		}
	}

	@media (prefers-reduced-motion: reduce) {
		.hero-mark img,
		.orbit-a {
			animation: none;
		}
	}
</style>
