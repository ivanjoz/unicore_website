import { browser } from '$app/environment';

export type Lang = 'es' | 'en';

const STORAGE_KEY = 'unicore:lang';

/*
 * Los textos bilingües viajan dentro de un único string separado por "|":
 *
 *   'Contáctanos|Contact us'  →  [0] español · [1] inglés
 *
 * Un string sin "|" se devuelve tal cual, así que los nombres propios, las
 * siglas y el código (Go, RocksDB, ERP…) no necesitan traducción.
 */
let current = $state<Lang>('es');

export const lang = {
	get value(): Lang {
		return current;
	},
	set value(next: Lang) {
		current = next;
		if (!browser) return;
		document.documentElement.lang = next;
		try {
			localStorage.setItem(STORAGE_KEY, next);
		} catch {
			// Navegación privada o almacenamiento bloqueado: el idioma dura la sesión.
		}
	}
};

export function t(text: string): string {
	if (typeof text !== 'string' || !text.includes('|')) return text;
	const variants = text.split('|');
	return (variants[current === 'es' ? 0 : 1] ?? variants[0]).trim();
}

export function toggleLang() {
	lang.value = current === 'es' ? 'en' : 'es';
}

/** Etiqueta BCP-47 para `Intl` y para el atributo `lang` del documento. */
export function localeTag(): string {
	return current === 'es' ? 'es-PE' : 'en-US';
}

/**
 * La página se prerenderiza en español, así que la preferencia real solo puede
 * aplicarse en el cliente: primero lo elegido antes, y si no hay nada guardado,
 * el idioma del navegador.
 */
export function restoreLang() {
	if (!browser) return;

	let stored: string | null = null;
	try {
		stored = localStorage.getItem(STORAGE_KEY);
	} catch {
		stored = null;
	}

	if (stored === 'es' || stored === 'en') {
		lang.value = stored;
		return;
	}

	lang.value = navigator.language?.toLowerCase().startsWith('es') ? 'es' : 'en';
}
