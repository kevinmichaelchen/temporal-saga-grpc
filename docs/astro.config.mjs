import { defineConfig } from "astro/config";
import starlight from "@astrojs/starlight";

// https://astro.build/config
export default defineConfig({
  site: "https://kevinmichaelchen.github.io",
  base: "/temporal-saga-grpc",
  integrations: [
    starlight({
      title: "temporal-saga-grpc",
      social: {
        github: "https://github.com/kevinmichaelchen/temporal-saga-grpc",
      },
      editLink: {
        baseUrl:
          "https://github.com/kevinmichaelchen/temporal-saga-grpc/tree/main/docs/",
      },
      pagination: true,
      lastUpdated: true,
      sidebar: [
        {
          label: "Architecture",
          items: [
            // Each item here is one entry in the navigation menu.
            {
              label: "Overview",
              link: "/architecture/overview/",
            },
            {
              label: "Why Temporal?",
              link: "/architecture/temporal/",
            },
            {
              label: "Why Buf?",
              link: "/architecture/buf/",
            },
          ],
        },
        {
          label: "Reference",
          autogenerate: { directory: "reference" },
        },
      ],
    }),
  ],
});
