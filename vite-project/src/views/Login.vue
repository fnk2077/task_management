<template>
    <div class="max-w-md mx-auto p-6 bg-white rounded shadow">
      <h2 class="text-2xl font-bold mb-4">Login</h2>
      <form @submit.prevent="handleLogin">
        <div class="mb-4">
          <label>Email</label>
          <input v-model="email" type="email" class="w-full border rounded p-2" />
        </div>
        <div class="mb-4">
          <label>Password</label>
          <input v-model="password" type="password" class="w-full border rounded p-2" />
        </div>
        <button type="submit" class="bg-blue-500 text-white py-2 px-4 rounded">Login</button>
      </form>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  import { useRouter } from 'vue-router';
  
  export default {
    // setup() {
    //   const router = useRouter();
    //   const email = ref('');
    //   const password = ref('');
  
    //   const handleLogin = async () => {
    //     try {
    //       await axios.post('http://localhost:8080/api/v1/login', { email: email.value, password: password.value }, { withCredentials: true });
    //       router.push('/dashboard');
    //     } catch (error) {
    //       alert('Login failed');
    //     }
    //   };
  
    //   return { email, password, handleLogin };
    // },
  };
  </script>
  