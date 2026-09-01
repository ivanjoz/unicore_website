import { browser } from '$app/environment';

export type Lang = 'es' | 'en';

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
		// El atributo del documento lo pone el servidor al prerenderizar; aquí sólo
		// hay que mantenerlo al día cuando se navega entre `/` y `/en/` sin recarga.
		if (browser) document.documentElement.lang = next;
	}
};

export function t(text: string): string {
	if (typeof text !== 'string' || !text.includes('|')) return text;
	const variants = text.split('|');
	return (variants[current === 'es' ? 0 : 1] ?? variants[0]).trim();
}

/** Etiqueta BCP-47 para `Intl` y para el atributo `lang` del documento. */
export function localeTag(): string {
	return current === 'es' ? 'es-PE' : 'en-US';
}
