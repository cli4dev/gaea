package src

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

const RouterString = `
{{range $i,$c:=.router -}}
    {
      path: '{{$i}}',
      name:'{{$i|vname}}',
      component: () => import('./pages/{{$c}}')
    },
{{- end -}}`
