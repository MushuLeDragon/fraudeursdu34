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
    <strong>Horaire de départ</strong> : {{ heureDepart.toLocaleString() }}
    <br><br>
    <button>
      Réserver
    </button>
  </template>
</div>
</div>  
</template>

<script>
const getTrain = numTrain => ({
  numTrain: 'ABCDEF123',
  villeDepart: 'Montpellier',
  villeArrivee: 'Paris',
  heureDepart: new Date()
})

const placeReservation = (username, numTrain) => {}

export default {
  name: 'train',

  data: () => ({
    numTrain: null,
    villeDepart: null,
    villeArrivee: null,
    heureDepart: null
  }),

  methods: {
    updateTrain () {
      const train = getTrain()

      this.numTrain = train.numTrain
      this.villeDepart = train.villeDepart
      this.villeArrivee = train.villeArrivee
      this.heureDepart = train.heureDepart
    },

    placeReservation () {
      placeReservation(this.$store.getters.username, this.numTrain)
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
