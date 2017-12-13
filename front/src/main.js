// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import Vuex from 'vuex'
import VueAxios from 'vue-axios'
import axios from 'axios'

import App from './App'
import router from './router'

Vue.config.productionTip = false

Vue.use(Vuex)
Vue.use(VueAxios, axios)

Vue.axios.defaults.baseURL = 'http://localhost:8880'

const store = new Vuex.Store({
  state: {
    username: null
  },
  mutations: {
    connect (state, username) {
      state.username = username
    },
    disconnect (state) {
      state.username = null
    }
  },
  getters: {
    username: state => state.username
  }
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  template: '<App/>',
  components: { App }
})
