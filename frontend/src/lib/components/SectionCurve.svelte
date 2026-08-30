<script lang="ts">
	/*
	 * Curved divider between two sections: one path, two colours.
	 *
	 * It paints the *previous* section's colour over the top of the section that
	 * hosts it, rather than painting this section's colour above it. That way the
	 * host paints its own background — gradients included — across the full strip,
	 * so a glow near the top edge stays continuous instead of breaking into a hard
	 * line where a flat-filled overlay would have met it.
	 *
	 * Drawn as an SVG overlay rather than a `clip-path` on the section itself:
	 * `clip-path: path()` takes absolute pixel coordinates, so the curve would not
	 * follow the viewport width, and clipping the section would also cut its
	 * content. This never touches the text.
	 *
	 * `preserveAspectRatio="none"` lets the path stretch across any width, so the
	 * height is controlled from CSS independently of how wide the screen is.
	 *
	 * Hosts must reserve `var(--curve-h)` of extra top padding.
	 */
	type Variant = 'wave' | 'dome' | 'valley';

	let {
		/** Background of the section above — the colour that spills into this one. */
		fill,
		/**
		 * wave   — S-shaped, climbing from left to right.
		 * dome   — arc bulging upward, so this section rises into the one above.
		 * valley — arc bulging downward, so the one above dips into this section.
		 */
		variant = 'wave',
		/** Mirror horizontally. */
		flip = false
	}: { fill: string; variant?: Variant; flip?: boolean } = $props();

	// Each path traces the boundary, then closes a few units above the viewBox so
	// the fill bleeds over the seam. Keep that bleed small: it is painted in the
	// previous section's colour, so a large one covers that section's content.
	const paths: Record<Variant, string> = {
		wave: 'M0 100 C 90 88 200 58 430 50 C 620 45 850 53 1150 46 C 1290 42 1370 22 1440 0 L1440 -4 L0 -4 Z',
		dome: 'M0 50 C 480 -22 1020 -18 1440 92 L1440 -4 L0 -4 Z',
		valley: 'M0 50 C 480 122 1020 118 1440 8 L1440 -4 L0 -4 Z'
	};
</script>

<div class="curve" class:flip style:--curve-fill={fill} aria-hidden="true">
	<svg viewBox="0 0 1440 100" preserveAspectRatio="none" focusable="false">
		<path d={paths[variant]} />
	</svg>
</div>

<style>
	.curve {
		position: absolute;
		top: 0;
		left: 0;
		width: 100%;
		height: var(--curve-h);
		line-height: 0;
		pointer-events: none;
	}

	.flip {
		scale: -1 1;
	}

	/*
	 * The path runs past the top of the viewBox and the svg does not clip, so the
	 * fill bleeds into the section above. Ending it exactly on the viewBox edge
	 * leaves an antialiased hairline along the seam.
	 */
	svg {
		display: block;
		width: 100%;
		height: 100%;
		overflow: visible;
	}

	path {
		fill: var(--curve-fill);
	}
</style>
