<template>
  <div class="hello">
    <h2>Mes réservations</h2>
    <center class="data-container">
      <div v-if="!reservations">
        Chargement ...
      </div>
      <template v-else>
        <table>
          <thead>
            <tr>
              <th>Numéro</th>
              <th>Train</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="res in reservations" :key="res.numero">
              <td v-html="res.numReservation"></td>
              <td>
                <a :href="`#/train/${res.currentTrain}`">{{ res.currentTrain }}</a>
              </td>
              <td>
                <button @click="cancelReservation(res.numReservation)">Annuler</button>
              </td>
            </tr>
          </tbody>
        </table>
      </template>
    </center>

    <h2>Trains</h2>

    <center class="data-container">
      <div v-if="!trains">
        Chargement ...
      </div>
      <template v-else>
        <table>
          <thead>
            <tr>
              <th>Train N°</th>
              <th>Départ</th>
              <th>Arrivée</th>
              <th>Heure départ</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="train in trains" :key="train.numTrain">
              <td>
                <a :href="`#/train/${train.numTrain}`">{{ train.numTrain }}</a>
              </td>
              <td v-html="train.villeDepart"></td>
              <td v-html="train.villeArrivee"></td>
              <td v-html="train.heureDepart && new Date(train.heureDepart).toLocaleString()"></td>
            </tr>
          </tbody>
        </table>
      </template>
    </center>
  </div>
</template>

<script>
// const getTrains = () => [{
//   numTrain: 'ABCDEF123',
//   villeDepart: 'Montpellier',
//   villeArrivee: 'Paris',
//   heureDepart: new Date()
// }, {
//   numTrain: 'ABCDEF123',
//   villeDepart: 'Montpellier',
//   villeArrivee: 'Paris',
//   heureDepart: new Date()
// }, {
//   numTrain: 'ABCDEF123',
//   villeDepart: 'Montpellier',
//   villeArrivee: 'Paris',
//   heureDepart: new Date()
// }, {
//   numTrain: 'ABCDEF123',
//   villeDepart: 'Montpellier',
//   villeArrivee: 'Paris',
//   heureDepart: new Date()
// }]

// const getReservations = (name) => [{
//   numero: '1',
//   date: new Date(),
//   numTrain: 'ABCDEF123'
// }, {
//   numero: '1',
//   date: new Date(),
//   numTrain: 'ABCDEF123'
// }, {
//   numero: '1',
//   date: new Date(),
//   numTrain: 'ABCDEF123'
// }]

export default {
  name: 'Index',
  data: () => ({
    trains: null,
    reservations: null
  }),

  async created () {
    this.updateReservations()

    try {
      const trainsResponse = await this.axios.get(`/api/trains`)

      if (trainsResponse.status !== 200) {
        throw new Error('Erreur de récupération des trains')
      }

      this.trains = trainsResponse.data
    } catch (e) {
      console.error(e)
      // this.trains = getTrains()
    }
  },

  methods: {
    async cancelReservation (numero) {
      await this.axios.delete(`/api/reservations/${numero}`)
      this.updateReservations()
    },

    async updateReservations () {
      try {
        const reservationsResponse = await this.axios.get(`/api/reservations?username=${this.$store.getters.username}`)

        if (reservationsResponse.status !== 200) {
          throw new Error('Erreur de récupération des réservations')
        }

        this.reservations = reservationsResponse.data.filter(x => x.username === this.$store.getters.username)
      } catch (e) {
        console.error(e)
        // this.reservations = getReservations(this.$store.getters.username)
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
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
</style>
