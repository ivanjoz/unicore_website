#!/usr/bin/env bash
# Genera la tarjeta 1200x630 que ven WhatsApp, Facebook, LinkedIn y X al
# compartir el enlace, una por idioma.
#
# Parte de design-assets/og_card_base.svg, que ya trae compuestos el fondo y el
# lockup del logo. Ese SVG no vive en static/ porque pesa 2,4 MB —lleva el PNG
# de la Tierra incrustado— y se publicaría entero sin que nadie lo pida: aquí
# es un máster de diseño, igual que los PNG del hero.
set -euo pipefail

here="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
src="$(cd "$here/../../design-assets" && pwd)"
out="$(cd "$here/../static/images" && pwd)"
tmp="$(mktemp -d)"
trap 'rm -rf "$tmp"' EXIT

# El claim del máster viene convertido a trazos (dos paths con los ids `text4`
# y `text4-9`), así que no se puede traducir editando el SVG. Se quitan y se
# repone por idioma más abajo, con la Inter del sitio.
python3 - "$src/og_card_base.svg" "$tmp/base.svg" <<'PY'
import re, sys
svg = open(sys.argv[1], encoding='utf-8').read()
stripped = re.sub(r'<path\b[^>]*?\bid="(?:text4|text4-9)"[^>]*?/?>', '', svg)
removed = len(re.findall(r'<path', svg)) - len(re.findall(r'<path', stripped))
if removed != 2:
	sys.exit(f'esperaba quitar los 2 paths del claim, quité {removed}: ¿cambió el máster?')
open(sys.argv[2], 'w', encoding='utf-8').write(stripped)
PY

# El renderizador SVG interno de ImageMagick ignora parte del documento, así
# que el rasterizado lo hace inkscape.
inkscape --export-type=png --export-width=1200 --export-background-opacity=0 \
	-o "$tmp/base.png" "$tmp/base.svg" >/dev/null 2>&1

render() {
	local name="$1" claim="$2"

	# El máster es 1.79:1 y la tarjeta 1.91:1: sobran 41 px de alto. Se recortan
	# casi todos por arriba, que es donde hay aire, para no comerse la Tierra.
	magick "$tmp/base.png" -crop 1200x630+0+28 +repage \
		\( -background none -fill white -font Inter-SemiBold -pointsize 47 \
			-interline-spacing 12 -gravity center label:"$claim" \
			\( +clone -background black -shadow 100x12+0+4 \) +swap \
			-background none -layers merge +repage \) \
		-gravity center -geometry +0+172 -composite \
		-strip -quality 88 "$out/$name.jpg"

	# JPEG y no AVIF/WebP: los crawlers de WhatsApp y Facebook siguen sin
	# soportarlos de forma fiable, y aquí una imagen que no carga deja la
	# previsualización sin imagen. Son los únicos JPEG que quedan en el sitio.
	printf '%-18s %s B\n' "$name.jpg" "$(stat -c%s "$out/$name.jpg")"
}

render og-card-es $'Innovación y desarrollo\nde sistemas Open Source'
render og-card-en $'Innovation and development\nof Open Source systems'
