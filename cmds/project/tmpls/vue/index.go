package vue

//IndexHTML .
const IndexHTML = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width,initial-scale=1.0">
    <link rel="icon" href="<%= BASE_URL %>favicon.ico">
    <title>vue3</title>
  </head>
  <body>
    <noscript>
      <strong>We're sorry but vue3 doesn't work properly without JavaScript enabled. Please enable it to continue.</strong>
    </noscript>
    <div id="app"></div>
    <!-- built files will be auto injected -->
  </body>
</html>`

const EnvProd = `NODE_ENV=production
VUE_APP_API_URL=http://{{.ip}}:9090
`
const EnvDev = `NODE_ENV=development
VUE_APP_API_URL=http://{{.ip}}:9090
`

//PackAgeJSON .
const PackAgeJSON = `
{
  "name": "{{.projectName}}",
  "version": "0.1.0",
  "private": true,
  "scripts": {
    "serve": "vue-cli-service serve --mode dev",
    "build": "vue-cli-service build --mode prod",
    "lint": "vue-cli-service lint"
  },
  "dependencies": {
    "bootstrap": "^4.1.3",
    "element-ui": "^2.4.5",
    "jquery": "^3.3.1",
    "popper.js": "^1.14.6",
    "vue": "^2.5.17",
    "vue-router": "^3.0.2",
    "vuex": "^3.0.1",
    "nav-menu":"^1.2.9",
    "login-with-up":"^1.1.1",
    "axios":"^0.18.0",
    "qs":"^6.6.0",
    "vue-cookies": "^1.5.12"
  },
  "devDependencies": {
    "@vue/cli-plugin-babel": "^3.2.0",
    "@vue/cli-plugin-eslint": "^3.2.0",
    "@vue/cli-service": "^3.2.0",
    "babel-eslint": "^10.0.1",
    "eslint": "^5.8.0",
    "eslint-plugin-vue": "^5.0.0-0",
    "vue-template-compiler": "^2.5.17"
  },
  "eslintConfig": {
    "root": true,
    "env": {
      "node": true
    },
    "extends": [
      "plugin:vue/essential",
      "eslint:recommended"
    ],
    "rules": {},
    "parserOptions": {
      "parser": "babel-eslint"
    }
  },
  "postcss": {
    "plugins": {
      "autoprefixer": {}
    }
  },
  "browserslist": [
    "> 1%",
    "last 2 versions",
    "not ie <= 8"
  ]
}
`

const Babel = `module.exports = {
	presets: [
	  '@vue/app'
	]
}`

const PostCss = `module.exports = {
	plugins: {
	  autoprefixer: {}
	}
}`

const MainGo = `package main

import (
	'github.com/micro-plat/hydra/hydra'
)

type mgrweb struct {
	*hydra.MicroApp
}

func main() {
	app := &mgrweb{
		hydra.NewApp(
			hydra.WithPlatName('{{.projectName}}'),
			hydra.WithSystemName('{{.projectName}}'),
			hydra.WithServerTypes('web'),
		),
	}

	app.install()
	app.Start()
}

func (s *mgrweb) install() {
	s.Conf.WEB.SetMainConf("{'address':':8089'}")
	s.Conf.WEB.SetSubConf('static', "{
			'dir':'./static',
			'rewriters':['*'],
			'exts':['.ttf','.woff','.woff2']			
	}")
}`

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

const RouterString = `{{range $i,$c:=.router -}}
    {
      path: '{{$i}}',
      name:'{{$i|vname}}',
      component: () => import('./pages/{{$c}}')
    },
{{- end -}}`

const Router = `import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'menu',
      component: () => import('./pages/menu/menu.vue'),
      children:[
        //page.router#//
        //#page.router//
      ]
    },{
      path:'/login',
      name:'login',
      component: () => import('./pages/login/login.vue'),
    }
  ]
})`

const Store = `import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {

  },
  mutations: {

  },
  actions: {

  }
})`

const AppVue = `<template>
  <div>
    <router-view/>
  </div>
</template>

<script>
  import VueCookies from 'vue-cookies'
  export default {
    methods: {
      beforeunloadHandler (e) {
        
        sessionStorage.removeItem("__jwt__")
        VueCookies.remove("__jwt__")
      }
    },
    destroyed() {
    window.removeEventListener('beforeunload', e => this.beforeunloadHandler(e))
    }
  }
</script>

<style>

</style>`
