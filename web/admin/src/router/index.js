import Vue from 'vue'
import VueRouter from 'vue-router'
import LoginView from '../views/LoginView.vue'
import AdminView from '../views/AdminView.vue'

import Index from '../components/admin/Index.vue'
import AddArticle from '../components/article/AddArticle.vue'
import ArticleList from '../components/article/ArticleList.vue'
import CategoryList from '../components/category/CategoryList.vue'
import UserList from '../components/user/UserList.vue'
import Profile from '../components/user/Profile.vue'

Vue.use(VueRouter)

const routes = [
    {
        path: '/login',
        name: 'login',
        component: LoginView
    },
    {
        path: '/',
        name: 'admin',
        component: AdminView,
        children: [
            { path: 'index', component: Index },
            { path: 'addarticle', component: AddArticle },
            { path: 'addarticle/:id', component: AddArticle, props: true },
            { path: 'articlelist', component: ArticleList },
            { path: 'categorylist', component: CategoryList },
            { path: 'userlist', component: UserList },
            { path: 'profile', component: Profile }
        ]
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
