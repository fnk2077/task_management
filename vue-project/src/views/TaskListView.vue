<template>
    <div class="p-6 bg-gray-100 min-h-screen">
      <h1 class="text-3xl font-bold text-blue-600 mb-6">Task List</h1>
  
      <div class="bg-white shadow-md rounded-lg overflow-hidden">
        <table class="min-w-full table-auto border-collapse border border-gray-200">
          <thead class="bg-blue-600 text-white">
            <tr>
              <th class="px-4 py-2 border">ID</th>
              <th class="px-4 py-2 border">Title</th>
              <th class="px-4 py-2 border">Description</th>
              <th class="px-4 py-2 border">Status</th>
              <th class="px-4 py-2 border">Priority</th>
              <th class="px-4 py-2 border">Assigned To</th>
              <th class="px-4 py-2 border">Created By</th>
              <th class="px-4 py-2 border">Created At</th>
              <th class="px-4 py-2 border">Due Date</th>
              <th class="px-4 py-2 border">Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="task in tasks"
              :key="task.ID"
              class="odd:bg-white even:bg-gray-50 hover:bg-gray-100"
            >
              <td class="px-4 py-2 border text-center">{{ task.ID }}</td>
              <td class="px-4 py-2 border">{{ task.title }}</td>
              <td class="px-4 py-2 border truncate max-w-xs">{{ task.description }}</td>
              <td class="px-4 py-2 border">{{ task.status }}</td>
              <td class="px-4 py-2 border">{{ task.priority }}</td>
              <td class="px-4 py-2 border">{{ task.assignedTo }}</td>
              <td class="px-4 py-2 border">{{ task.createdBy }}</td>
              <td class="px-4 py-2 border">{{ formatDate(task.CreatedAt) }}</td>
              <td class="px-4 py-2 border">{{ formatDate(task.dueDate) }}</td>
              <td class="px-4 py-2 border text-center">
                <button
                  @click="openViewModal(task)"
                  class="bg-blue-500 text-white px-2 py-1 rounded hover:bg-blue-600"
                >
                  View
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
  
      <div
        v-if="showViewModal"
        class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center"
      >
        <div class="bg-white rounded-lg p-6 w-96">
          <h2 class="text-xl font-bold mb-4">Task Details</h2>
          <p><strong>ID:</strong> {{ viewTask.ID }}</p>
          <p><strong>Title:</strong> {{ viewTask.title }}</p>
          <p><strong>Description:</strong> {{ viewTask.description }}</p>
          <p><strong>Status:</strong> {{ viewTask.status }}</p>
          <p><strong>Priority:</strong> {{ viewTask.priority }}</p>
          <p><strong>Assigned To:</strong> {{ viewTask.assignedTo }}</p>
          <p><strong>Created By:</strong> {{ viewTask.createdBy }}</p>
          <p><strong>Created At:</strong> {{ formatDate(viewTask.CreatedAt) }}</p>
          <p><strong>Due Date:</strong> {{ formatDate(viewTask.dueDate) }}</p>
          <div class="flex justify-end mt-4">
            <button
              @click="closeViewModal"
              class="bg-gray-300 text-gray-700 px-4 py-2 rounded"
            >
              Close
            </button>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import axios from "axios";
  
  export default {
    data() {
      return {
        tasks: [],
        showViewModal: false,
        viewTask: null,
      };
    },
    created() {
      this.fetchTasks();
    },
    methods: {
      async fetchTasks() {
        try {
          const response = await axios.get("http://localhost:8080/api/v1/task");
          this.tasks = response.data;
        } catch (error) {
          console.error("Error fetching tasks:", error);
        }
      },
      formatDate(date) {
        const newDate = new Date(date);
        return newDate.toLocaleDateString("en-GB", {
          year: "numeric",
          month: "short",
          day: "numeric",
        });
      },
      openViewModal(task) {
        this.viewTask = task;
        this.showViewModal = true;
      },
      closeViewModal() {
        this.showViewModal = false;
        this.viewTask = null;
      },
    },
  };
  </script>
  