import Vue from 'vue'
import Router from 'vue-router'
import Index from '@/components/Index'
import Train from '@/components/Train'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Index',
      component: Index
    }, {
      path: '/train/:id',
      name: 'Train',
      component: Train
    }
  ]
})
