<script lang="ts">
	import { base } from '$app/paths';
	import T from '$lib/components/T.svelte';
	import { t } from '$lib/i18n.svelte';
	import { ApiError, sendContactMessage } from '$lib/api';

	/* Los mínimos son los mismos que valida el backend (handlers/contact.go): pedirlos
	 * aquí convierte un 400 con texto genérico en un aviso del propio navegador. */
	const MIN_NAME = 2;
	const MIN_MESSAGE = 10;
	const MAX_MESSAGE = 4000;

	let form = $state({
		name: '',
		email: '',
		message: ''
	});
	let status = $state<'idle' | 'sending' | 'sent' | 'error'>('idle');
	let errorText = $state('');

	/*
	 * El backend responde en español y este sitio es bilingüe, así que el texto que
	 * se muestra se decide por el código HTTP —que sí es neutro— y no por el cuerpo.
	 */
	function messageForStatus(httpStatus: number): string {
		if (httpStatus === 429) {
			return t(
				'Has enviado varios mensajes seguidos. Inténtalo de nuevo en unos minutos.|You have sent several messages in a row. Please try again in a few minutes.'
			);
		}
		if (httpStatus === 400) {
			return t(
				'Revisa los datos del formulario e inténtalo de nuevo.|Please check the form fields and try again.'
			);
		}
		return t(
			'No pudimos enviar tu mensaje. Escríbenos a contacto@un.pe y lo vemos.|We could not send your message. Write to us at contacto@un.pe and we will take a look.'
		);
	}

	async function submit() {
		if (status === 'sending') return;

		status = 'sending';
		errorText = '';

		try {
			await sendContactMessage({
				Name: form.name.trim(),
				Email: form.email.trim(),
				Message: form.message.trim()
			});
			// `Notified: false` significa que el correo de aviso falló, no que el
			// mensaje se haya perdido: quedó guardado y alguien lo va a leer igual.
			status = 'sent';
			form = { name: '', email: '', message: '' };
		} catch (error) {
			status = 'error';
			errorText = messageForStatus(error instanceof ApiError ? error.status : 0);
		}
	}
</script>

<section class="contact-section" id="contacto">
	<div class="contact-copy">
		<h2><T text="Contáctanos|Contact us" /></h2>

		<p>
			<T
				text="Si tienes una iniciativa de código abierto o quieres participar en el desarrollo de alguna de las nuestras, escríbenos: podemos ayudarte y asesorarte.|If you have an open-source initiative, or you would like to take part in the development of one of ours, write to us: we can help and advise you."
			/>
		</p>

		<figure class="contact-art">
			<img
				src={`${base}/svg/conversation-dark.svg`}
				alt={t('Dos personas conversando frente a un café|Two people talking over a coffee')}
				loading="lazy"
			/>
		</figure>
	</div>

	<form onsubmit={(event) => { event.preventDefault(); submit(); }}>
		<div class="field">
			<label for="contact-name"><T text="Nombre|Name" /></label>
			<input
				id="contact-name"
				bind:value={form.name}
				autocomplete="name"
				minlength={MIN_NAME}
				maxlength="120"
				required
			/>
		</div>
		<div class="field">
			<label for="contact-email"><T text="Correo|Email" /></label>
			<input
				id="contact-email"
				type="email"
				bind:value={form.email}
				autocomplete="email"
				required
			/>
		</div>
		<div class="field">
			<label for="contact-message"><T text="Mensaje|Message" /></label>
			<textarea
				id="contact-message"
				bind:value={form.message}
				rows="5"
				minlength={MIN_MESSAGE}
				maxlength={MAX_MESSAGE}
				required
			></textarea>
		</div>
		<button type="submit" disabled={status === 'sending'}>
			{#if status === 'sending'}
				<T text="Enviando…|Sending…" />
			{:else}
				<T text="Enviar mensaje|Send message" />
			{/if}
			<span>↗</span>
		</button>
		{#if status === 'sent'}
			<p class="form-status" role="status">
				<T
					text="Recibimos tu mensaje. Te responderemos al correo que nos dejaste.|We have your message. We will reply to the address you left us."
				/>
			</p>
		{:else if status === 'error'}
			<p class="form-status is-error" role="alert">{errorText}</p>
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
		font-size: var(--fs-h1);
		font-weight: 500;
		line-height: var(--lh-display);
		letter-spacing: -0.04em;
	}

	.contact-copy > p:not(.eyebrow) {
		max-width: 31rem;
		color: rgba(255, 255, 255, 0.62);
		font-size: var(--fs-md);
		line-height: var(--lh-relaxed);
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
		font-size: var(--fs-2xs);
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
		font-size: var(--fs-xs);
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

	/* Es la flecha del botón, no texto: no entra en la escala. */
	form button span {
		font-size: 1.15rem;
	}

	form button:disabled {
		opacity: 0.65;
		transform: none;
		cursor: progress;
	}

	.form-status {
		margin: 1rem 0 0;
		color: var(--aqua);
		font-size: var(--fs-xs);
		line-height: var(--lh-normal);
	}

	.form-status.is-error {
		color: #ff9aa5;
	}

	@media (max-width: 760px) {
		.contact-section {
			grid-template-columns: 1fr;
		}
	}
</style>
