<template>
<div>
  <div class='chat-list'>
      <div
          v-for="chat in chats"
          :key="chat.Id"
      >
      <div  class='v-user' @click="toUserChat(chat)">
          <div class="v-user__avatar"></div>
          <div class="v-user__info">
            <p class="info">{{chat.Name}}</p>
            <p class="info__last-message">{{chat.Chat[chat.Chat.length-1].Text}}</p>
          </div>
          <div class="v-user__time">{{chat.Chat[chat.Chat.length-1].Time}}</div>
        </div>
    </div>
</div>
  <div class='v-user-chat'>
<span>{{name}}</span>
  <hr/>
    <v-message
        v-for="message in messages"
        :key="message.Id"
        :message="message"
    />
    <div class="input_field">
      <input
          type="text"
          class="v-user-chat__textfield"
          v-model="textValue"
          @keypress.enter="sendMessage">
      <i class="material-icons"
          @click="sendMessage">
        send
      </i>
    </div>
  </div>
  </div>
</template>



<script>
  import vMessage from './v-message'
  import {mapActions, mapState} from 'vuex'

  export default {
    name: "v-user-chat",
    components: {
      vMessage,
    },

    data() {
      return {
      userTo: null,
      name: 'Please select contact to start chat',
        messages:[],
        connection: null,
        textValue: '',
        chat: null,
        serverUrl: "ws://localhost:8080/socket"
      }
    },
  computed: {
      ...mapState([
        'chats',
        'userId'
       ])
    },
     mounted: function() {
          this.connectToWebsocket()
          this.FETCH_CHATS()
          this.defaultChat()
        },
    methods: {
    defaultChat(){
          this.messages = this.chats[0].Chat
          this.userTo = this.chats[0].Id
          this.name = this.chats[0].Name
    },

     ...mapActions([
            'FETCH_CHATS',
            'SEND_MSG_TO_CHAT'
          ]),
    toUserChat(chat) {
          this.messages = chat.Chat
            this.userTo = chat.Id
            this.name = chat.Name
           /* this.$router.push({
            query: {'id': this.userTo}
          });*/
          },
      sendMessage() {
        this.chat = {
                  Text: this.textValue,
                 UserId: this.userTo,
                 Type: "own"
               }
               this.messages.push(this.chat)
            this.connection.send(this.chat.Text);
             this.SEND_MSG_TO_CHAT({chat: this.chat})
                .then((response) => {
                  console.log (response);
                  this.textValue = ''
                })
                 this.FETCH_CHATS()
      },
            connectToWebsocket() {
                   this.connection = new WebSocket( this.serverUrl);
                   this.connection.addEventListener('open', (event) => { this.onWebsocketOpen(event) });
                   this.connection.addEventListener('message', (event) => { this.handleNewMessage(event) });
                 },
             onWebsocketOpen() {
                    console.log("connected to WS!");
                  },

              handleNewMessage(event) {
               this.FETCH_CHATS()
                console.log(event.data)
               //this.messages = this.chat.Chat
                //  let data = event.data;
               //    data = data.split(/\r?\n/);
               //    this.messages.push({Text: event.data});
                  }
    }
  }
</script>

<style>
.chat-list{
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


.v-user-chat {
    display: block;
    padding-left: 300px;
    top:4rem;
    margin-bottom: 200px;
  }
  .input_field {
    display: flex;
    align-items: center;
    position: fixed;
    bottom: 70px;
    right: 0;
   left: 25rem;
}
    .material-icons {
      font-size: 36px;
    }

  .v-user-chat__textfield {
    width: 80%;
    padding: 16px;
    border: 0;
    box-shadow: inset 0 0 5px 1px #bdbdbd;
    outline: transparent;
    border-radius: 8px;
    font-size: 20px;
  }
.v-user {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20px;
}
 .v-user__avatar {
    width: 50px;
    height: 50px;
    background: #e7e7e7;
    border-radius: 50%;
  }

  .v-user__info {
    text-align: left;
    max-width: 170px;
    flex-basis: 50%;
  }
    .info {
      font-weight: bold;
    }
    .info__last-message {
      white-space: nowrap;
      text-overflow: ellipsis;
      max-width: 150px;
      overflow: hidden;
    }
</style>
