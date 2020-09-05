<template>
  <el-menu mode="horizontal">
    <el-menu-item @click="jump('/')">
      <img class="nav-logo" src="@/assets/logo.png" alt="" />
    </el-menu-item>
    <el-menu-item
      v-for="item in navList"
      :key="item.id"
      @click="jump(item.link)"
      >{{ item.name }}</el-menu-item
    >
  </el-menu>
</template>

<script>
export default {
  data() {
    return {
      navList: [
        { id: 1, name: "车商城", link: "/shop" },
        { id: 2, name: "车管所", link: "/office" },
        { id: 3, name: "车护理", link: "/nurse" },
      ],
    };
  },
  methods: {
    jump(link) {
      if (this.$route.path == link) {
        this.$log.debug(`Skip: from ${this.$route.path} to ${link}`)
        return;
      }
      this.$router.push({ path: link });
    },
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