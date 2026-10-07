import type { Status } from "./generated/api.ts";

export type * from "./generated/api.ts";

export interface ClientOptions {
	/** Where the nullTorrent daemon is reachable, e.g. `http://nas:8080`. Defaults to the current origin. */
	baseUrl?: string;
	/** Override the `fetch` implementation, e.g. for tests. */
	fetch?: typeof fetch;
}

export class NullTorrentError extends Error {
	constructor(
		message: string,
		readonly status: number,
	) {
		super(message);
		this.name = "NullTorrentError";
	}
}

export function createClient(options: ClientOptions = {}) {
	const baseUrl = (options.baseUrl ?? "").replace(/\/+$/, "");
	const fetchImpl = options.fetch ?? globalThis.fetch;

	async function get<T>(path: string): Promise<T> {
		const res = await fetchImpl(baseUrl + path);
		if (!res.ok) {
			throw new NullTorrentError(`GET ${path} failed: ${res.status} ${res.statusText}`, res.status);
		}
		return res.json();
	}

	return {
		getStatus: () => get<Status>("/api/status"),
	};
}

export type Client = ReturnType<typeof createClient>;
