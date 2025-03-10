<template>
  <div class="flex w-full">
    <TicketEditor
      v-if="featureEditorVisible"
      :ticket="featureBeingEdited"
      :type="EditorTypes.FEATURE"
      :update="editorIsUpdating"
      @exit="editorExit"
      @finish="onEditorFinish"
    />
    <TicketEditor
      v-if="requirementEditorVisible"
      :ticket="requirementBeingEdited"
      :type="EditorTypes.REQUIREMENT"
      :update="editorIsUpdating"
      @exit="editorExit"
      @finish="onEditorFinish"
    />

    <div class="w-15 p-3"></div>
    <VerticalNavigationBar />
    <main class="flex flex-col w-full items-center gap-8 mt-8">
      <div class="flex flex-col gap-8 w-96 sm:w-130 md:w-160 lg:w-220 2xl:w-333">
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
          <div class="relative w-81 sm:w-115 md:w-145 lg:w-205 2xl:w-318 text-left table-fixed">
            <div class="grid grid-cols-20 font-medium text-xs pb-4 border-b-1 border-zinc-700 text-zinc-500">
              <div></div>
              <p class="col-span-1">Id</p>
              <p class="col-span-4">Assignee</p>
              <p class="col-span-2">State</p>
              <p class="col-span-9">Title</p>
              <p class="col-span-1">Estimate</p>
              <p class="col-span-2 text-right pr-2">Actions</p>
            </div>
            <div
              class="group"
              v-for="feature, i in features"
              :key="i"
            >
              <DraggableTableEntry
                :id="i"
                :draggable="true"
                :sibling-coords="draggableCoordinates"
                :sibling-being-dragged="draggableBeingDragged"
                class="group-even:bg-zinc-900/30 w-full"
                @coordinates="(v) => draggableCoordinates = v"
                @dragging="(v) => handleDragging(v, feature.id)"
              >
                <div class="pl-2">
                  <FormButton
                    v-if="requirements.filter(x => x.feature_id === feature.id).length"
                    :type="ButtonTypes.SECONDARY"
                    @click="featureCollapseToggles.find(x => x.id === feature.id)!.enabled = !featureCollapseToggles.find(x => x.id === feature.id)!.enabled"
                    @mousedown.stop=""
                    @mouseup.stop=""
                  >
                    <ChevronUpIcon
                      v-if="featureCollapseToggles.find(x => x.id === feature.id)?.enabled"
                      class="h-4 aspect-square"
                    />
                    <ChevronDownIcon
                      v-else
                      class="h-4 aspect-square"
                    />
                  </FormButton>
                </div>
                <div class="flex gap-1 text-xs col-span-1 text-zinc-500 my-auto">
                  <ClipboardDocumentListIcon
                    class="h-5 aspect-square text-red-400"
                  />
                  {{ feature.external_id }}
                </div>
                <div class="col-span-4">
                  <div
                    class="flex gap-2 items-center"
                    v-if="getUserWithId(feature.user_id)"
                  >
                    <div class="hover:cursor-pointer bg-blue-400 font-bold w-8 h-8 flex justify-center items-center rounded-full">
                      {{ getInitials(getUserWithId(feature.user_id)) }}
                    </div>
                    <span class="overflow-hidden text-ellipsis line-clamp-1">{{ getUserWithId(feature.user_id)?.first_name }} {{ getUserWithId(feature.user_id)?.last_name }}</span>
                  </div>
                  <div
                    class="flex gap-2 items-center"
                    v-else
                  >
                    <UserCircleIcon
                      class="h-8"
                    />
                    Unassigned
                  </div>
                </div>
                <div
                  class="col-span-2"
                  :class="{
                    'text-zinc-500': feature.state === States.PROPOSED,
                    'text-yellow-500': feature.state === States.IN_PROGRESS,
                    'text-cyan-500': feature.state === States.IN_TEST,
                    'text-lime-500': feature.state === States.DONE,
                  }"
                >{{ translateState(feature.state) }}</div>
                <div class="col-span-10 overflow-hidden text-ellipsis line-clamp-1">{{ feature.title }}</div>
                <div class="col-span-2 pr-2 flex items-center justify-end gap-2">
                  <FormButton
                    :type="ButtonTypes.SECONDARY"
                    @click.stop="newRequirementForFeature(feature)"
                    @mousedown.stop=""
                    @mouseup.stop=""
                  >
                    <DocumentPlusIcon
                      class="h-4"
                    />
                  </FormButton>
                  <FormButton
                    :type="ButtonTypes.SECONDARY"
                    @click.stop="addToEditor(EditorTypes.FEATURE, feature)"
                    @mousedown.stop=""
                    @mouseup.stop=""
                  >
                    <PencilIcon
                      class="h-4"
                    />
                  </FormButton>
                  <FormButton
                    :type="ButtonTypes.SECONDARY"
                    @mousedown.stop=""
                    @mouseup.stop=""
                  >
                    <TrashIcon
                      class="h-4"
                    />
                  </FormButton>
                </div>
              </DraggableTableEntry>
              <DraggableTableEntry
                v-if="featureCollapseToggles.find(x => x.id === feature.id)?.enabled"
                v-for="requirement, i in requirements.filter(x => x.feature_id === feature.id)"
                :key="i"
                :id="i"
                :draggable="false"
                :sibling-coords="draggableCoordinates"
                :sibling-being-dragged="draggableBeingDragged"
                @coordinates="(v) => draggableCoordinates = v"
                @dragging="(v) => handleDragging(v, requirement.id)"
                class="even:bg-zinc-900/70 odd:bg-zinc-900/90 w-full"
              >
                <div></div>
                <div class="flex gap-1 text-xs col-span-1 text-zinc-500 my-auto">
                  <DocumentPlusIcon
                    class="h-5 aspect-square text-blue-400"
                  />
                  {{ requirement.external_id }}
                </div>
                <div class="col-span-4">
                  <div
                    class="flex gap-2 items-center"
                    v-if="getUserWithId(requirement.user_id)"
                  >
                    <div class="hover:cursor-pointer bg-blue-400 font-bold w-8 h-8 flex justify-center items-center rounded-full">
                      {{ getInitials(getUserWithId(requirement.user_id)) }}
                    </div>
                    <span class="overflow-hidden text-ellipsis line-clamp-1">{{ getUserWithId(requirement.user_id)?.first_name }} {{ getUserWithId(requirement.user_id)?.last_name }}</span>
                  </div>
                  <div
                    class="flex gap-2 items-center"
                    v-else
                  >
                    <UserCircleIcon
                      class="h-8"
                    />
                    Unassigned
                  </div>
                </div>
                <div
                  class="col-span-2"
                  :class="{
                    'text-zinc-500': requirement.state === States.PROPOSED,
                    'text-yellow-500': requirement.state === States.IN_PROGRESS,
                    'text-cyan-500': requirement.state === States.IN_TEST,
                    'text-lime-500': requirement.state === States.DONE,
                  }"
                >{{ translateState(requirement.state) }}</div>
                <div class="col-span-9 overflow-hidden text-ellipsis line-clamp-1">{{ requirement.title }}</div>
                <div class="col-span-1">
                  <span v-if="requirement.estimate">
                    {{ requirement.estimate }}
                  </span>
                  <span v-else>
                  </span>
                </div>
                <div class="col-span-2 pr-2 flex items-center justify-end gap-2">
                  <FormButton
                    :type="ButtonTypes.SECONDARY"
                    @click.stop="addToEditor(EditorTypes.REQUIREMENT, requirement)"
                    @mousedown.stop=""
                    @mouseup.stop=""
                  >
                    <PencilIcon
                      class="h-4"
                    />
                  </FormButton>
                  <FormButton
                    :type="ButtonTypes.SECONDARY"
                    @mousedown.stop=""
                    @mouseup.stop=""
                  >
                    <TrashIcon
                      class="h-4"
                    />
                  </FormButton>
                </div>
              </DraggableTableEntry>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import DraggableTableEntry from '@/components/DraggableTableEntry.vue';
