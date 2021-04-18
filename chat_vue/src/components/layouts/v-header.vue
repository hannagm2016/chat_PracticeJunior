<template>
  <div class='v-header'>
    <div class="v-back"
     v-if="!isCoreRoute"
     @click="routeBack"
     >
         <span> -Back- </span>
    </div>
    <div class="v-name" v-if="!isCoreRoute">
        <span>{{this.AuthorizedUser}}</span>
    </div>
    <div class="right-side">
         <span v-if="isCookie">
              <a  href="/logout" class="btn my-2">Logout</a>
             <button @:click="logout" class="btn btn-primary my-2">Logout</button>
        </span>
         <a v-else href="/#/authorization" class="btn btn-primary my-2">Authorization</a>
     </div>
 </div>

</template>

<script>
 // import {mapState} from 'vuex'
 import VueCookies from 'vue-cookies';

  export default {
    name: "v-header",
   /* computed: {
      ...mapState([
        'currentUserChat'
      ]),*/
      computed:{
      isCoreRoute() {
        return this.$route.path === '/';
      },
      isCookie(){
        return VueCookies.isKey('Token')
      }
    },
    methods: {
      routeBack() {
        this.$router.go(-1);
      },
        logout() {
                  fetch("http://localhost:8080/logout", {
                      method:"GET"
                  })
                  console.log("llog")
                   .then(response => {
                       console.log(response)
                        VueCookies.remove('Token');
                       window.location = '/'

                   })

              },
    }
  }
</script>

<style>
    .v-header {
      height: 50px;
      display: flex;
      align-items: center;
      justify-content: flex-start;
      padding: 8px;
      background: grey;
      position: fixed;
      top: 0;
      right: 0;
      left: 0;
      z-index: 1;
      }
      .v-back {
        display: flex;
        align-items: center;
        flex-basis: 25%;
      }
      .v-name {
        flex-basis: 50%;
      }
      .right-side {
        flex-basis: 25%;
      }

</style>