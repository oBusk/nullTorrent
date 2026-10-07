import { createSignal, onCleanup } from "solid-js";

interface Status {
	bytesCompleted: number;
	length: number;
	activePeers: number;
	seeders: number;
}

function formatMB(bytes: number) {
	return (bytes / (1024 * 1024)).toFixed(1) + " MB";
}

export default function App() {
	const [status, setStatus] = createSignal<Status>({
		bytesCompleted: 0,
		length: 0,
		activePeers: 0,
		seeders: 0,
	});
	const percent = () =>
		status().length > 0 ? (status().bytesCompleted / status().length) * 100 : 0;

	const [error, setError] = createSignal<string | null>(null);

	let lastBytes: number | null = null;
	let lastTime = 0;
	let timer: ReturnType<typeof setTimeout> | undefined;
	let disposed = false;

	// Wait for each response before scheduling the next, so slow responses
	// can't overlap and arrive out of order.
	async function poll() {
		try {
			const res = await fetch("/api/status");
			if (!res.ok) {
				throw new Error(`GET /api/status failed: ${res.status} ${res.statusText}`);
			}
			const next: Status = await res.json();

			const now = Date.now();
			if (lastBytes !== null) {
				const rate = (next.bytesCompleted - lastBytes) / ((now - lastTime) / 1000);
				document.title = formatMB(rate) + "/s - Torrent Progress";
			}
			lastBytes = next.bytesCompleted;
			lastTime = now;

			setStatus(next);
			setError(null);
		} catch (err) {
			lastBytes = null;
			document.title = "Disconnected - Torrent Progress";
			setError(err instanceof Error ? err.message : String(err));
		}

		if (!disposed) timer = setTimeout(poll, 1000);
	}

	poll();
	onCleanup(() => {
		disposed = true;
		clearTimeout(timer);
	});

	return (
		<>
			{error() && <p class="error">Can't reach nullTorrent: {error()}</p>}
			<h1>{percent().toFixed(1)}%</h1>
			<progress value={status().bytesCompleted} max={status().length} />
			<div id="stats">
				<span>
					{formatMB(status().bytesCompleted)} / {formatMB(status().length)}
				</span>
				<span>
					{status().activePeers} peers, {status().seeders} seeders
				</span>
			</div>
		</>
	);
}
