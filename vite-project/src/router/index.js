import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomeView
        },
        {
            path: '/login',
            name: 'login',
            beforeEnter: (to, from, next) => {
                if (!localStorage.getItem('jwt_token')) {
                    next()
                } else {
                    next('/posts')
                }
            },
            component: () => import('@/views/login-register/LoginView.vue')
        },
        {
            path: '/logout',
            name: 'logout',
            component: () => import('@/views/LogoutView.vue')
        },
        {
            path: '/posts',
            name: 'posts',
            beforeEnter: (to, from, next) => {
                if (localStorage.getItem('jwt_token')) {
                    next()
                } else {
                    next('/login')
                }
            },
            component: () => import('@/views/posts/AllPostView.vue')
        },
        {
            path: '/chat',
            name: 'chat',
            beforeEnter: (to, from, next) => {
                if (localStorage.getItem('jwt_token') && useAuthStore().auth.role !== 'ADMIN') {
                    next()
                } else {
                    next('/login')
                }
            },
            component: () => import('@/views/chat/AllChatView.vue')
        },
        {
            path: '/register',
            name: 'register',
            component: () => import('@/views/login-register/RegisterView.vue')
        },
        {
            path: '/nearby',
            name: 'nearby',
            component: () => import('@/views/posts/NearByPostView.vue')
        },
        {
            path: '/posts/:id',
            name: 'posts.show',
            component: () => import('@/views/posts/DetailView.vue')
        },
        {
            path: '/posts/precreate',
            name: 'posts.precreate',
            beforeEnter: (to, from, next) => {
                if (useAuthStore().auth.role === 'ADMIN') {
                    next('/posts')
                } else if (localStorage.getItem('jwt_token')) {
                    next()
                } else {
                    next('/login')
                }
            },
            component: () => import('@/views/posts/PreCreateView.vue')
        },
        {
            path: '/posts/create/:is_lost',
            name: 'posts.create',
            beforeEnter: (to, from, next) => {
                if (useAuthStore().auth.role === 'ADMIN') {
                    next('/posts')
                } else if (localStorage.getItem('jwt_token')) {
                    next()
                } else {
                    next('/login')
                }
            },
            component: () => import('@/views/posts/CreateView.vue')
        },
        {
            path: '/posts/edit/:id',
            name: 'posts.edit',
            beforeEnter: (to, from, next) => {
                if (useAuthStore().auth.role === 'ADMIN') {
                    next('/posts')
                } else if (localStorage.getItem('jwt_token')) {
                    next()
                } else {
                    next('/login')
                }
            },
            component: () => import('@/views/posts/EditView.vue')
        },
    ]

})

router.beforeEach((to, from, next) => {
    if (to.name !== 'login' && to.name !== 'register' && !localStorage.getItem('jwt_token')) {
        next({name: 'login'})
    } else{
        next()
    }
})

export default router
