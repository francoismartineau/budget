import Vue from 'vue';
import Vuex from 'vuex'
import 'es6-promise/auto'
Vue.use(Vuex)

export const store = new Vuex.Store({
    state: {
      date: {
        day: 1,
        month: 10,
        year: 2018,
      },
    },
    mutations: {
      setDay(state, day) {
        state.date.day = day
      },
      setMonth(state, month) {
        state.date.month = month
      },
      setYear(state, year) {
        state.date.year = year
      }
    },
  })