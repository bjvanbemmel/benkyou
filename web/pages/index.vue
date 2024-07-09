<template>
  <ElContainer>
    <ElHeader>
      <h1>Hello, World!</h1>
    </ElHeader>
    <ElMain>
        <ElButton @click="toggleDark()">
          El Butane!
        </ElButton>
    </ElMain>
  </ElContainer>
</template>

<script setup lang="ts">
import { useDark, useToggle } from '@vueuse/core'

const isDark = useDark()
const toggleDark = () => {
  useToggle(isDark)
  console.log(isDark.value)
}
</script>
