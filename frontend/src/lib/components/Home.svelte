<script lang="ts">
	import { base } from '$app/paths';
	import Contact from '$lib/components/Contact.svelte';
	import LangToggle from '$lib/components/LangToggle.svelte';
	import Reveal from '$lib/components/Reveal.svelte';
	import SectionCurve from '$lib/components/SectionCurve.svelte';
	import T from '$lib/components/T.svelte';
	import { t } from '$lib/i18n.svelte';
	import { labProjects, statusLabels } from '$lib/data/site';
	import {
		almacen1,
		ecommerce,
		finanzas2,
		finanzas3,
		genixLogo,
		iconContinuity,
		iconPortability,
		iconProcesses,
		iconSupport,
		idea,
		nounInnovation,
		nounMind,
		nounResearch,
		people2,
		producto1,
		ventas1
	} from '$lib/svg/icons.gen';

	const doubts = [
		{
			text: '¿La empresa proveedora seguirá existiendo en los próximos años?|Will the vendor still be around in the coming years?',
			icon: iconContinuity
		},
		{
			text: '¿Puedo mover mis datos a mi propia nube?|Can I move my data to my own cloud?',
			icon: iconPortability
		},
		{
			text: '¿Por qué no puedo modificar el sistema para adecuarlo a mis procesos?|Why can I not modify the system to fit my own processes?',
			icon: iconProcesses
		},
		{
			text: '¿Dependo de un solo proveedor para cualquier soporte?|Do I depend on a single vendor for any kind of support?',
			icon: iconSupport
		}
	];

	// Módulos de Genix. `icon: ''` deja el hueco marcado en la tarjeta para colocar
	// un icono nuevo: basta con apuntarlo a un archivo dentro de /static/svg.
	const genixFeatures = [
		{
			title: 'Productos e insumos|Products and supplies',
			description:
				'Catálogo con presentaciones, lotes, series e insumos de producción.|Catalog with presentations, batches, serial numbers and production supplies.',
			icon: producto1
		},
		{
			title: 'Almacenes y logística|Warehouses and logistics',
			description:
				'Stock por almacén, órdenes de compra, recepciones y proveedores.|Stock by warehouse, purchase orders, goods receipts and suppliers.',
			icon: almacen1
		},
		{
			title: 'Punto de venta|Point of sale',
			description:
				'Cobro en segundos, venta al crédito y descuento de stock automático.|Checkout in seconds, credit sales and automatic stock deduction.',
			icon: ventas1
		},
		{
			title: 'Clientes y CRM|Customers and CRM',
			description:
				'Clientes y proveedores con RUC o DNI, historial y reporte por cliente.|Customers and suppliers by tax or national ID, with history and per-customer reporting.',
			icon: people2
		},
		{
			title: 'Caja y finanzas|Cash and finance',
			description:
				'Cajas y bancos, gastos, cuentas por cobrar y flujo de caja proyectado.|Cash registers and banks, expenses, receivables and projected cash flow.',
			icon: finanzas2
		},
		{
			title: 'Contabilidad|Accounting',
			description:
				'Facturación, libros, activos, estados financieros y balance general.|Invoicing, ledgers, assets, financial statements and balance sheet.',
			icon: finanzas3
		},
		{
			title: 'Indicadores y gráficos|Metrics and charts',
			description:
				'Dashboards que parten de lo general y bajan al detalle con un click.|Dashboards that start from the big picture and drill down to the detail in one click.',
			icon: idea
		},
		{
			title: 'Comercio electrónico|E-commerce',
			description:
				'Tienda en línea armada con IA sobre el mismo catálogo del ERP.|Online store assembled with AI on top of the very same ERP catalog.',
			icon: ecommerce
		}
	];

	/*
	 * Lo que sabe hacer el agente de Genix. Va dentro de la ficha del producto, así
	 * que cada entrada es de una línea: el detalle largo vive en la app, no aquí.
	 */
	const agentSkills = [
		{
			title: 'Navega por ti|Navigates for you',
			description:
				'Abre el módulo, filtra y te deja en la vista exacta.|Opens the module, filters and lands you on the exact view.'
		},
		{
			title: 'Llena formularios|Fills in forms',
			description:
				'Le dictas los datos y él completa los campos.|You dictate the data and it fills in the fields.'
		},
		{
			title: 'Te orienta|Guides you',
			description:
				'Resuelve dudas de uso sin sacarte de la pantalla.|Answers how-to questions without pulling you off the screen.'
		},
		{
			title: 'Arma tu tienda|Builds your store',
			description:
				'Es parte del constructor de la página y el catálogo.|It is part of the builder for your page and catalog.'
		},
		{
			title: 'Reportes y gráficos|Reports and charts',
			description:
				'Pregunta por tus ventas y responde con el gráfico.|Ask about your sales and it answers with the chart.'
		},
		{
			title: 'Recordatorios|Reminders',
			description:
				'Avisa de cobros por vencer y stock por reponer.|Flags payments coming due and stock to replenish.'
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
	// El titular se corta en sitios distintos según el ancho, así que lleva los dos
	// juegos de saltos: `w` (wide) sólo se ve en pantalla ancha y `n` (narrow) sólo
	// en móvil. El espacio tras cada <br /> es intencional: el salto que sobra se
	// oculta con display:none y sin ese espacio las palabras quedarían pegadas.
	const heroHeading =
		'Iniciativas de <em>código abierto</em><br class="n" /> que<br class="w" /> expanden<br class="n" /> el acceso a<br class="w" /> tecnologías<br class="n" /> de alto impacto|<em>Open source</em> initiatives<br class="n" /> that<br class="w" /> expands access<br class="n" /> to<br class="w" /> high-impact technology';

	const introHeading =
		'Un equipo profesional<br />descentralizado.|A decentralized team<br />of professionals.';

	const openSourceLead =
		'Liberar el código amplía el mercado, acelera la adopción, multiplica la visibilidad, construye comunidad y hace el software más seguro porque cualquiera puede auditarlo. <span class="os-split">Linux, Red Hat, Valve, WordPress, GitLab y Odoo lo probaron. <strong>El open source resuelve lo siguiente:</strong></span>|Releasing the code widens the market, speeds up adoption, multiplies visibility, builds community and makes the software safer, because anyone can audit it. <span class="os-split">Linux, Red Hat, Valve, WordPress, GitLab and Odoo proved it. <strong>Open source solves the following:</strong></span>';

	const openSourceAnswer =
		'Con open source, <strong>el proveedor deja de ser el cuello de botella</strong>, el soporte se descentraliza, el sistema es más flexible y la IA es más autónoma al poder inspeccionar el código. <span class="os-split">Estos proyectos crecen de forma orgánica con la colaboración de profesionales, investigadores y empresas SaaS que los incorporan como base de sus servicios.</span>|With open source, <strong>the vendor stops being the bottleneck</strong>, support becomes decentralized, the system is more flexible and AI is more autonomous because it can inspect the code. <span class="os-split">These projects grow organically through the collaboration of professionals, researchers and SaaS companies that adopt them as the base of their own services.</span>';

	const labsProduct =
		'Creemos en anteponer el desarrollo de producto por sobre cualquier otro eje empresarial. <strong>El producto es la empresa y la investigación da forma al producto.</strong> Las regalías comerciales sirven para financiar al producto y su investigación, el cual <strong>responde a un impacto esperado</strong>. El sufijo «labs» abraza estos conceptos.|We believe in putting product development ahead of every other axis of the business. <strong>The product is the company, and research shapes the product.</strong> Commercial royalties are there to fund the product and its research, which <strong>answers to an expected impact</strong>. The «labs» suffix embraces these ideas.';

	const genixNote =
		'Disponible desde el 8 de Febrero 2027, con límite para uso gratuito o con plan profesional de pago. <a href="#contacto">Escríbenos</a> para ser de los primeros en usarlo.|Available from 8 February 2027, with a capped free tier or a paid professional plan. <a href="#contacto">Write to us</a> to be among the first to use it.';
</script>

<svelte:head>
	<!--
		Aquí sólo van los preloads del hero: título, descripción y Open Graph los
		emite el layout, que es quien sabe en qué idioma está la URL.

		Un preload por condición, nunca dos para el mismo ancho: el atributo `type`
		hace que el navegador ignore el que no soporta, así que listar AVIF y WebP
		a la vez haría que Chrome —que soporta ambos— se descargara los dos.
	-->
	<link
		rel="preload"
		as="image"
		type="image/avif"
		media="(max-width: 720px)"
		href={`${base}/images/space_earth_mobile.avif`}
	/>
	<link
		rel="preload"
		as="image"
		type="image/avif"
		media="(min-width: 721px)"
		href={`${base}/images/space_earth.avif`}
	/>
</svelte:head>

<main>
	<section class="hero">
		<!--
			En vertical la foto apaisada del hero se recortaba a una franja, así que
			bajo 720px entra una toma en formato retrato. Va en <picture> y no como
			`background-image` porque `image-set()` no permite condicionar por ancho.
		-->
		<picture class="hero-bg">
			<source
				media="(max-width: 720px)"
				type="image/avif"
				srcset={`${base}/images/space_earth_mobile.avif`}
			/>
			<source
				media="(max-width: 720px)"
				type="image/webp"
				srcset={`${base}/images/space_earth_mobile.webp`}
			/>
			<source type="image/avif" srcset={`${base}/images/space_earth.avif`} />
			<!--
				El respaldo del <img> es WebP: cualquier navegador que entienda
				<picture> lo descodifica, así que un JPEG extra sólo añadiría peso al
				repositorio sin ampliar la cobertura.
			-->
			<img src={`${base}/images/space_earth.webp`} alt="" fetchpriority="high" decoding="async" />
		</picture>

		<div class="hero-grid" aria-hidden="true"></div>

		<!-- Sin cabecera en móvil, el selector de idioma se ancla al hero. -->
		<div class="hero-lang"><LangToggle /></div>

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
				<p class="intro-more">
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

				<p class="project-hint">
					<T text="Haz click en|Click on" />
					<span class="icon-[lucide--external-link]" aria-hidden="true"></span>
					<T text="para ir a la página de cada proyecto.|to go to each project’s page." />
				</p>
			</div>
		</div>

		<div class="project-grid">
			{#each labProjects as project, index}
				<Reveal delay={(index % 3) * 90}>
					<article class="project-card">
						{#if project.url}
							<a
								class="project-open"
								href={project.url}
								target="_blank"
								rel="noreferrer"
								aria-label={`${t('Ir a la página de|Go to the page of')} ${project.name}`}
							>
								<span class="icon-[lucide--external-link]" aria-hidden="true"></span>
							</a>
						{/if}

						<div class="project-logo">
							{#if project.logo}
								<img
									src={project.logo}
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

							<ul class="project-stack">
								{#each project.stack as item}
									<li><T text={item} /></li>
								{/each}
							</ul>
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
							<img src={doubt.icon} alt="" loading="lazy" />
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
				<!-- El arco acompaña la curva que trazan las cuatro tarjetas; apiladas en
				     una sola columna esa lectura se pierde y queda sólo la línea. -->
				<path
					class="horizon-arc"
					d="M0 72 C 400 8 1040 8 1440 72"
					fill="none"
					stroke="url(#horizonStroke)"
					stroke-width="2"
					vector-effect="non-scaling-stroke"
				/>
				<path
					class="horizon-flat"
					d="M0 72 H1440"
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
							style:--icon={`url("${nounResearch}")`}
							aria-hidden="true"
						>
							<span class="labs-point-glyph"></span>
						</span>
						<p>
							<T
								text="Inspirados en Bell Labs y Xerox PARC. Espacios de innovación donde nacieron los fundamentos tecnológicos que dieron forma a las décadas siguientes. Y en Google Research que produjo «Attention is All You Need», el origen de la IA moderna.|Inspired by Bell Labs and Xerox PARC: spaces of innovation where the technological foundations that shaped the following decades were born. And by Google Research, which produced «Attention is All You Need», the origin of modern AI."
							/>
						</p>
					</div>
					<div class="labs-point">
						<span
							class="labs-point-icon"
							style:--icon={`url("${nounInnovation}")`}
							aria-hidden="true"
						>
							<span class="labs-point-glyph"></span>
						</span>
						<p><T html text={labsProduct} /></p>
					</div>
					<div class="labs-point">
						<span
							class="labs-point-icon"
							style:--icon={`url("${nounMind}")`}
							aria-hidden="true"
						>
							<span class="labs-point-glyph"></span>
						</span>
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
					text="Genix Agentic UI es un ejemplo de cómo modelos pequeños y sin visión pueden servir como agentes que ayudan al usuario a navegar el sistema y ejecutar instrucciones. Ello reduce drásticamente los costos de inferencia.|Genix Agentic UI is an example of how small, vision-free models can act as agents that help the user navigate the system and carry out instructions. That is what makes it possible to squeeze every dollar of inference."
				/>
			</p>
		</div>

		<SectionCurve fill="#ffffff" variant="wave" side="bottom" flip />
	</section>

	<section class="genix">

		<div class="genix-intro">
			<img class="genix-logo" src={genixLogo} alt="Genix" />
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
							<span
								class="feature-icon"
								style:--icon={`url("${feature.icon}")`}
								aria-hidden="true"
							></span>
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

		<div class="genix-agent">
			<figure class="genix-agent-art">
				<img
					src={`${base}/svg/feature-ai.svg`}
					alt={t(
						'Un asistente virtual conversando con el usuario dentro de la aplicación|A virtual assistant chatting with the user inside the application'
					)}
					loading="lazy"
				/>
			</figure>

			<div class="genix-agent-copy">
				<h3>
					<T
						text="Un asistente virtual apoyándote en cada click.|A virtual assistant supporting you on every click."
					/>
				</h3>
				<p class="genix-agent-lead">
					<T
						text="Genix incorpora un agente que ve la misma pantalla que tú y actúa sobre ella: le pides algo en tus propias palabras y lo hace, así que aprender el sistema deja de ser un requisito para usarlo.|The agent sees the same screen you do and acts on it: you ask for something in your own words and it does it, so learning the system stops being a requirement for using it."
					/>
				</p>
			</div>

			<ul class="genix-agent-skills">
				{#each agentSkills as skill}
					<li>
						<strong><T text={skill.title} /></strong>
						<span><T text={skill.description} /></span>
					</li>
				{/each}
			</ul>
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
		background: #05061a;
		color: white;
	}

	/*
	 * Capas del hero, de abajo arriba: foto (0), velo y rejilla (1), contenido (2).
	 * Los z-index son explícitos porque la foto va después del ::before en el DOM
	 * y sin ellos taparía el velo que hace legible el texto.
	 */
	.hero-bg {
		position: absolute;
		z-index: 0;
		inset: 0;
	}

	.hero-bg img {
		width: 100%;
		height: 100%;
		object-fit: cover;
		object-position: center 32%;
	}

	.hero::before {
		position: absolute;
		z-index: 1;
		inset: 0;
		background:
			linear-gradient(100deg, rgba(5, 6, 26, 0.9) 0%, rgba(5, 6, 26, 0.56) 46%, rgba(5, 6, 26, 0.24) 100%),
			linear-gradient(0deg, rgba(5, 6, 26, 0.72), transparent 45%);
		content: '';
	}

	/* Sólo visible en móvil: en escritorio el selector sigue en la cabecera. */
	.hero-lang {
		display: none;
	}

	.hero-grid {
		position: absolute;
		z-index: 1;
		inset: 0;
		background-image:
			linear-gradient(rgba(150, 170, 255, 0.06) 1px, transparent 1px),
			linear-gradient(90deg, rgba(150, 170, 255, 0.06) 1px, transparent 1px);
		background-size: 7rem 7rem;
		mask-image: radial-gradient(circle at 20% 40%, black, transparent 70%);
	}

	.hero-inner {
		position: relative;
		z-index: 2;
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
		font-size: var(--fs-2xs);
		font-weight: 700;
		letter-spacing: 0.32em;
	}

	.hero-kicker span {
		width: 2.5rem;
		height: 1px;
		background: linear-gradient(90deg, transparent, var(--aqua));
	}

	/*
	 * Único titular fuera de --fs-h1..h6: su cuerpo no lo manda la jerarquía sino
	 * el corte del texto, que tiene que caer en las líneas marcadas con <br>. Lo
	 * mismo pasa con el override de móvil, más abajo.
	 */
	.hero h1 {
		margin: 0;
		font-family: var(--font-display);
		font-size: clamp(1.5rem, 2.9vw, 2.9rem);
		font-weight: 500;
		line-height: 1.15;
		letter-spacing: -0.025em;
	}

	/* Saltos de la versión estrecha: ocultos hasta que la media query los active. */
	.hero h1 :global(br.n) {
		display: none;
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
		font-size: var(--fs-lg);
		line-height: var(--lh-body);
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
		font-size: var(--fs-3xs);
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
		font-size: var(--fs-h3);
	}

	.intro-copy p:not(.eyebrow) {
		max-width: 44rem;
		margin: 0 0 1.2rem;
		color: var(--muted);
		font-size: var(--fs-md);
		line-height: var(--lh-relaxed);
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
		font-size: var(--fs-base);
		line-height: var(--lh-normal);
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
		font-size: var(--fs-2xs);
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

	.section-top-copy p:not(.eyebrow):not(.project-hint) {
		max-width: 38rem;
		margin: 1.4rem 0 0;
		color: var(--muted);
		line-height: var(--lh-relaxed);
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

	/* Cuelga del párrafo anterior, dentro del mismo bloque de texto. */
	.project-hint {
		margin: 0.85rem 0 0;
		color: var(--muted);
		font-size: var(--fs-sm);
	}

	/* El icono va dentro de la frase, así que se alinea con la línea base. */
	.project-hint span[class*='icon-'] {
		width: 1.15em;
		height: 1.15em;
		color: var(--accent);
		vertical-align: -0.28em;
	}

	.project-card {
		position: relative;
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

	/* Acceso directo a la web del proyecto, sobre la banda del logo. */
	.project-open {
		position: absolute;
		z-index: 1;
		top: 0.55rem;
		right: 0.55rem;
		display: inline-flex;
		width: 2.5rem;
		height: 2.5rem;
		align-items: center;
		justify-content: center;
		border-radius: 0.7rem;
		/* La esquina exterior acompaña la curva de la tarjeta. */
		border-top-right-radius: 1.125rem;
		background: rgba(255, 255, 255, 0.7);
		color: var(--accent);
		transition:
			background 180ms ease,
			color 180ms ease;
	}

	/* El hover de la tarjeta sólo marca el botón en gris... */
	.project-card:hover .project-open {
		background: var(--mist);
	}

	/*
	 * ...el azul se reserva para el hover del propio botón, que es lo que de verdad
	 * se pulsa. Necesita ganar en especificidad a la regla de arriba, porque al
	 * apuntar al botón también se está sobre la tarjeta.
	 */
	.project-card:hover .project-open:hover,
	.project-open:hover,
	.project-open:focus-visible {
		background: var(--accent);
		color: white;
	}

	.project-open span {
		width: 1.4rem;
		height: 1.4rem;
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
		font-size: var(--fs-h6);
		font-weight: 500;
		letter-spacing: 0.02em;
		text-align: center;
	}

	.project-body {
		display: flex;
		flex: 1;
		flex-direction: column;
		/* Arriba menos aire: el rótulo ya viene separado por el filete del logo. */
		padding: clamp(0.8rem, 1.3vw, 1.05rem) clamp(1rem, 2.2vw, 1.75rem)
			clamp(1rem, 2.2vw, 1.75rem);
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
		font-size: var(--fs-3xs);
		font-weight: 600;
		letter-spacing: 0.12em;
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
		font-size: var(--fs-3xs);
		font-weight: 700;
		letter-spacing: 0.08em;
		text-transform: uppercase;
		margin-right: -4px;
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
		margin: 0.5rem 0 0.7rem;
		font-family: var(--font-display);
		font-size: var(--fs-h5);
		font-weight: 500;
		line-height: var(--lh-tight);
	}

	.project-text {
		margin: 0 0 0.85rem;
		color: var(--muted);
		font-size: var(--fs-base);
		line-height: var(--lh-body);
	}

	/* `margin-top: auto` fija los tags al pie: así todas las tarjetas de una
	   fila alinean su última línea aunque el texto sea de distinto largo. */
	.project-stack {
		display: flex;
		min-width: 0;
		flex-wrap: wrap;
		gap: 0.4rem;
		margin: auto 0 0;
		padding: 0.5rem 0 0;
		list-style: none;
	}

	.project-stack li {
		padding: 0.28rem 0.6rem;
		border-radius: 0.45rem;
		background: var(--mist);
		color: var(--muted);
		font-size: var(--fs-3xs);
		letter-spacing: 0.04em;
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
		font-size: var(--fs-h2);
		font-weight: 500;
		line-height: var(--lh-display);
		letter-spacing: -0.04em;
	}

	.os-head > p:last-child {
		max-width: 46rem;
		margin: 0 auto;
		color: rgba(255, 255, 255, 0.6);
		line-height: var(--lh-relaxed);
		font-size: var(--fs-md);
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
		font-size: var(--fs-3xs);
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
		font-size: var(--fs-lg);
		line-height: var(--lh-snug);
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

	.horizon-flat {
		display: none;
	}

	.os-answer {
		max-width: 56rem;
		margin: clamp(1.5rem, 4vw, 3rem) auto 0;
		text-align: center;
	}

	.os-answer p {
		margin: 0;
		color: rgba(255, 255, 255, 0.66);
		font-size: var(--fs-md);
		line-height: var(--lh-relaxed);
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
		grid-template-columns: auto minmax(0, 1fr);
		gap: clamp(1.1rem, 2.2vw, 1.7rem);
		align-items: center;
	}

	/*
	 * Neumorfismo: la pastilla es del mismo color que la sección y sólo se separa
	 * del fondo por dos sombras opuestas —una luz arriba a la izquierda y una
	 * sombra abajo a la derecha—, así que el color de `--sand` y el de la sombra
	 * oscura tienen que moverse juntos si algún día cambia el fondo.
	 */
	.labs-point-icon {
		display: grid;
		place-items: center;
		width: clamp(4.1rem, 6.2vw, 5.4rem);
		aspect-ratio: 1;
		border-radius: 30%;
		background: var(--sand);
		box-shadow:
			6px 6px 14px rgba(105, 97, 168, 0.22),
			-6px -6px 14px #ffffff;
	}

	.labs-point-glyph {
		width: 70%;
		aspect-ratio: 1;
		background: var(--icon-ink);
		-webkit-mask: var(--icon) center / contain no-repeat;
		mask: var(--icon) center / contain no-repeat;
	}

	.labs-copy p {
		margin: 0;
		color: var(--muted);
		font-size: var(--fs-md);
		line-height: var(--lh-relaxed);
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
		font-size: var(--fs-xl);
		font-weight: 500;
		letter-spacing: -0.01em;
	}

	.legacy span {
		color: var(--muted);
		font-size: var(--fs-sm);
		line-height: var(--lh-body);
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

	/*
	 * La columna izquierda es más ancha que el logotipo, así que sin centrarlo el
	 * bloque quedaba pegado al margen. `align-self` mantiene la altura a la del
	 * titular; `justify-self` sólo lo mueve dentro de su propia columna.
	 */
	.genix-logo {
		grid-row: 1 / 3;
		grid-column: 1;
		width: min(15rem, 100%);
		align-self: center;
		justify-self: center;
	}

	.genix-cta {
		display: inline-flex;
		grid-row: 3;
		grid-column: 1;
		align-items: center;
		justify-self: center;
		gap: 0.5rem;
		margin-top: clamp(1.2rem, 2.5vw, 1.8rem);
		padding: 0.75rem 1.15rem;
		border: 1px solid var(--line);
		border-radius: 999px;
		background: var(--mist);
		color: var(--accent);
		font-size: var(--fs-2xs);
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
		font-size: var(--fs-sm);
		line-height: var(--lh-relaxed);
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
		font-size: var(--fs-h2);
		font-weight: 500;
		line-height: var(--lh-display);
		letter-spacing: -0.04em;
	}

	.genix-lead {
		grid-row: 3;
		grid-column: 3;
		max-width: 40rem;
		margin: clamp(1.2rem, 2.5vw, 1.8rem) 0 0;
		color: var(--muted);
		font-size: var(--fs-md);
		line-height: var(--lh-relaxed);
	}

	/*
	 * Sin panel de fondo: la sección es blanca y las tarjetas también, así que el
	 * relieve lo dan el filete y la sombra, no una tercera capa de color.
	 */
	.feature-grid {
		display: grid;
		grid-template-columns: repeat(4, minmax(0, 1fr));
		gap: clamp(1rem, 1.8vw, 1.4rem);
		margin-top: clamp(1.8rem, 3.5vw, 3rem);
	}

	/* El envoltorio de Reveal es el que cae en la rejilla: sin este `flex` la
	   tarjeta no se estira hasta el alto de su fila y los bordes quedan desiguales. */
	.feature-grid > :global(.reveal) {
		display: flex;
	}

	.feature-grid article {
		display: flex;
		flex: 1;
		flex-direction: column;
		position: relative;
		/* Recorta los realces contra el radio de la tarjeta. */
		overflow: hidden;
		padding: 1.6rem;
		border: 1px solid var(--line);
		border-radius: 1.1rem;
		background: #ffffff;
		box-shadow:
			0 1px 2px rgba(19, 22, 52, 0.04),
			0 12px 26px -16px rgba(91, 79, 214, 0.26);
		transition:
			transform 220ms ease,
			box-shadow 220ms ease,
			border-color 220ms ease;
	}

	/*
	 * Realce de esquina. Cada uno es un cuadrado anclado a su esquina, con un
	 * degradado radial centrado en el vértice INTERIOR y transparente hasta el 76 %
	 * del radio. Al quedar la esquina fuera de esa circunferencia, su borde es un
	 * arco que abomba hacia la esquina, y el realce se estrecha justo en la
	 * diagonal: eso es lo que lo mantiene pegado al vértice en vez de repartirse
	 * por toda la tarjeta, que es lo que pasaba al centrar el degradado en el
	 * elemento entero.
	 *
	 * El lado del cuadrado marca el radio del arco: cuanto menor, más cerrada es la
	 * curva. El 44 % es el punto en el que se lee como curva y no como diagonal.
	 *
	 * El cuadrado es la clave del truco: al medir lo mismo de alto que de ancho,
	 * sus otros dos vértices caen dentro del tramo transparente y el realce muere
	 * antes de llegar a los bordes de la caja, sin cortes secos.
	 */
	.feature-grid article::before,
	.feature-grid article::after {
		position: absolute;
		width: 44%;
		aspect-ratio: 1;
		content: '';
		pointer-events: none;
	}

	/* Luz violeta en la esquina superior derecha. */
	.feature-grid article::before {
		top: 0;
		right: 0;
		background: radial-gradient(
			circle farthest-corner at 0% 100%,
			rgba(98, 66, 173, 0) 0%,
			rgba(98, 66, 173, 0) 76%,
			rgba(98, 66, 173, 0.035) 90%,
			rgba(98, 66, 173, 0.16) 100%
		);
	}

	/* Sombra fría en la esquina inferior izquierda. */
	.feature-grid article::after {
		bottom: 0;
		left: 0;
		background: radial-gradient(
			circle farthest-corner at 100% 0%,
			rgba(29, 33, 68, 0) 0%,
			rgba(29, 33, 68, 0) 77%,
			rgba(29, 33, 68, 0.02) 90%,
			rgba(29, 33, 68, 0.1) 100%
		);
	}

	/* Los realces van absolutos y pintarían por encima del texto: este z-index los
	   devuelve al fondo sin necesidad de un contexto de apilamiento aparte. */
	.feature-grid article > * {
		position: relative;
		z-index: 1;
	}

	.feature-grid article:hover {
		border-color: rgba(91, 79, 214, 0.3);
		box-shadow:
			0 2px 4px rgba(19, 22, 52, 0.05),
			0 22px 40px -20px rgba(91, 79, 214, 0.4);
		transform: translateY(-4px);
	}

	/*
	 * Los iconos vienen dibujados en blanco, así que se usan como máscara y el
	 * color lo pone el fondo del propio elemento, en vez del negro que daba
	 * `invert()`. No llevan `--accent`: a este tamaño, ocho manchas de violeta
	 * saturado gritan más que el propio titular. Este índigo apagado conserva el
	 * matiz de la marca y deja el violeta pleno para enlaces y botones.
	 */
	.feature-icon {
		display: block;
		width: 3.8rem;
		height: 3.8rem;
		margin: 0.2rem 0 1.2rem;
		background: var(--icon-ink);
		mask: var(--icon) center / contain no-repeat;
		-webkit-mask: var(--icon) center / contain no-repeat;
		transition: filter 220ms ease;
	}

	/* Oscurecer con `filter` y no con otro `background`: así el degradado del icono
	   se mantiene en el hover en vez de aplanarse a un color liso. */
	.feature-grid article:hover .feature-icon {
		filter: brightness(0.82) saturate(1.1);
	}

	.feature-icon-slot {
		display: grid;
		width: 3.8rem;
		height: 3.8rem;
		margin: 0.2rem 0 1.2rem;
		border: 1px dashed rgba(61, 63, 102, 0.3);
		border-radius: 0.7rem;
		color: rgba(61, 63, 102, 0.5);
		font-size: var(--fs-3xs);
		letter-spacing: 0.08em;
		place-content: center;
		text-align: center;
		text-transform: uppercase;
	}

	/*
	 * Sin `margin-top: auto`: el título va pegado al icono en vez de caer al fondo
	 * de la tarjeta, que en las de texto corto abría un hueco enorme.
	 */
	.feature-grid h3 {
		margin: 0 0 0.6rem;
		font-family: var(--font-display);
		font-size: var(--fs-h6);
		font-weight: 500;
	}

	.feature-grid p {
		margin: 0;
		color: var(--muted);
		font-size: var(--fs-base);
		line-height: var(--lh-body);
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
		font-size: var(--fs-h2);
		font-weight: 500;
		line-height: var(--lh-display);
		letter-spacing: -0.04em;
	}

	.ai-copy p:not(.eyebrow) {
		max-width: 40rem;
		margin: 0 0 1.2rem;
		color: rgba(255, 255, 255, 0.6);
		font-size: var(--fs-md);
		line-height: var(--lh-relaxed);
	}

	/*
	 * Ficha del agente: cierra la sección de Genix, así que no es una sección más
	 * sino una tarjeta dentro de ella. El gris lila la despega del blanco de la
	 * sección sin competir con la mancha de la ilustración.
	 */
	.genix-agent {
		display: grid;
		grid-template-columns: minmax(0, 18rem) minmax(0, 1fr);
		gap: clamp(1.2rem, 3vw, 2.5rem);
		align-items: center;
		margin-top: clamp(2.5rem, 5vw, 4rem);
		padding: clamp(1.6rem, 3vw, 2.4rem);
		border-radius: 1.2rem;
		background: #f6f5fa;
	}

	.genix-agent-art {
		grid-row: 1 / 3;
		align-self: center;
		margin: 0;
	}

	.genix-agent-art img {
		width: 100%;
		height: auto;
	}

	.genix-agent-copy h3 {
		max-width: 22ch;
		margin: 1rem 0 1.4rem 0;
		font-family: var(--font-display);
		font-size: var(--fs-h4);
		font-weight: 500;
		line-height: var(--lh-tight);
		letter-spacing: -0.03em;
	}

	.genix-agent-lead {
		max-width: 46rem;
		margin: 0;
		color: var(--muted);
		font-size: var(--fs-md);
		line-height: var(--lh-relaxed);
	}

	.genix-agent-skills {
		display: grid;
		grid-template-columns: repeat(3, minmax(0, 1fr));
		gap: 1rem clamp(1.2rem, 2.5vw, 2rem);
		/* Sin margen propio: el hueco con el párrafo ya lo pone el `gap` de la ficha. */
		margin: 0;
		padding: 0;
		list-style: none;
	}

	.genix-agent-skills li {
		padding-top: 0.7rem;
		border-top: 1px solid #dcd9f2;
	}

	.genix-agent-skills strong {
		display: block;
		margin-bottom: 0.15rem;
		color: var(--accent);
		font-size: var(--fs-base);
		font-weight: 700;
	}

	.genix-agent-skills span {
		color: var(--muted);
		font-size: var(--fs-sm);
		line-height: var(--lh-body);
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
			font-size: var(--fs-3xs);
			letter-spacing: 0.17em;
		}

		.hero-kicker span {
			width: 1.5rem;
		}
	}

	@media (max-width: 860px) {
		.hero {
			padding-top: calc(var(--header-height) + 2.5rem);
		}

		/* El encuadre que antes daba `background-position` ahora es del <img>. */
		.hero-bg img {
			object-position: center 45%;
		}

		/* Apilado el texto ocupa menos alto, así que el velo puede aclararse. */
		.hero::before {
			background: linear-gradient(0deg, rgba(5, 6, 26, 0.65) 24%, rgba(5, 6, 26, 0.2));
		}

		/*
		 * Con el velo más claro, el texto cae sobre el limbo iluminado de la Tierra,
		 * que además es la zona con más detalle. La sombra hace dos cosas: la corta
		 * define el filo de la letra y la ancha abre un halo oscuro que la despega
		 * del fondo. No se aplica a los botones: el principal lleva texto oscuro
		 * sobre aguamarina y ahí la sombra sólo ensucia.
		 */
		.hero h1 {
			text-shadow: 0 2px 12px rgba(3, 4, 20, 0.55);
		}

		.hero-lead,
		.hero-foot {
			text-shadow:
				0 1px 3px rgba(3, 4, 20, 0.75),
				0 0 14px rgba(3, 4, 20, 0.6);
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

		.labs-body,
		.genix-intro {
			grid-template-columns: 1fr;
		}

		/* Apilada, la ilustración se queda arriba y a un tamaño de icono grande. */
		.genix-agent {
			grid-template-columns: 1fr;
		}

		.genix-agent-art {
			grid-row: auto;
			width: min(18rem, 72%);
		}

		.genix-agent-skills {
			grid-template-columns: repeat(2, minmax(0, 1fr));
		}

		/*
		 * Apilada, la ilustración sube entre el primer y el segundo párrafo. Para
		 * intercalarla hay que disolver las cajas intermedias (el envoltorio de
		 * Reveal y `.intro-copy`) y dejar que rótulo, titular, párrafos, figura y
		 * llamada a la acción sean hermanos directos del mismo contenedor flex.
		 */
		.intro {
			display: flex;
			flex-direction: column;
			align-items: stretch;
			gap: 0;
		}

		.intro > :global(.reveal),
		.intro-copy {
			display: contents;
		}

		/*
		 * El resto del texto pasa a un grupo posterior; todo lo demás se queda en el
		 * grupo por defecto conservando el orden del marcado, así que la figura cae
		 * justo detrás del primer párrafo.
		 */
		.intro-more,
		.intro-cta {
			order: 1;
		}

		/*
		 * Arriba basta con poco: el párrafo anterior ya aporta su propio margen
		 * inferior (en flex los márgenes no colapsan). Abajo, en cambio, el texto
		 * arranca sin margen superior y necesita todo el aire aquí.
		 */
		.intro-art {
			max-width: 20rem;
			margin: 0.25rem auto clamp(1.75rem, 6vw, 2.5rem);
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

		/*
		 * Apilado, la ilustración va entre la entradilla y la línea de ayuda, que se
		 * queda pegada a las tarjetas que explica. Como en «¿Quiénes somos?», hay que
		 * disolver `.section-top-copy` para poder intercalarla con `order`.
		 */
		.section-top-copy {
			display: contents;
		}

		.section-top-copy > :not(.project-hint) {
			order: -1;
		}

		.project-hint {
			order: 1;
			margin-top: clamp(2rem, 5vw, 4.5rem);
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
		 * las dos columnas; apilado, monta el rótulo sobre el botón. Y como aquí el
		 * rótulo y el titular presentan la sección, se colocan delante del logo: en
		 * una sola columna basta con `order`, la rejilla respeta ese orden.
		 */
		.genix-intro .eyebrow {
			order: -2;
			/* Sin el `transform` que lo montaba sobre el titular, necesita su propio aire. */
			margin-bottom: 0.6rem;
			transform: none;
		}

		.genix-intro h2 {
			order: -1;
		}

		/* Ya no abre el bloque: necesita separarse del titular que ahora va encima. */
		.genix-logo {
			margin-top: calc(clamp(1.2rem, 4vw, 2rem) + 0.5rem);
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
			display: flex;
			flex-direction: column;
			align-items: stretch;
			/* El espaciado lo llevan los márgenes de cada bloque, no el `gap`. */
			gap: 0;
			margin-bottom: clamp(1rem, 2.2vw, 1.5rem);
		}

		.section-top-art {
			margin: clamp(2rem, 5vw, 4.5rem) auto 0;
		}
	}

	/*
	 * Móvil: no hay cabecera (sólo el botón del menú flotando en la esquina), así
	 * que el hero recupera la parte alta de la pantalla y el selector de idioma se
	 * ancla a su esquina superior izquierda. Con el logo apilado encima, el texto
	 * queda centrado.
	 */
	@media (max-width: 720px) {
		/*
		 * El borde inferior es sólo el respiro del pie: la posición vertical del
		 * bloque la fija el `margin-top` de `.hero-inner`, no el reparto del
		 * centrado. Cuanto se le añada aquí se suma al alto del hero y lo saca de
		 * la pantalla.
		 */
		.hero {
			padding-top: 3.75rem;
			padding-bottom: 2rem;
		}

		.hero-bg img {
			object-position: center;
		}

		/*
		 * Con el logo ocupando media pantalla, el balanceo del `float` se lee como
		 * un temblor: en pantalla ancha es un detalle, aquí distrae.
		 */
		.hero-mark img {
			animation: none;
		}

		.hero-lang {
			position: absolute;
			z-index: 3;
			top: 1rem;
			left: var(--page-pad);
			display: block;
		}

		.hero-copy {
			text-align: center;
		}

		.hero-lead {
			margin-top: 2.5rem;
			margin-inline: auto;
		}

		/* Más aire entre el logo y el titular que el que reparte el layout ancho. */
		.hero-inner {
			gap: 20vw;
			margin-top: 14vw;
		}

		/*
		 * El botón fantasma cae justo sobre el limbo iluminado de la Tierra, donde
		 * el borde de 1px se pierde. Un velo negro le devuelve el contorno sin
		 * convertirlo en un botón sólido que compita con el principal.
		 */
		.hero-actions .button-secondary {
			background: rgba(0, 0, 0, 0.35);
		}

		.hero-actions,
		.hero-foot {
			justify-content: center;
		}

		.intro {
			padding-top: 2.75rem;
			text-align: center;
		}

		.ai {
			text-align: center;
		}

		/* Rótulo y titular presentan la sección de Genix: centrados como el resto. */
		.genix-intro .eyebrow,
		.genix-intro h2 {
			text-align: center;
		}

		/*
		 * El botón se recuesta en el costado derecho y sube con margen negativo para
		 * aprovechar el aire que deja el logo por debajo.
		 */
		.genix-cta {
			justify-self: end;
			margin-top: -1.5rem;
		}

		.ai-copy p:not(.eyebrow) {
			margin-inline: auto;
		}

		/*
		 * Los puntos de «labs» se reorganizan: la pastilla del icono encima, alineada
		 * al margen, y el texto a todo el ancho debajo.
		 */
		.labs-point {
			grid-template-columns: minmax(0, 1fr);
			gap: 0.95rem;
			align-items: center;
		}

		/* En su propia fila el icono ya no compite con el texto: puede crecer. */
		.labs-point-icon {
			width: 5.6rem;
		}

		/* Dos líneas largas: en móvil el titular de «labs» pide un cuerpo menor. */
		.labs-head .section-heading {
			font-size: clamp(1.95rem, 5vw, 5rem);
		}

		/* Rótulo, titular y entradilla de «Proyectos», también centrados. */
		.section-top-copy {
			text-align: center;
		}

		.section-top-copy p:not(.eyebrow) {
			margin-inline: auto;
		}

		.project-hint {
			text-align: center;
		}

		/*
		 * El bloque de llamada a la acción conserva su texto alineado a la izquierda
		 * y el botón cae a la derecha, casi pegado a la pregunta: apilados no
		 * necesitan el aire de la versión ancha.
		 */
		.intro-cta {
			justify-content: flex-end;
			gap: 0 1.2rem;
			padding: 1.15rem 1.2rem;
			text-align: left;
		}

		/* El alto mínimo del botón ya deja aire de sobra sobre su rótulo. */
		.intro-cta a {
			margin-top: -0.45rem;
		}
	}

	@media (max-width: 640px) {
		.project-grid,
		.doubts,
		.feature-grid,
		.genix-agent-skills {
			grid-template-columns: 1fr;
		}

		/*
		 * En una columna la ficha ya no necesita despegarse de nada: el recuadro
		 * gris solo roba ancho al texto. Se disuelve y su papel de cierre lo asume
		 * un filete con la rampa de la marca, difuminado en los extremos para que
		 * no choque con los bordes de la página.
		 */
		.genix-agent {
			padding: 0;
			border-radius: 0;
			background: none;
		}

		/* Como hijo directo de la rejilla ocupa su propia fila; el `gap` lo separa. */
		.genix-agent::before {
			display: block;
			height: 2px;
			border-radius: 999px;
			background: var(--brand-ramp);
			-webkit-mask: linear-gradient(90deg, transparent, #000 16%, #000 84%, transparent);
			mask: linear-gradient(90deg, transparent, #000 16%, #000 84%, transparent);
			content: '';
		}

		/*
		 * En una columna el titular abre el bloque y la ilustración lo cierra, así
		 * que la figura baja con `order`. El filete es el primer hijo de la rejilla
		 * y también entra en el reparto: se queda arriba con un `order` negativo.
		 */
		.genix-agent {
			text-align: center;
		}

		.genix-agent::before {
			order: -1;
		}

		.genix-agent-art {
			order: 1;
			margin-inline: auto;
		}

		/* Las capacidades van después de la ilustración y siguen leyéndose en fila. */
		.genix-agent-skills {
			order: 2;
			text-align: left;
		}

		.genix-agent-copy h3,
		.genix-agent-lead {
			margin-inline: auto;
		}


		/*
		 * Una tarjeta por fila: el icono se va al costado derecho y el texto ocupa la
		 * columna izquierda, en lugar de apilarse bajo un icono suelto a la izquierda.
		 */
		.feature-grid article {
			display: grid;
			min-height: 0;
			grid-template-columns: minmax(0, 1fr) auto;
			gap: 0 1.2rem;
			align-content: start;
			padding: 1.3rem 1.4rem;
		}

		.feature-icon,
		.feature-icon-slot {
			grid-row: 1 / span 2;
			grid-column: 2;
			align-self: start;
			margin: 0;
		}

		.feature-grid h3 {
			grid-row: 1;
			grid-column: 1;
		}

		.feature-grid p {
			grid-row: 2;
			grid-column: 1;
		}

		/* La sección ya aporta `--page-pad`: este sangrado extra sobraba. */
		.os-answer {
			margin-top: 2.25rem;
		}

		/* El arco es mucho más bajo aquí: necesita su propio aire a ambos lados. */
		.horizon {
			margin-top: 2.25rem;
		}

		.horizon-arc {
			display: none;
		}

		.horizon-flat {
			display: block;
		}

		/*
		 * Una recta no necesita caja alta ni los recortes que pedía la curva, y a todo
		 * el ancho pesa demasiado: se queda en un filete centrado.
		 */
		.horizon svg {
			--horizon-h: 0.75rem;

			width: 78%;
			margin: 0 auto;
		}

		/*
		 * A una tarjeta por fila el número no necesita su propia línea: se saca del
		 * flujo y se ancla a la izquierda del icono, que así sube y recupera el aire
		 * que le quitaba encima.
		 */
		.doubts blockquote {
			position: relative;
		}

		.doubt-num {
			position: absolute;
			top: 0.35rem;
			left: 15%;
		}

		.doubt-icon {
			margin: 0 0 1.5rem;
		}

		/*
		 * El cuerpo lo manda el corte pedido: «Iniciativas de código abierto» tiene
		 * que caber de una pieza en la primera línea. Manda el 7vw; el mínimo está
		 * en 1.6rem y no en 1.85 porque a partir de ahí el cuerpo dejaría de
		 * encoger con la pantalla mientras el ancho disponible sigue bajando, y por
		 * debajo de 400px la primera línea se partía en dos.
		 */
		.hero h1 {
			font-size: clamp(1.6rem, 7vw, 2.4rem);
			text-wrap: balance;
		}

		/* Se intercambian los dos juegos de saltos. */
		.hero h1 :global(br.w) {
			display: none;
		}

		.hero h1 :global(br.n) {
			display: inline;
		}

		/*
		 * En estrecho el último tramo de estos dos párrafos pasa a leerse aparte:
		 * el <span> se vuelve bloque, lo que ya fuerza el salto, y el margen abre
		 * el aire. En ancho no lleva estilo ninguno y sigue siendo texto corrido,
		 * así que no hace falta partir el string bilingüe en dos entradas.
		 */
		.os-head :global(.os-split),
		.os-answer :global(.os-split) {
			display: block;
			margin-top: 1em;
		}

		.hero-foot {
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