import FormButton from '@/components/FormButton.vue';
import TicketEditor from '@/components/TicketEditor.vue';
import VerticalNavigationBar from '@/components/VerticalNavigationBar.vue';
import router from '@/router';
import { ButtonTypes, EditorTypes, States } from '@/types/enums';
import type { Feature, Requirement, Response, User } from '@/types/interfaces';
import { ChevronDownIcon, ChevronUpIcon, UserCircleIcon } from '@heroicons/vue/20/solid';
import { ClipboardDocumentListIcon, DocumentPlusIcon, PencilIcon, TrashIcon } from '@heroicons/vue/24/outline';
import { vOnClickOutside } from '@vueuse/components';
import axios from 'axios';
import type _default from 'ink-mde/vue';
import { onMounted, ref, watch, type Ref } from 'vue';
import { useRoute } from 'vue-router';

const createDropdownVisible: Ref<boolean> = ref(false);
const featureEditorVisible: Ref<boolean> = ref(false);
const requirementEditorVisible: Ref<boolean> = ref(false);

const featureBeingEdited: Ref<Feature | null> = ref(null);
const requirementBeingEdited: Ref<Requirement | null> = ref(null);
const editorIsUpdating: Ref<boolean> = ref(false);

const features: Ref<Feature[]> = ref([]);
const requirements: Ref<Requirement[]> = ref([]);
const users: Ref<User[]> = ref([]);

const draggableCoordinates: Ref<number> = ref(0);
const draggableBeingDragged: Ref<number> = ref(-1);

