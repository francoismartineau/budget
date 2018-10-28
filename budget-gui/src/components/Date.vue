<!-- TODO
  .Changer le format des boîtes
  .d / Ctrl + d   pour modifier le jour
  .unit tests  
-->

<template>
  <div>
    <div v-shortkey="['arrowright']" @shortkey="increment()"></div>
    <div v-shortkey="['arrowleft']" @shortkey="decrement()"></div>
    
    <div class="title">Date</div>
    <el-input-number v-model="day" controls-position="right"></el-input-number>
    <el-input-number v-model="month" controls-position="right"></el-input-number>
    <el-input-number v-model="year" controls-position="right"></el-input-number>

  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: 'Date',

  data() {
    return {
      thirtyDays: [2, 4, 6, 9, 11],
    }
  },

  methods: {
    isLeap: function(year) {
      var isLeap = false;
      if (year % 400 == 0 || (year % 4 == 0 && year % 100 != 0)) {
        isLeap = true
      } 
      return isLeap
    },
    
    increment: function() {
      switch(++this.day) {
        case 29:
          if (this.month == 2 && !this.isLeap(this.year)) {
            this.day = 1
            this.month++
          }
          break
        case 30:
          if (this.month == 2) {
            this.day = 1
            this.month++
          }
          break
        case 31:
          if (this.thirtyDays.includes(this.month)) {
            this.day = 1
            this.month++
          }
          break
        case 32:
          this.day = 1
          this.month++
      }
      if (this.month > 12) {
        this.month = 1
        this.year++
      }
    },
    decrement: function() {
      this.day--
      if (this.day == 0) {
        this.month--
        if (this.month == 0) {
          this.day = 31
          this.month = 12
          this.year--
        } else if (this.month == 2) {
          if (this.isLeap(this.year)) {
            this.day = 29
          } else {
            this.day = 28
          }
        } else if (this.thirtyDays.includes(this.month)) {
          this.day = 30
        } else {
          this.day = 31
        }
      }
    },
  },

  computed: {
    ...mapState(['date']),
    day: {
      set(day) {
        this.$store.commit('setDay', day);
      },
      get() {
        return this.date.day
      },
    },
    month: {
      set(month) {
        this.$store.commit('setMonth', month);
      },
      get() {
        return this.date.month
      },
    },
    year: {
      set(year) {
        this.$store.commit('setYear', year);
      },
      get() {
        return this.date.year
      },
    }
  },
}
</script>

<style>

.title {
  font-size: 30px;
}

/* Date boxes */
.el-input-number {
    padding: 0px;
    width: 100px;
}

/* Increment buttons */
.el-input-number__decrease, .el-input-number__increase {
  width: 20px;
}
</style>
