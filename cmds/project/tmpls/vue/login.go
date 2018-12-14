package vue

const LoginTpl = `
<template>
  <div id="app">
    <login-with-up
      :copyright="copyright"
      :systemName="systemName"
      :conf="conf"
      @loginCall="call"
      ref="loginItem"
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
        systemName: "{{.projectName}} 系统",
        copyright: "{{.projectName}} Copyright&copy; 2018 版权所有 蜀ICP备 17001788号-2",
        conf: {loginNameType: "请输入邮箱或用户名", pwd: "输入密码"},   //输入框提示信息配置
      }
    },
    components: { //注册插件
      loginWithUp
    },
    mounted(){
      document.title = "登录 - gaea 系统"
      //sessionStorage.removeItem("__jwt__")
    },
    methods: {
      call: function (code) {
        this.$get("/member/login", code)
          .then(response => {
            sessionStorage.setItem("userinfo",JSON.stringify(response))
            this.$refs.loginItem.showMsg("登录中.....");

            setTimeout(() => {

              this.$router.push("/")

            }, 500);

          })
          .catch(err => {
            if (err.response.status == 403 || err.response.status == 406){
              this.$refs.loginItem.showMsg("用户名或密码错误");
            }else{
              this.$refs.loginItem.showMsg("系统繁忙")
            }
          })
      }
    }
  }
</script>`
