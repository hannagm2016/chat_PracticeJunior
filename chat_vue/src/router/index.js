import Vue from 'vue'
import Router from 'vue-router'
import store from '../store/index.js'

import vContactList from '../components/contacts/v-contact-list'
import vContactUserInfo from '../components/contacts/v-contact-info'
import vChat from '../components/chats/v-chat'
import Authorization from '@/components/Auth/Authorization'
import Registration from '@/components/Auth/Registration'

Vue.use(Router);

/*const ifNotAuthenticated = (to, from, next) => {
  if (!store.getters.isAuthenticated) {
    next();
    return;
  }
  next("/");
};

const ifAuthenticated = (to, from, next) => {
  if (store.getters.isAuthenticated) {
    next();
    return;
  }
  next("/authorization");
};*/
let router = new Router({

//export default new Router({
 // mode: 'history',
  routes: [
    {
      path: '/',
      name: 'contact',
      component: vContactUserInfo,
      props: true
    },
    {
      path: '/contacts',
      name: 'contacts',
      component: vContactList,
         meta: {
                    requiresAuth: true
                  }
    },
    {
      path: '/chat',
      name: 'chat',
      component: vChat,
      props: true,
//      beforeEnter: ifAuthenticated

       meta: {
              requiresAuth: true
            }
    },
       {
              path: '/authorization',
              name: 'Authorization',
              //props: true,
              component: Authorization,

            //  beforeEnter: ifNotAuthenticated

          },
          {
             path: '/registration',
             name: 'Registration',
            // props: true,
             component: Registration
          }
  ]
})

router.beforeEach((to, from, next) => {
  if(to.matched.some(record => record.meta.requiresAuth)) {
    if (store.getters.isLoggedIn) {
      next()
      return
    }
    next('/authorization')
  } else {
    next()
  }
})

export default router
