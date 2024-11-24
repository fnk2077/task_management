<template>
  <div class="home-container">
    <h1 class="page-title">Welcome to Task Management</h1>
  </div>
</template>

<script setup>
  // You can import the necessary components or icons if required
</script>

<style scoped>
.home-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
  font-family: Arial, sans-serif;
  animation: fadeIn 1s ease-out;
}

.page-title {
  text-align: center;
  font-size: 2.5em;
  margin-bottom: 30px;
  color: #2c3e50;
  animation: slideIn 1s ease-out;
}

.welcome-item {
  display: flex;
  align-items: center;
  background-color: #f9f9f9;
  border-radius: 8px;
  margin-bottom: 20px;
  padding: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  opacity: 0;
  animation: fadeInUp 0.5s ease-out forwards;
}

.welcome-item:nth-child(1) {
  animation-delay: 0.2s;
}
.welcome-item:nth-child(2) {
  animation-delay: 0.4s;
}
.welcome-item:nth-child(3) {
  animation-delay: 0.6s;
}

.icon {
  width: 100px;
  height: 100px;
  margin-right: 20px;
  border-radius: 50%;
  object-fit: cover;
}

.welcome-item .heading {
  font-size: 1.5em;
  color: #34495e;
  margin-bottom: 10px;
  animation: fadeIn 1s ease-out;
}

.welcome-item p {
  color: #7f8c8d;
  margin-bottom: 10px;
}

.btn-view {
  display: inline-block;
  padding: 10px 20px;
  background-color: #3498db;
  color: white;
  text-decoration: none;
  border-radius: 5px;
  font-weight: bold;
  transition: background-color 0.3s;
  animation: scaleUp 0.5s ease-out;
}

.btn-view:hover {
  background-color: #2980b9;
}

a {
  text-decoration: none;
}

/* Animations */
@keyframes fadeIn {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}

@keyframes slideIn {
  0% {
    transform: translateY(-50px);
  }
  100% {
    transform: translateY(0);
  }
}

@keyframes fadeInUp {
  0% {
    opacity: 0;
    transform: translateY(20px);
  }
  100% {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes scaleUp {
  0% {
    transform: scale(0.9);
  }
  100% {
    transform: scale(1);
  }
}
</style>
