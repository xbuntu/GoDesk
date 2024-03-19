<template>
  <div class="menu">
    <div class="list">
      <div class="item" v-for="(item,index) in list" @click="toPage(item.name)" :class="{'active':routeName === item.name}">
        <span><i :class="item.icon"></i>{{ item.title }}</span>
      </div>
      <div class="item" @click="logout()">
        <span><i class="el-icon-error"></i>退出登录</span>
      </div>
    </div>
  </div>
</template>

<script>
import {removeToken} from "@/utils/token";

export default {
  name: 'Menu',
  props: {
    routeName: String
  },
  data() {
    return {
      list: [
        {
          name: 'home',
          icon: 'el-icon-s-help',
          title: '网址导航'
        }, {
          name: 'mock',
          icon: 'el-icon-share',
          title: '网络请求'
        },{
          name: 'photo',
          icon: 'el-icon-picture',
          title: '照片预览'
        },{
          name: 'system',
          icon: 'el-icon-connection',
          title: '系统接口'
        },{
          name: 'table',
          icon: 'el-icon-s-data',
          title: '表格模版'
        },
      ]
    }
  },
  methods: {
    toPage(name) {
      this.$router.push({
        name
      });
    },
    logout() {
      removeToken()
      this.$router.push({
        name: 'login'
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.menu {
  height: calc(100vh - 40px);
  padding: 30px 10px 10px 10px;
  background: #f0f5f6;
  color: #76808c;

  .item {
    padding: 0 10px;
    margin-top: 10px;
    cursor: pointer;
    width: 100px;
    height: 40px;
    line-height: 40px;
    text-indent: 10px;
    font-size: 14px;
    border-radius: 10px;
    display: flex;
    align-items: center;

    i {
      margin: 0 5px 0 -15px;
    }

    &:hover {
      background: #e4e8ec;
      color: #4f5a6b;
    }

    &.active {
      background: #fc3d48;
      color: #ffffff;
    }
  }
}
</style>
