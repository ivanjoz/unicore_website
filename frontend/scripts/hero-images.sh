#!/usr/bin/env bash
# Genera los derivados del hero desde los PNG maestros de design-assets/.
#
# Sólo AVIF y WebP: el <img> de respaldo apunta al WebP, soportado por todo
# navegador que entienda <picture>, así que un JPEG extra no aporta cobertura.
#
# Las calidades están calibradas para dejar el AVIF en 150-180 kB. Bajar de ahí
# ensucia el degradado del espacio con banding, que es la zona sobre la que cae
# el texto del hero.
set -euo pipefail

src="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../design-assets" && pwd)"
out="$(cd "$(dirname "${BASH_SOURCE[0]}")/../static/images" && pwd)"

# origen                     destino               avif  webp
encode() {
	local input="$1" name="$2" q_avif="$3" q_webp="$4"
	avifenc -q "$q_avif" -s 4 --jobs all "$src/$input" "$out/$name.avif" >/dev/null
	magick "$src/$input" -strip -quality "$q_webp" -define webp:method=6 "$out/$name.webp"
	# Las dimensiones se leen del máster: ImageMagick aquí no decodifica AVIF.
	printf '%-26s %-9s  avif %6s B  webp %6s B\n' "$name" \
		"$(magick identify -format '%wx%h' "$src/$input")" \
		"$(stat -c%s "$out/$name.avif")" "$(stat -c%s "$out/$name.webp")"
}

# Se codifica al tamaño nativo del máster. El JPEG anterior iba a 2400 px de
# ancho, pero era un reescalado del mismo 1672 px: píxeles de más sin detalle
# de más, y el presupuesto de bytes se iba en el reescalado en vez de en la
# compresión.
encode hero_space_earth_original.png space_earth        84 88
encode hero_scape_vertical.png       space_earth_mobile 82 88
