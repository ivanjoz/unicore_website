<script lang="ts">
	import { base } from '$app/paths';
	import Contact from '$lib/components/Contact.svelte';
	import Reveal from '$lib/components/Reveal.svelte';
	import SectionCurve from '$lib/components/SectionCurve.svelte';
	import T from '$lib/components/T.svelte';
	import { t } from '$lib/i18n.svelte';
	import { labProjects, statusLabels } from '$lib/data/site';

	const doubts = [
		{
			text: '¿La empresa proveedora seguirá existiendo en los próximos años?|Will the vendor still be around in the coming years?',
			icon: '/svg/icons/icon-continuity.svg'
		},
		{
			text: '¿Puedo mover mis datos a mi propia nube?|Can I move my data to my own cloud?',
			icon: '/svg/icons/icon-portability.svg'
		},
		{
			text: '¿Por qué no puedo modificar el sistema para adecuarlo a mis procesos?|Why can I not modify the system to fit my own processes?',
			icon: '/svg/icons/icon-processes.svg'
		},
		{
			text: '¿Dependo de un solo proveedor para cualquier soporte?|Do I depend on a single vendor for any kind of support?',
			icon: '/svg/icons/icon-support.svg'
		}
	];

	// Módulos de Genix. `icon: ''` deja el hueco marcado en la tarjeta para colocar
	// un icono nuevo: basta con apuntarlo a un archivo dentro de /static/svg.
	const genixFeatures = [
		{
			title: 'Productos e insumos|Products and supplies',
			description:
				'Catálogo con presentaciones, lotes, series e insumos de producción.|Catalog with presentations, batches, serial numbers and production supplies.',
			icon: '/svg/producto1.svg'
		},
		{
			title: 'Almacenes y logística|Warehouses and logistics',
			description:
				'Stock por almacén, órdenes de compra, recepciones y proveedores.|Stock by warehouse, purchase orders, goods receipts and suppliers.',
			icon: '/svg/almacen1.svg'
		},
		{
			title: 'Punto de venta|Point of sale',
			description:
				'Cobro en segundos, venta al crédito y descuento de stock automático.|Checkout in seconds, credit sales and automatic stock deduction.',
			icon: '/svg/ventas1.svg'
		},
		{
			title: 'Clientes y CRM|Customers and CRM',
			description:
				'Clientes y proveedores con RUC o DNI, historial y reporte por cliente.|Customers and suppliers by tax or national ID, with history and per-customer reporting.',
			icon: '/svg/people2.svg'
		},
		{
			title: 'Caja y finanzas|Cash and finance',
			description:
				'Cajas y bancos, gastos, cuentas por cobrar y flujo de caja proyectado.|Cash registers and banks, expenses, receivables and projected cash flow.',
			icon: '/svg/finanzas2.svg'
		},
		{
			title: 'Contabilidad|Accounting',
			description:
				'Facturación, libros, activos, estados financieros y balance general.|Invoicing, ledgers, assets, financial statements and balance sheet.',
			icon: '/svg/finanzas3.svg'
		},
		{
			title: 'Indicadores y gráficos|Metrics and charts',
			description:
				'Dashboards que parten de lo general y bajan al detalle con un click.|Dashboards that start from the big picture and drill down to the detail in one click.',
			icon: '/svg/idea.svg'
		},
		{
			title: 'Comercio electrónico|E-commerce',
			description:
				'Tienda en línea armada con IA sobre el mismo catálogo del ERP.|Online store assembled with AI on top of the very same ERP catalog.',
			icon: ''
		}
	];

	const legacy = [
		{
			place: 'Bell Labs',
			note: 'El transistor, UNIX y la teoría de la información.|The transistor, UNIX and information theory.'
		},
		{
			place: 'Xerox PARC',
			note: 'La interfaz gráfica, el ratón y la red local.|The graphical interface, the mouse and the local network.'
		},
		{
			place: 'Google Research',
			note: '«Attention is All You Need», el origen de la IA moderna.|«Attention is All You Need», the origin of modern AI.'
		}
	];

	/*
	 * Textos con etiquetas propias: viven aquí porque llevan comillas dobles dentro
	 * y se pasan a <T html /> por variable en lugar de escribirlos en el atributo.
	 */
	// El espacio tras cada <br /> es intencional: en móvil el salto se oculta
	// (.hero h1 br { display: none }) y sin él las palabras quedarían pegadas.
	const heroHeading =
		'Iniciativas de <em>código abierto</em> que<br /> democratizan el acceso<br /> a tecnología de alto impacto|<em>Open source</em> initiatives that<br /> democratize access<br /> to high-impact technology';

	const introHeading =
		'Un equipo profesional<br />descentralizado.|A decentralized team<br />of professionals.';

	const openSourceLead =
		'Liberar el código amplía el mercado, acelera la adopción, multiplica la visibilidad, construye comunidad y hace el software más seguro porque cualquiera puede auditarlo. Linux, Red Hat, Valve, WordPress, GitLab y Odoo lo probaron. <strong>El open source resuelve lo siguiente:</strong>|Releasing the code widens the market, speeds up adoption, multiplies visibility, builds community and makes the software safer, because anyone can audit it. Linux, Red Hat, Valve, WordPress, GitLab and Odoo proved it. <strong>Open source solves the following:</strong>';

	const openSourceAnswer =
		'Con open source, <strong>el proveedor deja de ser el cuello de botella</strong>, el soporte se descentraliza, el sistema es más flexible y la IA es más autónoma al poder inspeccionar el código. Estos proyectos crecen de forma orgánica con la colaboración de profesionales, investigadores y empresas SaaS que los incorporan como base de sus servicios.|With open source, <strong>the vendor stops being the bottleneck</strong>, support becomes decentralized, the system is more flexible and AI is more autonomous because it can inspect the code. These projects grow organically through the collaboration of professionals, researchers and SaaS companies that adopt them as the base of their own services.';

	const labsProduct =
		'Creemos en anteponer el desarrollo de producto por sobre cualquier otro eje empresarial. <strong>El producto es la empresa y la investigación da forma al producto.</strong> Las regalías comerciales sirven para financiar al producto y su investigación, el cual <strong>responde a un impacto esperado</strong>. El sufijo «labs» abraza estos conceptos.|We believe in putting product development ahead of every other axis of the business. <strong>The product is the company, and research shapes the product.</strong> Commercial royalties are there to fund the product and its research, which <strong>answers to an expected impact</strong>. The «labs» suffix embraces these ideas.';

	const genixNote =
		'Disponible desde el 8 de Febrero 2027, con límite para uso gratuito o con plan profesional de pago. <a href="#contacto">Escríbenos</a> para ser de los primeros en usarlo.|Available from 8 February 2027, with a capped free tier or a paid professional plan. <a href="#contacto">Write to us</a> to be among the first to use it.';
