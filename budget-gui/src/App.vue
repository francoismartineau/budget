<template>
  <div id="app">
    <date></date>
    <bnc></bnc>
  </div>
</template>

<script>
import date from './components/Date'
import bnc from './components/Bnc'

export default {
  name: 'app',
  components: {
    date,
    bnc,
  }
}
</script>

<style>
#app {
  font-family: "Helvetica";
}
</style>
