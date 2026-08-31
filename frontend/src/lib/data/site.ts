/*
 * Los textos de este archivo son bilingües: 'español|english'. Se muestran con el
 * componente <T> o con la función t() de $lib/i18n.svelte. Lo que se escribe igual
 * en ambos idiomas (nombres, siglas, tecnologías) va sin separador.
 */

export type Article = {
	title: string;
	summary: string;
	image: string;
	date: string;
};

export type LabProject = {
	name: string;
	kind: string;
	description: string;
	url?: string;
	/** Ruta a un logo en /static. Si no existe, la tarjeta muestra el nombre. */
	logo?: string;
	status: 'Activo' | 'En desarrollo' | 'En diseño';
	stack: string[];
};

/** `status` se usa como clave para el estilo de la tarjeta; aquí vive su etiqueta visible. */
export const statusLabels: Record<LabProject['status'], string> = {
	Activo: 'Activo|Active',
	'En desarrollo': 'En desarrollo|In development',
	'En diseño': 'En diseño|In design'
};

export const articles: Article[] = [
	{
		title:
			'Inteligencia Artificial: Oportunidades y Riesgos|Artificial Intelligence: Opportunities and Risks',
		summary:
			'Un entendimiento del machine learning para sumarnos al debate sobre sus posibilidades y su futuro.|An understanding of machine learning so we can join the debate about its possibilities and its future.',
		image: '/blog/S_1566444701.jpeg',
		date: '2019-08-21'
	},
	{
		title: 'La dicotomía de los modelos económicos|The dichotomy of economic models',
		summary:
			'Cómo se articula un modelo económico a través de la generación de valor y qué podemos aprender de ello.|How an economic model is built around the creation of value, and what we can learn from it.',
		image: '/blog/S_1566444850.jpeg',
		date: '2019-08-21'
	},
	{
		title:
			'Retos de la evaluación de inversiones en el Perú|Challenges of investment appraisal in Peru',
		summary:
			'Criterios para mitigar el riesgo y tomar mejores decisiones financieras en el mercado peruano.|Criteria to mitigate risk and make better financial decisions in the Peruvian market.',
		image: '/blog/S_1566444836.jpeg',
		date: '2019-08-21'
	},
	{
		title: 'La informalidad empresarial, una cuestión de datos|Business informality, a matter of data',
		summary:
			'La gestión de información como base para construir criterios y profesionalizar las decisiones.|Information management as the basis for building criteria and professionalizing decisions.',
		image: '/blog/S_1566444823.jpeg',
		date: '2019-08-21'
	},
	{
		title:
			'Concepto de «mente abierta» desde la ciencia|The idea of the «open mind» seen from science',
		summary:
			'Una mirada al pensamiento crítico y a la interpretación no sesgada de la realidad.|A look at critical thinking and at the unbiased interpretation of reality.',
		image: '/blog/S_1566444812.jpeg',
		date: '2019-08-21'
	}
];

