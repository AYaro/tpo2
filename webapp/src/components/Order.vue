<template>
  <div class="container"> 
    <h1>Заказ {{ order.ID }}</h1>
    <div class="centerwrapper">
      <div v-if="!order.finished && !showError">
        Добавить блюдо
        <div>
          <input class = "text-input" v-model="dish.name" placeholder="НАЗВАНИЕ"><br/>
        </div>
        <div>
          <input class = "text-input" v-model.number="dish.price" placeholder="СТОИМОСТЬ"><br/>
        </div>
        <select class = "text-input" v-model="dish.users" multiple="true">
          <option v-for="user in order.users" v-bind:key="user">
            {{ user }}
          </option>
        </select>
        <div>
          <a class="button" v-on:click="addDish(dish)">Добавить</a>
        </div>
        <div class="button dish">
          <ul>
            <li class="button dish" v-for="dish in order.dishes" v-bind:key="dish">
              Назвние: <span class="green">{{ dish.name }}</span><br/>
              Цена: <span class="green">{{ dish.price }}</span><br/>
              <div>
                  Кто заказал: <div class="green" v-for="user of dish.users" v-bind:key="user">{{user}}</div>
              </div>
            </li>
          </ul>
        </div>
        <div>
          <a class="button" v-on:click="finishOrder(true)">Завершить заказ</a>
        </div>
      </div>
      <div v-if="order.finished && !showError" class="button">
        Кто сколько платит?
        <ul>
          <li class="button dish" v-for="user in order.result" v-bind:key="user">
            Имя: <span class="green">{{ user.name }}</span><br/>
            Сумма: <span class="green">{{ user.sum }}</span><br/>
           </li>
        </ul>
      </div>
      <div v-if="showError" class="button">
        Oшибка: {{error}}
      </div> 
    </div>
  </div>
</template>

<script>
import axios from 'axios'
// import Multiselect from 'vue-multiselect'

export default {
//  components: { Multiselect },  
  name: 'Order',
  props: ['orderID'],
  data: () => ({
    order: {
        ID: "",
        currency: "",
        users: [],
        dishes: [],
        finished: false,
        result: {}
    },
    dish: {
        name: "",
        price: 0,
        users: [],
    },
    error: "",
    showError: false,
  }),
  methods: {
    greet: function () {
      // `this` внутри методов указывает на экземпляр Vue
      alert('Привет!')
    },
    addDish: function (dish) {
      const self = this
      axios.post('http://localhost:8080/api/order/' + this.order.ID + "/addDish", dish)
      .then(function (response) {
        console.log(response);
        self.order = response.data
      })
    },
    finishOrder: function (isFinished) {
      const self = this
      axios.post("http://localhost:8080/api/order/" + this.order.ID + "/finishOrder", isFinished)
      .then(function (response) {
        console.log(response);
        self.order = response.data
      })
    }
  },
  created: function () {
    const self = this
    const orderId = this.$route.params.orderId
    axios.get('http://localhost:8080/api/order/'+ orderId)
    .then(function (response) {
        self.order = response.data
        self.order.ID = orderId
        console.log(self.order);
    })
    .catch(function (error) {
        console.log(error);
        self.showError = true
        self.error = error.message
    });
  },
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
  margin: 20px;
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

.dish {
  text-align: left;
}

.green {
  color: #bbddbb
}
</style>
