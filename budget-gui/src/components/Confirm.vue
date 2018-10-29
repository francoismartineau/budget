<template>
  <div>
    <el-button>Confirm</el-button>
  </div>
</template>

<script>
export default {
  name: 'Confirm',

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
