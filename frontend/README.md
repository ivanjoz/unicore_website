# Unicore frontend

Refactor del sitio de Unicore en SvelteKit, Svelte 5 con runes y TypeScript.

## Desarrollo

```sh
bun install
bun run dev
```

## Validación

```sh
bun run check
bun run build
```

## GitHub Pages

El workflow [deploy-pages.yml](../.github/workflows/deploy-pages.yml) publica el sitio
automáticamente cuando hay cambios en `frontend/` sobre la rama `main`. También puede
ejecutarse manualmente desde la pestaña **Actions**.

El build obtiene el subdirectorio de GitHub Pages automáticamente. Para reproducir
localmente un build de proyecto:

```sh
BASE_PATH=/nombre-del-repositorio bun run build
```

Rutas incluidas:

- `/` — inicio, servicios, GERP, publicaciones y contacto
- `/nosotros` — misión, visión, objetivos, cosmovisión y marca
- `/portafolio` — proyectos seleccionados
