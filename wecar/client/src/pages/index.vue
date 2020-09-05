<template>
  <div class="home">
    <div class="home-intro">
      <h1>WeCar, 全球领先的二手车交易平台</h1>
      <img src="@/assets/logo.png" alt="" />
    </div>
    <div class="home-intro">
      <h1>植根区块链平台，追踪爱车前世今生</h1>
      <img src="@/assets/chain.png" alt="" />
    </div>
    <div class="home-intro">
      <h1>落地北上广深，销售网络遍布大江南北</h1>
      <img src="@/assets/global.png" alt="" />
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
export default {
  name: "home",
  layout: "main",
};
</script>

<style lang="postcss" scoped>
.home {
  &-intro {
    &:nth-child(2){
     background: #eee; 
    }
    padding: 50px 0;
    img {
      margin: 50px 0;
      width: 450px;
      max-width: 60%;
    }
  }
}
</style>