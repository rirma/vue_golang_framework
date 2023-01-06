import { createRouter, createWebHistory } from 'vue-router'
import Page_1 from '../components/Page1.vue'
import Page_2 from '../components/Page2.vue'
import Login from '../components/Login.vue'
import HelloWorld from '../components/HelloWorld.vue'
// import Store from '../store/index.js'

const routes = [
    {
        path: '/',
        name: 'home',
        component: HelloWorld
    },
    {
        path: '/login',
        name: 'login',
        component: Login
    },
    {
        path: '/page1',
        name: 'page1',
        component: Page_1,
        meta: {
        requiresAuth: true   // 認証済の時のみ閲覧可能となるように定義
        }
    },
    {
        path: '/page2',
        name: 'page2',
        component: Page_2,
        meta: {
        requiresAuth: true   // 認証済の時のみ閲覧可能となるように定義
        }
    },
]

const router = createRouter({
    history: createWebHistory(),
    base: process.env.BASE_URL,
    routes
})

/*router.beforeEach((to, from, next) => {
    if (to.matched.some(record => record.meta.requiresAuth)) {
        if (!Store.state.login_user_information.email) {
            // 認証されていない時、認証画面へ遷移
            next({
                path: '/login',
                query: {
                    redirect: to.fullPath
                }
            })
        } else {
            next();
        }
    } else {
        next();
    }
});*/

export default router