<script setup>
import { RouterLink, RouterView } from 'vue-router'
import Navbar from './components/Navbar.vue'
</script>

<template>
  <header>
      <nav>
        <Navbar></Navbar>
      </nav>
  </header>

  <RouterView />
</template>