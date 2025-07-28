<script setup lang="ts">
import { ref } from 'vue'

const items = ref(Array.from({ length: 100000 }).map((_, i) => `Item #${i}`))
</script>

<template>
  <div class="w-full h-screen">
    <!--  Chat room header -->
    <div class="w-full flex items-center px-2" style="height: 32px">
      <div>Room title</div>
    </div>
    <!--  Message scroller -->
    <div>
      <div class="flex justify-center">
        <VirtualScroller
          :items="items"
          :itemSize="50"
          class="border border-surface-200 dark:border-surface-700 rounded"
          style="width: 100%; height: calc(100vh - 32px)"
        >
          <template v-slot:item="{ item, options }">
            <div
              :class="[
                'flex items-center p-2',
                { 'bg-surface-100 dark:bg-surface-700': options.odd },
              ]"
              style="height: 50px"
            >
              {{ item }}
            </div>
          </template>
        </VirtualScroller>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
