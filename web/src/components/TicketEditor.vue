<template>
  <div class="flex justify-center items-center absolute z-20 w-screen h-screen">
    <div class="absolute w-full h-full bg-black/40"></div>
    <div
      class="relative flex flex-col gap-4 z-30 mb-8 p-4 w-124 rounded-md bg-zinc-800 border-1 border-zinc-700 shadow-sm shadow-zinc-900"
      v-on-click-outside="() => emit('exit')"
    >
      <div class="flex justify-between items-center">
        <div
          class="flex gap-1 border-1 border-red-400 bg-red-500/80 p-1 rounded-md text-xs w-fit"
          v-if="props.type === EditorTypes.FEATURE"
        >
          <ClipboardDocumentListIcon class="h-4 aspect-square" />
          Feature
        </div>

        <div
          class="flex gap-1 border-1 border-blue-400 bg-blue-500/80 p-1 rounded-md text-xs w-fit"
          v-else
        >
          <DocumentPlusIcon class="h-4 aspect-square" />
          Requirement
        </div>

        <button
          class="group hover:cursor-pointer"
          @click="emit('exit')"
        >
          <XMarkIcon class="h-6 aspect-square fill-zinc-400 group-hover:fill-zinc-500" />
        </button>
      </div>
      <form
        class="flex flex-col gap-4"
        @submit.prevent=""
      >
        <FormTextInput
          :type="TextInputTypes.TEXT"
          @input="(v) => form.title.value = v"
          :error="form.title.error"
          :required="true"
          placeholder="Title"
          :class="props.type === EditorTypes.REQUIREMENT ? 'col-span-3' : 'col-span-4'"
        />
        <div class="grid grid-cols-6 gap-4">
          <div
            class="relative col-span-4 gap-4"
            v-on-click-outside="() => assigneeDropdownVisible = false"
          >
            <button
              class="flex items-center justify-between w-full border-zinc-500 gap-3 bg-zinc-700 p-2 text-sm hover:bg-zinc-800 hover:cursor-pointer"
              :class="assigneeDropdownVisible ? 'border-x-1 border-t-1 rounded-t-md' : 'border-1 rounded-md'"
              @click="assigneeDropdownVisible = !assigneeDropdownVisible"
            >
              <div
                class="flex items-center gap-3"
                v-if="currentDropdownUser"
              >
                <div class="hover:cursor-pointer bg-blue-400 font-bold w-5 h-5 text-xs flex justify-center items-center rounded-full">
                  {{ getInitials(currentDropdownUser) }}
                </div>
                {{ currentDropdownUser?.first_name }} {{ currentDropdownUser?.last_name}}
              </div>
              <div
                class="flex items-center gap-3"
                v-else
              >
                <UserCircleIcon
                  class="h-5"
                />
                Unassigned
              </div>
              <ChevronUpIcon
                v-if="assigneeDropdownVisible"
                class="h-4"
              />
              <ChevronDownIcon
                v-else
                class="h-4"
              />
            </button>
            <div
              class="z-10 w-full absolute"
              v-if="assigneeDropdownVisible"
            >
              <button
                v-if="currentDropdownUser"
                class="flex items-center gap-3 w-full p-2 text-sm bg-zinc-700 border-x-1 border-t-1 last:border-b-1 last:rounded-b-md border-zinc-500 hover:cursor-pointer hover:bg-zinc-800"
                @click="setAssignee(null)"
              >
                <UserCircleIcon
                  class="h-5"
                />
                Unassigned
              </button>
              <button
                v-if="assigneeDropdownVisible"
                v-for="user, i in dropdownUsers"
                :key="i"
                class="flex items-center gap-3 w-full p-2 text-sm bg-zinc-700 border-x-1 border-t-1 last:border-b-1 last:rounded-b-md border-zinc-500 hover:cursor-pointer hover:bg-zinc-800"
                @click="setAssignee(user)"
              >
                <div class="hover:cursor-pointer bg-blue-400 font-bold w-5 h-5 text-xs flex justify-center items-center rounded-full">
                  {{ getInitials(user) }}
                </div>
                <span>{{ user.first_name }} {{ user.last_name }}</span>
              </button>
            </div>
          </div>
          <div
            class="relative col-span-2"
            v-on-click-outside="() => stateDropdownVisible = false"
          >
            <button
              class="flex items-center justify-between w-full border-zinc-500 gap-3 bg-zinc-700 p-2 text-sm hover:bg-zinc-800 hover:cursor-pointer"
              :class="stateDropdownVisible ? 'border-x-1 border-t-1 rounded-t-md' : 'border-1 rounded-md'"
              @click="stateDropdownVisible = !stateDropdownVisible"
            >
              {{ translateState(currentDropdownState) }}
              <ChevronUpIcon
                v-if="stateDropdownVisible"
                class="h-4"
              />
              <ChevronDownIcon
                v-else
                class="h-4"
              />
            </button>
            <div
              class="z-10 w-full absolute"
              v-if="stateDropdownVisible"
            >
              <button
                v-for="state, i in dropdownStates"
                :key="i"
                class="flex items-center gap-3 w-full p-2 text-sm bg-zinc-700 border-x-1 border-t-1 last:border-b-1 last:rounded-b-md border-zinc-500 hover:cursor-pointer hover:bg-zinc-800"
                @click="setState(state)"
              >
                {{ translateState(state) }}
              </button>
            </div>
          </div>
        </div>
        <div
          class="grid grid-cols-6 gap-4"
          v-if="props.type === EditorTypes.REQUIREMENT"
        >
          <div
            class="relative col-span-4 gap-4"
            v-on-click-outside="() => featureDropdownVisible = false"
          >
            <button
              class="flex items-center justify-between w-full border-zinc-500 gap-3 bg-zinc-700 p-2 text-sm hover:bg-zinc-800 hover:cursor-pointer"
              :class="featureDropdownVisible ? 'border-x-1 border-t-1 rounded-t-md' : 'border-1 rounded-md'"
              @click="featureDropdownVisible = !featureDropdownVisible"
            >
              <div
                class="flex items-center gap-3"
                v-if="currentDropdownFeature"
              >
                <ClipboardDocumentListIcon
                  class="h-5 aspect-square text-red-400"
                />
                {{ currentDropdownFeature.title }}
              </div>
              <div
                class="flex items-center gap-3"
                v-else
              >
                <ClipboardDocumentListIcon
                  class="h-5 aspect-square"
                />
                Unassigned
              </div>
              <ChevronUpIcon
                v-if="featureDropdownVisible"
                class="h-4"
              />
              <ChevronDownIcon
                v-else
                class="h-4"
              />
            </button>
            <div
              class="z-5 w-full absolute overflow-scroll h-32 rounded-b-md border-b-1 border-zinc-500"
              v-if="featureDropdownVisible"
            >
              <button
                v-if="currentDropdownFeature"
                class="flex items-center gap-3 w-full p-2 text-sm bg-zinc-700 border-x-1 border-t-1 last:border-b-1 last:rounded-b-md border-zinc-500 hover:cursor-pointer hover:bg-zinc-800"
                @click="setFeature(null)"
              >
                <ClipboardDocumentListIcon
                  class="h-5 aspect-square"
                />
                Unassigned
              </button>
              <button
                v-if="featureDropdownVisible"
                v-for="feature, i in dropdownFeatures"
                :key="i"
                class="flex items-center gap-3 w-full p-2 text-sm bg-zinc-700 border-x-1 border-t-1 last:rounded-b-md border-zinc-500 hover:cursor-pointer hover:bg-zinc-800"
                :title="feature.id"
                @click="setFeature(feature)"
              >
                <ClipboardDocumentListIcon
                  class="h-5 aspect-square text-red-400"
                />
                {{ feature.title }}
              </button>
            </div>
          </div>
          <div class="col-span-2">
            <FormTextInput
              :type="TextInputTypes.NUMERICAL"
              @input="(v) => form.estimate.value = v as number"
              placeholder="Estimate"
            />
          </div>
        </div>
        <div class="max-h-72 overflow-scroll">
          <InkMde
            v-model="form.description.value"
            class="bg-zinc-700 rounded-md"
            style="-ink-internal-syntax-processing-instruction-color: rgb(255, 0, 0);"
            :options="{
              interface: {
                toolbar: true,
              }
            }"
          />
        </div>

        <div class="flex gap-4 ml-auto w-2/3">
          <FormButton
            class="w-full"
            :type="ButtonTypes.SECONDARY"
            @click="() => emit('exit')"
          >Cancel</FormButton>

          <FormButton
            :type="ButtonTypes.PRIMARY"
            class="w-full"
            @click="props.type === EditorTypes.FEATURE ? createFeature() : createRequirement()"
          >Save</FormButton>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Feature, FormValue, Requirement, Response, User } from '@/types/interfaces';
