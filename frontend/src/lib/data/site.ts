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
		description:
			'Plataforma autoalojable y multiempresa para pequeños negocios: punto de venta, inventarios, caja, clientes y flujos de caja proyectados. Cada empresa puede exportar una copia completa de sus datos cuando quiera.',
		url: 'https://github.com/ivanjoz/genix',
		status: 'En desarrollo',
		stack: ['Go', 'ScyllaDB', 'SvelteKit']
	},
	{
		name: 'Colbin',
		kind: 'Serialización',
		description:
			'Serializador binario columnar para arreglos de structs. Transpone los datos y los codifica columna por columna, reduciendo drásticamente el ancho de banda frente a formatos por fila como JSON o CBOR.',
		url: 'https://github.com/ivanjoz/colbin',
		status: 'Activo',
		stack: ['Go', 'Rust']
	},
	{
		name: 'Genix ORM',
		kind: 'Acceso a datos',
		description:
			'ORM con API de consultas verificada en tiempo de compilación para ScyllaDB / Cassandra y DynamoDB. Los patrones de acceso se declaran por adelantado y se sirven por partición, índice o vista.',
		url: 'https://github.com/ivanjoz/genix-orm',
		status: 'Activo',
		stack: ['Go', 'ScyllaDB', 'DynamoDB']
	},
	{
		name: 'Genix Agentic UI',
		kind: 'Componentes de IA',
		description:
			'Componentes de interfaz agénticos: modelos pequeños y sin visión que acompañan al usuario a navegar el sistema y a ejecutar instrucciones, exprimiendo cada dólar de inferencia.',
		url: 'https://github.com/ivanjoz/genix-ui',
		status: 'En desarrollo',
		stack: ['Svelte', 'LLM']
	},
	{
		name: 'Simple Vault',
		kind: 'Seguridad',
		description:
			'Gestor de contraseñas offline-first que cifra todo en el navegador con Argon2id y AES-256-GCM, y guarda la bóveda en el almacenamiento aislado de tu propio Google Drive. Sin servidor de aplicación.',
		url: 'https://github.com/ivanjoz/simple-vault',
		status: 'Activo',
		stack: ['SvelteKit', 'PWA', 'Criptografía']
	},
	{
		name: 'Factura-Go',
		kind: 'Facturación electrónica',
		description:
			'Librería en Go para la emisión de comprobantes electrónicos: construcción del documento, firma digital, envío y consulta del estado ante la administración tributaria.',
		url: 'https://github.com/ivanjoz/facturago',
		status: 'En desarrollo',
		stack: ['Go', 'XML/UBL']
	},
	{
		name: 'Genix Search',
		kind: 'Motor de búsqueda',
		description:
			'Backend de búsqueda compacto y rankeado, optimizado para nombres de productos y texto comercial en español. Índice por pares de letras, claves enteras directas y binarios estáticos sin dependencias.',
		url: 'https://github.com/ivanjoz/genix-search',
		status: 'Activo',
		stack: ['Rust', 'RocksDB', 'TCP']
	},
	{
		name: 'Classi-Cont',
		kind: 'Modelo de IA',
		description:
			'Clasificación contable asistida por IA: categoría contable, cuenta PCGE sugerida y taxonomía de producto, acompañada de un motor determinista y versionado para las reglas tributarias.',
		status: 'En diseño',
		stack: ['ML', 'PCGE', 'Modelo ad-hoc']
	},
	{
		name: 'Re-Limiter',
		kind: 'Infraestructura',
		description:
			'Servicio en Rust que recibe mensajes autenticados de tamaño fijo sobre TCP y aplica límites de CPU y de créditos de inferencia por empresa, por usuario y por ventana de tiempo.',
		url: 'https://github.com/ivanjoz/genix/tree/main/server_utils',
		status: 'En desarrollo',
		stack: ['Rust', 'TCP', 'ScyllaDB']
	},
	{
		name: 'Genix Support AI',
		kind: 'Agente autónomo',
		description:
			'Automatización de soporte para proyectos de código abierto: indexa el código, lee los tickets, reproduce el problema como un usuario real con BrowserOS, abre el PR con la corrección, espera el despliegue y responde el ticket.',
		status: 'En diseño',
		stack: ['Agentes', 'BrowserOS', 'CI/CD']
	}
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

export const worldview = [
	'Creemos en la tolerancia, el respeto y los derechos individuales.',
	'Creemos que las empresas son motores indispensables de generación de valor.',
	'Creemos que todas las profesiones son trascendentes y el arte, una expresión sublime.',
	'Creemos que la tecnología, el conocimiento y la información deben ser accesibles.',
	'Creemos en una población culta, educada e informada como base del desarrollo.',
	'Creemos en el método científico y en decisiones sustentadas en evidencias.',
	'Creemos en resolver los conflictos de nuestra época con conciencia y consenso.',
	'Creemos en la urgencia de detener el cambio climático y la degradación del planeta.',
	'Creemos que el poder de la tecnología debe manejarse con responsabilidad.',
	'Creemos en la humanidad y en su capacidad de corregirse y superarse.'
];
