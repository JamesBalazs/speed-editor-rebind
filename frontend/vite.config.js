import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import wails from "@wailsio/runtime/plugins/vite";
import { quasar, transformAssetUrls } from "@quasar/vite-plugin";

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        wails("./bindings"),
        vue({ template: { transformAssetUrls } }),
        quasar({ sassVariables: "src/quasar-variables.sass" }),
    ],
});
