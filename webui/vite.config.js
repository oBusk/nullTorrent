import { defineConfig } from "vite";

export default defineConfig({
	build: {
		outDir: "../internal/webserver/dist",
		emptyOutDir: true,
	},
	server: {
		proxy: {
			"/api": "http://localhost:8080",
		},
	},
});
