<template>
  <div id="app">
    <Car />
  </div>
</template>

<script>
import Car from "@/components/Car"

export default {
  name: 'App',
  components: {
    Car:Car
  },
  created() {
    document.title = "Car Roster"
  },
}
</script>

<style>
html,body{
  padding: 0;
  margin: 0;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
