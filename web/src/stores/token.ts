import axios from "axios";
import { defineStore } from "pinia";
import { ref, type Ref } from "vue";

export const useTokenStore = defineStore('token', () => {
  const token: Ref<string | null> = ref(null);

  function set(value: string): void {
    localStorage.setItem('token', value);
    token.value = value;
    updateAxios();
  }

  function updateAxios(): void {
    axios.defaults.headers.common['Authorization'] = `Bearer ${token.value}`;
  }

  function clear(): void {
    localStorage.removeItem('token');
    token.value = null;
  }

  function get(): string | null {
    return localStorage.getItem('token');
  }

  return { token, set, clear, get, updateAxios };
});
