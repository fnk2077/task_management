<template>
    <nav class="bg-blue-600 text-white">
      <div class="container mx-auto flex items-center justify-between p-4">
        <div class="text-xl font-bold">
          <RouterLink to="/" class="hover:text-blue-200">MyApp</RouterLink>
        </div>
        <ul class="flex space-x-4">
          <li>
            <RouterLink
              to="/"
              class="hover:text-blue-200 transition duration-300"
              active-class="font-semibold underline"
            >
              Home
            </RouterLink>
          </li>
          <li>
            <RouterLink
              to="/task"
              class="hover:text-blue-200 transition duration-300"
              active-class="font-semibold underline"
            >
              Task
            </RouterLink>
          </li>
          <li>
            <RouterLink
              to="/login"
              class="hover:text-blue-200 transition duration-300"
              active-class="font-semibold underline"
            >
              Login
            </RouterLink>
          </li>
          <li>
            <RouterLink
              to="/register"
              class="hover:text-blue-200 transition duration-300"
              active-class="font-semibold underline"
            >
              Register
            </RouterLink>
          </li>
        </ul>
      </div>
    </nav>
</template>
  