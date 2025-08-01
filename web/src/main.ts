import './assets/main.css'
import 'primeicons/primeicons.css'
import 'virtual:uno.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import PrimeVue from 'primevue/config'

import App from './App.vue'
import router from './router'
import { Noir } from '@/theme.ts'

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.use(PrimeVue, {
  theme: {
    preset: Noir,
    options: {
      darkModeSelector: 'system',
    },
  },
})

app.mount('#app')
