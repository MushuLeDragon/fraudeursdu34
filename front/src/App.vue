<template>
  <div id="app">
    <h1>Les fraudeurs du 3-4</h1>
    <span v-if="$store.getters.username" class="username">
      Bonjour {{ $store.getters.username }}
      <button style="margin-left: 30px;" @click="disconnect">Déconnexion</button>
    </span>
    <login v-if="!$store.getters.username"></login>
    <router-view v-else />
  </div>
</template>

<script>
import Login from '@/components/Login'

export default {
  name: 'app',

  components: {
    Login
  },

  methods: {
    disconnect () {
      this.$store.commit('disconnect')
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #fff;
  margin-top: 60px;
}

.username {
  position: fixed;
  top: 15px;
  right: 15px;
}

a {
  color: #42b983;
}

th, td {
  padding: 5px;
}

button {
  background-color: rgba(0, 0, 0, .2);
  border: none;
  color: white;
  padding: 5px 12px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  cursor: pointer;
}

button:hover {
  box-shadow: 0 1px #fff;
}

body {
  background: fixed linear-gradient(141deg, #355c7d -50%, #6c5b7b 50%, #c06c84 175%);
  color: #fff;
}

.data-container {
  background: rgba(0, 0, 0, .2) !important;
  padding: 20px !important;
}
</style>
