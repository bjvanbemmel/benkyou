<script setup lang="ts">
  import { CirclePlusFilled } from '@element-plus/icons-vue'
  import type { Ref } from 'vue'

  const search: Ref<string> = ref('')

  const workspaces: Array<{ id: string, title: string }> = [
    {
      id: '0',
      title: 'Test 1',
    },
    {
      id: '1',
      title: 'Test 2',
    },
    {
      id: '2',
      title: 'Patat',
    },
    {
      id: '3',
      title: 'Frikandel',
    },
    {
      id: '4',
      title: 'Schnitzel',
    },
    {
      id: '5',
      title: 'Varkenshaas',
    },
    {
      id: '6',
      title: 'RYSST',
    },
    {
      id: '7',
      title: 'Beer',
    },
  ]

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const searchWorkspaces = (query: string, cb: any): void => {
    const results = query
      ? workspaces.filter(x => x.title.toLowerCase().includes(query.toLowerCase()))
      : workspaces

      cb(results.map(x => ({ value: x.title })))
  }
</script>

<template>
  <NuxtLayout name="authenticated">
    <div
      class="mx-auto w-fit"
    >
      <div
        class="flex flex-col sm:gap-0 gap-6 sm:flex-row sm:justify-between"
      >
        <el-badge
          :value="workspaces.length"
          type="primary"
          :offset="[10, 5]"
        >
          <h1
            class="text-3xl"
          >
            Workspaces
          </h1>
        </el-badge>

        <el-autocomplete
          v-model="search"
          :fetch-suggestions="searchWorkspaces"
          :trigger-on-focus="false"
          clearable
          placeholder="Search for workspace..."
          class="w-full sm:max-w-64"
        />
      </div>
      <el-divider />
      <div
        class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 gap-4"
      >
        <button
          v-for="workspace in workspaces"
          :key="workspace.id"
          class=""
          >
          <el-card
            class="text-left w-96 sm:w-64 h-32"
          >
            <el-tag
              class="mb-2"
              >Workspace</el-tag>
            <h1
              class="text-ellipsis overflow-hidden"
              >{{ workspace.title }}</h1>
          </el-card>
        </button>
      </div>
      <el-divider />
    </div>
  </NuxtLayout>
</template>