import { vOnClickOutside } from '@vueuse/components';
import FormTextInput from './FormTextInput.vue';
import { ButtonTypes, EditorTypes, States, TextInputTypes } from '@/types/enums';
import { computed, onMounted, ref, type ComputedRef, type Ref } from 'vue';
import axios from 'axios';
import { ChevronDownIcon, ChevronUpIcon, UserCircleIcon, XMarkIcon } from '@heroicons/vue/20/solid';
import { ClipboardDocumentListIcon, DocumentPlusIcon } from '@heroicons/vue/24/outline';
import InkMde from 'ink-mde/vue';
import FormButton from './FormButton.vue';

const assigneeDropdownVisible: Ref<boolean> = ref(false);
const assigneeDropdownUsers: Ref<User[]> = ref([]);

const stateDropdownVisible: Ref<boolean> = ref(false);
const stateDropdownStates: Ref<States[]> = ref([
  States.PROPOSED,
  States.IN_PROGRESS,
  States.IN_TEST,
  States.DONE,
]);

const featureDropdownVisible: Ref<boolean> = ref(false);
const featureDropdownFeatures: Ref<Feature[]> = ref([]);

interface Form {
  user_id: FormValue<string>,
  feature_id: FormValue<string>,
  sprint_id: FormValue<string>,
  state: FormValue<States>,
  title: FormValue<string>,
  estimate: FormValue<number>,
  description: FormValue<string>,
}

