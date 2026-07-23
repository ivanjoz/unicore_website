export type Service = {
	title: string;
	description: string;
	icon: string;
	tag: string;
};

export type Feature = {
	title: string;
	description: string;
	icon: string;
};

export type Article = {
	title: string;
	summary: string;
	image: string;
	date: string;
};

export type Project = {
	type: string;
	image: string;
	title: string;
	description: string;
	url: string;
};

export const services: Service[] = [
	{
		tag: '01 / PRESENCIA DIGITAL',
		title: 'Web que trabaja por tu negocio',
		description:
			'Creamos sitios rápidos, adaptables y administrables, con dominio, correo corporativo, SSL y posicionamiento en buscadores.',
		icon: '/svg/web2.svg'
	},
	{
		tag: '02 / COMERCIO ELECTRÓNICO',
		title: 'Tu tienda abierta las 24 horas',
		description:
			'Catálogo, carrito, stock, variantes y pagos en línea en una experiencia clara para tus clientes y sencilla para tu equipo.',
		icon: '/svg/ecommerce2.svg'
	},
	{
		tag: '03 / SOFTWARE',
		title: 'Sistemas hechos para tus procesos',
		description:
			'Desarrollamos aplicaciones web a medida, portales y herramientas empresariales que conectan equipos, datos y decisiones.',
		icon: '/svg/chip1.svg'
	}
];

export const gerpFeatures: Feature[] = [
	{
		title: 'Productos y servicios',
		description: 'Catálogo, fabricación, insumos y suministros.',
		icon: '/svg/producto1.svg'
	},
	{
		title: 'Almacenes',
		description: 'Inventarios, traslados, consignaciones y guías.',
		icon: '/svg/almacen1.svg'
	},
	{
		title: 'Ventas',
		description: 'Clientes, facturación electrónica y venta rápida.',
		icon: '/svg/ventas1.svg'
	},
	{
		title: 'Equipos',
		description: 'Usuarios, roles, permisos y planillas.',
		icon: '/svg/people2.svg'
	},
	{
		title: 'Finanzas',
		description: 'Cajas, cuentas por cobrar y cuentas por pagar.',
		icon: '/svg/finanzas2.svg'
	},
	{
		title: 'Contabilidad',
		description: 'Libros, activos, resultados y balance general.',
		icon: '/svg/finanzas3.svg'
	},
	{
		title: 'Indicadores',
		description: 'Métricas de gestión, liquidez y flujo de caja.',
		icon: '/svg/idea.svg'
	}
];

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

export const projects: Project[] = [
	{
		type: 'Página web y sistema',
		image: '/webs/web_elpaisa.jpg',
		title: 'Restaurante Turístico El Paisa',
		description:
			'Sitio web optimizado con sistema de gestión de reservas y pagos con tarjeta.',
		url: 'https://elpaisa.pe/'
	},
	{
		type: 'Página web y correos',
		image: '/webs/amg.jpg',
		title: 'AMG Constructora',
		description:
			'Sitio web optimizado, panel de administración y correos corporativos.',
		url: 'https://amgconstructora.un.pe/'
	},
	{
		type: 'Sistema web empresarial',
		image: '/webs/jobfinder1.jpg',
		title: 'Jobfinder',
		description:
			'Buscador e indexador de empleos con herramientas de gestión de usuarios.',
		url: 'https://jobfinder.pe/'
	},
	{
		type: 'Página web',
		image: '/webs/esquimedsac.jpg',
		title: 'Esquimed SAC',
		description: 'Sitio web corporativo adaptable con panel de administración.',
		url: 'https://esquimedsac.un.pe/'
	},
	{
		type: 'Página web',
		image: '/webs/matizpe.jpg',
		title: 'Matiz Publicidad',
		description: 'Sitio web corporativo adaptable con panel de administración.',
		url: 'https://matiz.pe/'
	},
	{
		type: 'Página web',
		image: '/webs/recolocate1.jpg',
		title: 'Recolocate Job & Career',
		description: 'Sitio web de servicios profesionales con contenido administrable.',
		url: 'http://recolocate.pe/'
	},
	{
		type: 'Página web + redes',
		image: '/webs/mundosalud.jpg',
		title: 'Mundo Salud ONG',
		description:
			'Sitio web con panel de administración y despliegue de presencia en redes.',
		url: 'https://mundosalud.un.pe/'
	},
	{
		type: 'Página web y correos',
		image: '/webs/unicore1.jpg',
		title: 'Unicore Perú',
		description:
			'Sitio web administrable, optimización de rendimiento y correos corporativos.',
		url: 'https://un.pe/'
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
