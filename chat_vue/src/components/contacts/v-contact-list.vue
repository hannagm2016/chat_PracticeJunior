<template>
<div>

  <div class='v-contact-list'>
  <div class='friends'>
  <span>Friends</span>
  </div>
    <v-contact
        v-for="contact in contacts"
        :key="contact.id"
        :contact_data="contact"
        @to-contact-info="toContactInfo(contact)"
    />
     <div class='black_list'>
          <span>Black List</span>
          </div>
  </div>

      <v-contact-user-info/>

</div>

</template>

<script>
  import vContact from './v-contact'
  import vContactUserInfo from './v-contact-info'
  import {mapActions, mapState} from 'vuex'

  export default {
    name: "v-contact-list",
    components: {
      vContact,
      vContactUserInfo
    },
    computed: {
          ...mapState([
            'contacts'
          ])
        },
        methods: {
          ...mapActions([
            'FETCH_CONTACTS',
         //   'SET_USER_TO_HEADER'
          ]),
          toContactInfo(contact) {
            this.$router.push({
             // name: 'contact',
              query: {'id': contact.Id},
              params: {'contact': contact},
            });
          //  this.SET_USER_TO_HEADER(contact.Name)
          }
        },
        mounted() {
          this.FETCH_CONTACTS()
        }
      }
</script>

<style>
.v-contact-list{
    font-size: 16px;
    background-color: #fff;
    width: 20rem;
    position: fixed;
    z-index: 10;
    margin: 0;
    top: 66px;
    left: 0;
    bottom: 0;
    box-sizing: border-box;
    border-right: 1px solid #eaecef;
    overflow-y: auto;
  }

</style>
