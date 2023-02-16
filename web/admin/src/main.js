import Vue from 'vue'
import Antd from 'ant-design-vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import './assets/css/style.css'

import 'ant-design-vue/dist/antd.css'
Vue.config.productionTip = false

Vue.use(Antd)

Vue.prototype.$message.config({
    top: '60px',
    duration: 2,
    maxCount: 3
})

axios.defaults.baseURL = 'http://localhost:3000/api/v1'
axios.interceptors.request.use(config => {
    config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
    return config
})
// axios.interceptors.response.use(res => {
//     console.log(res)
//     if (res.status === 200) {
//         return res
//     } else {
//         alert(res.statusText)
//         return res
//     }
// })
Vue.prototype.$http = axios

new Vue({
    router,
    render: h => h(App)
}).$mount('#app')
