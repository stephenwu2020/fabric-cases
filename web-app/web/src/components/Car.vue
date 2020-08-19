<template>
  <div class="car">
    <h2 class="car-head">CAR ROSTER</h2>
    <div class="car-list">
      <div class="car-item" v-for="item in carList" :key="item.Key">
        <img :src="carImgMap[item.Record.model] || carImgTest" alt="" class="car-item__img">
        <div class="car-item-intro">
          <div>
            <span class="car-item-intro__label">Model: {{item.Record.model}}</span>
            <span class="car-item-intro__label">Make: {{item.Record.make}}</span>
          </div>
          <div>
            <span class="car-item-intro__label">Colour: {{item.Record.colour}}</span>
            <span class="car-item-intro__label">Owner: {{item.Record.owner}}</span>
          </div>
        </div>
      </div>
    </div>
    <el-button class="car-add" type="info" icon="el-icon-plus" circle @click="onBtnAdd"></el-button>
    <Add ref="add"  @addEvent="queryAllCars"/>
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
        Prius: require("@/assets/cars/Prius.png"),
        Mustang: require("@/assets/cars/Mustang.png"),
        Tucson: require("@/assets/cars/Tucson.png"),
        Passat: require("@/assets/cars/Passat.png"),
        S: require("@/assets/cars/ModelS.png"),
        205: require("@/assets/cars/205.png"),
        S22L: require("@/assets/cars/S22L.png"),
        Punto: require("@/assets/cars/Punto.png"),
        Nano: require("@/assets/cars/Nano.png"),
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
  position: relative;
  &-head{
    position: fixed;
    z-index: 1;
    top: 0;
    left: 0;
    right: 0;
    padding: 10px ;
    margin: 0;
    text-align: left;
    background: #000;
    color: #fff;
  }
  &-list{
    padding-top: 60px;
    margin: auto;
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  &-item{
    position: relative;
    width: 100%;
    max-width: 100%;
    height: 60vh;
    min-height: 500px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    &__img{
      max-width: 50%;
    }
    &-intro{
      position: absolute;
      bottom: 0;
      width: 100%;
      display: flex;
      flex-wrap: wrap;
      justify-content: center;
      div{
        min-width: 200px;
        flex: 1;
        display: flex;
        span{
          line-height: 50px;
          text-align: center;
          flex: 1;
          color: #fff;
          background: rgb(87, 86, 86);
          font-weight: bold;
        }
      }
    }
  }
  &-add{
    position: fixed;
    right: 10%;
    bottom: 10%;
    font-size: 20px;
  }
}
@media screen and (max-width: 500px){
.car{
  &-item{
    &__img{
      max-width: 80%;
    }
  }
}
}
</style>