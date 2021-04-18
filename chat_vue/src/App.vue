<template>
  <div id="app">
    <v-main/>
   </div>
</template>

<script>
    import vMain from './components/v-main'

    export default {
        name: 'App',
        components: {
            vMain
        }
    }
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 100px;
}

</style>
