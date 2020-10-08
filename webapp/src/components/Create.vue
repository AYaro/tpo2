<template>
  <div class="container"> 
    <div class="centerwrapper">
      <h1>Введите валюту</h1>
      <div>
        <input class = "text-input" v-model="check.currency" placeholder="ВАЛЮТА"><br/>
      </div>
      <h1>Выберите режим</h1>
      <div>
        <select class = "text-input" v-model="check.mode">
            <option>Рандомный</option>
            <option>Поровну</option>
            <option>Честный</option>
        </select>
      </div>
      <h1>Добавьте людей</h1>
      <input class = "text-input" v-model="name" placeholder="ИМЯ">
      <a class="button" v-on:click="addUser(name)">+</a>
      <div v-for="user in check.users"  :key="user" >
        <div class="name"> 
          {{ user }}
        </div>
      </div>
      <a class="button" v-on:click="createOrder(check)">Создать заказ</a>
      <div class="button" v-if="orderID.length > 0">
        Заказ {{ orderID }} создан!
         <router-link  class="button" :to="'/order/' + orderID"  event>Проложить</router-link>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Create',
  data: () => ({
    check: {
      currency: "",
      mode:"",
      users: [],
    },
    orderID: "",
    name:""
  }),
  methods: {
    createOrder: function (check) {
      const self = this
      axios.post('http://localhost:8080/api/order', check)
      .then(function (response) {
        console.log(response);
        self.orderID = response.data
      })
    },
    greet: function () {
      // `this` внутри методов указывает на экземпляр Vue
      alert('Привет!')
    },
    addUser: function (str) {
      if (str != "") {
        this.check.users.push(str)
      }
    }
  }
};
</script>

<style scoped>
.technologies {
  margin-top: 5px;
}
.text-input{
  padding: 40px 60px;
  text-align: center;
  border: none;
  color: #d5d5d5;
  display: inline-block;
  font-size: 40px;
  border-radius: 63px;
  background: linear-gradient(45deg, #484848, #565656);
  box-shadow:  12px -12px 29px #363636, 
              -12px 12px 29px #6a6a6a;
}
.button {
  font-family: 'Trebuchet MS', sans-serif;
  border: none;
  color: #d5d5d5;
  padding: 40px 60px;
  margin: 20px;
  text-align: center;
  display: inline-block;
  font-size: 40px;
  border-radius: 75px;
  background: linear-gradient(45deg, #565656, #484848);
  box-shadow:  12px -12px 29px #3a3a3a, 
              -12px 12px 29px #666666;
}

.name {
  font-family: 'Trebuchet MS', sans-serif;
  border: none;
  color: #d5d5d5;
  padding: 20px 40px;
  margin: 20px;
  text-align: center;
  display: inline-block;
  font-size: 40px;
  border-radius: 75px;
  background: linear-gradient(45deg, #565656, #484848);
  box-shadow:  12px -12px 29px #3a3a3a, 
              -12px 12px 29px #666666;

}

.container{
  text-align:center;
  width: 800pt;
  height: 800pt;
}
.centerwrapper{
  margin: auto;
  width: 600pt;
}

h1 {
  color: #d5d5d5;
}
</style>
