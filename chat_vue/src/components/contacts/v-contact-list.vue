<template>

<div>
  <div class='v-contact-list'>
  <div class='friends'>
  <span>Friends</span>
    <v-contact
        v-for="contact in contactFriends"
        :key="contact.id"
        :contact_data="contact"
        @to-contact-info="toContactInfo(contact)"
    />
    </div>
     <div class='black_list'>
          <span>Black List</span>
           <v-contact
                  v-for="contact in contactBlocked"
                  :key="contact.id"
                  :contact_data="contact"
                  @to-contact-info="toContactInfo(contact)"
              />
     </div>
       <div class='other_list'>
               <span>all Others</span>
                <v-contact
                       v-for="contact in contactNone"
                       :key="contact.id"
                       :contact_data="contact"
                       @to-contact-info="toContactInfo(contact)"
                   />
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
          ]),
           contactFriends() {
                  return this.$store.getters.contactFriends
           },
           contactBlocked() {
                  return this.$store.getters.contactBlocked
           },
           contactNone() {
                  return this.$store.getters.contactNone
           }
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
