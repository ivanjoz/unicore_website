<script lang="ts">
	import { onMount, type Snippet } from 'svelte';

	let {
		children,
		delay = 0,
		class: className = ''
	}: { children: Snippet; delay?: number; class?: string } = $props();

	let element = $state<HTMLElement>();
	let visible = $state(true);
	let enhanced = $state(false);

	onMount(() => {
		if (!element || !('IntersectionObserver' in window)) {
			visible = true;
			return;
		}

		enhanced = true;
		visible = element.getBoundingClientRect().top < window.innerHeight * 0.92;

		const observer = new IntersectionObserver(
			([entry]) => {
				if (entry.isIntersecting) {
					visible = true;
					observer.disconnect();
				}
			},
			{ rootMargin: '0px 0px -8% 0px' }
		);

		observer.observe(element);
		return () => observer.disconnect();
	});
</script>

<div
	bind:this={element}
	class="reveal {className}"
	class:visible
	class:enhanced
	style:--reveal-delay={`${delay}ms`}
>
	{@render children()}
</div>

<style>
	.reveal.enhanced {
		opacity: 0;
		transform: translateY(2rem);
		transition:
			opacity 700ms cubic-bezier(0.22, 1, 0.36, 1) var(--reveal-delay),
			transform 700ms cubic-bezier(0.22, 1, 0.36, 1) var(--reveal-delay);
	}

	.reveal.enhanced.visible {
		opacity: 1;
		transform: translateY(0);
	}

	@media (prefers-reduced-motion: reduce) {
		.reveal.enhanced {
			opacity: 1;
			transform: none;
			transition: none;
		}
	}
</style>
