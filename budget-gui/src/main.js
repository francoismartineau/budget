import Vue from 'vue'
import { store } from './store'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import VueRouter from 'vue-router'
import VueResource from 'vue-resource';
import App from './App.vue'

Vue.use(VueRouter)
Vue.use(VueResource)
Vue.use(require('vue-shortkey'))
Vue.use(ElementUI);
Vue.config.productionTip = false

new Vue({
  store,
  render: h => h(App)
}).$mount('#app')
