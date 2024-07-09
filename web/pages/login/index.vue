<template>
  <div class="flex justify-center items-center w-full h-full">
    <div class="min-w-96 w-full max-w-2xl">
      <el-row :gutter="0" class="min-h-96">
        <el-col :span="10">
          <div class="rounded-l-xl opacity-75 bg-cover bg-[url('/img/login-picture.jpg')] w-full h-full"/>
        </el-col>
        <el-col :span="14" class="bg-zinc-800 rounded-r-xl">
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
              <el-form-item>
                <el-button
                  type="primary"
                  class="mt-4 ml-auto"
                  @click="validateForm(formRef)"
                >
                  Login
                </el-button>
              </el-form-item>
            </el-form>
            <div class="flex gap-1 items-center m-8">
              <span class="text-zinc-300 text-sm">No account yet?</span>
              <NuxtLink class="text-sm text-blue-300 font-bold hover:underline" to="/register">
                Register here.
              </NuxtLink>
            </div>
          </div>
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
  username: string,
  password: string,
}

const formSize: Ref = ref<ComponentSize>('default')

const formRef: Ref = ref<FormInstance>()

const form: Partial<RuleForm> = reactive<RuleForm>({
  username: '',
  password: '',
})

const rules: Partial<FormRules<RuleForm>> = reactive<FormRules<RuleForm>>({
  username: [
    { required: true, message: 'Please enter your username', trigger: 'blur', },
  ],
  password: [
    { required: true, message: 'Please enter your password', trigger: 'blur', },
  ],
})

const validateForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate((valid, _) => {
    if (!valid) console.log('sadge')
  })
}
</script>
