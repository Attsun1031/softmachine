// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
require('materialize-css')
require('../node_modules/materialize-css/dist/css/materialize.min.css')
require('../node_modules/vis/dist/vis.min.css')
require('es6-promise').polyfill()

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
