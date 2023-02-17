import Vue from 'vue'
import Antd from 'ant-design-vue'
import App from './App.vue'
import router from './router'
import './assets/css/style.css'
import 'ant-design-vue/dist/antd.css'
import './plugin/http'

Vue.config.productionTip = false

Vue.use(Antd)

Vue.prototype.$message.config({
    top: '60px',
    duration: 2,
    maxCount: 3
})

new Vue({
    router,
    render: (h) => h(App)
}).$mount('#app')
