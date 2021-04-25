<template>
  <div class='v-my-info'>
  <div class="info__content">
    <div class="info__avatar"></div>
      <div class="info__name">
          <form @submit.prevent="UpdateInfo">
                   <h1 class="h3 mb-3 fw-normal">Edit your profile: </h1>
                   <span>Name: </span>
                    <input v-model="myContact.Name" type="text" class="form-control mb-3 fw-normal">
                    <span>Phone: </span>
                    <input v-model="myContact.Phone" type="text" class="form-control mb-3" >
                    <span>Email: </span>
                    <input v-model="myContact.Email" type="email" class="form-control mb-3" >
                    <span>Password: </span>
                    <input v-model="myContact.Password" type="password" class="form-control mb-3" >

                   <button class="w-100 btn btn-lg btn-primary mb-3" type="submit">Save</button>
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
    name: "v-my-info",
    data() {
      return {
        my_info: {
            Id:'',
            Name:'',
            Phone: '',
            Email: '',
            Password:'',
            Status:''
        }
      }
    },
      computed: {
         ...mapState([
           'myContact'
         ])
       },
       methods: {
         ...mapActions([
          'FETCH_MY_CONTACT',
          'SEND_MY_INFO'
         ]),
         UpdateInfo(){
         this.my_info={
            Name: this.myContact.Name,
            Email: this.myContact.Email,
            Phone: this.myContact.Phone,
            Password: this.myContact.Password,
         }
             this. SEND_MY_INFO ({myContact: this.my_info})
             .then (response => {
                         console.log (response, this.my_info)
                        alert ("Saved: " + this.myContact.Name)
             })
    //          window.location = '/'
         }
},
 mounted() {
          this.FETCH_MY_CONTACT()
        }
}
</script>

<style>
.v-my-info{
    display: block;
    width: 85%;
    top: 4rem;
    margin: 0 10%;
  }
  .info__avatar {
         width: 250px;
         height: 250px;
         background: #e7e7e7;
         border-radius: 50%;
  }
  .info__name{
   height: 200px;
   padding-left: 350px;
  }
.social {
  background: forestgreen;
  margin: 50px 5px;
  border-radius: 4px;
  color: #ffffff;
  outline: transparent;
}

</style>
