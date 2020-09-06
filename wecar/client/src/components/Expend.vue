<template>
  <el-tabs class="expend">
    <el-tab-pane label="登记记录">
      <div class="expend-owner" v-for="item in carOwner" :key="item.id">
        <span>{{ item.id }}</span>
        <span>{{ item.time }}</span>
        <span>{{ item.from }}</span>
        <span>{{ item.to }}</span>
      </div>
    </el-tab-pane>
    <el-tab-pane label="护理记录">
      <div class="expend-nurse" v-for="item in carRepair" :key="item.id">
        <span>{{ item.id }}</span>
        <span>{{ item.time }}</span>
        <span>{{ item.fatory }}</span>
        <span>{{ item.detail }}</span>
      </div>
    </el-tab-pane>
  </el-tabs>
</template>

<script>
import config from "@/scripts/config.js";
export default {
  computed: {
    carOwner() {
      return config.carOwner.h_1;
    },
    carRepair() {
      return config.carRepair.r_1;
    },
  },
};
</script>

<style lang="postcss" scoped>
.expend {
  padding: 20px;
  &-owner, &-nurse {
    justify-content: space-between;
    display: flex;
    span {
      display: inline-block;
      flex: 1;
    }
  }
}
</style>