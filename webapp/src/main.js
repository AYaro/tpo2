import { createApp } from 'vue'
import App from './App.vue'
import { createRouter, createWebHistory } from 'vue-router'
import Create from './components/Create.vue'
import Order from './components/Order.vue'
import Intro from './components/Intro.vue'

let routerHistory = createWebHistory()

let router = createRouter({
    history: routerHistory,
    routes: [
      {
        path: '/',
        component: Intro,
      },
      {
        path: '/order',
        component: Order,
      },
      {
        path: '/order/:orderId',
        component: Order,
      },
      {
        path: '/create',
        component: Create,
      }
    ]
})

createApp(App).use(router).mount('#app')
