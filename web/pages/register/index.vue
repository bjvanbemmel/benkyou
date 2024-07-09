<template>
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
              class="flex flex-col p-12"
              @keyup.enter="validateForm(formRef)"
            >
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
                  @click="validateForm(formRef)"
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
</template>

<script setup lang="ts">
import { reactive, ref, type Ref } from 'vue'
import type { FormInstance, ComponentSize, FormRules } from 'element-plus'
import { NuxtLink } from '#components'

interface RuleForm {
  email: string,
  username: string,
  password: string,
  confirmPassword: string,
}

const formSize: Ref = ref<ComponentSize>('default')

const formRef: Ref = ref<FormInstance>()

const form: Partial<RuleForm> = reactive<RuleForm>({
  email: '',
  username: '',
  password: '',
  confirmPassword: '',
})

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const validatePasswords = (_: any, value: string, callback: any) => {
  if (value === '') return callback(new Error('Please repeat the aforementioned password'))
  if (value !== form.password) return callback(new Error('Passwords must be identical'))

  return callback()
}

const rules: Partial<FormRules<RuleForm>> = reactive<FormRules<RuleForm>>({
  email: [
    { required: true, type: 'email', message: 'Please enter a valid e-mail address', trigger: 'blur', },
  ],
  username: [
    { required: true, message: 'Please enter your username', trigger: 'blur', },
  ],
  password: [
    { required: true, message: 'Please enter your password', trigger: 'blur', },
  ],
  confirmPassword: [
    { required: true, validator: validatePasswords, trigger: 'blur', },
  ],
})

const validateForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate((valid, _) => {
    if (!valid) console.log('sadge')
  })
}
</script>
