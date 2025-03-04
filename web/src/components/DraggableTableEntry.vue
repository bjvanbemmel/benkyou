<template>
  <div
    ref="el"
    class="*:flex *:items-center overflow-hidden text-ellipsis h-12 grid grid-cols-20"
    :class="customClasses"
    :style="`top: ${verticalCoord}px`"
    @mousedown="(e) => toggleDrag(e, true)"
  >
    <slot></slot>
  </div>
</template>

<script setup lang="ts">
import { useDraggable } from '@vueuse/core';
import { computed, ref, useTemplateRef, watch, type ComputedRef, type Ref } from 'vue';

const props = defineProps<{
  draggable?: boolean,
}>();

const emits = defineEmits<{
  (e: 'dragging', value: boolean): void,
}>();

const el = useTemplateRef<HTMLElement>('el');
const isDragging: Ref<boolean> = ref(false);
const customClasses: ComputedRef<string> = computed(() => {
  let base: string = props.draggable ? 'hover:cursor-pointer' : '';

  if (isDragging.value) {
    base += ' fixed select-none bg-zinc-900 border-blue-400 border-1 rounded-md shadow-md shadow-zinc-900 w-81 sm:w-115 md:w-145 lg:w-205 2xl:w-318';
    return base;
  }
  
  return base;
});

const verticalCoord: Ref<number> = ref(0);
const mouseRelatedToBoundingBox: Ref<number> = ref(0);
const { y } = useDraggable(el, {
  disabled: props.draggable,
});

watch(isDragging, async (to) => {
  emits('dragging', to);

  if (!to) {
    document.removeEventListener('mousemove', updateCoords);
    document.removeEventListener('mouseup', __toggleDragDisable);
    return;
  }

  document.addEventListener('mousemove', updateCoords);
  document.addEventListener('mouseup', __toggleDragDisable);
});

function updateCoords(event: MouseEvent): void {
  let newVal = y.value - (y.value - event.clientY) - mouseRelatedToBoundingBox.value;
  verticalCoord.value = newVal;
}

function toggleDrag(event: MouseEvent, enabled: boolean): void {
  if (!props.draggable) return;

  isDragging.value = enabled;
  if (!enabled) return;

  const boundingBoxY: number = el.value?.getBoundingClientRect().y ?? 0;
  mouseRelatedToBoundingBox.value = event.clientY - boundingBoxY;
  verticalCoord.value = el.value?.getBoundingClientRect().y ?? 0;
}

function __toggleDragDisable(event: MouseEvent): void {
  toggleDrag(event, false);
}
</script>