const form: Ref<Form> = ref({
  user_id: { value: '', error: null },
  feature_id: { value: '', error: null },
  sprint_id: { value: '', error: null },
  state: { value: States.PROPOSED, error: null },
  title: { value: '', error: null },
  estimate: { value: 1, error: null },
  description: { value: '', error: null },
});

onMounted(async () => {
  axios.get<Response<User[]>>('/users')
    .then((res) => assigneeDropdownUsers.value = res.data.data);

  if (props.type === EditorTypes.REQUIREMENT) {
    axios.get<Response<Feature[]>>('/features')
      .then((res) => {
        console.log(res.data.data);
        featureDropdownFeatures.value = res.data.data.sort(x => new Date(x.created_at).getMilliseconds());
      });
  }
})

const emit = defineEmits<{
  (e: 'exit' ): void,
  (e: 'finish', value: Feature | Requirement): void,
}>();

const props = defineProps<{
  type: EditorTypes,
}>();

const currentDropdownUser: ComputedRef<User | null> = computed(() => {
  return assigneeDropdownUsers.value.find(x => x.id === form.value.user_id.value) ?? null;
})

const currentDropdownState: ComputedRef<States> = computed(() => {
  return stateDropdownStates.value.find(x => x === form.value.state.value) ?? States.PROPOSED;
});

const currentDropdownFeature: ComputedRef<Feature | null> = computed(() => {
  return featureDropdownFeatures.value.find(x => x.id === form.value.feature_id.value) ?? null;
});

const dropdownUsers: ComputedRef<User[]> = computed(() => {
  return assigneeDropdownUsers.value.filter(x => x.id !== form.value.user_id.value);
});

