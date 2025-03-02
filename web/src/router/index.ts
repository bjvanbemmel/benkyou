import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import RegistrationView from '@/views/RegistrationView.vue'
import LoginView from '@/views/LoginView.vue'
import { useTokenStore } from '@/stores/token';
import axios from 'axios';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/auth/registration',
      name: 'registration',
      component: RegistrationView,
    },
    {
      path: '/auth/login',
      name: 'login',
      component: LoginView,
    },
  ],
});

router.beforeEach(async (to, from) => {
  if ((to.name === 'login' || to.name === 'registration') && await isAuthenticated()) {
    return { name: 'home' };
  }

  if (to.name !== 'login' && to.name !== 'registration' && !await isAuthenticated()) {
    return { name: 'login' };
  }
});

async function isAuthenticated(): Promise<boolean> {
  const tokenStore = useTokenStore();
  const token = tokenStore.token ?? tokenStore.get();

  try {
    await axios.get('/auth/identity', {
      headers: {
        Authorization: `Bearer ${token}`,
      }
    });
  } catch (e) {
    tokenStore.clear();
    return false;
  }

  return true;
}

export default router
