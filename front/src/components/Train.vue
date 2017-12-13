<template>
<div>
<a href="#/">Retour</a>
<br><br>
<div class="data-container">
  <div v-if="!numTrain">
    Chargement ...
  </div>
  <template v-else>
    <strong>Train numéro</strong> : {{ $route.params.id }}
    <br><br>
    <strong>Ville de départ</strong> : {{ villeDepart }}
    <br><br>
    <strong>Ville d'arrivée</strong> : {{ villeArrivee }}
    <br><br>
    <strong>Horaire de départ</strong> : {{ heureDepart && new Date(heureDepart).toLocaleString() }}
    <br><br>
    <button @click="placeReservation">
      Réserver
    </button>
  </template>
</div>
</div>  
</template>

<script>
// const getTrain = numTrain => ({
//   numTrain: 'ABCDEF123',
//   villeDepart: 'Montpellier',
//   villeArrivee: 'Paris',
//   heureDepart: new Date()
// })

export default {
  name: 'train',

  data: () => ({
    numTrain: null,
    villeDepart: null,
    villeArrivee: null,
    heureDepart: null
  }),

  methods: {
    async updateTrain () {
      let train

      try {
        const trainResponse = await this.axios.get(`/api/trains/${this.$route.params.id}`)

        if (trainResponse.status !== 200) {
          throw new Error('Erreur de récupération des trains')
        }

        train = trainResponse.data
      } catch (e) {
        console.error(e)
        // train = getTrain()
      }

      // console.log(train)

      this.numTrain = train.numTrain
      this.villeDepart = train.villeDepart
      this.villeArrivee = train.villeArrivee
      this.heureDepart = train.heureDepart
    },

    placeReservation () {
      try {
        const payload = {
          numReservation: Date.now(),
          username: this.$store.getters.username,
          currentTrain: this.numTrain,
          numberPlaces: 1
        }

        this.axios.post(`/api/reservations`, payload)
      } catch (e) {
        console.error(e)
        // Afficher "impossible de créer la réservation ..."
      }
    }
  },

  created () {
    this.updateTrain()
  },

  watch: {
    '$route': 'updateTrain'
  }
}
</script>
