<template>
  <el-menu mode="horizontal">
    <el-menu-item>
      <img class="nav-logo" src="@/assets/logo.png" alt="" />
    </el-menu-item>
    <el-menu-item v-for="item in navList" :key="item.id">{{
      item.name
    }}</el-menu-item>
  </el-menu>
</template>

<script>
export default {
  data() {
    return {
      navList: [
        { id: 1, name: "车商城" },
        { id: 2, name: "车管所" },
        { id: 3, name: "车护理" },
      ],
    };
  },
};
</script>

<style lang="postcss" scoped>
.nav {
  &-logo {
    width: 60px;
  }
}
</style>