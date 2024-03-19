<template>
  <div class="app">
    <div class="appMenu">
      <Menu v-if="routeName !== 'login'" :routeName="routeName"/>
    </div>
    <div class="appMain">
      <router-view/>
    </div>
  </div>
</template>

<script>
import Menu from './components/Menu.vue'
export default {
  name: 'AppView',
  components: {
    Menu
  },
  computed: {
    routeName() {
      return this.$route.name;
    }
  },
  data(){
    return {
    }
  },
  created() {

  },
  methods:{

  }
}
</script>


<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
html,body{
  margin: 0;
  padding: 0;
}
body{
  width: 100vw;
  height: 100vh;
}
.app{
  display: flex;
  align-items: center;
  .appMain{
    width: calc(100vw - 120px);
    height: 100vh;
    overflow: auto;
    flex: 1;
  }
}
</style>
