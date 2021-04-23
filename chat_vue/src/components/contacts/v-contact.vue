<template>
  <div class='v-contact' @click="toContactInfo">
    <div class="user__avatar"></div>
    <div class="user__name">
      <span>{{contact_data.Name}}</span>
    </div>
     <div class="user__status"
        :class="{ 'online': contact_data.Status === 'online',
         'offline': contact_data.Status === 'offline' }"
    ></div>
  </div>
</template>

<script>
  export default {
    name: "v-contact",
    props: {
      contact_data: {
        type: Object,
        default: () => {}
      }
    },
    data() {
      return {}
    },
    methods: {
      toContactInfo() {
        this.$emit('to-contact-info')
      }
    }
  }
</script>

<style>
.v-contact {
  display: flex;
  align-items: center;
  margin-bottom: 5px;
  padding: 5px;
  }
  .user__avatar {
      width: 50px;
      height: 50px;
      background: #e7e7e7;
      border-radius: 50%;
    }
    .user__name {
      font-weight: bold;
      text-align: left;
      flex-basis: 50%;
      padding-left: 15px;
    }
    .user__status {
      width: 10px;
      height: 10px;
      border-radius: 50%;
    }
    .user__status.online {
      background: green;
    }
    .user__status.offline {
      background: red;
    }


</style>
