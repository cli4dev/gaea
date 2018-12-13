package vue

//IndexHTML .
const IndexHTML = `
<!DOCTYPE html>
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
VUE_APP_API_URL=http://127.0.0.1:8090
`
const EnvDev = `NODE_ENV=development
VUE_APP_API_URL=http://127.0.0.1:8090
`

//PackAgeJSON .
const PackAgeJSON = `
{
  "name": "vue3",
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
    "qs":"^6.6.0"
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

const Babel = `
module.exports = {
	presets: [
	  '@vue/app'
	]
}`

const PostCss = `
module.exports = {
	plugins: {
	  autoprefixer: {}
	}
}`

const MainGo = `
package main

import (
	"github.com/micro-plat/hydra/hydra"
)

type mgrweb struct {
	*hydra.MicroApp
}

func main() {
	app := &mgrweb{
		hydra.NewApp(
			hydra.WithPlatName("coupon"),
			hydra.WithSystemName("mgrweb"),
			hydra.WithServerTypes("web"),
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

const MainJS = `
import "jquery"
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
}).$mount('#app')
`
const RouterString = `{{range $i,$c:=.router -}}
    {
      path: '{{$i}}',
      name:'{{$i|vname}}',
      component: () => import('./pages/{{$c}}')
    },
{{- end -}}`

const Router = `
import Vue from 'vue'
import Router from 'vue-router'
import Home from './pages/HelloWorld'

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
})
`

const Store = `
import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {

  },
  mutations: {

  },
  actions: {

  }
})
`
const AppVue = `
<template>
  <div>
    <router-view/>
  </div>
</template>

<style>

</style>
`

const HelloVue = `
<template>
  <div>
    <h1>{{ msg }}</h1>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  }
}
</script>
</style>
`

const PackAgeLock = `{
  "name": "vue3",
  "version": "0.1.0",
  "lockfileVersion": 1,
  "requires": true,
  "dependencies": {
    "@babel/code-frame": {
      "version": "7.0.0",
      "resolved": "http://registry.npm.taobao.org/@babel/code-frame/download/@babel/code-frame-7.0.0.tgz",
      "integrity": "sha1-BuKrGb21NThVWaq7W6WXKUgoAPg=",
      "dev": true,
      "requires": {
        "@babel/highlight": "^7.0.0"
      }
    },
    "@babel/core": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/core/download/@babel/core-7.2.0.tgz",
      "integrity": "sha1-pN04FJAZmOkzQPAIbphn/voWOto=",
      "dev": true,
      "requires": {
        "@babel/code-frame": "^7.0.0",
        "@babel/generator": "^7.2.0",
        "@babel/helpers": "^7.2.0",
        "@babel/parser": "^7.2.0",
        "@babel/template": "^7.1.2",
        "@babel/traverse": "^7.1.6",
        "@babel/types": "^7.2.0",
        "convert-source-map": "^1.1.0",
        "debug": "^4.1.0",
        "json5": "^2.1.0",
        "lodash": "^4.17.10",
        "resolve": "^1.3.2",
        "semver": "^5.4.1",
        "source-map": "^0.5.0"
      }
    },
    "@babel/generator": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/generator/download/@babel/generator-7.2.0.tgz",
      "integrity": "sha1-6vOCH6AwHZ1K74jmPUvMGbc7oWw=",
      "dev": true,
      "requires": {
        "@babel/types": "^7.2.0",
        "jsesc": "^2.5.1",
        "lodash": "^4.17.10",
        "source-map": "^0.5.0",
        "trim-right": "^1.0.1"
      }
    },
    "@babel/helper-annotate-as-pure": {
      "version": "7.0.0",
      "resolved": "http://registry.npm.taobao.org/@babel/helper-annotate-as-pure/download/@babel/helper-annotate-as-pure-7.0.0.tgz",
      "integrity": "sha1-Mj053QtQ4Qx8Bsp9djjmhk2MXDI=",
      "dev": true,
      "requires": {
        "@babel/types": "^7.0.0"
      }
    },
    "@babel/helper-builder-binary-assignment-operator-visitor": {
      "version": "7.1.0",
      "resolved": "http://registry.npm.taobao.org/@babel/helper-builder-binary-assignment-operator-visitor/download/@babel/helper-builder-binary-assignment-operator-visitor-7.1.0.tgz",
      "integrity": "sha1-a2lijf5Ah3mODE7Zjj1Kay+9L18=",
      "dev": true,
      "requires": {
        "@babel/helper-explode-assignable-expression": "^7.1.0",
        "@babel/types": "^7.0.0"
      }
    },
    "@babel/helper-call-delegate": {
      "version": "7.1.0",
      "resolved": "http://registry.npm.taobao.org/@babel/helper-call-delegate/download/@babel/helper-call-delegate-7.1.0.tgz",
      "integrity": "sha1-apV/EF83dV6GRTQ9MDiiLhRJzEo=",
      "dev": true,
      "requires": {
        "@babel/helper-hoist-variables": "^7.0.0",
        "@babel/traverse": "^7.1.0",
        "@babel/types": "^7.0.0"
      }
    },
    "@babel/helper-create-class-features-plugin": {
      "version": "7.2.1",
      "resolved": "http://registry.npm.taobao.org/@babel/helper-create-class-features-plugin/download/@babel/helper-create-class-features-plugin-7.2.1.tgz",
      "integrity": "sha1-9ugCcpFmnvZEMyINyDJ1MSM/EWE=",
      "dev": true,
      "requires": {
        "@babel/helper-function-name": "^7.1.0",
        "@babel/helper-member-expression-to-functions": "^7.0.0",
        "@babel/helper-optimise-call-expression": "^7.0.0",
        "@babel/helper-plugin-utils": "^7.0.0",
        "@babel/helper-replace-supers": "^7.1.0"
      }
    },
    "@babel/helper-define-map": {
      "version": "7.1.0",
      "resolved": "http://registry.npm.taobao.org/@babel/helper-define-map/download/@babel/helper-define-map-7.1.0.tgz",
      "integrity": "sha1-O3TK7DKbPIDBFikIh8DdmuRowgw=",
      "dev": true,
      "requires": {
        "@babel/helper-function-name": "^7.1.0",
        "@babel/types": "^7.0.0",
        "lodash": "^4.17.10"
      }
    },
    "@babel/helper-explode-assignable-expression": {
      "version": "7.1.0",
      "resolved": "http://registry.npm.taobao.org/@babel/helper-explode-assignable-expression/download/@babel/helper-explode-assignable-expression-7.1.0.tgz",
      "integrity": "sha1-U3+hP28WdN90WwwA7I/k6ZaByPY=",
      "dev": true,
      "requires": {
        "@babel/traverse": "^7.1.0",
        "@babel/types": "^7.0.0"
      }
    },
    "@babel/helper-function-name": {
      "version": "7.1.0",
      "resolved": "http://registry.npm.taobao.org/@babel/helper-function-name/download/@babel/helper-function-name-7.1.0.tgz",
      "integrity": "sha1-oM6wFoX3M1XUNgwSR/WCv6/I/1M=",
      "dev": true,
      "requires": {
        "@babel/helper-get-function-arity": "^7.0.0",
        "@babel/template": "^7.1.0",
        "@babel/types": "^7.0.0"
      }
    },
    "@babel/helper-get-function-arity": {
      "version": "7.0.0",
      "resolved": "http://registry.npm.taobao.org/@babel/helper-get-function-arity/download/@babel/helper-get-function-arity-7.0.0.tgz",
      "integrity": "sha1-g1ctQyDipGVyY3NBE8QoaLZOScM=",
      "dev": true,
      "requires": {
        "@babel/types": "^7.0.0"
      }
    },
    "@babel/helper-hoist-variables": {
      "version": "7.0.0",
      "resolved": "http://registry.npm.taobao.org/@babel/helper-hoist-variables/download/@babel/helper-hoist-variables-7.0.0.tgz",
      "integrity": "sha1-Rq3ExedYZFrnpF3rkrqwkYwju4g=",
      "dev": true,
      "requires": {
        "@babel/types": "^7.0.0"
      }
    },
    "@babel/helper-member-expression-to-functions": {
      "version": "7.0.0",
      "resolved": "http://registry.npm.taobao.org/@babel/helper-member-expression-to-functions/download/@babel/helper-member-expression-to-functions-7.0.0.tgz",
      "integrity": "sha1-jNFLCg33/wDwCefXpDaUX0fHoW8=",
      "dev": true,
      "requires": {
        "@babel/types": "^7.0.0"
      }
    },
    "@babel/helper-module-imports": {
      "version": "7.0.0",
      "resolved": "http://registry.npm.taobao.org/@babel/helper-module-imports/download/@babel/helper-module-imports-7.0.0.tgz",
      "integrity": "sha1-lggbcRHkhtpNLNlxrRpP4hbMLj0=",
      "dev": true,
      "requires": {
        "@babel/types": "^7.0.0"
      }
    },
    "@babel/helper-module-transforms": {
      "version": "7.1.0",
      "resolved": "http://registry.npm.taobao.org/@babel/helper-module-transforms/download/@babel/helper-module-transforms-7.1.0.tgz",
      "integrity": "sha1-Rw1PlnbZ+tULMkzczl+6u8PaV4c=",
      "dev": true,
      "requires": {
        "@babel/helper-module-imports": "^7.0.0",
        "@babel/helper-simple-access": "^7.1.0",
        "@babel/helper-split-export-declaration": "^7.0.0",
        "@babel/template": "^7.1.0",
        "@babel/types": "^7.0.0",
        "lodash": "^4.17.10"
      }
    },
    "@babel/helper-optimise-call-expression": {
      "version": "7.0.0",
      "resolved": "http://registry.npm.taobao.org/@babel/helper-optimise-call-expression/download/@babel/helper-optimise-call-expression-7.0.0.tgz",
      "integrity": "sha1-opIMVwKwc8Fd5REGIAqoytIEl9U=",
      "dev": true,
      "requires": {
        "@babel/types": "^7.0.0"
      }
    },
    "@babel/helper-plugin-utils": {
      "version": "7.0.0",
      "resolved": "http://registry.npm.taobao.org/@babel/helper-plugin-utils/download/@babel/helper-plugin-utils-7.0.0.tgz",
      "integrity": "sha1-u7P77phmHFaQNCN8wDlnupm08lA=",
      "dev": true
    },
    "@babel/helper-regex": {
      "version": "7.0.0",
      "resolved": "http://registry.npm.taobao.org/@babel/helper-regex/download/@babel/helper-regex-7.0.0.tgz",
      "integrity": "sha1-LBcYkjtX+bvmRwX/5WQKxk2b2yc=",
      "dev": true,
      "requires": {
        "lodash": "^4.17.10"
      }
    },
    "@babel/helper-remap-async-to-generator": {
      "version": "7.1.0",
      "resolved": "http://registry.npm.taobao.org/@babel/helper-remap-async-to-generator/download/@babel/helper-remap-async-to-generator-7.1.0.tgz",
      "integrity": "sha1-Nh2AghtvONp1vT8HheziCojF/n8=",
      "dev": true,
      "requires": {
        "@babel/helper-annotate-as-pure": "^7.0.0",
        "@babel/helper-wrap-function": "^7.1.0",
        "@babel/template": "^7.1.0",
        "@babel/traverse": "^7.1.0",
        "@babel/types": "^7.0.0"
      }
    },
    "@babel/helper-replace-supers": {
      "version": "7.1.0",
      "resolved": "http://registry.npm.taobao.org/@babel/helper-replace-supers/download/@babel/helper-replace-supers-7.1.0.tgz",
      "integrity": "sha1-X8Md5SLsDvCJncmz589qXdZV82I=",
      "dev": true,
      "requires": {
        "@babel/helper-member-expression-to-functions": "^7.0.0",
        "@babel/helper-optimise-call-expression": "^7.0.0",
        "@babel/traverse": "^7.1.0",
        "@babel/types": "^7.0.0"
      }
    },
    "@babel/helper-simple-access": {
      "version": "7.1.0",
      "resolved": "http://registry.npm.taobao.org/@babel/helper-simple-access/download/@babel/helper-simple-access-7.1.0.tgz",
      "integrity": "sha1-Ze65VMjCRb6qToWdphiPOdceWFw=",
      "dev": true,
      "requires": {
        "@babel/template": "^7.1.0",
        "@babel/types": "^7.0.0"
      }
    },
    "@babel/helper-split-export-declaration": {
      "version": "7.0.0",
      "resolved": "http://registry.npm.taobao.org/@babel/helper-split-export-declaration/download/@babel/helper-split-export-declaration-7.0.0.tgz",
      "integrity": "sha1-Oq4oXAMRwqsJXZl7jJqUytVH2BM=",
      "dev": true,
      "requires": {
        "@babel/types": "^7.0.0"
      }
    },
    "@babel/helper-wrap-function": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/helper-wrap-function/download/@babel/helper-wrap-function-7.2.0.tgz",
      "integrity": "sha1-xOABJEV2nigVtVKW6tQ6lYVJ9vo=",
      "dev": true,
      "requires": {
        "@babel/helper-function-name": "^7.1.0",
        "@babel/template": "^7.1.0",
        "@babel/traverse": "^7.1.0",
        "@babel/types": "^7.2.0"
      }
    },
    "@babel/helpers": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/helpers/download/@babel/helpers-7.2.0.tgz",
      "integrity": "sha1-gzXzFA8xRCcNxjxHMqT4sKULeiE=",
      "dev": true,
      "requires": {
        "@babel/template": "^7.1.2",
        "@babel/traverse": "^7.1.5",
        "@babel/types": "^7.2.0"
      }
    },
    "@babel/highlight": {
      "version": "7.0.0",
      "resolved": "http://registry.npm.taobao.org/@babel/highlight/download/@babel/highlight-7.0.0.tgz",
      "integrity": "sha1-9xDDjI1Fjm3ZogGvtjf8t4HOmeQ=",
      "dev": true,
      "requires": {
        "chalk": "^2.0.0",
        "esutils": "^2.0.2",
        "js-tokens": "^4.0.0"
      }
    },
    "@babel/parser": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/parser/download/@babel/parser-7.2.0.tgz",
      "integrity": "sha1-AtAdvDMLbL82t2rJPFB1LGkCcGU=",
      "dev": true
    },
    "@babel/plugin-proposal-async-generator-functions": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-proposal-async-generator-functions/download/@babel/plugin-proposal-async-generator-functions-7.2.0.tgz",
      "integrity": "sha1-somzBmadzkrSCwJSiJoVdoydQX4=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0",
        "@babel/helper-remap-async-to-generator": "^7.1.0",
        "@babel/plugin-syntax-async-generators": "^7.2.0"
      }
    },
    "@babel/plugin-proposal-class-properties": {
      "version": "7.2.1",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-proposal-class-properties/download/@babel/plugin-proposal-class-properties-7.2.1.tgz",
      "integrity": "sha1-xzSlPgoexA/lwi7lBp0m2jsYfQU=",
      "dev": true,
      "requires": {
        "@babel/helper-create-class-features-plugin": "^7.2.1",
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-proposal-decorators": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-proposal-decorators/download/@babel/plugin-proposal-decorators-7.2.0.tgz",
      "integrity": "sha1-a0J4KCpvXdCLXYm5TyGqFnH+oHE=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0",
        "@babel/helper-replace-supers": "^7.1.0",
        "@babel/helper-split-export-declaration": "^7.0.0",
        "@babel/plugin-syntax-decorators": "^7.2.0"
      }
    },
    "@babel/plugin-proposal-json-strings": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-proposal-json-strings/download/@babel/plugin-proposal-json-strings-7.2.0.tgz",
      "integrity": "sha1-Vo7MRGxhSK5rJn8CVREwiR4p8xc=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0",
        "@babel/plugin-syntax-json-strings": "^7.2.0"
      }
    },
    "@babel/plugin-proposal-object-rest-spread": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-proposal-object-rest-spread/download/@babel/plugin-proposal-object-rest-spread-7.2.0.tgz",
      "integrity": "sha1-iPX+w+etAZAUyX9+48mS8K2/f7g=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0",
        "@babel/plugin-syntax-object-rest-spread": "^7.2.0"
      }
    },
    "@babel/plugin-proposal-optional-catch-binding": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-proposal-optional-catch-binding/download/@babel/plugin-proposal-optional-catch-binding-7.2.0.tgz",
      "integrity": "sha1-E12B7baKCB5V5W7EhUHs6AZcOPU=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0",
        "@babel/plugin-syntax-optional-catch-binding": "^7.2.0"
      }
    },
    "@babel/plugin-proposal-unicode-property-regex": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-proposal-unicode-property-regex/download/@babel/plugin-proposal-unicode-property-regex-7.2.0.tgz",
      "integrity": "sha1-q+coH+Rsld3BQ6ZeU1hkd5IDlSA=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0",
        "@babel/helper-regex": "^7.0.0",
        "regexpu-core": "^4.2.0"
      }
    },
    "@babel/plugin-syntax-async-generators": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-syntax-async-generators/download/@babel/plugin-syntax-async-generators-7.2.0.tgz",
      "integrity": "sha1-aeHw2zTG9aDPfiszI78VmnbIy38=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-syntax-decorators": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-syntax-decorators/download/@babel/plugin-syntax-decorators-7.2.0.tgz",
      "integrity": "sha1-xQsblX3MaeSxEntl4cM+72FXDBs=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-syntax-dynamic-import": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-syntax-dynamic-import/download/@babel/plugin-syntax-dynamic-import-7.2.0.tgz",
      "integrity": "sha1-acFZ/69JmBIhYa2OvF5tH1XfhhI=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-syntax-json-strings": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-syntax-json-strings/download/@babel/plugin-syntax-json-strings-7.2.0.tgz",
      "integrity": "sha1-cr0T9v/h0lk4Ep0qGGsR/WKVFHA=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-syntax-jsx": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-syntax-jsx/download/@babel/plugin-syntax-jsx-7.2.0.tgz",
      "integrity": "sha1-C4WjtLx830zEuL8jYzW5B8oi58c=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-syntax-object-rest-spread": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-syntax-object-rest-spread/download/@babel/plugin-syntax-object-rest-spread-7.2.0.tgz",
      "integrity": "sha1-O3o+czUQxX6CC5FCpleayLDfrS4=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-syntax-optional-catch-binding": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-syntax-optional-catch-binding/download/@babel/plugin-syntax-optional-catch-binding-7.2.0.tgz",
      "integrity": "sha1-qUAT1u2okI3+akd+f57ahWVuz1w=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-transform-arrow-functions": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-arrow-functions/download/@babel/plugin-transform-arrow-functions-7.2.0.tgz",
      "integrity": "sha1-mur75Nb/xlY7+Pg3IJFijwB3lVA=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-transform-async-to-generator": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-async-to-generator/download/@babel/plugin-transform-async-to-generator-7.2.0.tgz",
      "integrity": "sha1-aLikOGY+iFGeZbd2+JOPNEWxov8=",
      "dev": true,
      "requires": {
        "@babel/helper-module-imports": "^7.0.0",
        "@babel/helper-plugin-utils": "^7.0.0",
        "@babel/helper-remap-async-to-generator": "^7.1.0"
      }
    },
    "@babel/plugin-transform-block-scoped-functions": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-block-scoped-functions/download/@babel/plugin-transform-block-scoped-functions-7.2.0.tgz",
      "integrity": "sha1-XTzBHo1d3XUqpkyRSNDbbLef0ZA=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-transform-block-scoping": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-block-scoping/download/@babel/plugin-transform-block-scoping-7.2.0.tgz",
      "integrity": "sha1-8XxJ2R7tvN9d1QWX0W9fL3cBMtQ=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0",
        "lodash": "^4.17.10"
      }
    },
    "@babel/plugin-transform-classes": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-classes/download/@babel/plugin-transform-classes-7.2.0.tgz",
      "integrity": "sha1-N0+IdgddfSH+pVrrXFNWElkWP5Y=",
      "dev": true,
      "requires": {
        "@babel/helper-annotate-as-pure": "^7.0.0",
        "@babel/helper-define-map": "^7.1.0",
        "@babel/helper-function-name": "^7.1.0",
        "@babel/helper-optimise-call-expression": "^7.0.0",
        "@babel/helper-plugin-utils": "^7.0.0",
        "@babel/helper-replace-supers": "^7.1.0",
        "@babel/helper-split-export-declaration": "^7.0.0",
        "globals": "^11.1.0"
      }
    },
    "@babel/plugin-transform-computed-properties": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-computed-properties/download/@babel/plugin-transform-computed-properties-7.2.0.tgz",
      "integrity": "sha1-g6ffamWIZbHI9kHVEMbzryICFto=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-transform-destructuring": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-destructuring/download/@babel/plugin-transform-destructuring-7.2.0.tgz",
      "integrity": "sha1-51JptLeInsOjMs0NDIz/j+0NxvM=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-transform-dotall-regex": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-dotall-regex/download/@babel/plugin-transform-dotall-regex-7.2.0.tgz",
      "integrity": "sha1-8Kq7k9EgqKxh6SXqC6RAgS2+Dkk=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0",
        "@babel/helper-regex": "^7.0.0",
        "regexpu-core": "^4.1.3"
      }
    },
    "@babel/plugin-transform-duplicate-keys": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-duplicate-keys/download/@babel/plugin-transform-duplicate-keys-7.2.0.tgz",
      "integrity": "sha1-2VLEkw8xKk2//xjwspFOYMNVMLM=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-transform-exponentiation-operator": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-exponentiation-operator/download/@babel/plugin-transform-exponentiation-operator-7.2.0.tgz",
      "integrity": "sha1-pjhoKJ5bQAf3BU1GSRr1FDV2YAg=",
      "dev": true,
      "requires": {
        "@babel/helper-builder-binary-assignment-operator-visitor": "^7.1.0",
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-transform-for-of": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-for-of/download/@babel/plugin-transform-for-of-7.2.0.tgz",
      "integrity": "sha1-q3RovvqA92S7A9PLXu+MyZjhytk=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-transform-function-name": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-function-name/download/@babel/plugin-transform-function-name-7.2.0.tgz",
      "integrity": "sha1-95MDYoKf+ZoxdMOfCvzAJO9Zcxo=",
      "dev": true,
      "requires": {
        "@babel/helper-function-name": "^7.1.0",
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-transform-literals": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-literals/download/@babel/plugin-transform-literals-7.2.0.tgz",
      "integrity": "sha1-aQNT6B+SZ9rU/Yz9d+r6hqulPqE=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-transform-modules-amd": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-modules-amd/download/@babel/plugin-transform-modules-amd-7.2.0.tgz",
      "integrity": "sha1-gqm85FuVRB9heiQBHcidEtp/TuY=",
      "dev": true,
      "requires": {
        "@babel/helper-module-transforms": "^7.1.0",
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-transform-modules-commonjs": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-modules-commonjs/download/@babel/plugin-transform-modules-commonjs-7.2.0.tgz",
      "integrity": "sha1-xPGTP1mR1RRenPrR39hI6hcn9AQ=",
      "dev": true,
      "requires": {
        "@babel/helper-module-transforms": "^7.1.0",
        "@babel/helper-plugin-utils": "^7.0.0",
        "@babel/helper-simple-access": "^7.1.0"
      }
    },
    "@babel/plugin-transform-modules-systemjs": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-modules-systemjs/download/@babel/plugin-transform-modules-systemjs-7.2.0.tgz",
      "integrity": "sha1-kSv+nl/5gpJMgdCTfJLSSZS7kGg=",
      "dev": true,
      "requires": {
        "@babel/helper-hoist-variables": "^7.0.0",
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-transform-modules-umd": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-modules-umd/download/@babel/plugin-transform-modules-umd-7.2.0.tgz",
      "integrity": "sha1-dnjOdRafCHe46yI1U4wHQmjdAa4=",
      "dev": true,
      "requires": {
        "@babel/helper-module-transforms": "^7.1.0",
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-transform-new-target": {
      "version": "7.0.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-new-target/download/@babel/plugin-transform-new-target-7.0.0.tgz",
      "integrity": "sha1-ro+9iVF/p4ktIOZWTmQeh3DDqko=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-transform-object-super": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-object-super/download/@babel/plugin-transform-object-super-7.2.0.tgz",
      "integrity": "sha1-s11MEPVrq11lAEfa0PHY6IFLZZg=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0",
        "@babel/helper-replace-supers": "^7.1.0"
      }
    },
    "@babel/plugin-transform-parameters": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-parameters/download/@babel/plugin-transform-parameters-7.2.0.tgz",
      "integrity": "sha1-DVrRXcgF4uqGbfTdZoK/520UCMI=",
      "dev": true,
      "requires": {
        "@babel/helper-call-delegate": "^7.1.0",
        "@babel/helper-get-function-arity": "^7.0.0",
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-transform-regenerator": {
      "version": "7.0.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-regenerator/download/@babel/plugin-transform-regenerator-7.0.0.tgz",
      "integrity": "sha1-W0Foa07UC++HTX7WqEvdhJwT4ME=",
      "dev": true,
      "requires": {
        "regenerator-transform": "^0.13.3"
      }
    },
    "@babel/plugin-transform-runtime": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-runtime/download/@babel/plugin-transform-runtime-7.2.0.tgz",
      "integrity": "sha1-VmvEP30K7ciA6t29KRaNDySJZuo=",
      "dev": true,
      "requires": {
        "@babel/helper-module-imports": "^7.0.0",
        "@babel/helper-plugin-utils": "^7.0.0",
        "resolve": "^1.8.1",
        "semver": "^5.5.1"
      }
    },
    "@babel/plugin-transform-shorthand-properties": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-shorthand-properties/download/@babel/plugin-transform-shorthand-properties-7.2.0.tgz",
      "integrity": "sha1-YzOu4vjW7n4oYVRXKYk0o7RhmPA=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-transform-spread": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-spread/download/@babel/plugin-transform-spread-7.2.0.tgz",
      "integrity": "sha1-DHbBKjtYJhMAeO6OyEp6jkr9ecQ=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-transform-sticky-regex": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-sticky-regex/download/@babel/plugin-transform-sticky-regex-7.2.0.tgz",
      "integrity": "sha1-oeRUtZlVYKnB4NU338FQYf0mh+E=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0",
        "@babel/helper-regex": "^7.0.0"
      }
    },
    "@babel/plugin-transform-template-literals": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-template-literals/download/@babel/plugin-transform-template-literals-7.2.0.tgz",
      "integrity": "sha1-2H7QG46qx6kkc/YIyXwIneK6Hls=",
      "dev": true,
      "requires": {
        "@babel/helper-annotate-as-pure": "^7.0.0",
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-transform-typeof-symbol": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-typeof-symbol/download/@babel/plugin-transform-typeof-symbol-7.2.0.tgz",
      "integrity": "sha1-EX0rzsL79ktLWdH5gZiUaC0p8rI=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0"
      }
    },
    "@babel/plugin-transform-unicode-regex": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/plugin-transform-unicode-regex/download/@babel/plugin-transform-unicode-regex-7.2.0.tgz",
      "integrity": "sha1-TrjbFvly+Ku1BiwWG4sRVUat4Is=",
      "dev": true,
      "requires": {
        "@babel/helper-plugin-utils": "^7.0.0",
        "@babel/helper-regex": "^7.0.0",
        "regexpu-core": "^4.1.3"
      }
    },
    "@babel/preset-env": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/preset-env/download/@babel/preset-env-7.2.0.tgz",
      "integrity": "sha1-pQMOfkMGr1opXdXXx43FRkrz/uI=",
      "dev": true,
      "requires": {
        "@babel/helper-module-imports": "^7.0.0",
        "@babel/helper-plugin-utils": "^7.0.0",
        "@babel/plugin-proposal-async-generator-functions": "^7.2.0",
        "@babel/plugin-proposal-json-strings": "^7.2.0",
        "@babel/plugin-proposal-object-rest-spread": "^7.2.0",
        "@babel/plugin-proposal-optional-catch-binding": "^7.2.0",
        "@babel/plugin-proposal-unicode-property-regex": "^7.2.0",
        "@babel/plugin-syntax-async-generators": "^7.2.0",
        "@babel/plugin-syntax-object-rest-spread": "^7.2.0",
        "@babel/plugin-syntax-optional-catch-binding": "^7.2.0",
        "@babel/plugin-transform-arrow-functions": "^7.2.0",
        "@babel/plugin-transform-async-to-generator": "^7.2.0",
        "@babel/plugin-transform-block-scoped-functions": "^7.2.0",
        "@babel/plugin-transform-block-scoping": "^7.2.0",
        "@babel/plugin-transform-classes": "^7.2.0",
        "@babel/plugin-transform-computed-properties": "^7.2.0",
        "@babel/plugin-transform-destructuring": "^7.2.0",
        "@babel/plugin-transform-dotall-regex": "^7.2.0",
        "@babel/plugin-transform-duplicate-keys": "^7.2.0",
        "@babel/plugin-transform-exponentiation-operator": "^7.2.0",
        "@babel/plugin-transform-for-of": "^7.2.0",
        "@babel/plugin-transform-function-name": "^7.2.0",
        "@babel/plugin-transform-literals": "^7.2.0",
        "@babel/plugin-transform-modules-amd": "^7.2.0",
        "@babel/plugin-transform-modules-commonjs": "^7.2.0",
        "@babel/plugin-transform-modules-systemjs": "^7.2.0",
        "@babel/plugin-transform-modules-umd": "^7.2.0",
        "@babel/plugin-transform-new-target": "^7.0.0",
        "@babel/plugin-transform-object-super": "^7.2.0",
        "@babel/plugin-transform-parameters": "^7.2.0",
        "@babel/plugin-transform-regenerator": "^7.0.0",
        "@babel/plugin-transform-shorthand-properties": "^7.2.0",
        "@babel/plugin-transform-spread": "^7.2.0",
        "@babel/plugin-transform-sticky-regex": "^7.2.0",
        "@babel/plugin-transform-template-literals": "^7.2.0",
        "@babel/plugin-transform-typeof-symbol": "^7.2.0",
        "@babel/plugin-transform-unicode-regex": "^7.2.0",
        "browserslist": "^4.3.4",
        "invariant": "^2.2.2",
        "js-levenshtein": "^1.1.3",
        "semver": "^5.3.0"
      }
    },
    "@babel/runtime": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/runtime/download/@babel/runtime-7.2.0.tgz",
      "integrity": "sha1-sD5C7t31iY4AZG5MhA+ge6jcrX8=",
      "dev": true,
      "requires": {
        "regenerator-runtime": "^0.12.0"
      }
    },
    "@babel/template": {
      "version": "7.1.2",
      "resolved": "http://registry.npm.taobao.org/@babel/template/download/@babel/template-7.1.2.tgz",
      "integrity": "sha1-CQSEpXT+9aLS13JqZ07O2lxbVkQ=",
      "dev": true,
      "requires": {
        "@babel/code-frame": "^7.0.0",
        "@babel/parser": "^7.1.2",
        "@babel/types": "^7.1.2"
      }
    },
    "@babel/traverse": {
      "version": "7.1.6",
      "resolved": "http://registry.npm.taobao.org/@babel/traverse/download/@babel/traverse-7.1.6.tgz",
      "integrity": "sha1-yNuZY6tM5biUIiQ1SCvY6oVLe1w=",
      "dev": true,
      "requires": {
        "@babel/code-frame": "^7.0.0",
        "@babel/generator": "^7.1.6",
        "@babel/helper-function-name": "^7.1.0",
        "@babel/helper-split-export-declaration": "^7.0.0",
        "@babel/parser": "^7.1.6",
        "@babel/types": "^7.1.6",
        "debug": "^4.1.0",
        "globals": "^11.1.0",
        "lodash": "^4.17.10"
      }
    },
    "@babel/types": {
      "version": "7.2.0",
      "resolved": "http://registry.npm.taobao.org/@babel/types/download/@babel/types-7.2.0.tgz",
      "integrity": "sha1-eUHFstgGDgb5YB1r58Ij7vkG1dg=",
      "dev": true,
      "requires": {
        "esutils": "^2.0.2",
        "lodash": "^4.17.10",
        "to-fast-properties": "^2.0.0"
      }
    },
    "@intervolga/optimize-cssnano-plugin": {
      "version": "1.0.6",
      "resolved": "http://registry.npm.taobao.org/@intervolga/optimize-cssnano-plugin/download/@intervolga/optimize-cssnano-plugin-1.0.6.tgz",
      "integrity": "sha1-vnx4RhKLiPapsdEmGgrQbrXA/fg=",
      "dev": true,
      "requires": {
        "cssnano": "^4.0.0",
        "cssnano-preset-default": "^4.0.0",
        "postcss": "^7.0.0"
      }
    },
    "@mrmlnc/readdir-enhanced": {
      "version": "2.2.1",
      "resolved": "http://registry.npm.taobao.org/@mrmlnc/readdir-enhanced/download/@mrmlnc/readdir-enhanced-2.2.1.tgz",
      "integrity": "sha1-UkryQNGjYFJ7cwR17PoTRKpUDd4=",
      "dev": true,
      "requires": {
        "call-me-maybe": "^1.0.1",
        "glob-to-regexp": "^0.3.0"
      }
    },
    "@nodelib/fs.stat": {
      "version": "1.1.3",
      "resolved": "http://registry.npm.taobao.org/@nodelib/fs.stat/download/@nodelib/fs.stat-1.1.3.tgz",
      "integrity": "sha1-K1o6s/kYzKSKjHVMCBaOPwPrphs=",
      "dev": true
    },
    "@types/q": {
      "version": "1.5.1",
      "resolved": "http://registry.npm.taobao.org/@types/q/download/@types/q-1.5.1.tgz",
      "integrity": "sha1-SP2YwVYf5xi2FzPa7Ub/EVtJbhg=",
      "dev": true
    },
    "@vue/babel-preset-app": {
      "version": "3.2.0",
      "resolved": "http://registry.npm.taobao.org/@vue/babel-preset-app/download/@vue/babel-preset-app-3.2.0.tgz",
      "integrity": "sha1-pEOs29NPZtdkXbJx2axY++T+hw0=",
      "dev": true,
      "requires": {
        "@babel/plugin-proposal-class-properties": "^7.0.0",
        "@babel/plugin-proposal-decorators": "^7.1.0",
        "@babel/plugin-syntax-dynamic-import": "^7.0.0",
        "@babel/plugin-syntax-jsx": "^7.0.0",
        "@babel/plugin-transform-runtime": "^7.0.0",
        "@babel/preset-env": "^7.0.0",
        "@babel/runtime": "^7.0.0",
        "babel-helper-vue-jsx-merge-props": "^2.0.3",
        "babel-plugin-dynamic-import-node": "^2.2.0",
        "babel-plugin-transform-vue-jsx": "^4.0.1",
        "core-js": "^2.5.7"
      }
    },
    "@vue/cli-overlay": {
      "version": "3.2.0",
      "resolved": "http://registry.npm.taobao.org/@vue/cli-overlay/download/@vue/cli-overlay-3.2.0.tgz",
      "integrity": "sha1-u114GRS7Wvl9kkELq7qjcgcHtyg=",
      "dev": true
    },
    "@vue/cli-plugin-babel": {
      "version": "3.2.0",
      "resolved": "http://registry.npm.taobao.org/@vue/cli-plugin-babel/download/@vue/cli-plugin-babel-3.2.0.tgz",
      "integrity": "sha1-VUjgUvB1EhVpQvUNzxiUitKb5+w=",
      "dev": true,
      "requires": {
        "@babel/core": "^7.0.0",
        "@vue/babel-preset-app": "^3.2.0",
        "babel-loader": "^8.0.4"
      }
    },
    "@vue/cli-plugin-eslint": {
      "version": "3.2.1",
      "resolved": "http://registry.npm.taobao.org/@vue/cli-plugin-eslint/download/@vue/cli-plugin-eslint-3.2.1.tgz",
      "integrity": "sha1-TcNTrdkwIzY79MazR+ZEctHy9DI=",
      "dev": true,
      "requires": {
        "@vue/cli-shared-utils": "^3.2.0",
        "babel-eslint": "^10.0.1",
        "eslint": "^4.19.1",
        "eslint-loader": "^2.1.1",
        "eslint-plugin-vue": "^4.7.1",
        "globby": "^8.0.1"
      },
      "dependencies": {
        "ajv": {
          "version": "5.5.2",
          "resolved": "http://registry.npm.taobao.org/ajv/download/ajv-5.5.2.tgz",
          "integrity": "sha1-c7Xuyj+rZT49P5Qis0GtQiBdyWU=",
          "dev": true,
          "requires": {
            "co": "^4.6.0",
            "fast-deep-equal": "^1.0.0",
            "fast-json-stable-stringify": "^2.0.0",
            "json-schema-traverse": "^0.3.0"
          }
        },
        "cross-spawn": {
          "version": "5.1.0",
          "resolved": "http://registry.npm.taobao.org/cross-spawn/download/cross-spawn-5.1.0.tgz",
          "integrity": "sha1-6L0O/uWPz/b4+UUQoKVUu/ojVEk=",
          "dev": true,
          "requires": {
            "lru-cache": "^4.0.1",
            "shebang-command": "^1.2.0",
            "which": "^1.2.9"
          }
        },
        "debug": {
          "version": "3.2.6",
          "resolved": "http://registry.npm.taobao.org/debug/download/debug-3.2.6.tgz",
          "integrity": "sha1-6D0X3hbYp++3cX7b5fsQE17uYps=",
          "dev": true,
          "requires": {
            "ms": "^2.1.1"
          }
        },
        "eslint": {
          "version": "4.19.1",
          "resolved": "http://registry.npm.taobao.org/eslint/download/eslint-4.19.1.tgz",
          "integrity": "sha1-MtHWU+HZBAiFS/spbwdux+GGowA=",
          "dev": true,
          "requires": {
            "ajv": "^5.3.0",
            "babel-code-frame": "^6.22.0",
            "chalk": "^2.1.0",
            "concat-stream": "^1.6.0",
            "cross-spawn": "^5.1.0",
            "debug": "^3.1.0",
            "doctrine": "^2.1.0",
            "eslint-scope": "^3.7.1",
            "eslint-visitor-keys": "^1.0.0",
            "espree": "^3.5.4",
            "esquery": "^1.0.0",
            "esutils": "^2.0.2",
            "file-entry-cache": "^2.0.0",
            "functional-red-black-tree": "^1.0.1",
            "glob": "^7.1.2",
            "globals": "^11.0.1",
            "ignore": "^3.3.3",
            "imurmurhash": "^0.1.4",
            "inquirer": "^3.0.6",
            "is-resolvable": "^1.0.0",
            "js-yaml": "^3.9.1",
            "json-stable-stringify-without-jsonify": "^1.0.1",
            "levn": "^0.3.0",
            "lodash": "^4.17.4",
            "minimatch": "^3.0.2",
            "mkdirp": "^0.5.1",
            "natural-compare": "^1.4.0",
            "optionator": "^0.8.2",
            "path-is-inside": "^1.0.2",
            "pluralize": "^7.0.0",
            "progress": "^2.0.0",
            "regexpp": "^1.0.1",
            "require-uncached": "^1.0.3",
            "semver": "^5.3.0",
            "strip-ansi": "^4.0.0",
            "strip-json-comments": "~2.0.1",
            "table": "4.0.2",
            "text-table": "~0.2.0"
          }
        },
        "eslint-plugin-vue": {
          "version": "4.7.1",
          "resolved": "http://registry.npm.taobao.org/eslint-plugin-vue/download/eslint-plugin-vue-4.7.1.tgz",
          "integrity": "sha1-yCm5/GJYLBiXtaC5Sv1E7MpRHmM=",
          "dev": true,
          "requires": {
            "vue-eslint-parser": "^2.0.3"
          }
        },
        "fast-deep-equal": {
          "version": "1.1.0",
          "resolved": "http://registry.npm.taobao.org/fast-deep-equal/download/fast-deep-equal-1.1.0.tgz",
          "integrity": "sha1-wFNHeBfIa1HaqFPIHgWbcz0CNhQ=",
          "dev": true
        },
        "json-schema-traverse": {
          "version": "0.3.1",
          "resolved": "http://registry.npm.taobao.org/json-schema-traverse/download/json-schema-traverse-0.3.1.tgz",
          "integrity": "sha1-NJptRMU6Ud6JtAgFxdXlm0F9M0A=",
          "dev": true
        }
      }
    },
    "@vue/cli-service": {
      "version": "3.2.0",
      "resolved": "http://registry.npm.taobao.org/@vue/cli-service/download/@vue/cli-service-3.2.0.tgz",
      "integrity": "sha1-lujuezOpEatxcQwSarVc1kwqTFE=",
      "dev": true,
      "requires": {
        "@intervolga/optimize-cssnano-plugin": "^1.0.5",
        "@vue/cli-overlay": "^3.2.0",
        "@vue/cli-shared-utils": "^3.2.0",
        "@vue/preload-webpack-plugin": "^1.1.0",
        "@vue/web-component-wrapper": "^1.2.0",
        "acorn": "^6.0.4",
        "acorn-walk": "^6.1.1",
        "address": "^1.0.3",
        "autoprefixer": "^8.6.5",
        "cache-loader": "^1.2.5",
        "case-sensitive-paths-webpack-plugin": "^2.1.2",
        "chalk": "^2.4.1",
        "clipboardy": "^1.2.3",
        "cliui": "^4.1.0",
        "copy-webpack-plugin": "^4.6.0",
        "css-loader": "^1.0.1",
        "cssnano": "^4.1.7",
        "debug": "^4.1.0",
        "escape-string-regexp": "^1.0.5",
        "file-loader": "^2.0.0",
        "friendly-errors-webpack-plugin": "^1.7.0",
        "fs-extra": "^7.0.1",
        "globby": "^8.0.1",
        "hash-sum": "^1.0.2",
        "html-webpack-plugin": "^3.2.0",
        "launch-editor-middleware": "^2.2.1",
        "lodash.defaultsdeep": "^4.6.0",
        "lodash.mapvalues": "^4.6.0",
        "lodash.transform": "^4.6.0",
        "mini-css-extract-plugin": "^0.4.5",
        "minimist": "^1.2.0",
        "ora": "^3.0.0",
        "portfinder": "^1.0.19",
        "postcss-loader": "^3.0.0",
        "read-pkg": "^4.0.1",
        "semver": "^5.6.0",
        "slash": "^2.0.0",
        "source-map-url": "^0.4.0",
        "ssri": "^6.0.1",
        "string.prototype.padend": "^3.0.0",
        "terser-webpack-plugin": "^1.1.0",
        "thread-loader": "^1.2.0",
        "url-loader": "^1.1.2",
        "vue-loader": "^15.4.2",
        "webpack": "^4.26.1",
        "webpack-bundle-analyzer": "^3.0.3",
        "webpack-chain": "^4.11.0",
        "webpack-dev-server": "^3.1.10",
        "webpack-merge": "^4.1.4",
        "yorkie": "^2.0.0"
      },
      "dependencies": {
        "acorn": {
          "version": "6.0.4",
          "resolved": "http://registry.npm.taobao.org/acorn/download/acorn-6.0.4.tgz",
          "integrity": "sha1-dzd+c1O3LsUQRVCqLSCXov1At1Q=",
          "dev": true
        },
        "ora": {
          "version": "3.0.0",
          "resolved": "http://registry.npm.taobao.org/ora/download/ora-3.0.0.tgz",
          "integrity": "sha1-gXnjUluar9mSQtY8wgb9ZHMnQdA=",
          "dev": true,
          "requires": {
            "chalk": "^2.3.1",
            "cli-cursor": "^2.1.0",
            "cli-spinners": "^1.1.0",
            "log-symbols": "^2.2.0",
            "strip-ansi": "^4.0.0",
            "wcwidth": "^1.0.1"
          }
        },
        "slash": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/slash/download/slash-2.0.0.tgz",
          "integrity": "sha1-3lUoUaF1nfOo8gZTVEL17E3eq0Q=",
          "dev": true
        }
      }
    },
    "@vue/cli-shared-utils": {
      "version": "3.2.0",
      "resolved": "http://registry.npm.taobao.org/@vue/cli-shared-utils/download/@vue/cli-shared-utils-3.2.0.tgz",
      "integrity": "sha1-UEA3BjwqSjRtw1yWUr04Y9pIFuc=",
      "dev": true,
      "requires": {
        "chalk": "^2.4.1",
        "execa": "^1.0.0",
        "joi": "^13.0.0",
        "launch-editor": "^2.2.1",
        "lru-cache": "^4.1.3",
        "node-ipc": "^9.1.1",
        "opn": "^5.3.0",
        "ora": "^2.1.0",
        "request": "^2.87.0",
        "request-promise-native": "^1.0.5",
        "semver": "^5.5.0",
        "string.prototype.padstart": "^3.0.0"
      }
    },
    "@vue/component-compiler-utils": {
      "version": "2.3.1",
      "resolved": "http://registry.npm.taobao.org/@vue/component-compiler-utils/download/@vue/component-compiler-utils-2.3.1.tgz",
      "integrity": "sha1-0cJiPwKtP+a2/Jw3Yr5VycYeOXc=",
      "dev": true,
      "requires": {
        "consolidate": "^0.15.1",
        "hash-sum": "^1.0.2",
        "lru-cache": "^4.1.2",
        "merge-source-map": "^1.1.0",
        "postcss": "^6.0.20",
        "postcss-selector-parser": "^3.1.1",
        "prettier": "1.13.7",
        "source-map": "^0.5.6",
        "vue-template-es2015-compiler": "^1.6.0"
      },
      "dependencies": {
        "postcss": {
          "version": "6.0.23",
          "resolved": "http://registry.npm.taobao.org/postcss/download/postcss-6.0.23.tgz",
          "integrity": "sha1-YcgswyisYOZ3ZF+XkFTrmLwOMyQ=",
          "dev": true,
          "requires": {
            "chalk": "^2.4.1",
            "source-map": "^0.6.1",
            "supports-color": "^5.4.0"
          },
          "dependencies": {
            "source-map": {
              "version": "0.6.1",
              "resolved": "http://registry.npm.taobao.org/source-map/download/source-map-0.6.1.tgz",
              "integrity": "sha1-dHIq8y6WFOnCh6jQu95IteLxomM=",
              "dev": true
            }
          }
        },
        "postcss-selector-parser": {
          "version": "3.1.1",
          "resolved": "http://registry.npm.taobao.org/postcss-selector-parser/download/postcss-selector-parser-3.1.1.tgz",
          "integrity": "sha1-T4dfSvsMllc9XPTXQBGu4lCn6GU=",
          "dev": true,
          "requires": {
            "dot-prop": "^4.1.1",
            "indexes-of": "^1.0.1",
            "uniq": "^1.0.1"
          }
        }
      }
    },
    "@vue/preload-webpack-plugin": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/@vue/preload-webpack-plugin/download/@vue/preload-webpack-plugin-1.1.0.tgz",
      "integrity": "sha1-12jboAQmHAKbU6d8XqLV+e5PPM4=",
      "dev": true
    },
    "@vue/web-component-wrapper": {
      "version": "1.2.0",
      "resolved": "http://registry.npm.taobao.org/@vue/web-component-wrapper/download/@vue/web-component-wrapper-1.2.0.tgz",
      "integrity": "sha1-uw5G8VhafiibTuYGfcxaauYvHdE=",
      "dev": true
    },
    "@webassemblyjs/ast": {
      "version": "1.7.11",
      "resolved": "http://registry.npm.taobao.org/@webassemblyjs/ast/download/@webassemblyjs/ast-1.7.11.tgz",
      "integrity": "sha1-uYhYLK+7Kwlei1VlJvMMkNBXys4=",
      "dev": true,
      "requires": {
        "@webassemblyjs/helper-module-context": "1.7.11",
        "@webassemblyjs/helper-wasm-bytecode": "1.7.11",
        "@webassemblyjs/wast-parser": "1.7.11"
      }
    },
    "@webassemblyjs/floating-point-hex-parser": {
      "version": "1.7.11",
      "resolved": "http://registry.npm.taobao.org/@webassemblyjs/floating-point-hex-parser/download/@webassemblyjs/floating-point-hex-parser-1.7.11.tgz",
      "integrity": "sha1-pp8K9lAuuaPARVVbGmEp09Py4xM=",
      "dev": true
    },
    "@webassemblyjs/helper-api-error": {
      "version": "1.7.11",
      "resolved": "http://registry.npm.taobao.org/@webassemblyjs/helper-api-error/download/@webassemblyjs/helper-api-error-1.7.11.tgz",
      "integrity": "sha1-x7a7gQX4QDlRGis5zklPGTgYoyo=",
      "dev": true
    },
    "@webassemblyjs/helper-buffer": {
      "version": "1.7.11",
      "resolved": "http://registry.npm.taobao.org/@webassemblyjs/helper-buffer/download/@webassemblyjs/helper-buffer-1.7.11.tgz",
      "integrity": "sha1-MSLUjcxslFbtmC3r4WyPNxAd85s=",
      "dev": true
    },
    "@webassemblyjs/helper-code-frame": {
      "version": "1.7.11",
      "resolved": "http://registry.npm.taobao.org/@webassemblyjs/helper-code-frame/download/@webassemblyjs/helper-code-frame-1.7.11.tgz",
      "integrity": "sha1-z48QbnRmYqDaKb3vY1/NPRJINks=",
      "dev": true,
      "requires": {
        "@webassemblyjs/wast-printer": "1.7.11"
      }
    },
    "@webassemblyjs/helper-fsm": {
      "version": "1.7.11",
      "resolved": "http://registry.npm.taobao.org/@webassemblyjs/helper-fsm/download/@webassemblyjs/helper-fsm-1.7.11.tgz",
      "integrity": "sha1-3ziIKmJAgNA/dQP5Pj8XrFrAEYE=",
      "dev": true
    },
    "@webassemblyjs/helper-module-context": {
      "version": "1.7.11",
      "resolved": "http://registry.npm.taobao.org/@webassemblyjs/helper-module-context/download/@webassemblyjs/helper-module-context-1.7.11.tgz",
      "integrity": "sha1-2HTXIuUeYqwgJHaTXWScgC+g4gk=",
      "dev": true
    },
    "@webassemblyjs/helper-wasm-bytecode": {
      "version": "1.7.11",
      "resolved": "http://registry.npm.taobao.org/@webassemblyjs/helper-wasm-bytecode/download/@webassemblyjs/helper-wasm-bytecode-1.7.11.tgz",
      "integrity": "sha1-3ZoegX8cLrEFtM8QEwk8ufPJywY=",
      "dev": true
    },
    "@webassemblyjs/helper-wasm-section": {
      "version": "1.7.11",
      "resolved": "http://registry.npm.taobao.org/@webassemblyjs/helper-wasm-section/download/@webassemblyjs/helper-wasm-section-1.7.11.tgz",
      "integrity": "sha1-nJrEHs+fvP/8lvbSZ14t4zgR5oo=",
      "dev": true,
      "requires": {
        "@webassemblyjs/ast": "1.7.11",
        "@webassemblyjs/helper-buffer": "1.7.11",
        "@webassemblyjs/helper-wasm-bytecode": "1.7.11",
        "@webassemblyjs/wasm-gen": "1.7.11"
      }
    },
    "@webassemblyjs/ieee754": {
      "version": "1.7.11",
      "resolved": "http://registry.npm.taobao.org/@webassemblyjs/ieee754/download/@webassemblyjs/ieee754-1.7.11.tgz",
      "integrity": "sha1-yVg562N1ejGICq7HtlEtQZGsZAs=",
      "dev": true,
      "requires": {
        "@xtuc/ieee754": "^1.2.0"
      }
    },
    "@webassemblyjs/leb128": {
      "version": "1.7.11",
      "resolved": "http://registry.npm.taobao.org/@webassemblyjs/leb128/download/@webassemblyjs/leb128-1.7.11.tgz",
      "integrity": "sha1-1yZ6HunEWU/T9+NymIGOxlaH22M=",
      "dev": true,
      "requires": {
        "@xtuc/long": "4.2.1"
      }
    },
    "@webassemblyjs/utf8": {
      "version": "1.7.11",
      "resolved": "http://registry.npm.taobao.org/@webassemblyjs/utf8/download/@webassemblyjs/utf8-1.7.11.tgz",
      "integrity": "sha1-Btchjqn9yUpnk6qSIIFg2z0m7oI=",
      "dev": true
    },
    "@webassemblyjs/wasm-edit": {
      "version": "1.7.11",
      "resolved": "http://registry.npm.taobao.org/@webassemblyjs/wasm-edit/download/@webassemblyjs/wasm-edit-1.7.11.tgz",
      "integrity": "sha1-jHTKR01PlR0B266b1wgU7iKoIAU=",
      "dev": true,
      "requires": {
        "@webassemblyjs/ast": "1.7.11",
        "@webassemblyjs/helper-buffer": "1.7.11",
        "@webassemblyjs/helper-wasm-bytecode": "1.7.11",
        "@webassemblyjs/helper-wasm-section": "1.7.11",
        "@webassemblyjs/wasm-gen": "1.7.11",
        "@webassemblyjs/wasm-opt": "1.7.11",
        "@webassemblyjs/wasm-parser": "1.7.11",
        "@webassemblyjs/wast-printer": "1.7.11"
      }
    },
    "@webassemblyjs/wasm-gen": {
      "version": "1.7.11",
      "resolved": "http://registry.npm.taobao.org/@webassemblyjs/wasm-gen/download/@webassemblyjs/wasm-gen-1.7.11.tgz",
      "integrity": "sha1-m7upQvIjdWhqb7dZr816ycRdoag=",
      "dev": true,
      "requires": {
        "@webassemblyjs/ast": "1.7.11",
        "@webassemblyjs/helper-wasm-bytecode": "1.7.11",
        "@webassemblyjs/ieee754": "1.7.11",
        "@webassemblyjs/leb128": "1.7.11",
        "@webassemblyjs/utf8": "1.7.11"
      }
    },
    "@webassemblyjs/wasm-opt": {
      "version": "1.7.11",
      "resolved": "http://registry.npm.taobao.org/@webassemblyjs/wasm-opt/download/@webassemblyjs/wasm-opt-1.7.11.tgz",
      "integrity": "sha1-szHo5874+OLwB9QsOjagWAp9bKc=",
      "dev": true,
      "requires": {
        "@webassemblyjs/ast": "1.7.11",
        "@webassemblyjs/helper-buffer": "1.7.11",
        "@webassemblyjs/wasm-gen": "1.7.11",
        "@webassemblyjs/wasm-parser": "1.7.11"
      }
    },
    "@webassemblyjs/wasm-parser": {
      "version": "1.7.11",
      "resolved": "http://registry.npm.taobao.org/@webassemblyjs/wasm-parser/download/@webassemblyjs/wasm-parser-1.7.11.tgz",
      "integrity": "sha1-bj0g+mo1GfawhO+Tka1YIR77Cho=",
      "dev": true,
      "requires": {
        "@webassemblyjs/ast": "1.7.11",
        "@webassemblyjs/helper-api-error": "1.7.11",
        "@webassemblyjs/helper-wasm-bytecode": "1.7.11",
        "@webassemblyjs/ieee754": "1.7.11",
        "@webassemblyjs/leb128": "1.7.11",
        "@webassemblyjs/utf8": "1.7.11"
      }
    },
    "@webassemblyjs/wast-parser": {
      "version": "1.7.11",
      "resolved": "http://registry.npm.taobao.org/@webassemblyjs/wast-parser/download/@webassemblyjs/wast-parser-1.7.11.tgz",
      "integrity": "sha1-Jb0RdWLKjAAnIP+BFu+QctnKhpw=",
      "dev": true,
      "requires": {
        "@webassemblyjs/ast": "1.7.11",
        "@webassemblyjs/floating-point-hex-parser": "1.7.11",
        "@webassemblyjs/helper-api-error": "1.7.11",
        "@webassemblyjs/helper-code-frame": "1.7.11",
        "@webassemblyjs/helper-fsm": "1.7.11",
        "@xtuc/long": "4.2.1"
      }
    },
    "@webassemblyjs/wast-printer": {
      "version": "1.7.11",
      "resolved": "http://registry.npm.taobao.org/@webassemblyjs/wast-printer/download/@webassemblyjs/wast-printer-1.7.11.tgz",
      "integrity": "sha1-xCRbbeJCy1CizJUBdP2/ZceNeBM=",
      "dev": true,
      "requires": {
        "@webassemblyjs/ast": "1.7.11",
        "@webassemblyjs/wast-parser": "1.7.11",
        "@xtuc/long": "4.2.1"
      }
    },
    "@xtuc/ieee754": {
      "version": "1.2.0",
      "resolved": "http://registry.npm.taobao.org/@xtuc/ieee754/download/@xtuc/ieee754-1.2.0.tgz",
      "integrity": "sha1-7vAUoxRa5Hehy8AM0eVSM23Ot5A=",
      "dev": true
    },
    "@xtuc/long": {
      "version": "4.2.1",
      "resolved": "http://registry.npm.taobao.org/@xtuc/long/download/@xtuc/long-4.2.1.tgz",
      "integrity": "sha1-XIXWYvdvodNFdXZsXc1mFavNMNg=",
      "dev": true
    },
    "accepts": {
      "version": "1.3.5",
      "resolved": "http://registry.npm.taobao.org/accepts/download/accepts-1.3.5.tgz",
      "integrity": "sha1-63d99gEXI6OxTopywIBcjoZ0a9I=",
      "dev": true,
      "requires": {
        "mime-types": "~2.1.18",
        "negotiator": "0.6.1"
      }
    },
    "acorn": {
      "version": "5.7.3",
      "resolved": "http://registry.npm.taobao.org/acorn/download/acorn-5.7.3.tgz",
      "integrity": "sha1-Z6ojG/iBKXS4UjWpZ3Hra9B+onk=",
      "dev": true
    },
    "acorn-dynamic-import": {
      "version": "3.0.0",
      "resolved": "http://registry.npm.taobao.org/acorn-dynamic-import/download/acorn-dynamic-import-3.0.0.tgz",
      "integrity": "sha1-kBzu5Mf6rvfgetKkfokGddpQong=",
      "dev": true,
      "requires": {
        "acorn": "^5.0.0"
      }
    },
    "acorn-jsx": {
      "version": "3.0.1",
      "resolved": "http://registry.npm.taobao.org/acorn-jsx/download/acorn-jsx-3.0.1.tgz",
      "integrity": "sha1-r9+UiPsezvyDSPb7IvRk4ypYs2s=",
      "dev": true,
      "requires": {
        "acorn": "^3.0.4"
      },
      "dependencies": {
        "acorn": {
          "version": "3.3.0",
          "resolved": "http://registry.npm.taobao.org/acorn/download/acorn-3.3.0.tgz",
          "integrity": "sha1-ReN/s56No/JbruP/U2niu18iAXo=",
          "dev": true
        }
      }
    },
    "acorn-walk": {
      "version": "6.1.1",
      "resolved": "http://registry.npm.taobao.org/acorn-walk/download/acorn-walk-6.1.1.tgz",
      "integrity": "sha1-02O2b1+sXwGP+cOh57b44xDMORM=",
      "dev": true
    },
    "address": {
      "version": "1.0.3",
      "resolved": "http://registry.npm.taobao.org/address/download/address-1.0.3.tgz",
      "integrity": "sha1-tfUGMfjWzsi9IMljljr7VeBsvOk=",
      "dev": true
    },
    "ajv": {
      "version": "6.6.1",
      "resolved": "http://registry.npm.taobao.org/ajv/download/ajv-6.6.1.tgz",
      "integrity": "sha1-Y2D17Q2A8jLMKylMNi1dwuU43WE=",
      "dev": true,
      "requires": {
        "fast-deep-equal": "^2.0.1",
        "fast-json-stable-stringify": "^2.0.0",
        "json-schema-traverse": "^0.4.1",
        "uri-js": "^4.2.2"
      }
    },
    "ajv-errors": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/ajv-errors/download/ajv-errors-1.0.1.tgz",
      "integrity": "sha1-81mGrOuRr63sQQL72FAUlQzvpk0=",
      "dev": true
    },
    "ajv-keywords": {
      "version": "2.1.1",
      "resolved": "http://registry.npm.taobao.org/ajv-keywords/download/ajv-keywords-2.1.1.tgz",
      "integrity": "sha1-YXmX/F9gV2iUxDX5QNgZ4TW4B2I=",
      "dev": true
    },
    "alphanum-sort": {
      "version": "1.0.2",
      "resolved": "http://registry.npm.taobao.org/alphanum-sort/download/alphanum-sort-1.0.2.tgz",
      "integrity": "sha1-l6ERlkmyEa0zaR2fn0hqjsn74KM=",
      "dev": true
    },
    "ansi-colors": {
      "version": "3.2.3",
      "resolved": "http://registry.npm.taobao.org/ansi-colors/download/ansi-colors-3.2.3.tgz",
      "integrity": "sha1-V9NbhoboUeLMBMQD8cACA5dqGBM=",
      "dev": true
    },
    "ansi-escapes": {
      "version": "3.1.0",
      "resolved": "http://registry.npm.taobao.org/ansi-escapes/download/ansi-escapes-3.1.0.tgz",
      "integrity": "sha1-9zIHu4EgfXX9bIPxJa8m7qN4yjA=",
      "dev": true
    },
    "ansi-html": {
      "version": "0.0.7",
      "resolved": "http://registry.npm.taobao.org/ansi-html/download/ansi-html-0.0.7.tgz",
      "integrity": "sha1-gTWEAhliqenm/QOflA0S9WynhZ4=",
      "dev": true
    },
    "ansi-regex": {
      "version": "3.0.0",
      "resolved": "http://registry.npm.taobao.org/ansi-regex/download/ansi-regex-3.0.0.tgz",
      "integrity": "sha1-7QMXwyIGT3lGbAKWa922Bas32Zg=",
      "dev": true
    },
    "ansi-styles": {
      "version": "3.2.1",
      "resolved": "http://registry.npm.taobao.org/ansi-styles/download/ansi-styles-3.2.1.tgz",
      "integrity": "sha1-QfuyAkPlCxK+DwS43tvwdSDOhB0=",
      "dev": true,
      "requires": {
        "color-convert": "^1.9.0"
      }
    },
    "anymatch": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/anymatch/download/anymatch-2.0.0.tgz",
      "integrity": "sha1-vLJLTzeTTZqnrBe0ra+J58du8us=",
      "dev": true,
      "requires": {
        "micromatch": "^3.1.4",
        "normalize-path": "^2.1.1"
      }
    },
    "aproba": {
      "version": "1.2.0",
      "resolved": "http://registry.npm.taobao.org/aproba/download/aproba-1.2.0.tgz",
      "integrity": "sha1-aALmJk79GMeQobDVF/DyYnvyyUo=",
      "dev": true
    },
    "arch": {
      "version": "2.1.1",
      "resolved": "http://registry.npm.taobao.org/arch/download/arch-2.1.1.tgz",
      "integrity": "sha1-j1wnMao1owkpIhuwZA7tZRdeyE4=",
      "dev": true
    },
    "argparse": {
      "version": "1.0.10",
      "resolved": "http://registry.npm.taobao.org/argparse/download/argparse-1.0.10.tgz",
      "integrity": "sha1-vNZ5HqWuCXJeF+WtmIE0zUCz2RE=",
      "dev": true,
      "requires": {
        "sprintf-js": "~1.0.2"
      }
    },
    "arr-diff": {
      "version": "4.0.0",
      "resolved": "http://registry.npm.taobao.org/arr-diff/download/arr-diff-4.0.0.tgz",
      "integrity": "sha1-1kYQdP6/7HHn4VI1dhoyml3HxSA=",
      "dev": true
    },
    "arr-flatten": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/arr-flatten/download/arr-flatten-1.1.0.tgz",
      "integrity": "sha1-NgSLv/TntH4TZkQxbJlmnqWukfE=",
      "dev": true
    },
    "arr-union": {
      "version": "3.1.0",
      "resolved": "http://registry.npm.taobao.org/arr-union/download/arr-union-3.1.0.tgz",
      "integrity": "sha1-45sJrqne+Gao8gbiiK9jkZuuOcQ=",
      "dev": true
    },
    "array-filter": {
      "version": "0.0.1",
      "resolved": "http://registry.npm.taobao.org/array-filter/download/array-filter-0.0.1.tgz",
      "integrity": "sha1-fajPLiZijtcygDWB/SH2fKzS7uw=",
      "dev": true
    },
    "array-flatten": {
      "version": "1.1.1",
      "resolved": "http://registry.npm.taobao.org/array-flatten/download/array-flatten-1.1.1.tgz",
      "integrity": "sha1-ml9pkFGx5wczKPKgCJaLZOopVdI=",
      "dev": true
    },
    "array-map": {
      "version": "0.0.0",
      "resolved": "http://registry.npm.taobao.org/array-map/download/array-map-0.0.0.tgz",
      "integrity": "sha1-iKK6tz0c97zVwbEYoAP2b2ZfpmI=",
      "dev": true
    },
    "array-reduce": {
      "version": "0.0.0",
      "resolved": "http://registry.npm.taobao.org/array-reduce/download/array-reduce-0.0.0.tgz",
      "integrity": "sha1-FziZ0//Rx9k4PkR5Ul2+J4yrXys=",
      "dev": true
    },
    "array-union": {
      "version": "1.0.2",
      "resolved": "http://registry.npm.taobao.org/array-union/download/array-union-1.0.2.tgz",
      "integrity": "sha1-mjRBDk9OPaI96jdb5b5w8kd47Dk=",
      "dev": true,
      "requires": {
        "array-uniq": "^1.0.1"
      }
    },
    "array-uniq": {
      "version": "1.0.3",
      "resolved": "http://registry.npm.taobao.org/array-uniq/download/array-uniq-1.0.3.tgz",
      "integrity": "sha1-r2rId6Jcx/dOBYiUdThY39sk/bY=",
      "dev": true
    },
    "array-unique": {
      "version": "0.3.2",
      "resolved": "http://registry.npm.taobao.org/array-unique/download/array-unique-0.3.2.tgz",
      "integrity": "sha1-qJS3XUvE9s1nnvMkSp/Y9Gri1Cg=",
      "dev": true
    },
    "arrify": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/arrify/download/arrify-1.0.1.tgz",
      "integrity": "sha1-iYUI2iIm84DfkEcoRWhJwVAaSw0=",
      "dev": true
    },
    "asn1": {
      "version": "0.2.4",
      "resolved": "http://registry.npm.taobao.org/asn1/download/asn1-0.2.4.tgz",
      "integrity": "sha1-jSR136tVO7M+d7VOWeiAu4ziMTY=",
      "dev": true,
      "requires": {
        "safer-buffer": "~2.1.0"
      }
    },
    "asn1.js": {
      "version": "4.10.1",
      "resolved": "http://registry.npm.taobao.org/asn1.js/download/asn1.js-4.10.1.tgz",
      "integrity": "sha1-ucK/WAXx5kqt7tbfOiv6+1pz9aA=",
      "dev": true,
      "requires": {
        "bn.js": "^4.0.0",
        "inherits": "^2.0.1",
        "minimalistic-assert": "^1.0.0"
      }
    },
    "assert": {
      "version": "1.4.1",
      "resolved": "http://registry.npm.taobao.org/assert/download/assert-1.4.1.tgz",
      "integrity": "sha1-mZEtWRg2tab1s0XA8H7vwI/GXZE=",
      "dev": true,
      "requires": {
        "util": "0.10.3"
      },
      "dependencies": {
        "inherits": {
          "version": "2.0.1",
          "resolved": "http://registry.npm.taobao.org/inherits/download/inherits-2.0.1.tgz",
          "integrity": "sha1-sX0I0ya0Qj5Wjv9xn5GwscvfafE=",
          "dev": true
        },
        "util": {
          "version": "0.10.3",
          "resolved": "http://registry.npm.taobao.org/util/download/util-0.10.3.tgz",
          "integrity": "sha1-evsa/lCAUkZInj23/g7TeTNqwPk=",
          "dev": true,
          "requires": {
            "inherits": "2.0.1"
          }
        }
      }
    },
    "assert-plus": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/assert-plus/download/assert-plus-1.0.0.tgz",
      "integrity": "sha1-8S4PPF13sLHN2RRpQuTpbB5N1SU=",
      "dev": true
    },
    "assign-symbols": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/assign-symbols/download/assign-symbols-1.0.0.tgz",
      "integrity": "sha1-WWZ/QfrdTyDMvCu5a41Pf3jsA2c=",
      "dev": true
    },
    "astral-regex": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/astral-regex/download/astral-regex-1.0.0.tgz",
      "integrity": "sha1-bIw/uCfdQ+45GPJ7gngqt2WKb9k=",
      "dev": true
    },
    "async": {
      "version": "1.5.2",
      "resolved": "http://registry.npm.taobao.org/async/download/async-1.5.2.tgz",
      "integrity": "sha1-7GphrlZIDAw8skHJVhjiCJL5Zyo=",
      "dev": true
    },
    "async-each": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/async-each/download/async-each-1.0.1.tgz",
      "integrity": "sha1-GdOGodntxufByF04iu28xW0zYC0=",
      "dev": true
    },
    "async-limiter": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/async-limiter/download/async-limiter-1.0.0.tgz",
      "integrity": "sha1-ePrtjD0HSrgfIrTphdeehzj3IPg=",
      "dev": true
    },
    "asynckit": {
      "version": "0.4.0",
      "resolved": "http://registry.npm.taobao.org/asynckit/download/asynckit-0.4.0.tgz",
      "integrity": "sha1-x57Zf380y48robyXkLzDZkdLS3k=",
      "dev": true
    },
    "atob": {
      "version": "2.1.2",
      "resolved": "http://registry.npm.taobao.org/atob/download/atob-2.1.2.tgz",
      "integrity": "sha1-bZUX654DDSQ2ZmZR6GvZ9vE1M8k=",
      "dev": true
    },
    "autoprefixer": {
      "version": "8.6.5",
      "resolved": "http://registry.npm.taobao.org/autoprefixer/download/autoprefixer-8.6.5.tgz",
      "integrity": "sha1-ND89GT7VaLMgjgARehuW62kdTuk=",
      "dev": true,
      "requires": {
        "browserslist": "^3.2.8",
        "caniuse-lite": "^1.0.30000864",
        "normalize-range": "^0.1.2",
        "num2fraction": "^1.2.2",
        "postcss": "^6.0.23",
        "postcss-value-parser": "^3.2.3"
      },
      "dependencies": {
        "browserslist": {
          "version": "3.2.8",
          "resolved": "http://registry.npm.taobao.org/browserslist/download/browserslist-3.2.8.tgz",
          "integrity": "sha1-sABTYdZHHw9ZUnl6dvyYXx+Xj8Y=",
          "dev": true,
          "requires": {
            "caniuse-lite": "^1.0.30000844",
            "electron-to-chromium": "^1.3.47"
          }
        },
        "postcss": {
          "version": "6.0.23",
          "resolved": "http://registry.npm.taobao.org/postcss/download/postcss-6.0.23.tgz",
          "integrity": "sha1-YcgswyisYOZ3ZF+XkFTrmLwOMyQ=",
          "dev": true,
          "requires": {
            "chalk": "^2.4.1",
            "source-map": "^0.6.1",
            "supports-color": "^5.4.0"
          }
        },
        "source-map": {
          "version": "0.6.1",
          "resolved": "http://registry.npm.taobao.org/source-map/download/source-map-0.6.1.tgz",
          "integrity": "sha1-dHIq8y6WFOnCh6jQu95IteLxomM=",
          "dev": true
        }
      }
    },
    "aws-sign2": {
      "version": "0.7.0",
      "resolved": "http://registry.npm.taobao.org/aws-sign2/download/aws-sign2-0.7.0.tgz",
      "integrity": "sha1-tG6JCTSpWR8tL2+G1+ap8bP+dqg=",
      "dev": true
    },
    "aws4": {
      "version": "1.8.0",
      "resolved": "http://registry.npm.taobao.org/aws4/download/aws4-1.8.0.tgz",
      "integrity": "sha1-8OAD2cqef1nHpQiUXXsu+aBKVC8=",
      "dev": true
    },
    "babel-code-frame": {
      "version": "6.26.0",
      "resolved": "http://registry.npm.taobao.org/babel-code-frame/download/babel-code-frame-6.26.0.tgz",
      "integrity": "sha1-Y/1D99weO7fONZR9uP42mj9Yx0s=",
      "dev": true,
      "requires": {
        "chalk": "^1.1.3",
        "esutils": "^2.0.2",
        "js-tokens": "^3.0.2"
      },
      "dependencies": {
        "ansi-regex": {
          "version": "2.1.1",
          "resolved": "http://registry.npm.taobao.org/ansi-regex/download/ansi-regex-2.1.1.tgz",
          "integrity": "sha1-w7M6te42DYbg5ijwRorn7yfWVN8=",
          "dev": true
        },
        "ansi-styles": {
          "version": "2.2.1",
          "resolved": "http://registry.npm.taobao.org/ansi-styles/download/ansi-styles-2.2.1.tgz",
          "integrity": "sha1-tDLdM1i2NM914eRmQ2gkBTPB3b4=",
          "dev": true
        },
        "chalk": {
          "version": "1.1.3",
          "resolved": "http://registry.npm.taobao.org/chalk/download/chalk-1.1.3.tgz",
          "integrity": "sha1-qBFcVeSnAv5NFQq9OHKCKn4J/Jg=",
          "dev": true,
          "requires": {
            "ansi-styles": "^2.2.1",
            "escape-string-regexp": "^1.0.2",
            "has-ansi": "^2.0.0",
            "strip-ansi": "^3.0.0",
            "supports-color": "^2.0.0"
          }
        },
        "js-tokens": {
          "version": "3.0.2",
          "resolved": "http://registry.npm.taobao.org/js-tokens/download/js-tokens-3.0.2.tgz",
          "integrity": "sha1-mGbfOVECEw449/mWvOtlRDIJwls=",
          "dev": true
        },
        "strip-ansi": {
          "version": "3.0.1",
          "resolved": "http://registry.npm.taobao.org/strip-ansi/download/strip-ansi-3.0.1.tgz",
          "integrity": "sha1-ajhfuIU9lS1f8F0Oiq+UJ43GPc8=",
          "dev": true,
          "requires": {
            "ansi-regex": "^2.0.0"
          }
        },
        "supports-color": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/supports-color/download/supports-color-2.0.0.tgz",
          "integrity": "sha1-U10EXOa2Nj+kARcIRimZXp3zJMc=",
          "dev": true
        }
      }
    },
    "babel-eslint": {
      "version": "10.0.1",
      "resolved": "http://registry.npm.taobao.org/babel-eslint/download/babel-eslint-10.0.1.tgz",
      "integrity": "sha1-kZaB3AmWFM19MdRciQhpUJKh+u0=",
      "dev": true,
      "requires": {
        "@babel/code-frame": "^7.0.0",
        "@babel/parser": "^7.0.0",
        "@babel/traverse": "^7.0.0",
        "@babel/types": "^7.0.0",
        "eslint-scope": "3.7.1",
        "eslint-visitor-keys": "^1.0.0"
      }
    },
    "babel-helper-vue-jsx-merge-props": {
      "version": "2.0.3",
      "resolved": "http://registry.npm.taobao.org/babel-helper-vue-jsx-merge-props/download/babel-helper-vue-jsx-merge-props-2.0.3.tgz",
      "integrity": "sha1-Iq69OzOQIyjlEyk6jkmSs4T58bY=",
      "dev": true
    },
    "babel-loader": {
      "version": "8.0.4",
      "resolved": "http://registry.npm.taobao.org/babel-loader/download/babel-loader-8.0.4.tgz",
      "integrity": "sha1-e78gy+RWBini5BU0FHaS0/7L3OY=",
      "dev": true,
      "requires": {
        "find-cache-dir": "^1.0.0",
        "loader-utils": "^1.0.2",
        "mkdirp": "^0.5.1",
        "util.promisify": "^1.0.0"
      }
    },
    "babel-plugin-dynamic-import-node": {
      "version": "2.2.0",
      "resolved": "http://registry.npm.taobao.org/babel-plugin-dynamic-import-node/download/babel-plugin-dynamic-import-node-2.2.0.tgz",
      "integrity": "sha1-wK37B9lfSkSV6aqsbsOGxNfCUk4=",
      "dev": true,
      "requires": {
        "object.assign": "^4.1.0"
      }
    },
    "babel-plugin-transform-vue-jsx": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/babel-plugin-transform-vue-jsx/download/babel-plugin-transform-vue-jsx-4.0.1.tgz",
      "integrity": "sha1-LIvdzoem7wnqpZhp/xv77q/F+I0=",
      "dev": true,
      "requires": {
        "esutils": "^2.0.2"
      }
    },
    "balanced-match": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/balanced-match/download/balanced-match-1.0.0.tgz",
      "integrity": "sha1-ibTRmasr7kneFk6gK4nORi1xt2c=",
      "dev": true
    },
    "base": {
      "version": "0.11.2",
      "resolved": "http://registry.npm.taobao.org/base/download/base-0.11.2.tgz",
      "integrity": "sha1-e95c7RRbbVUakNuH+DxVi060io8=",
      "dev": true,
      "requires": {
        "cache-base": "^1.0.1",
        "class-utils": "^0.3.5",
        "component-emitter": "^1.2.1",
        "define-property": "^1.0.0",
        "isobject": "^3.0.1",
        "mixin-deep": "^1.2.0",
        "pascalcase": "^0.1.1"
      },
      "dependencies": {
        "define-property": {
          "version": "1.0.0",
          "resolved": "http://registry.npm.taobao.org/define-property/download/define-property-1.0.0.tgz",
          "integrity": "sha1-dp66rz9KY6rTr56NMEybvnm/sOY=",
          "dev": true,
          "requires": {
            "is-descriptor": "^1.0.0"
          }
        },
        "is-accessor-descriptor": {
          "version": "1.0.0",
          "resolved": "http://registry.npm.taobao.org/is-accessor-descriptor/download/is-accessor-descriptor-1.0.0.tgz",
          "integrity": "sha1-FpwvbT3x+ZJhgHI2XJsOofaHhlY=",
          "dev": true,
          "requires": {
            "kind-of": "^6.0.0"
          }
        },
        "is-data-descriptor": {
          "version": "1.0.0",
          "resolved": "http://registry.npm.taobao.org/is-data-descriptor/download/is-data-descriptor-1.0.0.tgz",
          "integrity": "sha1-2Eh2Mh0Oet0DmQQGq7u9NrqSaMc=",
          "dev": true,
          "requires": {
            "kind-of": "^6.0.0"
          }
        },
        "is-descriptor": {
          "version": "1.0.2",
          "resolved": "http://registry.npm.taobao.org/is-descriptor/download/is-descriptor-1.0.2.tgz",
          "integrity": "sha1-OxWXRqZmBLBPjIFSS6NlxfFNhuw=",
          "dev": true,
          "requires": {
            "is-accessor-descriptor": "^1.0.0",
            "is-data-descriptor": "^1.0.0",
            "kind-of": "^6.0.2"
          }
        }
      }
    },
    "base64-js": {
      "version": "1.3.0",
      "resolved": "http://registry.npm.taobao.org/base64-js/download/base64-js-1.3.0.tgz",
      "integrity": "sha1-yrHmEY8FEJXli1KBrqjBzSK/wOM=",
      "dev": true
    },
    "batch": {
      "version": "0.6.1",
      "resolved": "http://registry.npm.taobao.org/batch/download/batch-0.6.1.tgz",
      "integrity": "sha1-3DQxT05nkxgJP8dgJyUl+UvyXBY=",
      "dev": true
    },
    "bcrypt-pbkdf": {
      "version": "1.0.2",
      "resolved": "http://registry.npm.taobao.org/bcrypt-pbkdf/download/bcrypt-pbkdf-1.0.2.tgz",
      "integrity": "sha1-pDAdOJtqQ/m2f/PKEaP2Y342Dp4=",
      "dev": true,
      "requires": {
        "tweetnacl": "^0.14.3"
      }
    },
    "bfj": {
      "version": "6.1.1",
      "resolved": "http://registry.npm.taobao.org/bfj/download/bfj-6.1.1.tgz",
      "integrity": "sha1-BaO3eE+9cs+jwi5WAC75kzZRbEg=",
      "dev": true,
      "requires": {
        "bluebird": "^3.5.1",
        "check-types": "^7.3.0",
        "hoopy": "^0.1.2",
        "tryer": "^1.0.0"
      }
    },
    "big.js": {
      "version": "3.2.0",
      "resolved": "http://registry.npm.taobao.org/big.js/download/big.js-3.2.0.tgz",
      "integrity": "sha1-pfwpi4G54Nyi5FiCR4S2XFK6WI4=",
      "dev": true
    },
    "binary-extensions": {
      "version": "1.12.0",
      "resolved": "http://registry.npm.taobao.org/binary-extensions/download/binary-extensions-1.12.0.tgz",
      "integrity": "sha1-wteA9T1Fu6gxeokC1M7q86Y4WxQ=",
      "dev": true
    },
    "bluebird": {
      "version": "3.5.3",
      "resolved": "http://registry.npm.taobao.org/bluebird/download/bluebird-3.5.3.tgz",
      "integrity": "sha1-fQHG+WFsmlGrD4xUmnnf5uwz76c=",
      "dev": true
    },
    "bn.js": {
      "version": "4.11.8",
      "resolved": "http://registry.npm.taobao.org/bn.js/download/bn.js-4.11.8.tgz",
      "integrity": "sha1-LN4J617jQfSEdGuwMJsyU7GxRC8=",
      "dev": true
    },
    "body-parser": {
      "version": "1.18.3",
      "resolved": "http://registry.npm.taobao.org/body-parser/download/body-parser-1.18.3.tgz",
      "integrity": "sha1-WykhmP/dVTs6DyDe0FkrlWlVyLQ=",
      "dev": true,
      "requires": {
        "bytes": "3.0.0",
        "content-type": "~1.0.4",
        "debug": "2.6.9",
        "depd": "~1.1.2",
        "http-errors": "~1.6.3",
        "iconv-lite": "0.4.23",
        "on-finished": "~2.3.0",
        "qs": "6.5.2",
        "raw-body": "2.3.3",
        "type-is": "~1.6.16"
      },
      "dependencies": {
        "debug": {
          "version": "2.6.9",
          "resolved": "http://registry.npm.taobao.org/debug/download/debug-2.6.9.tgz",
          "integrity": "sha1-XRKFFd8TT/Mn6QpMk/Tgd6U2NB8=",
          "dev": true,
          "requires": {
            "ms": "2.0.0"
          }
        },
        "iconv-lite": {
          "version": "0.4.23",
          "resolved": "http://registry.npm.taobao.org/iconv-lite/download/iconv-lite-0.4.23.tgz",
          "integrity": "sha1-KXhx9jvlB63Pv8pxXQzQ7thOmmM=",
          "dev": true,
          "requires": {
            "safer-buffer": ">= 2.1.2 < 3"
          }
        },
        "ms": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/ms/download/ms-2.0.0.tgz",
          "integrity": "sha1-VgiurfwAvmwpAd9fmGF4jeDVl8g=",
          "dev": true
        }
      }
    },
    "bonjour": {
      "version": "3.5.0",
      "resolved": "http://registry.npm.taobao.org/bonjour/download/bonjour-3.5.0.tgz",
      "integrity": "sha1-jokKGD2O6aI5OzhExpGkK897yfU=",
      "dev": true,
      "requires": {
        "array-flatten": "^2.1.0",
        "deep-equal": "^1.0.1",
        "dns-equal": "^1.0.0",
        "dns-txt": "^2.0.2",
        "multicast-dns": "^6.0.1",
        "multicast-dns-service-types": "^1.1.0"
      },
      "dependencies": {
        "array-flatten": {
          "version": "2.1.2",
          "resolved": "http://registry.npm.taobao.org/array-flatten/download/array-flatten-2.1.2.tgz",
          "integrity": "sha1-JO+AoowaiTYX4hSbDG0NeIKTsJk=",
          "dev": true
        }
      }
    },
    "boolbase": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/boolbase/download/boolbase-1.0.0.tgz",
      "integrity": "sha1-aN/1++YMUes3cl6p4+0xDcwed24=",
      "dev": true
    },
    "brace-expansion": {
      "version": "1.1.11",
      "resolved": "http://registry.npm.taobao.org/brace-expansion/download/brace-expansion-1.1.11.tgz",
      "integrity": "sha1-PH/L9SnYcibz0vUrlm/1Jx60Qd0=",
      "dev": true,
      "requires": {
        "balanced-match": "^1.0.0",
        "concat-map": "0.0.1"
      }
    },
    "braces": {
      "version": "2.3.2",
      "resolved": "http://registry.npm.taobao.org/braces/download/braces-2.3.2.tgz",
      "integrity": "sha1-WXn9PxTNUxVl5fot8av/8d+u5yk=",
      "dev": true,
      "requires": {
        "arr-flatten": "^1.1.0",
        "array-unique": "^0.3.2",
        "extend-shallow": "^2.0.1",
        "fill-range": "^4.0.0",
        "isobject": "^3.0.1",
        "repeat-element": "^1.1.2",
        "snapdragon": "^0.8.1",
        "snapdragon-node": "^2.0.1",
        "split-string": "^3.0.2",
        "to-regex": "^3.0.1"
      },
      "dependencies": {
        "extend-shallow": {
          "version": "2.0.1",
          "resolved": "http://registry.npm.taobao.org/extend-shallow/download/extend-shallow-2.0.1.tgz",
          "integrity": "sha1-Ua99YUrZqfYQ6huvu5idaxxWiQ8=",
          "dev": true,
          "requires": {
            "is-extendable": "^0.1.0"
          }
        }
      }
    },
    "brorand": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/brorand/download/brorand-1.1.0.tgz",
      "integrity": "sha1-EsJe/kCkXjwyPrhnWgoM5XsiNx8=",
      "dev": true
    },
    "browserify-aes": {
      "version": "1.2.0",
      "resolved": "http://registry.npm.taobao.org/browserify-aes/download/browserify-aes-1.2.0.tgz",
      "integrity": "sha1-Mmc0ZC9APavDADIJhTu3CtQo70g=",
      "dev": true,
      "requires": {
        "buffer-xor": "^1.0.3",
        "cipher-base": "^1.0.0",
        "create-hash": "^1.1.0",
        "evp_bytestokey": "^1.0.3",
        "inherits": "^2.0.1",
        "safe-buffer": "^5.0.1"
      }
    },
    "browserify-cipher": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/browserify-cipher/download/browserify-cipher-1.0.1.tgz",
      "integrity": "sha1-jWR0wbhwv9q807z8wZNKEOlPFfA=",
      "dev": true,
      "requires": {
        "browserify-aes": "^1.0.4",
        "browserify-des": "^1.0.0",
        "evp_bytestokey": "^1.0.0"
      }
    },
    "browserify-des": {
      "version": "1.0.2",
      "resolved": "http://registry.npm.taobao.org/browserify-des/download/browserify-des-1.0.2.tgz",
      "integrity": "sha1-OvTx9Zg5QDVy8cZiBDdfen9wPpw=",
      "dev": true,
      "requires": {
        "cipher-base": "^1.0.1",
        "des.js": "^1.0.0",
        "inherits": "^2.0.1",
        "safe-buffer": "^5.1.2"
      }
    },
    "browserify-rsa": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/browserify-rsa/download/browserify-rsa-4.0.1.tgz",
      "integrity": "sha1-IeCr+vbyApzy+vsTNWenAdQTVSQ=",
      "dev": true,
      "requires": {
        "bn.js": "^4.1.0",
        "randombytes": "^2.0.1"
      }
    },
    "browserify-sign": {
      "version": "4.0.4",
      "resolved": "http://registry.npm.taobao.org/browserify-sign/download/browserify-sign-4.0.4.tgz",
      "integrity": "sha1-qk62jl17ZYuqa/alfmMMvXqT0pg=",
      "dev": true,
      "requires": {
        "bn.js": "^4.1.1",
        "browserify-rsa": "^4.0.0",
        "create-hash": "^1.1.0",
        "create-hmac": "^1.1.2",
        "elliptic": "^6.0.0",
        "inherits": "^2.0.1",
        "parse-asn1": "^5.0.0"
      }
    },
    "browserify-zlib": {
      "version": "0.2.0",
      "resolved": "http://registry.npm.taobao.org/browserify-zlib/download/browserify-zlib-0.2.0.tgz",
      "integrity": "sha1-KGlFnZqjviRf6P4sofRuLn9U1z8=",
      "dev": true,
      "requires": {
        "pako": "~1.0.5"
      }
    },
    "browserslist": {
      "version": "4.3.5",
      "resolved": "http://registry.npm.taobao.org/browserslist/download/browserslist-4.3.5.tgz",
      "integrity": "sha1-GpF2eKzAe1VgZ0jqGt+YRuqJIPc=",
      "dev": true,
      "requires": {
        "caniuse-lite": "^1.0.30000912",
        "electron-to-chromium": "^1.3.86",
        "node-releases": "^1.0.5"
      }
    },
    "buffer": {
      "version": "4.9.1",
      "resolved": "http://registry.npm.taobao.org/buffer/download/buffer-4.9.1.tgz",
      "integrity": "sha1-bRu2AbB6TvztlwlBMgkwJ8lbwpg=",
      "dev": true,
      "requires": {
        "base64-js": "^1.0.2",
        "ieee754": "^1.1.4",
        "isarray": "^1.0.0"
      }
    },
    "buffer-from": {
      "version": "1.1.1",
      "resolved": "http://registry.npm.taobao.org/buffer-from/download/buffer-from-1.1.1.tgz",
      "integrity": "sha1-MnE7wCj3XAL9txDXx7zsHyxgcO8=",
      "dev": true
    },
    "buffer-indexof": {
      "version": "1.1.1",
      "resolved": "http://registry.npm.taobao.org/buffer-indexof/download/buffer-indexof-1.1.1.tgz",
      "integrity": "sha1-Uvq8xqYG0aADAoAmSO9o9jnaJow=",
      "dev": true
    },
    "buffer-xor": {
      "version": "1.0.3",
      "resolved": "http://registry.npm.taobao.org/buffer-xor/download/buffer-xor-1.0.3.tgz",
      "integrity": "sha1-JuYe0UIvtw3ULm42cp7VHYVf6Nk=",
      "dev": true
    },
    "builtin-modules": {
      "version": "1.1.1",
      "resolved": "http://registry.npm.taobao.org/builtin-modules/download/builtin-modules-1.1.1.tgz",
      "integrity": "sha1-Jw8HbFpywC9bZaR9+Uxf46J4iS8=",
      "dev": true
    },
    "builtin-status-codes": {
      "version": "3.0.0",
      "resolved": "http://registry.npm.taobao.org/builtin-status-codes/download/builtin-status-codes-3.0.0.tgz",
      "integrity": "sha1-hZgoeOIbmOHGZCXgPQF0eI9Wnug=",
      "dev": true
    },
    "bytes": {
      "version": "3.0.0",
      "resolved": "http://registry.npm.taobao.org/bytes/download/bytes-3.0.0.tgz",
      "integrity": "sha1-0ygVQE1olpn4Wk6k+odV3ROpYEg=",
      "dev": true
    },
    "cacache": {
      "version": "10.0.4",
      "resolved": "http://registry.npm.taobao.org/cacache/download/cacache-10.0.4.tgz",
      "integrity": "sha1-ZFI2eZnv+dQYiu/ZoU6dfGomNGA=",
      "dev": true,
      "requires": {
        "bluebird": "^3.5.1",
        "chownr": "^1.0.1",
        "glob": "^7.1.2",
        "graceful-fs": "^4.1.11",
        "lru-cache": "^4.1.1",
        "mississippi": "^2.0.0",
        "mkdirp": "^0.5.1",
        "move-concurrently": "^1.0.1",
        "promise-inflight": "^1.0.1",
        "rimraf": "^2.6.2",
        "ssri": "^5.2.4",
        "unique-filename": "^1.1.0",
        "y18n": "^4.0.0"
      },
      "dependencies": {
        "ssri": {
          "version": "5.3.0",
          "resolved": "http://registry.npm.taobao.org/ssri/download/ssri-5.3.0.tgz",
          "integrity": "sha1-ujhyycbTOgcEp9cf8EXl7EiZnQY=",
          "dev": true,
          "requires": {
            "safe-buffer": "^5.1.1"
          }
        }
      }
    },
    "cache-base": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/cache-base/download/cache-base-1.0.1.tgz",
      "integrity": "sha1-Cn9GQWgxyLZi7jb+TnxZ129marI=",
      "dev": true,
      "requires": {
        "collection-visit": "^1.0.0",
        "component-emitter": "^1.2.1",
        "get-value": "^2.0.6",
        "has-value": "^1.0.0",
        "isobject": "^3.0.1",
        "set-value": "^2.0.0",
        "to-object-path": "^0.3.0",
        "union-value": "^1.0.0",
        "unset-value": "^1.0.0"
      }
    },
    "cache-loader": {
      "version": "1.2.5",
      "resolved": "http://registry.npm.taobao.org/cache-loader/download/cache-loader-1.2.5.tgz",
      "integrity": "sha1-mrFbCuX1RvN2CDppX8GnX1RssmY=",
      "dev": true,
      "requires": {
        "loader-utils": "^1.1.0",
        "mkdirp": "^0.5.1",
        "neo-async": "^2.5.0",
        "schema-utils": "^0.4.2"
      }
    },
    "call-me-maybe": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/call-me-maybe/download/call-me-maybe-1.0.1.tgz",
      "integrity": "sha1-JtII6onje1y95gJQoV8DHBak1ms=",
      "dev": true
    },
    "caller-callsite": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/caller-callsite/download/caller-callsite-2.0.0.tgz",
      "integrity": "sha1-hH4PzgoiN1CpoCfFSzNzGtMVQTQ=",
      "dev": true,
      "requires": {
        "callsites": "^2.0.0"
      },
      "dependencies": {
        "callsites": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/callsites/download/callsites-2.0.0.tgz",
          "integrity": "sha1-BuuE8A7qQT2oav/vrL/7Ngk7PFA=",
          "dev": true
        }
      }
    },
    "caller-path": {
      "version": "0.1.0",
      "resolved": "http://registry.npm.taobao.org/caller-path/download/caller-path-0.1.0.tgz",
      "integrity": "sha1-lAhe9jWB7NPaqSREqP6U6CV3dR8=",
      "dev": true,
      "requires": {
        "callsites": "^0.2.0"
      }
    },
    "callsites": {
      "version": "0.2.0",
      "resolved": "http://registry.npm.taobao.org/callsites/download/callsites-0.2.0.tgz",
      "integrity": "sha1-r6uWJikQp/M8GaV3WCXGnzTjUMo=",
      "dev": true
    },
    "camel-case": {
      "version": "3.0.0",
      "resolved": "http://registry.npm.taobao.org/camel-case/download/camel-case-3.0.0.tgz",
      "integrity": "sha1-yjw2iKTpzzpM2nd9xNy8cTJJz3M=",
      "dev": true,
      "requires": {
        "no-case": "^2.2.0",
        "upper-case": "^1.1.1"
      }
    },
    "camelcase": {
      "version": "4.1.0",
      "resolved": "http://registry.npm.taobao.org/camelcase/download/camelcase-4.1.0.tgz",
      "integrity": "sha1-1UVjW+HjPFQmScaRc+Xeas+uNN0=",
      "dev": true
    },
    "caniuse-api": {
      "version": "3.0.0",
      "resolved": "http://registry.npm.taobao.org/caniuse-api/download/caniuse-api-3.0.0.tgz",
      "integrity": "sha1-Xk2Q4idJYdRikZl99Znj7QCO5MA=",
      "dev": true,
      "requires": {
        "browserslist": "^4.0.0",
        "caniuse-lite": "^1.0.0",
        "lodash.memoize": "^4.1.2",
        "lodash.uniq": "^4.5.0"
      }
    },
    "caniuse-lite": {
      "version": "1.0.30000918",
      "resolved": "http://registry.npm.taobao.org/caniuse-lite/download/caniuse-lite-1.0.30000918.tgz",
      "integrity": "sha1-Yoj3naPFyLReUC9HrY8+uR8Teak=",
      "dev": true
    },
    "case-sensitive-paths-webpack-plugin": {
      "version": "2.1.2",
      "resolved": "http://registry.npm.taobao.org/case-sensitive-paths-webpack-plugin/download/case-sensitive-paths-webpack-plugin-2.1.2.tgz",
      "integrity": "sha1-yJm1IXV2NokiRXHa13h0LhM/AZI=",
      "dev": true
    },
    "caseless": {
      "version": "0.12.0",
      "resolved": "http://registry.npm.taobao.org/caseless/download/caseless-0.12.0.tgz",
      "integrity": "sha1-G2gcIf+EAzyCZUMJBolCDRhxUdw=",
      "dev": true
    },
    "chalk": {
      "version": "2.4.1",
      "resolved": "http://registry.npm.taobao.org/chalk/download/chalk-2.4.1.tgz",
      "integrity": "sha1-GMSasWoDe26wFSzIPjRxM4IVtm4=",
      "dev": true,
      "requires": {
        "ansi-styles": "^3.2.1",
        "escape-string-regexp": "^1.0.5",
        "supports-color": "^5.3.0"
      }
    },
    "chardet": {
      "version": "0.4.2",
      "resolved": "http://registry.npm.taobao.org/chardet/download/chardet-0.4.2.tgz",
      "integrity": "sha1-tUc7M9yXxCTl2Y3IfVXU2KKci/I=",
      "dev": true
    },
    "check-types": {
      "version": "7.4.0",
      "resolved": "http://registry.npm.taobao.org/check-types/download/check-types-7.4.0.tgz",
      "integrity": "sha1-A3jsG5YW7HH3dJMaPGUW+tjBUvQ=",
      "dev": true
    },
    "chokidar": {
      "version": "2.0.4",
      "resolved": "http://registry.npm.taobao.org/chokidar/download/chokidar-2.0.4.tgz",
      "integrity": "sha1-NW/04rDo5D4yLRijckYLvPOszSY=",
      "dev": true,
      "requires": {
        "anymatch": "^2.0.0",
        "async-each": "^1.0.0",
        "braces": "^2.3.0",
        "fsevents": "^1.2.2",
        "glob-parent": "^3.1.0",
        "inherits": "^2.0.1",
        "is-binary-path": "^1.0.0",
        "is-glob": "^4.0.0",
        "lodash.debounce": "^4.0.8",
        "normalize-path": "^2.1.1",
        "path-is-absolute": "^1.0.0",
        "readdirp": "^2.0.0",
        "upath": "^1.0.5"
      }
    },
    "chownr": {
      "version": "1.1.1",
      "resolved": "http://registry.npm.taobao.org/chownr/download/chownr-1.1.1.tgz",
      "integrity": "sha1-VHJri4//TfBTxCGH6AH7RBLfFJQ=",
      "dev": true
    },
    "chrome-trace-event": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/chrome-trace-event/download/chrome-trace-event-1.0.0.tgz",
      "integrity": "sha1-Rakb0sIMlBHwljtarrmhuV4JzEg=",
      "dev": true,
      "requires": {
        "tslib": "^1.9.0"
      }
    },
    "ci-info": {
      "version": "1.6.0",
      "resolved": "http://registry.npm.taobao.org/ci-info/download/ci-info-1.6.0.tgz",
      "integrity": "sha1-LKINu5zrMtRSSmgzAzE/AwSx5Jc=",
      "dev": true
    },
    "cipher-base": {
      "version": "1.0.4",
      "resolved": "http://registry.npm.taobao.org/cipher-base/download/cipher-base-1.0.4.tgz",
      "integrity": "sha1-h2Dk7MJy9MNjUy+SbYdKriwTl94=",
      "dev": true,
      "requires": {
        "inherits": "^2.0.1",
        "safe-buffer": "^5.0.1"
      }
    },
    "circular-json": {
      "version": "0.3.3",
      "resolved": "http://registry.npm.taobao.org/circular-json/download/circular-json-0.3.3.tgz",
      "integrity": "sha1-gVyZ6oT2gJUp0vRXkb34JxE1LWY=",
      "dev": true
    },
    "class-utils": {
      "version": "0.3.6",
      "resolved": "http://registry.npm.taobao.org/class-utils/download/class-utils-0.3.6.tgz",
      "integrity": "sha1-+TNprouafOAv1B+q0MqDAzGQxGM=",
      "dev": true,
      "requires": {
        "arr-union": "^3.1.0",
        "define-property": "^0.2.5",
        "isobject": "^3.0.0",
        "static-extend": "^0.1.1"
      },
      "dependencies": {
        "define-property": {
          "version": "0.2.5",
          "resolved": "http://registry.npm.taobao.org/define-property/download/define-property-0.2.5.tgz",
          "integrity": "sha1-w1se+RjsPJkPmlvFe+BKrOxcgRY=",
          "dev": true,
          "requires": {
            "is-descriptor": "^0.1.0"
          }
        }
      }
    },
    "clean-css": {
      "version": "4.2.1",
      "resolved": "http://registry.npm.taobao.org/clean-css/download/clean-css-4.2.1.tgz",
      "integrity": "sha1-LUEe92uFabbQyEBo2r6FsKpeXBc=",
      "dev": true,
      "requires": {
        "source-map": "~0.6.0"
      },
      "dependencies": {
        "source-map": {
          "version": "0.6.1",
          "resolved": "http://registry.npm.taobao.org/source-map/download/source-map-0.6.1.tgz",
          "integrity": "sha1-dHIq8y6WFOnCh6jQu95IteLxomM=",
          "dev": true
        }
      }
    },
    "cli-cursor": {
      "version": "2.1.0",
      "resolved": "http://registry.npm.taobao.org/cli-cursor/download/cli-cursor-2.1.0.tgz",
      "integrity": "sha1-s12sN2R5+sw+lHR9QdDQ9SOP/LU=",
      "dev": true,
      "requires": {
        "restore-cursor": "^2.0.0"
      }
    },
    "cli-spinners": {
      "version": "1.3.1",
      "resolved": "http://registry.npm.taobao.org/cli-spinners/download/cli-spinners-1.3.1.tgz",
      "integrity": "sha1-ACwZkJEtDVlYDJO9NsBW3pnkJZo=",
      "dev": true
    },
    "cli-width": {
      "version": "2.2.0",
      "resolved": "http://registry.npm.taobao.org/cli-width/download/cli-width-2.2.0.tgz",
      "integrity": "sha1-/xnt6Kml5XkyQUewwR8PvLq+1jk=",
      "dev": true
    },
    "clipboardy": {
      "version": "1.2.3",
      "resolved": "http://registry.npm.taobao.org/clipboardy/download/clipboardy-1.2.3.tgz",
      "integrity": "sha1-BSY2G/eHJMHyC+JI1CjjZUM8B+8=",
      "dev": true,
      "requires": {
        "arch": "^2.1.0",
        "execa": "^0.8.0"
      },
      "dependencies": {
        "cross-spawn": {
          "version": "5.1.0",
          "resolved": "http://registry.npm.taobao.org/cross-spawn/download/cross-spawn-5.1.0.tgz",
          "integrity": "sha1-6L0O/uWPz/b4+UUQoKVUu/ojVEk=",
          "dev": true,
          "requires": {
            "lru-cache": "^4.0.1",
            "shebang-command": "^1.2.0",
            "which": "^1.2.9"
          }
        },
        "execa": {
          "version": "0.8.0",
          "resolved": "http://registry.npm.taobao.org/execa/download/execa-0.8.0.tgz",
          "integrity": "sha1-2NdrvBtVIX7RkP1t1J08d07PyNo=",
          "dev": true,
          "requires": {
            "cross-spawn": "^5.0.1",
            "get-stream": "^3.0.0",
            "is-stream": "^1.1.0",
            "npm-run-path": "^2.0.0",
            "p-finally": "^1.0.0",
            "signal-exit": "^3.0.0",
            "strip-eof": "^1.0.0"
          }
        },
        "get-stream": {
          "version": "3.0.0",
          "resolved": "http://registry.npm.taobao.org/get-stream/download/get-stream-3.0.0.tgz",
          "integrity": "sha1-jpQ9E1jcN1VQVOy+LtsFqhdO3hQ=",
          "dev": true
        }
      }
    },
    "cliui": {
      "version": "4.1.0",
      "resolved": "http://registry.npm.taobao.org/cliui/download/cliui-4.1.0.tgz",
      "integrity": "sha1-NIQi2+gtgAswIu709qwQvy5NG0k=",
      "dev": true,
      "requires": {
        "string-width": "^2.1.1",
        "strip-ansi": "^4.0.0",
        "wrap-ansi": "^2.0.0"
      }
    },
    "clone": {
      "version": "1.0.4",
      "resolved": "http://registry.npm.taobao.org/clone/download/clone-1.0.4.tgz",
      "integrity": "sha1-2jCcwmPfFZlMaIypAheco8fNfH4=",
      "dev": true
    },
    "co": {
      "version": "4.6.0",
      "resolved": "http://registry.npm.taobao.org/co/download/co-4.6.0.tgz",
      "integrity": "sha1-bqa989hTrlTMuOR7+gvz+QMfsYQ=",
      "dev": true
    },
    "coa": {
      "version": "2.0.2",
      "resolved": "http://registry.npm.taobao.org/coa/download/coa-2.0.2.tgz",
      "integrity": "sha1-Q/bCEVG07yv1cYfbDXPeIp4+fsM=",
      "dev": true,
      "requires": {
        "@types/q": "^1.5.1",
        "chalk": "^2.4.1",
        "q": "^1.1.2"
      }
    },
    "code-point-at": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/code-point-at/download/code-point-at-1.1.0.tgz",
      "integrity": "sha1-DQcLTQQ6W+ozovGkDi7bPZpMz3c=",
      "dev": true
    },
    "collection-visit": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/collection-visit/download/collection-visit-1.0.0.tgz",
      "integrity": "sha1-S8A3PBZLwykbTTaMgpzxqApZ3KA=",
      "dev": true,
      "requires": {
        "map-visit": "^1.0.0",
        "object-visit": "^1.0.0"
      }
    },
    "color": {
      "version": "3.1.0",
      "resolved": "http://registry.npm.taobao.org/color/download/color-3.1.0.tgz",
      "integrity": "sha1-2On7CWcyh1d0yEv5IoFd8DCND/w=",
      "dev": true,
      "requires": {
        "color-convert": "^1.9.1",
        "color-string": "^1.5.2"
      }
    },
    "color-convert": {
      "version": "1.9.3",
      "resolved": "http://registry.npm.taobao.org/color-convert/download/color-convert-1.9.3.tgz",
      "integrity": "sha1-u3GFBpDh8TZWfeYp0tVHHe2kweg=",
      "dev": true,
      "requires": {
        "color-name": "1.1.3"
      }
    },
    "color-name": {
      "version": "1.1.3",
      "resolved": "http://registry.npm.taobao.org/color-name/download/color-name-1.1.3.tgz",
      "integrity": "sha1-p9BVi9icQveV3UIyj3QIMcpTvCU=",
      "dev": true
    },
    "color-string": {
      "version": "1.5.3",
      "resolved": "http://registry.npm.taobao.org/color-string/download/color-string-1.5.3.tgz",
      "integrity": "sha1-ybvF8BtYtUkvPWhXRZy2WQziBMw=",
      "dev": true,
      "requires": {
        "color-name": "^1.0.0",
        "simple-swizzle": "^0.2.2"
      }
    },
    "colors": {
      "version": "1.1.2",
      "resolved": "http://registry.npm.taobao.org/colors/download/colors-1.1.2.tgz",
      "integrity": "sha1-FopHAXVran9RoSzgyXv6KMCE7WM=",
      "dev": true
    },
    "combined-stream": {
      "version": "1.0.7",
      "resolved": "http://registry.npm.taobao.org/combined-stream/download/combined-stream-1.0.7.tgz",
      "integrity": "sha1-LR0kMXr7ir6V1tLAsHtXgTU52Cg=",
      "dev": true,
      "requires": {
        "delayed-stream": "~1.0.0"
      }
    },
    "commander": {
      "version": "2.17.1",
      "resolved": "http://registry.npm.taobao.org/commander/download/commander-2.17.1.tgz",
      "integrity": "sha1-vXerfebelCBc6sxy8XFtKfIKd78=",
      "dev": true
    },
    "commondir": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/commondir/download/commondir-1.0.1.tgz",
      "integrity": "sha1-3dgA2gxmEnOTzKWVDqloo6rxJTs=",
      "dev": true
    },
    "component-emitter": {
      "version": "1.2.1",
      "resolved": "http://registry.npm.taobao.org/component-emitter/download/component-emitter-1.2.1.tgz",
      "integrity": "sha1-E3kY1teCg/ffemt8WmPhQOaUJeY=",
      "dev": true
    },
    "compressible": {
      "version": "2.0.15",
      "resolved": "http://registry.npm.taobao.org/compressible/download/compressible-2.0.15.tgz",
      "integrity": "sha1-hXqasKfloH2Ng37UP+Le//ZP4hI=",
      "dev": true,
      "requires": {
        "mime-db": ">= 1.36.0 < 2"
      }
    },
    "compression": {
      "version": "1.7.3",
      "resolved": "http://registry.npm.taobao.org/compression/download/compression-1.7.3.tgz",
      "integrity": "sha1-J+DhdqryYPfywoE8PkQK258Zk9s=",
      "dev": true,
      "requires": {
        "accepts": "~1.3.5",
        "bytes": "3.0.0",
        "compressible": "~2.0.14",
        "debug": "2.6.9",
        "on-headers": "~1.0.1",
        "safe-buffer": "5.1.2",
        "vary": "~1.1.2"
      },
      "dependencies": {
        "debug": {
          "version": "2.6.9",
          "resolved": "http://registry.npm.taobao.org/debug/download/debug-2.6.9.tgz",
          "integrity": "sha1-XRKFFd8TT/Mn6QpMk/Tgd6U2NB8=",
          "dev": true,
          "requires": {
            "ms": "2.0.0"
          }
        },
        "ms": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/ms/download/ms-2.0.0.tgz",
          "integrity": "sha1-VgiurfwAvmwpAd9fmGF4jeDVl8g=",
          "dev": true
        }
      }
    },
    "concat-map": {
      "version": "0.0.1",
      "resolved": "http://registry.npm.taobao.org/concat-map/download/concat-map-0.0.1.tgz",
      "integrity": "sha1-2Klr13/Wjfd5OnMDajug1UBdR3s=",
      "dev": true
    },
    "concat-stream": {
      "version": "1.6.2",
      "resolved": "http://registry.npm.taobao.org/concat-stream/download/concat-stream-1.6.2.tgz",
      "integrity": "sha1-kEvfGUzTEi/Gdcd/xKw9T/D9GjQ=",
      "dev": true,
      "requires": {
        "buffer-from": "^1.0.0",
        "inherits": "^2.0.3",
        "readable-stream": "^2.2.2",
        "typedarray": "^0.0.6"
      }
    },
    "connect-history-api-fallback": {
      "version": "1.5.0",
      "resolved": "http://registry.npm.taobao.org/connect-history-api-fallback/download/connect-history-api-fallback-1.5.0.tgz",
      "integrity": "sha1-sGhzk0vF40T+9hGhlqb6rgruAVo=",
      "dev": true
    },
    "console-browserify": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/console-browserify/download/console-browserify-1.1.0.tgz",
      "integrity": "sha1-8CQcRXMKn8YyOyBtvzjtx0HQuxA=",
      "dev": true,
      "requires": {
        "date-now": "^0.1.4"
      }
    },
    "consolidate": {
      "version": "0.15.1",
      "resolved": "http://registry.npm.taobao.org/consolidate/download/consolidate-0.15.1.tgz",
      "integrity": "sha1-IasEMjXHGgfUXZqtmFk7DbpWurc=",
      "dev": true,
      "requires": {
        "bluebird": "^3.1.1"
      }
    },
    "constants-browserify": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/constants-browserify/download/constants-browserify-1.0.0.tgz",
      "integrity": "sha1-wguW2MYXdIqvHBYCF2DNJ/y4y3U=",
      "dev": true
    },
    "content-disposition": {
      "version": "0.5.2",
      "resolved": "http://registry.npm.taobao.org/content-disposition/download/content-disposition-0.5.2.tgz",
      "integrity": "sha1-DPaLud318r55YcOoUXjLhdunjLQ=",
      "dev": true
    },
    "content-type": {
      "version": "1.0.4",
      "resolved": "http://registry.npm.taobao.org/content-type/download/content-type-1.0.4.tgz",
      "integrity": "sha1-4TjMdeBAxyexlm/l5fjJruJW/js=",
      "dev": true
    },
    "convert-source-map": {
      "version": "1.6.0",
      "resolved": "http://registry.npm.taobao.org/convert-source-map/download/convert-source-map-1.6.0.tgz",
      "integrity": "sha1-UbU3qMQ+DwTewZk7/83VBOdYrCA=",
      "dev": true,
      "requires": {
        "safe-buffer": "~5.1.1"
      }
    },
    "cookie": {
      "version": "0.3.1",
      "resolved": "http://registry.npm.taobao.org/cookie/download/cookie-0.3.1.tgz",
      "integrity": "sha1-5+Ch+e9DtMi6klxcWpboBtFoc7s=",
      "dev": true
    },
    "cookie-signature": {
      "version": "1.0.6",
      "resolved": "http://registry.npm.taobao.org/cookie-signature/download/cookie-signature-1.0.6.tgz",
      "integrity": "sha1-4wOogrNCzD7oylE6eZmXNNqzriw=",
      "dev": true
    },
    "copy-concurrently": {
      "version": "1.0.5",
      "resolved": "http://registry.npm.taobao.org/copy-concurrently/download/copy-concurrently-1.0.5.tgz",
      "integrity": "sha1-kilzmMrjSTf8r9bsgTnBgFHwteA=",
      "dev": true,
      "requires": {
        "aproba": "^1.1.1",
        "fs-write-stream-atomic": "^1.0.8",
        "iferr": "^0.1.5",
        "mkdirp": "^0.5.1",
        "rimraf": "^2.5.4",
        "run-queue": "^1.0.0"
      }
    },
    "copy-descriptor": {
      "version": "0.1.1",
      "resolved": "http://registry.npm.taobao.org/copy-descriptor/download/copy-descriptor-0.1.1.tgz",
      "integrity": "sha1-Z29us8OZl8LuGsOpJP1hJHSPV40=",
      "dev": true
    },
    "copy-webpack-plugin": {
      "version": "4.6.0",
      "resolved": "http://registry.npm.taobao.org/copy-webpack-plugin/download/copy-webpack-plugin-4.6.0.tgz",
      "integrity": "sha1-5/QN2KaEd9QF3Rt6hUquMksVi64=",
      "dev": true,
      "requires": {
        "cacache": "^10.0.4",
        "find-cache-dir": "^1.0.0",
        "globby": "^7.1.1",
        "is-glob": "^4.0.0",
        "loader-utils": "^1.1.0",
        "minimatch": "^3.0.4",
        "p-limit": "^1.0.0",
        "serialize-javascript": "^1.4.0"
      },
      "dependencies": {
        "globby": {
          "version": "7.1.1",
          "resolved": "http://registry.npm.taobao.org/globby/download/globby-7.1.1.tgz",
          "integrity": "sha1-+yzP+UAfhgCUXfral0QMypcrhoA=",
          "dev": true,
          "requires": {
            "array-union": "^1.0.1",
            "dir-glob": "^2.0.0",
            "glob": "^7.1.2",
            "ignore": "^3.3.5",
            "pify": "^3.0.0",
            "slash": "^1.0.0"
          }
        }
      }
    },
    "core-js": {
      "version": "2.6.0",
      "resolved": "http://registry.npm.taobao.org/core-js/download/core-js-2.6.0.tgz",
      "integrity": "sha1-HjB5Pp7leCswfjf/oi2g6s3dhNQ=",
      "dev": true
    },
    "core-util-is": {
      "version": "1.0.2",
      "resolved": "http://registry.npm.taobao.org/core-util-is/download/core-util-is-1.0.2.tgz",
      "integrity": "sha1-tf1UIgqivFq1eqtxQMlAdUUDwac=",
      "dev": true
    },
    "cosmiconfig": {
      "version": "5.0.7",
      "resolved": "http://registry.npm.taobao.org/cosmiconfig/download/cosmiconfig-5.0.7.tgz",
      "integrity": "sha1-OYJrKS7g147aE336MXO9HCGkOwQ=",
      "dev": true,
      "requires": {
        "import-fresh": "^2.0.0",
        "is-directory": "^0.3.1",
        "js-yaml": "^3.9.0",
        "parse-json": "^4.0.0"
      }
    },
    "create-ecdh": {
      "version": "4.0.3",
      "resolved": "http://registry.npm.taobao.org/create-ecdh/download/create-ecdh-4.0.3.tgz",
      "integrity": "sha1-yREbbzMEXEaX8UR4f5JUzcd8Rf8=",
      "dev": true,
      "requires": {
        "bn.js": "^4.1.0",
        "elliptic": "^6.0.0"
      }
    },
    "create-hash": {
      "version": "1.2.0",
      "resolved": "http://registry.npm.taobao.org/create-hash/download/create-hash-1.2.0.tgz",
      "integrity": "sha1-iJB4rxGmN1a8+1m9IhmWvjqe8ZY=",
      "dev": true,
      "requires": {
        "cipher-base": "^1.0.1",
        "inherits": "^2.0.1",
        "md5.js": "^1.3.4",
        "ripemd160": "^2.0.1",
        "sha.js": "^2.4.0"
      }
    },
    "create-hmac": {
      "version": "1.1.7",
      "resolved": "http://registry.npm.taobao.org/create-hmac/download/create-hmac-1.1.7.tgz",
      "integrity": "sha1-aRcMeLOrlXFHsriwRXLkfq0iQ/8=",
      "dev": true,
      "requires": {
        "cipher-base": "^1.0.3",
        "create-hash": "^1.1.0",
        "inherits": "^2.0.1",
        "ripemd160": "^2.0.0",
        "safe-buffer": "^5.0.1",
        "sha.js": "^2.4.8"
      }
    },
    "cross-spawn": {
      "version": "6.0.5",
      "resolved": "http://registry.npm.taobao.org/cross-spawn/download/cross-spawn-6.0.5.tgz",
      "integrity": "sha1-Sl7Hxk364iw6FBJNus3uhG2Ay8Q=",
      "dev": true,
      "requires": {
        "nice-try": "^1.0.4",
        "path-key": "^2.0.1",
        "semver": "^5.5.0",
        "shebang-command": "^1.2.0",
        "which": "^1.2.9"
      }
    },
    "crypto-browserify": {
      "version": "3.12.0",
      "resolved": "http://registry.npm.taobao.org/crypto-browserify/download/crypto-browserify-3.12.0.tgz",
      "integrity": "sha1-OWz58xN/A+S45TLFj2mCVOAPgOw=",
      "dev": true,
      "requires": {
        "browserify-cipher": "^1.0.0",
        "browserify-sign": "^4.0.0",
        "create-ecdh": "^4.0.0",
        "create-hash": "^1.1.0",
        "create-hmac": "^1.1.0",
        "diffie-hellman": "^5.0.0",
        "inherits": "^2.0.1",
        "pbkdf2": "^3.0.3",
        "public-encrypt": "^4.0.0",
        "randombytes": "^2.0.0",
        "randomfill": "^1.0.3"
      }
    },
    "css-color-names": {
      "version": "0.0.4",
      "resolved": "http://registry.npm.taobao.org/css-color-names/download/css-color-names-0.0.4.tgz",
      "integrity": "sha1-gIrcLnnPhHOAabZGyyDsJ762KeA=",
      "dev": true
    },
    "css-declaration-sorter": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/css-declaration-sorter/download/css-declaration-sorter-4.0.1.tgz",
      "integrity": "sha1-wZiUD2OnbX42wecQGLABchBUyyI=",
      "dev": true,
      "requires": {
        "postcss": "^7.0.1",
        "timsort": "^0.3.0"
      }
    },
    "css-loader": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/css-loader/download/css-loader-1.0.1.tgz",
      "integrity": "sha1-aIW7UjOzXsR7AGBX2gHMZAtref4=",
      "dev": true,
      "requires": {
        "babel-code-frame": "^6.26.0",
        "css-selector-tokenizer": "^0.7.0",
        "icss-utils": "^2.1.0",
        "loader-utils": "^1.0.2",
        "lodash": "^4.17.11",
        "postcss": "^6.0.23",
        "postcss-modules-extract-imports": "^1.2.0",
        "postcss-modules-local-by-default": "^1.2.0",
        "postcss-modules-scope": "^1.1.0",
        "postcss-modules-values": "^1.3.0",
        "postcss-value-parser": "^3.3.0",
        "source-list-map": "^2.0.0"
      },
      "dependencies": {
        "postcss": {
          "version": "6.0.23",
          "resolved": "http://registry.npm.taobao.org/postcss/download/postcss-6.0.23.tgz",
          "integrity": "sha1-YcgswyisYOZ3ZF+XkFTrmLwOMyQ=",
          "dev": true,
          "requires": {
            "chalk": "^2.4.1",
            "source-map": "^0.6.1",
            "supports-color": "^5.4.0"
          }
        },
        "source-map": {
          "version": "0.6.1",
          "resolved": "http://registry.npm.taobao.org/source-map/download/source-map-0.6.1.tgz",
          "integrity": "sha1-dHIq8y6WFOnCh6jQu95IteLxomM=",
          "dev": true
        }
      }
    },
    "css-select": {
      "version": "2.0.2",
      "resolved": "http://registry.npm.taobao.org/css-select/download/css-select-2.0.2.tgz",
      "integrity": "sha1-q0OGzsnh9miFVWSxfDcztDsqXt4=",
      "dev": true,
      "requires": {
        "boolbase": "^1.0.0",
        "css-what": "^2.1.2",
        "domutils": "^1.7.0",
        "nth-check": "^1.0.2"
      }
    },
    "css-select-base-adapter": {
      "version": "0.1.1",
      "resolved": "http://registry.npm.taobao.org/css-select-base-adapter/download/css-select-base-adapter-0.1.1.tgz",
      "integrity": "sha1-Oy/0lyzDYquIVhUHqVQIoUMhNdc=",
      "dev": true
    },
    "css-selector-tokenizer": {
      "version": "0.7.1",
      "resolved": "http://registry.npm.taobao.org/css-selector-tokenizer/download/css-selector-tokenizer-0.7.1.tgz",
      "integrity": "sha1-oXcnGovKUBkXL0+JH8bu2cv2jV0=",
      "dev": true,
      "requires": {
        "cssesc": "^0.1.0",
        "fastparse": "^1.1.1",
        "regexpu-core": "^1.0.0"
      },
      "dependencies": {
        "cssesc": {
          "version": "0.1.0",
          "resolved": "http://registry.npm.taobao.org/cssesc/download/cssesc-0.1.0.tgz",
          "integrity": "sha1-yBSQPkViM3GgR3tAEJqq++6t27Q=",
          "dev": true
        },
        "jsesc": {
          "version": "0.5.0",
          "resolved": "http://registry.npm.taobao.org/jsesc/download/jsesc-0.5.0.tgz",
          "integrity": "sha1-597mbjXW/Bb3EP6R1c9p9w8IkR0=",
          "dev": true
        },
        "regexpu-core": {
          "version": "1.0.0",
          "resolved": "http://registry.npm.taobao.org/regexpu-core/download/regexpu-core-1.0.0.tgz",
          "integrity": "sha1-hqdj9Y7k18L2sQLkdkBQ3n7ZDGs=",
          "dev": true,
          "requires": {
            "regenerate": "^1.2.1",
            "regjsgen": "^0.2.0",
            "regjsparser": "^0.1.4"
          }
        },
        "regjsgen": {
          "version": "0.2.0",
          "resolved": "http://registry.npm.taobao.org/regjsgen/download/regjsgen-0.2.0.tgz",
          "integrity": "sha1-bAFq3qxVT3WCP+N6wFuS1aTtsfc=",
          "dev": true
        },
        "regjsparser": {
          "version": "0.1.5",
          "resolved": "http://registry.npm.taobao.org/regjsparser/download/regjsparser-0.1.5.tgz",
          "integrity": "sha1-fuj4Tcb6eS0/0K4ijSS9lJ6tIFw=",
          "dev": true,
          "requires": {
            "jsesc": "~0.5.0"
          }
        }
      }
    },
    "css-tree": {
      "version": "1.0.0-alpha.28",
      "resolved": "http://registry.npm.taobao.org/css-tree/download/css-tree-1.0.0-alpha.28.tgz",
      "integrity": "sha1-joloGQ2IbJR3vI1h6W9hrz9/+n8=",
      "dev": true,
      "requires": {
        "mdn-data": "~1.1.0",
        "source-map": "^0.5.3"
      }
    },
    "css-unit-converter": {
      "version": "1.1.1",
      "resolved": "http://registry.npm.taobao.org/css-unit-converter/download/css-unit-converter-1.1.1.tgz",
      "integrity": "sha1-2bkoGtz9jO2TW9urqDeGiX9k6ZY=",
      "dev": true
    },
    "css-url-regex": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/css-url-regex/download/css-url-regex-1.1.0.tgz",
      "integrity": "sha1-g4NCMMyfdMRX3lnuvRVD/uuDt+w=",
      "dev": true
    },
    "css-what": {
      "version": "2.1.2",
      "resolved": "http://registry.npm.taobao.org/css-what/download/css-what-2.1.2.tgz",
      "integrity": "sha1-wIdtnQSAkn19SSDc1yrzWVZJVU0=",
      "dev": true
    },
    "cssesc": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/cssesc/download/cssesc-2.0.0.tgz",
      "integrity": "sha1-OxO9G7HLNuG8taTc0n9UxdyzVwM=",
      "dev": true
    },
    "cssnano": {
      "version": "4.1.7",
      "resolved": "http://registry.npm.taobao.org/cssnano/download/cssnano-4.1.7.tgz",
      "integrity": "sha1-C/ESKUvsEDq19o0/gFcyyDJaCxs=",
      "dev": true,
      "requires": {
        "cosmiconfig": "^5.0.0",
        "cssnano-preset-default": "^4.0.5",
        "is-resolvable": "^1.0.0",
        "postcss": "^7.0.0"
      }
    },
    "cssnano-preset-default": {
      "version": "4.0.5",
      "resolved": "http://registry.npm.taobao.org/cssnano-preset-default/download/cssnano-preset-default-4.0.5.tgz",
      "integrity": "sha1-0XVsAlnZitMR5gG6dulcYPZ3GsE=",
      "dev": true,
      "requires": {
        "css-declaration-sorter": "^4.0.1",
        "cssnano-util-raw-cache": "^4.0.1",
        "postcss": "^7.0.0",
        "postcss-calc": "^7.0.0",
        "postcss-colormin": "^4.0.2",
        "postcss-convert-values": "^4.0.1",
        "postcss-discard-comments": "^4.0.1",
        "postcss-discard-duplicates": "^4.0.2",
        "postcss-discard-empty": "^4.0.1",
        "postcss-discard-overridden": "^4.0.1",
        "postcss-merge-longhand": "^4.0.9",
        "postcss-merge-rules": "^4.0.2",
        "postcss-minify-font-values": "^4.0.2",
        "postcss-minify-gradients": "^4.0.1",
        "postcss-minify-params": "^4.0.1",
        "postcss-minify-selectors": "^4.0.1",
        "postcss-normalize-charset": "^4.0.1",
        "postcss-normalize-display-values": "^4.0.1",
        "postcss-normalize-positions": "^4.0.1",
        "postcss-normalize-repeat-style": "^4.0.1",
        "postcss-normalize-string": "^4.0.1",
        "postcss-normalize-timing-functions": "^4.0.1",
        "postcss-normalize-unicode": "^4.0.1",
        "postcss-normalize-url": "^4.0.1",
        "postcss-normalize-whitespace": "^4.0.1",
        "postcss-ordered-values": "^4.1.1",
        "postcss-reduce-initial": "^4.0.2",
        "postcss-reduce-transforms": "^4.0.1",
        "postcss-svgo": "^4.0.1",
        "postcss-unique-selectors": "^4.0.1"
      }
    },
    "cssnano-util-get-arguments": {
      "version": "4.0.0",
      "resolved": "http://registry.npm.taobao.org/cssnano-util-get-arguments/download/cssnano-util-get-arguments-4.0.0.tgz",
      "integrity": "sha1-7ToIKZ8h11dBsg87gfGU7UnMFQ8=",
      "dev": true
    },
    "cssnano-util-get-match": {
      "version": "4.0.0",
      "resolved": "http://registry.npm.taobao.org/cssnano-util-get-match/download/cssnano-util-get-match-4.0.0.tgz",
      "integrity": "sha1-wOTKB/U4a7F+xeUiULT1lhNlFW0=",
      "dev": true
    },
    "cssnano-util-raw-cache": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/cssnano-util-raw-cache/download/cssnano-util-raw-cache-4.0.1.tgz",
      "integrity": "sha1-sm1f1fcqEd/np4RvtMZyYPlr8oI=",
      "dev": true,
      "requires": {
        "postcss": "^7.0.0"
      }
    },
    "cssnano-util-same-parent": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/cssnano-util-same-parent/download/cssnano-util-same-parent-4.0.1.tgz",
      "integrity": "sha1-V0CC+yhZ0ttDOFWDXZqEVuoYu/M=",
      "dev": true
    },
    "csso": {
      "version": "3.5.1",
      "resolved": "http://registry.npm.taobao.org/csso/download/csso-3.5.1.tgz",
      "integrity": "sha1-e564vmFiiXPBsmHhadLwJACOdYs=",
      "dev": true,
      "requires": {
        "css-tree": "1.0.0-alpha.29"
      },
      "dependencies": {
        "css-tree": {
          "version": "1.0.0-alpha.29",
          "resolved": "http://registry.npm.taobao.org/css-tree/download/css-tree-1.0.0-alpha.29.tgz",
          "integrity": "sha1-P6nU7zFCy9HDAedmTB81K9gvWjk=",
          "dev": true,
          "requires": {
            "mdn-data": "~1.1.0",
            "source-map": "^0.5.3"
          }
        }
      }
    },
    "cyclist": {
      "version": "0.2.2",
      "resolved": "http://registry.npm.taobao.org/cyclist/download/cyclist-0.2.2.tgz",
      "integrity": "sha1-GzN5LhHpFKL9bW7WRHRkRE5fpkA=",
      "dev": true
    },
    "dashdash": {
      "version": "1.14.1",
      "resolved": "http://registry.npm.taobao.org/dashdash/download/dashdash-1.14.1.tgz",
      "integrity": "sha1-hTz6D3y+L+1d4gMmuN1YEDX24vA=",
      "dev": true,
      "requires": {
        "assert-plus": "^1.0.0"
      }
    },
    "date-now": {
      "version": "0.1.4",
      "resolved": "http://registry.npm.taobao.org/date-now/download/date-now-0.1.4.tgz",
      "integrity": "sha1-6vQ5/U1ISK105cx9vvIAZyueNFs=",
      "dev": true
    },
    "de-indent": {
      "version": "1.0.2",
      "resolved": "http://registry.npm.taobao.org/de-indent/download/de-indent-1.0.2.tgz",
      "integrity": "sha1-sgOOhG3DO6pXlhKNCAS0VbjB4h0=",
      "dev": true
    },
    "debug": {
      "version": "4.1.0",
      "resolved": "http://registry.npm.taobao.org/debug/download/debug-4.1.0.tgz",
      "integrity": "sha1-NzaHv/pnizixzZH4YbY4UANd3Ic=",
      "dev": true,
      "requires": {
        "ms": "^2.1.1"
      }
    },
    "decamelize": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/decamelize/download/decamelize-2.0.0.tgz",
      "integrity": "sha1-ZW17vICUxMeI6lPFhAkIycfQY8c=",
      "dev": true,
      "requires": {
        "xregexp": "4.0.0"
      }
    },
    "decode-uri-component": {
      "version": "0.2.0",
      "resolved": "http://registry.npm.taobao.org/decode-uri-component/download/decode-uri-component-0.2.0.tgz",
      "integrity": "sha1-6zkTMzRYd1y4TNGh+uBiEGu4dUU=",
      "dev": true
    },
    "deep-equal": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/deep-equal/download/deep-equal-1.0.1.tgz",
      "integrity": "sha1-9dJgKStmDghO/0zbyfCK0yR0SLU=",
      "dev": true
    },
    "deep-is": {
      "version": "0.1.3",
      "resolved": "http://registry.npm.taobao.org/deep-is/download/deep-is-0.1.3.tgz",
      "integrity": "sha1-s2nW+128E+7PUk+RsHD+7cNXzzQ=",
      "dev": true
    },
    "deepmerge": {
      "version": "1.5.2",
      "resolved": "http://registry.npm.taobao.org/deepmerge/download/deepmerge-1.5.2.tgz",
      "integrity": "sha1-EEmdhohEza1P7ghC34x/bwyVp1M=",
      "dev": true
    },
    "default-gateway": {
      "version": "2.7.2",
      "resolved": "http://registry.npm.taobao.org/default-gateway/download/default-gateway-2.7.2.tgz",
      "integrity": "sha1-t+8znl4CSwRUZ69APVA0jbRkLQ8=",
      "dev": true,
      "requires": {
        "execa": "^0.10.0",
        "ip-regex": "^2.1.0"
      },
      "dependencies": {
        "execa": {
          "version": "0.10.0",
          "resolved": "http://registry.npm.taobao.org/execa/download/execa-0.10.0.tgz",
          "integrity": "sha1-/0Vqj1P5D47MxxqW0Rvfx/CCy1A=",
          "dev": true,
          "requires": {
            "cross-spawn": "^6.0.0",
            "get-stream": "^3.0.0",
            "is-stream": "^1.1.0",
            "npm-run-path": "^2.0.0",
            "p-finally": "^1.0.0",
            "signal-exit": "^3.0.0",
            "strip-eof": "^1.0.0"
          }
        },
        "get-stream": {
          "version": "3.0.0",
          "resolved": "http://registry.npm.taobao.org/get-stream/download/get-stream-3.0.0.tgz",
          "integrity": "sha1-jpQ9E1jcN1VQVOy+LtsFqhdO3hQ=",
          "dev": true
        }
      }
    },
    "defaults": {
      "version": "1.0.3",
      "resolved": "http://registry.npm.taobao.org/defaults/download/defaults-1.0.3.tgz",
      "integrity": "sha1-xlYFHpgX2f8I7YgUd/P+QBnz730=",
      "dev": true,
      "requires": {
        "clone": "^1.0.2"
      }
    },
    "define-properties": {
      "version": "1.1.3",
      "resolved": "http://registry.npm.taobao.org/define-properties/download/define-properties-1.1.3.tgz",
      "integrity": "sha1-z4jabL7ib+bbcJT2HYcMvYTO6fE=",
      "dev": true,
      "requires": {
        "object-keys": "^1.0.12"
      }
    },
    "define-property": {
      "version": "2.0.2",
      "resolved": "http://registry.npm.taobao.org/define-property/download/define-property-2.0.2.tgz",
      "integrity": "sha1-1Flono1lS6d+AqgX+HENcCyxbp0=",
      "dev": true,
      "requires": {
        "is-descriptor": "^1.0.2",
        "isobject": "^3.0.1"
      },
      "dependencies": {
        "is-accessor-descriptor": {
          "version": "1.0.0",
          "resolved": "http://registry.npm.taobao.org/is-accessor-descriptor/download/is-accessor-descriptor-1.0.0.tgz",
          "integrity": "sha1-FpwvbT3x+ZJhgHI2XJsOofaHhlY=",
          "dev": true,
          "requires": {
            "kind-of": "^6.0.0"
          }
        },
        "is-data-descriptor": {
          "version": "1.0.0",
          "resolved": "http://registry.npm.taobao.org/is-data-descriptor/download/is-data-descriptor-1.0.0.tgz",
          "integrity": "sha1-2Eh2Mh0Oet0DmQQGq7u9NrqSaMc=",
          "dev": true,
          "requires": {
            "kind-of": "^6.0.0"
          }
        },
        "is-descriptor": {
          "version": "1.0.2",
          "resolved": "http://registry.npm.taobao.org/is-descriptor/download/is-descriptor-1.0.2.tgz",
          "integrity": "sha1-OxWXRqZmBLBPjIFSS6NlxfFNhuw=",
          "dev": true,
          "requires": {
            "is-accessor-descriptor": "^1.0.0",
            "is-data-descriptor": "^1.0.0",
            "kind-of": "^6.0.2"
          }
        }
      }
    },
    "del": {
      "version": "3.0.0",
      "resolved": "http://registry.npm.taobao.org/del/download/del-3.0.0.tgz",
      "integrity": "sha1-U+z2mf/LyzljdpGrE7rxYIGXZuU=",
      "dev": true,
      "requires": {
        "globby": "^6.1.0",
        "is-path-cwd": "^1.0.0",
        "is-path-in-cwd": "^1.0.0",
        "p-map": "^1.1.1",
        "pify": "^3.0.0",
        "rimraf": "^2.2.8"
      },
      "dependencies": {
        "globby": {
          "version": "6.1.0",
          "resolved": "http://registry.npm.taobao.org/globby/download/globby-6.1.0.tgz",
          "integrity": "sha1-9abXDoOV4hyFj7BInWTfAkJNUGw=",
          "dev": true,
          "requires": {
            "array-union": "^1.0.1",
            "glob": "^7.0.3",
            "object-assign": "^4.0.1",
            "pify": "^2.0.0",
            "pinkie-promise": "^2.0.0"
          },
          "dependencies": {
            "pify": {
              "version": "2.3.0",
              "resolved": "http://registry.npm.taobao.org/pify/download/pify-2.3.0.tgz",
              "integrity": "sha1-7RQaasBDqEnqWISY59yosVMw6Qw=",
              "dev": true
            }
          }
        }
      }
    },
    "delayed-stream": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/delayed-stream/download/delayed-stream-1.0.0.tgz",
      "integrity": "sha1-3zrhmayt+31ECqrgsp4icrJOxhk=",
      "dev": true
    },
    "depd": {
      "version": "1.1.2",
      "resolved": "http://registry.npm.taobao.org/depd/download/depd-1.1.2.tgz",
      "integrity": "sha1-m81S4UwJd2PnSbJ0xDRu0uVgtak=",
      "dev": true
    },
    "des.js": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/des.js/download/des.js-1.0.0.tgz",
      "integrity": "sha1-wHTS4qpqipoH29YfmhXCzYPsjsw=",
      "dev": true,
      "requires": {
        "inherits": "^2.0.1",
        "minimalistic-assert": "^1.0.0"
      }
    },
    "destroy": {
      "version": "1.0.4",
      "resolved": "http://registry.npm.taobao.org/destroy/download/destroy-1.0.4.tgz",
      "integrity": "sha1-l4hXRCxEdJ5CBmE+N5RiBYJqvYA=",
      "dev": true
    },
    "detect-node": {
      "version": "2.0.4",
      "resolved": "http://registry.npm.taobao.org/detect-node/download/detect-node-2.0.4.tgz",
      "integrity": "sha1-AU7o+PZpxcWAI9pkuBecCDooxGw=",
      "dev": true
    },
    "diffie-hellman": {
      "version": "5.0.3",
      "resolved": "http://registry.npm.taobao.org/diffie-hellman/download/diffie-hellman-5.0.3.tgz",
      "integrity": "sha1-QOjumPVaIUlgcUaSHGPhrl89KHU=",
      "dev": true,
      "requires": {
        "bn.js": "^4.1.0",
        "miller-rabin": "^4.0.0",
        "randombytes": "^2.0.0"
      }
    },
    "dir-glob": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/dir-glob/download/dir-glob-2.0.0.tgz",
      "integrity": "sha1-CyBdK2rvmCOMooZZioIE0p0KADQ=",
      "dev": true,
      "requires": {
        "arrify": "^1.0.1",
        "path-type": "^3.0.0"
      }
    },
    "dns-equal": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/dns-equal/download/dns-equal-1.0.0.tgz",
      "integrity": "sha1-s55/HabrCnW6nBcySzR1PEfgZU0=",
      "dev": true
    },
    "dns-packet": {
      "version": "1.3.1",
      "resolved": "http://registry.npm.taobao.org/dns-packet/download/dns-packet-1.3.1.tgz",
      "integrity": "sha1-EqpCaYEHW+UAuRDu3NC0fdfe2lo=",
      "dev": true,
      "requires": {
        "ip": "^1.1.0",
        "safe-buffer": "^5.0.1"
      }
    },
    "dns-txt": {
      "version": "2.0.2",
      "resolved": "http://registry.npm.taobao.org/dns-txt/download/dns-txt-2.0.2.tgz",
      "integrity": "sha1-uR2Ab10nGI5Ks+fRB9iBocxGQrY=",
      "dev": true,
      "requires": {
        "buffer-indexof": "^1.0.0"
      }
    },
    "doctrine": {
      "version": "2.1.0",
      "resolved": "http://registry.npm.taobao.org/doctrine/download/doctrine-2.1.0.tgz",
      "integrity": "sha1-XNAfwQFiG0LEzX9dGmYkNxbT850=",
      "dev": true,
      "requires": {
        "esutils": "^2.0.2"
      }
    },
    "dom-converter": {
      "version": "0.2.0",
      "resolved": "http://registry.npm.taobao.org/dom-converter/download/dom-converter-0.2.0.tgz",
      "integrity": "sha1-ZyGp2u4uKTaClVtq/kFncWJ7t2g=",
      "dev": true,
      "requires": {
        "utila": "~0.4"
      }
    },
    "dom-serializer": {
      "version": "0.1.0",
      "resolved": "http://registry.npm.taobao.org/dom-serializer/download/dom-serializer-0.1.0.tgz",
      "integrity": "sha1-BzxpdUbOB4DOI75KKOKT5AvDDII=",
      "dev": true,
      "requires": {
        "domelementtype": "~1.1.1",
        "entities": "~1.1.1"
      },
      "dependencies": {
        "domelementtype": {
          "version": "1.1.3",
          "resolved": "http://registry.npm.taobao.org/domelementtype/download/domelementtype-1.1.3.tgz",
          "integrity": "sha1-vSh3PiZCiBrsUVRJJCmcXNgiGFs=",
          "dev": true
        }
      }
    },
    "domain-browser": {
      "version": "1.2.0",
      "resolved": "http://registry.npm.taobao.org/domain-browser/download/domain-browser-1.2.0.tgz",
      "integrity": "sha1-PTH1AZGmdJ3RN1p/Ui6CPULlTto=",
      "dev": true
    },
    "domelementtype": {
      "version": "1.3.1",
      "resolved": "http://registry.npm.taobao.org/domelementtype/download/domelementtype-1.3.1.tgz",
      "integrity": "sha1-0EjESzew0Qp/Kj1f7j9DM9eQSB8=",
      "dev": true
    },
    "domhandler": {
      "version": "2.1.0",
      "resolved": "http://registry.npm.taobao.org/domhandler/download/domhandler-2.1.0.tgz",
      "integrity": "sha1-0mRvXlf2w7qxHPbLBdPArPdBJZQ=",
      "dev": true,
      "requires": {
        "domelementtype": "1"
      }
    },
    "domutils": {
      "version": "1.7.0",
      "resolved": "http://registry.npm.taobao.org/domutils/download/domutils-1.7.0.tgz",
      "integrity": "sha1-Vuo0HoNOBuZ0ivehyyXaZ+qfjCo=",
      "dev": true,
      "requires": {
        "dom-serializer": "0",
        "domelementtype": "1"
      }
    },
    "dot-prop": {
      "version": "4.2.0",
      "resolved": "http://registry.npm.taobao.org/dot-prop/download/dot-prop-4.2.0.tgz",
      "integrity": "sha1-HxngwuGqDjJ5fEl5nyg3rGr2nFc=",
      "dev": true,
      "requires": {
        "is-obj": "^1.0.0"
      }
    },
    "duplexer": {
      "version": "0.1.1",
      "resolved": "http://registry.npm.taobao.org/duplexer/download/duplexer-0.1.1.tgz",
      "integrity": "sha1-rOb/gIwc5mtX0ev5eXessCM0z8E=",
      "dev": true
    },
    "duplexify": {
      "version": "3.6.1",
      "resolved": "http://registry.npm.taobao.org/duplexify/download/duplexify-3.6.1.tgz",
      "integrity": "sha1-saeinEq/1jlYXvrszoDWZrHjQSU=",
      "dev": true,
      "requires": {
        "end-of-stream": "^1.0.0",
        "inherits": "^2.0.1",
        "readable-stream": "^2.0.0",
        "stream-shift": "^1.0.0"
      }
    },
    "easy-stack": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/easy-stack/download/easy-stack-1.0.0.tgz",
      "integrity": "sha1-EskbMIWjfwuqM26UhurEv5Tj54g=",
      "dev": true
    },
    "ecc-jsbn": {
      "version": "0.1.2",
      "resolved": "http://registry.npm.taobao.org/ecc-jsbn/download/ecc-jsbn-0.1.2.tgz",
      "integrity": "sha1-OoOpBOVDUyh4dMVkt1SThoSamMk=",
      "dev": true,
      "requires": {
        "jsbn": "~0.1.0",
        "safer-buffer": "^2.1.0"
      }
    },
    "ee-first": {
      "version": "1.1.1",
      "resolved": "http://registry.npm.taobao.org/ee-first/download/ee-first-1.1.1.tgz",
      "integrity": "sha1-WQxhFWsK4vTwJVcyoViyZrxWsh0=",
      "dev": true
    },
    "ejs": {
      "version": "2.6.1",
      "resolved": "http://registry.npm.taobao.org/ejs/download/ejs-2.6.1.tgz",
      "integrity": "sha1-SY7A1JVlWrxvI81hho2SZGQHGqA=",
      "dev": true
    },
    "electron-to-chromium": {
      "version": "1.3.90",
      "resolved": "http://registry.npm.taobao.org/electron-to-chromium/download/electron-to-chromium-1.3.90.tgz",
      "integrity": "sha1-tMUbgwO+/xjyt0gXQCv0iY4JVYo=",
      "dev": true
    },
    "elliptic": {
      "version": "6.4.1",
      "resolved": "http://registry.npm.taobao.org/elliptic/download/elliptic-6.4.1.tgz",
      "integrity": "sha1-wtC3d2kRuGcixjLDwGxg8vgZk5o=",
      "dev": true,
      "requires": {
        "bn.js": "^4.4.0",
        "brorand": "^1.0.1",
        "hash.js": "^1.0.0",
        "hmac-drbg": "^1.0.0",
        "inherits": "^2.0.1",
        "minimalistic-assert": "^1.0.0",
        "minimalistic-crypto-utils": "^1.0.0"
      }
    },
    "emojis-list": {
      "version": "2.1.0",
      "resolved": "http://registry.npm.taobao.org/emojis-list/download/emojis-list-2.1.0.tgz",
      "integrity": "sha1-TapNnbAPmBmIDHn6RXrlsJof04k=",
      "dev": true
    },
    "encodeurl": {
      "version": "1.0.2",
      "resolved": "http://registry.npm.taobao.org/encodeurl/download/encodeurl-1.0.2.tgz",
      "integrity": "sha1-rT/0yG7C0CkyL1oCw6mmBslbP1k=",
      "dev": true
    },
    "end-of-stream": {
      "version": "1.4.1",
      "resolved": "http://registry.npm.taobao.org/end-of-stream/download/end-of-stream-1.4.1.tgz",
      "integrity": "sha1-7SljTRm6ukY7bOa4CjchPqtx7EM=",
      "dev": true,
      "requires": {
        "once": "^1.4.0"
      }
    },
    "enhanced-resolve": {
      "version": "4.1.0",
      "resolved": "http://registry.npm.taobao.org/enhanced-resolve/download/enhanced-resolve-4.1.0.tgz",
      "integrity": "sha1-Qcfgv9/nSsH/4eV61qXGyfN0Kn8=",
      "dev": true,
      "requires": {
        "graceful-fs": "^4.1.2",
        "memory-fs": "^0.4.0",
        "tapable": "^1.0.0"
      }
    },
    "entities": {
      "version": "1.1.2",
      "resolved": "http://registry.npm.taobao.org/entities/download/entities-1.1.2.tgz",
      "integrity": "sha1-vfpzUplmTfr9NFKe1PhSKidf6lY=",
      "dev": true
    },
    "errno": {
      "version": "0.1.7",
      "resolved": "http://registry.npm.taobao.org/errno/download/errno-0.1.7.tgz",
      "integrity": "sha1-RoTXF3mtOa8Xfj8AeZb3xnyFJhg=",
      "dev": true,
      "requires": {
        "prr": "~1.0.1"
      }
    },
    "error-ex": {
      "version": "1.3.2",
      "resolved": "http://registry.npm.taobao.org/error-ex/download/error-ex-1.3.2.tgz",
      "integrity": "sha1-tKxAZIEH/c3PriQvQovqihTU8b8=",
      "dev": true,
      "requires": {
        "is-arrayish": "^0.2.1"
      }
    },
    "error-stack-parser": {
      "version": "2.0.2",
      "resolved": "http://registry.npm.taobao.org/error-stack-parser/download/error-stack-parser-2.0.2.tgz",
      "integrity": "sha1-Sujbqiv5CotFBwe5FJ3KvKE1Ug0=",
      "dev": true,
      "requires": {
        "stackframe": "^1.0.4"
      }
    },
    "es-abstract": {
      "version": "1.12.0",
      "resolved": "http://registry.npm.taobao.org/es-abstract/download/es-abstract-1.12.0.tgz",
      "integrity": "sha1-nbvdJ8aFbwABQhyhh4LXhr+KYWU=",
      "dev": true,
      "requires": {
        "es-to-primitive": "^1.1.1",
        "function-bind": "^1.1.1",
        "has": "^1.0.1",
        "is-callable": "^1.1.3",
        "is-regex": "^1.0.4"
      }
    },
    "es-to-primitive": {
      "version": "1.2.0",
      "resolved": "http://registry.npm.taobao.org/es-to-primitive/download/es-to-primitive-1.2.0.tgz",
      "integrity": "sha1-7fckeAM0VujdqO8J4ArZZQcH83c=",
      "dev": true,
      "requires": {
        "is-callable": "^1.1.4",
        "is-date-object": "^1.0.1",
        "is-symbol": "^1.0.2"
      }
    },
    "escape-html": {
      "version": "1.0.3",
      "resolved": "http://registry.npm.taobao.org/escape-html/download/escape-html-1.0.3.tgz",
      "integrity": "sha1-Aljq5NPQwJdN4cFpGI7wBR0dGYg=",
      "dev": true
    },
    "escape-string-regexp": {
      "version": "1.0.5",
      "resolved": "http://registry.npm.taobao.org/escape-string-regexp/download/escape-string-regexp-1.0.5.tgz",
      "integrity": "sha1-G2HAViGQqN/2rjuyzwIAyhMLhtQ=",
      "dev": true
    },
    "eslint": {
      "version": "5.10.0",
      "resolved": "http://registry.npm.taobao.org/eslint/download/eslint-5.10.0.tgz",
      "integrity": "sha1-JK3L6Sv16x/C0vKx7r4MXgcTkDo=",
      "dev": true,
      "requires": {
        "@babel/code-frame": "^7.0.0",
        "ajv": "^6.5.3",
        "chalk": "^2.1.0",
        "cross-spawn": "^6.0.5",
        "debug": "^4.0.1",
        "doctrine": "^2.1.0",
        "eslint-scope": "^4.0.0",
        "eslint-utils": "^1.3.1",
        "eslint-visitor-keys": "^1.0.0",
        "espree": "^5.0.0",
        "esquery": "^1.0.1",
        "esutils": "^2.0.2",
        "file-entry-cache": "^2.0.0",
        "functional-red-black-tree": "^1.0.1",
        "glob": "^7.1.2",
        "globals": "^11.7.0",
        "ignore": "^4.0.6",
        "imurmurhash": "^0.1.4",
        "inquirer": "^6.1.0",
        "js-yaml": "^3.12.0",
        "json-stable-stringify-without-jsonify": "^1.0.1",
        "levn": "^0.3.0",
        "lodash": "^4.17.5",
        "minimatch": "^3.0.4",
        "mkdirp": "^0.5.1",
        "natural-compare": "^1.4.0",
        "optionator": "^0.8.2",
        "path-is-inside": "^1.0.2",
        "pluralize": "^7.0.0",
        "progress": "^2.0.0",
        "regexpp": "^2.0.1",
        "require-uncached": "^1.0.3",
        "semver": "^5.5.1",
        "strip-ansi": "^4.0.0",
        "strip-json-comments": "^2.0.1",
        "table": "^5.0.2",
        "text-table": "^0.2.0"
      },
      "dependencies": {
        "acorn": {
          "version": "6.0.4",
          "resolved": "http://registry.npm.taobao.org/acorn/download/acorn-6.0.4.tgz",
          "integrity": "sha1-dzd+c1O3LsUQRVCqLSCXov1At1Q=",
          "dev": true
        },
        "acorn-jsx": {
          "version": "5.0.1",
          "resolved": "http://registry.npm.taobao.org/acorn-jsx/download/acorn-jsx-5.0.1.tgz",
          "integrity": "sha1-MqBk/ZJUKSFqCbFBECv90YX65A4=",
          "dev": true
        },
        "ansi-regex": {
          "version": "4.0.0",
          "resolved": "http://registry.npm.taobao.org/ansi-regex/download/ansi-regex-4.0.0.tgz",
          "integrity": "sha1-cN55Ht8CFATD/WFaqJEYrgQy5ak=",
          "dev": true
        },
        "chardet": {
          "version": "0.7.0",
          "resolved": "http://registry.npm.taobao.org/chardet/download/chardet-0.7.0.tgz",
          "integrity": "sha1-kAlISfCTfy7twkJdDSip5fDLrZ4=",
          "dev": true
        },
        "eslint-scope": {
          "version": "4.0.0",
          "resolved": "http://registry.npm.taobao.org/eslint-scope/download/eslint-scope-4.0.0.tgz",
          "integrity": "sha1-UL8wcekzi83EMzF5Sgy1M/ATYXI=",
          "dev": true,
          "requires": {
            "esrecurse": "^4.1.0",
            "estraverse": "^4.1.1"
          }
        },
        "espree": {
          "version": "5.0.0",
          "resolved": "http://registry.npm.taobao.org/espree/download/espree-5.0.0.tgz",
          "integrity": "sha1-/H+YS2Kzag9UOxP7nNe59Kf1tlw=",
          "dev": true,
          "requires": {
            "acorn": "^6.0.2",
            "acorn-jsx": "^5.0.0",
            "eslint-visitor-keys": "^1.0.0"
          }
        },
        "external-editor": {
          "version": "3.0.3",
          "resolved": "http://registry.npm.taobao.org/external-editor/download/external-editor-3.0.3.tgz",
          "integrity": "sha1-WGbbKal4Jtvkvzr9JAcOrZ6kOic=",
          "dev": true,
          "requires": {
            "chardet": "^0.7.0",
            "iconv-lite": "^0.4.24",
            "tmp": "^0.0.33"
          }
        },
        "ignore": {
          "version": "4.0.6",
          "resolved": "http://registry.npm.taobao.org/ignore/download/ignore-4.0.6.tgz",
          "integrity": "sha1-dQ49tYYgh7RzfrrIIH/9HvJ7Jfw=",
          "dev": true
        },
        "inquirer": {
          "version": "6.2.1",
          "resolved": "http://registry.npm.taobao.org/inquirer/download/inquirer-6.2.1.tgz",
          "integrity": "sha1-mUP8SIIWG9sLDJJ2dpx1sy2/zVI=",
          "dev": true,
          "requires": {
            "ansi-escapes": "^3.0.0",
            "chalk": "^2.0.0",
            "cli-cursor": "^2.1.0",
            "cli-width": "^2.0.0",
            "external-editor": "^3.0.0",
            "figures": "^2.0.0",
            "lodash": "^4.17.10",
            "mute-stream": "0.0.7",
            "run-async": "^2.2.0",
            "rxjs": "^6.1.0",
            "string-width": "^2.1.0",
            "strip-ansi": "^5.0.0",
            "through": "^2.3.6"
          },
          "dependencies": {
            "strip-ansi": {
              "version": "5.0.0",
              "resolved": "http://registry.npm.taobao.org/strip-ansi/download/strip-ansi-5.0.0.tgz",
              "integrity": "sha1-949otdCGbCCyybjGG1KYUI3IdW8=",
              "dev": true,
              "requires": {
                "ansi-regex": "^4.0.0"
              }
            }
          }
        },
        "regexpp": {
          "version": "2.0.1",
          "resolved": "http://registry.npm.taobao.org/regexpp/download/regexpp-2.0.1.tgz",
          "integrity": "sha1-jRnTHPYySCtYkEn4KB+T28uk0H8=",
          "dev": true
        },
        "slice-ansi": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/slice-ansi/download/slice-ansi-2.0.0.tgz",
          "integrity": "sha1-U3O9uFWbRWduhUHGaRbN1iUWEuc=",
          "dev": true,
          "requires": {
            "ansi-styles": "^3.2.0",
            "astral-regex": "^1.0.0",
            "is-fullwidth-code-point": "^2.0.0"
          }
        },
        "table": {
          "version": "5.1.1",
          "resolved": "http://registry.npm.taobao.org/table/download/table-5.1.1.tgz",
          "integrity": "sha1-kgMBkvG3tRtu6rI+1BaGLke3CDc=",
          "dev": true,
          "requires": {
            "ajv": "^6.6.1",
            "lodash": "^4.17.11",
            "slice-ansi": "2.0.0",
            "string-width": "^2.1.1"
          }
        }
      }
    },
    "eslint-loader": {
      "version": "2.1.1",
      "resolved": "http://registry.npm.taobao.org/eslint-loader/download/eslint-loader-2.1.1.tgz",
      "integrity": "sha1-KpJRUjZSQwv91kPv2wr8GiqJVGo=",
      "dev": true,
      "requires": {
        "loader-fs-cache": "^1.0.0",
        "loader-utils": "^1.0.2",
        "object-assign": "^4.0.1",
        "object-hash": "^1.1.4",
        "rimraf": "^2.6.1"
      }
    },
    "eslint-plugin-vue": {
      "version": "5.0.0",
      "resolved": "http://registry.npm.taobao.org/eslint-plugin-vue/download/eslint-plugin-vue-5.0.0.tgz",
      "integrity": "sha1-SizBwOcepF4b2cGmD5Jb/mi7VxA=",
      "dev": true,
      "requires": {
        "vue-eslint-parser": "^4.0.2"
      },
      "dependencies": {
        "acorn": {
          "version": "6.0.4",
          "resolved": "http://registry.npm.taobao.org/acorn/download/acorn-6.0.4.tgz",
          "integrity": "sha1-dzd+c1O3LsUQRVCqLSCXov1At1Q=",
          "dev": true
        },
        "acorn-jsx": {
          "version": "5.0.1",
          "resolved": "http://registry.npm.taobao.org/acorn-jsx/download/acorn-jsx-5.0.1.tgz",
          "integrity": "sha1-MqBk/ZJUKSFqCbFBECv90YX65A4=",
          "dev": true
        },
        "eslint-scope": {
          "version": "4.0.0",
          "resolved": "http://registry.npm.taobao.org/eslint-scope/download/eslint-scope-4.0.0.tgz",
          "integrity": "sha1-UL8wcekzi83EMzF5Sgy1M/ATYXI=",
          "dev": true,
          "requires": {
            "esrecurse": "^4.1.0",
            "estraverse": "^4.1.1"
          }
        },
        "espree": {
          "version": "4.1.0",
          "resolved": "http://registry.npm.taobao.org/espree/download/espree-4.1.0.tgz",
          "integrity": "sha1-co1UUeD9FWwEOEp62J7VH/VOsl8=",
          "dev": true,
          "requires": {
            "acorn": "^6.0.2",
            "acorn-jsx": "^5.0.0",
            "eslint-visitor-keys": "^1.0.0"
          }
        },
        "vue-eslint-parser": {
          "version": "4.0.3",
          "resolved": "http://registry.npm.taobao.org/vue-eslint-parser/download/vue-eslint-parser-4.0.3.tgz",
          "integrity": "sha1-gM8WLkhDh7JkA3GtIbofhuDBCmE=",
          "dev": true,
          "requires": {
            "debug": "^4.1.0",
            "eslint-scope": "^4.0.0",
            "eslint-visitor-keys": "^1.0.0",
            "espree": "^4.1.0",
            "esquery": "^1.0.1",
            "lodash": "^4.17.11"
          }
        }
      }
    },
    "eslint-scope": {
      "version": "3.7.1",
      "resolved": "http://registry.npm.taobao.org/eslint-scope/download/eslint-scope-3.7.1.tgz",
      "integrity": "sha1-PWPD7f2gLgbgGkUq2IyqzHzctug=",
      "dev": true,
      "requires": {
        "esrecurse": "^4.1.0",
        "estraverse": "^4.1.1"
      }
    },
    "eslint-utils": {
      "version": "1.3.1",
      "resolved": "http://registry.npm.taobao.org/eslint-utils/download/eslint-utils-1.3.1.tgz",
      "integrity": "sha1-moUbqJ7nxGA0b5fPiTnHKYgn5RI=",
      "dev": true
    },
    "eslint-visitor-keys": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/eslint-visitor-keys/download/eslint-visitor-keys-1.0.0.tgz",
      "integrity": "sha1-PzGA+y4pEBdxastMnW1bXDSmqB0=",
      "dev": true
    },
    "espree": {
      "version": "3.5.4",
      "resolved": "http://registry.npm.taobao.org/espree/download/espree-3.5.4.tgz",
      "integrity": "sha1-sPRHGHyKi+2US4FaZgvd9d610ac=",
      "dev": true,
      "requires": {
        "acorn": "^5.5.0",
        "acorn-jsx": "^3.0.0"
      }
    },
    "esprima": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/esprima/download/esprima-4.0.1.tgz",
      "integrity": "sha1-E7BM2z5sXRnfkatph6hpVhmwqnE=",
      "dev": true
    },
    "esquery": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/esquery/download/esquery-1.0.1.tgz",
      "integrity": "sha1-QGxRZYsfWZGl+bYrHcJbAOPlxwg=",
      "dev": true,
      "requires": {
        "estraverse": "^4.0.0"
      }
    },
    "esrecurse": {
      "version": "4.2.1",
      "resolved": "http://registry.npm.taobao.org/esrecurse/download/esrecurse-4.2.1.tgz",
      "integrity": "sha1-AHo7n9vCs7uH5IeeoZyS/b05Qs8=",
      "dev": true,
      "requires": {
        "estraverse": "^4.1.0"
      }
    },
    "estraverse": {
      "version": "4.2.0",
      "resolved": "http://registry.npm.taobao.org/estraverse/download/estraverse-4.2.0.tgz",
      "integrity": "sha1-De4/7TH81GlhjOc0IJn8GvoL2xM=",
      "dev": true
    },
    "esutils": {
      "version": "2.0.2",
      "resolved": "http://registry.npm.taobao.org/esutils/download/esutils-2.0.2.tgz",
      "integrity": "sha1-Cr9PHKpbyx96nYrMbepPqqBLrJs=",
      "dev": true
    },
    "etag": {
      "version": "1.8.1",
      "resolved": "http://registry.npm.taobao.org/etag/download/etag-1.8.1.tgz",
      "integrity": "sha1-Qa4u62XvpiJorr/qg6x9eSmbCIc=",
      "dev": true
    },
    "event-pubsub": {
      "version": "4.3.0",
      "resolved": "http://registry.npm.taobao.org/event-pubsub/download/event-pubsub-4.3.0.tgz",
      "integrity": "sha1-9o2Ba8KfHsAsU53FjI3UDOcss24=",
      "dev": true
    },
    "eventemitter3": {
      "version": "3.1.0",
      "resolved": "http://registry.npm.taobao.org/eventemitter3/download/eventemitter3-3.1.0.tgz",
      "integrity": "sha1-CQtNbNvWRe0Qv3UNS1QHlC17oWM=",
      "dev": true
    },
    "events": {
      "version": "1.1.1",
      "resolved": "http://registry.npm.taobao.org/events/download/events-1.1.1.tgz",
      "integrity": "sha1-nr23Y1rQmccNzEwqH1AEKI6L2SQ=",
      "dev": true
    },
    "eventsource": {
      "version": "1.0.7",
      "resolved": "http://registry.npm.taobao.org/eventsource/download/eventsource-1.0.7.tgz",
      "integrity": "sha1-j7xyyT/NNAiAkLwKTmT0tc7m2NA=",
      "dev": true,
      "requires": {
        "original": "^1.0.0"
      }
    },
    "evp_bytestokey": {
      "version": "1.0.3",
      "resolved": "http://registry.npm.taobao.org/evp_bytestokey/download/evp_bytestokey-1.0.3.tgz",
      "integrity": "sha1-f8vbGY3HGVlDLv4ThCaE4FJaywI=",
      "dev": true,
      "requires": {
        "md5.js": "^1.3.4",
        "safe-buffer": "^5.1.1"
      }
    },
    "execa": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/execa/download/execa-1.0.0.tgz",
      "integrity": "sha1-xiNqW7TfbW8V6I5/AXeYIWdJ3dg=",
      "dev": true,
      "requires": {
        "cross-spawn": "^6.0.0",
        "get-stream": "^4.0.0",
        "is-stream": "^1.1.0",
        "npm-run-path": "^2.0.0",
        "p-finally": "^1.0.0",
        "signal-exit": "^3.0.0",
        "strip-eof": "^1.0.0"
      }
    },
    "expand-brackets": {
      "version": "2.1.4",
      "resolved": "http://registry.npm.taobao.org/expand-brackets/download/expand-brackets-2.1.4.tgz",
      "integrity": "sha1-t3c14xXOMPa27/D4OwQVGiJEliI=",
      "dev": true,
      "requires": {
        "debug": "^2.3.3",
        "define-property": "^0.2.5",
        "extend-shallow": "^2.0.1",
        "posix-character-classes": "^0.1.0",
        "regex-not": "^1.0.0",
        "snapdragon": "^0.8.1",
        "to-regex": "^3.0.1"
      },
      "dependencies": {
        "debug": {
          "version": "2.6.9",
          "resolved": "http://registry.npm.taobao.org/debug/download/debug-2.6.9.tgz",
          "integrity": "sha1-XRKFFd8TT/Mn6QpMk/Tgd6U2NB8=",
          "dev": true,
          "requires": {
            "ms": "2.0.0"
          }
        },
        "define-property": {
          "version": "0.2.5",
          "resolved": "http://registry.npm.taobao.org/define-property/download/define-property-0.2.5.tgz",
          "integrity": "sha1-w1se+RjsPJkPmlvFe+BKrOxcgRY=",
          "dev": true,
          "requires": {
            "is-descriptor": "^0.1.0"
          }
        },
        "extend-shallow": {
          "version": "2.0.1",
          "resolved": "http://registry.npm.taobao.org/extend-shallow/download/extend-shallow-2.0.1.tgz",
          "integrity": "sha1-Ua99YUrZqfYQ6huvu5idaxxWiQ8=",
          "dev": true,
          "requires": {
            "is-extendable": "^0.1.0"
          }
        },
        "ms": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/ms/download/ms-2.0.0.tgz",
          "integrity": "sha1-VgiurfwAvmwpAd9fmGF4jeDVl8g=",
          "dev": true
        }
      }
    },
    "express": {
      "version": "4.16.4",
      "resolved": "http://registry.npm.taobao.org/express/download/express-4.16.4.tgz",
      "integrity": "sha1-/d72GSYQniTFFeqX/S8b2/Yt8S4=",
      "dev": true,
      "requires": {
        "accepts": "~1.3.5",
        "array-flatten": "1.1.1",
        "body-parser": "1.18.3",
        "content-disposition": "0.5.2",
        "content-type": "~1.0.4",
        "cookie": "0.3.1",
        "cookie-signature": "1.0.6",
        "debug": "2.6.9",
        "depd": "~1.1.2",
        "encodeurl": "~1.0.2",
        "escape-html": "~1.0.3",
        "etag": "~1.8.1",
        "finalhandler": "1.1.1",
        "fresh": "0.5.2",
        "merge-descriptors": "1.0.1",
        "methods": "~1.1.2",
        "on-finished": "~2.3.0",
        "parseurl": "~1.3.2",
        "path-to-regexp": "0.1.7",
        "proxy-addr": "~2.0.4",
        "qs": "6.5.2",
        "range-parser": "~1.2.0",
        "safe-buffer": "5.1.2",
        "send": "0.16.2",
        "serve-static": "1.13.2",
        "setprototypeof": "1.1.0",
        "statuses": "~1.4.0",
        "type-is": "~1.6.16",
        "utils-merge": "1.0.1",
        "vary": "~1.1.2"
      },
      "dependencies": {
        "debug": {
          "version": "2.6.9",
          "resolved": "http://registry.npm.taobao.org/debug/download/debug-2.6.9.tgz",
          "integrity": "sha1-XRKFFd8TT/Mn6QpMk/Tgd6U2NB8=",
          "dev": true,
          "requires": {
            "ms": "2.0.0"
          }
        },
        "ms": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/ms/download/ms-2.0.0.tgz",
          "integrity": "sha1-VgiurfwAvmwpAd9fmGF4jeDVl8g=",
          "dev": true
        }
      }
    },
    "extend": {
      "version": "3.0.2",
      "resolved": "http://registry.npm.taobao.org/extend/download/extend-3.0.2.tgz",
      "integrity": "sha1-+LETa0Bx+9jrFAr/hYsQGewpFfo=",
      "dev": true
    },
    "extend-shallow": {
      "version": "3.0.2",
      "resolved": "http://registry.npm.taobao.org/extend-shallow/download/extend-shallow-3.0.2.tgz",
      "integrity": "sha1-Jqcarwc7OfshJxcnRhMcJwQCjbg=",
      "dev": true,
      "requires": {
        "assign-symbols": "^1.0.0",
        "is-extendable": "^1.0.1"
      },
      "dependencies": {
        "is-extendable": {
          "version": "1.0.1",
          "resolved": "http://registry.npm.taobao.org/is-extendable/download/is-extendable-1.0.1.tgz",
          "integrity": "sha1-p0cPnkJnM9gb2B4RVSZOOjUHyrQ=",
          "dev": true,
          "requires": {
            "is-plain-object": "^2.0.4"
          }
        }
      }
    },
    "external-editor": {
      "version": "2.2.0",
      "resolved": "http://registry.npm.taobao.org/external-editor/download/external-editor-2.2.0.tgz",
      "integrity": "sha1-BFURz9jRM/OEZnPRBHwVTiFK09U=",
      "dev": true,
      "requires": {
        "chardet": "^0.4.0",
        "iconv-lite": "^0.4.17",
        "tmp": "^0.0.33"
      }
    },
    "extglob": {
      "version": "2.0.4",
      "resolved": "http://registry.npm.taobao.org/extglob/download/extglob-2.0.4.tgz",
      "integrity": "sha1-rQD+TcYSqSMuhxhxHcXLWrAoVUM=",
      "dev": true,
      "requires": {
        "array-unique": "^0.3.2",
        "define-property": "^1.0.0",
        "expand-brackets": "^2.1.4",
        "extend-shallow": "^2.0.1",
        "fragment-cache": "^0.2.1",
        "regex-not": "^1.0.0",
        "snapdragon": "^0.8.1",
        "to-regex": "^3.0.1"
      },
      "dependencies": {
        "define-property": {
          "version": "1.0.0",
          "resolved": "http://registry.npm.taobao.org/define-property/download/define-property-1.0.0.tgz",
          "integrity": "sha1-dp66rz9KY6rTr56NMEybvnm/sOY=",
          "dev": true,
          "requires": {
            "is-descriptor": "^1.0.0"
          }
        },
        "extend-shallow": {
          "version": "2.0.1",
          "resolved": "http://registry.npm.taobao.org/extend-shallow/download/extend-shallow-2.0.1.tgz",
          "integrity": "sha1-Ua99YUrZqfYQ6huvu5idaxxWiQ8=",
          "dev": true,
          "requires": {
            "is-extendable": "^0.1.0"
          }
        },
        "is-accessor-descriptor": {
          "version": "1.0.0",
          "resolved": "http://registry.npm.taobao.org/is-accessor-descriptor/download/is-accessor-descriptor-1.0.0.tgz",
          "integrity": "sha1-FpwvbT3x+ZJhgHI2XJsOofaHhlY=",
          "dev": true,
          "requires": {
            "kind-of": "^6.0.0"
          }
        },
        "is-data-descriptor": {
          "version": "1.0.0",
          "resolved": "http://registry.npm.taobao.org/is-data-descriptor/download/is-data-descriptor-1.0.0.tgz",
          "integrity": "sha1-2Eh2Mh0Oet0DmQQGq7u9NrqSaMc=",
          "dev": true,
          "requires": {
            "kind-of": "^6.0.0"
          }
        },
        "is-descriptor": {
          "version": "1.0.2",
          "resolved": "http://registry.npm.taobao.org/is-descriptor/download/is-descriptor-1.0.2.tgz",
          "integrity": "sha1-OxWXRqZmBLBPjIFSS6NlxfFNhuw=",
          "dev": true,
          "requires": {
            "is-accessor-descriptor": "^1.0.0",
            "is-data-descriptor": "^1.0.0",
            "kind-of": "^6.0.2"
          }
        }
      }
    },
    "extsprintf": {
      "version": "1.3.0",
      "resolved": "http://registry.npm.taobao.org/extsprintf/download/extsprintf-1.3.0.tgz",
      "integrity": "sha1-lpGEQOMEGnpBT4xS48V06zw+HgU=",
      "dev": true
    },
    "fast-deep-equal": {
      "version": "2.0.1",
      "resolved": "http://registry.npm.taobao.org/fast-deep-equal/download/fast-deep-equal-2.0.1.tgz",
      "integrity": "sha1-ewUhjd+WZ79/Nwv3/bLLFf3Qqkk=",
      "dev": true
    },
    "fast-glob": {
      "version": "2.2.4",
      "resolved": "http://registry.npm.taobao.org/fast-glob/download/fast-glob-2.2.4.tgz",
      "integrity": "sha1-5U9LZtN4BA4OTWpo7Da7xbBDY8A=",
      "dev": true,
      "requires": {
        "@mrmlnc/readdir-enhanced": "^2.2.1",
        "@nodelib/fs.stat": "^1.1.2",
        "glob-parent": "^3.1.0",
        "is-glob": "^4.0.0",
        "merge2": "^1.2.3",
        "micromatch": "^3.1.10"
      }
    },
    "fast-json-stable-stringify": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/fast-json-stable-stringify/download/fast-json-stable-stringify-2.0.0.tgz",
      "integrity": "sha1-1RQsDK7msRifh9OnYREGT4bIu/I=",
      "dev": true
    },
    "fast-levenshtein": {
      "version": "2.0.6",
      "resolved": "http://registry.npm.taobao.org/fast-levenshtein/download/fast-levenshtein-2.0.6.tgz",
      "integrity": "sha1-PYpcZog6FqMMqGQ+hR8Zuqd5eRc=",
      "dev": true
    },
    "fastparse": {
      "version": "1.1.2",
      "resolved": "http://registry.npm.taobao.org/fastparse/download/fastparse-1.1.2.tgz",
      "integrity": "sha1-kXKMWllC7O2FMSg8eUQe5BIsNak=",
      "dev": true
    },
    "faye-websocket": {
      "version": "0.10.0",
      "resolved": "http://registry.npm.taobao.org/faye-websocket/download/faye-websocket-0.10.0.tgz",
      "integrity": "sha1-TkkvjQTftviQA1B/btvy1QHnxvQ=",
      "dev": true,
      "requires": {
        "websocket-driver": ">=0.5.1"
      }
    },
    "figgy-pudding": {
      "version": "3.5.1",
      "resolved": "http://registry.npm.taobao.org/figgy-pudding/download/figgy-pudding-3.5.1.tgz",
      "integrity": "sha1-hiRwESkBxyeg5JWoB0S9W6odZ5A=",
      "dev": true
    },
    "figures": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/figures/download/figures-2.0.0.tgz",
      "integrity": "sha1-OrGi0qYsi/tDGgyUy3l6L84nyWI=",
      "dev": true,
      "requires": {
        "escape-string-regexp": "^1.0.5"
      }
    },
    "file-entry-cache": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/file-entry-cache/download/file-entry-cache-2.0.0.tgz",
      "integrity": "sha1-w5KZDD5oR4PYOLjISkXYoEhFg2E=",
      "dev": true,
      "requires": {
        "flat-cache": "^1.2.1",
        "object-assign": "^4.0.1"
      }
    },
    "file-loader": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/file-loader/download/file-loader-2.0.0.tgz",
      "integrity": "sha1-OXScgvAguehZAdz/mOgATmQBz94=",
      "dev": true,
      "requires": {
        "loader-utils": "^1.0.2",
        "schema-utils": "^1.0.0"
      },
      "dependencies": {
        "ajv-keywords": {
          "version": "3.2.0",
          "resolved": "http://registry.npm.taobao.org/ajv-keywords/download/ajv-keywords-3.2.0.tgz",
          "integrity": "sha1-6GuBnGAs+IIa1jdBNpjx3sAhhHo=",
          "dev": true
        },
        "schema-utils": {
          "version": "1.0.0",
          "resolved": "http://registry.npm.taobao.org/schema-utils/download/schema-utils-1.0.0.tgz",
          "integrity": "sha1-C3mpMgTXtgDUsoUNH2bCo0lRx3A=",
          "dev": true,
          "requires": {
            "ajv": "^6.1.0",
            "ajv-errors": "^1.0.0",
            "ajv-keywords": "^3.1.0"
          }
        }
      }
    },
    "filesize": {
      "version": "3.6.1",
      "resolved": "http://registry.npm.taobao.org/filesize/download/filesize-3.6.1.tgz",
      "integrity": "sha1-CQuz7gG2+AGoqL6Z0xcQs0Irsxc=",
      "dev": true
    },
    "fill-range": {
      "version": "4.0.0",
      "resolved": "http://registry.npm.taobao.org/fill-range/download/fill-range-4.0.0.tgz",
      "integrity": "sha1-1USBHUKPmOsGpj3EAtJAPDKMOPc=",
      "dev": true,
      "requires": {
        "extend-shallow": "^2.0.1",
        "is-number": "^3.0.0",
        "repeat-string": "^1.6.1",
        "to-regex-range": "^2.1.0"
      },
      "dependencies": {
        "extend-shallow": {
          "version": "2.0.1",
          "resolved": "http://registry.npm.taobao.org/extend-shallow/download/extend-shallow-2.0.1.tgz",
          "integrity": "sha1-Ua99YUrZqfYQ6huvu5idaxxWiQ8=",
          "dev": true,
          "requires": {
            "is-extendable": "^0.1.0"
          }
        }
      }
    },
    "finalhandler": {
      "version": "1.1.1",
      "resolved": "http://registry.npm.taobao.org/finalhandler/download/finalhandler-1.1.1.tgz",
      "integrity": "sha1-7r9O2EAHnIP0JJA4ydcDAIMBsQU=",
      "dev": true,
      "requires": {
        "debug": "2.6.9",
        "encodeurl": "~1.0.2",
        "escape-html": "~1.0.3",
        "on-finished": "~2.3.0",
        "parseurl": "~1.3.2",
        "statuses": "~1.4.0",
        "unpipe": "~1.0.0"
      },
      "dependencies": {
        "debug": {
          "version": "2.6.9",
          "resolved": "http://registry.npm.taobao.org/debug/download/debug-2.6.9.tgz",
          "integrity": "sha1-XRKFFd8TT/Mn6QpMk/Tgd6U2NB8=",
          "dev": true,
          "requires": {
            "ms": "2.0.0"
          }
        },
        "ms": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/ms/download/ms-2.0.0.tgz",
          "integrity": "sha1-VgiurfwAvmwpAd9fmGF4jeDVl8g=",
          "dev": true
        }
      }
    },
    "find-cache-dir": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/find-cache-dir/download/find-cache-dir-1.0.0.tgz",
      "integrity": "sha1-kojj6ePMN0hxfTnq3hfPcfww7m8=",
      "dev": true,
      "requires": {
        "commondir": "^1.0.1",
        "make-dir": "^1.0.0",
        "pkg-dir": "^2.0.0"
      }
    },
    "find-up": {
      "version": "2.1.0",
      "resolved": "http://registry.npm.taobao.org/find-up/download/find-up-2.1.0.tgz",
      "integrity": "sha1-RdG35QbHF93UgndaK3eSCjwMV6c=",
      "dev": true,
      "requires": {
        "locate-path": "^2.0.0"
      }
    },
    "flat-cache": {
      "version": "1.3.4",
      "resolved": "http://registry.npm.taobao.org/flat-cache/download/flat-cache-1.3.4.tgz",
      "integrity": "sha1-LC73dSXMKSkAff/6HdMUqpyd7m8=",
      "dev": true,
      "requires": {
        "circular-json": "^0.3.1",
        "graceful-fs": "^4.1.2",
        "rimraf": "~2.6.2",
        "write": "^0.2.1"
      }
    },
    "flush-write-stream": {
      "version": "1.0.3",
      "resolved": "http://registry.npm.taobao.org/flush-write-stream/download/flush-write-stream-1.0.3.tgz",
      "integrity": "sha1-xdWG7zivYJdlC0m8QbVfq7GfNb0=",
      "dev": true,
      "requires": {
        "inherits": "^2.0.1",
        "readable-stream": "^2.0.4"
      }
    },
    "follow-redirects": {
      "version": "1.5.10",
      "resolved": "http://registry.npm.taobao.org/follow-redirects/download/follow-redirects-1.5.10.tgz",
      "integrity": "sha1-e3qfmuov3/NnhqlP9kPtB/T/Xio=",
      "dev": true,
      "requires": {
        "debug": "=3.1.0"
      },
      "dependencies": {
        "debug": {
          "version": "3.1.0",
          "resolved": "http://registry.npm.taobao.org/debug/download/debug-3.1.0.tgz",
          "integrity": "sha1-W7WgZyYotkFJVmuhaBnmFRjGcmE=",
          "dev": true,
          "requires": {
            "ms": "2.0.0"
          }
        },
        "ms": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/ms/download/ms-2.0.0.tgz",
          "integrity": "sha1-VgiurfwAvmwpAd9fmGF4jeDVl8g=",
          "dev": true
        }
      }
    },
    "for-in": {
      "version": "1.0.2",
      "resolved": "http://registry.npm.taobao.org/for-in/download/for-in-1.0.2.tgz",
      "integrity": "sha1-gQaNKVqBQuwKxybG4iAMMPttXoA=",
      "dev": true
    },
    "forever-agent": {
      "version": "0.6.1",
      "resolved": "http://registry.npm.taobao.org/forever-agent/download/forever-agent-0.6.1.tgz",
      "integrity": "sha1-+8cfDEGt6zf5bFd60e1C2P2sypE=",
      "dev": true
    },
    "form-data": {
      "version": "2.3.3",
      "resolved": "http://registry.npm.taobao.org/form-data/download/form-data-2.3.3.tgz",
      "integrity": "sha1-3M5SwF9kTymManq5Nr1yTO/786Y=",
      "dev": true,
      "requires": {
        "asynckit": "^0.4.0",
        "combined-stream": "^1.0.6",
        "mime-types": "^2.1.12"
      }
    },
    "forwarded": {
      "version": "0.1.2",
      "resolved": "http://registry.npm.taobao.org/forwarded/download/forwarded-0.1.2.tgz",
      "integrity": "sha1-mMI9qxF1ZXuMBXPozszZGw/xjIQ=",
      "dev": true
    },
    "fragment-cache": {
      "version": "0.2.1",
      "resolved": "http://registry.npm.taobao.org/fragment-cache/download/fragment-cache-0.2.1.tgz",
      "integrity": "sha1-QpD60n8T6Jvn8zeZxrxaCr//DRk=",
      "dev": true,
      "requires": {
        "map-cache": "^0.2.2"
      }
    },
    "fresh": {
      "version": "0.5.2",
      "resolved": "http://registry.npm.taobao.org/fresh/download/fresh-0.5.2.tgz",
      "integrity": "sha1-PYyt2Q2XZWn6g1qx+OSyOhBWBac=",
      "dev": true
    },
    "friendly-errors-webpack-plugin": {
      "version": "1.7.0",
      "resolved": "http://registry.npm.taobao.org/friendly-errors-webpack-plugin/download/friendly-errors-webpack-plugin-1.7.0.tgz",
      "integrity": "sha1-78hsu4FiJFZYYaG+ep2E0Kr+oTY=",
      "dev": true,
      "requires": {
        "chalk": "^1.1.3",
        "error-stack-parser": "^2.0.0",
        "string-width": "^2.0.0"
      },
      "dependencies": {
        "ansi-regex": {
          "version": "2.1.1",
          "resolved": "http://registry.npm.taobao.org/ansi-regex/download/ansi-regex-2.1.1.tgz",
          "integrity": "sha1-w7M6te42DYbg5ijwRorn7yfWVN8=",
          "dev": true
        },
        "ansi-styles": {
          "version": "2.2.1",
          "resolved": "http://registry.npm.taobao.org/ansi-styles/download/ansi-styles-2.2.1.tgz",
          "integrity": "sha1-tDLdM1i2NM914eRmQ2gkBTPB3b4=",
          "dev": true
        },
        "chalk": {
          "version": "1.1.3",
          "resolved": "http://registry.npm.taobao.org/chalk/download/chalk-1.1.3.tgz",
          "integrity": "sha1-qBFcVeSnAv5NFQq9OHKCKn4J/Jg=",
          "dev": true,
          "requires": {
            "ansi-styles": "^2.2.1",
            "escape-string-regexp": "^1.0.2",
            "has-ansi": "^2.0.0",
            "strip-ansi": "^3.0.0",
            "supports-color": "^2.0.0"
          }
        },
        "strip-ansi": {
          "version": "3.0.1",
          "resolved": "http://registry.npm.taobao.org/strip-ansi/download/strip-ansi-3.0.1.tgz",
          "integrity": "sha1-ajhfuIU9lS1f8F0Oiq+UJ43GPc8=",
          "dev": true,
          "requires": {
            "ansi-regex": "^2.0.0"
          }
        },
        "supports-color": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/supports-color/download/supports-color-2.0.0.tgz",
          "integrity": "sha1-U10EXOa2Nj+kARcIRimZXp3zJMc=",
          "dev": true
        }
      }
    },
    "from2": {
      "version": "2.3.0",
      "resolved": "http://registry.npm.taobao.org/from2/download/from2-2.3.0.tgz",
      "integrity": "sha1-i/tVAr3kpNNs/e6gB/zKIdfjgq8=",
      "dev": true,
      "requires": {
        "inherits": "^2.0.1",
        "readable-stream": "^2.0.0"
      }
    },
    "fs-extra": {
      "version": "7.0.1",
      "resolved": "http://registry.npm.taobao.org/fs-extra/download/fs-extra-7.0.1.tgz",
      "integrity": "sha1-TxicRKoSO4lfcigE9V6iPq3DSOk=",
      "dev": true,
      "requires": {
        "graceful-fs": "^4.1.2",
        "jsonfile": "^4.0.0",
        "universalify": "^0.1.0"
      }
    },
    "fs-write-stream-atomic": {
      "version": "1.0.10",
      "resolved": "http://registry.npm.taobao.org/fs-write-stream-atomic/download/fs-write-stream-atomic-1.0.10.tgz",
      "integrity": "sha1-tH31NJPvkR33VzHnCp3tAYnbQMk=",
      "dev": true,
      "requires": {
        "graceful-fs": "^4.1.2",
        "iferr": "^0.1.5",
        "imurmurhash": "^0.1.4",
        "readable-stream": "1 || 2"
      }
    },
    "fs.realpath": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/fs.realpath/download/fs.realpath-1.0.0.tgz",
      "integrity": "sha1-FQStJSMVjKpA20onh8sBQRmU6k8=",
      "dev": true
    },
    "fsevents": {
      "version": "1.2.4",
      "resolved": "http://registry.npm.taobao.org/fsevents/download/fsevents-1.2.4.tgz",
      "integrity": "sha1-9B3LGvJYKvNpLaNvxVy9jhBBxCY=",
      "dev": true,
      "optional": true,
      "requires": {
        "nan": "^2.9.2",
        "node-pre-gyp": "^0.10.0"
      },
      "dependencies": {
        "abbrev": {
          "version": "1.1.1",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "ansi-regex": {
          "version": "2.1.1",
          "bundled": true,
          "dev": true
        },
        "aproba": {
          "version": "1.2.0",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "are-we-there-yet": {
          "version": "1.1.4",
          "bundled": true,
          "dev": true,
          "optional": true,
          "requires": {
            "delegates": "^1.0.0",
            "readable-stream": "^2.0.6"
          }
        },
        "balanced-match": {
          "version": "1.0.0",
          "bundled": true,
          "dev": true
        },
        "brace-expansion": {
          "version": "1.1.11",
          "bundled": true,
          "dev": true,
          "requires": {
            "balanced-match": "^1.0.0",
            "concat-map": "0.0.1"
          }
        },
        "chownr": {
          "version": "1.0.1",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "code-point-at": {
          "version": "1.1.0",
          "bundled": true,
          "dev": true
        },
        "concat-map": {
          "version": "0.0.1",
          "bundled": true,
          "dev": true
        },
        "console-control-strings": {
          "version": "1.1.0",
          "bundled": true,
          "dev": true
        },
        "core-util-is": {
          "version": "1.0.2",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "debug": {
          "version": "2.6.9",
          "bundled": true,
          "dev": true,
          "optional": true,
          "requires": {
            "ms": "2.0.0"
          }
        },
        "deep-extend": {
          "version": "0.5.1",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "delegates": {
          "version": "1.0.0",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "detect-libc": {
          "version": "1.0.3",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "fs-minipass": {
          "version": "1.2.5",
          "bundled": true,
          "dev": true,
          "optional": true,
          "requires": {
            "minipass": "^2.2.1"
          }
        },
        "fs.realpath": {
          "version": "1.0.0",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "gauge": {
          "version": "2.7.4",
          "bundled": true,
          "dev": true,
          "optional": true,
          "requires": {
            "aproba": "^1.0.3",
            "console-control-strings": "^1.0.0",
            "has-unicode": "^2.0.0",
            "object-assign": "^4.1.0",
            "signal-exit": "^3.0.0",
            "string-width": "^1.0.1",
            "strip-ansi": "^3.0.1",
            "wide-align": "^1.1.0"
          }
        },
        "glob": {
          "version": "7.1.2",
          "bundled": true,
          "dev": true,
          "optional": true,
          "requires": {
            "fs.realpath": "^1.0.0",
            "inflight": "^1.0.4",
            "inherits": "2",
            "minimatch": "^3.0.4",
            "once": "^1.3.0",
            "path-is-absolute": "^1.0.0"
          }
        },
        "has-unicode": {
          "version": "2.0.1",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "iconv-lite": {
          "version": "0.4.21",
          "bundled": true,
          "dev": true,
          "optional": true,
          "requires": {
            "safer-buffer": "^2.1.0"
          }
        },
        "ignore-walk": {
          "version": "3.0.1",
          "bundled": true,
          "dev": true,
          "optional": true,
          "requires": {
            "minimatch": "^3.0.4"
          }
        },
        "inflight": {
          "version": "1.0.6",
          "bundled": true,
          "dev": true,
          "optional": true,
          "requires": {
            "once": "^1.3.0",
            "wrappy": "1"
          }
        },
        "inherits": {
          "version": "2.0.3",
          "bundled": true,
          "dev": true
        },
        "ini": {
          "version": "1.3.5",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "is-fullwidth-code-point": {
          "version": "1.0.0",
          "bundled": true,
          "dev": true,
          "requires": {
            "number-is-nan": "^1.0.0"
          }
        },
        "isarray": {
          "version": "1.0.0",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "minimatch": {
          "version": "3.0.4",
          "bundled": true,
          "dev": true,
          "requires": {
            "brace-expansion": "^1.1.7"
          }
        },
        "minimist": {
          "version": "0.0.8",
          "bundled": true,
          "dev": true
        },
        "minipass": {
          "version": "2.2.4",
          "bundled": true,
          "dev": true,
          "requires": {
            "safe-buffer": "^5.1.1",
            "yallist": "^3.0.0"
          }
        },
        "minizlib": {
          "version": "1.1.0",
          "bundled": true,
          "dev": true,
          "optional": true,
          "requires": {
            "minipass": "^2.2.1"
          }
        },
        "mkdirp": {
          "version": "0.5.1",
          "bundled": true,
          "dev": true,
          "requires": {
            "minimist": "0.0.8"
          }
        },
        "ms": {
          "version": "2.0.0",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "needle": {
          "version": "2.2.0",
          "bundled": true,
          "dev": true,
          "optional": true,
          "requires": {
            "debug": "^2.1.2",
            "iconv-lite": "^0.4.4",
            "sax": "^1.2.4"
          }
        },
        "node-pre-gyp": {
          "version": "0.10.0",
          "bundled": true,
          "dev": true,
          "optional": true,
          "requires": {
            "detect-libc": "^1.0.2",
            "mkdirp": "^0.5.1",
            "needle": "^2.2.0",
            "nopt": "^4.0.1",
            "npm-packlist": "^1.1.6",
            "npmlog": "^4.0.2",
            "rc": "^1.1.7",
            "rimraf": "^2.6.1",
            "semver": "^5.3.0",
            "tar": "^4"
          }
        },
        "nopt": {
          "version": "4.0.1",
          "bundled": true,
          "dev": true,
          "optional": true,
          "requires": {
            "abbrev": "1",
            "osenv": "^0.1.4"
          }
        },
        "npm-bundled": {
          "version": "1.0.3",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "npm-packlist": {
          "version": "1.1.10",
          "bundled": true,
          "dev": true,
          "optional": true,
          "requires": {
            "ignore-walk": "^3.0.1",
            "npm-bundled": "^1.0.1"
          }
        },
        "npmlog": {
          "version": "4.1.2",
          "bundled": true,
          "dev": true,
          "optional": true,
          "requires": {
            "are-we-there-yet": "~1.1.2",
            "console-control-strings": "~1.1.0",
            "gauge": "~2.7.3",
            "set-blocking": "~2.0.0"
          }
        },
        "number-is-nan": {
          "version": "1.0.1",
          "bundled": true,
          "dev": true
        },
        "object-assign": {
          "version": "4.1.1",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "once": {
          "version": "1.4.0",
          "bundled": true,
          "dev": true,
          "requires": {
            "wrappy": "1"
          }
        },
        "os-homedir": {
          "version": "1.0.2",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "os-tmpdir": {
          "version": "1.0.2",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "osenv": {
          "version": "0.1.5",
          "bundled": true,
          "dev": true,
          "optional": true,
          "requires": {
            "os-homedir": "^1.0.0",
            "os-tmpdir": "^1.0.0"
          }
        },
        "path-is-absolute": {
          "version": "1.0.1",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "process-nextick-args": {
          "version": "2.0.0",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "rc": {
          "version": "1.2.7",
          "bundled": true,
          "dev": true,
          "optional": true,
          "requires": {
            "deep-extend": "^0.5.1",
            "ini": "~1.3.0",
            "minimist": "^1.2.0",
            "strip-json-comments": "~2.0.1"
          },
          "dependencies": {
            "minimist": {
              "version": "1.2.0",
              "bundled": true,
              "dev": true,
              "optional": true
            }
          }
        },
        "readable-stream": {
          "version": "2.3.6",
          "bundled": true,
          "dev": true,
          "optional": true,
          "requires": {
            "core-util-is": "~1.0.0",
            "inherits": "~2.0.3",
            "isarray": "~1.0.0",
            "process-nextick-args": "~2.0.0",
            "safe-buffer": "~5.1.1",
            "string_decoder": "~1.1.1",
            "util-deprecate": "~1.0.1"
          }
        },
        "rimraf": {
          "version": "2.6.2",
          "bundled": true,
          "dev": true,
          "optional": true,
          "requires": {
            "glob": "^7.0.5"
          }
        },
        "safe-buffer": {
          "version": "5.1.1",
          "bundled": true,
          "dev": true
        },
        "safer-buffer": {
          "version": "2.1.2",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "sax": {
          "version": "1.2.4",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "semver": {
          "version": "5.5.0",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "set-blocking": {
          "version": "2.0.0",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "signal-exit": {
          "version": "3.0.2",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "string-width": {
          "version": "1.0.2",
          "bundled": true,
          "dev": true,
          "requires": {
            "code-point-at": "^1.0.0",
            "is-fullwidth-code-point": "^1.0.0",
            "strip-ansi": "^3.0.0"
          }
        },
        "string_decoder": {
          "version": "1.1.1",
          "bundled": true,
          "dev": true,
          "optional": true,
          "requires": {
            "safe-buffer": "~5.1.0"
          }
        },
        "strip-ansi": {
          "version": "3.0.1",
          "bundled": true,
          "dev": true,
          "requires": {
            "ansi-regex": "^2.0.0"
          }
        },
        "strip-json-comments": {
          "version": "2.0.1",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "tar": {
          "version": "4.4.1",
          "bundled": true,
          "dev": true,
          "optional": true,
          "requires": {
            "chownr": "^1.0.1",
            "fs-minipass": "^1.2.5",
            "minipass": "^2.2.4",
            "minizlib": "^1.1.0",
            "mkdirp": "^0.5.0",
            "safe-buffer": "^5.1.1",
            "yallist": "^3.0.2"
          }
        },
        "util-deprecate": {
          "version": "1.0.2",
          "bundled": true,
          "dev": true,
          "optional": true
        },
        "wide-align": {
          "version": "1.1.2",
          "bundled": true,
          "dev": true,
          "optional": true,
          "requires": {
            "string-width": "^1.0.2"
          }
        },
        "wrappy": {
          "version": "1.0.2",
          "bundled": true,
          "dev": true
        },
        "yallist": {
          "version": "3.0.2",
          "bundled": true,
          "dev": true
        }
      }
    },
    "function-bind": {
      "version": "1.1.1",
      "resolved": "http://registry.npm.taobao.org/function-bind/download/function-bind-1.1.1.tgz",
      "integrity": "sha1-pWiZ0+o8m6uHS7l3O3xe3pL0iV0=",
      "dev": true
    },
    "functional-red-black-tree": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/functional-red-black-tree/download/functional-red-black-tree-1.0.1.tgz",
      "integrity": "sha1-GwqzvVU7Kg1jmdKcDj6gslIHgyc=",
      "dev": true
    },
    "get-caller-file": {
      "version": "1.0.3",
      "resolved": "http://registry.npm.taobao.org/get-caller-file/download/get-caller-file-1.0.3.tgz",
      "integrity": "sha1-+Xj6TJDR3+f/LWvtoqUV5xO9z0o=",
      "dev": true
    },
    "get-stream": {
      "version": "4.1.0",
      "resolved": "http://registry.npm.taobao.org/get-stream/download/get-stream-4.1.0.tgz",
      "integrity": "sha1-wbJVV189wh1Zv8ec09K0axw6VLU=",
      "dev": true,
      "requires": {
        "pump": "^3.0.0"
      }
    },
    "get-value": {
      "version": "2.0.6",
      "resolved": "http://registry.npm.taobao.org/get-value/download/get-value-2.0.6.tgz",
      "integrity": "sha1-3BXKHGcjh8p2vTesCjlbogQqLCg=",
      "dev": true
    },
    "getpass": {
      "version": "0.1.7",
      "resolved": "http://registry.npm.taobao.org/getpass/download/getpass-0.1.7.tgz",
      "integrity": "sha1-Xv+OPmhNVprkyysSgmBOi6YhSfo=",
      "dev": true,
      "requires": {
        "assert-plus": "^1.0.0"
      }
    },
    "glob": {
      "version": "7.1.3",
      "resolved": "http://registry.npm.taobao.org/glob/download/glob-7.1.3.tgz",
      "integrity": "sha1-OWCDLT8VdBCDQtr9OmezMsCWnfE=",
      "dev": true,
      "requires": {
        "fs.realpath": "^1.0.0",
        "inflight": "^1.0.4",
        "inherits": "2",
        "minimatch": "^3.0.4",
        "once": "^1.3.0",
        "path-is-absolute": "^1.0.0"
      }
    },
    "glob-parent": {
      "version": "3.1.0",
      "resolved": "http://registry.npm.taobao.org/glob-parent/download/glob-parent-3.1.0.tgz",
      "integrity": "sha1-nmr2KZ2NO9K9QEMIMr0RPfkGxa4=",
      "dev": true,
      "requires": {
        "is-glob": "^3.1.0",
        "path-dirname": "^1.0.0"
      },
      "dependencies": {
        "is-glob": {
          "version": "3.1.0",
          "resolved": "http://registry.npm.taobao.org/is-glob/download/is-glob-3.1.0.tgz",
          "integrity": "sha1-e6WuJCF4BKxwcHuWkiVnSGzD6Eo=",
          "dev": true,
          "requires": {
            "is-extglob": "^2.1.0"
          }
        }
      }
    },
    "glob-to-regexp": {
      "version": "0.3.0",
      "resolved": "http://registry.npm.taobao.org/glob-to-regexp/download/glob-to-regexp-0.3.0.tgz",
      "integrity": "sha1-jFoUlNIGbFcMw7/kSWF1rMTVAqs=",
      "dev": true
    },
    "globals": {
      "version": "11.9.0",
      "resolved": "http://registry.npm.taobao.org/globals/download/globals-11.9.0.tgz",
      "integrity": "sha1-veI2gI6YfykHaKk9BlBg145qskk=",
      "dev": true
    },
    "globby": {
      "version": "8.0.1",
      "resolved": "http://registry.npm.taobao.org/globby/download/globby-8.0.1.tgz",
      "integrity": "sha1-ta1IuKqAs1uBT8EoHsyFHx0rW1A=",
      "dev": true,
      "requires": {
        "array-union": "^1.0.1",
        "dir-glob": "^2.0.0",
        "fast-glob": "^2.0.2",
        "glob": "^7.1.2",
        "ignore": "^3.3.5",
        "pify": "^3.0.0",
        "slash": "^1.0.0"
      }
    },
    "graceful-fs": {
      "version": "4.1.15",
      "resolved": "http://registry.npm.taobao.org/graceful-fs/download/graceful-fs-4.1.15.tgz",
      "integrity": "sha1-/7cD4QZuig7qpMi4C6klPu77+wA=",
      "dev": true
    },
    "gzip-size": {
      "version": "5.0.0",
      "resolved": "http://registry.npm.taobao.org/gzip-size/download/gzip-size-5.0.0.tgz",
      "integrity": "sha1-pV7NmSIvTEj9jAHGJc47NJ0KDoA=",
      "dev": true,
      "requires": {
        "duplexer": "^0.1.1",
        "pify": "^3.0.0"
      }
    },
    "handle-thing": {
      "version": "1.2.5",
      "resolved": "http://registry.npm.taobao.org/handle-thing/download/handle-thing-1.2.5.tgz",
      "integrity": "sha1-/Xqtcmvxpf0W38KbL3pmAdJxOcQ=",
      "dev": true
    },
    "har-schema": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/har-schema/download/har-schema-2.0.0.tgz",
      "integrity": "sha1-qUwiJOvKwEeCoNkDVSHyRzW37JI=",
      "dev": true
    },
    "har-validator": {
      "version": "5.1.3",
      "resolved": "http://registry.npm.taobao.org/har-validator/download/har-validator-5.1.3.tgz",
      "integrity": "sha1-HvievT5JllV2de7ZiTEQ3DUPoIA=",
      "dev": true,
      "requires": {
        "ajv": "^6.5.5",
        "har-schema": "^2.0.0"
      }
    },
    "has": {
      "version": "1.0.3",
      "resolved": "http://registry.npm.taobao.org/has/download/has-1.0.3.tgz",
      "integrity": "sha1-ci18v8H2qoJB8W3YFOAR4fQeh5Y=",
      "dev": true,
      "requires": {
        "function-bind": "^1.1.1"
      }
    },
    "has-ansi": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/has-ansi/download/has-ansi-2.0.0.tgz",
      "integrity": "sha1-NPUEnOHs3ysGSa8+8k5F7TVBbZE=",
      "dev": true,
      "requires": {
        "ansi-regex": "^2.0.0"
      },
      "dependencies": {
        "ansi-regex": {
          "version": "2.1.1",
          "resolved": "http://registry.npm.taobao.org/ansi-regex/download/ansi-regex-2.1.1.tgz",
          "integrity": "sha1-w7M6te42DYbg5ijwRorn7yfWVN8=",
          "dev": true
        }
      }
    },
    "has-flag": {
      "version": "3.0.0",
      "resolved": "http://registry.npm.taobao.org/has-flag/download/has-flag-3.0.0.tgz",
      "integrity": "sha1-tdRU3CGZriJWmfNGfloH87lVuv0=",
      "dev": true
    },
    "has-symbols": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/has-symbols/download/has-symbols-1.0.0.tgz",
      "integrity": "sha1-uhqPGvKg/DllD1yFA2dwQSIGO0Q=",
      "dev": true
    },
    "has-value": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/has-value/download/has-value-1.0.0.tgz",
      "integrity": "sha1-GLKB2lhbHFxR3vJMkw7SmgvmsXc=",
      "dev": true,
      "requires": {
        "get-value": "^2.0.6",
        "has-values": "^1.0.0",
        "isobject": "^3.0.0"
      }
    },
    "has-values": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/has-values/download/has-values-1.0.0.tgz",
      "integrity": "sha1-lbC2P+whRmGab+V/51Yo1aOe/k8=",
      "dev": true,
      "requires": {
        "is-number": "^3.0.0",
        "kind-of": "^4.0.0"
      },
      "dependencies": {
        "kind-of": {
          "version": "4.0.0",
          "resolved": "http://registry.npm.taobao.org/kind-of/download/kind-of-4.0.0.tgz",
          "integrity": "sha1-IIE989cSkosgc3hpGkUGb65y3Vc=",
          "dev": true,
          "requires": {
            "is-buffer": "^1.1.5"
          }
        }
      }
    },
    "hash-base": {
      "version": "3.0.4",
      "resolved": "http://registry.npm.taobao.org/hash-base/download/hash-base-3.0.4.tgz",
      "integrity": "sha1-X8hoaEfs1zSZQDMZprCj8/auSRg=",
      "dev": true,
      "requires": {
        "inherits": "^2.0.1",
        "safe-buffer": "^5.0.1"
      }
    },
    "hash-sum": {
      "version": "1.0.2",
      "resolved": "http://registry.npm.taobao.org/hash-sum/download/hash-sum-1.0.2.tgz",
      "integrity": "sha1-M7QHd3VMZDJXPBIMw4CLvRDUfwQ=",
      "dev": true
    },
    "hash.js": {
      "version": "1.1.7",
      "resolved": "http://registry.npm.taobao.org/hash.js/download/hash.js-1.1.7.tgz",
      "integrity": "sha1-C6vKU46NTuSg+JiNaIZlN6ADz0I=",
      "dev": true,
      "requires": {
        "inherits": "^2.0.3",
        "minimalistic-assert": "^1.0.1"
      }
    },
    "he": {
      "version": "1.2.0",
      "resolved": "http://registry.npm.taobao.org/he/download/he-1.2.0.tgz",
      "integrity": "sha1-hK5l+n6vsWX922FWauFLrwVmTw8=",
      "dev": true
    },
    "hex-color-regex": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/hex-color-regex/download/hex-color-regex-1.1.0.tgz",
      "integrity": "sha1-TAb8y0YC/iYCs8k9+C1+fb8aio4=",
      "dev": true
    },
    "hmac-drbg": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/hmac-drbg/download/hmac-drbg-1.0.1.tgz",
      "integrity": "sha1-0nRXAQJabHdabFRXk+1QL8DGSaE=",
      "dev": true,
      "requires": {
        "hash.js": "^1.0.3",
        "minimalistic-assert": "^1.0.0",
        "minimalistic-crypto-utils": "^1.0.1"
      }
    },
    "hoek": {
      "version": "5.0.4",
      "resolved": "http://registry.npm.taobao.org/hoek/download/hoek-5.0.4.tgz",
      "integrity": "sha1-D3+icKHK/rNkpLLd+qM/hk5BV9o=",
      "dev": true
    },
    "hoopy": {
      "version": "0.1.4",
      "resolved": "http://registry.npm.taobao.org/hoopy/download/hoopy-0.1.4.tgz",
      "integrity": "sha1-YJIH1mEQADOpqUAq096mdzgcGx0=",
      "dev": true
    },
    "hosted-git-info": {
      "version": "2.7.1",
      "resolved": "http://registry.npm.taobao.org/hosted-git-info/download/hosted-git-info-2.7.1.tgz",
      "integrity": "sha1-l/I2l3vW4SVAiTD/bePuxigewEc=",
      "dev": true
    },
    "hpack.js": {
      "version": "2.1.6",
      "resolved": "http://registry.npm.taobao.org/hpack.js/download/hpack.js-2.1.6.tgz",
      "integrity": "sha1-h3dMCUnlE/QuhFdbPEVoH63ioLI=",
      "dev": true,
      "requires": {
        "inherits": "^2.0.1",
        "obuf": "^1.0.0",
        "readable-stream": "^2.0.1",
        "wbuf": "^1.1.0"
      }
    },
    "hsl-regex": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/hsl-regex/download/hsl-regex-1.0.0.tgz",
      "integrity": "sha1-1JMwx4ntgZ4nakwNJy3/owsY/m4=",
      "dev": true
    },
    "hsla-regex": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/hsla-regex/download/hsla-regex-1.0.0.tgz",
      "integrity": "sha1-wc56MWjIxmFAM6S194d/OyJfnDg=",
      "dev": true
    },
    "html-comment-regex": {
      "version": "1.1.2",
      "resolved": "http://registry.npm.taobao.org/html-comment-regex/download/html-comment-regex-1.1.2.tgz",
      "integrity": "sha1-l9RoiutcgYhqNk+qDK0d2hTUM6c=",
      "dev": true
    },
    "html-entities": {
      "version": "1.2.1",
      "resolved": "http://registry.npm.taobao.org/html-entities/download/html-entities-1.2.1.tgz",
      "integrity": "sha1-DfKTUfByEWNRXfueVUPl9u7VFi8=",
      "dev": true
    },
    "html-minifier": {
      "version": "3.5.21",
      "resolved": "http://registry.npm.taobao.org/html-minifier/download/html-minifier-3.5.21.tgz",
      "integrity": "sha1-0AQOBUcw41TbAIRjWTGUAVIS0gw=",
      "dev": true,
      "requires": {
        "camel-case": "3.0.x",
        "clean-css": "4.2.x",
        "commander": "2.17.x",
        "he": "1.2.x",
        "param-case": "2.1.x",
        "relateurl": "0.2.x",
        "uglify-js": "3.4.x"
      }
    },
    "html-webpack-plugin": {
      "version": "3.2.0",
      "resolved": "http://registry.npm.taobao.org/html-webpack-plugin/download/html-webpack-plugin-3.2.0.tgz",
      "integrity": "sha1-sBq71yOsqqeze2r0SS69oD2d03s=",
      "dev": true,
      "requires": {
        "html-minifier": "^3.2.3",
        "loader-utils": "^0.2.16",
        "lodash": "^4.17.3",
        "pretty-error": "^2.0.2",
        "tapable": "^1.0.0",
        "toposort": "^1.0.0",
        "util.promisify": "1.0.0"
      },
      "dependencies": {
        "json5": {
          "version": "0.5.1",
          "resolved": "http://registry.npm.taobao.org/json5/download/json5-0.5.1.tgz",
          "integrity": "sha1-Hq3nrMASA0rYTiOWdn6tn6VJWCE=",
          "dev": true
        },
        "loader-utils": {
          "version": "0.2.17",
          "resolved": "http://registry.npm.taobao.org/loader-utils/download/loader-utils-0.2.17.tgz",
          "integrity": "sha1-+G5jdNQyBabmxg6RlvF8Apm/s0g=",
          "dev": true,
          "requires": {
            "big.js": "^3.1.3",
            "emojis-list": "^2.0.0",
            "json5": "^0.5.0",
            "object-assign": "^4.0.1"
          }
        }
      }
    },
    "htmlparser2": {
      "version": "3.3.0",
      "resolved": "http://registry.npm.taobao.org/htmlparser2/download/htmlparser2-3.3.0.tgz",
      "integrity": "sha1-zHDQWln2VC5D8OaFyYLhTJJKnv4=",
      "dev": true,
      "requires": {
        "domelementtype": "1",
        "domhandler": "2.1",
        "domutils": "1.1",
        "readable-stream": "1.0"
      },
      "dependencies": {
        "domutils": {
          "version": "1.1.6",
          "resolved": "http://registry.npm.taobao.org/domutils/download/domutils-1.1.6.tgz",
          "integrity": "sha1-vdw94Jm5ou+sxRxiPyj0FuzFdIU=",
          "dev": true,
          "requires": {
            "domelementtype": "1"
          }
        },
        "isarray": {
          "version": "0.0.1",
          "resolved": "http://registry.npm.taobao.org/isarray/download/isarray-0.0.1.tgz",
          "integrity": "sha1-ihis/Kmo9Bd+Cav8YDiTmwXR7t8=",
          "dev": true
        },
        "readable-stream": {
          "version": "1.0.34",
          "resolved": "http://registry.npm.taobao.org/readable-stream/download/readable-stream-1.0.34.tgz",
          "integrity": "sha1-Elgg40vIQtLyqq+v5MKRbuMsFXw=",
          "dev": true,
          "requires": {
            "core-util-is": "~1.0.0",
            "inherits": "~2.0.1",
            "isarray": "0.0.1",
            "string_decoder": "~0.10.x"
          }
        },
        "string_decoder": {
          "version": "0.10.31",
          "resolved": "http://registry.npm.taobao.org/string_decoder/download/string_decoder-0.10.31.tgz",
          "integrity": "sha1-YuIDvEF2bGwoyfyEMB2rHFMQ+pQ=",
          "dev": true
        }
      }
    },
    "http-deceiver": {
      "version": "1.2.7",
      "resolved": "http://registry.npm.taobao.org/http-deceiver/download/http-deceiver-1.2.7.tgz",
      "integrity": "sha1-+nFolEq5pRnTN8sL7HKE3D5yPYc=",
      "dev": true
    },
    "http-errors": {
      "version": "1.6.3",
      "resolved": "http://registry.npm.taobao.org/http-errors/download/http-errors-1.6.3.tgz",
      "integrity": "sha1-i1VoC7S+KDoLW/TqLjhYC+HZMg0=",
      "dev": true,
      "requires": {
        "depd": "~1.1.2",
        "inherits": "2.0.3",
        "setprototypeof": "1.1.0",
        "statuses": ">= 1.4.0 < 2"
      }
    },
    "http-parser-js": {
      "version": "0.5.0",
      "resolved": "http://registry.npm.taobao.org/http-parser-js/download/http-parser-js-0.5.0.tgz",
      "integrity": "sha1-1l7b7ehDSdDcMDIIFaFdOcw8u9g=",
      "dev": true
    },
    "http-proxy": {
      "version": "1.17.0",
      "resolved": "http://registry.npm.taobao.org/http-proxy/download/http-proxy-1.17.0.tgz",
      "integrity": "sha1-etOElGWPhGBeL220Q230EPTlvpo=",
      "dev": true,
      "requires": {
        "eventemitter3": "^3.0.0",
        "follow-redirects": "^1.0.0",
        "requires-port": "^1.0.0"
      }
    },
    "http-proxy-middleware": {
      "version": "0.18.0",
      "resolved": "http://registry.npm.taobao.org/http-proxy-middleware/download/http-proxy-middleware-0.18.0.tgz",
      "integrity": "sha1-CYfmu1pWBuWmkWjY+WeofxXdiqs=",
      "dev": true,
      "requires": {
        "http-proxy": "^1.16.2",
        "is-glob": "^4.0.0",
        "lodash": "^4.17.5",
        "micromatch": "^3.1.9"
      }
    },
    "http-signature": {
      "version": "1.2.0",
      "resolved": "http://registry.npm.taobao.org/http-signature/download/http-signature-1.2.0.tgz",
      "integrity": "sha1-muzZJRFHcvPZW2WmCruPfBj7rOE=",
      "dev": true,
      "requires": {
        "assert-plus": "^1.0.0",
        "jsprim": "^1.2.2",
        "sshpk": "^1.7.0"
      }
    },
    "https-browserify": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/https-browserify/download/https-browserify-1.0.0.tgz",
      "integrity": "sha1-7AbBDgo0wPL68Zn3/X/Hj//QPHM=",
      "dev": true
    },
    "iconv-lite": {
      "version": "0.4.24",
      "resolved": "http://registry.npm.taobao.org/iconv-lite/download/iconv-lite-0.4.24.tgz",
      "integrity": "sha1-ICK0sl+93CHS9SSXSkdKr+czkIs=",
      "dev": true,
      "requires": {
        "safer-buffer": ">= 2.1.2 < 3"
      }
    },
    "icss-replace-symbols": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/icss-replace-symbols/download/icss-replace-symbols-1.1.0.tgz",
      "integrity": "sha1-Bupvg2ead0njhs/h/oEq5dsiPe0=",
      "dev": true
    },
    "icss-utils": {
      "version": "2.1.0",
      "resolved": "http://registry.npm.taobao.org/icss-utils/download/icss-utils-2.1.0.tgz",
      "integrity": "sha1-g/Cg7DeL8yRheLbCrZE28TWxyWI=",
      "dev": true,
      "requires": {
        "postcss": "^6.0.1"
      },
      "dependencies": {
        "postcss": {
          "version": "6.0.23",
          "resolved": "http://registry.npm.taobao.org/postcss/download/postcss-6.0.23.tgz",
          "integrity": "sha1-YcgswyisYOZ3ZF+XkFTrmLwOMyQ=",
          "dev": true,
          "requires": {
            "chalk": "^2.4.1",
            "source-map": "^0.6.1",
            "supports-color": "^5.4.0"
          }
        },
        "source-map": {
          "version": "0.6.1",
          "resolved": "http://registry.npm.taobao.org/source-map/download/source-map-0.6.1.tgz",
          "integrity": "sha1-dHIq8y6WFOnCh6jQu95IteLxomM=",
          "dev": true
        }
      }
    },
    "ieee754": {
      "version": "1.1.12",
      "resolved": "http://registry.npm.taobao.org/ieee754/download/ieee754-1.1.12.tgz",
      "integrity": "sha1-UL8k5bnIu5ivSWTJQc2wkY2ntgs=",
      "dev": true
    },
    "iferr": {
      "version": "0.1.5",
      "resolved": "http://registry.npm.taobao.org/iferr/download/iferr-0.1.5.tgz",
      "integrity": "sha1-xg7taebY/bazEEofy8ocGS3FtQE=",
      "dev": true
    },
    "ignore": {
      "version": "3.3.10",
      "resolved": "http://registry.npm.taobao.org/ignore/download/ignore-3.3.10.tgz",
      "integrity": "sha1-Cpf7h2mG6AgcYxFg+PnziRV/AEM=",
      "dev": true
    },
    "import-cwd": {
      "version": "2.1.0",
      "resolved": "http://registry.npm.taobao.org/import-cwd/download/import-cwd-2.1.0.tgz",
      "integrity": "sha1-qmzzbnInYShcs3HsZRn1PiQ1sKk=",
      "dev": true,
      "requires": {
        "import-from": "^2.1.0"
      }
    },
    "import-fresh": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/import-fresh/download/import-fresh-2.0.0.tgz",
      "integrity": "sha1-2BNVwVYS04bGH53dOSLUMEgipUY=",
      "dev": true,
      "requires": {
        "caller-path": "^2.0.0",
        "resolve-from": "^3.0.0"
      },
      "dependencies": {
        "caller-path": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/caller-path/download/caller-path-2.0.0.tgz",
          "integrity": "sha1-Ro+DBE42mrIBD6xfBs7uFbsssfQ=",
          "dev": true,
          "requires": {
            "caller-callsite": "^2.0.0"
          }
        },
        "resolve-from": {
          "version": "3.0.0",
          "resolved": "http://registry.npm.taobao.org/resolve-from/download/resolve-from-3.0.0.tgz",
          "integrity": "sha1-six699nWiBvItuZTM17rywoYh0g=",
          "dev": true
        }
      }
    },
    "import-from": {
      "version": "2.1.0",
      "resolved": "http://registry.npm.taobao.org/import-from/download/import-from-2.1.0.tgz",
      "integrity": "sha1-M1238qev/VOqpHHUuAId7ja387E=",
      "dev": true,
      "requires": {
        "resolve-from": "^3.0.0"
      },
      "dependencies": {
        "resolve-from": {
          "version": "3.0.0",
          "resolved": "http://registry.npm.taobao.org/resolve-from/download/resolve-from-3.0.0.tgz",
          "integrity": "sha1-six699nWiBvItuZTM17rywoYh0g=",
          "dev": true
        }
      }
    },
    "import-local": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/import-local/download/import-local-2.0.0.tgz",
      "integrity": "sha1-VQcL44pZk88Y72236WH1vuXFoJ0=",
      "dev": true,
      "requires": {
        "pkg-dir": "^3.0.0",
        "resolve-cwd": "^2.0.0"
      },
      "dependencies": {
        "find-up": {
          "version": "3.0.0",
          "resolved": "http://registry.npm.taobao.org/find-up/download/find-up-3.0.0.tgz",
          "integrity": "sha1-SRafHXmTQwZG2mHsxa41XCHJe3M=",
          "dev": true,
          "requires": {
            "locate-path": "^3.0.0"
          }
        },
        "locate-path": {
          "version": "3.0.0",
          "resolved": "http://registry.npm.taobao.org/locate-path/download/locate-path-3.0.0.tgz",
          "integrity": "sha1-2+w7OrdZdYBxtY/ln8QYca8hQA4=",
          "dev": true,
          "requires": {
            "p-locate": "^3.0.0",
            "path-exists": "^3.0.0"
          }
        },
        "p-limit": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/p-limit/download/p-limit-2.0.0.tgz",
          "integrity": "sha1-5iTtVO6MRgp3izyfNnBJb/ileuw=",
          "dev": true,
          "requires": {
            "p-try": "^2.0.0"
          }
        },
        "p-locate": {
          "version": "3.0.0",
          "resolved": "http://registry.npm.taobao.org/p-locate/download/p-locate-3.0.0.tgz",
          "integrity": "sha1-Mi1poFwCZLJZl9n0DNiokasAZKQ=",
          "dev": true,
          "requires": {
            "p-limit": "^2.0.0"
          }
        },
        "p-try": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/p-try/download/p-try-2.0.0.tgz",
          "integrity": "sha1-hQgLuHxkaI+keZb+j3376CEXYLE=",
          "dev": true
        },
        "pkg-dir": {
          "version": "3.0.0",
          "resolved": "http://registry.npm.taobao.org/pkg-dir/download/pkg-dir-3.0.0.tgz",
          "integrity": "sha1-J0kCDyOe2ZCIGx9xIQ1R62UjvqM=",
          "dev": true,
          "requires": {
            "find-up": "^3.0.0"
          }
        }
      }
    },
    "imurmurhash": {
      "version": "0.1.4",
      "resolved": "http://registry.npm.taobao.org/imurmurhash/download/imurmurhash-0.1.4.tgz",
      "integrity": "sha1-khi5srkoojixPcT7a21XbyMUU+o=",
      "dev": true
    },
    "indexes-of": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/indexes-of/download/indexes-of-1.0.1.tgz",
      "integrity": "sha1-8w9xbI4r00bHtn0985FVZqfAVgc=",
      "dev": true
    },
    "indexof": {
      "version": "0.0.1",
      "resolved": "http://registry.npm.taobao.org/indexof/download/indexof-0.0.1.tgz",
      "integrity": "sha1-gtwzbSMrkGIXnQWrMpOmYFn9Q10=",
      "dev": true
    },
    "inflight": {
      "version": "1.0.6",
      "resolved": "http://registry.npm.taobao.org/inflight/download/inflight-1.0.6.tgz",
      "integrity": "sha1-Sb1jMdfQLQwJvJEKEHW6gWW1bfk=",
      "dev": true,
      "requires": {
        "once": "^1.3.0",
        "wrappy": "1"
      }
    },
    "inherits": {
      "version": "2.0.3",
      "resolved": "http://registry.npm.taobao.org/inherits/download/inherits-2.0.3.tgz",
      "integrity": "sha1-Yzwsg+PaQqUC9SRmAiSA9CCCYd4=",
      "dev": true
    },
    "inquirer": {
      "version": "3.3.0",
      "resolved": "http://registry.npm.taobao.org/inquirer/download/inquirer-3.3.0.tgz",
      "integrity": "sha1-ndLyrXZdyrH/BEO0kUQqILoifck=",
      "dev": true,
      "requires": {
        "ansi-escapes": "^3.0.0",
        "chalk": "^2.0.0",
        "cli-cursor": "^2.1.0",
        "cli-width": "^2.0.0",
        "external-editor": "^2.0.4",
        "figures": "^2.0.0",
        "lodash": "^4.3.0",
        "mute-stream": "0.0.7",
        "run-async": "^2.2.0",
        "rx-lite": "^4.0.8",
        "rx-lite-aggregates": "^4.0.8",
        "string-width": "^2.1.0",
        "strip-ansi": "^4.0.0",
        "through": "^2.3.6"
      }
    },
    "internal-ip": {
      "version": "3.0.1",
      "resolved": "http://registry.npm.taobao.org/internal-ip/download/internal-ip-3.0.1.tgz",
      "integrity": "sha1-31yZh24dLrLqLXT1IOP2aaAOzic=",
      "dev": true,
      "requires": {
        "default-gateway": "^2.6.0",
        "ipaddr.js": "^1.5.2"
      }
    },
    "invariant": {
      "version": "2.2.4",
      "resolved": "http://registry.npm.taobao.org/invariant/download/invariant-2.2.4.tgz",
      "integrity": "sha1-YQ88ksk1nOHbYW5TgAjSP/NRWOY=",
      "dev": true,
      "requires": {
        "loose-envify": "^1.0.0"
      }
    },
    "invert-kv": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/invert-kv/download/invert-kv-2.0.0.tgz",
      "integrity": "sha1-c5P1r6Weyf9fZ6J2INEcIm4+7AI=",
      "dev": true
    },
    "ip": {
      "version": "1.1.5",
      "resolved": "http://registry.npm.taobao.org/ip/download/ip-1.1.5.tgz",
      "integrity": "sha1-vd7XARQpCCjAoDnnLvJfWq7ENUo=",
      "dev": true
    },
    "ip-regex": {
      "version": "2.1.0",
      "resolved": "http://registry.npm.taobao.org/ip-regex/download/ip-regex-2.1.0.tgz",
      "integrity": "sha1-+ni/XS5pE8kRzp+BnuUUa7bYROk=",
      "dev": true
    },
    "ipaddr.js": {
      "version": "1.8.0",
      "resolved": "http://registry.npm.taobao.org/ipaddr.js/download/ipaddr.js-1.8.0.tgz",
      "integrity": "sha1-6qM9bd16zo9/b+DJygRA5wZzix4=",
      "dev": true
    },
    "is-absolute-url": {
      "version": "2.1.0",
      "resolved": "http://registry.npm.taobao.org/is-absolute-url/download/is-absolute-url-2.1.0.tgz",
      "integrity": "sha1-UFMN+4T8yap9vnhS6Do3uTufKqY=",
      "dev": true
    },
    "is-accessor-descriptor": {
      "version": "0.1.6",
      "resolved": "http://registry.npm.taobao.org/is-accessor-descriptor/download/is-accessor-descriptor-0.1.6.tgz",
      "integrity": "sha1-qeEss66Nh2cn7u84Q/igiXtcmNY=",
      "dev": true,
      "requires": {
        "kind-of": "^3.0.2"
      },
      "dependencies": {
        "kind-of": {
          "version": "3.2.2",
          "resolved": "http://registry.npm.taobao.org/kind-of/download/kind-of-3.2.2.tgz",
          "integrity": "sha1-MeohpzS6ubuw8yRm2JOupR5KPGQ=",
          "dev": true,
          "requires": {
            "is-buffer": "^1.1.5"
          }
        }
      }
    },
    "is-arrayish": {
      "version": "0.2.1",
      "resolved": "http://registry.npm.taobao.org/is-arrayish/download/is-arrayish-0.2.1.tgz",
      "integrity": "sha1-d8mYQFJ6qOyxqLppe4BkWnqSap0=",
      "dev": true
    },
    "is-binary-path": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/is-binary-path/download/is-binary-path-1.0.1.tgz",
      "integrity": "sha1-dfFmQrSA8YenEcgUFh/TpKdlWJg=",
      "dev": true,
      "requires": {
        "binary-extensions": "^1.0.0"
      }
    },
    "is-buffer": {
      "version": "1.1.6",
      "resolved": "http://registry.npm.taobao.org/is-buffer/download/is-buffer-1.1.6.tgz",
      "integrity": "sha1-76ouqdqg16suoTqXsritUf776L4=",
      "dev": true
    },
    "is-builtin-module": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/is-builtin-module/download/is-builtin-module-1.0.0.tgz",
      "integrity": "sha1-VAVy0096wxGfj3bDDLwbHgN6/74=",
      "dev": true,
      "requires": {
        "builtin-modules": "^1.0.0"
      }
    },
    "is-callable": {
      "version": "1.1.4",
      "resolved": "http://registry.npm.taobao.org/is-callable/download/is-callable-1.1.4.tgz",
      "integrity": "sha1-HhrfIZ4e62hNaR+dagX/DTCiTXU=",
      "dev": true
    },
    "is-ci": {
      "version": "1.2.1",
      "resolved": "http://registry.npm.taobao.org/is-ci/download/is-ci-1.2.1.tgz",
      "integrity": "sha1-43ecjuF/zPQoSI9uKBGH8uYyhBw=",
      "dev": true,
      "requires": {
        "ci-info": "^1.5.0"
      }
    },
    "is-color-stop": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/is-color-stop/download/is-color-stop-1.1.0.tgz",
      "integrity": "sha1-z/9HGu5N1cnhWFmPvhKWe1za00U=",
      "dev": true,
      "requires": {
        "css-color-names": "^0.0.4",
        "hex-color-regex": "^1.1.0",
        "hsl-regex": "^1.0.0",
        "hsla-regex": "^1.0.0",
        "rgb-regex": "^1.0.1",
        "rgba-regex": "^1.0.0"
      }
    },
    "is-data-descriptor": {
      "version": "0.1.4",
      "resolved": "http://registry.npm.taobao.org/is-data-descriptor/download/is-data-descriptor-0.1.4.tgz",
      "integrity": "sha1-C17mSDiOLIYCgueT8YVv7D8wG1Y=",
      "dev": true,
      "requires": {
        "kind-of": "^3.0.2"
      },
      "dependencies": {
        "kind-of": {
          "version": "3.2.2",
          "resolved": "http://registry.npm.taobao.org/kind-of/download/kind-of-3.2.2.tgz",
          "integrity": "sha1-MeohpzS6ubuw8yRm2JOupR5KPGQ=",
          "dev": true,
          "requires": {
            "is-buffer": "^1.1.5"
          }
        }
      }
    },
    "is-date-object": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/is-date-object/download/is-date-object-1.0.1.tgz",
      "integrity": "sha1-mqIOtq7rv/d/vTPnTKAbM1gdOhY=",
      "dev": true
    },
    "is-descriptor": {
      "version": "0.1.6",
      "resolved": "http://registry.npm.taobao.org/is-descriptor/download/is-descriptor-0.1.6.tgz",
      "integrity": "sha1-Nm2CQN3kh8pRgjsaufB6EKeCUco=",
      "dev": true,
      "requires": {
        "is-accessor-descriptor": "^0.1.6",
        "is-data-descriptor": "^0.1.4",
        "kind-of": "^5.0.0"
      },
      "dependencies": {
        "kind-of": {
          "version": "5.1.0",
          "resolved": "http://registry.npm.taobao.org/kind-of/download/kind-of-5.1.0.tgz",
          "integrity": "sha1-cpyR4thXt6QZofmqZWhcTDP1hF0=",
          "dev": true
        }
      }
    },
    "is-directory": {
      "version": "0.3.1",
      "resolved": "http://registry.npm.taobao.org/is-directory/download/is-directory-0.3.1.tgz",
      "integrity": "sha1-YTObbyR1/Hcv2cnYP1yFddwVSuE=",
      "dev": true
    },
    "is-extendable": {
      "version": "0.1.1",
      "resolved": "http://registry.npm.taobao.org/is-extendable/download/is-extendable-0.1.1.tgz",
      "integrity": "sha1-YrEQ4omkcUGOPsNqYX1HLjAd/Ik=",
      "dev": true
    },
    "is-extglob": {
      "version": "2.1.1",
      "resolved": "http://registry.npm.taobao.org/is-extglob/download/is-extglob-2.1.1.tgz",
      "integrity": "sha1-qIwCU1eR8C7TfHahueqXc8gz+MI=",
      "dev": true
    },
    "is-fullwidth-code-point": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/is-fullwidth-code-point/download/is-fullwidth-code-point-2.0.0.tgz",
      "integrity": "sha1-o7MKXE8ZkYMWeqq5O+764937ZU8=",
      "dev": true
    },
    "is-glob": {
      "version": "4.0.0",
      "resolved": "http://registry.npm.taobao.org/is-glob/download/is-glob-4.0.0.tgz",
      "integrity": "sha1-lSHHaEXMJhCoUgPd8ICpWML/q8A=",
      "dev": true,
      "requires": {
        "is-extglob": "^2.1.1"
      }
    },
    "is-number": {
      "version": "3.0.0",
      "resolved": "http://registry.npm.taobao.org/is-number/download/is-number-3.0.0.tgz",
      "integrity": "sha1-JP1iAaR4LPUFYcgQJ2r8fRLXEZU=",
      "dev": true,
      "requires": {
        "kind-of": "^3.0.2"
      },
      "dependencies": {
        "kind-of": {
          "version": "3.2.2",
          "resolved": "http://registry.npm.taobao.org/kind-of/download/kind-of-3.2.2.tgz",
          "integrity": "sha1-MeohpzS6ubuw8yRm2JOupR5KPGQ=",
          "dev": true,
          "requires": {
            "is-buffer": "^1.1.5"
          }
        }
      }
    },
    "is-obj": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/is-obj/download/is-obj-1.0.1.tgz",
      "integrity": "sha1-PkcprB9f3gJc19g6iW2rn09n2w8=",
      "dev": true
    },
    "is-path-cwd": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/is-path-cwd/download/is-path-cwd-1.0.0.tgz",
      "integrity": "sha1-0iXsIxMuie3Tj9p2dHLmLmXxEG0=",
      "dev": true
    },
    "is-path-in-cwd": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/is-path-in-cwd/download/is-path-in-cwd-1.0.1.tgz",
      "integrity": "sha1-WsSLNF72dTOb1sekipEhELJBz1I=",
      "dev": true,
      "requires": {
        "is-path-inside": "^1.0.0"
      }
    },
    "is-path-inside": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/is-path-inside/download/is-path-inside-1.0.1.tgz",
      "integrity": "sha1-jvW33lBDej/cprToZe96pVy0gDY=",
      "dev": true,
      "requires": {
        "path-is-inside": "^1.0.1"
      }
    },
    "is-plain-object": {
      "version": "2.0.4",
      "resolved": "http://registry.npm.taobao.org/is-plain-object/download/is-plain-object-2.0.4.tgz",
      "integrity": "sha1-LBY7P6+xtgbZ0Xko8FwqHDjgdnc=",
      "dev": true,
      "requires": {
        "isobject": "^3.0.1"
      }
    },
    "is-promise": {
      "version": "2.1.0",
      "resolved": "http://registry.npm.taobao.org/is-promise/download/is-promise-2.1.0.tgz",
      "integrity": "sha1-eaKp7OfwlugPNtKy87wWwf9L8/o=",
      "dev": true
    },
    "is-regex": {
      "version": "1.0.4",
      "resolved": "http://registry.npm.taobao.org/is-regex/download/is-regex-1.0.4.tgz",
      "integrity": "sha1-VRdIm1RwkbCTDglWVM7SXul+lJE=",
      "dev": true,
      "requires": {
        "has": "^1.0.1"
      }
    },
    "is-resolvable": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/is-resolvable/download/is-resolvable-1.1.0.tgz",
      "integrity": "sha1-+xj4fOH+uSUWnJpAfBkxijIG7Yg=",
      "dev": true
    },
    "is-stream": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/is-stream/download/is-stream-1.1.0.tgz",
      "integrity": "sha1-EtSj3U5o4Lec6428hBc66A2RykQ=",
      "dev": true
    },
    "is-svg": {
      "version": "3.0.0",
      "resolved": "http://registry.npm.taobao.org/is-svg/download/is-svg-3.0.0.tgz",
      "integrity": "sha1-kyHb0pwhLlypnE+peUxxS8r6L3U=",
      "dev": true,
      "requires": {
        "html-comment-regex": "^1.1.0"
      }
    },
    "is-symbol": {
      "version": "1.0.2",
      "resolved": "http://registry.npm.taobao.org/is-symbol/download/is-symbol-1.0.2.tgz",
      "integrity": "sha1-oFX2rlcZLK7jKeeoYBGLSXqVDzg=",
      "dev": true,
      "requires": {
        "has-symbols": "^1.0.0"
      }
    },
    "is-typedarray": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/is-typedarray/download/is-typedarray-1.0.0.tgz",
      "integrity": "sha1-5HnICFjfDBsR3dppQPlgEfzaSpo=",
      "dev": true
    },
    "is-windows": {
      "version": "1.0.2",
      "resolved": "http://registry.npm.taobao.org/is-windows/download/is-windows-1.0.2.tgz",
      "integrity": "sha1-0YUOuXkezRjmGCzhKjDzlmNLsZ0=",
      "dev": true
    },
    "is-wsl": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/is-wsl/download/is-wsl-1.1.0.tgz",
      "integrity": "sha1-HxbkqiKwTRM2tmGIpmrzxgDDpm0=",
      "dev": true
    },
    "isarray": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/isarray/download/isarray-1.0.0.tgz",
      "integrity": "sha1-u5NdSFgsuhaMBoNJV6VKPgcSTxE=",
      "dev": true
    },
    "isemail": {
      "version": "3.2.0",
      "resolved": "http://registry.npm.taobao.org/isemail/download/isemail-3.2.0.tgz",
      "integrity": "sha1-WTEKAhkxqfsGu7UeFVzgs/I2gyw=",
      "dev": true,
      "requires": {
        "punycode": "2.x.x"
      }
    },
    "isexe": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/isexe/download/isexe-2.0.0.tgz",
      "integrity": "sha1-6PvzdNxVb/iUehDcsFctYz8s+hA=",
      "dev": true
    },
    "isobject": {
      "version": "3.0.1",
      "resolved": "http://registry.npm.taobao.org/isobject/download/isobject-3.0.1.tgz",
      "integrity": "sha1-TkMekrEalzFjaqH5yNHMvP2reN8=",
      "dev": true
    },
    "isstream": {
      "version": "0.1.2",
      "resolved": "http://registry.npm.taobao.org/isstream/download/isstream-0.1.2.tgz",
      "integrity": "sha1-R+Y/evVa+m+S4VAOaQ64uFKcCZo=",
      "dev": true
    },
    "javascript-stringify": {
      "version": "1.6.0",
      "resolved": "http://registry.npm.taobao.org/javascript-stringify/download/javascript-stringify-1.6.0.tgz",
      "integrity": "sha1-FC0RHzpuPa6PSpr9d9RYVbWpzOM=",
      "dev": true
    },
    "joi": {
      "version": "13.7.0",
      "resolved": "http://registry.npm.taobao.org/joi/download/joi-13.7.0.tgz",
      "integrity": "sha1-z9hev+Z+ihkAQyQAtNA7vZP7h58=",
      "dev": true,
      "requires": {
        "hoek": "5.x.x",
        "isemail": "3.x.x",
        "topo": "3.x.x"
      }
    },
    "js-levenshtein": {
      "version": "1.1.4",
      "resolved": "http://registry.npm.taobao.org/js-levenshtein/download/js-levenshtein-1.1.4.tgz",
      "integrity": "sha1-Olbjy/WJygCB6yLNm6CxKQoW0m4=",
      "dev": true
    },
    "js-message": {
      "version": "1.0.5",
      "resolved": "http://registry.npm.taobao.org/js-message/download/js-message-1.0.5.tgz",
      "integrity": "sha1-IwDSSxrwjondCVvBpMnJz8uJLRU=",
      "dev": true
    },
    "js-queue": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/js-queue/download/js-queue-2.0.0.tgz",
      "integrity": "sha1-NiITz4YPRo8BJfxslqvBdCUx+Ug=",
      "dev": true,
      "requires": {
        "easy-stack": "^1.0.0"
      }
    },
    "js-tokens": {
      "version": "4.0.0",
      "resolved": "http://registry.npm.taobao.org/js-tokens/download/js-tokens-4.0.0.tgz",
      "integrity": "sha1-GSA/tZmR35jjoocFDUZHzerzJJk=",
      "dev": true
    },
    "js-yaml": {
      "version": "3.12.0",
      "resolved": "http://registry.npm.taobao.org/js-yaml/download/js-yaml-3.12.0.tgz",
      "integrity": "sha1-6u1lbsg0TxD1J8a/obbiJE3hZ9E=",
      "dev": true,
      "requires": {
        "argparse": "^1.0.7",
        "esprima": "^4.0.0"
      }
    },
    "jsbn": {
      "version": "0.1.1",
      "resolved": "http://registry.npm.taobao.org/jsbn/download/jsbn-0.1.1.tgz",
      "integrity": "sha1-peZUwuWi3rXyAdls77yoDA7y9RM=",
      "dev": true
    },
    "jsesc": {
      "version": "2.5.2",
      "resolved": "http://registry.npm.taobao.org/jsesc/download/jsesc-2.5.2.tgz",
      "integrity": "sha1-gFZNLkg9rPbo7yCWUKZ98/DCg6Q=",
      "dev": true
    },
    "json-parse-better-errors": {
      "version": "1.0.2",
      "resolved": "http://registry.npm.taobao.org/json-parse-better-errors/download/json-parse-better-errors-1.0.2.tgz",
      "integrity": "sha1-u4Z8+zRQ5pEHwTHRxRS6s9yLyqk=",
      "dev": true
    },
    "json-schema": {
      "version": "0.2.3",
      "resolved": "http://registry.npm.taobao.org/json-schema/download/json-schema-0.2.3.tgz",
      "integrity": "sha1-tIDIkuWaLwWVTOcnvT8qTogvnhM=",
      "dev": true
    },
    "json-schema-traverse": {
      "version": "0.4.1",
      "resolved": "http://registry.npm.taobao.org/json-schema-traverse/download/json-schema-traverse-0.4.1.tgz",
      "integrity": "sha1-afaofZUTq4u4/mO9sJecRI5oRmA=",
      "dev": true
    },
    "json-stable-stringify-without-jsonify": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/json-stable-stringify-without-jsonify/download/json-stable-stringify-without-jsonify-1.0.1.tgz",
      "integrity": "sha1-nbe1lJatPzz+8wp1FC0tkwrXJlE=",
      "dev": true
    },
    "json-stringify-safe": {
      "version": "5.0.1",
      "resolved": "http://registry.npm.taobao.org/json-stringify-safe/download/json-stringify-safe-5.0.1.tgz",
      "integrity": "sha1-Epai1Y/UXxmg9s4B1lcB4sc1tus=",
      "dev": true
    },
    "json3": {
      "version": "3.3.2",
      "resolved": "http://registry.npm.taobao.org/json3/download/json3-3.3.2.tgz",
      "integrity": "sha1-PAQ0dD35Pi9cQq7nsZvLSDV19OE=",
      "dev": true
    },
    "json5": {
      "version": "2.1.0",
      "resolved": "http://registry.npm.taobao.org/json5/download/json5-2.1.0.tgz",
      "integrity": "sha1-56DGLEgoXGKNIKELhcibuAfDKFA=",
      "dev": true,
      "requires": {
        "minimist": "^1.2.0"
      }
    },
    "jsonfile": {
      "version": "4.0.0",
      "resolved": "http://registry.npm.taobao.org/jsonfile/download/jsonfile-4.0.0.tgz",
      "integrity": "sha1-h3Gq4HmbZAdrdmQPygWPnBDjPss=",
      "dev": true,
      "requires": {
        "graceful-fs": "^4.1.6"
      }
    },
    "jsonify": {
      "version": "0.0.0",
      "resolved": "http://registry.npm.taobao.org/jsonify/download/jsonify-0.0.0.tgz",
      "integrity": "sha1-LHS27kHZPKUbe1qu6PUDYx0lKnM=",
      "dev": true
    },
    "jsprim": {
      "version": "1.4.1",
      "resolved": "http://registry.npm.taobao.org/jsprim/download/jsprim-1.4.1.tgz",
      "integrity": "sha1-MT5mvB5cwG5Di8G3SZwuXFastqI=",
      "dev": true,
      "requires": {
        "assert-plus": "1.0.0",
        "extsprintf": "1.3.0",
        "json-schema": "0.2.3",
        "verror": "1.10.0"
      }
    },
    "killable": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/killable/download/killable-1.0.1.tgz",
      "integrity": "sha1-TIzkQRh6Bhx0dPuHygjipjgZSJI=",
      "dev": true
    },
    "kind-of": {
      "version": "6.0.2",
      "resolved": "http://registry.npm.taobao.org/kind-of/download/kind-of-6.0.2.tgz",
      "integrity": "sha1-ARRrNqYhjmTljzqNZt5df8b20FE=",
      "dev": true
    },
    "launch-editor": {
      "version": "2.2.1",
      "resolved": "http://registry.npm.taobao.org/launch-editor/download/launch-editor-2.2.1.tgz",
      "integrity": "sha1-hxtaPuOdZoD8wm03kwtu7aidsMo=",
      "dev": true,
      "requires": {
        "chalk": "^2.3.0",
        "shell-quote": "^1.6.1"
      }
    },
    "launch-editor-middleware": {
      "version": "2.2.1",
      "resolved": "http://registry.npm.taobao.org/launch-editor-middleware/download/launch-editor-middleware-2.2.1.tgz",
      "integrity": "sha1-4UsH5scVSwpLhqD9NFeE5FgEwVc=",
      "dev": true,
      "requires": {
        "launch-editor": "^2.2.1"
      }
    },
    "lcid": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/lcid/download/lcid-2.0.0.tgz",
      "integrity": "sha1-bvXS32DlL4LrIopMNz6NHzlyU88=",
      "dev": true,
      "requires": {
        "invert-kv": "^2.0.0"
      }
    },
    "levn": {
      "version": "0.3.0",
      "resolved": "http://registry.npm.taobao.org/levn/download/levn-0.3.0.tgz",
      "integrity": "sha1-OwmSTt+fCDwEkP3UwLxEIeBHZO4=",
      "dev": true,
      "requires": {
        "prelude-ls": "~1.1.2",
        "type-check": "~0.3.2"
      }
    },
    "loader-fs-cache": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/loader-fs-cache/download/loader-fs-cache-1.0.1.tgz",
      "integrity": "sha1-VuC/CL2XCLJqdltoUJhAyN7J/bw=",
      "dev": true,
      "requires": {
        "find-cache-dir": "^0.1.1",
        "mkdirp": "0.5.1"
      },
      "dependencies": {
        "find-cache-dir": {
          "version": "0.1.1",
          "resolved": "http://registry.npm.taobao.org/find-cache-dir/download/find-cache-dir-0.1.1.tgz",
          "integrity": "sha1-yN765XyKUqinhPnjHFfHQumToLk=",
          "dev": true,
          "requires": {
            "commondir": "^1.0.1",
            "mkdirp": "^0.5.1",
            "pkg-dir": "^1.0.0"
          }
        },
        "find-up": {
          "version": "1.1.2",
          "resolved": "http://registry.npm.taobao.org/find-up/download/find-up-1.1.2.tgz",
          "integrity": "sha1-ay6YIrGizgpgq2TWEOzK1TyyTQ8=",
          "dev": true,
          "requires": {
            "path-exists": "^2.0.0",
            "pinkie-promise": "^2.0.0"
          }
        },
        "path-exists": {
          "version": "2.1.0",
          "resolved": "http://registry.npm.taobao.org/path-exists/download/path-exists-2.1.0.tgz",
          "integrity": "sha1-D+tsZPD8UY2adU3V77YscCJ2H0s=",
          "dev": true,
          "requires": {
            "pinkie-promise": "^2.0.0"
          }
        },
        "pkg-dir": {
          "version": "1.0.0",
          "resolved": "http://registry.npm.taobao.org/pkg-dir/download/pkg-dir-1.0.0.tgz",
          "integrity": "sha1-ektQio1bstYp1EcFb/TpyTFM89Q=",
          "dev": true,
          "requires": {
            "find-up": "^1.0.0"
          }
        }
      }
    },
    "loader-runner": {
      "version": "2.3.1",
      "resolved": "http://registry.npm.taobao.org/loader-runner/download/loader-runner-2.3.1.tgz",
      "integrity": "sha1-Am8S/nwxFZkolqwCugIrqSlxuXk=",
      "dev": true
    },
    "loader-utils": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/loader-utils/download/loader-utils-1.1.0.tgz",
      "integrity": "sha1-yYrvSIvM7aL/teLeZG1qdUQp9c0=",
      "dev": true,
      "requires": {
        "big.js": "^3.1.3",
        "emojis-list": "^2.0.0",
        "json5": "^0.5.0"
      },
      "dependencies": {
        "json5": {
          "version": "0.5.1",
          "resolved": "http://registry.npm.taobao.org/json5/download/json5-0.5.1.tgz",
          "integrity": "sha1-Hq3nrMASA0rYTiOWdn6tn6VJWCE=",
          "dev": true
        }
      }
    },
    "locate-path": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/locate-path/download/locate-path-2.0.0.tgz",
      "integrity": "sha1-K1aLJl7slExtnA3pw9u7ygNUzY4=",
      "dev": true,
      "requires": {
        "p-locate": "^2.0.0",
        "path-exists": "^3.0.0"
      }
    },
    "lodash": {
      "version": "4.17.11",
      "resolved": "http://registry.npm.taobao.org/lodash/download/lodash-4.17.11.tgz",
      "integrity": "sha1-s56mIp72B+zYniyN8SU2iRysm40=",
      "dev": true
    },
    "lodash.debounce": {
      "version": "4.0.8",
      "resolved": "http://registry.npm.taobao.org/lodash.debounce/download/lodash.debounce-4.0.8.tgz",
      "integrity": "sha1-gteb/zCmfEAF/9XiUVMArZyk168=",
      "dev": true
    },
    "lodash.defaultsdeep": {
      "version": "4.6.0",
      "resolved": "http://registry.npm.taobao.org/lodash.defaultsdeep/download/lodash.defaultsdeep-4.6.0.tgz",
      "integrity": "sha1-vsECT4WxvZbL6kBbI8FK1kQ6b4E=",
      "dev": true
    },
    "lodash.mapvalues": {
      "version": "4.6.0",
      "resolved": "http://registry.npm.taobao.org/lodash.mapvalues/download/lodash.mapvalues-4.6.0.tgz",
      "integrity": "sha1-G6+lAF3p3W9PJmaMMMo3IwzJaJw=",
      "dev": true
    },
    "lodash.memoize": {
      "version": "4.1.2",
      "resolved": "http://registry.npm.taobao.org/lodash.memoize/download/lodash.memoize-4.1.2.tgz",
      "integrity": "sha1-vMbEmkKihA7Zl/Mj6tpezRguC/4=",
      "dev": true
    },
    "lodash.transform": {
      "version": "4.6.0",
      "resolved": "http://registry.npm.taobao.org/lodash.transform/download/lodash.transform-4.6.0.tgz",
      "integrity": "sha1-EjBkIvYzJK7YSD0/ODMrX2cFR6A=",
      "dev": true
    },
    "lodash.uniq": {
      "version": "4.5.0",
      "resolved": "http://registry.npm.taobao.org/lodash.uniq/download/lodash.uniq-4.5.0.tgz",
      "integrity": "sha1-0CJTc662Uq3BvILklFM5qEJ1R3M=",
      "dev": true
    },
    "log-symbols": {
      "version": "2.2.0",
      "resolved": "http://registry.npm.taobao.org/log-symbols/download/log-symbols-2.2.0.tgz",
      "integrity": "sha1-V0Dhxdbw39pK2TI7UzIQfva0xAo=",
      "dev": true,
      "requires": {
        "chalk": "^2.0.1"
      }
    },
    "loglevel": {
      "version": "1.6.1",
      "resolved": "http://registry.npm.taobao.org/loglevel/download/loglevel-1.6.1.tgz",
      "integrity": "sha1-4PyVEztu8nbNyIh82vJKpvFW+Po=",
      "dev": true
    },
    "loose-envify": {
      "version": "1.4.0",
      "resolved": "http://registry.npm.taobao.org/loose-envify/download/loose-envify-1.4.0.tgz",
      "integrity": "sha1-ce5R+nvkyuwaY4OffmgtgTLTDK8=",
      "dev": true,
      "requires": {
        "js-tokens": "^3.0.0 || ^4.0.0"
      }
    },
    "lower-case": {
      "version": "1.1.4",
      "resolved": "http://registry.npm.taobao.org/lower-case/download/lower-case-1.1.4.tgz",
      "integrity": "sha1-miyr0bno4K6ZOkv31YdcOcQujqw=",
      "dev": true
    },
    "lru-cache": {
      "version": "4.1.5",
      "resolved": "http://registry.npm.taobao.org/lru-cache/download/lru-cache-4.1.5.tgz",
      "integrity": "sha1-i75Q6oW+1ZvJ4z3KuCNe6bz0Q80=",
      "dev": true,
      "requires": {
        "pseudomap": "^1.0.2",
        "yallist": "^2.1.2"
      }
    },
    "make-dir": {
      "version": "1.3.0",
      "resolved": "http://registry.npm.taobao.org/make-dir/download/make-dir-1.3.0.tgz",
      "integrity": "sha1-ecEDO4BRW9bSTsmTPoYMp17ifww=",
      "dev": true,
      "requires": {
        "pify": "^3.0.0"
      }
    },
    "map-age-cleaner": {
      "version": "0.1.3",
      "resolved": "http://registry.npm.taobao.org/map-age-cleaner/download/map-age-cleaner-0.1.3.tgz",
      "integrity": "sha1-fVg6cwZDTAVf5HSw9FB45uG0uSo=",
      "dev": true,
      "requires": {
        "p-defer": "^1.0.0"
      }
    },
    "map-cache": {
      "version": "0.2.2",
      "resolved": "http://registry.npm.taobao.org/map-cache/download/map-cache-0.2.2.tgz",
      "integrity": "sha1-wyq9C9ZSXZsFFkW7TyasXcmKDb8=",
      "dev": true
    },
    "map-visit": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/map-visit/download/map-visit-1.0.0.tgz",
      "integrity": "sha1-7Nyo8TFE5mDxtb1B8S80edmN+48=",
      "dev": true,
      "requires": {
        "object-visit": "^1.0.0"
      }
    },
    "md5.js": {
      "version": "1.3.5",
      "resolved": "http://registry.npm.taobao.org/md5.js/download/md5.js-1.3.5.tgz",
      "integrity": "sha1-tdB7jjIW4+J81yjXL3DR5qNCAF8=",
      "dev": true,
      "requires": {
        "hash-base": "^3.0.0",
        "inherits": "^2.0.1",
        "safe-buffer": "^5.1.2"
      }
    },
    "mdn-data": {
      "version": "1.1.4",
      "resolved": "http://registry.npm.taobao.org/mdn-data/download/mdn-data-1.1.4.tgz",
      "integrity": "sha1-ULXU/8RXUnZXPE7tuHgIEqhBnwE=",
      "dev": true
    },
    "media-typer": {
      "version": "0.3.0",
      "resolved": "http://registry.npm.taobao.org/media-typer/download/media-typer-0.3.0.tgz",
      "integrity": "sha1-hxDXrwqmJvj/+hzgAWhUUmMlV0g=",
      "dev": true
    },
    "mem": {
      "version": "4.0.0",
      "resolved": "http://registry.npm.taobao.org/mem/download/mem-4.0.0.tgz",
      "integrity": "sha1-ZDdpDZRxZ49syDZZwAy6/Nawza8=",
      "dev": true,
      "requires": {
        "map-age-cleaner": "^0.1.1",
        "mimic-fn": "^1.0.0",
        "p-is-promise": "^1.1.0"
      }
    },
    "memory-fs": {
      "version": "0.4.1",
      "resolved": "http://registry.npm.taobao.org/memory-fs/download/memory-fs-0.4.1.tgz",
      "integrity": "sha1-OpoguEYlI+RHz7x+i7gO1me/xVI=",
      "dev": true,
      "requires": {
        "errno": "^0.1.3",
        "readable-stream": "^2.0.1"
      }
    },
    "merge-descriptors": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/merge-descriptors/download/merge-descriptors-1.0.1.tgz",
      "integrity": "sha1-sAqqVW3YtEVoFQ7J0blT8/kMu2E=",
      "dev": true
    },
    "merge-source-map": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/merge-source-map/download/merge-source-map-1.1.0.tgz",
      "integrity": "sha1-L93n5gIJOfcJBqaPLXrmheTIxkY=",
      "dev": true,
      "requires": {
        "source-map": "^0.6.1"
      },
      "dependencies": {
        "source-map": {
          "version": "0.6.1",
          "resolved": "http://registry.npm.taobao.org/source-map/download/source-map-0.6.1.tgz",
          "integrity": "sha1-dHIq8y6WFOnCh6jQu95IteLxomM=",
          "dev": true
        }
      }
    },
    "merge2": {
      "version": "1.2.3",
      "resolved": "http://registry.npm.taobao.org/merge2/download/merge2-1.2.3.tgz",
      "integrity": "sha1-fumdvWm7ZIFoklPwGEiKG5ArDtU=",
      "dev": true
    },
    "methods": {
      "version": "1.1.2",
      "resolved": "http://registry.npm.taobao.org/methods/download/methods-1.1.2.tgz",
      "integrity": "sha1-VSmk1nZUE07cxSZmVoNbD4Ua/O4=",
      "dev": true
    },
    "micromatch": {
      "version": "3.1.10",
      "resolved": "http://registry.npm.taobao.org/micromatch/download/micromatch-3.1.10.tgz",
      "integrity": "sha1-cIWbyVyYQJUvNZoGij/En57PrCM=",
      "dev": true,
      "requires": {
        "arr-diff": "^4.0.0",
        "array-unique": "^0.3.2",
        "braces": "^2.3.1",
        "define-property": "^2.0.2",
        "extend-shallow": "^3.0.2",
        "extglob": "^2.0.4",
        "fragment-cache": "^0.2.1",
        "kind-of": "^6.0.2",
        "nanomatch": "^1.2.9",
        "object.pick": "^1.3.0",
        "regex-not": "^1.0.0",
        "snapdragon": "^0.8.1",
        "to-regex": "^3.0.2"
      }
    },
    "miller-rabin": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/miller-rabin/download/miller-rabin-4.0.1.tgz",
      "integrity": "sha1-8IA1HIZbDcViqEYpZtqlNUPHik0=",
      "dev": true,
      "requires": {
        "bn.js": "^4.0.0",
        "brorand": "^1.0.1"
      }
    },
    "mime": {
      "version": "2.4.0",
      "resolved": "http://registry.npm.taobao.org/mime/download/mime-2.4.0.tgz",
      "integrity": "sha1-4FH9iBNYWF8yed8zP+aU2gvP/dY=",
      "dev": true
    },
    "mime-db": {
      "version": "1.37.0",
      "resolved": "http://registry.npm.taobao.org/mime-db/download/mime-db-1.37.0.tgz",
      "integrity": "sha1-C2oM5v2+lXbiXx8tL96IMNwK0Ng=",
      "dev": true
    },
    "mime-types": {
      "version": "2.1.21",
      "resolved": "http://registry.npm.taobao.org/mime-types/download/mime-types-2.1.21.tgz",
      "integrity": "sha1-KJlaoey3cHQv5q5+WPkYHHRLP5Y=",
      "dev": true,
      "requires": {
        "mime-db": "~1.37.0"
      }
    },
    "mimic-fn": {
      "version": "1.2.0",
      "resolved": "http://registry.npm.taobao.org/mimic-fn/download/mimic-fn-1.2.0.tgz",
      "integrity": "sha1-ggyGo5M0ZA6ZUWkovQP8qIBX0CI=",
      "dev": true
    },
    "mini-css-extract-plugin": {
      "version": "0.4.5",
      "resolved": "http://registry.npm.taobao.org/mini-css-extract-plugin/download/mini-css-extract-plugin-0.4.5.tgz",
      "integrity": "sha1-yZ6eeNVPP6d1YzruWTOuqk6AcZo=",
      "dev": true,
      "requires": {
        "loader-utils": "^1.1.0",
        "schema-utils": "^1.0.0",
        "webpack-sources": "^1.1.0"
      },
      "dependencies": {
        "ajv-keywords": {
          "version": "3.2.0",
          "resolved": "http://registry.npm.taobao.org/ajv-keywords/download/ajv-keywords-3.2.0.tgz",
          "integrity": "sha1-6GuBnGAs+IIa1jdBNpjx3sAhhHo=",
          "dev": true
        },
        "schema-utils": {
          "version": "1.0.0",
          "resolved": "http://registry.npm.taobao.org/schema-utils/download/schema-utils-1.0.0.tgz",
          "integrity": "sha1-C3mpMgTXtgDUsoUNH2bCo0lRx3A=",
          "dev": true,
          "requires": {
            "ajv": "^6.1.0",
            "ajv-errors": "^1.0.0",
            "ajv-keywords": "^3.1.0"
          }
        }
      }
    },
    "minimalistic-assert": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/minimalistic-assert/download/minimalistic-assert-1.0.1.tgz",
      "integrity": "sha1-LhlN4ERibUoQ5/f7wAznPoPk1cc=",
      "dev": true
    },
    "minimalistic-crypto-utils": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/minimalistic-crypto-utils/download/minimalistic-crypto-utils-1.0.1.tgz",
      "integrity": "sha1-9sAMHAsIIkblxNmd+4x8CDsrWCo=",
      "dev": true
    },
    "minimatch": {
      "version": "3.0.4",
      "resolved": "http://registry.npm.taobao.org/minimatch/download/minimatch-3.0.4.tgz",
      "integrity": "sha1-UWbihkV/AzBgZL5Ul+jbsMPTIIM=",
      "dev": true,
      "requires": {
        "brace-expansion": "^1.1.7"
      }
    },
    "minimist": {
      "version": "1.2.0",
      "resolved": "http://registry.npm.taobao.org/minimist/download/minimist-1.2.0.tgz",
      "integrity": "sha1-o1AIsg9BOD7sH7kU9M1d95omQoQ=",
      "dev": true
    },
    "mississippi": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/mississippi/download/mississippi-2.0.0.tgz",
      "integrity": "sha1-NEKlCPr8KFAEhv7qmUCWduTuWm8=",
      "dev": true,
      "requires": {
        "concat-stream": "^1.5.0",
        "duplexify": "^3.4.2",
        "end-of-stream": "^1.1.0",
        "flush-write-stream": "^1.0.0",
        "from2": "^2.1.0",
        "parallel-transform": "^1.1.0",
        "pump": "^2.0.1",
        "pumpify": "^1.3.3",
        "stream-each": "^1.1.0",
        "through2": "^2.0.0"
      },
      "dependencies": {
        "pump": {
          "version": "2.0.1",
          "resolved": "http://registry.npm.taobao.org/pump/download/pump-2.0.1.tgz",
          "integrity": "sha1-Ejma3W5M91Jtlzy8i1zi4pCLOQk=",
          "dev": true,
          "requires": {
            "end-of-stream": "^1.1.0",
            "once": "^1.3.1"
          }
        }
      }
    },
    "mixin-deep": {
      "version": "1.3.1",
      "resolved": "http://registry.npm.taobao.org/mixin-deep/download/mixin-deep-1.3.1.tgz",
      "integrity": "sha1-pJ5yaNzhoNlpjkUybFYm3zVD0P4=",
      "dev": true,
      "requires": {
        "for-in": "^1.0.2",
        "is-extendable": "^1.0.1"
      },
      "dependencies": {
        "is-extendable": {
          "version": "1.0.1",
          "resolved": "http://registry.npm.taobao.org/is-extendable/download/is-extendable-1.0.1.tgz",
          "integrity": "sha1-p0cPnkJnM9gb2B4RVSZOOjUHyrQ=",
          "dev": true,
          "requires": {
            "is-plain-object": "^2.0.4"
          }
        }
      }
    },
    "mkdirp": {
      "version": "0.5.1",
      "resolved": "http://registry.npm.taobao.org/mkdirp/download/mkdirp-0.5.1.tgz",
      "integrity": "sha1-MAV0OOrGz3+MR2fzhkjWaX11yQM=",
      "dev": true,
      "requires": {
        "minimist": "0.0.8"
      },
      "dependencies": {
        "minimist": {
          "version": "0.0.8",
          "resolved": "http://registry.npm.taobao.org/minimist/download/minimist-0.0.8.tgz",
          "integrity": "sha1-hX/Kv8M5fSYluCKCYuhqp6ARsF0=",
          "dev": true
        }
      }
    },
    "move-concurrently": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/move-concurrently/download/move-concurrently-1.0.1.tgz",
      "integrity": "sha1-viwAX9oy4LKa8fBdfEszIUxwH5I=",
      "dev": true,
      "requires": {
        "aproba": "^1.1.1",
        "copy-concurrently": "^1.0.0",
        "fs-write-stream-atomic": "^1.0.8",
        "mkdirp": "^0.5.1",
        "rimraf": "^2.5.4",
        "run-queue": "^1.0.3"
      }
    },
    "ms": {
      "version": "2.1.1",
      "resolved": "http://registry.npm.taobao.org/ms/download/ms-2.1.1.tgz",
      "integrity": "sha1-MKWGTrPrsKZvLr5tcnrwagnYbgo=",
      "dev": true
    },
    "multicast-dns": {
      "version": "6.2.3",
      "resolved": "http://registry.npm.taobao.org/multicast-dns/download/multicast-dns-6.2.3.tgz",
      "integrity": "sha1-oOx72QVcQoL3kMPIL04o2zsxsik=",
      "dev": true,
      "requires": {
        "dns-packet": "^1.3.1",
        "thunky": "^1.0.2"
      }
    },
    "multicast-dns-service-types": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/multicast-dns-service-types/download/multicast-dns-service-types-1.1.0.tgz",
      "integrity": "sha1-iZ8R2WhuXgXLkbNdXw5jt3PPyQE=",
      "dev": true
    },
    "mute-stream": {
      "version": "0.0.7",
      "resolved": "http://registry.npm.taobao.org/mute-stream/download/mute-stream-0.0.7.tgz",
      "integrity": "sha1-MHXOk7whuPq0PhvE2n6BFe0ee6s=",
      "dev": true
    },
    "nan": {
      "version": "2.11.1",
      "resolved": "http://registry.npm.taobao.org/nan/download/nan-2.11.1.tgz",
      "integrity": "sha1-kOIrzLjKV+pM03zIPTgZtS7qZ2Y=",
      "dev": true,
      "optional": true
    },
    "nanomatch": {
      "version": "1.2.13",
      "resolved": "http://registry.npm.taobao.org/nanomatch/download/nanomatch-1.2.13.tgz",
      "integrity": "sha1-uHqKpPwN6P5r6IiVs4mD/yZb0Rk=",
      "dev": true,
      "requires": {
        "arr-diff": "^4.0.0",
        "array-unique": "^0.3.2",
        "define-property": "^2.0.2",
        "extend-shallow": "^3.0.2",
        "fragment-cache": "^0.2.1",
        "is-windows": "^1.0.2",
        "kind-of": "^6.0.2",
        "object.pick": "^1.3.0",
        "regex-not": "^1.0.0",
        "snapdragon": "^0.8.1",
        "to-regex": "^3.0.1"
      }
    },
    "natural-compare": {
      "version": "1.4.0",
      "resolved": "http://registry.npm.taobao.org/natural-compare/download/natural-compare-1.4.0.tgz",
      "integrity": "sha1-Sr6/7tdUHywnrPspvbvRXI1bpPc=",
      "dev": true
    },
    "negotiator": {
      "version": "0.6.1",
      "resolved": "http://registry.npm.taobao.org/negotiator/download/negotiator-0.6.1.tgz",
      "integrity": "sha1-KzJxhOiZIQEXeyhWP7XnECrNDKk=",
      "dev": true
    },
    "neo-async": {
      "version": "2.6.0",
      "resolved": "http://registry.npm.taobao.org/neo-async/download/neo-async-2.6.0.tgz",
      "integrity": "sha1-udFeTXHGdikIZUtRg+04t1M0CDU=",
      "dev": true
    },
    "nice-try": {
      "version": "1.0.5",
      "resolved": "http://registry.npm.taobao.org/nice-try/download/nice-try-1.0.5.tgz",
      "integrity": "sha1-ozeKdpbOfSI+iPybdkvX7xCJ42Y=",
      "dev": true
    },
    "no-case": {
      "version": "2.3.2",
      "resolved": "http://registry.npm.taobao.org/no-case/download/no-case-2.3.2.tgz",
      "integrity": "sha1-YLgTOWvjmz8SiKTB7V0efSi0ZKw=",
      "dev": true,
      "requires": {
        "lower-case": "^1.1.1"
      }
    },
    "node-forge": {
      "version": "0.7.5",
      "resolved": "http://registry.npm.taobao.org/node-forge/download/node-forge-0.7.5.tgz",
      "integrity": "sha1-bBUsNFzhHFL0ZcKr2VfoY5zWdN8=",
      "dev": true
    },
    "node-ipc": {
      "version": "9.1.1",
      "resolved": "http://registry.npm.taobao.org/node-ipc/download/node-ipc-9.1.1.tgz",
      "integrity": "sha1-TiRe1pOOZRAOWV68XcNLFujdXWk=",
      "dev": true,
      "requires": {
        "event-pubsub": "4.3.0",
        "js-message": "1.0.5",
        "js-queue": "2.0.0"
      }
    },
    "node-libs-browser": {
      "version": "2.1.0",
      "resolved": "http://registry.npm.taobao.org/node-libs-browser/download/node-libs-browser-2.1.0.tgz",
      "integrity": "sha1-X5QmPUBPbkR2fXJpAf/wVHjWAN8=",
      "dev": true,
      "requires": {
        "assert": "^1.1.1",
        "browserify-zlib": "^0.2.0",
        "buffer": "^4.3.0",
        "console-browserify": "^1.1.0",
        "constants-browserify": "^1.0.0",
        "crypto-browserify": "^3.11.0",
        "domain-browser": "^1.1.1",
        "events": "^1.0.0",
        "https-browserify": "^1.0.0",
        "os-browserify": "^0.3.0",
        "path-browserify": "0.0.0",
        "process": "^0.11.10",
        "punycode": "^1.2.4",
        "querystring-es3": "^0.2.0",
        "readable-stream": "^2.3.3",
        "stream-browserify": "^2.0.1",
        "stream-http": "^2.7.2",
        "string_decoder": "^1.0.0",
        "timers-browserify": "^2.0.4",
        "tty-browserify": "0.0.0",
        "url": "^0.11.0",
        "util": "^0.10.3",
        "vm-browserify": "0.0.4"
      },
      "dependencies": {
        "punycode": {
          "version": "1.4.1",
          "resolved": "http://registry.npm.taobao.org/punycode/download/punycode-1.4.1.tgz",
          "integrity": "sha1-wNWmOycYgArY4esPpSachN1BhF4=",
          "dev": true
        }
      }
    },
    "node-releases": {
      "version": "1.1.1",
      "resolved": "http://registry.npm.taobao.org/node-releases/download/node-releases-1.1.1.tgz",
      "integrity": "sha1-j/+K6hz8rR+0IF+AUUkFT79zyv0=",
      "dev": true,
      "requires": {
        "semver": "^5.3.0"
      }
    },
    "normalize-package-data": {
      "version": "2.4.0",
      "resolved": "http://registry.npm.taobao.org/normalize-package-data/download/normalize-package-data-2.4.0.tgz",
      "integrity": "sha1-EvlaMH1YNSB1oEkHuErIvpisAS8=",
      "dev": true,
      "requires": {
        "hosted-git-info": "^2.1.4",
        "is-builtin-module": "^1.0.0",
        "semver": "2 || 3 || 4 || 5",
        "validate-npm-package-license": "^3.0.1"
      }
    },
    "normalize-path": {
      "version": "2.1.1",
      "resolved": "http://registry.npm.taobao.org/normalize-path/download/normalize-path-2.1.1.tgz",
      "integrity": "sha1-GrKLVW4Zg2Oowab35vogE3/mrtk=",
      "dev": true,
      "requires": {
        "remove-trailing-separator": "^1.0.1"
      }
    },
    "normalize-range": {
      "version": "0.1.2",
      "resolved": "http://registry.npm.taobao.org/normalize-range/download/normalize-range-0.1.2.tgz",
      "integrity": "sha1-LRDAa9/TEuqXd2laTShDlFa3WUI=",
      "dev": true
    },
    "normalize-url": {
      "version": "3.3.0",
      "resolved": "http://registry.npm.taobao.org/normalize-url/download/normalize-url-3.3.0.tgz",
      "integrity": "sha1-suHE3E98bVd0PfczpPWXjRhlBVk=",
      "dev": true
    },
    "npm-run-path": {
      "version": "2.0.2",
      "resolved": "http://registry.npm.taobao.org/npm-run-path/download/npm-run-path-2.0.2.tgz",
      "integrity": "sha1-NakjLfo11wZ7TLLd8jV7GHFTbF8=",
      "dev": true,
      "requires": {
        "path-key": "^2.0.0"
      }
    },
    "nth-check": {
      "version": "1.0.2",
      "resolved": "http://registry.npm.taobao.org/nth-check/download/nth-check-1.0.2.tgz",
      "integrity": "sha1-sr0pXDfj3VijvwcAN2Zjuk2c8Fw=",
      "dev": true,
      "requires": {
        "boolbase": "~1.0.0"
      }
    },
    "num2fraction": {
      "version": "1.2.2",
      "resolved": "http://registry.npm.taobao.org/num2fraction/download/num2fraction-1.2.2.tgz",
      "integrity": "sha1-b2gragJ6Tp3fpFZM0lidHU5mnt4=",
      "dev": true
    },
    "number-is-nan": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/number-is-nan/download/number-is-nan-1.0.1.tgz",
      "integrity": "sha1-CXtgK1NCKlIsGvuHkDGDNpQaAR0=",
      "dev": true
    },
    "oauth-sign": {
      "version": "0.9.0",
      "resolved": "http://registry.npm.taobao.org/oauth-sign/download/oauth-sign-0.9.0.tgz",
      "integrity": "sha1-R6ewFrqmi1+g7PPe4IqFxnmsZFU=",
      "dev": true
    },
    "object-assign": {
      "version": "4.1.1",
      "resolved": "http://registry.npm.taobao.org/object-assign/download/object-assign-4.1.1.tgz",
      "integrity": "sha1-IQmtx5ZYh8/AXLvUQsrIv7s2CGM=",
      "dev": true
    },
    "object-copy": {
      "version": "0.1.0",
      "resolved": "http://registry.npm.taobao.org/object-copy/download/object-copy-0.1.0.tgz",
      "integrity": "sha1-fn2Fi3gb18mRpBupde04EnVOmYw=",
      "dev": true,
      "requires": {
        "copy-descriptor": "^0.1.0",
        "define-property": "^0.2.5",
        "kind-of": "^3.0.3"
      },
      "dependencies": {
        "define-property": {
          "version": "0.2.5",
          "resolved": "http://registry.npm.taobao.org/define-property/download/define-property-0.2.5.tgz",
          "integrity": "sha1-w1se+RjsPJkPmlvFe+BKrOxcgRY=",
          "dev": true,
          "requires": {
            "is-descriptor": "^0.1.0"
          }
        },
        "kind-of": {
          "version": "3.2.2",
          "resolved": "http://registry.npm.taobao.org/kind-of/download/kind-of-3.2.2.tgz",
          "integrity": "sha1-MeohpzS6ubuw8yRm2JOupR5KPGQ=",
          "dev": true,
          "requires": {
            "is-buffer": "^1.1.5"
          }
        }
      }
    },
    "object-hash": {
      "version": "1.3.1",
      "resolved": "http://registry.npm.taobao.org/object-hash/download/object-hash-1.3.1.tgz",
      "integrity": "sha1-/eRSCYqVHLFF8Dm7fUVUSd3BJt8=",
      "dev": true
    },
    "object-keys": {
      "version": "1.0.12",
      "resolved": "http://registry.npm.taobao.org/object-keys/download/object-keys-1.0.12.tgz",
      "integrity": "sha1-CcU4VTd1dTEMymL1W7M0q/97PtI=",
      "dev": true
    },
    "object-visit": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/object-visit/download/object-visit-1.0.1.tgz",
      "integrity": "sha1-95xEk68MU3e1n+OdOV5BBC3QRbs=",
      "dev": true,
      "requires": {
        "isobject": "^3.0.0"
      }
    },
    "object.assign": {
      "version": "4.1.0",
      "resolved": "http://registry.npm.taobao.org/object.assign/download/object.assign-4.1.0.tgz",
      "integrity": "sha1-lovxEA15Vrs8oIbwBvhGs7xACNo=",
      "dev": true,
      "requires": {
        "define-properties": "^1.1.2",
        "function-bind": "^1.1.1",
        "has-symbols": "^1.0.0",
        "object-keys": "^1.0.11"
      }
    },
    "object.getownpropertydescriptors": {
      "version": "2.0.3",
      "resolved": "http://registry.npm.taobao.org/object.getownpropertydescriptors/download/object.getownpropertydescriptors-2.0.3.tgz",
      "integrity": "sha1-h1jIRvW0B62rDyNuCYbxSwUcqhY=",
      "dev": true,
      "requires": {
        "define-properties": "^1.1.2",
        "es-abstract": "^1.5.1"
      }
    },
    "object.pick": {
      "version": "1.3.0",
      "resolved": "http://registry.npm.taobao.org/object.pick/download/object.pick-1.3.0.tgz",
      "integrity": "sha1-h6EKxMFpS9Lhy/U1kaZhQftd10c=",
      "dev": true,
      "requires": {
        "isobject": "^3.0.1"
      }
    },
    "object.values": {
      "version": "1.0.4",
      "resolved": "http://registry.npm.taobao.org/object.values/download/object.values-1.0.4.tgz",
      "integrity": "sha1-5STaCbT2b/Bd9FdUbscqyZ8TBpo=",
      "dev": true,
      "requires": {
        "define-properties": "^1.1.2",
        "es-abstract": "^1.6.1",
        "function-bind": "^1.1.0",
        "has": "^1.0.1"
      }
    },
    "obuf": {
      "version": "1.1.2",
      "resolved": "http://registry.npm.taobao.org/obuf/download/obuf-1.1.2.tgz",
      "integrity": "sha1-Cb6jND1BhZ69RGKS0RydTbYZCE4=",
      "dev": true
    },
    "on-finished": {
      "version": "2.3.0",
      "resolved": "http://registry.npm.taobao.org/on-finished/download/on-finished-2.3.0.tgz",
      "integrity": "sha1-IPEzZIGwg811M3mSoWlxqi2QaUc=",
      "dev": true,
      "requires": {
        "ee-first": "1.1.1"
      }
    },
    "on-headers": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/on-headers/download/on-headers-1.0.1.tgz",
      "integrity": "sha1-ko9dD0cNSTQmUepnlLCFfBAGk/c=",
      "dev": true
    },
    "once": {
      "version": "1.4.0",
      "resolved": "http://registry.npm.taobao.org/once/download/once-1.4.0.tgz",
      "integrity": "sha1-WDsap3WWHUsROsF9nFC6753Xa9E=",
      "dev": true,
      "requires": {
        "wrappy": "1"
      }
    },
    "onetime": {
      "version": "2.0.1",
      "resolved": "http://registry.npm.taobao.org/onetime/download/onetime-2.0.1.tgz",
      "integrity": "sha1-BnQoIw/WdEOyeUsiu6UotoZ5YtQ=",
      "dev": true,
      "requires": {
        "mimic-fn": "^1.0.0"
      }
    },
    "opener": {
      "version": "1.5.1",
      "resolved": "http://registry.npm.taobao.org/opener/download/opener-1.5.1.tgz",
      "integrity": "sha1-bS8Od/GgrwAyrKcWwsH7uOfoq+0=",
      "dev": true
    },
    "opn": {
      "version": "5.4.0",
      "resolved": "http://registry.npm.taobao.org/opn/download/opn-5.4.0.tgz",
      "integrity": "sha1-y1Reeqt4VivrEao7+rxwQuF2EDU=",
      "dev": true,
      "requires": {
        "is-wsl": "^1.1.0"
      }
    },
    "optionator": {
      "version": "0.8.2",
      "resolved": "http://registry.npm.taobao.org/optionator/download/optionator-0.8.2.tgz",
      "integrity": "sha1-NkxeQJ0/TWMB1sC0wFu6UBgK62Q=",
      "dev": true,
      "requires": {
        "deep-is": "~0.1.3",
        "fast-levenshtein": "~2.0.4",
        "levn": "~0.3.0",
        "prelude-ls": "~1.1.2",
        "type-check": "~0.3.2",
        "wordwrap": "~1.0.0"
      }
    },
    "ora": {
      "version": "2.1.0",
      "resolved": "http://registry.npm.taobao.org/ora/download/ora-2.1.0.tgz",
      "integrity": "sha1-bK8oMOuSSUGGHsU6FzeZ4Ai1Hls=",
      "dev": true,
      "requires": {
        "chalk": "^2.3.1",
        "cli-cursor": "^2.1.0",
        "cli-spinners": "^1.1.0",
        "log-symbols": "^2.2.0",
        "strip-ansi": "^4.0.0",
        "wcwidth": "^1.0.1"
      }
    },
    "original": {
      "version": "1.0.2",
      "resolved": "http://registry.npm.taobao.org/original/download/original-1.0.2.tgz",
      "integrity": "sha1-5EKmHP/hxf0gpl8yYcJmY7MD8l8=",
      "dev": true,
      "requires": {
        "url-parse": "^1.4.3"
      }
    },
    "os-browserify": {
      "version": "0.3.0",
      "resolved": "http://registry.npm.taobao.org/os-browserify/download/os-browserify-0.3.0.tgz",
      "integrity": "sha1-hUNzx/XCMVkU/Jv8a9gjj92h7Cc=",
      "dev": true
    },
    "os-locale": {
      "version": "3.0.1",
      "resolved": "http://registry.npm.taobao.org/os-locale/download/os-locale-3.0.1.tgz",
      "integrity": "sha1-OwFPvwHYf2Ch5TSNgP6HDcgsRiA=",
      "dev": true,
      "requires": {
        "execa": "^0.10.0",
        "lcid": "^2.0.0",
        "mem": "^4.0.0"
      },
      "dependencies": {
        "execa": {
          "version": "0.10.0",
          "resolved": "http://registry.npm.taobao.org/execa/download/execa-0.10.0.tgz",
          "integrity": "sha1-/0Vqj1P5D47MxxqW0Rvfx/CCy1A=",
          "dev": true,
          "requires": {
            "cross-spawn": "^6.0.0",
            "get-stream": "^3.0.0",
            "is-stream": "^1.1.0",
            "npm-run-path": "^2.0.0",
            "p-finally": "^1.0.0",
            "signal-exit": "^3.0.0",
            "strip-eof": "^1.0.0"
          }
        },
        "get-stream": {
          "version": "3.0.0",
          "resolved": "http://registry.npm.taobao.org/get-stream/download/get-stream-3.0.0.tgz",
          "integrity": "sha1-jpQ9E1jcN1VQVOy+LtsFqhdO3hQ=",
          "dev": true
        }
      }
    },
    "os-tmpdir": {
      "version": "1.0.2",
      "resolved": "http://registry.npm.taobao.org/os-tmpdir/download/os-tmpdir-1.0.2.tgz",
      "integrity": "sha1-u+Z0BseaqFxc/sdm/lc0VV36EnQ=",
      "dev": true
    },
    "p-defer": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/p-defer/download/p-defer-1.0.0.tgz",
      "integrity": "sha1-n26xgvbJqozXQwBKfU+WsZaw+ww=",
      "dev": true
    },
    "p-finally": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/p-finally/download/p-finally-1.0.0.tgz",
      "integrity": "sha1-P7z7FbiZpEEjs0ttzBi3JDNqLK4=",
      "dev": true
    },
    "p-is-promise": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/p-is-promise/download/p-is-promise-1.1.0.tgz",
      "integrity": "sha1-nJRWmJ6fZYgBewQ01WCXZ1w9oF4=",
      "dev": true
    },
    "p-limit": {
      "version": "1.3.0",
      "resolved": "http://registry.npm.taobao.org/p-limit/download/p-limit-1.3.0.tgz",
      "integrity": "sha1-uGvV8MJWkJEcdZD8v8IBDVSzzLg=",
      "dev": true,
      "requires": {
        "p-try": "^1.0.0"
      }
    },
    "p-locate": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/p-locate/download/p-locate-2.0.0.tgz",
      "integrity": "sha1-IKAQOyIqcMj9OcwuWAaA893l7EM=",
      "dev": true,
      "requires": {
        "p-limit": "^1.1.0"
      }
    },
    "p-map": {
      "version": "1.2.0",
      "resolved": "http://registry.npm.taobao.org/p-map/download/p-map-1.2.0.tgz",
      "integrity": "sha1-5OlPMR6rvIYzoeeZCBZfyiYkG2s=",
      "dev": true
    },
    "p-try": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/p-try/download/p-try-1.0.0.tgz",
      "integrity": "sha1-y8ec26+P1CKOE/Yh8rGiN8GyB7M=",
      "dev": true
    },
    "pako": {
      "version": "1.0.7",
      "resolved": "http://registry.npm.taobao.org/pako/download/pako-1.0.7.tgz",
      "integrity": "sha1-JHNDkCG1fxUWyC9YvnJ1rY7xuyc=",
      "dev": true
    },
    "parallel-transform": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/parallel-transform/download/parallel-transform-1.1.0.tgz",
      "integrity": "sha1-1BDwZbBdojCB/NEPKIVMKb2jOwY=",
      "dev": true,
      "requires": {
        "cyclist": "~0.2.2",
        "inherits": "^2.0.3",
        "readable-stream": "^2.1.5"
      }
    },
    "param-case": {
      "version": "2.1.1",
      "resolved": "http://registry.npm.taobao.org/param-case/download/param-case-2.1.1.tgz",
      "integrity": "sha1-35T9jPZTHs915r75oIWPvHK+Ikc=",
      "dev": true,
      "requires": {
        "no-case": "^2.2.0"
      }
    },
    "parse-asn1": {
      "version": "5.1.1",
      "resolved": "http://registry.npm.taobao.org/parse-asn1/download/parse-asn1-5.1.1.tgz",
      "integrity": "sha1-9r8pOBgzK9DatU77Fgh3JHRebKg=",
      "dev": true,
      "requires": {
        "asn1.js": "^4.0.0",
        "browserify-aes": "^1.0.0",
        "create-hash": "^1.1.0",
        "evp_bytestokey": "^1.0.0",
        "pbkdf2": "^3.0.3"
      }
    },
    "parse-json": {
      "version": "4.0.0",
      "resolved": "http://registry.npm.taobao.org/parse-json/download/parse-json-4.0.0.tgz",
      "integrity": "sha1-vjX1Qlvh9/bHRxhPmKeIy5lHfuA=",
      "dev": true,
      "requires": {
        "error-ex": "^1.3.1",
        "json-parse-better-errors": "^1.0.1"
      }
    },
    "parseurl": {
      "version": "1.3.2",
      "resolved": "http://registry.npm.taobao.org/parseurl/download/parseurl-1.3.2.tgz",
      "integrity": "sha1-/CidTtiZMRlGDBViUyYs3I3mW/M=",
      "dev": true
    },
    "pascalcase": {
      "version": "0.1.1",
      "resolved": "http://registry.npm.taobao.org/pascalcase/download/pascalcase-0.1.1.tgz",
      "integrity": "sha1-s2PlXoAGym/iF4TS2yK9FdeRfxQ=",
      "dev": true
    },
    "path-browserify": {
      "version": "0.0.0",
      "resolved": "http://registry.npm.taobao.org/path-browserify/download/path-browserify-0.0.0.tgz",
      "integrity": "sha1-oLhwcpquIUAFt9UDLsLLuw+0RRo=",
      "dev": true
    },
    "path-dirname": {
      "version": "1.0.2",
      "resolved": "http://registry.npm.taobao.org/path-dirname/download/path-dirname-1.0.2.tgz",
      "integrity": "sha1-zDPSTVJeCZpTiMAzbG4yuRYGCeA=",
      "dev": true
    },
    "path-exists": {
      "version": "3.0.0",
      "resolved": "http://registry.npm.taobao.org/path-exists/download/path-exists-3.0.0.tgz",
      "integrity": "sha1-zg6+ql94yxiSXqfYENe1mwEP1RU=",
      "dev": true
    },
    "path-is-absolute": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/path-is-absolute/download/path-is-absolute-1.0.1.tgz",
      "integrity": "sha1-F0uSaHNVNP+8es5r9TpanhtcX18=",
      "dev": true
    },
    "path-is-inside": {
      "version": "1.0.2",
      "resolved": "http://registry.npm.taobao.org/path-is-inside/download/path-is-inside-1.0.2.tgz",
      "integrity": "sha1-NlQX3t5EQw0cEa9hAn+s8HS9/FM=",
      "dev": true
    },
    "path-key": {
      "version": "2.0.1",
      "resolved": "http://registry.npm.taobao.org/path-key/download/path-key-2.0.1.tgz",
      "integrity": "sha1-QRyttXTFoUDTpLGRDUDYDMn0C0A=",
      "dev": true
    },
    "path-parse": {
      "version": "1.0.6",
      "resolved": "http://registry.npm.taobao.org/path-parse/download/path-parse-1.0.6.tgz",
      "integrity": "sha1-1i27VnlAXXLEc37FhgDp3c8G0kw=",
      "dev": true
    },
    "path-to-regexp": {
      "version": "0.1.7",
      "resolved": "http://registry.npm.taobao.org/path-to-regexp/download/path-to-regexp-0.1.7.tgz",
      "integrity": "sha1-32BBeABfUi8V60SQ5yR6G/qmf4w=",
      "dev": true
    },
    "path-type": {
      "version": "3.0.0",
      "resolved": "http://registry.npm.taobao.org/path-type/download/path-type-3.0.0.tgz",
      "integrity": "sha1-zvMdyOCho7sNEFwM2Xzzv0f0428=",
      "dev": true,
      "requires": {
        "pify": "^3.0.0"
      }
    },
    "pbkdf2": {
      "version": "3.0.17",
      "resolved": "http://registry.npm.taobao.org/pbkdf2/download/pbkdf2-3.0.17.tgz",
      "integrity": "sha1-l2wgZTBhexTrsyEUI597CTNuk6Y=",
      "dev": true,
      "requires": {
        "create-hash": "^1.1.2",
        "create-hmac": "^1.1.4",
        "ripemd160": "^2.0.1",
        "safe-buffer": "^5.0.1",
        "sha.js": "^2.4.8"
      }
    },
    "performance-now": {
      "version": "2.1.0",
      "resolved": "http://registry.npm.taobao.org/performance-now/download/performance-now-2.1.0.tgz",
      "integrity": "sha1-Ywn04OX6kT7BxpMHrjZLSzd8nns=",
      "dev": true
    },
    "pify": {
      "version": "3.0.0",
      "resolved": "http://registry.npm.taobao.org/pify/download/pify-3.0.0.tgz",
      "integrity": "sha1-5aSs0sEB/fPZpNB/DbxNtJ3SgXY=",
      "dev": true
    },
    "pinkie": {
      "version": "2.0.4",
      "resolved": "http://registry.npm.taobao.org/pinkie/download/pinkie-2.0.4.tgz",
      "integrity": "sha1-clVrgM+g1IqXToDnckjoDtT3+HA=",
      "dev": true
    },
    "pinkie-promise": {
      "version": "2.0.1",
      "resolved": "http://registry.npm.taobao.org/pinkie-promise/download/pinkie-promise-2.0.1.tgz",
      "integrity": "sha1-ITXW36ejWMBprJsXh3YogihFD/o=",
      "dev": true,
      "requires": {
        "pinkie": "^2.0.0"
      }
    },
    "pkg-dir": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/pkg-dir/download/pkg-dir-2.0.0.tgz",
      "integrity": "sha1-9tXREJ4Z1j7fQo4L1X4Sd3YVM0s=",
      "dev": true,
      "requires": {
        "find-up": "^2.1.0"
      }
    },
    "pluralize": {
      "version": "7.0.0",
      "resolved": "http://registry.npm.taobao.org/pluralize/download/pluralize-7.0.0.tgz",
      "integrity": "sha1-KYuJ34uTsCIdv0Ia0rGx6iP8Z3c=",
      "dev": true
    },
    "portfinder": {
      "version": "1.0.20",
      "resolved": "http://registry.npm.taobao.org/portfinder/download/portfinder-1.0.20.tgz",
      "integrity": "sha1-vqaGMuVLLhOrewxHdem0G/Jw5Eo=",
      "dev": true,
      "requires": {
        "async": "^1.5.2",
        "debug": "^2.2.0",
        "mkdirp": "0.5.x"
      },
      "dependencies": {
        "debug": {
          "version": "2.6.9",
          "resolved": "http://registry.npm.taobao.org/debug/download/debug-2.6.9.tgz",
          "integrity": "sha1-XRKFFd8TT/Mn6QpMk/Tgd6U2NB8=",
          "dev": true,
          "requires": {
            "ms": "2.0.0"
          }
        },
        "ms": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/ms/download/ms-2.0.0.tgz",
          "integrity": "sha1-VgiurfwAvmwpAd9fmGF4jeDVl8g=",
          "dev": true
        }
      }
    },
    "posix-character-classes": {
      "version": "0.1.1",
      "resolved": "http://registry.npm.taobao.org/posix-character-classes/download/posix-character-classes-0.1.1.tgz",
      "integrity": "sha1-AerA/jta9xoqbAL+q7jB/vfgDqs=",
      "dev": true
    },
    "postcss": {
      "version": "7.0.6",
      "resolved": "http://registry.npm.taobao.org/postcss/download/postcss-7.0.6.tgz",
      "integrity": "sha1-bcqh6ZnN1KJV3NfU2VR/TKAQzcI=",
      "dev": true,
      "requires": {
        "chalk": "^2.4.1",
        "source-map": "^0.6.1",
        "supports-color": "^5.5.0"
      },
      "dependencies": {
        "source-map": {
          "version": "0.6.1",
          "resolved": "http://registry.npm.taobao.org/source-map/download/source-map-0.6.1.tgz",
          "integrity": "sha1-dHIq8y6WFOnCh6jQu95IteLxomM=",
          "dev": true
        }
      }
    },
    "postcss-calc": {
      "version": "7.0.1",
      "resolved": "http://registry.npm.taobao.org/postcss-calc/download/postcss-calc-7.0.1.tgz",
      "integrity": "sha1-Ntd7qwI7Dsu5eJ2E3LI8SUEUVDY=",
      "dev": true,
      "requires": {
        "css-unit-converter": "^1.1.1",
        "postcss": "^7.0.5",
        "postcss-selector-parser": "^5.0.0-rc.4",
        "postcss-value-parser": "^3.3.1"
      }
    },
    "postcss-colormin": {
      "version": "4.0.2",
      "resolved": "http://registry.npm.taobao.org/postcss-colormin/download/postcss-colormin-4.0.2.tgz",
      "integrity": "sha1-k80foRKAAIaWiH2xpSgEixjn7Zk=",
      "dev": true,
      "requires": {
        "browserslist": "^4.0.0",
        "color": "^3.0.0",
        "has": "^1.0.0",
        "postcss": "^7.0.0",
        "postcss-value-parser": "^3.0.0"
      }
    },
    "postcss-convert-values": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/postcss-convert-values/download/postcss-convert-values-4.0.1.tgz",
      "integrity": "sha1-yjgT7U2g+BL51DcDWE5Enr4Ymn8=",
      "dev": true,
      "requires": {
        "postcss": "^7.0.0",
        "postcss-value-parser": "^3.0.0"
      }
    },
    "postcss-discard-comments": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/postcss-discard-comments/download/postcss-discard-comments-4.0.1.tgz",
      "integrity": "sha1-MGl3NbDEdoUqehEFDrhDh6Z+9V0=",
      "dev": true,
      "requires": {
        "postcss": "^7.0.0"
      }
    },
    "postcss-discard-duplicates": {
      "version": "4.0.2",
      "resolved": "http://registry.npm.taobao.org/postcss-discard-duplicates/download/postcss-discard-duplicates-4.0.2.tgz",
      "integrity": "sha1-P+EzzTyCKC5VD8myORdqkge3hOs=",
      "dev": true,
      "requires": {
        "postcss": "^7.0.0"
      }
    },
    "postcss-discard-empty": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/postcss-discard-empty/download/postcss-discard-empty-4.0.1.tgz",
      "integrity": "sha1-yMlR6fc+2UKAGUWERKAq2Qu592U=",
      "dev": true,
      "requires": {
        "postcss": "^7.0.0"
      }
    },
    "postcss-discard-overridden": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/postcss-discard-overridden/download/postcss-discard-overridden-4.0.1.tgz",
      "integrity": "sha1-ZSrvipZybwKfXj4AFG7npOdV/1c=",
      "dev": true,
      "requires": {
        "postcss": "^7.0.0"
      }
    },
    "postcss-load-config": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/postcss-load-config/download/postcss-load-config-2.0.0.tgz",
      "integrity": "sha1-8TEt2/WRLNdHF3CDxe96GdYu5IQ=",
      "dev": true,
      "requires": {
        "cosmiconfig": "^4.0.0",
        "import-cwd": "^2.0.0"
      },
      "dependencies": {
        "cosmiconfig": {
          "version": "4.0.0",
          "resolved": "http://registry.npm.taobao.org/cosmiconfig/download/cosmiconfig-4.0.0.tgz",
          "integrity": "sha1-dgORVJWAu9LfHlYrwXexPCkJctw=",
          "dev": true,
          "requires": {
            "is-directory": "^0.3.1",
            "js-yaml": "^3.9.0",
            "parse-json": "^4.0.0",
            "require-from-string": "^2.0.1"
          }
        }
      }
    },
    "postcss-loader": {
      "version": "3.0.0",
      "resolved": "http://registry.npm.taobao.org/postcss-loader/download/postcss-loader-3.0.0.tgz",
      "integrity": "sha1-a5eUPkfHLYRfqeA/Jzdz1OjdbC0=",
      "dev": true,
      "requires": {
        "loader-utils": "^1.1.0",
        "postcss": "^7.0.0",
        "postcss-load-config": "^2.0.0",
        "schema-utils": "^1.0.0"
      },
      "dependencies": {
        "ajv-keywords": {
          "version": "3.2.0",
          "resolved": "http://registry.npm.taobao.org/ajv-keywords/download/ajv-keywords-3.2.0.tgz",
          "integrity": "sha1-6GuBnGAs+IIa1jdBNpjx3sAhhHo=",
          "dev": true
        },
        "schema-utils": {
          "version": "1.0.0",
          "resolved": "http://registry.npm.taobao.org/schema-utils/download/schema-utils-1.0.0.tgz",
          "integrity": "sha1-C3mpMgTXtgDUsoUNH2bCo0lRx3A=",
          "dev": true,
          "requires": {
            "ajv": "^6.1.0",
            "ajv-errors": "^1.0.0",
            "ajv-keywords": "^3.1.0"
          }
        }
      }
    },
    "postcss-merge-longhand": {
      "version": "4.0.9",
      "resolved": "http://registry.npm.taobao.org/postcss-merge-longhand/download/postcss-merge-longhand-4.0.9.tgz",
      "integrity": "sha1-wkKLmUgz/7KgcvKQymQudc6rzW8=",
      "dev": true,
      "requires": {
        "css-color-names": "0.0.4",
        "postcss": "^7.0.0",
        "postcss-value-parser": "^3.0.0",
        "stylehacks": "^4.0.0"
      }
    },
    "postcss-merge-rules": {
      "version": "4.0.2",
      "resolved": "http://registry.npm.taobao.org/postcss-merge-rules/download/postcss-merge-rules-4.0.2.tgz",
      "integrity": "sha1-K+REAb8ZhW8n8yuLEsDfWvG4jnQ=",
      "dev": true,
      "requires": {
        "browserslist": "^4.0.0",
        "caniuse-api": "^3.0.0",
        "cssnano-util-same-parent": "^4.0.0",
        "postcss": "^7.0.0",
        "postcss-selector-parser": "^3.0.0",
        "vendors": "^1.0.0"
      },
      "dependencies": {
        "postcss-selector-parser": {
          "version": "3.1.1",
          "resolved": "http://registry.npm.taobao.org/postcss-selector-parser/download/postcss-selector-parser-3.1.1.tgz",
          "integrity": "sha1-T4dfSvsMllc9XPTXQBGu4lCn6GU=",
          "dev": true,
          "requires": {
            "dot-prop": "^4.1.1",
            "indexes-of": "^1.0.1",
            "uniq": "^1.0.1"
          }
        }
      }
    },
    "postcss-minify-font-values": {
      "version": "4.0.2",
      "resolved": "http://registry.npm.taobao.org/postcss-minify-font-values/download/postcss-minify-font-values-4.0.2.tgz",
      "integrity": "sha1-zUw0TM5HQ0P6xdgiBqssvLiv1aY=",
      "dev": true,
      "requires": {
        "postcss": "^7.0.0",
        "postcss-value-parser": "^3.0.0"
      }
    },
    "postcss-minify-gradients": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/postcss-minify-gradients/download/postcss-minify-gradients-4.0.1.tgz",
      "integrity": "sha1-balcbpKoCflWu3a/DARJSVPhp90=",
      "dev": true,
      "requires": {
        "cssnano-util-get-arguments": "^4.0.0",
        "is-color-stop": "^1.0.0",
        "postcss": "^7.0.0",
        "postcss-value-parser": "^3.0.0"
      }
    },
    "postcss-minify-params": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/postcss-minify-params/download/postcss-minify-params-4.0.1.tgz",
      "integrity": "sha1-Wy4tAmTdZF711o+P7A1MOMHPk9I=",
      "dev": true,
      "requires": {
        "alphanum-sort": "^1.0.0",
        "browserslist": "^4.0.0",
        "cssnano-util-get-arguments": "^4.0.0",
        "postcss": "^7.0.0",
        "postcss-value-parser": "^3.0.0",
        "uniqs": "^2.0.0"
      }
    },
    "postcss-minify-selectors": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/postcss-minify-selectors/download/postcss-minify-selectors-4.0.1.tgz",
      "integrity": "sha1-qJHBl5d8w3q/YLPqBrhCSLHB6c0=",
      "dev": true,
      "requires": {
        "alphanum-sort": "^1.0.0",
        "has": "^1.0.0",
        "postcss": "^7.0.0",
        "postcss-selector-parser": "^3.0.0"
      },
      "dependencies": {
        "postcss-selector-parser": {
          "version": "3.1.1",
          "resolved": "http://registry.npm.taobao.org/postcss-selector-parser/download/postcss-selector-parser-3.1.1.tgz",
          "integrity": "sha1-T4dfSvsMllc9XPTXQBGu4lCn6GU=",
          "dev": true,
          "requires": {
            "dot-prop": "^4.1.1",
            "indexes-of": "^1.0.1",
            "uniq": "^1.0.1"
          }
        }
      }
    },
    "postcss-modules-extract-imports": {
      "version": "1.2.1",
      "resolved": "http://registry.npm.taobao.org/postcss-modules-extract-imports/download/postcss-modules-extract-imports-1.2.1.tgz",
      "integrity": "sha1-3IfjQUjsfqtfeR981YSYMzdbdBo=",
      "dev": true,
      "requires": {
        "postcss": "^6.0.1"
      },
      "dependencies": {
        "postcss": {
          "version": "6.0.23",
          "resolved": "http://registry.npm.taobao.org/postcss/download/postcss-6.0.23.tgz",
          "integrity": "sha1-YcgswyisYOZ3ZF+XkFTrmLwOMyQ=",
          "dev": true,
          "requires": {
            "chalk": "^2.4.1",
            "source-map": "^0.6.1",
            "supports-color": "^5.4.0"
          }
        },
        "source-map": {
          "version": "0.6.1",
          "resolved": "http://registry.npm.taobao.org/source-map/download/source-map-0.6.1.tgz",
          "integrity": "sha1-dHIq8y6WFOnCh6jQu95IteLxomM=",
          "dev": true
        }
      }
    },
    "postcss-modules-local-by-default": {
      "version": "1.2.0",
      "resolved": "http://registry.npm.taobao.org/postcss-modules-local-by-default/download/postcss-modules-local-by-default-1.2.0.tgz",
      "integrity": "sha1-99gMOYxaOT+nlkRmvRlQCn1hwGk=",
      "dev": true,
      "requires": {
        "css-selector-tokenizer": "^0.7.0",
        "postcss": "^6.0.1"
      },
      "dependencies": {
        "postcss": {
          "version": "6.0.23",
          "resolved": "http://registry.npm.taobao.org/postcss/download/postcss-6.0.23.tgz",
          "integrity": "sha1-YcgswyisYOZ3ZF+XkFTrmLwOMyQ=",
          "dev": true,
          "requires": {
            "chalk": "^2.4.1",
            "source-map": "^0.6.1",
            "supports-color": "^5.4.0"
          }
        },
        "source-map": {
          "version": "0.6.1",
          "resolved": "http://registry.npm.taobao.org/source-map/download/source-map-0.6.1.tgz",
          "integrity": "sha1-dHIq8y6WFOnCh6jQu95IteLxomM=",
          "dev": true
        }
      }
    },
    "postcss-modules-scope": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/postcss-modules-scope/download/postcss-modules-scope-1.1.0.tgz",
      "integrity": "sha1-1upkmUx5+XtipytCb75gVqGUu5A=",
      "dev": true,
      "requires": {
        "css-selector-tokenizer": "^0.7.0",
        "postcss": "^6.0.1"
      },
      "dependencies": {
        "postcss": {
          "version": "6.0.23",
          "resolved": "http://registry.npm.taobao.org/postcss/download/postcss-6.0.23.tgz",
          "integrity": "sha1-YcgswyisYOZ3ZF+XkFTrmLwOMyQ=",
          "dev": true,
          "requires": {
            "chalk": "^2.4.1",
            "source-map": "^0.6.1",
            "supports-color": "^5.4.0"
          }
        },
        "source-map": {
          "version": "0.6.1",
          "resolved": "http://registry.npm.taobao.org/source-map/download/source-map-0.6.1.tgz",
          "integrity": "sha1-dHIq8y6WFOnCh6jQu95IteLxomM=",
          "dev": true
        }
      }
    },
    "postcss-modules-values": {
      "version": "1.3.0",
      "resolved": "http://registry.npm.taobao.org/postcss-modules-values/download/postcss-modules-values-1.3.0.tgz",
      "integrity": "sha1-7P+p1+GSUYOJ9CrQ6D9yrsRW6iA=",
      "dev": true,
      "requires": {
        "icss-replace-symbols": "^1.1.0",
        "postcss": "^6.0.1"
      },
      "dependencies": {
        "postcss": {
          "version": "6.0.23",
          "resolved": "http://registry.npm.taobao.org/postcss/download/postcss-6.0.23.tgz",
          "integrity": "sha1-YcgswyisYOZ3ZF+XkFTrmLwOMyQ=",
          "dev": true,
          "requires": {
            "chalk": "^2.4.1",
            "source-map": "^0.6.1",
            "supports-color": "^5.4.0"
          }
        },
        "source-map": {
          "version": "0.6.1",
          "resolved": "http://registry.npm.taobao.org/source-map/download/source-map-0.6.1.tgz",
          "integrity": "sha1-dHIq8y6WFOnCh6jQu95IteLxomM=",
          "dev": true
        }
      }
    },
    "postcss-normalize-charset": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/postcss-normalize-charset/download/postcss-normalize-charset-4.0.1.tgz",
      "integrity": "sha1-izWt067oOhNrBHHg1ZvlilAoXdQ=",
      "dev": true,
      "requires": {
        "postcss": "^7.0.0"
      }
    },
    "postcss-normalize-display-values": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/postcss-normalize-display-values/download/postcss-normalize-display-values-4.0.1.tgz",
      "integrity": "sha1-2ag9R8cW6KmA8i9jLIsEWM+0ikw=",
      "dev": true,
      "requires": {
        "cssnano-util-get-match": "^4.0.0",
        "postcss": "^7.0.0",
        "postcss-value-parser": "^3.0.0"
      }
    },
    "postcss-normalize-positions": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/postcss-normalize-positions/download/postcss-normalize-positions-4.0.1.tgz",
      "integrity": "sha1-7i1LZ4GMlhlkxr4J0XmJS5T9a6E=",
      "dev": true,
      "requires": {
        "cssnano-util-get-arguments": "^4.0.0",
        "has": "^1.0.0",
        "postcss": "^7.0.0",
        "postcss-value-parser": "^3.0.0"
      }
    },
    "postcss-normalize-repeat-style": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/postcss-normalize-repeat-style/download/postcss-normalize-repeat-style-4.0.1.tgz",
      "integrity": "sha1-UpPyNLlNdmmp+AVJXTW4KlgcUOU=",
      "dev": true,
      "requires": {
        "cssnano-util-get-arguments": "^4.0.0",
        "cssnano-util-get-match": "^4.0.0",
        "postcss": "^7.0.0",
        "postcss-value-parser": "^3.0.0"
      }
    },
    "postcss-normalize-string": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/postcss-normalize-string/download/postcss-normalize-string-4.0.1.tgz",
      "integrity": "sha1-I8UDDCzCQXX2bJFPpRmeLjwQ/vM=",
      "dev": true,
      "requires": {
        "has": "^1.0.0",
        "postcss": "^7.0.0",
        "postcss-value-parser": "^3.0.0"
      }
    },
    "postcss-normalize-timing-functions": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/postcss-normalize-timing-functions/download/postcss-normalize-timing-functions-4.0.1.tgz",
      "integrity": "sha1-i+g+C5yz/y0avd7gMqSRCPBfldc=",
      "dev": true,
      "requires": {
        "cssnano-util-get-match": "^4.0.0",
        "postcss": "^7.0.0",
        "postcss-value-parser": "^3.0.0"
      }
    },
    "postcss-normalize-unicode": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/postcss-normalize-unicode/download/postcss-normalize-unicode-4.0.1.tgz",
      "integrity": "sha1-hBvUj9zzAZrUuqdJOj02O1KuHPs=",
      "dev": true,
      "requires": {
        "browserslist": "^4.0.0",
        "postcss": "^7.0.0",
        "postcss-value-parser": "^3.0.0"
      }
    },
    "postcss-normalize-url": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/postcss-normalize-url/download/postcss-normalize-url-4.0.1.tgz",
      "integrity": "sha1-EOQ3+GvHx+WPe5ZS7YeNqqlfquE=",
      "dev": true,
      "requires": {
        "is-absolute-url": "^2.0.0",
        "normalize-url": "^3.0.0",
        "postcss": "^7.0.0",
        "postcss-value-parser": "^3.0.0"
      }
    },
    "postcss-normalize-whitespace": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/postcss-normalize-whitespace/download/postcss-normalize-whitespace-4.0.1.tgz",
      "integrity": "sha1-0Uy2ObYSOEGKyLyNO3vdZfyGV14=",
      "dev": true,
      "requires": {
        "postcss": "^7.0.0",
        "postcss-value-parser": "^3.0.0"
      }
    },
    "postcss-ordered-values": {
      "version": "4.1.1",
      "resolved": "http://registry.npm.taobao.org/postcss-ordered-values/download/postcss-ordered-values-4.1.1.tgz",
      "integrity": "sha1-LjtDLvPkibGDM67KHxKV64m+n8I=",
      "dev": true,
      "requires": {
        "cssnano-util-get-arguments": "^4.0.0",
        "postcss": "^7.0.0",
        "postcss-value-parser": "^3.0.0"
      }
    },
    "postcss-reduce-initial": {
      "version": "4.0.2",
      "resolved": "http://registry.npm.taobao.org/postcss-reduce-initial/download/postcss-reduce-initial-4.0.2.tgz",
      "integrity": "sha1-usjjJdZ1EO4B+kYGdtyOqeO0DxU=",
      "dev": true,
      "requires": {
        "browserslist": "^4.0.0",
        "caniuse-api": "^3.0.0",
        "has": "^1.0.0",
        "postcss": "^7.0.0"
      }
    },
    "postcss-reduce-transforms": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/postcss-reduce-transforms/download/postcss-reduce-transforms-4.0.1.tgz",
      "integrity": "sha1-hgDVVTvdOtZA9Dv/getS+HYNRWE=",
      "dev": true,
      "requires": {
        "cssnano-util-get-match": "^4.0.0",
        "has": "^1.0.0",
        "postcss": "^7.0.0",
        "postcss-value-parser": "^3.0.0"
      }
    },
    "postcss-selector-parser": {
      "version": "5.0.0-rc.4",
      "resolved": "http://registry.npm.taobao.org/postcss-selector-parser/download/postcss-selector-parser-5.0.0-rc.4.tgz",
      "integrity": "sha1-yl53I4vxUpZjeME+ka1tYRVo6oc=",
      "dev": true,
      "requires": {
        "cssesc": "^2.0.0",
        "indexes-of": "^1.0.1",
        "uniq": "^1.0.1"
      }
    },
    "postcss-svgo": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/postcss-svgo/download/postcss-svgo-4.0.1.tgz",
      "integrity": "sha1-VijNs48BXea1iM5tC/ByS0krWB0=",
      "dev": true,
      "requires": {
        "is-svg": "^3.0.0",
        "postcss": "^7.0.0",
        "postcss-value-parser": "^3.0.0",
        "svgo": "^1.0.0"
      }
    },
    "postcss-unique-selectors": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/postcss-unique-selectors/download/postcss-unique-selectors-4.0.1.tgz",
      "integrity": "sha1-lEaRHzKJv9ZMbWgPBzwDsfnuS6w=",
      "dev": true,
      "requires": {
        "alphanum-sort": "^1.0.0",
        "postcss": "^7.0.0",
        "uniqs": "^2.0.0"
      }
    },
    "postcss-value-parser": {
      "version": "3.3.1",
      "resolved": "http://registry.npm.taobao.org/postcss-value-parser/download/postcss-value-parser-3.3.1.tgz",
      "integrity": "sha1-n/giVH4okyE88cMO+lGsX9G6goE=",
      "dev": true
    },
    "prelude-ls": {
      "version": "1.1.2",
      "resolved": "http://registry.npm.taobao.org/prelude-ls/download/prelude-ls-1.1.2.tgz",
      "integrity": "sha1-IZMqVJ9eUv/ZqCf1cOBL5iqX2lQ=",
      "dev": true
    },
    "prettier": {
      "version": "1.13.7",
      "resolved": "http://registry.npm.taobao.org/prettier/download/prettier-1.13.7.tgz",
      "integrity": "sha1-hQ87iveEpJpuotLqp+0UKKNLcoE=",
      "dev": true
    },
    "pretty-error": {
      "version": "2.1.1",
      "resolved": "http://registry.npm.taobao.org/pretty-error/download/pretty-error-2.1.1.tgz",
      "integrity": "sha1-X0+HyPkeWuPzuoerTPXgOxoX8aM=",
      "dev": true,
      "requires": {
        "renderkid": "^2.0.1",
        "utila": "~0.4"
      }
    },
    "private": {
      "version": "0.1.8",
      "resolved": "http://registry.npm.taobao.org/private/download/private-0.1.8.tgz",
      "integrity": "sha1-I4Hts2ifelPWUxkAYPz4ItLzaP8=",
      "dev": true
    },
    "process": {
      "version": "0.11.10",
      "resolved": "http://registry.npm.taobao.org/process/download/process-0.11.10.tgz",
      "integrity": "sha1-czIwDoQBYb2j5podHZGn1LwW8YI=",
      "dev": true
    },
    "process-nextick-args": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/process-nextick-args/download/process-nextick-args-2.0.0.tgz",
      "integrity": "sha1-o31zL0JxtKsa0HDTVQjoKQeI/6o=",
      "dev": true
    },
    "progress": {
      "version": "2.0.3",
      "resolved": "http://registry.npm.taobao.org/progress/download/progress-2.0.3.tgz",
      "integrity": "sha1-foz42PW48jnBvGi+tOt4Vn1XLvg=",
      "dev": true
    },
    "promise-inflight": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/promise-inflight/download/promise-inflight-1.0.1.tgz",
      "integrity": "sha1-mEcocL8igTL8vdhoEputEsPAKeM=",
      "dev": true
    },
    "proxy-addr": {
      "version": "2.0.4",
      "resolved": "http://registry.npm.taobao.org/proxy-addr/download/proxy-addr-2.0.4.tgz",
      "integrity": "sha1-7PxzO/Iv+Mb0B/onUye5q2fki5M=",
      "dev": true,
      "requires": {
        "forwarded": "~0.1.2",
        "ipaddr.js": "1.8.0"
      }
    },
    "prr": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/prr/download/prr-1.0.1.tgz",
      "integrity": "sha1-0/wRS6BplaRexok/SEzrHXj19HY=",
      "dev": true
    },
    "pseudomap": {
      "version": "1.0.2",
      "resolved": "http://registry.npm.taobao.org/pseudomap/download/pseudomap-1.0.2.tgz",
      "integrity": "sha1-8FKijacOYYkX7wqKw0wa5aaChrM=",
      "dev": true
    },
    "psl": {
      "version": "1.1.31",
      "resolved": "http://registry.npm.taobao.org/psl/download/psl-1.1.31.tgz",
      "integrity": "sha1-6aqG0BAbWxBcvpOsa3hM1UcnYYQ=",
      "dev": true
    },
    "public-encrypt": {
      "version": "4.0.3",
      "resolved": "http://registry.npm.taobao.org/public-encrypt/download/public-encrypt-4.0.3.tgz",
      "integrity": "sha1-T8ydd6B+SLp1J+fL4N4z0HATMeA=",
      "dev": true,
      "requires": {
        "bn.js": "^4.1.0",
        "browserify-rsa": "^4.0.0",
        "create-hash": "^1.1.0",
        "parse-asn1": "^5.0.0",
        "randombytes": "^2.0.1",
        "safe-buffer": "^5.1.2"
      }
    },
    "pump": {
      "version": "3.0.0",
      "resolved": "http://registry.npm.taobao.org/pump/download/pump-3.0.0.tgz",
      "integrity": "sha1-tKIRaBW94vTh6mAjVOjHVWUQemQ=",
      "dev": true,
      "requires": {
        "end-of-stream": "^1.1.0",
        "once": "^1.3.1"
      }
    },
    "pumpify": {
      "version": "1.5.1",
      "resolved": "http://registry.npm.taobao.org/pumpify/download/pumpify-1.5.1.tgz",
      "integrity": "sha1-NlE74karJ1cLGjdKXOJ4v9dDcM4=",
      "dev": true,
      "requires": {
        "duplexify": "^3.6.0",
        "inherits": "^2.0.3",
        "pump": "^2.0.0"
      },
      "dependencies": {
        "pump": {
          "version": "2.0.1",
          "resolved": "http://registry.npm.taobao.org/pump/download/pump-2.0.1.tgz",
          "integrity": "sha1-Ejma3W5M91Jtlzy8i1zi4pCLOQk=",
          "dev": true,
          "requires": {
            "end-of-stream": "^1.1.0",
            "once": "^1.3.1"
          }
        }
      }
    },
    "punycode": {
      "version": "2.1.1",
      "resolved": "http://registry.npm.taobao.org/punycode/download/punycode-2.1.1.tgz",
      "integrity": "sha1-tYsBCsQMIsVldhbI0sLALHv0eew=",
      "dev": true
    },
    "q": {
      "version": "1.5.1",
      "resolved": "http://registry.npm.taobao.org/q/download/q-1.5.1.tgz",
      "integrity": "sha1-fjL3W0E4EpHQRhHxvxQQmsAGUdc=",
      "dev": true
    },
    "qs": {
      "version": "6.5.2",
      "resolved": "http://registry.npm.taobao.org/qs/download/qs-6.5.2.tgz",
      "integrity": "sha1-yzroBuh0BERYTvFUzo7pjUA/PjY=",
      "dev": true
    },
    "querystring": {
      "version": "0.2.0",
      "resolved": "http://registry.npm.taobao.org/querystring/download/querystring-0.2.0.tgz",
      "integrity": "sha1-sgmEkgO7Jd+CDadW50cAWHhSFiA=",
      "dev": true
    },
    "querystring-es3": {
      "version": "0.2.1",
      "resolved": "http://registry.npm.taobao.org/querystring-es3/download/querystring-es3-0.2.1.tgz",
      "integrity": "sha1-nsYfeQSYdXB9aUFFlv2Qek1xHnM=",
      "dev": true
    },
    "querystringify": {
      "version": "2.1.0",
      "resolved": "http://registry.npm.taobao.org/querystringify/download/querystringify-2.1.0.tgz",
      "integrity": "sha1-fe2N+/eHncxg0KZErGdUsoOtF+8=",
      "dev": true
    },
    "randombytes": {
      "version": "2.0.6",
      "resolved": "http://registry.npm.taobao.org/randombytes/download/randombytes-2.0.6.tgz",
      "integrity": "sha1-0wLFIpSFiISKjTAMkytEwkIx2oA=",
      "dev": true,
      "requires": {
        "safe-buffer": "^5.1.0"
      }
    },
    "randomfill": {
      "version": "1.0.4",
      "resolved": "http://registry.npm.taobao.org/randomfill/download/randomfill-1.0.4.tgz",
      "integrity": "sha1-ySGW/IarQr6YPxvzF3giSTHWFFg=",
      "dev": true,
      "requires": {
        "randombytes": "^2.0.5",
        "safe-buffer": "^5.1.0"
      }
    },
    "range-parser": {
      "version": "1.2.0",
      "resolved": "http://registry.npm.taobao.org/range-parser/download/range-parser-1.2.0.tgz",
      "integrity": "sha1-9JvmtIeJTdxA3MlKMi9hEJLgDV4=",
      "dev": true
    },
    "raw-body": {
      "version": "2.3.3",
      "resolved": "http://registry.npm.taobao.org/raw-body/download/raw-body-2.3.3.tgz",
      "integrity": "sha1-GzJOzmtXBuFThVvBFIxlu39uoMM=",
      "dev": true,
      "requires": {
        "bytes": "3.0.0",
        "http-errors": "1.6.3",
        "iconv-lite": "0.4.23",
        "unpipe": "1.0.0"
      },
      "dependencies": {
        "iconv-lite": {
          "version": "0.4.23",
          "resolved": "http://registry.npm.taobao.org/iconv-lite/download/iconv-lite-0.4.23.tgz",
          "integrity": "sha1-KXhx9jvlB63Pv8pxXQzQ7thOmmM=",
          "dev": true,
          "requires": {
            "safer-buffer": ">= 2.1.2 < 3"
          }
        }
      }
    },
    "read-pkg": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/read-pkg/download/read-pkg-4.0.1.tgz",
      "integrity": "sha1-ljYlN48+HE1IyFhytabsfV0JMjc=",
      "dev": true,
      "requires": {
        "normalize-package-data": "^2.3.2",
        "parse-json": "^4.0.0",
        "pify": "^3.0.0"
      }
    },
    "readable-stream": {
      "version": "2.3.6",
      "resolved": "http://registry.npm.taobao.org/readable-stream/download/readable-stream-2.3.6.tgz",
      "integrity": "sha1-sRwn2IuP8fvgcGQ8+UsMea4bCq8=",
      "dev": true,
      "requires": {
        "core-util-is": "~1.0.0",
        "inherits": "~2.0.3",
        "isarray": "~1.0.0",
        "process-nextick-args": "~2.0.0",
        "safe-buffer": "~5.1.1",
        "string_decoder": "~1.1.1",
        "util-deprecate": "~1.0.1"
      }
    },
    "readdirp": {
      "version": "2.2.1",
      "resolved": "http://registry.npm.taobao.org/readdirp/download/readdirp-2.2.1.tgz",
      "integrity": "sha1-DodiKjMlqjPokihcr4tOhGUppSU=",
      "dev": true,
      "requires": {
        "graceful-fs": "^4.1.11",
        "micromatch": "^3.1.10",
        "readable-stream": "^2.0.2"
      }
    },
    "regenerate": {
      "version": "1.4.0",
      "resolved": "http://registry.npm.taobao.org/regenerate/download/regenerate-1.4.0.tgz",
      "integrity": "sha1-SoVuxLVuQHfFV1icroXnpMiGmhE=",
      "dev": true
    },
    "regenerate-unicode-properties": {
      "version": "7.0.0",
      "resolved": "http://registry.npm.taobao.org/regenerate-unicode-properties/download/regenerate-unicode-properties-7.0.0.tgz",
      "integrity": "sha1-EHQFr8xKGQ7F7UUOyqAO0Mr6ekw=",
      "dev": true,
      "requires": {
        "regenerate": "^1.4.0"
      }
    },
    "regenerator-runtime": {
      "version": "0.12.1",
      "resolved": "http://registry.npm.taobao.org/regenerator-runtime/download/regenerator-runtime-0.12.1.tgz",
      "integrity": "sha1-+hpxVEdkwDb4xJsToIsllMn4oN4=",
      "dev": true
    },
    "regenerator-transform": {
      "version": "0.13.3",
      "resolved": "http://registry.npm.taobao.org/regenerator-transform/download/regenerator-transform-0.13.3.tgz",
      "integrity": "sha1-JkvZ/zioziSwbgY2SWsshWtXvLs=",
      "dev": true,
      "requires": {
        "private": "^0.1.6"
      }
    },
    "regex-not": {
      "version": "1.0.2",
      "resolved": "http://registry.npm.taobao.org/regex-not/download/regex-not-1.0.2.tgz",
      "integrity": "sha1-H07OJ+ALC2XgJHpoEOaoXYOldSw=",
      "dev": true,
      "requires": {
        "extend-shallow": "^3.0.2",
        "safe-regex": "^1.1.0"
      }
    },
    "regexpp": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/regexpp/download/regexpp-1.1.0.tgz",
      "integrity": "sha1-DjUW3Qt5BPQT0tQZPc5GGMOmias=",
      "dev": true
    },
    "regexpu-core": {
      "version": "4.4.0",
      "resolved": "http://registry.npm.taobao.org/regexpu-core/download/regexpu-core-4.4.0.tgz",
      "integrity": "sha1-jUPg0SZog5aXIDRecMJ17grsDTI=",
      "dev": true,
      "requires": {
        "regenerate": "^1.4.0",
        "regenerate-unicode-properties": "^7.0.0",
        "regjsgen": "^0.5.0",
        "regjsparser": "^0.6.0",
        "unicode-match-property-ecmascript": "^1.0.4",
        "unicode-match-property-value-ecmascript": "^1.0.2"
      }
    },
    "regjsgen": {
      "version": "0.5.0",
      "resolved": "http://registry.npm.taobao.org/regjsgen/download/regjsgen-0.5.0.tgz",
      "integrity": "sha1-p2NNwI+JIJwgSa3aNSVxH7lyZd0=",
      "dev": true
    },
    "regjsparser": {
      "version": "0.6.0",
      "resolved": "http://registry.npm.taobao.org/regjsparser/download/regjsparser-0.6.0.tgz",
      "integrity": "sha1-8eaui32iuulsmTmbhozWyTOiupw=",
      "dev": true,
      "requires": {
        "jsesc": "~0.5.0"
      },
      "dependencies": {
        "jsesc": {
          "version": "0.5.0",
          "resolved": "http://registry.npm.taobao.org/jsesc/download/jsesc-0.5.0.tgz",
          "integrity": "sha1-597mbjXW/Bb3EP6R1c9p9w8IkR0=",
          "dev": true
        }
      }
    },
    "relateurl": {
      "version": "0.2.7",
      "resolved": "http://registry.npm.taobao.org/relateurl/download/relateurl-0.2.7.tgz",
      "integrity": "sha1-VNvzd+UUQKypCkzSdGANP/LYiKk=",
      "dev": true
    },
    "remove-trailing-separator": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/remove-trailing-separator/download/remove-trailing-separator-1.1.0.tgz",
      "integrity": "sha1-wkvOKig62tW8P1jg1IJJuSN52O8=",
      "dev": true
    },
    "renderkid": {
      "version": "2.0.2",
      "resolved": "http://registry.npm.taobao.org/renderkid/download/renderkid-2.0.2.tgz",
      "integrity": "sha1-EtMQ8lU2DAetj94lP2yeneNy0qo=",
      "dev": true,
      "requires": {
        "css-select": "^1.1.0",
        "dom-converter": "~0.2",
        "htmlparser2": "~3.3.0",
        "strip-ansi": "^3.0.0",
        "utila": "^0.4.0"
      },
      "dependencies": {
        "ansi-regex": {
          "version": "2.1.1",
          "resolved": "http://registry.npm.taobao.org/ansi-regex/download/ansi-regex-2.1.1.tgz",
          "integrity": "sha1-w7M6te42DYbg5ijwRorn7yfWVN8=",
          "dev": true
        },
        "css-select": {
          "version": "1.2.0",
          "resolved": "http://registry.npm.taobao.org/css-select/download/css-select-1.2.0.tgz",
          "integrity": "sha1-KzoRBTnFNV8c2NMUYj6HCxIeyFg=",
          "dev": true,
          "requires": {
            "boolbase": "~1.0.0",
            "css-what": "2.1",
            "domutils": "1.5.1",
            "nth-check": "~1.0.1"
          }
        },
        "domutils": {
          "version": "1.5.1",
          "resolved": "http://registry.npm.taobao.org/domutils/download/domutils-1.5.1.tgz",
          "integrity": "sha1-3NhIiib1Y9YQeeSMn3t+Mjc2gs8=",
          "dev": true,
          "requires": {
            "dom-serializer": "0",
            "domelementtype": "1"
          }
        },
        "strip-ansi": {
          "version": "3.0.1",
          "resolved": "http://registry.npm.taobao.org/strip-ansi/download/strip-ansi-3.0.1.tgz",
          "integrity": "sha1-ajhfuIU9lS1f8F0Oiq+UJ43GPc8=",
          "dev": true,
          "requires": {
            "ansi-regex": "^2.0.0"
          }
        }
      }
    },
    "repeat-element": {
      "version": "1.1.3",
      "resolved": "http://registry.npm.taobao.org/repeat-element/download/repeat-element-1.1.3.tgz",
      "integrity": "sha1-eC4NglwMWjuzlzH4Tv7mt0Lmsc4=",
      "dev": true
    },
    "repeat-string": {
      "version": "1.6.1",
      "resolved": "http://registry.npm.taobao.org/repeat-string/download/repeat-string-1.6.1.tgz",
      "integrity": "sha1-jcrkcOHIirwtYA//Sndihtp15jc=",
      "dev": true
    },
    "request": {
      "version": "2.88.0",
      "resolved": "http://registry.npm.taobao.org/request/download/request-2.88.0.tgz",
      "integrity": "sha1-nC/KT301tZLv5Xx/ClXoEFIST+8=",
      "dev": true,
      "requires": {
        "aws-sign2": "~0.7.0",
        "aws4": "^1.8.0",
        "caseless": "~0.12.0",
        "combined-stream": "~1.0.6",
        "extend": "~3.0.2",
        "forever-agent": "~0.6.1",
        "form-data": "~2.3.2",
        "har-validator": "~5.1.0",
        "http-signature": "~1.2.0",
        "is-typedarray": "~1.0.0",
        "isstream": "~0.1.2",
        "json-stringify-safe": "~5.0.1",
        "mime-types": "~2.1.19",
        "oauth-sign": "~0.9.0",
        "performance-now": "^2.1.0",
        "qs": "~6.5.2",
        "safe-buffer": "^5.1.2",
        "tough-cookie": "~2.4.3",
        "tunnel-agent": "^0.6.0",
        "uuid": "^3.3.2"
      }
    },
    "request-promise-core": {
      "version": "1.1.1",
      "resolved": "http://registry.npm.taobao.org/request-promise-core/download/request-promise-core-1.1.1.tgz",
      "integrity": "sha1-Pu4AssWqgyOc+wTFcA2jb4HNCLY=",
      "dev": true,
      "requires": {
        "lodash": "^4.13.1"
      }
    },
    "request-promise-native": {
      "version": "1.0.5",
      "resolved": "http://registry.npm.taobao.org/request-promise-native/download/request-promise-native-1.0.5.tgz",
      "integrity": "sha1-UoF3D2jgyXGeUWP9P6tIIhX0/aU=",
      "dev": true,
      "requires": {
        "request-promise-core": "1.1.1",
        "stealthy-require": "^1.1.0",
        "tough-cookie": ">=2.3.3"
      }
    },
    "require-directory": {
      "version": "2.1.1",
      "resolved": "http://registry.npm.taobao.org/require-directory/download/require-directory-2.1.1.tgz",
      "integrity": "sha1-jGStX9MNqxyXbiNE/+f3kqam30I=",
      "dev": true
    },
    "require-from-string": {
      "version": "2.0.2",
      "resolved": "http://registry.npm.taobao.org/require-from-string/download/require-from-string-2.0.2.tgz",
      "integrity": "sha1-iaf92TgmEmcxjq/hT5wy5ZjDaQk=",
      "dev": true
    },
    "require-main-filename": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/require-main-filename/download/require-main-filename-1.0.1.tgz",
      "integrity": "sha1-l/cXtp1IeE9fUmpsWqj/3aBVpNE=",
      "dev": true
    },
    "require-uncached": {
      "version": "1.0.3",
      "resolved": "http://registry.npm.taobao.org/require-uncached/download/require-uncached-1.0.3.tgz",
      "integrity": "sha1-Tg1W1slmL9MeQwEcS5WqSZVUIdM=",
      "dev": true,
      "requires": {
        "caller-path": "^0.1.0",
        "resolve-from": "^1.0.0"
      }
    },
    "requires-port": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/requires-port/download/requires-port-1.0.0.tgz",
      "integrity": "sha1-kl0mAdOaxIXgkc8NpcbmlNw9yv8=",
      "dev": true
    },
    "resolve": {
      "version": "1.8.1",
      "resolved": "http://registry.npm.taobao.org/resolve/download/resolve-1.8.1.tgz",
      "integrity": "sha1-gvHsGaQjrB+9CAsLqwa6NuhKeiY=",
      "dev": true,
      "requires": {
        "path-parse": "^1.0.5"
      }
    },
    "resolve-cwd": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/resolve-cwd/download/resolve-cwd-2.0.0.tgz",
      "integrity": "sha1-AKn3OHVW4nA46uIyyqNypqWbZlo=",
      "dev": true,
      "requires": {
        "resolve-from": "^3.0.0"
      },
      "dependencies": {
        "resolve-from": {
          "version": "3.0.0",
          "resolved": "http://registry.npm.taobao.org/resolve-from/download/resolve-from-3.0.0.tgz",
          "integrity": "sha1-six699nWiBvItuZTM17rywoYh0g=",
          "dev": true
        }
      }
    },
    "resolve-from": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/resolve-from/download/resolve-from-1.0.1.tgz",
      "integrity": "sha1-Jsv+k10a7uq7Kbw/5a6wHpPUQiY=",
      "dev": true
    },
    "resolve-url": {
      "version": "0.2.1",
      "resolved": "http://registry.npm.taobao.org/resolve-url/download/resolve-url-0.2.1.tgz",
      "integrity": "sha1-LGN/53yJOv0qZj/iGqkIAGjiBSo=",
      "dev": true
    },
    "restore-cursor": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/restore-cursor/download/restore-cursor-2.0.0.tgz",
      "integrity": "sha1-n37ih/gv0ybU/RYpI9YhKe7g368=",
      "dev": true,
      "requires": {
        "onetime": "^2.0.0",
        "signal-exit": "^3.0.2"
      }
    },
    "ret": {
      "version": "0.1.15",
      "resolved": "http://registry.npm.taobao.org/ret/download/ret-0.1.15.tgz",
      "integrity": "sha1-uKSCXVvbH8P29Twrwz+BOIaBx7w=",
      "dev": true
    },
    "rgb-regex": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/rgb-regex/download/rgb-regex-1.0.1.tgz",
      "integrity": "sha1-wODWiC3w4jviVKR16O3UGRX+rrE=",
      "dev": true
    },
    "rgba-regex": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/rgba-regex/download/rgba-regex-1.0.0.tgz",
      "integrity": "sha1-QzdOLiyglosO8VI0YLfXMP8i7rM=",
      "dev": true
    },
    "rimraf": {
      "version": "2.6.2",
      "resolved": "http://registry.npm.taobao.org/rimraf/download/rimraf-2.6.2.tgz",
      "integrity": "sha1-LtgVDSShbqhlHm1u8PR8QVjOejY=",
      "dev": true,
      "requires": {
        "glob": "^7.0.5"
      }
    },
    "ripemd160": {
      "version": "2.0.2",
      "resolved": "http://registry.npm.taobao.org/ripemd160/download/ripemd160-2.0.2.tgz",
      "integrity": "sha1-ocGm9iR1FXe6XQeRTLyShQWFiQw=",
      "dev": true,
      "requires": {
        "hash-base": "^3.0.0",
        "inherits": "^2.0.1"
      }
    },
    "run-async": {
      "version": "2.3.0",
      "resolved": "http://registry.npm.taobao.org/run-async/download/run-async-2.3.0.tgz",
      "integrity": "sha1-A3GrSuC91yDUFm19/aZP96RFpsA=",
      "dev": true,
      "requires": {
        "is-promise": "^2.1.0"
      }
    },
    "run-queue": {
      "version": "1.0.3",
      "resolved": "http://registry.npm.taobao.org/run-queue/download/run-queue-1.0.3.tgz",
      "integrity": "sha1-6Eg5bwV9Ij8kOGkkYY4laUFh7Ec=",
      "dev": true,
      "requires": {
        "aproba": "^1.1.1"
      }
    },
    "rx-lite": {
      "version": "4.0.8",
      "resolved": "http://registry.npm.taobao.org/rx-lite/download/rx-lite-4.0.8.tgz",
      "integrity": "sha1-Cx4Rr4vESDbwSmQH6S2kJGe3lEQ=",
      "dev": true
    },
    "rx-lite-aggregates": {
      "version": "4.0.8",
      "resolved": "http://registry.npm.taobao.org/rx-lite-aggregates/download/rx-lite-aggregates-4.0.8.tgz",
      "integrity": "sha1-dTuHqJoRyVRnxKwWJsTvxOBcZ74=",
      "dev": true,
      "requires": {
        "rx-lite": "*"
      }
    },
    "rxjs": {
      "version": "6.3.3",
      "resolved": "http://registry.npm.taobao.org/rxjs/download/rxjs-6.3.3.tgz",
      "integrity": "sha1-PGp/pCDoRKgTkPsRWKnsYU9LrVU=",
      "dev": true,
      "requires": {
        "tslib": "^1.9.0"
      }
    },
    "safe-buffer": {
      "version": "5.1.2",
      "resolved": "http://registry.npm.taobao.org/safe-buffer/download/safe-buffer-5.1.2.tgz",
      "integrity": "sha1-mR7GnSluAxN0fVm9/St0XDX4go0=",
      "dev": true
    },
    "safe-regex": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/safe-regex/download/safe-regex-1.1.0.tgz",
      "integrity": "sha1-QKNmnzsHfR6UPURinhV91IAjvy4=",
      "dev": true,
      "requires": {
        "ret": "~0.1.10"
      }
    },
    "safer-buffer": {
      "version": "2.1.2",
      "resolved": "http://registry.npm.taobao.org/safer-buffer/download/safer-buffer-2.1.2.tgz",
      "integrity": "sha1-RPoWGwGHuVSd2Eu5GAL5vYOFzWo=",
      "dev": true
    },
    "sax": {
      "version": "1.2.4",
      "resolved": "http://registry.npm.taobao.org/sax/download/sax-1.2.4.tgz",
      "integrity": "sha1-KBYjTiN4vdxOU1T6tcqold9xANk=",
      "dev": true
    },
    "schema-utils": {
      "version": "0.4.7",
      "resolved": "http://registry.npm.taobao.org/schema-utils/download/schema-utils-0.4.7.tgz",
      "integrity": "sha1-unT1l9K+LqiAExdG7hfQoJPGgYc=",
      "dev": true,
      "requires": {
        "ajv": "^6.1.0",
        "ajv-keywords": "^3.1.0"
      },
      "dependencies": {
        "ajv-keywords": {
          "version": "3.2.0",
          "resolved": "http://registry.npm.taobao.org/ajv-keywords/download/ajv-keywords-3.2.0.tgz",
          "integrity": "sha1-6GuBnGAs+IIa1jdBNpjx3sAhhHo=",
          "dev": true
        }
      }
    },
    "select-hose": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/select-hose/download/select-hose-2.0.0.tgz",
      "integrity": "sha1-Yl2GWPhlr0Psliv8N2o3NZpJlMo=",
      "dev": true
    },
    "selfsigned": {
      "version": "1.10.4",
      "resolved": "http://registry.npm.taobao.org/selfsigned/download/selfsigned-1.10.4.tgz",
      "integrity": "sha1-zdfsz8pO12NdR6CL8tXTB0CS4s0=",
      "dev": true,
      "requires": {
        "node-forge": "0.7.5"
      }
    },
    "semver": {
      "version": "5.6.0",
      "resolved": "http://registry.npm.taobao.org/semver/download/semver-5.6.0.tgz",
      "integrity": "sha1-fnQlb7qknHWqfHogXMInmcrIAAQ=",
      "dev": true
    },
    "send": {
      "version": "0.16.2",
      "resolved": "http://registry.npm.taobao.org/send/download/send-0.16.2.tgz",
      "integrity": "sha1-bsyh4PjBVtFBWXVZhI32RzCmu8E=",
      "dev": true,
      "requires": {
        "debug": "2.6.9",
        "depd": "~1.1.2",
        "destroy": "~1.0.4",
        "encodeurl": "~1.0.2",
        "escape-html": "~1.0.3",
        "etag": "~1.8.1",
        "fresh": "0.5.2",
        "http-errors": "~1.6.2",
        "mime": "1.4.1",
        "ms": "2.0.0",
        "on-finished": "~2.3.0",
        "range-parser": "~1.2.0",
        "statuses": "~1.4.0"
      },
      "dependencies": {
        "debug": {
          "version": "2.6.9",
          "resolved": "http://registry.npm.taobao.org/debug/download/debug-2.6.9.tgz",
          "integrity": "sha1-XRKFFd8TT/Mn6QpMk/Tgd6U2NB8=",
          "dev": true,
          "requires": {
            "ms": "2.0.0"
          }
        },
        "mime": {
          "version": "1.4.1",
          "resolved": "http://registry.npm.taobao.org/mime/download/mime-1.4.1.tgz",
          "integrity": "sha1-Eh+evEnjdm8xGnbh+hyAA8SwOqY=",
          "dev": true
        },
        "ms": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/ms/download/ms-2.0.0.tgz",
          "integrity": "sha1-VgiurfwAvmwpAd9fmGF4jeDVl8g=",
          "dev": true
        }
      }
    },
    "serialize-javascript": {
      "version": "1.5.0",
      "resolved": "http://registry.npm.taobao.org/serialize-javascript/download/serialize-javascript-1.5.0.tgz",
      "integrity": "sha1-GqM2FiyIqJDdrVOEuuvJOmVRYf4=",
      "dev": true
    },
    "serve-index": {
      "version": "1.9.1",
      "resolved": "http://registry.npm.taobao.org/serve-index/download/serve-index-1.9.1.tgz",
      "integrity": "sha1-03aNabHn2C5c4FD/9bRTvqEqkjk=",
      "dev": true,
      "requires": {
        "accepts": "~1.3.4",
        "batch": "0.6.1",
        "debug": "2.6.9",
        "escape-html": "~1.0.3",
        "http-errors": "~1.6.2",
        "mime-types": "~2.1.17",
        "parseurl": "~1.3.2"
      },
      "dependencies": {
        "debug": {
          "version": "2.6.9",
          "resolved": "http://registry.npm.taobao.org/debug/download/debug-2.6.9.tgz",
          "integrity": "sha1-XRKFFd8TT/Mn6QpMk/Tgd6U2NB8=",
          "dev": true,
          "requires": {
            "ms": "2.0.0"
          }
        },
        "ms": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/ms/download/ms-2.0.0.tgz",
          "integrity": "sha1-VgiurfwAvmwpAd9fmGF4jeDVl8g=",
          "dev": true
        }
      }
    },
    "serve-static": {
      "version": "1.13.2",
      "resolved": "http://registry.npm.taobao.org/serve-static/download/serve-static-1.13.2.tgz",
      "integrity": "sha1-CV6Ecv1bRiN9tQzkhqQ/S4bGzsE=",
      "dev": true,
      "requires": {
        "encodeurl": "~1.0.2",
        "escape-html": "~1.0.3",
        "parseurl": "~1.3.2",
        "send": "0.16.2"
      }
    },
    "set-blocking": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/set-blocking/download/set-blocking-2.0.0.tgz",
      "integrity": "sha1-BF+XgtARrppoA93TgrJDkrPYkPc=",
      "dev": true
    },
    "set-value": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/set-value/download/set-value-2.0.0.tgz",
      "integrity": "sha1-ca5KiPD+77v1LR6mBPP7MV67YnQ=",
      "dev": true,
      "requires": {
        "extend-shallow": "^2.0.1",
        "is-extendable": "^0.1.1",
        "is-plain-object": "^2.0.3",
        "split-string": "^3.0.1"
      },
      "dependencies": {
        "extend-shallow": {
          "version": "2.0.1",
          "resolved": "http://registry.npm.taobao.org/extend-shallow/download/extend-shallow-2.0.1.tgz",
          "integrity": "sha1-Ua99YUrZqfYQ6huvu5idaxxWiQ8=",
          "dev": true,
          "requires": {
            "is-extendable": "^0.1.0"
          }
        }
      }
    },
    "setimmediate": {
      "version": "1.0.5",
      "resolved": "http://registry.npm.taobao.org/setimmediate/download/setimmediate-1.0.5.tgz",
      "integrity": "sha1-KQy7Iy4waULX1+qbg3Mqt4VvgoU=",
      "dev": true
    },
    "setprototypeof": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/setprototypeof/download/setprototypeof-1.1.0.tgz",
      "integrity": "sha1-0L2FU2iHtv58DYGMuWLZ2RxU5lY=",
      "dev": true
    },
    "sha.js": {
      "version": "2.4.11",
      "resolved": "http://registry.npm.taobao.org/sha.js/download/sha.js-2.4.11.tgz",
      "integrity": "sha1-N6XPC4HsvGlD3hCbopYNGyZYSuc=",
      "dev": true,
      "requires": {
        "inherits": "^2.0.1",
        "safe-buffer": "^5.0.1"
      }
    },
    "shebang-command": {
      "version": "1.2.0",
      "resolved": "http://registry.npm.taobao.org/shebang-command/download/shebang-command-1.2.0.tgz",
      "integrity": "sha1-RKrGW2lbAzmJaMOfNj/uXer98eo=",
      "dev": true,
      "requires": {
        "shebang-regex": "^1.0.0"
      }
    },
    "shebang-regex": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/shebang-regex/download/shebang-regex-1.0.0.tgz",
      "integrity": "sha1-2kL0l0DAtC2yypcoVxyxkMmO/qM=",
      "dev": true
    },
    "shell-quote": {
      "version": "1.6.1",
      "resolved": "http://registry.npm.taobao.org/shell-quote/download/shell-quote-1.6.1.tgz",
      "integrity": "sha1-9HgZSczkAmlxJ0MOo7PFR29IF2c=",
      "dev": true,
      "requires": {
        "array-filter": "~0.0.0",
        "array-map": "~0.0.0",
        "array-reduce": "~0.0.0",
        "jsonify": "~0.0.0"
      }
    },
    "signal-exit": {
      "version": "3.0.2",
      "resolved": "http://registry.npm.taobao.org/signal-exit/download/signal-exit-3.0.2.tgz",
      "integrity": "sha1-tf3AjxKH6hF4Yo5BXiUTK3NkbG0=",
      "dev": true
    },
    "simple-swizzle": {
      "version": "0.2.2",
      "resolved": "http://registry.npm.taobao.org/simple-swizzle/download/simple-swizzle-0.2.2.tgz",
      "integrity": "sha1-pNprY1/8zMoz9w0Xy5JZLeleVXo=",
      "dev": true,
      "requires": {
        "is-arrayish": "^0.3.1"
      },
      "dependencies": {
        "is-arrayish": {
          "version": "0.3.2",
          "resolved": "http://registry.npm.taobao.org/is-arrayish/download/is-arrayish-0.3.2.tgz",
          "integrity": "sha1-RXSirlb3qyBolvtDHq7tBm/fjwM=",
          "dev": true
        }
      }
    },
    "slash": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/slash/download/slash-1.0.0.tgz",
      "integrity": "sha1-xB8vbDn8FtHNF61LXYlhFK5HDVU=",
      "dev": true
    },
    "slice-ansi": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/slice-ansi/download/slice-ansi-1.0.0.tgz",
      "integrity": "sha1-BE8aSdiEL/MHqta1Be0Xi9lQE00=",
      "dev": true,
      "requires": {
        "is-fullwidth-code-point": "^2.0.0"
      }
    },
    "snapdragon": {
      "version": "0.8.2",
      "resolved": "http://registry.npm.taobao.org/snapdragon/download/snapdragon-0.8.2.tgz",
      "integrity": "sha1-ZJIufFZbDhQgS6GqfWlkJ40lGC0=",
      "dev": true,
      "requires": {
        "base": "^0.11.1",
        "debug": "^2.2.0",
        "define-property": "^0.2.5",
        "extend-shallow": "^2.0.1",
        "map-cache": "^0.2.2",
        "source-map": "^0.5.6",
        "source-map-resolve": "^0.5.0",
        "use": "^3.1.0"
      },
      "dependencies": {
        "debug": {
          "version": "2.6.9",
          "resolved": "http://registry.npm.taobao.org/debug/download/debug-2.6.9.tgz",
          "integrity": "sha1-XRKFFd8TT/Mn6QpMk/Tgd6U2NB8=",
          "dev": true,
          "requires": {
            "ms": "2.0.0"
          }
        },
        "define-property": {
          "version": "0.2.5",
          "resolved": "http://registry.npm.taobao.org/define-property/download/define-property-0.2.5.tgz",
          "integrity": "sha1-w1se+RjsPJkPmlvFe+BKrOxcgRY=",
          "dev": true,
          "requires": {
            "is-descriptor": "^0.1.0"
          }
        },
        "extend-shallow": {
          "version": "2.0.1",
          "resolved": "http://registry.npm.taobao.org/extend-shallow/download/extend-shallow-2.0.1.tgz",
          "integrity": "sha1-Ua99YUrZqfYQ6huvu5idaxxWiQ8=",
          "dev": true,
          "requires": {
            "is-extendable": "^0.1.0"
          }
        },
        "ms": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/ms/download/ms-2.0.0.tgz",
          "integrity": "sha1-VgiurfwAvmwpAd9fmGF4jeDVl8g=",
          "dev": true
        }
      }
    },
    "snapdragon-node": {
      "version": "2.1.1",
      "resolved": "http://registry.npm.taobao.org/snapdragon-node/download/snapdragon-node-2.1.1.tgz",
      "integrity": "sha1-bBdfhv8UvbByRWPo88GwIaKGhTs=",
      "dev": true,
      "requires": {
        "define-property": "^1.0.0",
        "isobject": "^3.0.0",
        "snapdragon-util": "^3.0.1"
      },
      "dependencies": {
        "define-property": {
          "version": "1.0.0",
          "resolved": "http://registry.npm.taobao.org/define-property/download/define-property-1.0.0.tgz",
          "integrity": "sha1-dp66rz9KY6rTr56NMEybvnm/sOY=",
          "dev": true,
          "requires": {
            "is-descriptor": "^1.0.0"
          }
        },
        "is-accessor-descriptor": {
          "version": "1.0.0",
          "resolved": "http://registry.npm.taobao.org/is-accessor-descriptor/download/is-accessor-descriptor-1.0.0.tgz",
          "integrity": "sha1-FpwvbT3x+ZJhgHI2XJsOofaHhlY=",
          "dev": true,
          "requires": {
            "kind-of": "^6.0.0"
          }
        },
        "is-data-descriptor": {
          "version": "1.0.0",
          "resolved": "http://registry.npm.taobao.org/is-data-descriptor/download/is-data-descriptor-1.0.0.tgz",
          "integrity": "sha1-2Eh2Mh0Oet0DmQQGq7u9NrqSaMc=",
          "dev": true,
          "requires": {
            "kind-of": "^6.0.0"
          }
        },
        "is-descriptor": {
          "version": "1.0.2",
          "resolved": "http://registry.npm.taobao.org/is-descriptor/download/is-descriptor-1.0.2.tgz",
          "integrity": "sha1-OxWXRqZmBLBPjIFSS6NlxfFNhuw=",
          "dev": true,
          "requires": {
            "is-accessor-descriptor": "^1.0.0",
            "is-data-descriptor": "^1.0.0",
            "kind-of": "^6.0.2"
          }
        }
      }
    },
    "snapdragon-util": {
      "version": "3.0.1",
      "resolved": "http://registry.npm.taobao.org/snapdragon-util/download/snapdragon-util-3.0.1.tgz",
      "integrity": "sha1-+VZHlIbyrNeXAGk/b3uAXkWrVuI=",
      "dev": true,
      "requires": {
        "kind-of": "^3.2.0"
      },
      "dependencies": {
        "kind-of": {
          "version": "3.2.2",
          "resolved": "http://registry.npm.taobao.org/kind-of/download/kind-of-3.2.2.tgz",
          "integrity": "sha1-MeohpzS6ubuw8yRm2JOupR5KPGQ=",
          "dev": true,
          "requires": {
            "is-buffer": "^1.1.5"
          }
        }
      }
    },
    "sockjs": {
      "version": "0.3.19",
      "resolved": "http://registry.npm.taobao.org/sockjs/download/sockjs-0.3.19.tgz",
      "integrity": "sha1-2Xa76ACve9IK4IWY1YI5NQiZPA0=",
      "dev": true,
      "requires": {
        "faye-websocket": "^0.10.0",
        "uuid": "^3.0.1"
      }
    },
    "sockjs-client": {
      "version": "1.3.0",
      "resolved": "http://registry.npm.taobao.org/sockjs-client/download/sockjs-client-1.3.0.tgz",
      "integrity": "sha1-EvydbLZj2lc509xftuhofalcsXc=",
      "dev": true,
      "requires": {
        "debug": "^3.2.5",
        "eventsource": "^1.0.7",
        "faye-websocket": "~0.11.1",
        "inherits": "^2.0.3",
        "json3": "^3.3.2",
        "url-parse": "^1.4.3"
      },
      "dependencies": {
        "debug": {
          "version": "3.2.6",
          "resolved": "http://registry.npm.taobao.org/debug/download/debug-3.2.6.tgz",
          "integrity": "sha1-6D0X3hbYp++3cX7b5fsQE17uYps=",
          "dev": true,
          "requires": {
            "ms": "^2.1.1"
          }
        },
        "faye-websocket": {
          "version": "0.11.1",
          "resolved": "http://registry.npm.taobao.org/faye-websocket/download/faye-websocket-0.11.1.tgz",
          "integrity": "sha1-8O/hjE9W5PQK/H4Gxxn9XuYYjzg=",
          "dev": true,
          "requires": {
            "websocket-driver": ">=0.5.1"
          }
        }
      }
    },
    "source-list-map": {
      "version": "2.0.1",
      "resolved": "http://registry.npm.taobao.org/source-list-map/download/source-list-map-2.0.1.tgz",
      "integrity": "sha1-OZO9hzv8SEecyp6jpUeDXHwVSzQ=",
      "dev": true
    },
    "source-map": {
      "version": "0.5.7",
      "resolved": "http://registry.npm.taobao.org/source-map/download/source-map-0.5.7.tgz",
      "integrity": "sha1-igOdLRAh0i0eoUyA2OpGi6LvP8w=",
      "dev": true
    },
    "source-map-resolve": {
      "version": "0.5.2",
      "resolved": "http://registry.npm.taobao.org/source-map-resolve/download/source-map-resolve-0.5.2.tgz",
      "integrity": "sha1-cuLMNAlVQ+Q7LGKyxMENSpBU8lk=",
      "dev": true,
      "requires": {
        "atob": "^2.1.1",
        "decode-uri-component": "^0.2.0",
        "resolve-url": "^0.2.1",
        "source-map-url": "^0.4.0",
        "urix": "^0.1.0"
      }
    },
    "source-map-support": {
      "version": "0.5.9",
      "resolved": "http://registry.npm.taobao.org/source-map-support/download/source-map-support-0.5.9.tgz",
      "integrity": "sha1-QbyVOyU0Jn6i1gW8z6e/oxEc7V8=",
      "dev": true,
      "requires": {
        "buffer-from": "^1.0.0",
        "source-map": "^0.6.0"
      },
      "dependencies": {
        "source-map": {
          "version": "0.6.1",
          "resolved": "http://registry.npm.taobao.org/source-map/download/source-map-0.6.1.tgz",
          "integrity": "sha1-dHIq8y6WFOnCh6jQu95IteLxomM=",
          "dev": true
        }
      }
    },
    "source-map-url": {
      "version": "0.4.0",
      "resolved": "http://registry.npm.taobao.org/source-map-url/download/source-map-url-0.4.0.tgz",
      "integrity": "sha1-PpNdfd1zYxuXZZlW1VEo6HtQhKM=",
      "dev": true
    },
    "spdx-correct": {
      "version": "3.1.0",
      "resolved": "http://registry.npm.taobao.org/spdx-correct/download/spdx-correct-3.1.0.tgz",
      "integrity": "sha1-+4PlBERSaPFUsHTiGMh8ADzTHfQ=",
      "dev": true,
      "requires": {
        "spdx-expression-parse": "^3.0.0",
        "spdx-license-ids": "^3.0.0"
      }
    },
    "spdx-exceptions": {
      "version": "2.2.0",
      "resolved": "http://registry.npm.taobao.org/spdx-exceptions/download/spdx-exceptions-2.2.0.tgz",
      "integrity": "sha1-LqRQrudPKom/uUUZwH/Nb0EyKXc=",
      "dev": true
    },
    "spdx-expression-parse": {
      "version": "3.0.0",
      "resolved": "http://registry.npm.taobao.org/spdx-expression-parse/download/spdx-expression-parse-3.0.0.tgz",
      "integrity": "sha1-meEZt6XaAOBUkcn6M4t5BII7QdA=",
      "dev": true,
      "requires": {
        "spdx-exceptions": "^2.1.0",
        "spdx-license-ids": "^3.0.0"
      }
    },
    "spdx-license-ids": {
      "version": "3.0.2",
      "resolved": "http://registry.npm.taobao.org/spdx-license-ids/download/spdx-license-ids-3.0.2.tgz",
      "integrity": "sha1-pZ78CXhMKlutoTz+r1x13SFARNI=",
      "dev": true
    },
    "spdy": {
      "version": "3.4.7",
      "resolved": "http://registry.npm.taobao.org/spdy/download/spdy-3.4.7.tgz",
      "integrity": "sha1-Qv9B7OXMD5mjpsKKq7c/XDsDrLw=",
      "dev": true,
      "requires": {
        "debug": "^2.6.8",
        "handle-thing": "^1.2.5",
        "http-deceiver": "^1.2.7",
        "safe-buffer": "^5.0.1",
        "select-hose": "^2.0.0",
        "spdy-transport": "^2.0.18"
      },
      "dependencies": {
        "debug": {
          "version": "2.6.9",
          "resolved": "http://registry.npm.taobao.org/debug/download/debug-2.6.9.tgz",
          "integrity": "sha1-XRKFFd8TT/Mn6QpMk/Tgd6U2NB8=",
          "dev": true,
          "requires": {
            "ms": "2.0.0"
          }
        },
        "ms": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/ms/download/ms-2.0.0.tgz",
          "integrity": "sha1-VgiurfwAvmwpAd9fmGF4jeDVl8g=",
          "dev": true
        }
      }
    },
    "spdy-transport": {
      "version": "2.1.1",
      "resolved": "http://registry.npm.taobao.org/spdy-transport/download/spdy-transport-2.1.1.tgz",
      "integrity": "sha1-xUgV1zhYqt0GzmMAHn0l+mRBYjs=",
      "dev": true,
      "requires": {
        "debug": "^2.6.8",
        "detect-node": "^2.0.3",
        "hpack.js": "^2.1.6",
        "obuf": "^1.1.1",
        "readable-stream": "^2.2.9",
        "safe-buffer": "^5.0.1",
        "wbuf": "^1.7.2"
      },
      "dependencies": {
        "debug": {
          "version": "2.6.9",
          "resolved": "http://registry.npm.taobao.org/debug/download/debug-2.6.9.tgz",
          "integrity": "sha1-XRKFFd8TT/Mn6QpMk/Tgd6U2NB8=",
          "dev": true,
          "requires": {
            "ms": "2.0.0"
          }
        },
        "ms": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/ms/download/ms-2.0.0.tgz",
          "integrity": "sha1-VgiurfwAvmwpAd9fmGF4jeDVl8g=",
          "dev": true
        }
      }
    },
    "split-string": {
      "version": "3.1.0",
      "resolved": "http://registry.npm.taobao.org/split-string/download/split-string-3.1.0.tgz",
      "integrity": "sha1-fLCd2jqGWFcFxks5pkZgOGguj+I=",
      "dev": true,
      "requires": {
        "extend-shallow": "^3.0.0"
      }
    },
    "sprintf-js": {
      "version": "1.0.3",
      "resolved": "http://registry.npm.taobao.org/sprintf-js/download/sprintf-js-1.0.3.tgz",
      "integrity": "sha1-BOaSb2YolTVPPdAVIDYzuFcpfiw=",
      "dev": true
    },
    "sshpk": {
      "version": "1.15.2",
      "resolved": "http://registry.npm.taobao.org/sshpk/download/sshpk-1.15.2.tgz",
      "integrity": "sha1-yUbWvZsaOdDoY1dj9SQtbtbctik=",
      "dev": true,
      "requires": {
        "asn1": "~0.2.3",
        "assert-plus": "^1.0.0",
        "bcrypt-pbkdf": "^1.0.0",
        "dashdash": "^1.12.0",
        "ecc-jsbn": "~0.1.1",
        "getpass": "^0.1.1",
        "jsbn": "~0.1.0",
        "safer-buffer": "^2.0.2",
        "tweetnacl": "~0.14.0"
      }
    },
    "ssri": {
      "version": "6.0.1",
      "resolved": "http://registry.npm.taobao.org/ssri/download/ssri-6.0.1.tgz",
      "integrity": "sha1-KjxBso3UW2K2Nnbst0ABJlrp7dg=",
      "dev": true,
      "requires": {
        "figgy-pudding": "^3.5.1"
      }
    },
    "stable": {
      "version": "0.1.8",
      "resolved": "http://registry.npm.taobao.org/stable/download/stable-0.1.8.tgz",
      "integrity": "sha1-g26zyDgv4pNv6vVEYxAXzn1Ho88=",
      "dev": true
    },
    "stackframe": {
      "version": "1.0.4",
      "resolved": "http://registry.npm.taobao.org/stackframe/download/stackframe-1.0.4.tgz",
      "integrity": "sha1-NXskqZL5Qny6a1RdlqFO0svKGHs=",
      "dev": true
    },
    "static-extend": {
      "version": "0.1.2",
      "resolved": "http://registry.npm.taobao.org/static-extend/download/static-extend-0.1.2.tgz",
      "integrity": "sha1-YICcOcv/VTNyJv1eC1IPNB8ftcY=",
      "dev": true,
      "requires": {
        "define-property": "^0.2.5",
        "object-copy": "^0.1.0"
      },
      "dependencies": {
        "define-property": {
          "version": "0.2.5",
          "resolved": "http://registry.npm.taobao.org/define-property/download/define-property-0.2.5.tgz",
          "integrity": "sha1-w1se+RjsPJkPmlvFe+BKrOxcgRY=",
          "dev": true,
          "requires": {
            "is-descriptor": "^0.1.0"
          }
        }
      }
    },
    "statuses": {
      "version": "1.4.0",
      "resolved": "http://registry.npm.taobao.org/statuses/download/statuses-1.4.0.tgz",
      "integrity": "sha1-u3PURtonlhBu/MG2AaJT1sRr0Ic=",
      "dev": true
    },
    "stealthy-require": {
      "version": "1.1.1",
      "resolved": "http://registry.npm.taobao.org/stealthy-require/download/stealthy-require-1.1.1.tgz",
      "integrity": "sha1-NbCYdbT/SfJqd35QmzCQoyJr8ks=",
      "dev": true
    },
    "stream-browserify": {
      "version": "2.0.1",
      "resolved": "http://registry.npm.taobao.org/stream-browserify/download/stream-browserify-2.0.1.tgz",
      "integrity": "sha1-ZiZu5fm9uZQKTkUUyvtDu3Hlyds=",
      "dev": true,
      "requires": {
        "inherits": "~2.0.1",
        "readable-stream": "^2.0.2"
      }
    },
    "stream-each": {
      "version": "1.2.3",
      "resolved": "http://registry.npm.taobao.org/stream-each/download/stream-each-1.2.3.tgz",
      "integrity": "sha1-6+J6DDibBPvMIzZClS4Qcxr6m64=",
      "dev": true,
      "requires": {
        "end-of-stream": "^1.1.0",
        "stream-shift": "^1.0.0"
      }
    },
    "stream-http": {
      "version": "2.8.3",
      "resolved": "http://registry.npm.taobao.org/stream-http/download/stream-http-2.8.3.tgz",
      "integrity": "sha1-stJCRpKIpaJ+xP6JM6z2I95lFPw=",
      "dev": true,
      "requires": {
        "builtin-status-codes": "^3.0.0",
        "inherits": "^2.0.1",
        "readable-stream": "^2.3.6",
        "to-arraybuffer": "^1.0.0",
        "xtend": "^4.0.0"
      }
    },
    "stream-shift": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/stream-shift/download/stream-shift-1.0.0.tgz",
      "integrity": "sha1-1cdSgl5TZ+eG944Y5EXqIjoVWVI=",
      "dev": true
    },
    "string-width": {
      "version": "2.1.1",
      "resolved": "http://registry.npm.taobao.org/string-width/download/string-width-2.1.1.tgz",
      "integrity": "sha1-q5Pyeo3BPSjKyBXEYhQ6bZASrp4=",
      "dev": true,
      "requires": {
        "is-fullwidth-code-point": "^2.0.0",
        "strip-ansi": "^4.0.0"
      }
    },
    "string.prototype.padend": {
      "version": "3.0.0",
      "resolved": "http://registry.npm.taobao.org/string.prototype.padend/download/string.prototype.padend-3.0.0.tgz",
      "integrity": "sha1-86rvfBcZ8XDF6rHDK/eA2W4h8vA=",
      "dev": true,
      "requires": {
        "define-properties": "^1.1.2",
        "es-abstract": "^1.4.3",
        "function-bind": "^1.0.2"
      }
    },
    "string.prototype.padstart": {
      "version": "3.0.0",
      "resolved": "http://registry.npm.taobao.org/string.prototype.padstart/download/string.prototype.padstart-3.0.0.tgz",
      "integrity": "sha1-W8+tOfRkm7LQMSkuGbzwtRDUskI=",
      "dev": true,
      "requires": {
        "define-properties": "^1.1.2",
        "es-abstract": "^1.4.3",
        "function-bind": "^1.0.2"
      }
    },
    "string_decoder": {
      "version": "1.1.1",
      "resolved": "http://registry.npm.taobao.org/string_decoder/download/string_decoder-1.1.1.tgz",
      "integrity": "sha1-nPFhG6YmhdcDCunkujQUnDrwP8g=",
      "dev": true,
      "requires": {
        "safe-buffer": "~5.1.0"
      }
    },
    "strip-ansi": {
      "version": "4.0.0",
      "resolved": "http://registry.npm.taobao.org/strip-ansi/download/strip-ansi-4.0.0.tgz",
      "integrity": "sha1-qEeQIusaw2iocTibY1JixQXuNo8=",
      "dev": true,
      "requires": {
        "ansi-regex": "^3.0.0"
      }
    },
    "strip-eof": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/strip-eof/download/strip-eof-1.0.0.tgz",
      "integrity": "sha1-u0P/VZim6wXYm1n80SnJgzE2Br8=",
      "dev": true
    },
    "strip-indent": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/strip-indent/download/strip-indent-2.0.0.tgz",
      "integrity": "sha1-XvjbKV0B5u1sv3qrlpmNeCJSe2g=",
      "dev": true
    },
    "strip-json-comments": {
      "version": "2.0.1",
      "resolved": "http://registry.npm.taobao.org/strip-json-comments/download/strip-json-comments-2.0.1.tgz",
      "integrity": "sha1-PFMZQukIwml8DsNEhYwobHygpgo=",
      "dev": true
    },
    "stylehacks": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/stylehacks/download/stylehacks-4.0.1.tgz",
      "integrity": "sha1-MYZZXQR6sN+BPSE+Uci5TguQEPI=",
      "dev": true,
      "requires": {
        "browserslist": "^4.0.0",
        "postcss": "^7.0.0",
        "postcss-selector-parser": "^3.0.0"
      },
      "dependencies": {
        "postcss-selector-parser": {
          "version": "3.1.1",
          "resolved": "http://registry.npm.taobao.org/postcss-selector-parser/download/postcss-selector-parser-3.1.1.tgz",
          "integrity": "sha1-T4dfSvsMllc9XPTXQBGu4lCn6GU=",
          "dev": true,
          "requires": {
            "dot-prop": "^4.1.1",
            "indexes-of": "^1.0.1",
            "uniq": "^1.0.1"
          }
        }
      }
    },
    "supports-color": {
      "version": "5.5.0",
      "resolved": "http://registry.npm.taobao.org/supports-color/download/supports-color-5.5.0.tgz",
      "integrity": "sha1-4uaaRKyHcveKHsCzW2id9lMO/I8=",
      "dev": true,
      "requires": {
        "has-flag": "^3.0.0"
      }
    },
    "svgo": {
      "version": "1.1.1",
      "resolved": "http://registry.npm.taobao.org/svgo/download/svgo-1.1.1.tgz",
      "integrity": "sha1-EjhLAzNbzs2Fz6X04zdf7WccuYU=",
      "dev": true,
      "requires": {
        "coa": "~2.0.1",
        "colors": "~1.1.2",
        "css-select": "^2.0.0",
        "css-select-base-adapter": "~0.1.0",
        "css-tree": "1.0.0-alpha.28",
        "css-url-regex": "^1.1.0",
        "csso": "^3.5.0",
        "js-yaml": "^3.12.0",
        "mkdirp": "~0.5.1",
        "object.values": "^1.0.4",
        "sax": "~1.2.4",
        "stable": "~0.1.6",
        "unquote": "~1.1.1",
        "util.promisify": "~1.0.0"
      }
    },
    "table": {
      "version": "4.0.2",
      "resolved": "http://registry.npm.taobao.org/table/download/table-4.0.2.tgz",
      "integrity": "sha1-ozRHN1OR52atNNNIbm4q7chNLjY=",
      "dev": true,
      "requires": {
        "ajv": "^5.2.3",
        "ajv-keywords": "^2.1.0",
        "chalk": "^2.1.0",
        "lodash": "^4.17.4",
        "slice-ansi": "1.0.0",
        "string-width": "^2.1.1"
      },
      "dependencies": {
        "ajv": {
          "version": "5.5.2",
          "resolved": "http://registry.npm.taobao.org/ajv/download/ajv-5.5.2.tgz",
          "integrity": "sha1-c7Xuyj+rZT49P5Qis0GtQiBdyWU=",
          "dev": true,
          "requires": {
            "co": "^4.6.0",
            "fast-deep-equal": "^1.0.0",
            "fast-json-stable-stringify": "^2.0.0",
            "json-schema-traverse": "^0.3.0"
          }
        },
        "fast-deep-equal": {
          "version": "1.1.0",
          "resolved": "http://registry.npm.taobao.org/fast-deep-equal/download/fast-deep-equal-1.1.0.tgz",
          "integrity": "sha1-wFNHeBfIa1HaqFPIHgWbcz0CNhQ=",
          "dev": true
        },
        "json-schema-traverse": {
          "version": "0.3.1",
          "resolved": "http://registry.npm.taobao.org/json-schema-traverse/download/json-schema-traverse-0.3.1.tgz",
          "integrity": "sha1-NJptRMU6Ud6JtAgFxdXlm0F9M0A=",
          "dev": true
        }
      }
    },
    "tapable": {
      "version": "1.1.1",
      "resolved": "http://registry.npm.taobao.org/tapable/download/tapable-1.1.1.tgz",
      "integrity": "sha1-TSl5I8WnKkI2DeKrUtrfquwAAY4=",
      "dev": true
    },
    "terser": {
      "version": "3.11.0",
      "resolved": "http://registry.npm.taobao.org/terser/download/terser-3.11.0.tgz",
      "integrity": "sha1-YHgok+H01niKzGljUfQGNtDjevA=",
      "dev": true,
      "requires": {
        "commander": "~2.17.1",
        "source-map": "~0.6.1",
        "source-map-support": "~0.5.6"
      },
      "dependencies": {
        "source-map": {
          "version": "0.6.1",
          "resolved": "http://registry.npm.taobao.org/source-map/download/source-map-0.6.1.tgz",
          "integrity": "sha1-dHIq8y6WFOnCh6jQu95IteLxomM=",
          "dev": true
        }
      }
    },
    "terser-webpack-plugin": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/terser-webpack-plugin/download/terser-webpack-plugin-1.1.0.tgz",
      "integrity": "sha1-z3wloe7iW/Eh9KWHu54ATj+A5Sg=",
      "dev": true,
      "requires": {
        "cacache": "^11.0.2",
        "find-cache-dir": "^2.0.0",
        "schema-utils": "^1.0.0",
        "serialize-javascript": "^1.4.0",
        "source-map": "^0.6.1",
        "terser": "^3.8.1",
        "webpack-sources": "^1.1.0",
        "worker-farm": "^1.5.2"
      },
      "dependencies": {
        "ajv-keywords": {
          "version": "3.2.0",
          "resolved": "http://registry.npm.taobao.org/ajv-keywords/download/ajv-keywords-3.2.0.tgz",
          "integrity": "sha1-6GuBnGAs+IIa1jdBNpjx3sAhhHo=",
          "dev": true
        },
        "cacache": {
          "version": "11.3.1",
          "resolved": "http://registry.npm.taobao.org/cacache/download/cacache-11.3.1.tgz",
          "integrity": "sha1-0J0l9sSsp6bTBdFBrjMmE6odUV8=",
          "dev": true,
          "requires": {
            "bluebird": "^3.5.1",
            "chownr": "^1.0.1",
            "figgy-pudding": "^3.1.0",
            "glob": "^7.1.2",
            "graceful-fs": "^4.1.11",
            "lru-cache": "^4.1.3",
            "mississippi": "^3.0.0",
            "mkdirp": "^0.5.1",
            "move-concurrently": "^1.0.1",
            "promise-inflight": "^1.0.1",
            "rimraf": "^2.6.2",
            "ssri": "^6.0.0",
            "unique-filename": "^1.1.0",
            "y18n": "^4.0.0"
          }
        },
        "find-cache-dir": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/find-cache-dir/download/find-cache-dir-2.0.0.tgz",
          "integrity": "sha1-TB+u1Z9FGEUw+51/oSOk0EqYRy0=",
          "dev": true,
          "requires": {
            "commondir": "^1.0.1",
            "make-dir": "^1.0.0",
            "pkg-dir": "^3.0.0"
          }
        },
        "find-up": {
          "version": "3.0.0",
          "resolved": "http://registry.npm.taobao.org/find-up/download/find-up-3.0.0.tgz",
          "integrity": "sha1-SRafHXmTQwZG2mHsxa41XCHJe3M=",
          "dev": true,
          "requires": {
            "locate-path": "^3.0.0"
          }
        },
        "locate-path": {
          "version": "3.0.0",
          "resolved": "http://registry.npm.taobao.org/locate-path/download/locate-path-3.0.0.tgz",
          "integrity": "sha1-2+w7OrdZdYBxtY/ln8QYca8hQA4=",
          "dev": true,
          "requires": {
            "p-locate": "^3.0.0",
            "path-exists": "^3.0.0"
          }
        },
        "mississippi": {
          "version": "3.0.0",
          "resolved": "http://registry.npm.taobao.org/mississippi/download/mississippi-3.0.0.tgz",
          "integrity": "sha1-6goykfl+C16HdrNj1fChLZTGcCI=",
          "dev": true,
          "requires": {
            "concat-stream": "^1.5.0",
            "duplexify": "^3.4.2",
            "end-of-stream": "^1.1.0",
            "flush-write-stream": "^1.0.0",
            "from2": "^2.1.0",
            "parallel-transform": "^1.1.0",
            "pump": "^3.0.0",
            "pumpify": "^1.3.3",
            "stream-each": "^1.1.0",
            "through2": "^2.0.0"
          }
        },
        "p-limit": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/p-limit/download/p-limit-2.0.0.tgz",
          "integrity": "sha1-5iTtVO6MRgp3izyfNnBJb/ileuw=",
          "dev": true,
          "requires": {
            "p-try": "^2.0.0"
          }
        },
        "p-locate": {
          "version": "3.0.0",
          "resolved": "http://registry.npm.taobao.org/p-locate/download/p-locate-3.0.0.tgz",
          "integrity": "sha1-Mi1poFwCZLJZl9n0DNiokasAZKQ=",
          "dev": true,
          "requires": {
            "p-limit": "^2.0.0"
          }
        },
        "p-try": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/p-try/download/p-try-2.0.0.tgz",
          "integrity": "sha1-hQgLuHxkaI+keZb+j3376CEXYLE=",
          "dev": true
        },
        "pkg-dir": {
          "version": "3.0.0",
          "resolved": "http://registry.npm.taobao.org/pkg-dir/download/pkg-dir-3.0.0.tgz",
          "integrity": "sha1-J0kCDyOe2ZCIGx9xIQ1R62UjvqM=",
          "dev": true,
          "requires": {
            "find-up": "^3.0.0"
          }
        },
        "schema-utils": {
          "version": "1.0.0",
          "resolved": "http://registry.npm.taobao.org/schema-utils/download/schema-utils-1.0.0.tgz",
          "integrity": "sha1-C3mpMgTXtgDUsoUNH2bCo0lRx3A=",
          "dev": true,
          "requires": {
            "ajv": "^6.1.0",
            "ajv-errors": "^1.0.0",
            "ajv-keywords": "^3.1.0"
          }
        },
        "source-map": {
          "version": "0.6.1",
          "resolved": "http://registry.npm.taobao.org/source-map/download/source-map-0.6.1.tgz",
          "integrity": "sha1-dHIq8y6WFOnCh6jQu95IteLxomM=",
          "dev": true
        }
      }
    },
    "text-table": {
      "version": "0.2.0",
      "resolved": "http://registry.npm.taobao.org/text-table/download/text-table-0.2.0.tgz",
      "integrity": "sha1-f17oI66AUgfACvLfSoTsP8+lcLQ=",
      "dev": true
    },
    "thread-loader": {
      "version": "1.2.0",
      "resolved": "http://registry.npm.taobao.org/thread-loader/download/thread-loader-1.2.0.tgz",
      "integrity": "sha1-Nd7bI88pSvu85sRcEzm5UO0X56Q=",
      "dev": true,
      "requires": {
        "async": "^2.3.0",
        "loader-runner": "^2.3.0",
        "loader-utils": "^1.1.0"
      },
      "dependencies": {
        "async": {
          "version": "2.6.1",
          "resolved": "http://registry.npm.taobao.org/async/download/async-2.6.1.tgz",
          "integrity": "sha1-skWiPKcZMAROxT+kaqAKPofGphA=",
          "dev": true,
          "requires": {
            "lodash": "^4.17.10"
          }
        }
      }
    },
    "through": {
      "version": "2.3.8",
      "resolved": "http://registry.npm.taobao.org/through/download/through-2.3.8.tgz",
      "integrity": "sha1-DdTJ/6q8NXlgsbckEV1+Doai4fU=",
      "dev": true
    },
    "through2": {
      "version": "2.0.5",
      "resolved": "http://registry.npm.taobao.org/through2/download/through2-2.0.5.tgz",
      "integrity": "sha1-AcHjnrMdB8t9A6lqcIIyYLIxMs0=",
      "dev": true,
      "requires": {
        "readable-stream": "~2.3.6",
        "xtend": "~4.0.1"
      }
    },
    "thunky": {
      "version": "1.0.3",
      "resolved": "http://registry.npm.taobao.org/thunky/download/thunky-1.0.3.tgz",
      "integrity": "sha1-9d9zJFNAewkZHa5z4qjMc/OBqCY=",
      "dev": true
    },
    "timers-browserify": {
      "version": "2.0.10",
      "resolved": "http://registry.npm.taobao.org/timers-browserify/download/timers-browserify-2.0.10.tgz",
      "integrity": "sha1-HSjj0qrfHVpZlsTp+VYBzQU0gK4=",
      "dev": true,
      "requires": {
        "setimmediate": "^1.0.4"
      }
    },
    "timsort": {
      "version": "0.3.0",
      "resolved": "http://registry.npm.taobao.org/timsort/download/timsort-0.3.0.tgz",
      "integrity": "sha1-QFQRqOfmM5/mTbmiNN4R3DHgK9Q=",
      "dev": true
    },
    "tmp": {
      "version": "0.0.33",
      "resolved": "http://registry.npm.taobao.org/tmp/download/tmp-0.0.33.tgz",
      "integrity": "sha1-bTQzWIl2jSGyvNoKonfO07G/rfk=",
      "dev": true,
      "requires": {
        "os-tmpdir": "~1.0.2"
      }
    },
    "to-arraybuffer": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/to-arraybuffer/download/to-arraybuffer-1.0.1.tgz",
      "integrity": "sha1-fSKbH8xjfkZsoIEYCDanqr/4P0M=",
      "dev": true
    },
    "to-fast-properties": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/to-fast-properties/download/to-fast-properties-2.0.0.tgz",
      "integrity": "sha1-3F5pjL0HkmW8c+A3doGk5Og/YW4=",
      "dev": true
    },
    "to-object-path": {
      "version": "0.3.0",
      "resolved": "http://registry.npm.taobao.org/to-object-path/download/to-object-path-0.3.0.tgz",
      "integrity": "sha1-KXWIt7Dn4KwI4E5nL4XB9JmeF68=",
      "dev": true,
      "requires": {
        "kind-of": "^3.0.2"
      },
      "dependencies": {
        "kind-of": {
          "version": "3.2.2",
          "resolved": "http://registry.npm.taobao.org/kind-of/download/kind-of-3.2.2.tgz",
          "integrity": "sha1-MeohpzS6ubuw8yRm2JOupR5KPGQ=",
          "dev": true,
          "requires": {
            "is-buffer": "^1.1.5"
          }
        }
      }
    },
    "to-regex": {
      "version": "3.0.2",
      "resolved": "http://registry.npm.taobao.org/to-regex/download/to-regex-3.0.2.tgz",
      "integrity": "sha1-E8/dmzNlUvMLUfM6iuG0Knp1mc4=",
      "dev": true,
      "requires": {
        "define-property": "^2.0.2",
        "extend-shallow": "^3.0.2",
        "regex-not": "^1.0.2",
        "safe-regex": "^1.1.0"
      }
    },
    "to-regex-range": {
      "version": "2.1.1",
      "resolved": "http://registry.npm.taobao.org/to-regex-range/download/to-regex-range-2.1.1.tgz",
      "integrity": "sha1-fIDBe53+vlmeJzZ+DU3VWQFB2zg=",
      "dev": true,
      "requires": {
        "is-number": "^3.0.0",
        "repeat-string": "^1.6.1"
      }
    },
    "topo": {
      "version": "3.0.3",
      "resolved": "http://registry.npm.taobao.org/topo/download/topo-3.0.3.tgz",
      "integrity": "sha1-1aZ/suaTB+vusIQC7Coqb1962Vw=",
      "dev": true,
      "requires": {
        "hoek": "6.x.x"
      },
      "dependencies": {
        "hoek": {
          "version": "6.1.2",
          "resolved": "http://registry.npm.taobao.org/hoek/download/hoek-6.1.2.tgz",
          "integrity": "sha1-mebQcFYYOd507kJ7YapHa9a939Y=",
          "dev": true
        }
      }
    },
    "toposort": {
      "version": "1.0.7",
      "resolved": "http://registry.npm.taobao.org/toposort/download/toposort-1.0.7.tgz",
      "integrity": "sha1-LmhELZ9k7HILjMieZEOsbKqVACk=",
      "dev": true
    },
    "tough-cookie": {
      "version": "2.4.3",
      "resolved": "http://registry.npm.taobao.org/tough-cookie/download/tough-cookie-2.4.3.tgz",
      "integrity": "sha1-U/Nto/R3g7CSWvoG/587FlKA94E=",
      "dev": true,
      "requires": {
        "psl": "^1.1.24",
        "punycode": "^1.4.1"
      },
      "dependencies": {
        "punycode": {
          "version": "1.4.1",
          "resolved": "http://registry.npm.taobao.org/punycode/download/punycode-1.4.1.tgz",
          "integrity": "sha1-wNWmOycYgArY4esPpSachN1BhF4=",
          "dev": true
        }
      }
    },
    "trim-right": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/trim-right/download/trim-right-1.0.1.tgz",
      "integrity": "sha1-yy4SAwZ+DI3h9hQJS5/kVwTqYAM=",
      "dev": true
    },
    "tryer": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/tryer/download/tryer-1.0.1.tgz",
      "integrity": "sha1-8shUBoALmw90yfdGW4HqrSQSUvg=",
      "dev": true
    },
    "tslib": {
      "version": "1.9.3",
      "resolved": "http://registry.npm.taobao.org/tslib/download/tslib-1.9.3.tgz",
      "integrity": "sha1-1+TdeSRdhUKMTX5IIqeZF5VMooY=",
      "dev": true
    },
    "tty-browserify": {
      "version": "0.0.0",
      "resolved": "http://registry.npm.taobao.org/tty-browserify/download/tty-browserify-0.0.0.tgz",
      "integrity": "sha1-oVe6QC2iTpv5V/mqadUk7tQpAaY=",
      "dev": true
    },
    "tunnel-agent": {
      "version": "0.6.0",
      "resolved": "http://registry.npm.taobao.org/tunnel-agent/download/tunnel-agent-0.6.0.tgz",
      "integrity": "sha1-J6XeoGs2sEoKmWZ3SykIaPD8QP0=",
      "dev": true,
      "requires": {
        "safe-buffer": "^5.0.1"
      }
    },
    "tweetnacl": {
      "version": "0.14.5",
      "resolved": "http://registry.npm.taobao.org/tweetnacl/download/tweetnacl-0.14.5.tgz",
      "integrity": "sha1-WuaBd/GS1EViadEIr6k/+HQ/T2Q=",
      "dev": true
    },
    "type-check": {
      "version": "0.3.2",
      "resolved": "http://registry.npm.taobao.org/type-check/download/type-check-0.3.2.tgz",
      "integrity": "sha1-WITKtRLPHTVeP7eE8wgEsrUg23I=",
      "dev": true,
      "requires": {
        "prelude-ls": "~1.1.2"
      }
    },
    "type-is": {
      "version": "1.6.16",
      "resolved": "http://registry.npm.taobao.org/type-is/download/type-is-1.6.16.tgz",
      "integrity": "sha1-+JzjQVQcZysl7nrjxz3uOyvlAZQ=",
      "dev": true,
      "requires": {
        "media-typer": "0.3.0",
        "mime-types": "~2.1.18"
      }
    },
    "typedarray": {
      "version": "0.0.6",
      "resolved": "http://registry.npm.taobao.org/typedarray/download/typedarray-0.0.6.tgz",
      "integrity": "sha1-hnrHTjhkGHsdPUfZlqeOxciDB3c=",
      "dev": true
    },
    "uglify-js": {
      "version": "3.4.9",
      "resolved": "http://registry.npm.taobao.org/uglify-js/download/uglify-js-3.4.9.tgz",
      "integrity": "sha1-rwLxgMEgfXZDLkc+0koo9KeCuuM=",
      "dev": true,
      "requires": {
        "commander": "~2.17.1",
        "source-map": "~0.6.1"
      },
      "dependencies": {
        "source-map": {
          "version": "0.6.1",
          "resolved": "http://registry.npm.taobao.org/source-map/download/source-map-0.6.1.tgz",
          "integrity": "sha1-dHIq8y6WFOnCh6jQu95IteLxomM=",
          "dev": true
        }
      }
    },
    "unicode-canonical-property-names-ecmascript": {
      "version": "1.0.4",
      "resolved": "http://registry.npm.taobao.org/unicode-canonical-property-names-ecmascript/download/unicode-canonical-property-names-ecmascript-1.0.4.tgz",
      "integrity": "sha1-JhmADEyCWADv3YNDr33Zkzy+KBg=",
      "dev": true
    },
    "unicode-match-property-ecmascript": {
      "version": "1.0.4",
      "resolved": "http://registry.npm.taobao.org/unicode-match-property-ecmascript/download/unicode-match-property-ecmascript-1.0.4.tgz",
      "integrity": "sha1-jtKjJWmWG86SJ9Cc0/+7j+1fAgw=",
      "dev": true,
      "requires": {
        "unicode-canonical-property-names-ecmascript": "^1.0.4",
        "unicode-property-aliases-ecmascript": "^1.0.4"
      }
    },
    "unicode-match-property-value-ecmascript": {
      "version": "1.0.2",
      "resolved": "http://registry.npm.taobao.org/unicode-match-property-value-ecmascript/download/unicode-match-property-value-ecmascript-1.0.2.tgz",
      "integrity": "sha1-nx3HaSbWzPRSMQVk/YNKzgWWY9Q=",
      "dev": true
    },
    "unicode-property-aliases-ecmascript": {
      "version": "1.0.4",
      "resolved": "http://registry.npm.taobao.org/unicode-property-aliases-ecmascript/download/unicode-property-aliases-ecmascript-1.0.4.tgz",
      "integrity": "sha1-WlM/MbQxfqdvF9gH+g0RZUYRHdA=",
      "dev": true
    },
    "union-value": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/union-value/download/union-value-1.0.0.tgz",
      "integrity": "sha1-XHHDTLW61dzr4+oM0IIHulqhrqQ=",
      "dev": true,
      "requires": {
        "arr-union": "^3.1.0",
        "get-value": "^2.0.6",
        "is-extendable": "^0.1.1",
        "set-value": "^0.4.3"
      },
      "dependencies": {
        "extend-shallow": {
          "version": "2.0.1",
          "resolved": "http://registry.npm.taobao.org/extend-shallow/download/extend-shallow-2.0.1.tgz",
          "integrity": "sha1-Ua99YUrZqfYQ6huvu5idaxxWiQ8=",
          "dev": true,
          "requires": {
            "is-extendable": "^0.1.0"
          }
        },
        "set-value": {
          "version": "0.4.3",
          "resolved": "http://registry.npm.taobao.org/set-value/download/set-value-0.4.3.tgz",
          "integrity": "sha1-fbCPnT0i3H945Trzw79GZuzfzPE=",
          "dev": true,
          "requires": {
            "extend-shallow": "^2.0.1",
            "is-extendable": "^0.1.1",
            "is-plain-object": "^2.0.1",
            "to-object-path": "^0.3.0"
          }
        }
      }
    },
    "uniq": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/uniq/download/uniq-1.0.1.tgz",
      "integrity": "sha1-sxxa6CVIRKOoKBVBzisEuGWnNP8=",
      "dev": true
    },
    "uniqs": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/uniqs/download/uniqs-2.0.0.tgz",
      "integrity": "sha1-/+3ks2slKQaW5uFl1KWe25mOawI=",
      "dev": true
    },
    "unique-filename": {
      "version": "1.1.1",
      "resolved": "http://registry.npm.taobao.org/unique-filename/download/unique-filename-1.1.1.tgz",
      "integrity": "sha1-HWl2k2mtoFgxA6HmrodoG1ZXMjA=",
      "dev": true,
      "requires": {
        "unique-slug": "^2.0.0"
      }
    },
    "unique-slug": {
      "version": "2.0.1",
      "resolved": "http://registry.npm.taobao.org/unique-slug/download/unique-slug-2.0.1.tgz",
      "integrity": "sha1-Xp7cbRzo+yZNsYpQfvm9hURFHKY=",
      "dev": true,
      "requires": {
        "imurmurhash": "^0.1.4"
      }
    },
    "universalify": {
      "version": "0.1.2",
      "resolved": "http://registry.npm.taobao.org/universalify/download/universalify-0.1.2.tgz",
      "integrity": "sha1-tkb2m+OULavOzJ1mOcgNwQXvqmY=",
      "dev": true
    },
    "unpipe": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/unpipe/download/unpipe-1.0.0.tgz",
      "integrity": "sha1-sr9O6FFKrmFltIF4KdIbLvSZBOw=",
      "dev": true
    },
    "unquote": {
      "version": "1.1.1",
      "resolved": "http://registry.npm.taobao.org/unquote/download/unquote-1.1.1.tgz",
      "integrity": "sha1-j97XMk7G6IoP+LkF58CYzcCG1UQ=",
      "dev": true
    },
    "unset-value": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/unset-value/download/unset-value-1.0.0.tgz",
      "integrity": "sha1-g3aHP30jNRef+x5vw6jtDfyKtVk=",
      "dev": true,
      "requires": {
        "has-value": "^0.3.1",
        "isobject": "^3.0.0"
      },
      "dependencies": {
        "has-value": {
          "version": "0.3.1",
          "resolved": "http://registry.npm.taobao.org/has-value/download/has-value-0.3.1.tgz",
          "integrity": "sha1-ex9YutpiyoJ+wKIHgCVlSEWZXh8=",
          "dev": true,
          "requires": {
            "get-value": "^2.0.3",
            "has-values": "^0.1.4",
            "isobject": "^2.0.0"
          },
          "dependencies": {
            "isobject": {
              "version": "2.1.0",
              "resolved": "http://registry.npm.taobao.org/isobject/download/isobject-2.1.0.tgz",
              "integrity": "sha1-8GVWEJaj8dou9GJy+BXIQNh+DIk=",
              "dev": true,
              "requires": {
                "isarray": "1.0.0"
              }
            }
          }
        },
        "has-values": {
          "version": "0.1.4",
          "resolved": "http://registry.npm.taobao.org/has-values/download/has-values-0.1.4.tgz",
          "integrity": "sha1-bWHeldkd/Km5oCCJrThL/49it3E=",
          "dev": true
        }
      }
    },
    "upath": {
      "version": "1.1.0",
      "resolved": "http://registry.npm.taobao.org/upath/download/upath-1.1.0.tgz",
      "integrity": "sha1-NSVll+RqWB20eT0M5H+prr/J+r0=",
      "dev": true
    },
    "upper-case": {
      "version": "1.1.3",
      "resolved": "http://registry.npm.taobao.org/upper-case/download/upper-case-1.1.3.tgz",
      "integrity": "sha1-9rRQHC7EzdJrp4vnIilh3ndiFZg=",
      "dev": true
    },
    "uri-js": {
      "version": "4.2.2",
      "resolved": "http://registry.npm.taobao.org/uri-js/download/uri-js-4.2.2.tgz",
      "integrity": "sha1-lMVA4f93KVbiKZUHwBCupsiDjrA=",
      "dev": true,
      "requires": {
        "punycode": "^2.1.0"
      }
    },
    "urix": {
      "version": "0.1.0",
      "resolved": "http://registry.npm.taobao.org/urix/download/urix-0.1.0.tgz",
      "integrity": "sha1-2pN/emLiH+wf0Y1Js1wpNQZ6bHI=",
      "dev": true
    },
    "url": {
      "version": "0.11.0",
      "resolved": "http://registry.npm.taobao.org/url/download/url-0.11.0.tgz",
      "integrity": "sha1-ODjpfPxgUh63PFJajlW/3Z4uKPE=",
      "dev": true,
      "requires": {
        "punycode": "1.3.2",
        "querystring": "0.2.0"
      },
      "dependencies": {
        "punycode": {
          "version": "1.3.2",
          "resolved": "http://registry.npm.taobao.org/punycode/download/punycode-1.3.2.tgz",
          "integrity": "sha1-llOgNvt8HuQjQvIyXM7v6jkmxI0=",
          "dev": true
        }
      }
    },
    "url-loader": {
      "version": "1.1.2",
      "resolved": "http://registry.npm.taobao.org/url-loader/download/url-loader-1.1.2.tgz",
      "integrity": "sha1-uXHRkbg69pPF4/6kBkvp4fLX+Ng=",
      "dev": true,
      "requires": {
        "loader-utils": "^1.1.0",
        "mime": "^2.0.3",
        "schema-utils": "^1.0.0"
      },
      "dependencies": {
        "ajv-keywords": {
          "version": "3.2.0",
          "resolved": "http://registry.npm.taobao.org/ajv-keywords/download/ajv-keywords-3.2.0.tgz",
          "integrity": "sha1-6GuBnGAs+IIa1jdBNpjx3sAhhHo=",
          "dev": true
        },
        "schema-utils": {
          "version": "1.0.0",
          "resolved": "http://registry.npm.taobao.org/schema-utils/download/schema-utils-1.0.0.tgz",
          "integrity": "sha1-C3mpMgTXtgDUsoUNH2bCo0lRx3A=",
          "dev": true,
          "requires": {
            "ajv": "^6.1.0",
            "ajv-errors": "^1.0.0",
            "ajv-keywords": "^3.1.0"
          }
        }
      }
    },
    "url-parse": {
      "version": "1.4.4",
      "resolved": "http://registry.npm.taobao.org/url-parse/download/url-parse-1.4.4.tgz",
      "integrity": "sha1-ysFVbpX6oDA2kf7Fz51aG8NGSPg=",
      "dev": true,
      "requires": {
        "querystringify": "^2.0.0",
        "requires-port": "^1.0.0"
      }
    },
    "use": {
      "version": "3.1.1",
      "resolved": "http://registry.npm.taobao.org/use/download/use-3.1.1.tgz",
      "integrity": "sha1-1QyMrHmhn7wg8pEfVuuXP04QBw8=",
      "dev": true
    },
    "util": {
      "version": "0.10.4",
      "resolved": "http://registry.npm.taobao.org/util/download/util-0.10.4.tgz",
      "integrity": "sha1-OqASW/5mikZy3liFfTrOJ+y3aQE=",
      "dev": true,
      "requires": {
        "inherits": "2.0.3"
      }
    },
    "util-deprecate": {
      "version": "1.0.2",
      "resolved": "http://registry.npm.taobao.org/util-deprecate/download/util-deprecate-1.0.2.tgz",
      "integrity": "sha1-RQ1Nyfpw3nMnYvvS1KKJgUGaDM8=",
      "dev": true
    },
    "util.promisify": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/util.promisify/download/util.promisify-1.0.0.tgz",
      "integrity": "sha1-RA9xZaRZyaFtwUXrjnLzVocJcDA=",
      "dev": true,
      "requires": {
        "define-properties": "^1.1.2",
        "object.getownpropertydescriptors": "^2.0.3"
      }
    },
    "utila": {
      "version": "0.4.0",
      "resolved": "http://registry.npm.taobao.org/utila/download/utila-0.4.0.tgz",
      "integrity": "sha1-ihagXURWV6Oupe7MWxKk+lN5dyw=",
      "dev": true
    },
    "utils-merge": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/utils-merge/download/utils-merge-1.0.1.tgz",
      "integrity": "sha1-n5VxD1CiZ5R7LMwSR0HBAoQn5xM=",
      "dev": true
    },
    "uuid": {
      "version": "3.3.2",
      "resolved": "http://registry.npm.taobao.org/uuid/download/uuid-3.3.2.tgz",
      "integrity": "sha1-G0r0lV6zB3xQHCOHL8ZROBFYcTE=",
      "dev": true
    },
    "validate-npm-package-license": {
      "version": "3.0.4",
      "resolved": "http://registry.npm.taobao.org/validate-npm-package-license/download/validate-npm-package-license-3.0.4.tgz",
      "integrity": "sha1-/JH2uce6FchX9MssXe/uw51PQQo=",
      "dev": true,
      "requires": {
        "spdx-correct": "^3.0.0",
        "spdx-expression-parse": "^3.0.0"
      }
    },
    "vary": {
      "version": "1.1.2",
      "resolved": "http://registry.npm.taobao.org/vary/download/vary-1.1.2.tgz",
      "integrity": "sha1-IpnwLG3tMNSllhsLn3RSShj2NPw=",
      "dev": true
    },
    "vendors": {
      "version": "1.0.2",
      "resolved": "http://registry.npm.taobao.org/vendors/download/vendors-1.0.2.tgz",
      "integrity": "sha1-f8te759WI7FWvOqJ7DfWNnbyGAE=",
      "dev": true
    },
    "verror": {
      "version": "1.10.0",
      "resolved": "http://registry.npm.taobao.org/verror/download/verror-1.10.0.tgz",
      "integrity": "sha1-OhBcoXBTr1XW4nDB+CiGguGNpAA=",
      "dev": true,
      "requires": {
        "assert-plus": "^1.0.0",
        "core-util-is": "1.0.2",
        "extsprintf": "^1.2.0"
      }
    },
    "vm-browserify": {
      "version": "0.0.4",
      "resolved": "http://registry.npm.taobao.org/vm-browserify/download/vm-browserify-0.0.4.tgz",
      "integrity": "sha1-XX6kW7755Kb/ZflUOOCofDV9WnM=",
      "dev": true,
      "requires": {
        "indexof": "0.0.1"
      }
    },
    "vue": {
      "version": "2.5.21",
      "resolved": "http://registry.npm.taobao.org/vue/download/vue-2.5.21.tgz",
      "integrity": "sha1-PTPc0Du4E5Es6JSoMDq1U2mcSoU="
    },
    "vue-eslint-parser": {
      "version": "2.0.3",
      "resolved": "http://registry.npm.taobao.org/vue-eslint-parser/download/vue-eslint-parser-2.0.3.tgz",
      "integrity": "sha1-wmjJbG2Uz+PZOKX3WTlZsMozYNE=",
      "dev": true,
      "requires": {
        "debug": "^3.1.0",
        "eslint-scope": "^3.7.1",
        "eslint-visitor-keys": "^1.0.0",
        "espree": "^3.5.2",
        "esquery": "^1.0.0",
        "lodash": "^4.17.4"
      },
      "dependencies": {
        "debug": {
          "version": "3.2.6",
          "resolved": "http://registry.npm.taobao.org/debug/download/debug-3.2.6.tgz",
          "integrity": "sha1-6D0X3hbYp++3cX7b5fsQE17uYps=",
          "dev": true,
          "requires": {
            "ms": "^2.1.1"
          }
        }
      }
    },
    "vue-hot-reload-api": {
      "version": "2.3.1",
      "resolved": "http://registry.npm.taobao.org/vue-hot-reload-api/download/vue-hot-reload-api-2.3.1.tgz",
      "integrity": "sha1-stPZVAKoEWAjgHg+pPVm64dVaaI=",
      "dev": true
    },
    "vue-loader": {
      "version": "15.4.2",
      "resolved": "http://registry.npm.taobao.org/vue-loader/download/vue-loader-15.4.2.tgz",
      "integrity": "sha1-gSuybkR907hMSF62NBkNkUzhJeI=",
      "dev": true,
      "requires": {
        "@vue/component-compiler-utils": "^2.0.0",
        "hash-sum": "^1.0.2",
        "loader-utils": "^1.1.0",
        "vue-hot-reload-api": "^2.3.0",
        "vue-style-loader": "^4.1.0"
      }
    },
    "vue-router": {
      "version": "3.0.2",
      "resolved": "https://registry.npmjs.org/vue-router/-/vue-router-3.0.2.tgz",
      "integrity": "sha512-opKtsxjp9eOcFWdp6xLQPLmRGgfM932Tl56U9chYTnoWqKxQ8M20N7AkdEbM5beUh6wICoFGYugAX9vQjyJLFg=="
    },
    "vue-style-loader": {
      "version": "4.1.2",
      "resolved": "http://registry.npm.taobao.org/vue-style-loader/download/vue-style-loader-4.1.2.tgz",
      "integrity": "sha1-3t80mAbyXOtOZPOtfApE+6c1/Pg=",
      "dev": true,
      "requires": {
        "hash-sum": "^1.0.2",
        "loader-utils": "^1.0.2"
      }
    },
    "vue-template-compiler": {
      "version": "2.5.21",
      "resolved": "http://registry.npm.taobao.org/vue-template-compiler/download/vue-template-compiler-2.5.21.tgz",
      "integrity": "sha1-pXzrkDF36PZDVgqNY5oPjbZHBUo=",
      "dev": true,
      "requires": {
        "de-indent": "^1.0.2",
        "he": "^1.1.0"
      }
    },
    "vue-template-es2015-compiler": {
      "version": "1.6.0",
      "resolved": "http://registry.npm.taobao.org/vue-template-es2015-compiler/download/vue-template-es2015-compiler-1.6.0.tgz",
      "integrity": "sha1-3EJpcTMwLOMBdSQ1amxht7abShg=",
      "dev": true
    },
    "vuex": {
      "version": "3.0.1",
      "resolved": "https://registry.npmjs.org/vuex/-/vuex-3.0.1.tgz",
      "integrity": "sha512-wLoqz0B7DSZtgbWL1ShIBBCjv22GV5U+vcBFox658g6V0s4wZV9P4YjCNyoHSyIBpj1f29JBoNQIqD82cR4O3w=="
    },
    "watchpack": {
      "version": "1.6.0",
      "resolved": "http://registry.npm.taobao.org/watchpack/download/watchpack-1.6.0.tgz",
      "integrity": "sha1-S8EsLr6KonenHx0/FNaFx7RGzQA=",
      "dev": true,
      "requires": {
        "chokidar": "^2.0.2",
        "graceful-fs": "^4.1.2",
        "neo-async": "^2.5.0"
      }
    },
    "wbuf": {
      "version": "1.7.3",
      "resolved": "http://registry.npm.taobao.org/wbuf/download/wbuf-1.7.3.tgz",
      "integrity": "sha1-wdjRSTFtPqhShIiVy2oL/oh7h98=",
      "dev": true,
      "requires": {
        "minimalistic-assert": "^1.0.0"
      }
    },
    "wcwidth": {
      "version": "1.0.1",
      "resolved": "http://registry.npm.taobao.org/wcwidth/download/wcwidth-1.0.1.tgz",
      "integrity": "sha1-8LDc+RW8X/FSivrbLA4XtTLaL+g=",
      "dev": true,
      "requires": {
        "defaults": "^1.0.3"
      }
    },
    "webpack": {
      "version": "4.27.1",
      "resolved": "http://registry.npm.taobao.org/webpack/download/webpack-4.27.1.tgz",
      "integrity": "sha1-Xy4ttEbSJmN2+hXX0id6GpwuErs=",
      "dev": true,
      "requires": {
        "@webassemblyjs/ast": "1.7.11",
        "@webassemblyjs/helper-module-context": "1.7.11",
        "@webassemblyjs/wasm-edit": "1.7.11",
        "@webassemblyjs/wasm-parser": "1.7.11",
        "acorn": "^5.6.2",
        "acorn-dynamic-import": "^3.0.0",
        "ajv": "^6.1.0",
        "ajv-keywords": "^3.1.0",
        "chrome-trace-event": "^1.0.0",
        "enhanced-resolve": "^4.1.0",
        "eslint-scope": "^4.0.0",
        "json-parse-better-errors": "^1.0.2",
        "loader-runner": "^2.3.0",
        "loader-utils": "^1.1.0",
        "memory-fs": "~0.4.1",
        "micromatch": "^3.1.8",
        "mkdirp": "~0.5.0",
        "neo-async": "^2.5.0",
        "node-libs-browser": "^2.0.0",
        "schema-utils": "^0.4.4",
        "tapable": "^1.1.0",
        "terser-webpack-plugin": "^1.1.0",
        "watchpack": "^1.5.0",
        "webpack-sources": "^1.3.0"
      },
      "dependencies": {
        "ajv-keywords": {
          "version": "3.2.0",
          "resolved": "http://registry.npm.taobao.org/ajv-keywords/download/ajv-keywords-3.2.0.tgz",
          "integrity": "sha1-6GuBnGAs+IIa1jdBNpjx3sAhhHo=",
          "dev": true
        },
        "eslint-scope": {
          "version": "4.0.0",
          "resolved": "http://registry.npm.taobao.org/eslint-scope/download/eslint-scope-4.0.0.tgz",
          "integrity": "sha1-UL8wcekzi83EMzF5Sgy1M/ATYXI=",
          "dev": true,
          "requires": {
            "esrecurse": "^4.1.0",
            "estraverse": "^4.1.1"
          }
        }
      }
    },
    "webpack-bundle-analyzer": {
      "version": "3.0.3",
      "resolved": "http://registry.npm.taobao.org/webpack-bundle-analyzer/download/webpack-bundle-analyzer-3.0.3.tgz",
      "integrity": "sha1-28f/+PUgWLZxSiD93zCdB5Dj4KA=",
      "dev": true,
      "requires": {
        "acorn": "^5.7.3",
        "bfj": "^6.1.1",
        "chalk": "^2.4.1",
        "commander": "^2.18.0",
        "ejs": "^2.6.1",
        "express": "^4.16.3",
        "filesize": "^3.6.1",
        "gzip-size": "^5.0.0",
        "lodash": "^4.17.10",
        "mkdirp": "^0.5.1",
        "opener": "^1.5.1",
        "ws": "^6.0.0"
      },
      "dependencies": {
        "commander": {
          "version": "2.19.0",
          "resolved": "http://registry.npm.taobao.org/commander/download/commander-2.19.0.tgz",
          "integrity": "sha1-9hmKqE5bg8RgVLlN3tv+1e6f8So=",
          "dev": true
        }
      }
    },
    "webpack-chain": {
      "version": "4.12.1",
      "resolved": "http://registry.npm.taobao.org/webpack-chain/download/webpack-chain-4.12.1.tgz",
      "integrity": "sha1-bIQ5u7KrVQlS1g4eqTGRQZBsAqY=",
      "dev": true,
      "requires": {
        "deepmerge": "^1.5.2",
        "javascript-stringify": "^1.6.0"
      }
    },
    "webpack-dev-middleware": {
      "version": "3.4.0",
      "resolved": "http://registry.npm.taobao.org/webpack-dev-middleware/download/webpack-dev-middleware-3.4.0.tgz",
      "integrity": "sha1-ETL+zJAm/ZDw7O2sXL/3XR+0WJA=",
      "dev": true,
      "requires": {
        "memory-fs": "~0.4.1",
        "mime": "^2.3.1",
        "range-parser": "^1.0.3",
        "webpack-log": "^2.0.0"
      }
    },
    "webpack-dev-server": {
      "version": "3.1.10",
      "resolved": "http://registry.npm.taobao.org/webpack-dev-server/download/webpack-dev-server-3.1.10.tgz",
      "integrity": "sha1-UHQRvucn7o0v3/3GIbZqZKs96is=",
      "dev": true,
      "requires": {
        "ansi-html": "0.0.7",
        "bonjour": "^3.5.0",
        "chokidar": "^2.0.0",
        "compression": "^1.5.2",
        "connect-history-api-fallback": "^1.3.0",
        "debug": "^3.1.0",
        "del": "^3.0.0",
        "express": "^4.16.2",
        "html-entities": "^1.2.0",
        "http-proxy-middleware": "~0.18.0",
        "import-local": "^2.0.0",
        "internal-ip": "^3.0.1",
        "ip": "^1.1.5",
        "killable": "^1.0.0",
        "loglevel": "^1.4.1",
        "opn": "^5.1.0",
        "portfinder": "^1.0.9",
        "schema-utils": "^1.0.0",
        "selfsigned": "^1.9.1",
        "serve-index": "^1.7.2",
        "sockjs": "0.3.19",
        "sockjs-client": "1.3.0",
        "spdy": "^3.4.1",
        "strip-ansi": "^3.0.0",
        "supports-color": "^5.1.0",
        "webpack-dev-middleware": "3.4.0",
        "webpack-log": "^2.0.0",
        "yargs": "12.0.2"
      },
      "dependencies": {
        "ajv-keywords": {
          "version": "3.2.0",
          "resolved": "http://registry.npm.taobao.org/ajv-keywords/download/ajv-keywords-3.2.0.tgz",
          "integrity": "sha1-6GuBnGAs+IIa1jdBNpjx3sAhhHo=",
          "dev": true
        },
        "ansi-regex": {
          "version": "2.1.1",
          "resolved": "http://registry.npm.taobao.org/ansi-regex/download/ansi-regex-2.1.1.tgz",
          "integrity": "sha1-w7M6te42DYbg5ijwRorn7yfWVN8=",
          "dev": true
        },
        "debug": {
          "version": "3.2.6",
          "resolved": "http://registry.npm.taobao.org/debug/download/debug-3.2.6.tgz",
          "integrity": "sha1-6D0X3hbYp++3cX7b5fsQE17uYps=",
          "dev": true,
          "requires": {
            "ms": "^2.1.1"
          }
        },
        "schema-utils": {
          "version": "1.0.0",
          "resolved": "http://registry.npm.taobao.org/schema-utils/download/schema-utils-1.0.0.tgz",
          "integrity": "sha1-C3mpMgTXtgDUsoUNH2bCo0lRx3A=",
          "dev": true,
          "requires": {
            "ajv": "^6.1.0",
            "ajv-errors": "^1.0.0",
            "ajv-keywords": "^3.1.0"
          }
        },
        "strip-ansi": {
          "version": "3.0.1",
          "resolved": "http://registry.npm.taobao.org/strip-ansi/download/strip-ansi-3.0.1.tgz",
          "integrity": "sha1-ajhfuIU9lS1f8F0Oiq+UJ43GPc8=",
          "dev": true,
          "requires": {
            "ansi-regex": "^2.0.0"
          }
        }
      }
    },
    "webpack-log": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/webpack-log/download/webpack-log-2.0.0.tgz",
      "integrity": "sha1-W3ko4GN1k/EZ0y9iJ8HgrDHhtH8=",
      "dev": true,
      "requires": {
        "ansi-colors": "^3.0.0",
        "uuid": "^3.3.2"
      }
    },
    "webpack-merge": {
      "version": "4.1.5",
      "resolved": "http://registry.npm.taobao.org/webpack-merge/download/webpack-merge-4.1.5.tgz",
      "integrity": "sha1-K+MehGwgdn0b71a9ymTDKKaBGQo=",
      "dev": true,
      "requires": {
        "lodash": "^4.17.5"
      }
    },
    "webpack-sources": {
      "version": "1.3.0",
      "resolved": "http://registry.npm.taobao.org/webpack-sources/download/webpack-sources-1.3.0.tgz",
      "integrity": "sha1-KijcufH0X+lg2PFJMlK17mUw+oU=",
      "dev": true,
      "requires": {
        "source-list-map": "^2.0.0",
        "source-map": "~0.6.1"
      },
      "dependencies": {
        "source-map": {
          "version": "0.6.1",
          "resolved": "http://registry.npm.taobao.org/source-map/download/source-map-0.6.1.tgz",
          "integrity": "sha1-dHIq8y6WFOnCh6jQu95IteLxomM=",
          "dev": true
        }
      }
    },
    "websocket-driver": {
      "version": "0.7.0",
      "resolved": "http://registry.npm.taobao.org/websocket-driver/download/websocket-driver-0.7.0.tgz",
      "integrity": "sha1-DK+dLXVdk67gSdS90NP+LMoqJOs=",
      "dev": true,
      "requires": {
        "http-parser-js": ">=0.4.0",
        "websocket-extensions": ">=0.1.1"
      }
    },
    "websocket-extensions": {
      "version": "0.1.3",
      "resolved": "http://registry.npm.taobao.org/websocket-extensions/download/websocket-extensions-0.1.3.tgz",
      "integrity": "sha1-XS/yKXcAPsaHpLhwc9+7rBRszyk=",
      "dev": true
    },
    "which": {
      "version": "1.3.1",
      "resolved": "http://registry.npm.taobao.org/which/download/which-1.3.1.tgz",
      "integrity": "sha1-pFBD1U9YBTFtqNYvn1CRjT2nCwo=",
      "dev": true,
      "requires": {
        "isexe": "^2.0.0"
      }
    },
    "which-module": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/which-module/download/which-module-2.0.0.tgz",
      "integrity": "sha1-2e8H3Od7mQK4o6j6SzHD4/fm6Ho=",
      "dev": true
    },
    "wordwrap": {
      "version": "1.0.0",
      "resolved": "http://registry.npm.taobao.org/wordwrap/download/wordwrap-1.0.0.tgz",
      "integrity": "sha1-J1hIEIkUVqQXHI0CJkQa3pDLyus=",
      "dev": true
    },
    "worker-farm": {
      "version": "1.6.0",
      "resolved": "http://registry.npm.taobao.org/worker-farm/download/worker-farm-1.6.0.tgz",
      "integrity": "sha1-rsxAWXb6talVJhgIRvDboojzpKA=",
      "dev": true,
      "requires": {
        "errno": "~0.1.7"
      }
    },
    "wrap-ansi": {
      "version": "2.1.0",
      "resolved": "http://registry.npm.taobao.org/wrap-ansi/download/wrap-ansi-2.1.0.tgz",
      "integrity": "sha1-2Pw9KE3QV5T+hJc8rs3Rz4JP3YU=",
      "dev": true,
      "requires": {
        "string-width": "^1.0.1",
        "strip-ansi": "^3.0.1"
      },
      "dependencies": {
        "ansi-regex": {
          "version": "2.1.1",
          "resolved": "http://registry.npm.taobao.org/ansi-regex/download/ansi-regex-2.1.1.tgz",
          "integrity": "sha1-w7M6te42DYbg5ijwRorn7yfWVN8=",
          "dev": true
        },
        "is-fullwidth-code-point": {
          "version": "1.0.0",
          "resolved": "http://registry.npm.taobao.org/is-fullwidth-code-point/download/is-fullwidth-code-point-1.0.0.tgz",
          "integrity": "sha1-754xOG8DGn8NZDr4L95QxFfvAMs=",
          "dev": true,
          "requires": {
            "number-is-nan": "^1.0.0"
          }
        },
        "string-width": {
          "version": "1.0.2",
          "resolved": "http://registry.npm.taobao.org/string-width/download/string-width-1.0.2.tgz",
          "integrity": "sha1-EYvfW4zcUaKn5w0hHgfisLmxB9M=",
          "dev": true,
          "requires": {
            "code-point-at": "^1.0.0",
            "is-fullwidth-code-point": "^1.0.0",
            "strip-ansi": "^3.0.0"
          }
        },
        "strip-ansi": {
          "version": "3.0.1",
          "resolved": "http://registry.npm.taobao.org/strip-ansi/download/strip-ansi-3.0.1.tgz",
          "integrity": "sha1-ajhfuIU9lS1f8F0Oiq+UJ43GPc8=",
          "dev": true,
          "requires": {
            "ansi-regex": "^2.0.0"
          }
        }
      }
    },
    "wrappy": {
      "version": "1.0.2",
      "resolved": "http://registry.npm.taobao.org/wrappy/download/wrappy-1.0.2.tgz",
      "integrity": "sha1-tSQ9jz7BqjXxNkYFvA0QNuMKtp8=",
      "dev": true
    },
    "write": {
      "version": "0.2.1",
      "resolved": "http://registry.npm.taobao.org/write/download/write-0.2.1.tgz",
      "integrity": "sha1-X8A4KOJkzqP+kUVUdvejxWbLB1c=",
      "dev": true,
      "requires": {
        "mkdirp": "^0.5.1"
      }
    },
    "ws": {
      "version": "6.1.2",
      "resolved": "http://registry.npm.taobao.org/ws/download/ws-6.1.2.tgz",
      "integrity": "sha1-PMdGLph5LwrGeUJBSJA97TucOtg=",
      "dev": true,
      "requires": {
        "async-limiter": "~1.0.0"
      }
    },
    "xregexp": {
      "version": "4.0.0",
      "resolved": "http://registry.npm.taobao.org/xregexp/download/xregexp-4.0.0.tgz",
      "integrity": "sha1-5pgYneSd0qGMxWh7BeF8jkOUMCA=",
      "dev": true
    },
    "xtend": {
      "version": "4.0.1",
      "resolved": "http://registry.npm.taobao.org/xtend/download/xtend-4.0.1.tgz",
      "integrity": "sha1-pcbVMr5lbiPbgg77lDofBJmNY68=",
      "dev": true
    },
    "y18n": {
      "version": "4.0.0",
      "resolved": "http://registry.npm.taobao.org/y18n/download/y18n-4.0.0.tgz",
      "integrity": "sha1-le+U+F7MgdAHwmThkKEg8KPIVms=",
      "dev": true
    },
    "yallist": {
      "version": "2.1.2",
      "resolved": "http://registry.npm.taobao.org/yallist/download/yallist-2.1.2.tgz",
      "integrity": "sha1-HBH5IY8HYImkfdUS+TxmmaaoHVI=",
      "dev": true
    },
    "yargs": {
      "version": "12.0.2",
      "resolved": "http://registry.npm.taobao.org/yargs/download/yargs-12.0.2.tgz",
      "integrity": "sha1-/lgjQ2k5KvM+y+9TgZFx7/D1qtw=",
      "dev": true,
      "requires": {
        "cliui": "^4.0.0",
        "decamelize": "^2.0.0",
        "find-up": "^3.0.0",
        "get-caller-file": "^1.0.1",
        "os-locale": "^3.0.0",
        "require-directory": "^2.1.1",
        "require-main-filename": "^1.0.1",
        "set-blocking": "^2.0.0",
        "string-width": "^2.0.0",
        "which-module": "^2.0.0",
        "y18n": "^3.2.1 || ^4.0.0",
        "yargs-parser": "^10.1.0"
      },
      "dependencies": {
        "find-up": {
          "version": "3.0.0",
          "resolved": "http://registry.npm.taobao.org/find-up/download/find-up-3.0.0.tgz",
          "integrity": "sha1-SRafHXmTQwZG2mHsxa41XCHJe3M=",
          "dev": true,
          "requires": {
            "locate-path": "^3.0.0"
          }
        },
        "locate-path": {
          "version": "3.0.0",
          "resolved": "http://registry.npm.taobao.org/locate-path/download/locate-path-3.0.0.tgz",
          "integrity": "sha1-2+w7OrdZdYBxtY/ln8QYca8hQA4=",
          "dev": true,
          "requires": {
            "p-locate": "^3.0.0",
            "path-exists": "^3.0.0"
          }
        },
        "p-limit": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/p-limit/download/p-limit-2.0.0.tgz",
          "integrity": "sha1-5iTtVO6MRgp3izyfNnBJb/ileuw=",
          "dev": true,
          "requires": {
            "p-try": "^2.0.0"
          }
        },
        "p-locate": {
          "version": "3.0.0",
          "resolved": "http://registry.npm.taobao.org/p-locate/download/p-locate-3.0.0.tgz",
          "integrity": "sha1-Mi1poFwCZLJZl9n0DNiokasAZKQ=",
          "dev": true,
          "requires": {
            "p-limit": "^2.0.0"
          }
        },
        "p-try": {
          "version": "2.0.0",
          "resolved": "http://registry.npm.taobao.org/p-try/download/p-try-2.0.0.tgz",
          "integrity": "sha1-hQgLuHxkaI+keZb+j3376CEXYLE=",
          "dev": true
        }
      }
    },
    "yargs-parser": {
      "version": "10.1.0",
      "resolved": "http://registry.npm.taobao.org/yargs-parser/download/yargs-parser-10.1.0.tgz",
      "integrity": "sha1-cgImW4n36eny5XZeD+c1qQXtuqg=",
      "dev": true,
      "requires": {
        "camelcase": "^4.1.0"
      }
    },
    "yorkie": {
      "version": "2.0.0",
      "resolved": "http://registry.npm.taobao.org/yorkie/download/yorkie-2.0.0.tgz",
      "integrity": "sha1-kkEZEtQ1IU4SxRwq4Qk+VLa7g9k=",
      "dev": true,
      "requires": {
        "execa": "^0.8.0",
        "is-ci": "^1.0.10",
        "normalize-path": "^1.0.0",
        "strip-indent": "^2.0.0"
      },
      "dependencies": {
        "cross-spawn": {
          "version": "5.1.0",
          "resolved": "http://registry.npm.taobao.org/cross-spawn/download/cross-spawn-5.1.0.tgz",
          "integrity": "sha1-6L0O/uWPz/b4+UUQoKVUu/ojVEk=",
          "dev": true,
          "requires": {
            "lru-cache": "^4.0.1",
            "shebang-command": "^1.2.0",
            "which": "^1.2.9"
          }
        },
        "execa": {
          "version": "0.8.0",
          "resolved": "http://registry.npm.taobao.org/execa/download/execa-0.8.0.tgz",
          "integrity": "sha1-2NdrvBtVIX7RkP1t1J08d07PyNo=",
          "dev": true,
          "requires": {
            "cross-spawn": "^5.0.1",
            "get-stream": "^3.0.0",
            "is-stream": "^1.1.0",
            "npm-run-path": "^2.0.0",
            "p-finally": "^1.0.0",
            "signal-exit": "^3.0.0",
            "strip-eof": "^1.0.0"
          }
        },
        "get-stream": {
          "version": "3.0.0",
          "resolved": "http://registry.npm.taobao.org/get-stream/download/get-stream-3.0.0.tgz",
          "integrity": "sha1-jpQ9E1jcN1VQVOy+LtsFqhdO3hQ=",
          "dev": true
        },
        "normalize-path": {
          "version": "1.0.0",
          "resolved": "http://registry.npm.taobao.org/normalize-path/download/normalize-path-1.0.0.tgz",
          "integrity": "sha1-MtDkcvkf80VwHBWoMRAY07CpA3k=",
          "dev": true
        }
      }
    }
  }
}
`
