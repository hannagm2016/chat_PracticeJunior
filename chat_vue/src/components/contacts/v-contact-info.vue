<template>
  <div class='v-contact-info'>
    <div class="info__avatar"></div>
    <div class="info__content">
      <div class="info__name">
        <span>{{contact_info.Phone}}</span>
      </div>
      <div class="info__tools">
        <button class="start-call">Add to Black list</button>
        <button class="start-call">Add to Friends List</button>
        <button class="start-chat" @click="checkChats">
          Start chat
        </button>
      </div>
    </div>
  </div>
</template>

<script>

  export default {
    name: "v-contact-user-info",
    data() {
      return {
        contact_info: {}
      }
    },
    methods: {
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
    },
    /*mounted() {
      this.contacts.find((contact) => {
        if (contact.Id === this.$route.query.id) {
          this.contact_info = contact
        }
      })
    }*/
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
