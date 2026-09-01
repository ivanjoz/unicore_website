#!/usr/bin/env bash
# Genera la tarjeta 1200x630 que ven WhatsApp, Facebook, LinkedIn y X al
# compartir el enlace, una por idioma.
#
# Se compone en vez de reusar el logo suelto: en la previsualización de
# WhatsApp el texto del mensaje queda diminuto, así que la imagen es lo único
# que se lee de un vistazo y tiene que decir de quién es el sitio y qué hace.
set -euo pipefail

here="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
src="$(cd "$here/../../design-assets" && pwd)"
svg="$(cd "$here/../static/svg" && pwd)"
out="$(cd "$here/../static/images" && pwd)"
tmp="$(mktemp -d)"
trap 'rm -rf "$tmp"' EXIT

# El logo se rasteriza con inkscape y no con ImageMagick: el renderizador SVG
# interno de IM ignora parte del documento y devuelve el logo recortado sobre
# una caja blanca.
inkscape --export-type=png --export-height=150 --export-background-opacity=0 \
	-o "$tmp/logo.png" "$svg/logo_unicore_horizontal_light.svg" >/dev/null 2>&1

# Velo con los mismos degradados que .hero::before: oscurece la izquierda, donde
# va el texto, y deja limpia la derecha, que es donde se ve la Tierra.
magick -size 630x1200 gradient:'rgba(5,6,26,0.94)-rgba(5,6,26,0.10)' \
	-rotate 90 "$tmp/veil.png"

render() {
	local name="$1" title="$2" claim="$3"

	# El máster es 16:9 y la tarjeta 1.91:1, así que sobra alto: se recorta desde
	# +40 para no cortar el limbo iluminado, que es lo que da profundidad.
	magick "$src/hero_space_earth_original.png" \
		-gravity north -crop 1672x878+0+40 +repage -resize 1200x630! \
		"$tmp/veil.png" -compose over -composite \
		"$tmp/logo.png" -gravity northwest -geometry +70+62 -compose over -composite \
		-font Inter-SemiBold -fill white -pointsize 64 -interline-spacing 16 \
		-gravity northwest -annotate +72+250 "$title" \
		-font Inter-Regular -fill '#a8bacf' -pointsize 30 -interline-spacing 10 \
		-gravity northwest -annotate +75+468 "$claim" \
		-strip -quality 88 "$out/$name.jpg"

	# JPEG y no AVIF/WebP: los crawlers de WhatsApp y Facebook siguen sin
	# soportarlos de forma fiable, y aquí una imagen que no carga deja la
	# previsualización sin imagen. Son los únicos JPEG que quedan en el sitio.
	printf '%-18s %s B\n' "$name.jpg" "$(stat -c%s "$out/$name.jpg")"
}

render og-card-es \
	$'Código abierto\nde alto impacto' \
	$'Sistemas y herramientas que expanden el acceso\na tecnologías de alto impacto.'

render og-card-en \
	$'High-impact\nopen source' \
	$'Systems and tools that expand access\nto high-impact technology.'
