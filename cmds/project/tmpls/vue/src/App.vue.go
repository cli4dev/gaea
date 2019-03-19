package src

const AppVue = `<template>
  <div>
    <router-view/>
  </div>
</template>

<script>
  import VueCookies from 'vue-cookies'
  export default {
    created(){
      if (!sessionStorage.getItem("systeminfo")) {
        this.$get("/member/getsysinfo",{})
          .then(res=>{
          sessionStorage.setItem("systeminfo",JSON.stringify(res))
        }).catch(err=>{
          console.log(err)
        });
      }
    },
    moutend(){
      if (!sessionStorage.getItem("userinfo")) {
          this.$router.push("/login")
      }
    },
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
