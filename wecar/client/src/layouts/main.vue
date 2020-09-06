<template>
  <el-container>
    <el-header>
      <Navigator />
    </el-header>
    <el-main>
      <router-view />
    </el-main>
    <el-footer>
      <Footer />
    </el-footer>
  </el-container>
</template>

<script>
import Navigator from "@/components/Navigator.vue"
import Footer from "@/components/Footer.vue"
export default {
  components: {
    Navigator,
    Footer,
  } 
}
</script>

<style scoped>
.el-header{
  position: fixed;
  z-index: 999;
  width: 100%;
  padding: 0;
}
.el-main{
  padding: 0;
  margin-top: 60px;
  min-height: 85vh;
}
</style>