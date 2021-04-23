<template>
  <div class='v-header'>
    <div class="v-back"
     v-if="!isCoreRoute"
     @click="routeBack">
         <span> Back </span>
         <router-link :to="{name: 'contacts'}">
                   <i class="link">Contacts</i>
                 </router-link>

                 <router-link :to="{name: 'chat'}">
                   <i class="link">Chats</i>
                 </router-link>

    </div>
     <div class="v-name">

    </div>
    <div class="right-side">
     <span>{{user}} </span>
         <span v-if="isLoggedIn">
             <button @click="logout" class="btn btn-primary my-2">Logout</button>
        </span>
         <a v-else href="/#/authorization" class="btn btn-primary my-2">Authorization</a>
     </div>
 </div>

</template>

<script>
  import {mapState} from 'vuex'
// import VueCookies from 'vue-cookies';

  export default {
    name: "v-header",
    computed: {
      ...mapState([
        'user'
      ]),
      isLoggedIn : function()
      { return this.$store.getters.isLoggedIn},
       authUser : function()
      { return this.$store.getters.authUser},
       isCoreRoute() {
       console.log(this.user,"________")
        return this.$route.path === '/';
      },
     },
    methods: {
      routeBack() {
        this.$router.go(-1);
      },
       logout: function () {
              this.$store.dispatch('logout')
              .then(() => {
                this.$router.push('/authorization')
              })
            }
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
      margin: 10px;
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
      .link {
      margin: 10px;
      }

</style>