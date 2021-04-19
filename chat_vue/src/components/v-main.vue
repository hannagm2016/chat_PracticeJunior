<template>
  <div class="v-main">
    <v-header/>
    <router-view></router-view>
     <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css">
  </div>
</template>

<script>
import vHeader from './layouts/v-header'
export default {
  name: 'v-main',
  components: {
    vHeader
  },
}
</script>

<style>

.v-main {
  padding: 10 px;
  margin: 0 auto;

}

</style>