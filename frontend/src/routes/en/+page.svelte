<script lang="ts">
	// La portada vive en $lib/components/Home.svelte porque la comparten las dos
	// rutas de idioma: `/` (español) y `/en/` (inglés). El idioma no se pasa por
	// props, lo fija el layout a partir de la URL antes de renderizar esto.
	import Home from '$lib/components/Home.svelte';
</script>

<Home />
