<template>
  <div class="v-main">
    <v-header/>
    <router-view></router-view>
  </div>
</template>

<script>
import vHeader from './layouts/v-header'
export default {
  name: 'v-main',
  components: {
    vHeader
  }
}
</script>

<style>

.v-main {
  padding: 10 px;
  margin: 0 auto;

}

</style>