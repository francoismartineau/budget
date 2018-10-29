<template>
  <div id="app">
    <bnc></bnc>
    <date></date>
    <description></description>
  </div>
</template>

<script>
import date from './components/Date'
import description from './components/Description'
import bnc from './components/Bnc'

export default {
  name: 'app',
  components: {
    date,
    bnc,
    description,
  }
}
</script>

<style>
#app {
  font-family: "Helvetica";
}
</style>
