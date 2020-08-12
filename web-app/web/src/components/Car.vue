<template>
  <div class="car">
    <h1 class="car-head">Car Roster</h1>
    <div class="car-list">
      <div class="car-item" v-for="item in carList" :key="item.Key">
        <img :src="carImgMap[item.Record.model] || carImgTest" alt="" class="car-item__img">
        <div class="car-item-intro">
          <span class="car-item-intro__label">Model: {{item.Record.model}}</span>
          <span class="car-item-intro__label">Make: {{item.Record.make}}</span>
          <span class="car-item-intro__label">Colour: {{item.Record.colour}}</span>
          <span class="car-item-intro__label">Owner: {{item.Record.owner}}</span>
        </div>
      </div>
    </div>
    <el-button class="car-add" type="primary" icon="el-icon-plus" circle @click="onBtnAdd"></el-button>
    <Add ref="add" @addEvent="queryAllCars"/>
  </div>
</template>

<script>
import Add from "@/components/Add"
export default {
  components: {Add},
  data() {
    return {
      carList: []
    }
  },
  computed: {
    carImgMap(){
      return {
        Prius: require("@/assets/cars/Prius2.jpeg"),
        Mustang: require("@/assets/cars/Mustang.png"),
        Tucson: require("@/assets/cars/Tucson.jpeg"),
        Passat: require("@/assets/cars/Passat.png"),
        S: require("@/assets/cars/ModelS.jpg"),
        205: require("@/assets/cars/205.png"),
        S22L: require("@/assets/cars/S22L.jpg"),
        Punto: require("@/assets/cars/Punto.jpeg"),
        Nano: require("@/assets/cars/Nano.jpg"),
        Barina: require("@/assets/cars/Barina.png"),
      }
    },
    carImgTest() {
      return require("@/assets/cars/test.png")
    }
  },
  created() {
    this.queryAllCars()
  },
  methods: {
    queryAllCars(){
      this.$axios.post("/queryAllCars")
        .then(res => {
          let tmp = res.data
          tmp.sort((a, b) => {
            let aid = a.Key.slice(3)
            let bid = b.Key.slice(3)
            return Number(bid) - Number(aid)
          })
          this.carList = tmp
      })
    },
    onBtnAdd(){
      this.$refs.add.show()
    } 
  },
}
</script>

<style lang="postcss" scoped>
.car{
  &-list{
    margin: auto;
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    padding: 0 5%;
  }
  &-item{
    width: 500px;
    padding: 20px;
    display: flex;
    &__img{
      width: 300px;
    }
    &-intro{
      flex: 1;
      display: flex;
      flex-direction: column;
      align-items: flex-start;
    }
  }
  &-add{
    position: fixed;
    right: 10%;
    bottom: 10%;
    font-size: 20px;
  }
}
</style>