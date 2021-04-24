import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex);

let store = new Vuex.Store({
  state: {
    contacts: [],
    chats: [],
    currentUserChat: {},
      status: '',
      token: localStorage.getItem('token') || '',
      user : localStorage.getItem('user') || '',
      userId : localStorage.getItem('userId') || ''
  },
  mutations: {
    SET_CONTACTS_TO_STORE(state, contacts) {
      state.contacts = contacts;
    },
    SET_CHATS_TO_STORE(state, chats) {
      state.chats = chats;
    },
    SET_USER_TO_HEAD(state, user) {
      if (user) {
        state.user = user;
      } else {
        state.user = '';
      }
    },
      auth_request(state){
        state.status = 'loading'
      },
      auth_success(state, token, user, userId){
        state.status = 'success'
        state.token = token
        state.user = user
        state.userId = userId
      },
      auth_error(state){
        state.status = 'error'
      },
      logout(state){
        state.status = ''
        state.token = ''
        state.user = ''
      }
  },
  actions: {
      login({commit}, user){
        return new Promise((resolve, reject) => {
          commit('auth_request')
          axios({url: 'http://localhost:8080/login', data: user, method: 'POST' })
          .then(resp => {
            const token = resp.data.token
            const user = resp.data.user.Name
            const userId = resp.data.user.Id

            localStorage.setItem('token', token)
            localStorage.setItem('user', user)
            localStorage.setItem('userId', userId)
            console.log(localStorage, "hhhhhhhh", userId)
            axios.defaults.headers.common['Authorization'] = token
            axios.defaults.headers.common['Authorization'] = user
            commit('auth_success', token, user, userId)
            resolve(resp)
          })
          .catch(err => {
            commit('auth_error')
            localStorage.removeItem('token')
            reject(err)
          })
        })
      },
       logout({commit}){
        return new Promise((resolve, reject) => {
          commit('logout')
          console.log(reject, "%%%%%", this.state.user)
          localStorage.removeItem('token')
          localStorage.removeItem('user')
          delete axios.defaults.headers.common['Authorization']
          resolve()
        })
       },
    FETCH_CONTACTS({commit}) {
      return axios.get('http://localhost:8080/contacts/'+this.state.userId)
        .then((contacts) => {
                console.log(contacts.data)

          commit('SET_CONTACTS_TO_STORE', contacts.data)
        })
    },
    FETCH_CHATS({commit}) {
      return axios.get('http://localhost:8080/chats/'+this.state.userId)
        .then((chats) => {
        console.log(chats.data)
          commit('SET_CHATS_TO_STORE', chats.data),
                    console.log(this.state.token, "UUUUser", this.state.user)

        })
    },
   /* SET_USER_TO_HEADER({commit}, user) {
      commit('SET_USER_TO_HEAD', user)
    },*/
   SEND_MSG_TO_CHAT({commit}, {chat}) {
   console.log (commit, chat, "tuituiuiuit")
chat.Time= new Date().toLocaleTimeString('en-US',
                       {
                         hour12: false,
                         hour: "numeric",
                         minute: "numeric"
                       }
                     )

      return axios.post('http://localhost:8080/message/'+ this.state.userId, chat)
        .then((response) => {
          return response;
        })
    }
  },
   getters : {
     isLoggedIn: state => !!state.token,
     authStatus: state => state.status,
     authUser: state => !!state.user
   }
})

export default store;