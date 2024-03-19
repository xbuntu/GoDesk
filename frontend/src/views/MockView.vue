<template>
  <div class="mock">
    <div class="mb-4">
      <el-button type="primary" @click="getMockData">模拟mock请求</el-button>
      <el-button type="success" @click="getApiData('get')">GET 请求</el-button>
      <el-button type="warning" @click="getApiData('post')">POST 请求</el-button>
      <el-button type="danger" @click="getGoData">调用Go代码</el-button>
    </div>
    <div class="result">
      <json-viewer
          :value="jsonData"
          :expand-depth="5"
          boxed
          sort
          copyable
      />
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import {GetConfig} from "../../wailsjs/go/controller/Config";
import {mockIndex , getApi, postApi} from "@/api/mock";
import JsonViewer from 'vue-json-viewer'
export default {
  name: 'MockView',
  components:{
    JsonViewer
  },
  data(){
    return {
      jsonData:{}
    }
  },
  created() {
    this.getMockData()
  },
  methods:{
    getMockData(){
      mockIndex().then(res => {
        this.jsonData = res
      })
    },
    getGoData(){
      GetConfig().then(res => {
        this.jsonData = res
      })
    },
    getApiData(method){
      let res = method === 'get'?getApi():postApi();
      res.then(res => {
        console.log(res)
      }).catch(res => {
        this.jsonData = {
          method:method,
          message:res.message
        }
      }).finally(()=>{})
    }
  }
}
</script>

<style lang="scss" scoped>
  .mock{
    padding: 20px;
    .result{
      margin-top: 20px;
    }
  }
</style>
