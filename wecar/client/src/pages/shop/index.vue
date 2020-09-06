<template>
  <div class="shop">
    <div class="shop-item" v-for="item in carList" :key="item.id">
      <div class="shop-car">
        <img src="@/assets/logo.png" alt="" />
        <span>{{ item.brand }}·{{ item.model }}</span>
        <span>{{ item.seller }}</span>
        <span>{{ item.age }}年</span>
        <span>{{ item.price }}元</span>
        <span class="shop-car__detail" @click="expand(item.id)"
          >详情<el-icon
            class="el-icon-caret-top"
            :class="item.id == expandId ? 'r180' : ''"
        /></span>
        <el-button type="primary" size="mini">购买</el-button>
      </div>
      <div
        class="shop-expand"
        :class="expandId == item.id ? '' : 'shop-expand--hide'"
      >
        <Expand />
      </div>
    </div>
  </div>
</template>

<script>
import Expand from "@/components/Expend.vue"
import config from "@/scripts/config.js";
export default {
  layout: "main",
  components: {Expand},
  data() {
    return {
      expandId: "",
    };
  },
  computed: {
    carList() {
      return config.carList;
    },
  },
  methods: {
    expand(id) {
      this.expandId == id ? (this.expandId = "") : (this.expandId = id);
    },
  },
};
</script>

<style lang="postcss" scoped>
.shop {
  overflow-y: hidden;
  &-car {
    position: relative;
    height: 120px;
    width: 100%;
    font-weight: bold;
    align-items: center;
    justify-content: space-around;
    display: flex;
    img {
      width: 100px;
    }
    &__detail {
      cursor: pointer;
    }
  }
  &-expand {
    background: #eee;
    height: 300px;
    transition: height 0.5s ease-in-out, opacity 0.5s;
    &--hide {
      opacity: 0;
      height: 0;
    }
  }
}

.r180 {
  transform: rotate(180deg);
}
</style>