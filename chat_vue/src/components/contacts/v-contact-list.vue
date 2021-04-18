<template>
  <div class='v-contact-list'>
    <v-contact
        v-for="contact in contacts"
        :key="contact.id"
        :contact_data="contact"
        @to-contact-info="toContactInfo(contact)"
    />
  </div>
</template>

<script>
  import vContact from './v-contact'
 // import {mapActions, mapState} from 'vuex'
import axios from 'axios'

  export default {
    name: "v-contact-list",
    components: {
      vContact
    },
    data: function(){
        return {
        contacts: []
        }
    },
    methods: {

      toContactInfo(contact) {
        this.$router.push({
          name: 'contact',
          query: {'id': contact.Id}
        });
      }
    },
    mounted() {
        axios.get('http://localhost:8080/contacts')
               .then((response)=>  {
                this.contacts = response.data;
                console.log(this.contacts)
            })
    }
  }
</script>

<style>

</style>
