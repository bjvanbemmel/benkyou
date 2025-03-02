import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import RegistrationView from '@/views/RegistrationView.vue'
import LoginView from '@/views/LoginView.vue'
import { useTokenStore } from '@/stores/token';
import axios from 'axios';
import { useUserStore } from '@/stores/user';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: { name: 'board' },
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
    {
      path: '/backlog',
      name: 'backlog',
      component: HomeView,
    },
    {
      path: '/sprints',
      name: 'sprints',
      component: HomeView,
    },
    {
      path: '/board',
      name: 'board',
      component: HomeView,
    },
    {
      path: '/braindump',
      name: 'braindump',
      component: HomeView,
    },
    {
      path: '/:pathMatch(.*)*',
      redirect: { name: 'board' },
    },
  ],
});

router.beforeEach(async (to, from) => {
  if ((to.name === 'login' || to.name === 'registration') && await isAuthenticated()) {
    return { name: 'board' };
  }

  if (to.name !== 'login' && to.name !== 'registration' && !await isAuthenticated()) {
    return { name: 'login' };
  }
});

async function isAuthenticated(): Promise<boolean> {
  const tokenStore = useTokenStore();
  const token = tokenStore.token ?? tokenStore.get();

  try {
    let response = await axios.get('/auth/identity', {
      headers: {
        Authorization: `Bearer ${token}`,
      }
    });

    const userStore = useUserStore();
    userStore.set(response.data.data);

  } catch (e) {
    tokenStore.clear();
    return false;
  }

  return true;
}

export default router