export const labProjects: LabProject[] = [
	{
		name: 'Genix ERP',
		kind: 'ERP + E-commerce',
		logo: '/svg/genix_logo.svg',
		description:
			'Sistema de gestión empresarial, CRM y E-commerce para pequeñas empresas. Punto de venta, logistica, gestión de cuentas, clientes y flujos de caja proyectados y facturación electrónica. Permite a cada empresa exportar su data. Arquitectura híbrida: cloud-native o self-host. Hecho en Go + Svelte.js.|Business management, CRM and e-commerce system for small companies. Point of sale, logistics, account management, customers, projected cash flows and electronic invoicing. Every company can export its own data. Hybrid architecture: cloud-native or self-hosted. Built with Go + Svelte.js.',
		url: 'https://github.com/ivanjoz/genix',
		status: 'En desarrollo',
		stack: ['Go', 'ScyllaDB', 'AI']
	},
	{
		name: 'Colbin',
		kind: 'Serialización|Serialization',
		logo: '/svg/colbin_icon.svg',
		description:
			'Serializador binario columnar optimizado para arrays de registros. Hecho en Go y AssemblyScript. Transpone los datos y los codifica por columna usando varint, delta encoding y 5-bit chars, reduciendo el ancho de banda frente a formatos como JSON, CBOR o Protobuf.|Columnar binary serializer optimized for arrays of records. Written in Go and AssemblyScript. It transposes the data and encodes it column by column using varint, delta encoding and 5-bit chars, cutting bandwidth compared with formats such as JSON, CBOR or Protobuf.',
		url: 'https://github.com/ivanjoz/colbin',
		status: 'Activo',
		stack: ['Go', 'Rust']
	},
	{
		name: 'Genix Search',
		kind: 'Motor de búsqueda|Search engine',
		logo: '/svg/genix_search.svg',
		description:
			'Backend de búsqueda compacto y rankeado, optimizado para textos cortos y multi-índice con bajo consumo de RAM. Usa bigramas computados en español, claves enteras directas y RocksDB como storage. Binario estático sin dependencias.|Compact ranked search backend, optimized for short texts and multiple indexes with low RAM usage. It uses bigrams computed for Spanish, direct integer keys and RocksDB as storage. A static binary with no dependencies.',
		url: 'https://github.com/ivanjoz/genix-search',
		status: 'Activo',
		stack: ['Rust', 'RocksDB', 'TCP']
	},
	{
		name: 'Genix Agentic UI',
		kind: 'Componentes de IA|AI components',
		logo: '/svg/genix_ui.svg',
		description:
			'Componentes de interfaz agénticos: modelos pequeños y sin visión que acompañan al usuario a navegar el sistema y a ejecutar instrucciones, exprimiendo cada dólar de inferencia.|Agentic interface components: small, vision-free models that guide the user through the system and carry out instructions, squeezing every dollar of inference.',
		url: 'https://github.com/ivanjoz/genix-ui',
		status: 'En desarrollo',
		stack: ['Svelte', 'LLM']
	},
	{
		name: 'Factura-Go',
		kind: 'Facturación electrónica|Electronic invoicing',
		logo: '/svg/factura_go.svg',
		description:
			'Librería en Go para la emisión de comprobantes electrónicos: construcción del documento, firma digital, envío y consulta del estado ante la SUNAT (Perú).|Go library for issuing electronic tax documents: building the document, signing it digitally, sending it and checking its status with SUNAT (Peru).',
		url: 'https://github.com/ivanjoz/facturago',
		status: 'En desarrollo',
		stack: ['Go', 'XML/UBL']
	},
	{
		name: 'Simple Vault',
		kind: 'Seguridad|Security',
		logo: '/svg/simple_vault_logo.png',
		description:
			'Gestor de contraseñas offline-first que cifra todo en el navegador con Argon2id y AES-256-GCM, y guarda la bóveda en el almacenamiento aislado de tu propio Google Drive. Sin servidor de aplicación.|Offline-first password manager that encrypts everything in the browser with Argon2id and AES-256-GCM, and keeps the vault in the isolated storage of your own Google Drive. No application server.',
		url: 'https://github.com/ivanjoz/simple-vault',
		status: 'Activo',
		stack: ['SvelteKit', 'PWA', 'Criptografía|Cryptography']
	},
	{
		name: 'Genix ORM',
		kind: 'Bases de datos|Databases',
		logo: '/svg/genix_orm.svg',
		description:
			'ORM con API de consultas verificada en tiempo de compilación para ScyllaDB / Cassandra y DynamoDB. Los patrones de acceso se declaran por adelantado y se sirven por partición, índice o vista.|ORM with a compile-time checked query API for ScyllaDB / Cassandra and DynamoDB. Access patterns are declared up front and served by partition, index or view.',
		url: 'https://github.com/ivanjoz/genix-orm',
		status: 'Activo',
		stack: ['Go', 'ScyllaDB', 'DynamoDB']
	},
	{
		name: 'Auth-Limiter',
		kind: 'Infraestructura|Infrastructure',
		logo: '/svg/relimiter_logo.svg',
		description:
			'Servicio en memoria que valida accesos y descuenta créditos de uso en una sola llamada TCP, en base a quotas por minuto. Proceso en Rust con tramas HMAC, locks entre Lambdas concurrentes, request logs, reportería y puente SSE.|In-memory service that validates access and deducts usage credits in a single TCP call, based on per-minute quotas. A Rust process with HMAC frames, locks across concurrent Lambdas, request logs, reporting and an SSE bridge.',
		url: 'https://github.com/ivanjoz/auth-limiter',
		status: 'En desarrollo',
		stack: ['Rust', 'TCP', 'SSE', 'ScyllaDB']
	},
	{
		name: 'Genix Support AI',
		kind: 'Agente autónomo|Autonomous agent',
		logo: '/svg/genix_support_agent.svg',
		description:
			'Automatización de soporte para proyectos de código abierto: indexa el código, lee los tickets, reproduce el problema como un usuario real con BrowserOS, abre el PR con la corrección, espera el despliegue y responde el ticket.|Support automation for open-source projects: it indexes the code, reads the tickets, reproduces the problem like a real user with BrowserOS, opens the PR with the fix, waits for the deployment and answers the ticket.',
		status: 'En diseño',
		stack: ['Agentes|Agents', 'BrowserOS', 'CI/CD']
	},
	{
		name: 'Classi-Cont',
		kind: 'Modelo de IA|AI model',
		description:
			'Clasificación contable asistida por IA: categoría contable, cuenta PCGE sugerida y taxonomía de producto, acompañada de un motor determinista y versionado para las reglas tributarias.|AI-assisted accounting classification: accounting category, suggested PCGE account and product taxonomy, backed by a deterministic, versioned engine for tax rules.',
		status: 'En diseño',
		stack: ['ML', 'PCGE', 'Modelo ad-hoc|Ad-hoc model']
	}
];

