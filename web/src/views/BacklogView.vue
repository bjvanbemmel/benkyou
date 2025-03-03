<template>
  <div class="flex w-full">
    <TicketEditor
      v-if="featureEditorVisible"
      :type="EditorTypes.FEATURE"
      @exit="featureEditorVisible = false"
      @finish="featureEditorVisible = false"
    />
    <TicketEditor
      v-if="requirementEditorVisible"
      :type="EditorTypes.REQUIREMENT"
      @exit="requirementEditorVisible = false"
      @finish="requirementEditorVisible = false"
    />

    <div class="w-15"></div>
    <VerticalNavigationBar />
    <main class="flex flex-col gap-8 w-full mx-8 mt-8">
      <div class="flex items-center justify-between">
        <h1 class="text-xl font-bold">Backlog</h1>
        <div class="relative w-max">
          <button
            class="flex items-center justify-center w-36 gap-2 bg-zinc-800 border-zinc-700 border-1 rounded-md p-2 hover:bg-zinc-900 hover:cursor-pointer"
            @click="createDropdownVisible = !createDropdownVisible"
            v-on-click-outside="() => createDropdownVisible = false"
          >
            Create new
            <ChevronUpIcon
              v-if="createDropdownVisible"
              class="h-4"
            />
            <ChevronDownIcon
              v-else
              class="h-4"
            />
          </button>
          <div
            v-if="createDropdownVisible"
            class="w-full flex flex-col w-full absolute bg-zinc-800 border-1 border-zinc-700 top-10 left-0 rounded-md shadow-sm shadow-zinc-900"
          >
            <button
              class="flex items-center gap-2 text-white w-full hover:cursor-pointer hover:bg-zinc-900 rounded-md px-2 py-1"
              @click="() => featureEditorVisible = true"
            >
              <ClipboardDocumentListIcon class="h-4 aspect-square text-red-400" />
              Feature
            </button>
            <button
              class="flex items-center gap-2 text-white w-full border-t-1 border-zinc-700 hover:cursor-pointer hover:bg-zinc-900 rounded-md px-2 py-1"
              @click="() => requirementEditorVisible = true"
            >
              <DocumentPlusIcon class="h-4 aspect-square text-blue-400" />
              Requirement
            </button>
          </div>
        </div>
      </div>
      <div class="bg-zinc-800 border-1 border-zinc-700 p-8 w-full h-full rounded-md">
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import TicketEditor from '@/components/TicketEditor.vue';
import VerticalNavigationBar from '@/components/VerticalNavigationBar.vue';
import { EditorTypes } from '@/types/enums';
import { ChevronDownIcon, ChevronUpIcon } from '@heroicons/vue/20/solid';
import { ClipboardDocumentListIcon, DocumentPlusIcon } from '@heroicons/vue/24/outline';
import { vOnClickOutside } from '@vueuse/components';
import { ref, type Ref } from 'vue';

const createDropdownVisible: Ref<boolean> = ref(false);
const featureEditorVisible: Ref<boolean> = ref(false);
const requirementEditorVisible: Ref<boolean> = ref(true);

</script>