const dropdownStates: ComputedRef<States[]> = computed(() => {
  return stateDropdownStates.value.filter(x => x !== form.value.state.value);
});

const dropdownFeatures: ComputedRef<Feature[]> = computed(() => {
  return featureDropdownFeatures.value.filter(x => x.id !== form.value.feature_id.value);
});

function getInitials(user: User): string {
  const split_last_name: string[] = user.last_name.split(' ');
  const last_name: string = split_last_name.length > 1 ? split_last_name[split_last_name.length - 1] : split_last_name[0]

  return user.first_name.substring(0, 1) + last_name.substring(0, 1);
}

function setAssignee(assignee: User | null): void {
  form.value.user_id.value = assignee?.id ?? '';
  assigneeDropdownVisible.value = false;
}

function setState(state: States): void {
  form.value.state.value = state;
  stateDropdownVisible.value = false;
}

function setFeature(feature: Feature | null): void {
  form.value.feature_id.value = feature?.id ?? '';
  featureDropdownVisible.value = false;
}

function translateState(state: States): string {
  switch(state) {
    case States.PROPOSED:
      return 'Proposed';
    case States.IN_PROGRESS:
      return 'In progress';
    case States.IN_TEST:
      return 'In test';
    case States.DONE:
      return 'Done';
  }
}

function createFeature(): void {
  form.value.user_id.error = null;
  form.value.sprint_id.error = null;
  form.value.state.error = null;
  form.value.title.error = null;
  form.value.description.error = null;

  axios.post<Response<Feature>>('/features', {
    user_id: form.value.user_id.value,
    sprint_id: form.value.sprint_id.value,
    state: form.value.state.value,
    title: form.value.title.value,
    description: form.value.description.value,
  })
    .then((res) => {
      emit('finish', res.data.data);
    })
    .catch((err) => {
      const message: string = err.response.data.message;

      if (message.includes('title')) {
        form.value.title.error = message;
        return
      }

      if (message.includes('user_id')) {
        form.value.user_id.error = message;
      }

      if (message.includes('sprint_id')) {
        form.value.sprint_id.error = message;
      }

      if (message.includes('state')) {
        form.value.state.error = message;
        return;
      }

      if (message.includes('description')) {
        form.value.description.error = message;
        return;
      }

      if (message.includes('resource not found')) {
        form.value.user_id.error = message;
        form.value.sprint_id.error = message;
        return
      }
    });
}

function createRequirement(): void {
  form.value.user_id.error = null;
  form.value.feature_id.error = null;
  form.value.sprint_id.error = null;
  form.value.state.error = null;
  form.value.title.error = null;
  form.value.description.error = null;

  axios.post<Response<Requirement>>('/requirements', {
    user_id: form.value.user_id.value,
    sprint_id: form.value.sprint_id.value,
    feature_id: form.value.feature_id.value,
    state: form.value.state.value,
    title: form.value.title.value,
    description: form.value.description.value,
  })
    .then((res) => {
      emit('finish', res.data.data);
    })
    .catch((err) => {
      const message: string = err.response.data.message;

      if (message.includes('title')) {
        form.value.title.error = message;
        return
      }

      if (message.includes('user_id')) {
        form.value.user_id.error = message;
      }

      if (message.includes('sprint_id')) {
        form.value.sprint_id.error = message;
      }

      if (message.includes('feature_id')) {
        form.value.feature_id.error = message;
      }

      if (message.includes('state')) {
        form.value.state.error = message;
        return;
      }

      if (message.includes('description')) {
        form.value.description.error = message;
        return;
      }

      if (message.includes('resource not found')) {
        form.value.user_id.error = message;
        form.value.sprint_id.error = message;
        form.value.feature_id.error = message;
        return
      }
    });
}
</script>

<style>
:root {
  --ink-syntax-processing-instruction-color: var(--color-zinc-500);
  --ink-block-background-color: var(--color-zinc-900);
  --ink-border-radius: var(--radius-md);
}
</style>
