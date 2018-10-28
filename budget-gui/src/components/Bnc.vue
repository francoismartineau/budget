<template>
  <div>
    <a :href="url" target="_blank">BNC</a>
  </div>
</template>

<script>
export default {
  name: 'Bnc',

  data() {
    return {
      url: "https://www.bnc.ca",
    }
  },

  methods: {
      open: function() {
        let routeData = this.$router.resolve({name: 'routeName', query: {data: "someData"}});
        window.open(routeData.href, '_blank');
      }
  },
}
</script>

<style>
</style>
