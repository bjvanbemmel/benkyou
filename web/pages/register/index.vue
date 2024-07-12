<script setup lang="ts">
import { reactive, ref, type Ref } from 'vue'
import type { FormInstance, ComponentSize, FormRules } from 'element-plus'
import { NuxtLink } from '#components'
import { match, P } from 'ts-pattern';

const { $fetchApi } = useNuxtApp()

interface RuleForm {
  email: string,
  username: string,
  password: string,
  confirmPassword: string,
}

const formSize: Ref = ref<ComponentSize>('default')
const formRef: Ref = ref<FormInstance>()
const formError: Ref<Error | null> = ref(null)
const form: Partial<RuleForm> = reactive<RuleForm>({
  email: '',
  username: '',
  password: '',
  confirmPassword: '',
})

const disableSubmit: Ref<boolean> = ref(false);

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const validatePasswords = (_: any, value: string, callback: any) => {
  if (value === '') return callback(new Error('Please repeat the aforementioned password'))
  if (value !== form.password) return callback(new Error('Passwords must be identical'))

  return callback()
}

const rules: Partial<FormRules<RuleForm>> = reactive<FormRules<RuleForm>>({
  email: [
    {
      required: true,
      type: 'email',
      message: 'Please enter a valid e-mail address',
      trigger: 'blur',
    },
  ],
  username: [
    {
      required: true,
      min: 3,
      max: 30,
      message: 'A valid username must be between 3 and 30 characters long',
      trigger: 'blur',
    },
  ],
  password: [
    {
      required: true,
      min: 8,
      max: 255,
      message: 'A valid password must be between 8 and 255 characters long',
      trigger: 'blur',
    },
  ],
  confirmPassword: [
    {
      required: true,
      validator: validatePasswords,
      trigger: 'blur',
    },
  ],
})

const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) {
    return
  }

  formError.value = null

  await formEl.validate(async (valid, _) => {
    if (!valid) {
      return
    }

    const err = await createUser()
    if (err !== null) {
      formError.value = err.status === 409
        ? formError.value = new Error('A user with the given e-mail address or username already exists')
        : formError.value = new Error('Something went wrong while registering your account')

      formEl.resetFields([ 'password', 'confirmPassword' ])

      return
    }

    if (await getToken() !== null) {
      formError.value = new Error('Something went wrong while registering your account')
      formEl.resetFields([ 'password', 'confirmPassword' ])

      return
    }

    navigateTo('/', {
      external: true,
    })
  })
}

async function createUser(): Promise<ErrorResponse | null> {
  const data: Response<User> | ErrorResponse = await $fetchApi<Response<User>>('users', {
    method: 'POST',
    body: {
      username: form.username,
      email: form.email,
      password: form.password,
      confirm_password: form.confirmPassword,
    },
  }).catch(e => e.data)

  const _user = match(data)
    .with({ data: P.not(undefined) }, () => data as Response<User>)
    .otherwise(() => undefined)

  if (typeof _user === 'undefined') return data as ErrorResponse

  return null
}

async function getToken(): Promise<ErrorResponse | null> {
  const data: Response<Token> | ErrorResponse = await $fetchApi<Response<Token>>('auth/login', {
    method: 'POST',
    body: {
      username: form.username,
      password: form.password,
    },
  }).catch(e => e.data)

  const _token = match(data)
    .with({ data: P.not(undefined) }, () => data as Response<Token>)
    .otherwise(() => undefined)

  if (typeof _token === 'undefined') return data as ErrorResponse

  const token = useCookie('token')
  token.value = _token.data.value

  return null
}
</script>

<template>
  <NuxtLayout name="unauthenticated">
    <div class="flex justify-center items-center w-full h-full">
      <div class="min-w-96 w-full max-w-2xl">
        <el-row :gutter="0" class="min-h-96">
          <el-col :span="16" class="bg-zinc-800 rounded-l-xl">
            <div class="h-full flex flex-col justify-between">
              <el-form
                ref="formRef"
                :rules="rules"
                :model="form"
                :size="formSize"
                status-icon
                label-width="auto"
                class="flex gap-4 flex-col p-12"
                @keyup.enter="submitForm(formRef)"
              >
                <el-form-item>
                  <el-alert v-if="formError" :title="formError.message" type="error" />
                </el-form-item>
                <el-form-item label="E-mail" prop="email">
                  <el-input
                    v-model="form.email"
                    placeholder="E-mail"
                    class="bg-zinc-900/50"
                  />
                </el-form-item>
                <el-form-item label="Username" prop="username">
                  <el-input
                    v-model="form.username"
                    placeholder="Username"
                    class="bg-zinc-900/50"
                  />
                </el-form-item>
                <el-form-item label="Password" prop="password">
                  <el-input
                    v-model="form.password"
                    placeholder="Password"
                    type="password"
                    class="bg-zinc-900/50"
                  />
                </el-form-item>
                <el-form-item label="Repeat password" prop="confirmPassword">
                  <el-input
                    v-model="form.confirmPassword"
                    placeholder="Repeat password"
                    type="password"
                    class="bg-zinc-900/50"
                  />
                </el-form-item>
                <el-form-item>
                  <el-button
                    type="primary"
                    class="mt-4 ml-auto"
                    :disabled="disableSubmit"
                    @click="submitForm(formRef)"
                  >
                    Create account
                  </el-button>
                </el-form-item>
              </el-form>
              <div class="flex gap-1 items-center m-8">
                <span class="text-zinc-300 text-sm">Already have an account?</span>
                <NuxtLink class="text-sm text-blue-300 font-bold hover:underline" to="/login">
                  Return to login.
                </NuxtLink>
              </div>
            </div>
          </el-col>
          <el-col :span="8">
            <div class="rounded-r-xl opacity-75 bg-cover bg-right bg-[url('/img/login-picture.jpg')] w-full h-full"/>
          </el-col>
        </el-row>
      </div>
    </div>
  </NuxtLayout>
</template>