const featureCollapseToggles: Ref<{ id: string, enabled: boolean }[]> = ref(
  JSON.parse(window.localStorage.getItem('backlog_feature_collapse_toggles') ?? '[]') ?? []
);

const route = useRoute();
onMounted(async () => {
  refreshData();

  if (route.query.ticketId && route.query.ticketType) {
    addToEditorById(route.query.ticketType as EditorTypes, route.query.ticketId as string);
  }
});

watch(featureCollapseToggles, (to) => {
  const serialized: string = JSON.stringify(to);
  window.localStorage.setItem('backlog_feature_collapse_toggles', serialized);
}, { deep: true});

async function refreshData(): Promise<void> {
  const featuresResponse = await axios.get<Response<Feature[]>>('/features');
  features.value = featuresResponse.data.data.sort((a, b) => a.position - b.position);

  const requirementsResponse = await axios.get<Response<Requirement[]>>('/requirements');
  // Temporary sorting solution. Should make use of position (or derivative) once this has been implemented at the requirement level.
  requirements.value = requirementsResponse.data.data.sort((a, b) => new Date(a.created_at).getMilliseconds() - new Date(b.created_at).getMilliseconds());

  const usersResponse = await axios.get<Response<User[]>>('/users');
  users.value = usersResponse.data.data;

  const newCollapseToggles = features.value.map(x => ({ id: x.id, enabled: false }));
  newCollapseToggles.forEach((x) => {
    if (featureCollapseToggles.value.includes(x)) {
      return;
    }

    featureCollapseToggles.value.push(x);
  });

  // TODO: Delete collapseToggles for feaetures no longer present in the backlog
  featureCollapseToggles.value.forEach((x) => {
    if (newCollapseToggles.includes(x)) {
      return;
    }

    const index: number = featureCollapseToggles.value.findIndex(y => x.id === y.id);
    if (index === -1) return;

    featureCollapseToggles.value = featureCollapseToggles.value.splice(index);
  });
}

function onEditorFinish(): void {
  editorExit()
  refreshData();
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

function getInitials(user: User | null): string {
  if (user === null) return '';

  const split_last_name: string[] = user.last_name.split(' ');
  const last_name: string = split_last_name.length > 1 ? split_last_name[split_last_name.length - 1] : split_last_name[0]

  return user.first_name.substring(0, 1) + last_name.substring(0, 1);
}

function getUserWithId(id: string):  User | null {
  return users.value.find(x => x.id === id) ?? null;
}

function addToEditor(type: EditorTypes, ticket: Requirement | Feature | null): void {
  if (ticket === null) {
    featureEditorVisible.value = false;
    requirementEditorVisible.value = false;
    editorIsUpdating.value = false;
  }
  
  if (type === EditorTypes.FEATURE) {
    featureBeingEdited.value = ticket as Feature;
    featureEditorVisible.value = true;
    editorIsUpdating.value = true;
    return
  }
  
  if (type === EditorTypes.REQUIREMENT) {
    requirementBeingEdited.value = ticket as Requirement;
    requirementEditorVisible.value = true;
    editorIsUpdating.value = true;
    return
  }
}

function addToEditorById(type: EditorTypes, id: string): void {
  if (type === EditorTypes.FEATURE) {
    axios.get<Response<Feature>>(`/features/${id}`)
      .then(res => {
        featureBeingEdited.value = res.data.data
        featureEditorVisible.value = true;
        editorIsUpdating.value = true;
      });
    return;
  }

  if (type === EditorTypes.REQUIREMENT) {
    axios.get<Response<Requirement>>(`/requirements/${id}`)
      .then(res => {
        requirementBeingEdited.value = res.data.data
        requirementEditorVisible.value = true;
        editorIsUpdating.value = true;
      });
    return;
  }
}

function editorExit(): void {
  featureEditorVisible.value = false;
  featureBeingEdited.value = null;
  requirementEditorVisible.value = false;
  requirementBeingEdited.value = null;
  editorIsUpdating.value = false;
  router.replace({ query: {} });
}

function handleDragging(dragging: number, featureId: string): void {
  draggableBeingDragged.value = dragging;
  disableCollapseIfEnabled(dragging, featureId);
}

function disableCollapseIfEnabled(dragging: number, featureId: string): void {
  if (dragging === -1) return;

  const toggle = featureCollapseToggles.value.find(x => x.id === featureId);
  if (typeof toggle === 'undefined') return;

  toggle.enabled = false;
}

function newRequirementForFeature(feature: Feature): void {
  const requirement: Requirement = {
    id: '',
    external_id: '',
    user_id: '',
    feature_id: feature.id,
    sprint_id: '',
    state: States.PROPOSED,
    title: '',
    estimate: 0,
    description: '',
    position: 0,
    created_at: '',
  };

  requirementBeingEdited.value = requirement;
  requirementEditorVisible.value = true;
}
</script>
