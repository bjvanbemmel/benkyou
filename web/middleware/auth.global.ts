export default defineNuxtRouteMiddleware(async (to, _) => {
  const { $fetchApi } = useNuxtApp()
  let error: ErrorResponse | undefined

  await $fetchApi('auth/identity')
    .catch(e => error = e.data)

  // Block off login and registration page when authenticated
  if ((to.path === '/login' || to.path === '/register') && typeof error === 'undefined') {
    return navigateTo('/')
  }

  // When not authenticated, only allow these pages
  if (to.path === '/login' || to.path === '/register') return;

  if (error) {
    return navigateTo('/login')
  }
})
