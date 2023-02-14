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

router.beforeEach((to, from, next) => {
    const token = window.sessionStorage.getItem('token')
    if (to.path === '/login') return next()
    if (!token && to.path === '/admin') {
        next('/login')
    } else {
        next()
    }
})

export default router