</script>

<svelte:head>
	<title>Unicore Labs | {t('Código abierto de alto impacto|High-impact open source')}</title>
	<meta
		name="description"
		content={t(
			'Iniciativas de código abierto que democratizan el acceso a tecnología de alto impacto. Sistemas y herramientas que incentivan la participación comunitaria y la innovación tecnológica.|Open-source initiatives that democratize access to high-impact technology. Systems and tools that encourage community participation and technological innovation.'
		)}
	/>
	<meta
		property="og:title"
		content={`Unicore Labs | ${t('Código abierto de alto impacto|High-impact open source')}`}
	/>
	<meta
		property="og:description"
		content={t(
			'Desarrollamos sistemas y herramientas de código abierto que incentivan la participación comunitaria y la innovación tecnológica.|We build open-source systems and tools that encourage community participation and technological innovation.'
		)}
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
				<h1><T html text={heroHeading} /></h1>
				<p class="hero-lead">
					<T
						text="Desarrollamos sistemas y herramientas que incentivan la participación comunitaria y la innovación tecnológica.|We build systems and tools that encourage community participation and technological innovation."
					/>
				</p>
				<div class="hero-actions">
					<a class="button-primary" href="#proyectos">
						<T text="Ver proyectos|See projects" /> <span>↓</span>
					</a>
					<a class="button-secondary" href="#contacto">
						<T text="Colaborar|Collaborate" /> <span>↗</span>
					</a>
				</div>
			</div>

			<div class="hero-mark">
				<span class="orbit orbit-a" aria-hidden="true"></span>
				<span class="orbit orbit-b" aria-hidden="true"></span>
				<img src={`${base}/svg/logo_unicore_vertical_light_2.svg`} alt="Unicore Labs" />
			</div>
		</div>

		<div class="hero-foot">
			<span>GPL · MIT</span>
			<span class="hero-foot-line"></span>
			<span>{labProjects.length} <T text="PROYECTOS ABIERTOS|OPEN PROJECTS" /></span>
		</div>
	</section>

	<section class="intro" id="nosotros">
		<Reveal>
			<div class="intro-copy">
				<p class="eyebrow"><T text="¿Quiénes somos?|Who are we?" /></p>
				<h2 class="section-heading"><T html text={introHeading} /></h2>
				<p>
					<T
						text="Somos personas apasionadas por el desarrollo de tecnologías de código abierto y apostamos por iniciativas de interés general, con fines comerciales, científicos y/o sociales.|We are people passionate about building open-source technology, and we back initiatives of general interest with commercial, scientific and/or social aims."
					/>
				</p>
				<p>
					<T
						text="Desde sistemas de facturación, logística y finanzas para pequeñas empresas, aplicaciones de comercio electrónico y automatización de soporte con IA, hasta algoritmos de serialización para reducir el uso de ancho de banda.|From invoicing, logistics and finance systems for small companies, e-commerce applications and AI support automation, to serialization algorithms that cut bandwidth usage."
					/>
				</p>
				<div class="intro-cta">
					<p>
						<T
							text="¿Tienes un proyecto de código abierto o quieres participar en el desarrollo de alguno?|Do you have an open-source project, or would you like to take part in one?"
						/>
					</p>
					<a href="#contacto"><T text="Conversemos|Let’s talk" /> <span>→</span></a>
				</div>
			</div>
		</Reveal>
		<Reveal delay={120}>
			<figure class="intro-art">
				<img
					src={`${base}/svg/coworking.svg`}
					alt={t(
						'Tres personas trabajando con laptops en un espacio compartido|Three people working on laptops in a shared space'
					)}
					loading="lazy"
				/>
			</figure>
		</Reveal>
	</section>

	<section class="portfolio" id="proyectos">
		<SectionCurve fill="#ffffff" />

		<div class="section-top">
			<figure class="section-top-art">
				<img
					src={`${base}/svg/team-work.svg`}
					alt={t(
						'Un equipo trabajando junto sobre una pizarra de tareas|A team working together at a task board'
					)}
					loading="lazy"
				/>
			</figure>

			<div class="section-top-copy">
				<p class="eyebrow"><T text="PROYECTOS|PROJECTS" /></p>
				<h2 class="section-heading">
					<T text="Lo que estamos construyendo.|What we are building." />
				</h2>
				<p>
					<T
						text="Cada pieza resuelve un problema concreto y vive en su propio repositorio, con licencia abierta y documentación pública.|Each piece solves a specific problem and lives in its own repository, with an open license and public documentation."
					/>
				</p>
			</div>
		</div>

		<div class="project-grid">
			{#each labProjects as project, index}
				<Reveal delay={(index % 3) * 90}>
					<article class="project-card">
						<div class="project-logo">
							{#if project.logo}
								<img
									src={`${base}${project.logo}`}
									alt={`${t('Logo de|Logo of')} ${project.name}`}
								/>
							{:else}
								<span>{project.name}</span>
							{/if}
						</div>

						<div class="project-body">
							<div class="project-meta">
								<p class="project-kind"><T text={project.kind} /></p>
								<span
									class="project-status"
									class:live={project.status === 'Activo'}
									class:draft={project.status === 'En diseño'}
								>
									<T text={statusLabels[project.status]} />
								</span>
							</div>
							<h3>{project.name}</h3>
							<p class="project-text"><T text={project.description} /></p>

							<div class="project-foot">
								<ul class="project-stack">
									{#each project.stack as item}
										<li><T text={item} /></li>
									{/each}
								</ul>

								{#if project.url}
									<a class="project-link" href={project.url} target="_blank" rel="noreferrer">
										<T text="Ver|View" /> <span>↗</span>
									</a>
								{:else}
									<span class="project-link muted"><T text="Próximamente|Coming soon" /></span>
								{/if}
							</div>
						</div>
					</article>
				</Reveal>
			{/each}
		</div>
	</section>

	<section class="open-source">
		<SectionCurve fill="var(--mist)" />

		<div class="os-head">
			<p class="eyebrow light"><T text="¿Por qué código abierto?|Why open source?" /></p>
			<h2>
				<T
					html
					text="Las empresas <em>open source</em> son disruptivas por naturaleza.|<em>Open source</em> companies are disruptive by nature."
				/>
			</h2>
			<p><T html text={openSourceLead} /></p>
		</div>

		<div class="doubts">
			{#each doubts as doubt, index}
				<Reveal delay={index * 80}>
					<blockquote>
						<span class="doubt-num">{String(index + 1).padStart(2, '0')}</span>
						<div class="doubt-icon">
							<img src={`${base}${doubt.icon}`} alt="" loading="lazy" />
						</div>
						<p><T text={doubt.text} /></p>
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
						<stop offset="0" stop-color="#00d8b3" stop-opacity="0" />
						<stop offset="0.12" stop-color="#00d8b3" stop-opacity="0.55" />
						<stop offset="0.5" stop-color="#e6f6ff" stop-opacity="1" />
						<stop offset="0.88" stop-color="#a69aff" stop-opacity="0.55" />
						<stop offset="1" stop-color="#a69aff" stop-opacity="0" />
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
				<p><T html text={openSourceAnswer} /></p>
			</div>
		</Reveal>
	</section>

	<section class="labs">
		<SectionCurve fill="#0a0b1f" variant="dome" flip />

		<Reveal>
			<div class="labs-head">
				<p class="eyebrow"><T text="¿Por qué «labs»?|Why «labs»?" /></p>
				<h2 class="section-heading">
					<T text="Un espacio de innovación e investigación|A space for innovation and research" />
				</h2>
			</div>
		</Reveal>

		<div class="labs-body">
			<Reveal delay={100}>
				<div class="labs-copy">
					<div class="labs-point">
						<span
							class="labs-point-icon"
							style={`--icon: url(${base}/svg/noun_Research.svg)`}
							aria-hidden="true"
						></span>
						<span class="labs-point-rule" aria-hidden="true"></span>
						<p>
							<T
								text="Inspirados en Bell Labs y Xerox PARC. Espacios de innovación donde nacieron los fundamentos tecnológicos que dieron forma a las décadas siguientes. Y en Google Research que produjo «Attention is All You Need», el origen de la IA moderna.|Inspired by Bell Labs and Xerox PARC: spaces of innovation where the technological foundations that shaped the following decades were born. And by Google Research, which produced «Attention is All You Need», the origin of modern AI."
							/>
						</p>
					</div>
					<div class="labs-point">
						<span
							class="labs-point-icon"
							style={`--icon: url(${base}/svg/noun_innovation.svg)`}
							aria-hidden="true"
						></span>
						<span class="labs-point-rule" aria-hidden="true"></span>
						<p><T html text={labsProduct} /></p>
					</div>
					<div class="labs-point accent">
						<span
							class="labs-point-icon"
							style={`--icon: url(${base}/svg/noun_mind.svg)`}
							aria-hidden="true"
						></span>
						<span class="labs-point-rule" aria-hidden="true"></span>
						<p>
							<T
								text="Si eres una mente curiosa, con habilidades para la programación y quieres aportar al código abierto, escríbenos. Eres libre de hacer un fork de cualquiera de nuestros proyectos y usarlos como base de nuevos desarrollos. Háznoslo saber en GitHub.|If you are a curious mind with programming skills and you want to contribute to open source, write to us. You are free to fork any of our projects and use them as the base for new work. Let us know on GitHub."
							/>
						</p>
					</div>
				</div>
			</Reveal>
			<Reveal delay={160}>
				<figure class="labs-art">
					<img
						src={`${base}/svg/connecting-teams.svg`}
						alt={t(
							'Personas colaborando y compartiendo información entre sí|People collaborating and sharing information with each other'
						)}
						loading="lazy"
					/>
				</figure>
			</Reveal>
		</div>
	</section>

	<section class="ai">
		<SectionCurve fill="var(--sand)" variant="valley" />

		<figure class="ai-art">
			<img
				src={`${base}/svg/voice-assistant-dark.svg`}
				alt={t(
					'Persona conversando con un asistente de voz|A person talking to a voice assistant'
				)}
				loading="lazy"
			/>
		</figure>

		<div class="ai-copy">
			<p class="eyebrow light"><T text="Inteligencia artificial|Artificial intelligence" /></p>
			<h2><T text="IA al alcance de cualquier usuario.|AI within reach of every user." /></h2>
			<p>
				<T
					text="Enfrentamos el reto de democratizar el acceso al software reduciendo el costo de la IA por proceso. Usamos IA de frontera para desarrollar, pero implementamos métodos más determinísticos para desplegar servicios eficientes basados en IA.|We take on the challenge of democratizing access to software by cutting the cost of AI per process. We use frontier AI to build, but we deploy efficient AI-based services with more deterministic methods."
				/>
			</p>
			<p>
				<T
					text="Genix Agentic UI es un ejemplo de cómo modelos pequeños y sin visión pueden servir como agentes que ayudan al usuario a navegar el sistema y ejecutar instrucciones. Ello permite exprimir cada dólar de inferencia.|Genix Agentic UI is an example of how small, vision-free models can act as agents that help the user navigate the system and carry out instructions. That is what makes it possible to squeeze every dollar of inference."
				/>
			</p>
		</div>

		<SectionCurve fill="#ffffff" variant="wave" side="bottom" flip />
	</section>

	<section class="genix">

		<div class="genix-intro">
			<img class="genix-logo" src={`${base}/svg/genix_logo.svg`} alt="Genix" />
			<a class="genix-cta" href="https://genix.un.pe/" target="_blank" rel="noreferrer">
				<T text="Ir a la aplicación|Go to the app" /> <span>↗</span>
			</a>

			<span class="genix-rule" aria-hidden="true"></span>

			<p class="eyebrow"><T text="Producto insignia|Flagship product" /></p>
			<h2><T text="Gestione cada proceso de su empresa.|Manage every process in your company." /></h2>
			<p class="genix-lead">
				<T
					text="Genix reúne ventas, inventario, compras, finanzas y comercio electrónico para pequeñas empresas que buscan control sin complejidad. Autoalojable, multiempresa y con exportación completa de sus datos cuando quiera.|Genix brings sales, inventory, purchasing, finance and e-commerce together for small companies that want control without complexity. Self-hostable, multi-company, and with a full export of your data whenever you want."
				/>
			</p>

			<p class="genix-note">
				<span class="genix-note-mark" aria-hidden="true"></span>
				<T html text={genixNote} />
			</p>
		</div>

		<div class="feature-grid">
			{#each genixFeatures as feature, index}
				<Reveal delay={(index % 4) * 70}>
					<article>
						{#if feature.icon}
							<img src={`${base}${feature.icon}`} alt="" loading="lazy" />
						{:else}
							<span class="feature-icon-slot" aria-hidden="true">
								<T html text="icono<br />pendiente|icon<br />pending" />
							</span>
						{/if}
						<h3><T text={feature.title} /></h3>
						<p><T text={feature.description} /></p>
					</article>
				</Reveal>
			{/each}
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
		background: #05061a var(--hero-jpg) center 32% / cover no-repeat;
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
			linear-gradient(100deg, rgba(5, 6, 26, 0.9) 0%, rgba(5, 6, 26, 0.56) 46%, rgba(5, 6, 26, 0.24) 100%),
			linear-gradient(0deg, rgba(5, 6, 26, 0.72), transparent 45%);
		content: '';
	}

	.hero-grid {
		position: absolute;
		inset: 0;
		background-image:
			linear-gradient(rgba(150, 170, 255, 0.06) 1px, transparent 1px),
			linear-gradient(90deg, rgba(150, 170, 255, 0.06) 1px, transparent 1px);
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
		font-size: calc(0.68rem + var(--fs-bump));
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
		font-size: clamp(1.5rem, 2.9vw, 2.9rem);
		font-weight: 500;
		line-height: 1.15;
		letter-spacing: -0.035em;
	}

	.hero h1 :global(em) {
		position: relative;
		background: var(--brand-ramp);
		-webkit-background-clip: text;
		background-clip: text;
		color: transparent;
		font-style: normal;
	}

	.hero-lead {
		max-width: 34rem;
		margin: 1.8rem 0 0;
		color: rgba(228, 230, 255, 0.72);
		font-size: calc(clamp(1rem, 1.35vw, 1.15rem) + var(--fs-bump));
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
			rgba(4, 5, 22, 0.78) 0%,
			rgba(4, 5, 22, 0.68) 30%,
			rgba(8, 8, 34, 0.38) 52%,
			transparent 74%
		);
		content: '';
		translate: -50% -50%;
	}

	.orbit {
		position: absolute;
		top: 44.2%;
		left: 50%;
		border: 1px solid rgba(166, 154, 255, 0.24);
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
		border-color: rgba(166, 154, 255, 0.12);
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
		color: rgba(228, 230, 255, 0.45);
		font-size: calc(0.6rem + var(--fs-bump));
		letter-spacing: 0.24em;
	}

	.hero-foot-line {
		width: clamp(2rem, 12vw, 9rem);
		height: 1px;
		background: rgba(228, 230, 255, 0.22);
	}

	/* --------------------------------------------------------------- intro */
	.intro {
		display: grid;
		background: #ffffff;
		grid-template-columns: minmax(0, 1.12fr) minmax(0, 0.88fr);
		gap: clamp(1.75rem, 3.2vw, 3.25rem);
		align-items: center;
		padding: clamp(5rem, 10vw, 7rem) var(--page-pad) clamp(3rem, 6vw, 5.5rem);
		margin-bottom: -1rem;
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
		font-size: calc(clamp(0.98rem, 1.3vw, 1.1rem) + var(--fs-bump));
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
		border-left: 3px solid var(--accent);
		border-radius: 0 1rem 1rem 0;
		background: var(--mist);
	}

	.intro-cta p {
		flex: 1 1 20rem;
		margin: 0;
		color: var(--ink);
		font-size: calc(clamp(0.95rem, 1.2vw, 1.05rem) + var(--fs-bump));
		line-height: 1.5;
	}

	.intro-cta a {
		display: inline-flex;
		flex: none;
		gap: 0.7rem;
		align-items: center;
		padding: 0.9rem 1.4rem;
		border-radius: 999px;
		background: var(--accent);
		color: white;
		font-size: calc(0.75rem + var(--fs-bump));
		font-weight: 800;
		letter-spacing: 0.1em;
		text-decoration: none;
		text-transform: uppercase;
		transition:
			gap 180ms ease,
			background 180ms ease;
	}

	.intro-cta a:hover {
		background: var(--accent-dark);
		gap: 1.1rem;
	}

	/* ----------------------------------------------------------- portfolio */
	.portfolio {
		position: relative;
		/* Por encima de la sección siguiente: si no, su fondo recorta la sombra
		   que las tarjetas proyectan por debajo de la última fila. */
		z-index: 10;
		padding-top: calc(var(--curve-h) + 10px);
		padding-right: var(--page-pad);
		padding-bottom: 3rem;
		padding-left: var(--page-pad);
		background: var(--mist);
	}

	.section-top {
		display: grid;
		grid-template-columns: minmax(0, 0.82fr) minmax(0, 1.18fr);
		gap: clamp(2rem, 5vw, 4.5rem);
		align-items: center;
		margin-bottom: clamp(2rem, 4.5vw, 3.5rem);
	}

	.section-top-art {
		max-width: 26rem;
		margin: 0;
	}

	.section-top-art img {
		width: 100%;
		height: auto;
	}

	/* El titular ya no compite con la columna de texto: puede usar todo su ancho. */
	.section-top-copy .section-heading {
		max-width: none;
	}

	.section-top-copy p:not(.eyebrow) {
		max-width: 32rem;
		margin: 1.4rem 0 0;
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
		/* Borde transparente, no ausente: reservarlo evita que el hover mueva el layout. */
		border: 1px solid transparent;
		border-radius: var(--radius-lg);
		background: white;
		box-shadow: var(--shadow);
		transition:
			transform 250ms ease,
			box-shadow 250ms ease;
	}

	.project-card:hover {
		box-shadow: 0 34px 90px rgba(28, 22, 74, 0.18);
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
			radial-gradient(circle at 30% 20%, rgba(0, 216, 179, 0.12), transparent 60%),
			linear-gradient(160deg, #f8f7fe, #edeffb);
	}

	.project-logo img {
		max-width: 12rem;
    max-height: 4.8rem;
		object-fit: contain;
	}

	.project-logo span {
		padding-inline: 1rem;
		color: #9a9cbe;
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
		color: var(--accent);
		font-size: calc(0.6rem + var(--fs-bump));
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
		background: #fbfaff;
		color: var(--muted);
		font-size: calc(0.58rem + var(--fs-bump));
		font-weight: 700;
		letter-spacing: 0.08em;
		text-transform: uppercase;
	}

	.project-status::before {
		width: 0.4rem;
		height: 0.4rem;
		border-radius: 50%;
		background: #b9bcd8;
		content: '';
	}

	.project-status.live::before {
		background: #0fbf9f;
	}

	.project-status.draft::before {
		background: var(--brand-magenta);
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
		font-size: calc(0.94rem + var(--fs-bump));
		line-height: 1.65;
	}

	/* Los tags y el enlace comparten fila; `margin-top: auto` la fija al pie. */
	.project-foot {
		display: flex;
		align-items: center;
		justify-content: space-between;
		gap: 0.75rem;
		margin-top: auto;
		padding-top: 1.2rem;
	}

	.project-stack {
		display: flex;
		min-width: 0;
		flex-wrap: wrap;
		gap: 0.4rem;
		margin: 0;
		padding: 0;
		list-style: none;
	}

	.project-stack li {
		padding: 0.28rem 0.6rem;
		border-radius: 0.45rem;
		background: var(--mist);
		color: var(--muted);
		font-size: calc(0.65rem + var(--fs-bump));
		letter-spacing: 0.04em;
	}

	.project-link {
		display: inline-flex;
		align-items: center;
		flex: none;
		gap: 0.35rem;
		color: var(--accent);
		white-space: nowrap;
		font-size: calc(0.7rem + var(--fs-bump));
		font-weight: 800;
		letter-spacing: 0.1em;
		text-decoration: none;
		text-transform: uppercase;
		transition: gap 180ms ease;
	}

	.project-link:hover {
		gap: 0.7rem;
	}

	.project-link.muted {
		color: #9a9cbe;
	}

	/* --------------------------------------------------------- open source */
	.open-source {
		position: relative;
		padding: calc(var(--curve-h) + clamp(2.5rem, 5vw, 4.5rem)) var(--page-pad)
			clamp(3rem, 6vw, 5.5rem);
		background:
			radial-gradient(circle at 80% 15%, rgba(100, 105, 238, 0.26), transparent 32rem),
			#0a0b1f;
		color: white;
	}

	.os-head {
		text-align: center;
	}

	.os-head h2 :global(em) {
		background: var(--brand-ramp);
		-webkit-background-clip: text;
		background-clip: text;
		color: transparent;
		font-style: normal;
	}

	.os-head h2 {
		max-width: 26ch;
		margin: 0.6rem auto 1.8rem;
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
		font-size: calc(1.05rem + var(--fs-bump));
	}

	.os-head > p:last-child :global(strong) {
		color: rgba(255, 255, 255, 0.92);
		font-weight: 600;
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
		font-size: calc(0.65rem + var(--fs-bump));
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
		filter: drop-shadow(0 0.7rem 1.5rem rgba(0, 216, 179, 0.24));
		transition: transform 320ms cubic-bezier(0.22, 1, 0.36, 1);
	}

	.doubts blockquote:hover .doubt-icon img {
		transform: translateY(-0.3rem) scale(1.05);
	}

	.doubts p {
		margin: auto 0 0;
		color: rgba(255, 255, 255, 0.86);
		font-family: var(--font-display);
		font-size: calc(clamp(1rem, 1.5vw, 1.2rem) + var(--fs-bump));
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
			rgba(0, 216, 179, 0.14),
			transparent 62%
		);
		content: '';
		pointer-events: none;
		translate: -50% -50%;
	}

	/*
	 * El trazo sólo ocupa la franja baja del viewBox, así que los márgenes negativos
	 * recortan el aire sobrante de la caja. Van en proporción a la altura: en píxeles
	 * fijos, al encoger el arco en móvil se restaba más de lo que medía y el arco se
	 * montaba sobre el texto de arriba y de abajo.
	 */
	.horizon svg {
		--horizon-h: clamp(2.75rem, 5.5vw, 5.5rem);

		display: block;
		width: 100%;
		height: var(--horizon-h);
		overflow: visible;
		filter: drop-shadow(0 0 1rem rgba(0, 216, 179, 0.38));
		margin-top: calc(var(--horizon-h) * -0.32);
		margin-bottom: calc(var(--horizon-h) * -0.36);
	}

	.os-answer {
		max-width: 56rem;
		margin: clamp(1.5rem, 4vw, 3rem) auto 0;
		text-align: center;
	}

	.os-answer p {
		margin: 0;
		color: rgba(255, 255, 255, 0.66);
		font-size: calc(clamp(1rem, 1.35vw, 1.15rem) + var(--fs-bump));
		line-height: 1.7;
	}

	.os-answer :global(strong) {
		color: white;
		font-weight: 600;
	}

	/* ---------------------------------------------------------------- labs */
	.labs {
		position: relative;
		padding-top: 120px;
		padding-right: var(--page-pad);
		padding-bottom: 30px;
		padding-left: var(--page-pad);
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

	.labs-copy {
		display: grid;
		gap: clamp(1.7rem, 3.4vw, 2.6rem);
	}

	.labs-point {
		display: grid;
		grid-template-columns: auto 2px minmax(0, 1fr);
		gap: clamp(0.9rem, 2vw, 1.35rem);
		align-items: center;
	}

	.labs-point-icon {
		width: clamp(2.3rem, 3.6vw, 3.1rem);
		aspect-ratio: 1;
		background: var(--accent);
		-webkit-mask: var(--icon) center / contain no-repeat;
		mask: var(--icon) center / contain no-repeat;
		opacity: 0.75;
	}

	.labs-point-rule {
		align-self: stretch;
		border-radius: 2px;
		background: linear-gradient(
			180deg,
			rgba(100, 105, 238, 0.08),
			var(--brand-blue) 26%,
			var(--brand-violet) 74%,
			rgba(166, 154, 255, 0.08)
		);
	}

	.labs-point.accent .labs-point-icon {
		opacity: 1;
	}

	.labs-point.accent p {
		color: var(--ink);
	}

	.labs-copy p {
		margin: 0;
		color: var(--muted);
		font-size: calc(clamp(0.98rem, 1.3vw, 1.1rem) + var(--fs-bump));
		line-height: 1.85;
	}

	.labs-copy :global(strong) {
		color: var(--ink);
		font-weight: 600;
		background-image: linear-gradient(
			transparent 62%,
			rgba(0, 216, 179, 0.34) 62%,
			rgba(0, 216, 179, 0.34) 92%,
			transparent 92%
		);
		box-decoration-break: clone;
		-webkit-box-decoration-break: clone;
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
		border-top: 1px solid var(--line);
	}

	.legacy strong {
		font-family: var(--font-display);
		font-size: 1.35rem;
		font-weight: 500;
		letter-spacing: -0.01em;
	}

	.legacy span {
		color: var(--muted);
		font-size: calc(0.85rem + var(--fs-bump));
		line-height: 1.6;
	}

	/* ------------------------------------------------------------------ ai */
	/* --------------------------------------------------------------- genix */
	.genix {
		position: relative;
		padding: clamp(3rem, 6vw, 5rem) var(--page-pad) clamp(3.5rem, 7vw, 6rem);
		background: #ffffff;
		color: var(--ink);
	}

	/*
	 * Una sola rejilla para toda la cabecera. Las filas se declaran a mano porque
	 * la columna izquierda (logo y botón) y la derecha (rótulo, titular y texto)
	 * no llevan el mismo número de elementos, y el filete tiene que cruzarlas todas.
	 */
	.genix-intro {
		display: grid;
		grid-template-columns: auto 1px minmax(0, 1fr);
		gap: 0 clamp(1.5rem, 4vw, 3.5rem);
		align-items: start;
	}

	.genix-logo {
		grid-row: 1 / 3;
		grid-column: 1;
		width: min(15rem, 100%);
		align-self: center;
	}

	.genix-cta {
		display: inline-flex;
		grid-row: 3;
		grid-column: 1;
		align-items: center;
		justify-self: start;
		gap: 0.5rem;
		margin-top: clamp(1.2rem, 2.5vw, 1.8rem);
		padding: 0.75rem 1.15rem;
		border: 1px solid var(--line);
		border-radius: 999px;
		background: var(--mist);
		color: var(--accent);
		font-size: calc(0.72rem + var(--fs-bump));
		font-weight: 800;
		letter-spacing: 0.1em;
		text-decoration: none;
		text-transform: uppercase;
		transition:
			background 180ms ease,
			border-color 180ms ease,
			transform 180ms ease;
	}

	.genix-cta:hover {
		border-color: var(--accent);
		background: white;
		transform: translateY(-2px);
	}

	/*
	 * Único elemento naranja del sitio: es un aviso de disponibilidad y tiene que
	 * leerse aparte de la paleta violeta, no integrarse en ella.
	 */
	.genix-note {
		position: relative;
		grid-row: 4;
		grid-column: 3;
		/* Misma medida que el párrafo de arriba: los dos bloques alinean sus bordes. */
		max-width: 40rem;
		margin: clamp(1.4rem, 2.5vw, 1.9rem) 0 0;
		padding: 1.1rem 1.2rem 1.1rem 2.6rem;
		border: 1px solid #f6d3b8;
		border-radius: 0.9rem;
		background: #fff5ed;
		color: #7a3d10;
		font-size: calc(0.84rem + var(--fs-bump));
		line-height: 1.65;
	}

	.genix-note-mark {
		position: absolute;
		top: 1.45rem;
		left: 1.15rem;
		width: 0.62rem;
		height: 0.62rem;
		border-radius: 0.15rem;
		background: #f97316;
	}

	.genix-note :global(a) {
		color: #b4470b;
		font-weight: 600;
		text-underline-offset: 0.2em;
	}

	.genix-rule {
		grid-row: 1 / -1;
		grid-column: 2;
		align-self: stretch;
		background: linear-gradient(180deg, transparent, var(--line) 8%, var(--line) 92%, transparent);
	}

	.genix-intro .eyebrow {
		grid-row: 1;
		grid-column: 3;
		/*
		 * El margen negativo cancela su propia caja para que la fila 1 no aporte
		 * altura: sólo con la fila a cero el titular y el logo comparten centro.
		 * El desplazamiento lo sube por encima del titular, y al ser un transform
		 * no vuelve a ocupar sitio en la rejilla.
		 */
		margin-bottom: -1rem;
		transform: translateY(-1.25rem);
	}

	.genix-intro h2 {
		grid-row: 2;
		grid-column: 3;
		margin: 0;
		font-family: var(--font-display);
		font-size: clamp(2.2rem, 4.6vw, 4.2rem);
		font-weight: 500;
		line-height: 1.02;
		letter-spacing: -0.04em;
	}

	.genix-lead {
		grid-row: 3;
		grid-column: 3;
		max-width: 40rem;
		margin: clamp(1.2rem, 2.5vw, 1.8rem) 0 0;
		color: var(--muted);
		line-height: 1.8;
	}

	.feature-grid {
		display: grid;
		grid-template-columns: repeat(4, minmax(0, 1fr));
		gap: 1px;
		margin-top: clamp(3rem, 7vw, 5.5rem);
		background: var(--line);
	}

	.feature-grid article {
		display: flex;
		min-height: 17rem;
		flex-direction: column;
		padding: 1.6rem;
		background: #ffffff;
	}

	.feature-grid img {
		width: 4.6rem;
		height: 4.6rem;
		margin: 0.4rem 0 1.9rem;
		object-fit: contain;
		/* los iconos están dibujados en blanco para fondo oscuro */
		filter: invert(1) opacity(0.82);
	}

	.feature-icon-slot {
		display: grid;
		width: 4.6rem;
		height: 4.6rem;
		margin: 0.4rem 0 1.9rem;
		border: 1px dashed var(--line);
		border-radius: 0.5rem;
		color: rgba(92, 110, 119, 0.55);
		font-size: 0.6rem;
		letter-spacing: 0.08em;
		place-content: center;
		text-align: center;
		text-transform: uppercase;
	}

	.feature-grid h3 {
		margin: auto 0 0.6rem;
		font-family: var(--font-display);
		font-size: 1.1rem;
		font-weight: 500;
	}

	.feature-grid p {
		margin: 0;
		color: var(--muted);
		font-size: calc(0.82rem + var(--fs-bump));
		line-height: 1.6;
	}

	.ai {
		position: relative;
		display: grid;
		grid-template-columns: minmax(0, 0.92fr) minmax(0, 1.08fr);
		gap: clamp(2.5rem, 7vw, 6rem);
		align-items: center;
		padding: calc(var(--curve-h) + clamp(2.5rem, 5vw, 4.5rem)) var(--page-pad)
			calc(var(--curve-h) + clamp(2rem, 4vw, 3.5rem));
		background:
			radial-gradient(circle at 12% 85%, rgba(125, 98, 217, 0.26), transparent 30rem),
			#0f1030;
		color: white;
	}

	.ai-art {
		max-width: 32rem;
		margin: 0;
	}

	.ai-art img {
		width: 100%;
		height: auto;
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
			background: linear-gradient(0deg, rgba(5, 6, 26, 0.88) 24%, rgba(5, 6, 26, 0.42));
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
		.genix-intro {
			grid-template-columns: 1fr;
		}

		/*
		 * La ilustración va entre el titular y el cuerpo de texto, pero en el
		 * marcado es hermana de `.ai-copy`. `display: contents` disuelve esa caja
		 * para que rótulo, titular y párrafos sean hermanos directos de la
		 * ilustración y `order` pueda intercalarla.
		 */
		.ai {
			display: flex;
			flex-direction: column;
			align-items: stretch;
			gap: 0;
		}

		.ai-copy {
			display: contents;
		}

		.ai-copy .eyebrow {
			order: 1;
		}

		.ai-copy h2 {
			order: 2;
			max-width: none;
			margin-bottom: 0;
		}

		.ai-art {
			order: 3;
			max-width: 22rem;
			margin: clamp(1.5rem, 5vw, 2.25rem) auto;
		}

		.ai-copy p:not(.eyebrow) {
			order: 4;
		}

		/* El titular por delante de la ilustración al apilarse. */
		.section-top-copy {
			order: -1;
		}

		/*
		 * El `padding-top` de escritorio es un valor fijo pensado para una cúpula de
		 * 8rem; al encoger la curva con el ancho dejaba un hueco creciente.
		 */
		.labs {
			padding-top: calc(var(--curve-h) + 1.5rem);
		}

		/* La ilustración sube justo debajo del titular de la sección. */
		.labs-body > :global(.reveal:last-child) {
			order: -1;
		}

		.labs-art {
			max-width: 20rem;
			margin: 0 auto;
		}

		/* Con el texto a todo el ancho, centrar el icono lo deja flotando. */
		.labs-point {
			align-items: start;
		}

		/*
		 * El desplazamiento que alinea el rótulo con el logo sólo tiene sentido con
		 * las dos columnas; apilado, monta el rótulo sobre el botón.
		 */
		.genix-intro .eyebrow {
			margin-bottom: 0;
			transform: none;
		}

		/* Sin las tres columnas, el filete no separa nada y el resto se apila. */
		.genix-rule {
			display: none;
		}

		/* Una sola columna: las filas explícitas dejan de servir y todo se apila. */
		.genix-logo,
		.genix-cta,
		.genix-note,
		.genix-intro .eyebrow,
		.genix-intro h2,
		.genix-lead {
			grid-row: auto;
			grid-column: 1;
		}

		.genix-cta {
			margin-bottom: clamp(1.2rem, 3vw, 2rem);
		}

		.feature-grid {
			grid-template-columns: repeat(2, minmax(0, 1fr));
		}

		.legacy {
			grid-template-columns: 1fr;
		}

		.section-top {
			grid-template-columns: 1fr;
			align-items: start;
		}
	}

	@media (max-width: 640px) {
		.project-grid,
		.doubts,
		.feature-grid {
			grid-template-columns: 1fr;
		}

		/* Apilado, el texto y el botón no necesitan el aire de la versión ancha. */
		.intro-cta {
			gap: 0.9rem;
			padding: 1.15rem 1.2rem;
		}

		/* Centrado y a una medida corta, el texto pedía separarse de los bordes. */
		.os-answer {
			margin-top: 2.25rem;
			padding: 0 1.1rem;
		}

		/* El arco es mucho más bajo aquí: necesita su propio aire a ambos lados. */
		.horizon {
			margin-top: 2.25rem;
		}

		.doubt-icon {
			margin: 1.3rem 0 1.5rem;
		}

		.hero h1 {
			font-size: clamp(1.85rem, 7.6vw, 2.4rem);
			text-wrap: balance;
		}

		.hero h1 :global(br) {
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
