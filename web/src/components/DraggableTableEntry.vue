<template>
  <div
    ref="el"
    class="flex flex-col w-full"
    :class="props.siblingBeingDragged === props.id ? '' : ''"
  >
    <div
      ref="dropArea"
      class="flex justify-center items-center border-blue-400 border-dashed rounded-md bg-blue-400/20"
      :class="shouldAllowDrop ? 'border-2' : 'border-0'"
    >
      <div
        class="origin-center transition-all duration-200"
        :class="shouldAllowDrop ? 'h-16' : 'h-0'"
      ></div>
    </div>
    <div
      class="*:flex *:items-center gap-2 overflow-hidden text-ellipsis h-12 grid grid-cols-20"
      :class="customClasses"
      :style="`top: ${verticalCoord}px`"
      @mousedown="(e) => toggleDrag(e, true)"
    >
      <slot></slot>
    </div>
    <div
      ref="dropArea"
      class="flex justify-center items-center border-blue-400 border-dashed rounded-md bg-blue-400/20"
      :class="shouldAllowDrop ? 'border-2' : 'border-0'"
    >
      <div
        class="origin-center transition-all duration-200"
        :class="shouldAllowDrop ? 'h-16' : 'h-0'"
      ></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useDraggable } from '@vueuse/core';
import { computed, ref, useTemplateRef, watch, type ComputedRef, type Ref } from 'vue';

const props = defineProps<{
  id: number,
  draggable?: boolean,
  siblingCoords?: number,
  siblingBeingDragged?: number,
}>();

const emits = defineEmits<{
  (e: 'dragging', value: number): void,
  (e: 'coordinates', value: number): void,
}>();

const el = useTemplateRef<HTMLElement>('el');
const dropArea = useTemplateRef<HTMLElement>('dropArea');
const isDragging: Ref<boolean> = ref(false);
const verticalCoord: Ref<number> = ref(0);
const mouseRelatedToBoundingBox: Ref<number> = ref(0);
const { y } = useDraggable(el, {
  disabled: !props.draggable,
});

const customClasses: ComputedRef<string> = computed(() => {
  let base: string = props.draggable ? 'transition-transform duration-200 hover:cursor-pointer' : '';

  if (isDragging.value) {
    base += ' fixed select-none bg-zinc-900/50 border-blue-400 border-1 rounded-md shadow-md shadow-zinc-900 w-81 sm:w-115 md:w-145 lg:w-205 2xl:w-318 scale-95';
    return base;
  }
  
  return base;
});

const shouldAllowDrop: ComputedRef<boolean> = computed(() => {
  if (!props.draggable) return false;
  if (props.siblingBeingDragged === -1) return false;
  if (props.siblingBeingDragged === props.id) return false;
  if (props.siblingCoords === null || typeof props.siblingCoords === 'undefined') return false;

  const boundingBoxYTop: number = el.value?.getBoundingClientRect().top ?? 0;
  const boundingBoxYBottom: number = el.value?.getBoundingClientRect().bottom ?? 0;
  
  return (props.siblingCoords) > boundingBoxYTop && props.siblingCoords < boundingBoxYBottom;
});

watch(isDragging, async (to) => {
  emits('dragging', to ? props.id : -1);

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
  emits('coordinates', verticalCoord.value);
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
  if (shouldAllowDrop.value) {
    alert('Updated position!');
  }

  toggleDrag(event, false);
}
</script>
