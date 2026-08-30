<script lang="ts">
	import { missionTerms } from '$lib/data/site';

	let selected = $state(0);
	let activeTerm = $derived(missionTerms[selected]);
</script>

<div class="mission-explorer">
	<div class="terms" aria-label="Explora nuestra misión">
		{#each missionTerms as term, index}
			<button
				type="button"
				class:active={selected === index}
				onmouseenter={() => (selected = index)}
				onfocus={() => (selected = index)}
				onclick={() => (selected = index)}
			>
				{term.name}
			</button>
		{/each}
	</div>

	<div class="explanation" aria-live="polite">
		<span>{String(selected + 1).padStart(2, '0')}</span>
		<div>
			<p class="eyebrow">Lo que significa para nosotros</p>
			<h3>{activeTerm.name}</h3>
			<p>{activeTerm.description}</p>
		</div>
	</div>
</div>

<style>
	.mission-explorer {
		display: grid;
		grid-template-columns: minmax(0, 1.1fr) minmax(18rem, 0.9fr);
		gap: clamp(2rem, 7vw, 7rem);
		align-items: center;
	}

	.terms {
		display: flex;
		flex-wrap: wrap;
		gap: 0.65rem;
	}

	.terms button {
		padding: 0.75rem 1rem;
		border: 1px solid var(--line);
		border-radius: 999px;
		background: transparent;
		color: var(--ink);
		font: inherit;
		font-size: clamp(0.96rem, 1.4vw, 1.15rem);
		transition:
			background 180ms ease,
			border-color 180ms ease,
			color 180ms ease,
			transform 180ms ease;
	}

	.terms button:hover,
	.terms button.active {
		border-color: var(--accent);
		background: var(--accent);
		color: white;
		transform: translateY(-2px);
	}

	.explanation {
		display: grid;
		grid-template-columns: auto 1fr;
		gap: 1.4rem;
		min-height: 17rem;
		padding: 2rem;
		border: 1px solid var(--line);
		border-radius: var(--radius-lg);
		background: white;
		box-shadow: var(--shadow);
	}

	.explanation > span {
		color: var(--accent-soft);
		font-family: var(--font-display);
		font-size: 0.9rem;
	}

	h3 {
		margin: 0.8rem 0 1rem;
		font-family: var(--font-display);
		font-size: clamp(1.75rem, 3vw, 2.6rem);
		line-height: 1.05;
		text-transform: capitalize;
	}

	.explanation div > p:last-child {
		margin: 0;
		color: var(--muted);
		line-height: 1.75;
	}

	@media (max-width: 780px) {
		.mission-explorer {
			grid-template-columns: 1fr;
		}
	}
</style>
