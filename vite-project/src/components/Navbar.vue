<template>
    <div class="px-4 py-5 mx-auto sm:max-w-xl md:max-w-full lg:max-w-screen-xl md:px-24 lg:px-8 border-2">
        <div class="relative flex items-center justify-between">
            <a href="/" aria-label="Company" title="Company" class="inline-flex items-center"> 
                <span class="ml-2 text-xl font-bold tracking-wide text-gray-800 uppercase">Memorize</span>
            </a>
            <ul class="flex items-center hidden space-x-8 lg:flex">
                <li>
                    <RouterLink to="/" href="/" aria-label="Our product" title="Our product" class="font-medium tracking-wide text-gray-700 transition-colors duration-200 hover:text-deep-purple-accent-400">
                        Home
                    </RouterLink>
                </li>
                <li>
                    <RouterLink to="/posts" href="/" aria-label="About us" title="About us" class="font-medium tracking-wide text-gray-700 transition-colors duration-200 hover:text-deep-purple-accent-400">
                        Posts
                    </RouterLink>
                </li>
                <li>
                    <RouterLink to="/posts/precreate" href="/" aria-label="Create post" title="Create post" class="font-medium tracking-wide text-gray-700 transition-colors duration-200 hover:text-deep-purple-accent-400">
                        Create post
                    </RouterLink>
                </li>
            </ul>
        </div>
    </div>
</template>

<script>
// import Profile from './components/profile/Profile.vue'
// import {useAuthStore} from './stores/auth.js'

export default {
    // setup() {
    //     // const auth_store = useAuthStore()
    //     // return {auth_store}
    // },
    // // components: {Profile},
    // data() {
    //     // return {
    //     //     auth: null,
    //     //     isMenuOpen: false,
    //     // };
    // },
    // watch: {
    //     // auth_store: {
    //     //     immediate: true,
    //     //     deep: true,
    //     //     handler(newValue, oldValue) {
    //     //         console.log(newValue.getAuth)
    //     //         this.auth = this.auth_store.getAuth
    //     //     }
    //     // }
    // },
    // mounted() {
    //     // if (this.auth_store.isAuthen) {
    //     //     this.auth = this.auth_store.getAuth
    //     // } else {
    //     //     this.auth = null
    //     // }

    // }
};
</script>
