<template>
  <div>
    <el-input v-model="description"></el-input>
    <p>Descri: {{ description }}</p>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: 'Description',
  computed: {
    description: {
      set(description) {
        this.$store.commit('setDescription', description);
      },
      get() {
        return this.$store.state.description;
      }
    },
  },

}
</script>

<style>
</style>
