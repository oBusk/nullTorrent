import { defineConfig } from "vite";
import checker from "vite-plugin-checker";
import solid from "vite-plugin-solid";

export default defineConfig({
	plugins: [solid(), checker({ typescript: true })],
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
