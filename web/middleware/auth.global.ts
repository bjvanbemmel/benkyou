export default defineNuxtRouteMiddleware(async (to, _) => {
  const { error: error } = await useFetchApi('auth/identity')

  // Block off login and registration page when authenticated
  if ((to.path === '/login' || to.path === '/register') && error.value === null) {
    return navigateTo('/')
  }

  // When not authenticated, only allow these pages
  if (to.path === '/login' || to.path === '/register') return;

  if (error.value) {
    return navigateTo('/login')
  }
})
