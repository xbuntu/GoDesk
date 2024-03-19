import "../mockjs/index";
import request from "../utils/request";

//Mock模拟
export function mockIndex() {
    return request({
        url: '/mock/index',
        method: 'get',
        baseURL:'' //mock请求，需要去掉接口域名
    })
}

//Get请求
export function getApi() {
    return request({
        url: '/get',
        method: 'get'
    })
}

//Post请求
export function postApi(data) {
    return request({
        url: '/create',
        method: 'post',
        data
    })
}
