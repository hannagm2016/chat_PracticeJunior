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
      curUser : {}
  },
 getters : {
   isLoggedIn: state => !!state.token,
   authStatus: state => state.status,
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
        state.currentUserChat = user;
      } else {
        state.currentUserChat = '';
      }
    },
      auth_request(state){
        state.status = 'loading'
      },
      auth_success(state, token, curUser){
        state.status = 'success'
        state.token = token
        state.user = curUser
      },
      auth_error(state){
        state.status = 'error'
      },
      logout(state){
        state.status = ''
        state.token = ''
      }
  },
  actions: {
      login({commit}, curUser){
        return new Promise((resolve, reject) => {
          commit('auth_request')
          axios({url: 'http://localhost:8080/login', data: curUser, method: 'POST' })
          .then(resp => {
          console.log (resp.data,"OOOOOO")
            const token = resp.data.token
            const curUser = resp.data.user
            localStorage.setItem('token', token)
            axios.defaults.headers.common['Authorization'] = token
            commit('auth_success', token, curUser)
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
          console.log(reject, "%%%%%")
          localStorage.removeItem('token')
          delete axios.defaults.headers.common['Authorization']
          resolve()
        })
       },
    FETCH_CONTACTS({commit}) {
      return axios.get('http://localhost:8080/contacts')
        .then((contacts) => {
          commit('SET_CONTACTS_TO_STORE', contacts.data)
        })
    },
    FETCH_CHATS({commit}) {
      return axios.get('http://localhost:8080/chats')
        .then((chats) => {
          commit('SET_CHATS_TO_STORE', chats.data)
        })
    },
    SET_USER_TO_HEADER({commit}, user) {
      commit('SET_USER_TO_HEAD', user)
    },
 /*  SEND_MSG_TO_CHAT({chat}) {
   console.log (chat, "tuituiuiuit")
      return axios.post('http://localhost:8080/message', chat)
        .then((response) => {
          return response;
        })
    }*/
  }
})

export default store;