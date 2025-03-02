<template>
  <button
    class="p-2 rounded-md focus:outline-1 hover:cursor-pointer"
    :class="{
      'bg-blue-500 font-bold text-white focus:outline-blue-300 active:bg-blue-400': props.type === ButtonTypes.PRIMARY,
      'border-1 border-zinc-500 focus:outline-zinc-300 active:bg-zinc-700': props.type === ButtonTypes.SECUNDARY,
    }"
    :type="props.submit ? 'submit' : 'button'"
  >
    <slot/>
  </button>
</template>

<script setup lang="ts">
import { ButtonTypes } from '@/types/enums';

const props = defineProps<{
  submit: boolean,
  type: ButtonTypes,
}>();
</script>