export const missionTerms = [
	{
		name: 'Desarrollar|To develop',
		description:
			'El desarrollo es un proceso metódico y constructivo que lleva a concretar o mejorar una idea. En Unicore estamos constantemente desarrollando y desarrollándonos.|Development is a methodical, constructive process that turns an idea into something real or makes it better. At Unicore we are constantly developing, and developing ourselves.'
	},
	{
		name: 'soluciones tecnológicas|technology solutions',
		description:
			'Toda solución apunta a resolver un problema e incentivar un cambio. Nuestros proyectos nacen del diagnóstico de una necesidad y buscan agregar valor.|Every solution aims to solve a problem and encourage change. Our projects start from diagnosing a need and set out to add value.'
	},
	{
		name: 'basadas en la web|built on the web',
		description:
			'La web permite que el software funcione en cualquier dispositivo y que la infraestructura sea más accesible, distribuida y potente.|The web lets software run on any device and makes infrastructure more accessible, distributed and powerful.'
	},
	{
		name: 'escalables|scalable',
		description:
			'Un sistema escalable está preparado para incorporar nuevas funcionalidades, expandirse y actualizarse con facilidad.|A scalable system is ready to take on new features, grow and be updated with ease.'
	},
	{
		name: 'accesibles|accessible',
		description:
			'Creemos que la tecnología debe ser inclusiva y estar al alcance de las empresas y usuarios que quieren transformar su actividad.|We believe technology must be inclusive and within reach of the companies and users who want to transform what they do.'
	},
	{
		name: 'con soporte y asesoría|with support and advice',
		description:
			'La tecnología es una herramienta. La acompañamos con diagnóstico, aprendizaje y planificación para obtener resultados reales.|Technology is a tool. We pair it with diagnosis, learning and planning to get real results.'
	},
	{
		name: 'para mejorar decisiones|to improve decisions',
		description:
			'Ordenar procesos e información permite medir mejor la actividad y tomar decisiones sustentadas en evidencias.|Putting processes and information in order makes activity easier to measure and decisions evidence-based.'
	},
	{
		name: 'en las empresas peruanas|in Peruvian companies',
		description:
			'Nos sumamos al reto de modernizar la gestión del empresariado peruano y llevar sus capacidades a nuevos horizontes.|We are taking on the challenge of modernizing management in Peruvian business and taking its capabilities to new horizons.'
	}
];
