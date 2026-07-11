import tailwindcss from "@tailwindcss/vite";

export default defineNuxtConfig({
  // Register the Tailwind Vite plugin
  vite: {
    plugins: [tailwindcss()],
  },

  // Register your global CSS file
  css: ["~/assets/css/main.css"],

  // Apollo Client configuration
  modules: ["@vue3-apollo/nuxt"],
  apollo: {
    clients: {
      default: {
        httpEndpoint: "http://localhost:8080/v1/graphql",
        wsEndpoint: "ws://localhost:8080/v1/graphql",
      },
    },
    auth: {
      tokenName: "auth-token",
      authType: "Bearer",
      authHeader: "Authorization",
    },
  },

  compatibilityDate: "2025-07-15",
});
