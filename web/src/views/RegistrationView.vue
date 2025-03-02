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
          @input="(val) => form.email.value = val"
        />

        <div class="gap-4 grid grid-rows-1 grid-cols-5">
          <FormTextInput
            :type="TextInputTypes.TEXT"
            :required="true"
            :error="form.firstName.error"
            autocomplete="first_name"
            placeholder="First name"
            class="col-span-2"
            @input="(val) => form.firstName.value = val"
          />

          <FormTextInput
            :type="TextInputTypes.TEXT"
            :required="true"
            :error="form.lastName.error"
            autocomplete="last_name"
            placeholder="Last name"
            class="col-span-3"
            @input="(val) => form.lastName.value = val"
          />
        </div>

        <FormTextInput
          :type="TextInputTypes.PASSWORD"
          :required="true"
          :error="form.password.error"
          autocomplete="password"
          placeholder="Password"
          @input="(val) => form.password.value = val"
        />

        <FormTextInput
          :type="TextInputTypes.PASSWORD"
          :required="true"
          :error="form.confirm_password.error"
          placeholder="Confirm password"
          @input="(val) => form.confirm_password.value = val"
        />

        <FormTextInput
          :type="TextInputTypes.TEXT"
          :required="true"
          :error="form.accessToken.error"
          label="This token must be provided by the site administrator to gain access."
          placeholder="Access token"
          @input="(val) => form.accessToken.value = val"
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
            :type="ButtonTypes.SECUNDARY"
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
  email: FormValue,
  firstName: FormValue,
  lastName: FormValue,
  password: FormValue,
  confirm_password: FormValue,
  accessToken: FormValue,
}

const loading: Ref<boolean> = ref(false);

const form: Ref<Form> = ref({
  email: { value: '', error: null },
  firstName: { value: '', error: null },
  lastName: { value: '', error: null },
  password: { value: '', error: null },
  confirm_password: { value: '', error: null },
  accessToken: { value: '', error: null },
});

function registerAccount() {
  loading.value = true;
  form.value.email.error = null;
  form.value.firstName.error = null;
  form.value.lastName.error = null;
  form.value.password.error = null;
  form.value.confirm_password.error = null;
  form.value.accessToken.error = null;

  axios.post('/auth/register', {
    email: form.value.email.value,
    first_name: form.value.firstName.value,
    last_name: form.value.lastName.value,
    password: form.value.password.value,
    confirm_password: form.value.confirm_password.value,
    access_token: form.value.accessToken.value,
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
        form.value.firstName.error = message;
        return;
      }

      if (message.includes('last_name', 0)) {
        form.value.lastName.error = message;
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
        form.value.accessToken.error = message;
        return;
      }
  })
  .finally(() => loading.value = false);
}
</script>
