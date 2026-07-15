import tailwindcss from "@tailwindcss/vite";

export default defineNuxtConfig({
  vite: {
    plugins: [tailwindcss()],
  },

  css: ["~/assets/css/main.css"],

  runtimeConfig: {
    public: {
      graphqlEndpoint: "http://localhost:8080/v1/graphql",
    },
  },

  compatibilityDate: "2025-07-15",
});