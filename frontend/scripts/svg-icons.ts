/*
 * Empaqueta los SVG pequeños de src/lib/svg en un único módulo (icons.gen.ts).
 * Cada SVG se optimiza con svgo y se deja pre-codificado como data-URI, así el
 * navegador no hace una petición por icono y en runtime no queda nada por hacer.
 * Los SVG grandes (ilustraciones) siguen en /static y se cargan por URL.
 */
import { readdirSync, readFileSync, writeFileSync, existsSync } from 'node:fs';
import { join, relative, dirname, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { optimize } from 'svgo';

const root = dirname(dirname(fileURLToPath(import.meta.url)));
export const SVG_DIR = join(root, 'src/lib/svg');
const OUT_FILE = join(SVG_DIR, 'icons.gen.ts');

/** Sobre este tamaño (bytes ya optimizados) inlinear deja de compensar. */
const MAX_INLINE_BYTES = 16 * 1024;

const walk = (dir: string): string[] =>
	readdirSync(dir, { withFileTypes: true }).flatMap((e) => {
		const full = join(dir, e.name);
		if (e.isDirectory()) return walk(full);
		return e.name.endsWith('.svg') ? [full] : [];
	});

/** producto1.svg -> producto1 | icons/icon-continuity.svg -> iconContinuity */
const toIdent = (file: string) =>
	file
		.replace(/\.svg$/, '')
		.split(/[/\\]/)
		.pop()!
		.replace(/[^a-zA-Z0-9]+(.)?/g, (_, c: string | undefined) => (c ? c.toUpperCase() : ''))
		.replace(/^[0-9]/, (d) => `_${d}`);

/*
 * Codificación mínima: sólo los caracteres que romperían el atributo o el url()
 * de CSS. Sale bastante más corto que encodeURIComponent y vale en ambos sitios.
 */
const encodeSvg = (svg: string) =>
	svg
		.trim()
		.replace(/[\r\n\t]+/g, ' ')
		.replace(/[%#<>?[\]^`{|}\\"]/g, (c) => `%${c.charCodeAt(0).toString(16).toUpperCase()}`);

export function generateSvgIcons(): string {
	if (!existsSync(SVG_DIR)) throw new Error(`No existe ${SVG_DIR}`);

	const files = walk(SVG_DIR).sort();
	const seen = new Map<string, string>();
	const entries: { ident: string; rel: string; uri: string; bytes: number }[] = [];

	for (const file of files) {
		const rel = relative(SVG_DIR, file).replaceAll('\\', '/');
		const ident = toIdent(rel);

		const prev = seen.get(ident);
		if (prev) throw new Error(`Colisión de nombre "${ident}": ${prev} y ${rel}`);
		seen.set(ident, rel);

		const { data } = optimize(readFileSync(file, 'utf8'), { path: file, multipass: true });
		const bytes = Buffer.byteLength(data);
		if (bytes > MAX_INLINE_BYTES) {
			throw new Error(
				`${rel} pesa ${bytes} bytes optimizado (máx ${MAX_INLINE_BYTES}). ` +
					`Muévelo a static/svg y cárgalo por URL.`
			);
		}

		entries.push({ ident, rel, uri: `data:image/svg+xml,${encodeSvg(data)}`, bytes });
	}

	const total = entries.reduce((n, e) => n + e.bytes, 0);
	const body = entries
		.map((e) => `/** ${e.rel} — ${e.bytes} B */\nexport const ${e.ident}: string = ${JSON.stringify(e.uri)};`)
		.join('\n\n');

	const out = `// GENERADO por scripts/svg-icons.ts — no editar a mano.
// Fuente: src/lib/svg/**/*.svg (${entries.length} iconos, ${total} B optimizados).
// Regenerar: npm run svg  (o solo con arrancar vite).

${body}
`;

	const current = existsSync(OUT_FILE) ? readFileSync(OUT_FILE, 'utf8') : '';
	if (current !== out) writeFileSync(OUT_FILE, out);
	return `${entries.length} iconos, ${(total / 1024).toFixed(1)} kB optimizados`;
}

/** Plugin de Vite: genera al arrancar y al tocar cualquier .svg en dev. */
export function svgIcons() {
	return {
		name: 'unicore:svg-icons',
		buildStart() {
			generateSvgIcons();
		},
		configureServer(server: {
			watcher: { add(p: string): void; on(ev: string, cb: (f: string) => void): void };
		}) {
			// chokidar 4 ya no acepta globs, así que se vigila el directorio entero.
			server.watcher.add(SVG_DIR);
			const onChange = (file: string) => {
				if (!file.endsWith('.svg') || !resolve(file).startsWith(SVG_DIR)) return;
				try {
					generateSvgIcons();
				} catch (err) {
					console.error(`[svg-icons] ${(err as Error).message}`);
				}
			};
			for (const ev of ['add', 'change', 'unlink']) server.watcher.on(ev, onChange);
		}
	};
}

if (process.argv[1] && fileURLToPath(import.meta.url) === process.argv[1]) {
	console.log(generateSvgIcons());
}
