<script lang="ts">
	import { articles } from '$lib/data/site';
	import { base } from '$app/paths';
	import T from '$lib/components/T.svelte';
	import { localeTag, t } from '$lib/i18n.svelte';

	let current = $state(0);
	let article = $derived(articles[current]);

	const previous = () => {
		current = (current - 1 + articles.length) % articles.length;
	};

	const next = () => {
		current = (current + 1) % articles.length;
	};
</script>

<div class="carousel">
	<div class="image-wrap">
		<img src={`${base}${article.image}`} alt="" />
		<span>{String(current + 1).padStart(2, '0')} / {String(articles.length).padStart(2, '0')}</span>
	</div>

	<div class="content" aria-live="polite">
		<p class="eyebrow"><T text="Ideas y publicaciones|Ideas and writing" /></p>
		<time datetime={article.date}
			>{new Intl.DateTimeFormat(localeTag(), { dateStyle: 'long' }).format(
				new Date(article.date)
			)}</time
		>
		<h3><T text={article.title} /></h3>
		<p><T text={article.summary} /></p>

		<div class="controls">
			<button type="button" onclick={previous} aria-label={t('Publicación anterior|Previous post')}>←</button>
			<div class="track">
				{#each articles as _, index}
					<button
						type="button"
						class:active={index === current}
						onclick={() => (current = index)}
						aria-label={`${t('Ver publicación|View post')} ${index + 1}`}
					></button>
				{/each}
			</div>
			<button type="button" onclick={next} aria-label={t('Siguiente publicación|Next post')}>→</button>
		</div>
	</div>
</div>

<style>
	.carousel {
		display: grid;
		grid-template-columns: minmax(0, 1.05fr) minmax(20rem, 0.95fr);
		min-height: 31rem;
		overflow: hidden;
		border: 1px solid var(--line);
		border-radius: var(--radius-lg);
		background: white;
		box-shadow: var(--shadow);
	}

	.image-wrap {
		position: relative;
		min-height: 28rem;
		overflow: hidden;
		background: #091c29;
	}

	.image-wrap::after {
		position: absolute;
		inset: 0;
		background: linear-gradient(180deg, transparent 55%, rgba(3, 20, 31, 0.55));
		content: '';
	}

	.image-wrap img {
		width: 100%;
		height: 100%;
		object-fit: cover;
		transition: transform 500ms ease;
	}

	.image-wrap span {
		position: absolute;
		z-index: 1;
		right: 1.5rem;
		bottom: 1.25rem;
		color: white;
		font-size: calc(0.72rem + var(--fs-bump));
		letter-spacing: 0.14em;
	}

	.content {
		display: flex;
		flex-direction: column;
		justify-content: center;
		padding: clamp(2rem, 5vw, 4.5rem);
	}

	time {
		margin: 1.4rem 0 0.75rem;
		color: var(--muted);
		font-size: calc(0.75rem + var(--fs-bump));
		text-transform: capitalize;
	}

	h3 {
		max-width: 18ch;
		margin: 0;
		color: var(--ink);
		font-family: var(--font-display);
		font-size: clamp(1.7rem, 3vw, 2.7rem);
		line-height: 1.08;
	}

	.content > p:last-of-type {
		margin: 1.25rem 0 0;
		color: var(--muted);
		line-height: 1.75;
	}

	.controls {
		display: flex;
		align-items: center;
		gap: 1rem;
		margin-top: 2rem;
	}

	.controls > button {
		display: grid;
		width: 2.8rem;
		height: 2.8rem;
		place-items: center;
		border: 1px solid var(--line);
		border-radius: 999px;
		background: white;
		color: var(--ink);
		font-size: 1.1rem;
		transition:
			border-color 160ms ease,
			background 160ms ease,
			color 160ms ease;
	}

	.controls > button:hover {
		border-color: var(--accent);
		background: var(--accent);
		color: white;
	}

	.track {
		display: flex;
		align-items: center;
		gap: 0.4rem;
	}

	.track button {
		width: 0.45rem;
		height: 0.45rem;
		padding: 0;
		border: 0;
		border-radius: 99px;
		background: #cad4d8;
		transition:
			width 180ms ease,
			background 180ms ease;
	}

	.track button.active {
		width: 1.5rem;
		background: var(--accent);
	}

	@media (max-width: 760px) {
		.carousel {
			grid-template-columns: 1fr;
		}

		.image-wrap {
			min-height: 17rem;
		}
	}
</style>
