import type { UseFetchOptions } from 'nuxt/app';

export function useFetchApi<T>(
  url: string | (() => string),
  options: UseFetchOptions<T> = {}
) {
  return useFetch(url, {
    ...options,
    $fetch: useNuxtApp().$fetchApi,
  })
}
