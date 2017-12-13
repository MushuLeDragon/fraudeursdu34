<template>
  <div class="hello">
    <h2>Mes réservations</h2>
    <center class="data-container">
      <table>
        <thead>
          <tr>
            <th>Numéro</th>
            <th>Date</th>
            <th>Train</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="res in reservations" :key="res.numero">
            <td v-html="res.numero"></td>
            <td v-html="res.date.toLocaleDateString()"></td>
            <td>
              <a :href="`#/train/${res.numTrain}`">{{ res.numTrain }}</a>
            </td>
            <td>
              <button>Annuler</button>
            </td>
          </tr>
        </tbody>
      </table>
    </center>

    <h2>Trains</h2>

    <center class="data-container">
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
            <td v-html="train.heureDepart.toLocaleString()"></td>
          </tr>
        </tbody>
      </table>
    </center>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  data: () => ({
    trains: [{
      numTrain: 'ABCDEF123',
      villeDepart: 'Montpellier',
      villeArrivee: 'Paris',
      heureDepart: new Date()
    }, {
      numTrain: 'ABCDEF123',
      villeDepart: 'Montpellier',
      villeArrivee: 'Paris',
      heureDepart: new Date()
    }, {
      numTrain: 'ABCDEF123',
      villeDepart: 'Montpellier',
      villeArrivee: 'Paris',
      heureDepart: new Date()
    }, {
      numTrain: 'ABCDEF123',
      villeDepart: 'Montpellier',
      villeArrivee: 'Paris',
      heureDepart: new Date()
    }],

    reservations: [{
      numero: '1',
      date: new Date(),
      numTrain: 'ABCDEF123'
    }, {
      numero: '1',
      date: new Date(),
      numTrain: 'ABCDEF123'
    }, {
      numero: '1',
      date: new Date(),
      numTrain: 'ABCDEF123'
    }]
  })
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
