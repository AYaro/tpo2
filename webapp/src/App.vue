<template>
  <div id="app">
    <h2 class="title">tpo lab 2</h2>
    <div>
      <router-view/>
    </div>
  </div>
</template>

<script>
// import Tech from './components/Tech.vue'

export default {
  name: 'App',
  components: {
//    Tech
  },
  data() {
    return {
      logoSVG: require('./assets/logo.svg')
    }
  }
};
</script>

<style>
body {
  margin-top: 5%;
  padding-right: 5%;
  padding-left: 5%;
  font-size: larger;
  font-family: 'Trebuchet MS', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  background: #505050;
}

@media screen and (min-width: 800px) {
  body {
    padding-right: 15%;
    padding-left: 15%;
  }
}

@media screen and (min-width: 1600px) {
  body {
    padding-right: 30%;
    padding-left: 30%;
  }
}

code {
  font-family: source-code-pro, Menlo, Monaco, Consolas, "Courier New",
    monospace;
  background-color: #b3e6ff;
}

.title {
  color: white;
  text-align: center;
}

.logo {
  text-align: center;
}
</style>
