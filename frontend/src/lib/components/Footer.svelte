<script lang="ts">
	import { base } from '$app/paths';

	const year = new Date().getFullYear();
</script>

<footer>
	<div class="footer-brand">
		<img src={`${base}/svg/unicore_logo5.svg`} alt="" />
		<div>
			<strong>UNICORE</strong>
			<span>Soluciones informáticas y empresariales</span>
		</div>
	</div>

	<div class="footer-links">
		<a href={`${base}/`}>Inicio</a>
		<a href={`${base}/nosotros/`}>Nosotros</a>
		<a href={`${base}/portafolio/`}>Portafolio</a>
		<a href="mailto:contacto@un.pe">contacto@un.pe</a>
	</div>

	<div class="legal">
		<span>© {year} Unicore Perú EIRL</span>
		<span>Trujillo · La Libertad · Perú</span>
	</div>
</footer>

<style>
	footer {
		padding: clamp(3rem, 7vw, 5rem) var(--page-pad) 2rem;
		border-top: 1px solid rgba(255, 255, 255, 0.1);
		background: #04141f;
		color: white;
	}

	.footer-brand {
		display: flex;
		align-items: center;
		gap: 1rem;
	}

	.footer-brand img {
		width: 3.8rem;
		height: 3.8rem;
	}

	.footer-brand div {
		display: grid;
		gap: 0.3rem;
	}

	.footer-brand strong {
		font-family: var(--font-display);
		font-size: 1.4rem;
		letter-spacing: 0.16em;
	}

	.footer-brand span {
		color: rgba(255, 255, 255, 0.58);
		font-size: 0.84rem;
	}

	.footer-links {
		display: flex;
		flex-wrap: wrap;
		gap: 1rem 2rem;
		margin: 2.5rem 0 3.5rem;
	}

	.footer-links a {
		color: rgba(255, 255, 255, 0.72);
		font-size: 0.82rem;
		text-decoration: none;
	}

	.footer-links a:hover {
		color: var(--aqua);
	}

	.legal {
		display: flex;
		justify-content: space-between;
		gap: 1rem;
		padding-top: 1.5rem;
		border-top: 1px solid rgba(255, 255, 255, 0.1);
		color: rgba(255, 255, 255, 0.42);
		font-size: 0.74rem;
	}

	@media (max-width: 600px) {
		.legal {
			flex-direction: column;
		}
	}
</style>
