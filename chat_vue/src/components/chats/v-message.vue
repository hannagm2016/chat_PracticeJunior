<template>
  <div :class="className">
    <div class="v-message__content">
      <p>{{message.Text}}</p>
      <p class="content__time">{{message.Time}}</p>
    </div>
  </div>
</template>

<script>
  export default {
    name: "v-message",
    props: {
      message: {
        type: Object,
        default: () => {}
      }
    },
    data() {
      return {}
    },
    computed: {

      className() {
        return {
          'chat__own': this.message.Type === 'own',
          'chat__other': this.message.Type === 'other'
        }
      }
    }
  }
</script>

<style>
  .chat__own{
    margin: 20px;
    display: flex;
    justify-content: flex-end;
}
  .chat__other{
      margin: 20px;
    display: flex;
    justify-content: flex-start;
}
  .v-message__content {
    background: #bad3ff;
    border-radius: 20px;
    padding: 0 20px;
    position: relative;
  }
  .content__time{
      position: absolute;
      bottom: -35px;
      right: 0px;
    }

</style>
