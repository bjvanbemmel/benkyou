export default defineNuxtPlugin(() => {
  const token = useCookie('token')
  const config = useRuntimeConfig()

  const $fetchApi = $fetch.create({
    baseURL: import.meta.server ? config.baseURL : config.public.baseURL,
    onRequest({ options }) {
      if (token.value) {
        options.headers = options.headers || {}
        options.headers.Authorization = `Bearer ${token.value}`
      }
    },
  })

  return {
    provide: {
      fetchApi: $fetchApi,
    }
  }
})
