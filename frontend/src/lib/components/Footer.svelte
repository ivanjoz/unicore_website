<script lang="ts">
	import { base } from '$app/paths';
	import T from '$lib/components/T.svelte';

	const year = new Date().getFullYear();
</script>

<footer>
	<div class="footer-top">
		<div class="footer-brand">
			<img src={`${base}/svg/logo_unicore_horizontal_light_mini.svg`} alt="Unicore Labs" />
			<span>
				<T
					text="Iniciativas de código abierto que democratizan el acceso a la tecnología|Open-source initiatives that democratize access to technology"
				/>
			</span>
		</div>

		<div class="footer-contact">
			<div class="footer-detail">
				<span><T text="ESCRÍBENOS|WRITE TO US" /></span>
				<a href="mailto:contacto@un.pe">contacto@un.pe</a>
			</div>
			<div class="footer-detail">
				<span><T text="CÓDIGO|CODE" /></span>
				<a href="https://github.com/ivanjoz?tab=repositories" target="_blank" rel="noreferrer">
					github.com/ivanjoz ↗
				</a>
			</div>
		</div>
	</div>

	<div class="legal">
		<span>© {year} Unicore Labs</span>
		<span><T text="Trujillo · La Libertad · Perú|Trujillo · La Libertad · Peru" /></span>
	</div>
</footer>

<style>
	footer {
		padding: clamp(3rem, 7vw, 5rem) var(--page-pad) 2rem;
		border-top: 1px solid rgba(255, 255, 255, 0.1);
		background: #07081a;
		color: white;
	}

	.footer-top {
		display: flex;
		flex-wrap: wrap;
		align-items: start;
		justify-content: space-between;
		gap: 2.5rem clamp(2rem, 6vw, 5rem);
		margin-bottom: clamp(2.5rem, 5vw, 3.5rem);
	}

	.footer-brand {
		display: grid;
		gap: 1rem;
		justify-items: start;
	}

	.footer-brand img {
		width: auto;
		height: 3.4rem;
	}

	.footer-brand span {
		max-width: 34rem;
		color: rgba(255, 255, 255, 0.58);
		font-size: calc(0.84rem + var(--fs-bump));
		line-height: 1.6;
	}

	.footer-contact {
		display: flex;
		flex-wrap: wrap;
		gap: 1.6rem clamp(2rem, 5vw, 4rem);
	}

	.footer-detail {
		display: grid;
		gap: 0.35rem;
	}

	.footer-detail span {
		color: var(--aqua);
		font-size: calc(0.64rem + var(--fs-bump));
		font-weight: 700;
		letter-spacing: 0.16em;
	}

	.footer-detail a {
		color: white;
		font-size: calc(0.94rem + var(--fs-bump));
		text-decoration: none;
	}

	.footer-detail a:hover {
		color: var(--aqua);
	}

	.legal {
		display: flex;
		justify-content: space-between;
		gap: 1rem;
		padding-top: 1.5rem;
		border-top: 1px solid rgba(255, 255, 255, 0.1);
		color: rgba(255, 255, 255, 0.42);
		font-size: calc(0.74rem + var(--fs-bump));
	}

	@media (max-width: 600px) {
		.legal {
			flex-direction: column;
		}
	}
</style>
