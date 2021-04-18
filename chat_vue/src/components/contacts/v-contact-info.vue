<template>
  <div class='v-contact-info'>
    <div class="info__avatar"></div>
    <div class="info__content">
      <div class="info__name">
        <span>{{contact_info.Phone}}</span>
      </div>
      <div class="info__tools">
        <button class="start-call">Call</button>
        <button class="start-chat" @click="checkChats">
          Start chat
        </button>
      </div>
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
        'SET_USER_TO_HEADER'
      ]),
      toUserChat() {
        this.chats.map((chat) => {
          if (chat.id === this.contact_info.id) {
            this.$router.push({
              name: 'user',
              params: {'messages': chat.chat, 'user': chat},
              query: {'id': this.contact_info.id}
            })
          }
        })
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
      }
    },
    mounted() {
      this.contacts.find((contact) => {
        if (contact.Id === this.$route.query.id) {
          this.contact_info = contact
        }
      })
    }
  }
</script>

<style>
.info__avatar {
    width: 100%;
    height: 300px;
    background: #e7e7e7;
  }
.info__content {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-top: 5px;
  }

.start-chat {
  background: cornflowerblue;
  padding: $spacer $spacer*2;
  border-radius: 4px;
  color: #ffffff;
  outline: transparent;
}

.start-call {
  background: forestgreen;
  padding: $spacer $spacer*2;
  border-radius: 4px;
  color: #ffffff;
  outline: transparent;
}

</style>
