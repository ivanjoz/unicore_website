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

export const articles: Article[] = [
	{
		title: 'Inteligencia Artificial: Oportunidades y Riesgos',
		summary:
			'Un entendimiento del machine learning para sumarnos al debate sobre sus posibilidades y su futuro.',
		image: '/blog/S_1566444701.jpeg',
		date: '2019-08-21'
	},
	{
		title: 'La dicotomía de los modelos económicos',
		summary:
			'Cómo se articula un modelo económico a través de la generación de valor y qué podemos aprender de ello.',
		image: '/blog/S_1566444850.jpeg',
		date: '2019-08-21'
	},
	{
		title: 'Retos de la evaluación de inversiones en el Perú',
		summary:
			'Criterios para mitigar el riesgo y tomar mejores decisiones financieras en el mercado peruano.',
		image: '/blog/S_1566444836.jpeg',
		date: '2019-08-21'
	},
	{
		title: 'La informalidad empresarial, una cuestión de datos',
		summary:
			'La gestión de información como base para construir criterios y profesionalizar las decisiones.',
		image: '/blog/S_1566444823.jpeg',
		date: '2019-08-21'
	},
	{
		title: 'Concepto de «mente abierta» desde la ciencia',
		summary:
			'Una mirada al pensamiento crítico y a la interpretación no sesgada de la realidad.',
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
			'Sistema de gestión empresarial, CRM y E-commerce para pequeñas empresas. Punto de venta, logistica, gestión de cuentas, clientes y flujos de caja proyectados y facturación electrónica. Permite a cada empresa exportar su data. Arquitectura híbrida: cloud-native o self-host. Hecho en Go + Svelte.js.',
		url: 'https://github.com/ivanjoz/genix',
		status: 'En desarrollo',
		stack: ['Go', 'ScyllaDB', 'AI']
	},
	{
		name: 'Colbin',
		kind: 'Serialización',
		logo: '/svg/colbin_icon.svg',
		description:
			'Serializador binario columnar optimizado para arrays de registros. Hecho en Go y AssemblyScript. Transpone los datos y los codifica por columna usando varint, delta encoding y 5-bit chars, reduciendo el ancho de banda frente a formatos como JSON, CBOR o Protobuf.',
		url: 'https://github.com/ivanjoz/colbin',
		status: 'Activo',
		stack: ['Go', 'Rust']
	},
	{
		name: 'Genix Search',
		kind: 'Motor de búsqueda',
		logo: '/svg/genix_search.svg',
		description:
			'Backend de búsqueda compacto y rankeado, optimizado para textos cortos y multi-índice con bajo consumo de RAM. Usa bigramas computados en español, claves enteras directas y RocksDB como storage. Binario estático sin dependencias.',
		url: 'https://github.com/ivanjoz/genix-search',
		status: 'Activo',
		stack: ['Rust', 'RocksDB', 'TCP']
	},
	{
		name: 'Genix Agentic UI',
		kind: 'Componentes de IA',
		logo: '/svg/genix_ui.svg',
		description:
			'Componentes de interfaz agénticos: modelos pequeños y sin visión que acompañan al usuario a navegar el sistema y a ejecutar instrucciones, exprimiendo cada dólar de inferencia.',
		url: 'https://github.com/ivanjoz/genix-ui',
		status: 'En desarrollo',
		stack: ['Svelte', 'LLM']
	},
	{
		name: 'Factura-Go',
		kind: 'Facturación electrónica',
		logo: '/svg/factura_go.svg',
		description:
			'Librería en Go para la emisión de comprobantes electrónicos: construcción del documento, firma digital, envío y consulta del estado ante la SUNAT (Perú).',
		url: 'https://github.com/ivanjoz/facturago',
		status: 'En desarrollo',
		stack: ['Go', 'XML/UBL']
	},
	{
		name: 'Simple Vault',
		kind: 'Seguridad',
		logo: '/svg/simple_vault_logo.png',
		description:
			'Gestor de contraseñas offline-first que cifra todo en el navegador con Argon2id y AES-256-GCM, y guarda la bóveda en el almacenamiento aislado de tu propio Google Drive. Sin servidor de aplicación.',
		url: 'https://github.com/ivanjoz/simple-vault',
		status: 'Activo',
		stack: ['SvelteKit', 'PWA', 'Criptografía']
	},
	{
		name: 'Genix ORM',
		kind: 'Bases de datos',
		logo: '/svg/genix_orm.svg',
		description:
			'ORM con API de consultas verificada en tiempo de compilación para ScyllaDB / Cassandra y DynamoDB. Los patrones de acceso se declaran por adelantado y se sirven por partición, índice o vista.',
		url: 'https://github.com/ivanjoz/genix-orm',
		status: 'Activo',
		stack: ['Go', 'ScyllaDB', 'DynamoDB']
	},
	{
		name: 'Auth-Limiter',
		kind: 'Infraestructura',
		logo: '/svg/relimiter_logo.svg',
		description:
			'Servicio en memoria que valida accesos y descuenta créditos de uso en una sola llamada TCP, en base a quotas por minuto. Proceso en Rust con tramas HMAC, locks entre Lambdas concurrentes, request logs, reportería y puente SSE.',
		url: 'https://github.com/ivanjoz/auth-limiter',
		status: 'En desarrollo',
		stack: ['Rust', 'TCP', 'SSE', 'ScyllaDB']
	},
	{
		name: 'Genix Support AI',
		kind: 'Agente autónomo',
		logo: '/svg/genix_support_agent.svg',
		description:
			'Automatización de soporte para proyectos de código abierto: indexa el código, lee los tickets, reproduce el problema como un usuario real con BrowserOS, abre el PR con la corrección, espera el despliegue y responde el ticket.',
		status: 'En diseño',
		stack: ['Agentes', 'BrowserOS', 'CI/CD']
	},
	{
		name: 'Classi-Cont',
		kind: 'Modelo de IA',
		description:
			'Clasificación contable asistida por IA: categoría contable, cuenta PCGE sugerida y taxonomía de producto, acompañada de un motor determinista y versionado para las reglas tributarias.',
		status: 'En diseño',
		stack: ['ML', 'PCGE', 'Modelo ad-hoc']
	},
];

