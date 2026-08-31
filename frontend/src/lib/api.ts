/*
 * Punto de entrada al backend (Go sobre Lambda, tras una Function URL).
 *
 * La URL viaja como variable de build (`VITE_API_URL`) porque el sitio es
 * estático: no hay servidor que la resuelva en tiempo de ejecución. El valor
 * por defecto es la Function URL desplegada, así que `bun run build` sin
 * configurar nada produce un sitio que ya apunta a producción.
 *
 * El origen desde el que se llama debe estar en `backend.allowed_origins`
 * (config.toml); de lo contrario el navegador bloquea la respuesta por CORS.
 */
const DEFAULT_API_URL = 'https://4zmdd4q2jxvhs252k2z2oxfnyi0wxdtd.lambda-url.us-east-1.on.aws/';

export const API_URL = (import.meta.env.VITE_API_URL || DEFAULT_API_URL).replace(/\/+$/, '');

export type ContactPayload = {
	Name: string;
	Email: string;
	Company?: string;
	Message: string;
};

/** Lo que responde `POST /p-contact-message` cuando el mensaje se guardó. */
export type ContactResult = {
	/** El mensaje quedó registrado. */
	Received: boolean;
	/** El correo de aviso salió. `false` no significa que se haya perdido nada. */
	Notified: boolean;
};

/**
 * Error de una llamada al backend, con el código HTTP a la vista para que quien
 * llama pueda distinguir «datos inválidos» de «límite alcanzado» sin leer texto.
 */
export class ApiError extends Error {
	status: number;

	constructor(status: number, message: string) {
		super(message);
		this.name = 'ApiError';
		this.status = status;
	}
}

export async function sendContactMessage(payload: ContactPayload): Promise<ContactResult> {
	let response: Response;
	try {
		response = await fetch(`${API_URL}/p-contact-message`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(payload)
		});
	} catch (error) {
		// Sin red, DNS caído o CORS rechazado: `fetch` lanza antes de que haya
		// respuesta que leer, y el código 0 marca justamente ese caso.
		throw new ApiError(0, error instanceof Error ? error.message : 'network error');
	}

	// El cuerpo de error es `{"error":"…"}`, pero un fallo de infraestructura
	// (un 502 del propio Lambda) llega como HTML; por eso el parseo no decide.
	let data: Record<string, unknown> = {};
	try {
		data = await response.json();
	} catch {
		data = {};
	}

	if (!response.ok) {
		throw new ApiError(response.status, typeof data.error === 'string' ? data.error : '');
	}

	return { Received: data.Received === true, Notified: data.Notified === true };
}
