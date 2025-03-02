<template>
  <div>
    <div class="relative">
      <input
        :id="props.id"
        :type="exposed ? TextInputTypes.TEXT : props.type"
        :placeholder="props.placeholder"
        :required="props.required"
        autocomplete="props.autocomplete"
        class="w-full focus:outline-1 focus:outline-zinc-400 p-2 bg-zinc-700 rounded-md text-sm text-zinc-100"
        :class="{
          'pr-8': props.type === TextInputTypes.PASSWORD,
          'border-1 border-red-400': props.error,
        }"
        v-model="value"
        @blur="emit('blur')"
        @input="emit('input', value)"
      />

      <button
        v-if="props.type === TextInputTypes.PASSWORD"
        type="button"
        class="absolute top-0 right-2 h-full *:h-5 *:aspect-square *:fill-zinc-300 *:active:fill-zinc-400 flex justify-center items-center hover:cursor-pointer"
        @click="() => exposed = !exposed"
      >
        <EyeIcon
          v-if="exposed"
        />
        <EyeSlashIcon
          v-else
        />
      </button>
    </div>

    <label
      v-if="error"
      :for="props.id"
      class="text-xs/1 text-red-400 mx-1"
    >
      {{ error }}
    </label>

    <label
      v-if="props.label"
      :for="props.id"
      class="text-xs/1 text-zinc-400 mx-1"
    >
      {{ props.label }}
    </label>
  </div>

</template>

<script setup lang="ts">
import { TextInputTypes } from '@/types/enums';
import { EyeSlashIcon, EyeIcon } from '@heroicons/vue/20/solid';
import { ref, type Ref } from 'vue';

const value: Ref<string> = ref('');
const exposed: Ref<boolean> = ref(false);

const emit = defineEmits<{
  (e: 'input', value: string): void,
  (e: 'blur'): void,
}>();

const props = withDefaults(defineProps<{
  id?: string,
  type: TextInputTypes,
  placeholder?: string,
  label?: string,
  required?: boolean,
  autocomplete?: string,
  error?: string | null,
}>(), {
    id: () => window.crypto.randomUUID(),
});
</script>
