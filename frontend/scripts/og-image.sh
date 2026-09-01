#!/usr/bin/env bash
# Genera la tarjeta 1200x630 que ven WhatsApp, Facebook, LinkedIn y X al
# compartir el enlace, una por idioma.
#
# Parte de design-assets/og_card_base.svg, que ya trae compuestos el fondo, el
# lockup del logo y el claim. Ese SVG no vive en static/ porque pesa 2,4 MB
# —lleva el PNG de la Tierra incrustado— y se publicaría entero al sitio sin
# que ninguna página lo pidiera: aquí es un máster de diseño, como los PNG del
# hero.
set -euo pipefail

here="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
src="$(cd "$here/../../design-assets" && pwd)"
out="$(cd "$here/../static/images" && pwd)"
tmp="$(mktemp -d)"
trap 'rm -rf "$tmp"' EXIT

# El español sale del máster sin tocar nada, así que editar el claim en Inkscape
# actualiza la tarjeta ES sin pasar por aquí. Sólo el inglés necesita sustituir
# el contenido de los dos <tspan> de #text13. Van alineados a la derecha
# (`text-anchor:end`) y comparten la misma x, así que las líneas en inglés, de
# otro largo, siguen cuadrando solas.
translate() {
	python3 - "$src/og_card_base.svg" "$1" "$2" "$3" <<'PY'
import re, sys

svg_path, out_path, line1, line2 = sys.argv[1:5]
svg = open(svg_path, encoding='utf-8').read()

block = re.search(r'<text\b[^>]*\bid="text13".*?</text>', svg, re.S)
if not block:
	sys.exit('no encuentro #text13 en el máster: ¿cambió el SVG?')

spans = re.findall(r'(<tspan\b[^>]*>)(.*?)(</tspan>)', block.group(0), re.S)
if len(spans) != 2:
	sys.exit(f'esperaba 2 líneas de claim, encontré {len(spans)}: ¿cambió el máster?')

new = block.group(0)
for (open_tag, _, close_tag), text in zip(spans, (line1, line2)):
	new = new.replace(f'{open_tag}{_}{close_tag}', f'{open_tag}{text}{close_tag}', 1)

open(out_path, 'w', encoding='utf-8').write(svg.replace(block.group(0), new, 1))
PY
}

render() {
	local name="$1" line1="${2:-}" line2="${3:-}"
	local svg="$src/og_card_base.svg"

	if [[ -n "$line1" ]]; then
		svg="$tmp/$name.svg"
		translate "$svg" "$line1" "$line2"
	fi

	# El rasterizado lo hace inkscape: el renderizador SVG interno de ImageMagick
	# ignora parte del documento y devuelve la tarjeta incompleta.
	inkscape --export-type=png --export-width=1200 --export-background-opacity=0 \
		-o "$tmp/$name.png" "$svg" >/dev/null 2>&1

	# El máster es 1.79:1 y la tarjeta 1.91:1: sobran 41 px de alto. Se reparten
	# entre el aire de arriba (sobre el logo) y el de abajo (bajo el claim), que
	# es donde hay margen de sobra en la composición.
	magick "$tmp/$name.png" -crop 1200x630+0+14 +repage \
		-strip -quality 88 "$out/$name.jpg"

	# JPEG y no AVIF/WebP: los crawlers de WhatsApp y Facebook siguen sin
	# soportarlos de forma fiable, y aquí una imagen que no carga deja la
	# previsualización sin imagen. Son los únicos JPEG que quedan en el sitio.
	printf '%-18s %s B\n' "$name.jpg" "$(stat -c%s "$out/$name.jpg")"
}

# Sin argumentos: el claim del máster, en español.
render og-card-es
render og-card-en 'Innovation and development' 'of open source systems'
