<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const menuRef = ref()
const items = ref([
  {
    separator: true,
  },
  {
    label: 'View Takeout',
    icon: 'pi pi-file',
    command: () => {
      router.push({
        name: 'loader',
      })
    },
  },
  { label: 'Upload Takeout', icon: 'pi pi-cloud-upload' },
  { label: 'Settings', icon: 'pi pi-cog' },
])

const toggle = (event) => {
  menuRef.value?.toggle(event)
}

const searchText = ref('')
</script>

<template>
  <div class="p-2 flex items-center gap-2">
    <div>
      <Button
        type="button"
        icon="pi pi-ellipsis-v"
        severity="secondary"
        @click="toggle"
        aria-haspopup="true"
        aria-controls="main_menu"
      />
      <Menu ref="menuRef" id="main_menu" :model="items" :popup="true">
        <template #start>
          <span class="flex items-center gap-1 px-2 py-2">
            <span class="font-semibold">Telegram Archive</span>
          </span>
        </template>
      </Menu>
    </div>
    <InputText type="text" v-model="searchText" placeholder="Search" class="w-full" />
  </div>
</template>

<style scoped></style>
