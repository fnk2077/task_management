<template>
  <nav class="bg-blue-600 text-white">
    <div class="container mx-auto flex items-center justify-between p-4">
      <div class="text-xl font-bold">
        <RouterLink to="/" class="hover:text-blue-200">MyApp</RouterLink>
      </div>
      <ul class="flex space-x-4">
        <li>
          <RouterLink to="/" class="hover:text-blue-200 transition duration-300" active-class="font-semibold underline">
            Home
          </RouterLink>
        </li>
        <li>
          <RouterLink to="/task" class="hover:text-blue-200 transition duration-300" active-class="font-semibold underline">
            Task
          </RouterLink>
        </li>
        <li v-if="isAuthenticated">
          <RouterLink to="/mytask" class="hover:text-blue-200 transition duration-300" active-class="font-semibold underline">
            MyTask
          </RouterLink>
        </li>

        <li v-if="!isAuthenticated">
          <RouterLink to="/login" class="hover:text-blue-200 transition duration-300" active-class="font-semibold underline">
            Login
          </RouterLink>
        </li>
        <li v-if="!isAuthenticated">
          <RouterLink to="/register" class="hover:text-blue-200 transition duration-300" active-class="font-semibold underline">
            Register
          </RouterLink>
        </li>

        <li v-if="isAuthenticated">
          <span class="text-white">{{ user.email }}</span>
        </li>
        <li v-if="isAuthenticated">
          <button @click="logout" class="text-white hover:text-gray-300">Logout</button>
        </li>
      </ul>
    </div>
  </nav>
</template>

<script>
export default {
  data() {
    return {
      userEmail: localStorage.getItem("email"),
      accessToken: localStorage.getItem("accessToken")
    };
  },
  computed: {
    isAuthenticated() {
      return this.userEmail && this.accessToken;
    },
    user() {
      return {
        email: this.userEmail
      };
    }
  },
  methods: {
    logout() {
      localStorage.removeItem("email");
      localStorage.removeItem('accessToken');
      localStorage.removeItem('refreshToken');
      
      this.userEmail = null;
      this.accessToken = null;

      this.$router.push('/login');  
    }
  }
};
</script>
