<template>
<div>
<a href="#/">Retour</a>
<br><br>
<div class="data-container">
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
</div>
</div>  
</template>

<script>
export default {
  name: 'train',

  data: () => ({
    numTrain: 'ABCDEF123',
    villeDepart: 'Montpellier',
    villeArrivee: 'Paris',
    heureDepart: new Date()
  })
}
</script>
