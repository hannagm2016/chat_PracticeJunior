<template>
  <div class='v-contact-info'>
  <div class="info__content">
    <div class="info__avatar"></div>
      <div class="info__name">
           <span>{{contact_info.Name}}<br/></span>
           <span>{{contact_info.Phone}}</span>
      </div>
   </div>
      <div class="info__tools">
        <button class="social">Add to Black list</button>
        <button class="social">Add to Friends List</button>
        <button class="start-chat" @click="checkChats">
          Start chat
        </button>
      </div>
    </div>
</template>

<script>
  import {mapState, mapActions} from 'vuex'
  export default {
    name: "v-contact-user-info",
    data() {
      return {
        contact_info: {}
      }
    },
      computed: {
         ...mapState([
           'chats',
           'contacts'
         ])
       },
       methods: {
         ...mapActions([
          'FETCH_CHATS',
          'SEND_MSG_TO_CHAT'
         ]),
         checkChats() {
           this.chats.map((chat) => {
             if (chat.Id === this.contact_info.Id) {
                 console.log (chat, "HHHH", this.contact_info.Id)
              this.chat = {
              Text: 'hi, '+this.contact_info.Name,
              UserId: this.contact_info.Id,
              Type: "own"
              }
              }
              })
              if (this.chat.Text !='') {
          //  this.SEND_MSG_TO_CHAT({chat: this.chat})
             console.log (this.chat, "HHHH", this.contact_info.Id)
}
             console.log("OOOOO")
             this.$router.push({
             name: 'chat',
            // params: {'messages': chat.Chat, 'user': chat},
           // query: {'id': this.contact_info.Id}
                })
             }
       },
       watch: {
           $route(to, from) {
                console.log (to, from)
                this.contact_info = this.$route.params.contact
           }
     }
}
   /* methods: {
      checkChats() {
        this.chats.map((chat) => {
          if (chat.id === this.contact_info.id) {
            this.$router.push({
              name: 'chat',
            //  params: {'messages': chat.chat, 'user': chat},
          //    query: {'id': this.contact_info.id}
            })
          }
        })
      },
    }*/
</script>

<style>
.v-contact-info{
    display: block;
    padding-left: 350px;
    top:4rem;
    margin-bottom: 200px;
  }
  .info__avatar {
      position: fixed;
         width: 200px;
         height: 200px;
         background: #e7e7e7;
         border-radius: 50%;
  }
  .info__name{
   height: 200px;
       padding-left: 200px;
       align-items: center;

  }
.start-chat {
  background: cornflowerblue;
  margin: 50px 5px;
  border-radius: 4px;
  color: #ffffff;
  outline: transparent;
}
.social {
  background: forestgreen;
  margin: 50px 5px;
  border-radius: 4px;
  color: #ffffff;
  outline: transparent;
}

</style>
