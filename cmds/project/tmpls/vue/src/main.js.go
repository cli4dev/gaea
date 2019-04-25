package src

const MainJS = `import "jquery"
import "bootstrap"
import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
Vue.use(ElementUI);
Vue.config.productionTip = false
console.log("当前环境：", process.env.NODE_ENV)
import {get, post, patch, put, del } from './util/http'

Vue.prototype.$get = get;
Vue.prototype.$post = post;
Vue.prototype.$patch = patch;
Vue.prototype.$put = put;
Vue.prototype.$del = del;


new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')`
