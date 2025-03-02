import type { User } from "@/types/interfaces";
import { defineStore } from "pinia";
import { ref, type Ref } from "vue";

export const useUserStore = defineStore('user', () => {
  const user: Ref<User | null> = ref(null);

  function set(value: User): void {
    user.value = value;
  }

  return { user, set };
});
