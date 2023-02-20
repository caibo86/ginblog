import Vue from 'vue'
import axios from 'axios'

axios.defaults.baseURL = 'http://localhost:3000/api/v1'
axios.interceptors.request.use((config) => {
    // config.headers.Authorization = `Bearer ${window.sessionStorage.getItem(
    //     'token'
    // )}`
    config.headers.Authorization = 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImNhaWJvIiwiaXNzIjoiZ2luYmxvZyIsImV4cCI6MTY3OTQ3MjM3OH0.oGQdHJVWusgXDnymtMS4Zj2vGgyl9QnOxUhgBhwX1ec'
    return config
})
Vue.prototype.$http = axios
