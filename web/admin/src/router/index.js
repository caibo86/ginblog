import Vue from 'vue'
import VueRouter from 'vue-router'
import LoginView from '../views/LoginView.vue'
import AdminView from '../views/AdminView.vue'

Vue.use(VueRouter)

const routes = [
    {
        path: '/login',
        name: 'login',
        component: LoginView
    },
    {
        path: '/admin',
        name: 'admin',
        component: AdminView
    }
]

const router = new VueRouter({
    routes
})

export default router
