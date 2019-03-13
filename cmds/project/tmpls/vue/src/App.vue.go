package src

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
