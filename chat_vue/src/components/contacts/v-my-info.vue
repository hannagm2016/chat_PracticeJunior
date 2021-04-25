<template>
  <div class='v-my-info'>
  <div class="info__content">
    <div class="info__avatar"></div>
      <div class="info__name">
          <form @submit.prevent="submit">
                   <h1 class="h3 mb-3 fw-normal">Edit your profile</h1>
                    <input v-model="Name" type="text" class="form-control" placeholder="Ann"  autofocus="">
                    <input v-model="Phone" type="text" class="form-control" placeholder="911">
                    <input v-model="Email" type="email" class="form-control" placeholder="ann@n.com" >
                    <input v-model="Password" type="password" class="form-control" placeholder="***" >
                   <button class="w-100 btn btn-lg btn-primary" type="submit">Save</button>
          </form>

      </div>
   </div>
      <!--<div class="hidden">
        <button  class="social" @click="EditProfile">Edit profile</button>
        <button  class="social" @click="SaveChanges">Save profile changes</button>
      </div>-->
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
.v-my-info{
    display: block;
    width:50%;
    top:4rem;
margin: 0 25%;
  }
  .info__avatar {
      position: fixed;
         width: 300px;
         height: 300px;
         background: #e7e7e7;
         border-radius: 50%;
  }
  .info__name{
   height: 200px;
       padding-left: 350px;

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
