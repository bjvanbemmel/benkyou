<template>
  <div class="h-screen">
    <div
      class="flex w-min flex-col p-3 h-full bg-zinc-800 border-r-1 border-zinc-700"
    >
      <nav
        class="flex gap-4 w-full justify-start items-center flex-col"
      >
        <div
          class="relative flex justify-center w-full"
        >
          <div>
            <button
              class="hover:cursor-pointer bg-blue-400 font-bold p-1.5 rounded-full"
              @click="userDropdownVisible = !userDropdownVisible"
            >
              {{ userInitials }}
            </button>
          </div>
          <OnClickOutside
            v-if="userDropdownVisible"
            @trigger="() => userDropdownVisible = false"
          >
            <div
              class="flex flex-col w-min absolute bg-zinc-700 border-1 border-zinc-400 top-8 left-0 rounded-md"
            >
              <div class="px-2 py-1 text-zinc-300 w-max">
                <p>{{ userStore.user?.first_name }} {{ userStore.user?.last_name }}</p>
                <p class="text-xs text-zinc-400">{{ userStore.user?.email }}</p>
              </div>
              <div class="w-full border-b-1 border-zinc-500"></div>
              <button
                class="flex items-center gap-2 text-white w-full hover:cursor-pointer hover:bg-zinc-800 rounded-md px-2 py-1"
                @click="logout"
              >
                <ArrowLeftEndOnRectangleIcon
                  class="h-5 aspect-square"
                />
                Logout
              </button>
            </div>
          </OnClickOutside>
        </div>
        <div class="w-full border-b-1 border-zinc-500"></div>
        <div class="group flex justify-center w-full">
          <RouterLink
            :to="{ name: 'backlog' }"
          >
            <RectangleStackIcon
              class="group-hover:fill-cyan-800 h-8 aspect-square fill-cyan-600"
            />
          </RouterLink>
        </div>
        <div class="group flex justify-center w-full">
          <RouterLink
            :to="{ name: 'sprints' }"
          >
            <CalendarDateRangeIcon
              class="group-hover:fill-orange-800 h-8 aspect-square fill-orange-600"
            />
          </RouterLink>
        </div>
        <div class="group flex justify-center w-full">
          <RouterLink
            :to="{ name: 'board' }"
          >
            <DocumentCheckIcon
              class="group-hover:fill-lime-800 h-8 aspect-square fill-lime-600"
            />
          </RouterLink>
        </div>
        <div class="group flex justify-center w-full">
          <RouterLink
            :to="{ name: 'braindump' }"
          >
            <ChatBubbleLeftRightIcon
              class="group-hover:fill-yellow-800 h-8 aspect-square fill-yellow-600"
            />
          </RouterLink>
        </div>
      </nav>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useUserStore } from '@/stores/user';
import { ArrowLeftEndOnRectangleIcon, CalendarDateRangeIcon, ChatBubbleLeftRightIcon, DocumentCheckIcon, RectangleStackIcon } from '@heroicons/vue/24/outline';
import { computed, ref, type ComputedRef, type Ref } from 'vue';
import { OnClickOutside } from '@vueuse/components';
import { useTokenStore } from '@/stores/token';
import router from '@/router';

const userDropdownVisible: Ref<boolean> = ref(false);

const userStore = useUserStore();
const tokenStore = useTokenStore();

const userInitials: ComputedRef<string> = computed(() => {
  const user = userStore.user;
  if (user === null) return '';
  const split_last_name: string[] = user.last_name.split(' ');
  const last_name: string = split_last_name.length > 1 ? split_last_name[split_last_name.length - 1] : split_last_name[0]

  return user.first_name.substring(0, 1) + last_name.substring(0, 1);
});

function logout(): void {
  tokenStore.clear();
  router.push({ name: 'login' });
}
</script>
