<template>
  <div class='v-contact-info' v-if="contact_info.Name" >
  <div class="info__content">
    <div class="info_avatar"></div>
      <div class="info_name">
           <span>{{contact_info.Name}}<br/></span>
           <span>{{contact_info.Relation}}<br/></span>
           <span>{{contact_info.Phone}}</span>
           <span>{{contact_info.Email}}</span>
      </div>
   </div>
      <div class="info__tools">
        <button v-if="contact_info.Relation == 'Friend'" class="social" @click="addToBlackList">Add to Black list</button>
        <button v-if="contact_info.Relation == ''" class="social" @click="addToFriendsList">Add to Friends List</button>
        <button v-if="contact_info.Relation == 'Friend'" class="social" @click="removeFriendsList">Remove from Friends List</button>
        <button v-if="contact_info.Relation == 'Blocked'" class="social" @click="removeBlackToFriends">Remove from Black to Friends</button>
        <button v-if="contact_info.Relation == 'Friend'" class="start-chat" @click="checkChats">
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
          'FETCH_CONTACTS',
          'SEND_MSG_TO_CHAT',
          'SET_RELATION',
          'DELETE_RELATION',
          'CHANGE_RELATION'
         ]),
         addToBlackList(){
               let rel={Relation: "Blocked",
                UserTo: this.contact_info.Id}
                console.log(rel),
                this.CHANGE_RELATION({relation: rel})
                this.FETCH_CONTACTS()
         },
         addToFriendsList(){
            let rel={Relation: "Friend",
             UserTo: this.contact_info.Id}
             console.log(rel),
             this.SET_RELATION({relation: rel}),
             this.FETCH_CONTACTS()
         },
          removeFriendsList(){
            let rel={
             UserTo: this.contact_info.Id}
             console.log(rel),
             this.DELETE_RELATION({relation: rel})
             this.FETCH_CONTACTS()
         },
          removeBlackToFriends(){
            let rel={Relation: "Friend",
             UserTo: this.contact_info.Id}
             console.log(rel),
             this.CHANGE_RELATION({relation: rel})
             this.FETCH_CONTACTS()
         },
         checkChats() {
                 if (!this.chats.length) {
                   this.FETCH_CHATS()
                     .then(() => {
                       this.toUserChat()
                     })
                 } else {
                   this.toUserChat()
                 }
               },
         toUserChat() {
         let index=false
           this.chats.map((chat) => {
             if (chat.Id === this.contact_info.Id) {
             index=true
                 console.log (chat, "HHHH", this.contact_info.Id, index)}
                 })
                 if (!index) {
              this.chat = {
              Text: 'hi, '+this.contact_info.Name,
              UserId: this.contact_info.Id,
              Type: "own"
              }
              console.log(this.chat, "UUUUUUU")
              this.SEND_MSG_TO_CHAT({chat: this.chat})
              }
             console.log("OOOOO"),
             this.$router.push({
             name: 'chat',
            // params: {'messages': chat.Chat, 'user': chat},
           // query: {'id': this.contact_info.Id}
                })
           }
       },
       watch: {
           $route(to, from) {
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
  .info_avatar {
      position: fixed;
         width: 200px;
         height: 200px;
         background: #e7e7e7;
         border-radius: 50%;
  }
  .info_name{
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
