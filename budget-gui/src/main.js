import Vue from 'vue'
import 'element-ui/lib/theme-chalk/index.css';
import { store } from './store'
import App from './App.vue'

Vue.use(require('vue-router'))
Vue.use(require('vue-resource'))
Vue.use(require('vue-shortkey'))
Vue.use(require('element-ui'))
Vue.config.productionTip = false

new Vue({
  store,
  render: h => h(App)
}).$mount('#app')
