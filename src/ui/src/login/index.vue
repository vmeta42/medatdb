<template>
  <div class="wrapper">
    <form class="form">
      <h1 class="title">
        <img class="logo" src="../assets/images/shijihulianLogo.png" alt="logo" height="120">
      </h1>
      <div class="form-error" ref="error">{{error}}</div>
      <div class="form-item">
        <img class="form-item-icon" src="./assets/user.svg" width="16" height="16">
        <input id="user" type="text" name="username" placeholder="用户名" autocomplete="off" v-model.trim="username">
      </div>
      <div class="form-item">
        <img class="form-item-icon" src="./assets/password.svg" width="16" height="16">
        <input class="password" id="password" type="password" name="password" placeholder="密码" v-model.trim="password">
      </div>
      <button class="form-submit" type="submit" @click.stop.prevent="handleSubmit">登录</button>
    </form>
  </div>
</template>

<script>
  // import Axios from 'axios'

  export default {
    data() {
      return {
        username: '',
        password: '',
        error: window.LOGIN_ERROR
      }
    },
    methods: {
      handleSubmit() {
        if (!this.username.length) {
          this.error = '请输入用户名'
        } else if (!this.password.length) {
          this.error = '请输入密码'
        } else {
          if (this.username === 'admin' && this.password === 'admin') {
            window.User.name = this.username
            this.$router.push('/')
          } else {
            this.error = '鉴权失败'
          }
          // axios实例
          // const axiosInstance = Axios.create({
          //   baseURL: '/',
          //   xsrfCookieName: 'data_csrftoken',
          //   xsrfHeaderName: 'X-CSRFToken',
          //   withCredentials: true
          // })
          // const formData = new FormData()
          // formData.set('username', window.btoa(this.username))
          // formData.set('password', window.btoa(this.password))
          // axiosInstance.post('ldap/auth', formData).then(() => {
          //   window.User.name = this.username
          //   this.$router.push('/')
          // }).catch(() => {
          //   this.error = '鉴权失败'
          // })
        }
      }
    }
  }
</script>

<style lang="scss" scoped>
.wrapper {
    height: 100%;
    overflow: hidden;
    background: url(./assets/login_bg.png) center center no-repeat;
    background-size: 100% 100%;
}
.form {
    width: 400px;
    height: 400px;
    background-color: #fff;
    border-radius: 2px;
    margin: 0 auto;
    margin-top: calc((100vh - 400px) / 3);
    .title {
        height: 110px;
        display: flex;
        align-items: center;
        justify-content: center;
        border-bottom: 1px solid #F0F1F5;
    }
    .form-error {
        width: 290px;
        height: 42px;
        margin: 0 auto;
        line-height: 42px;
        color: #EA3636;
    }
    .form-item {
        width: 290px;
        margin: 0 auto 15px;
        position: relative;
        input {
            width: 290px;
            height: 42px;
            padding: 0 20px 0 37px;
            border-radius: 2px;
            outline: 0;
            font-size: 14px;
            border: 1px solid #C4C6CC;
        }
    }
    .form-item-icon {
        position: absolute;
        top: 13px;
        left: 13px;
    }
    .form-submit {
        width: 290px;
        height: 42px;
        display: block;
        margin: 30px auto 0;
        background-color: #5c7ac6;
        border-radius: 2px;
        outline: 0;
        border: none;
        font-size: 14px;
        line-height: 18px;
        letter-spacing: 0;
        color: #fff;
        cursor: pointer;
    }
}

@media screen and (max-width: 1280px) {
    .wrapper {
        background-image: url(./assets/login_bg_1280.png);
    }
}
</style>
