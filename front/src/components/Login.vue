<template>
  <div class="data-container">
    <p>
      <input type="text" v-model="username" placeholder="Nom d'utilisateur" />
      <br><br>
      <button type="button" @click="connect">Se connecter</button>
    </p>
  </div>
</template>

<script>
export default {
  name: 'login',

  data: () => ({
    username: null
  }),

  methods: {
    connect () {
      this.$store.commit('connect', this.username)
    }
  }
}
</script>

<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
