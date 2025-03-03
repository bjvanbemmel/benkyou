import { createRouter, createWebHistory } from 'vue-router'
import { useTokenStore } from '@/stores/token';
import axios from 'axios';
import { useUserStore } from '@/stores/user';
import RegistrationView from '@/views/RegistrationView.vue'
import LoginView from '@/views/LoginView.vue'
import BacklogView from '@/views/BacklogView.vue';

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
      component: BacklogView,
    },
    {
      path: '/sprints',
      name: 'sprints',
      component: BacklogView,
    },
    {
      path: '/board',
      name: 'board',
      component: BacklogView,
    },
    {
      path: '/braindump',
      name: 'braindump',
      component: BacklogView,
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

  if (token === null) return false;

  try {
    let response = await axios.get('/auth/identity', {
      headers: {
        Authorization: `Bearer ${token}`,
      }
    });

    tokenStore.set(token);

    const userStore = useUserStore();
    userStore.set(response.data.data);

  } catch (e) {
    tokenStore.clear();
    return false;
  }

  return true;
}

export default router
