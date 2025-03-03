<template>
  <main class="h-screen w-screen flex justify-center items-center">
    <div class="p-8 w-96 rounded-md bg-zinc-800">
      <form
        class="flex flex-col gap-4"
        @submit.prevent="login"
      >
        <FormTextInput
          :type="TextInputTypes.EMAIL"
          :required="true"
          :error="form.email.error"
          autocomplete="email"
          placeholder="Email"
          @input="(val) => form.email.value = val"
        />

        <FormTextInput
          :type="TextInputTypes.PASSWORD"
          :required="true"
          :error="form.password.error"
          autocomplete="password"
          placeholder="Password"
          @input="(val) => form.password.value = val"
        />

        <div class="w-full my-2 border-1 border-zinc-600"></div>

        <FormButton
          :submit="true"
          :type="ButtonTypes.PRIMARY"
          :disabled="loading"
        >
          Login
        </FormButton>

        <RouterLink
          draggable="false"
          :to="{ name: 'registration' }"
        >
          <FormButton
            class="w-full"
            :submit="true"
            :type="ButtonTypes.SECONDARY"
          >
            I don't have an account yet
          </FormButton>
        </RouterLink>
      </form>
    </div>
  </main>
</template>

<script setup lang="ts">
import FormButton from '@/components/FormButton.vue';
import FormTextInput from '@/components/FormTextInput.vue';
import router from '@/router';
import { useTokenStore } from '@/stores/token';
import { ButtonTypes, TextInputTypes } from '@/types/enums';
import type { FormValue } from '@/types/interfaces';
import axios from 'axios';
import { ref, type Ref } from 'vue';

const tokenStore = useTokenStore();

interface Form {
  email: FormValue<string>,
  password: FormValue<string>,
}

const form: Ref<Form> = ref({
  email: { value: '', error: null },
  password: { value: '', error: null },
});

const loading: Ref<boolean> = ref(false);

function login() {
  loading.value = true;
  form.value.email.error = null;
  form.value.password.error = null;

  axios.post('/auth/login', {
    email: form.value.email.value,
    password: form.value.password.value,
  }).then((res) => {
      tokenStore.set(res?.data?.data?.value);
      router.push({ name: 'board' });
  }).catch((err) => {
      const message: string = err.response.data.message;

      if (message.includes('invalid credentials')) {
        form.value.email.error = message;
        form.value.password.error = message;
        return;
      }

      if (message.includes('email', 0)) {
        form.value.email.error = message;
        return;
      }

      if (message.includes('password', 0)) {
        form.value.password.error = message;
        return;
      }
  }).finally(() => {
      loading.value = false;
  });
}
</script>
