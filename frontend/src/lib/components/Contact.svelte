<script lang="ts">
	import { base } from '$app/paths';

	let form = $state({
		name: '',
		email: '',
		message: ''
	});
	let status = $state<'idle' | 'ready'>('idle');

	function submit() {
		const subject = encodeURIComponent(`Consulta web de ${form.name}`);
		const body = encodeURIComponent(
			`Nombre: ${form.name}\nEmail: ${form.email}\n\nMensaje:\n${form.message}`
		);
		status = 'ready';
		window.location.href = `mailto:contacto@un.pe?subject=${subject}&body=${body}`;
	}
</script>

<section class="contact-section" id="contacto">
	<div class="contact-copy">
		<h2>Contáctanos</h2>

		<p>
			Si tienes una iniciativa de código abierto o quieres participar en el desarrollo de
			alguna de las nuestras, escríbenos: podemos ayudarte y asesorarte.
		</p>

		<figure class="contact-art">
			<img
				src={`${base}/svg/conversation-dark.svg`}
				alt="Dos personas conversando frente a un café"
				loading="lazy"
			/>
		</figure>
	</div>

	<form onsubmit={(event) => { event.preventDefault(); submit(); }}>
		<div class="field">
			<label for="contact-name">Nombre</label>
			<input id="contact-name" bind:value={form.name} autocomplete="name" required />
		</div>
		<div class="field">
			<label for="contact-email">Correo</label>
			<input
				id="contact-email"
				type="email"
				bind:value={form.email}
				autocomplete="email"
				required
			/>
		</div>
		<div class="field">
			<label for="contact-message">Mensaje</label>
			<textarea id="contact-message" bind:value={form.message} rows="5" required></textarea>
		</div>
		<button type="submit">Preparar mensaje <span>↗</span></button>
		{#if status === 'ready'}
			<p class="form-status" role="status">
				Hemos preparado tu mensaje en la aplicación de correo del dispositivo.
			</p>
		{/if}
	</form>
</section>

<style>
	.contact-section {
		display: grid;
		grid-template-columns: minmax(0, 0.9fr) minmax(20rem, 1.1fr);
		gap: clamp(3rem, 9vw, 9rem);
		padding: clamp(5rem, 10vw, 9rem) var(--page-pad);
		background:
			radial-gradient(circle at 8% 88%, rgba(100, 105, 238, 0.24), transparent 30rem),
			radial-gradient(circle at 78% 12%, rgba(0, 216, 179, 0.1), transparent 26rem),
			#0a0b1f;
		color: white;
	}

	.contact-copy h2 {
		margin: 0 0 1.5rem;
		font-family: var(--font-display);
		font-size: clamp(2.5rem, 5vw, 5rem);
		font-weight: 500;
		line-height: 0.98;
		letter-spacing: -0.04em;
	}

	.contact-copy > p:not(.eyebrow) {
		max-width: 31rem;
		color: rgba(255, 255, 255, 0.62);
		line-height: 1.75;
	}

	.contact-art {
		max-width: 24.3rem;
		margin: 2rem 0 0;
	}

	.contact-art img {
		width: 100%;
		height: auto;
	}

	form {
		align-self: center;
		padding: clamp(1.5rem, 4vw, 3rem);
		border: 1px solid rgba(255, 255, 255, 0.12);
		border-radius: var(--radius-lg);
		background: rgba(255, 255, 255, 0.055);
		backdrop-filter: blur(12px);
	}

	.field {
		display: grid;
		gap: 0.65rem;
		margin-bottom: 1.4rem;
	}

	label {
		color: rgba(255, 255, 255, 0.62);
		font-size: 0.7rem;
		font-weight: 700;
		letter-spacing: 0.12em;
		text-transform: uppercase;
	}

	input,
	textarea {
		width: 100%;
		padding: 0.85rem 0;
		border: 0;
		border-bottom: 1px solid rgba(255, 255, 255, 0.22);
		border-radius: 0;
		outline: 0;
		background: transparent;
		color: white;
		font: inherit;
		resize: vertical;
		transition: border-color 160ms ease;
	}

	input:focus,
	textarea:focus {
		border-color: var(--aqua);
	}

	form button {
		display: flex;
		width: 100%;
		align-items: center;
		justify-content: space-between;
		margin-top: 2rem;
		padding: 1rem 1.2rem;
		border: 0;
		border-radius: 999px;
		background: var(--accent);
		color: white;
		font: inherit;
		font-size: 0.78rem;
		font-weight: 800;
		letter-spacing: 0.1em;
		text-transform: uppercase;
		transition:
			background 180ms ease,
			transform 180ms ease;
	}

	form button:hover {
		background: var(--accent-soft);
		transform: translateY(-2px);
	}

	form button span {
		font-size: 1.15rem;
	}

	.form-status {
		margin: 1rem 0 0;
		color: var(--aqua);
		font-size: 0.78rem;
		line-height: 1.5;
	}

	@media (max-width: 760px) {
		.contact-section {
			grid-template-columns: 1fr;
		}
	}
</style>
