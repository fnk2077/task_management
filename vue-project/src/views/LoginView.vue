<template>
    <div class="flex items-center justify-center min-h-screen bg-blue-50">
      <div class="w-full max-w-sm p-6 bg-white rounded-lg shadow-md">
        <h1 class="text-2xl font-bold text-center text-blue-600 mb-6">Login</h1>
        <form @submit.prevent="onSubmit">
          <div class="mb-4">
            <label for="email" class="block text-sm font-medium text-gray-700">Email</label>
            <input
              id="email"
              type="email"
              v-model="email"
              class="block w-full px-3 py-2 mt-1 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="you@example.com"
              required
            />
          </div>
          <div class="mb-6">
            <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
            <input
              id="password"
              type="password"
              v-model="password"
              class="block w-full px-3 py-2 mt-1 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="••••••••"
              required
            />
          </div>
          <button
            type="submit"
            class="w-full px-4 py-2 text-white bg-blue-600 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          >
            Login
          </button>
        </form>
        <p class="mt-4 text-sm text-center text-gray-600">
          Don't have an account?
          <a href="/register" class="text-blue-500 hover:underline">Sign up</a>
        </p>
      </div>
    </div>
</template>
  
<script>
  export default {
    data() {
      return {
        email: "",
        password: "",
      };
    },
    methods: {
      onSubmit() {
        console.log("Email:", this.email);
        console.log("Password:", this.password);
        // Handle the login logic here
      },
    },
  };
</script>
  