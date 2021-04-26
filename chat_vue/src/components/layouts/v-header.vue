<template>
  <div class='v-header'>
    <div class="v-back"
     v-if="!isCoreRoute"
     @click="routeBack">
      <i class="material-icons">keyboard_arrow_left</i>
         <span> Back </span>
     </div>
     <div class='v-link'>
       <router-link to="contacts">
       <i class="material-icons">call</i>
       <span class="mr-3">Contacts</span>
       </router-link>

      <router-link :to="{name: 'chat'}">
      <i class="material-icons">chat</i>
      <span class="mr-3">Chats</span>
      </router-link>
</div>
     <div class="v-name">
  <router-link :to="{name: 'myProfile'}">
          <span>{{user}}</span>
           <i class="material-icons">person_pin</i>
        </router-link>
    </div>
    <div class="right-side">

         <span v-if="isLoggedIn">
             <button @click="logout" class="btn btn-primary my-2 ml-2">Logout</button>
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
        'user',
        'token'
      ]),
      isLoggedIn : function()
      { return this.$store.getters.isLoggedIn},

       isCoreRoute() {
        return this.$route.path === '/authorization';
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
        flex-basis: 15%;
      }
      .v-name {
        flex-basis: 10%;
      }
      .right-side {
        flex-basis: 10%;

      }
      .v-link {
 text-align: start;
      flex-basis: 65%;

      }

</style>