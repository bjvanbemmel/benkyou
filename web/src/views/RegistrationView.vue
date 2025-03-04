<template>
  <main class="h-screen w-screen flex justify-center items-center">
    <div class="p-8 w-96 rounded-md bg-zinc-800">
      <form
        class="flex flex-col gap-4"
        @submit.prevent="registerAccount"
      >

        <FormTextInput
          :type="TextInputTypes.EMAIL"
          :required="true"
          :error="form.email.error"
          autocomplete="email"
          placeholder="Email"
          v-model="form.email.value"
        />

        <div class="gap-4 grid grid-rows-1 grid-cols-5">
          <FormTextInput
            :type="TextInputTypes.TEXT"
            :required="true"
            :error="form.first_name.error"
            autocomplete="first_name"
            placeholder="First name"
            class="col-span-2"
            v-model="form.first_name.value"
          />

          <FormTextInput
            :type="TextInputTypes.TEXT"
            :required="true"
            :error="form.last_name.error"
            autocomplete="last_name"
            placeholder="Last name"
            class="col-span-3"
            v-model="form.last_name.value"
          />
        </div>

        <FormTextInput
          :type="TextInputTypes.PASSWORD"
          :required="true"
          :error="form.password.error"
          autocomplete="password"
          placeholder="Password"
          v-model="form.password.value"
        />

        <FormTextInput
          :type="TextInputTypes.PASSWORD"
          :required="true"
          :error="form.confirm_password.error"
          placeholder="Confirm password"
          v-model="form.confirm_password.value"
        />

        <FormTextInput
          :type="TextInputTypes.TEXT"
          :required="true"
          :error="form.access_token.error"
          label="This token must be provided by the site administrator to gain access."
          placeholder="Access token"
          v-model="form.access_token.value"
        />

        <div class="w-full my-2 border-1 border-zinc-600"></div>

        <FormButton
          :submit="true"
          :type="ButtonTypes.PRIMARY"
          :disabled="loading"
        >
          Create account
        </FormButton>

        <RouterLink
          draggable="false"
          :to="{ name: 'login' }"
        >
          <FormButton
            class="w-full"
            :submit="true"
            :type="ButtonTypes.SECONDARY"
          >
            I already have an account
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
import { ButtonTypes, TextInputTypes } from '@/types/enums';
import type { FormValue } from '@/types/interfaces';
import axios from 'axios';
import { ref, type Ref } from 'vue';
import { RouterLink } from 'vue-router';

interface Form {
  email: FormValue<string>,
  first_name: FormValue<string>,
  last_name: FormValue<string>,
  password: FormValue<string>,
  confirm_password: FormValue<string>,
  access_token: FormValue<string>,
}

const loading: Ref<boolean> = ref(false);

const form: Ref<Form> = ref({
  email: { value: '', error: null },
  first_name: { value: '', error: null },
  last_name: { value: '', error: null },
  password: { value: '', error: null },
  confirm_password: { value: '', error: null },
  access_token: { value: '', error: null },
});

function registerAccount() {
  loading.value = true;
  form.value.email.error = null;
  form.value.first_name.error = null;
  form.value.last_name.error = null;
  form.value.password.error = null;
  form.value.confirm_password.error = null;
  form.value.access_token.error = null;

  axios.post('/auth/register', {
    email: form.value.email.value,
    first_name: form.value.first_name.value,
    last_name: form.value.last_name.value,
    password: form.value.password.value,
    confirm_password: form.value.confirm_password.value,
    access_token: form.value.access_token.value,
  }).then(() => router.push({ name: 'login' }))
    .catch((err) => {
      const message: string = err.response.data.message;

      if (message.includes('email', 0)) {
        form.value.email.error = message;
        return;
      }

      if (message === 'resource already exists') {
        form.value.email.error = 'this email has already been registered';
        return;
      }

      if (message.includes('first_name', 0)) {
        form.value.first_name.error = message;
        return;
      }

      if (message.includes('last_name', 0)) {
        form.value.last_name.error = message;
        return;
      }

      if (message.includes('confirm_password', 0)) {
        form.value.confirm_password.error = message;
        return;
      }

      if (message.includes('password', 0)) {
        form.value.password.error = message;
        return;
      }

      if (message.includes('access_token', 0) || message.includes('access token', 0)) {
        form.value.access_token.error = message;
        return;
      }
  })
  .finally(() => loading.value = false);
}
</script>
