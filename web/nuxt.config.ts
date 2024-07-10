// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: {
    enabled: true
  },
  modules: [ '@nuxt/eslint', '@element-plus/nuxt', "@nuxtjs/tailwindcss" ],
  runtimeConfig: {
    public: {
      baseURL: process.env.CLIENT_BASE_URL,
    },
    baseURL: process.env.SERVER_BASE_URL,
  },
  app: {
    head: {
      htmlAttrs: {
        class: 'dark',
      },
    },
  },
  eslint: {},
  elementPlus: {
    themes: [ 'dark' ],
  },
})
