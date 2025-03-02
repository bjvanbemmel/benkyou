<template>
  <main class="h-screen w-screen flex justify-center items-center">
    <div class="p-8 w-96 rounded-md bg-zinc-800">
      <form
        class="flex flex-col gap-4"
      >

        <FormTextInput
          :type="TextInputTypes.EMAIL"
          :required="true"
          placeholder="Email"
          class="w-full"
          v-model="email"
        />

        <div class="gap-4 grid grid-rows-1 grid-cols-5">
          <FormTextInput
            :type="TextInputTypes.TEXT"
            :required="true"
            placeholder="First name"
            class="col-span-2"
            v-model="firstName"
          />

          <FormTextInput
            :type="TextInputTypes.TEXT"
            :required="true"
            placeholder="Last name"
            class="col-span-3"
            v-model="lastName"
          />
        </div>

        <FormTextInput
          :type="TextInputTypes.PASSWORD"
          :required="true"
          placeholder="Password"
          class="w-full"
          v-model="password"
        />


        <FormTextInput
          :type="TextInputTypes.TEXT"
          :required="true"
          label="This token must be provided by the site administrator to gain access."
          placeholder="Access token"
          class="w-full"
          v-model="accessToken"
        />

        <div class="w-full my-2 border-1 border-zinc-600"></div>

        <FormButton
          :submit="true"
          :type="ButtonTypes.PRIMARY"
        >
          Create account
        </FormButton>

        <RouterLink class="w-full" :to="{ name: 'login' }">
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
import { ButtonTypes, TextInputTypes } from '@/types/enums';
import { ref, type Ref } from 'vue';
import { RouterLink } from 'vue-router';

const email: Ref<string> = ref('');
const firstName: Ref<string> = ref('');
const lastName: Ref<string> = ref('');
const password: Ref<string> = ref('');
const accessToken: Ref<string> = ref('');
</script>
