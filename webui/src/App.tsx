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

	let lastBytes: number | null = null;
	let lastTime = 0;

	async function poll() {
		const res = await fetch("/api/status");
		const next: Status = await res.json();

		const now = Date.now();
		if (lastBytes !== null) {
			const rate = (next.bytesCompleted - lastBytes) / ((now - lastTime) / 1000);
			document.title = formatMB(rate) + "/s - Torrent Progress";
		}
		lastBytes = next.bytesCompleted;
		lastTime = now;

		setStatus(next);
	}

	poll();
	const timer = setInterval(poll, 1000);
	onCleanup(() => clearInterval(timer));

	return (
		<>
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
