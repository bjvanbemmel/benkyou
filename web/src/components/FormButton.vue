<template>
  <button
    class="p-2 rounded-md focus:outline-1 hover:cursor-pointer"
    :disabled="props.disabled"
    :class="{
      'bg-blue-500/80 hover:bg-blue-500/60 border-1 border-blue-400 disabled:bg-blue-900 disabled:text-blue-500 disabled:border-blue-500 text-white focus:outline-blue-300 active:bg-blue-500': props.type === ButtonTypes.PRIMARY,
      'bg-zinc-800 border-1 border-zinc-500 hover:bg-zinc-700 disabled:bg-zinc-700 disabled:text-zinc-500 focus:outline-zinc-300 active:bg-zinc-700': props.type === ButtonTypes.SECONDARY,
    }"
    :type="props.submit ? 'submit' : 'button'"
  >
    <slot/>
  </button>
</template>

<script setup lang="ts">
import { ButtonTypes } from '@/types/enums';

const props = defineProps<{
  submit?: boolean,
  type: ButtonTypes,
  disabled?: boolean,
}>();
</script>