export const missionTerms = [
	{
		name: 'Desarrollar',
		description:
			'El desarrollo es un proceso metódico y constructivo que lleva a concretar o mejorar una idea. En Unicore estamos constantemente desarrollando y desarrollándonos.'
	},
	{
		name: 'soluciones tecnológicas',
		description:
			'Toda solución apunta a resolver un problema e incentivar un cambio. Nuestros proyectos nacen del diagnóstico de una necesidad y buscan agregar valor.'
	},
	{
		name: 'basadas en la web',
		description:
			'La web permite que el software funcione en cualquier dispositivo y que la infraestructura sea más accesible, distribuida y potente.'
	},
	{
		name: 'escalables',
		description:
			'Un sistema escalable está preparado para incorporar nuevas funcionalidades, expandirse y actualizarse con facilidad.'
	},
	{
		name: 'accesibles',
		description:
			'Creemos que la tecnología debe ser inclusiva y estar al alcance de las empresas y usuarios que quieren transformar su actividad.'
	},
	{
		name: 'con soporte y asesoría',
		description:
			'La tecnología es una herramienta. La acompañamos con diagnóstico, aprendizaje y planificación para obtener resultados reales.'
	},
	{
		name: 'para mejorar decisiones',
		description:
			'Ordenar procesos e información permite medir mejor la actividad y tomar decisiones sustentadas en evidencias.'
	},
	{
		name: 'en las empresas peruanas',
		description:
			'Nos sumamos al reto de modernizar la gestión del empresariado peruano y llevar sus capacidades a nuevos horizontes.'
	}
];
