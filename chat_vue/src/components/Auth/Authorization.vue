<template>
<div>

<main class="form-signin">
<form @submit.prevent="submit">
         <h1 class="h3 mb-3 fw-normal">Please Authorize</h1>
          <input v-model="Name" type="text" class="form-control" placeholder="Email address" required="" autofocus="">
          <input v-model="Password" type="password" class="form-control" placeholder="Password" >

         <button class="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>

</form>
       <div class="text-muted py-2">
        <a v-bind:href="link.FB">
        <button class="w-45 btn  btn-secondary">Login with Facebook!</button> </a>
         <a v-bind:href= "link.Google">
        <button class="w-45 btn  btn-secondary">Login with Google!</button> </a>
</div>
    <div class="text-muted py-3">
        <a href="/#/registration">Registration</a>
    </div>

    </main>
  </div>
</template>

<script>
//import axios from 'axios'
//import VueCookies from 'vue-cookies';

export default {
   data: () => ({
        Name:'',
        Password: '', //'123',
        link: []
    }),
    mutations: {
      auth_request(state){
        state.status = 'loading'
      },
      auth_success(state, token, user){
        state.status = 'success'
        state.token = token
        state.user = user
      },
      auth_error(state){
        state.status = 'error'
      },
      logout(state){
        state.status = ''
        state.token = ''
        state.user = ''
      },
    },
      methods: {
                    submit() {
                      let Name = this.Name
                      let Password = this.Password
                      this.$store.dispatch('login', { Name, Password })

                     .then(() => this.$router.push('/contacts'),

                   //   VueCookies.set('Token' , response.data.Token, '1d');
                     )
                     .catch(err => console.log(err))

             },
            },
    mounted() {
       fetch("http://localhost:8080/authorization")
         .then(response =>response.json())
         .then((data)=>{
          this.link = data;
         })
}
}

</script>

<style>


.form-signin {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: auto;
  margin-top:100px;
 }
.form-signin .form-control {
  position: relative;
  box-sizing: border-box;
  height: auto;
  padding: 10px;
  font-size: 16px;
}
.form-signin .form-control:focus {
  z-index: 2;
}
.form-signin input[type="email"] {
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}
.form-signin input[type="password"] {
  margin-bottom: 10px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}

</style>
