package pages

const LoginTpl = `
<template>
  <div id="app">
    <login-with-up
      :call="call"
      :err-msg.sync="errMsg"
      ref="loginUp"
    >
    </login-with-up>
  </div>
</template>

<script>
  import loginWithUp from 'login-with-up'; // 引入
  export default {
    name: 'app',
    data() {
      return {
        errMsg: {},
      }
    },
    components: { //注册插件
      loginWithUp
    },
    mounted(){
      document.title = "登录 - "+JSON.parse(sessionStorage.getItem("systeminfo")).name
    },
    methods: {
      call(code) {
        this.$get("/member/login", code)
          .then(response => {
            sessionStorage.setItem("userinfo",JSON.stringify(response))
            this.$refs.loginUp.success("/")
          })
          .catch(err => {
            this.errMsg = {message:err.response.data,timestamp:  Date.parse(new Date())}
          })
      }
    }
  }
</script>`
