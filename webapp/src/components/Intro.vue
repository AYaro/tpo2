<template>
  <div class="container"> 
    <div class="centerwrapper">
      <h1>Введите номер заказа</h1>
      <div>
        <input class = "text-input" v-model="id" placeholder="Номер заказа"><br/>
      </div>
      <router-link  class="button" :to="'/order/' + id"  event>Найти</router-link>
      <h1>Или создайте новый</h1>
      <router-link  class="button" to="'/create"  event>Создать</router-link>
    </div>
  </div>
</template>

<script>
// import axios from 'axios'
//import Create from './Create'
//import Order from './Order'

export default {
  name: 'Intro',
  components: {
    //Create,
    //Order
  },
  data: () => ({
    id: "",
    order: {},
  }),
  methods: {
    greet: function () {
      // `this` внутри методов указывает на экземпляр Vue
      alert('Привет!')
      this.isHidden = true
    }
  }
};
</script>

<style scoped>
.technologies {
  margin-top: 5px;
}
.text-input{
  text-transform: uppercase;
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
  margin: 20pt;
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
  width:600pt;
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
