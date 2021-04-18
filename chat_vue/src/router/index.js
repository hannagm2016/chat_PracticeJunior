import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router);

import vContactList from '../components/contacts/v-contact-list'
import vContactUserInfo from '../components/contacts/v-contact-info'
import vChat from '../components/chats/v-chat'

let router = new Router({
  routes: [
    {
      path: '/',
      name: 'contacts',
      component: vContactList
    },
    {
      path: '/contact',
      name: 'contact',
      component: vContactUserInfo,
      props: true
    },
    {
      path: '/chat',
      name: 'chat',
      component: vChat,
      props: true
    }
  ]
})

export default router;
