<template>
  <div class='v-header'>
    <div class="v-back"
    v-if="!isCoreRoute"
     @click="routeBack"
    >
      <span> -Back- </span>
    </div>
    <div class="v-name"     v-if="!isCoreRoute">
      <span>TOM</span>
    </div>
    <div class="right-side"></div>
  </div>
</template>

<script>
 // import {mapState} from 'vuex'
  export default {
    name: "v-header",
   /* computed: {
      ...mapState([
        'currentUserChat'
      ]),*/
      computed:{
      isCoreRoute() {
        return this.$route.path === '/';
      }
    },
    methods: {
      routeBack() {
        this.$router.go(-1);
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