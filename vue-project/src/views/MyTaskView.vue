<template>
  <div class="p-6 bg-gray-100 min-h-screen">
    <h1 class="text-3xl font-bold text-blue-600 mb-6">Task List</h1>

    <div class="mb-4">
      <button
        @click="openCreateModal"
        class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
      >
        + Create Task
      </button>
    </div>

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
            <td class="px-4 py-2 border">{{ task.assigned_to }}</td>
            <td class="px-4 py-2 border">{{ task.created_by }}</td>
            <td class="px-4 py-2 border">{{ formatDate(task.CreatedAt) }}</td>
            <td class="px-4 py-2 border">{{ formatDate(task.due_date) }}</td>
            <td class="px-4 py-2 border text-center">
              <button
                @click="openViewModal(task)"
                class="bg-blue-500 text-white px-2 py-1 rounded hover:bg-blue-600 mr-2"
              >
                View
              </button>
              <button
                @click="openEditModal(task)"
                class="bg-yellow-500 text-white px-2 py-1 rounded hover:bg-yellow-600 mr-2"
              >
                Edit
              </button>
              <button
                v-if="userEmail === task.created_by"
                @click="confirmDelete(task.ID)"
                class="bg-red-600 text-white px-2 py-1 rounded hover:bg-red-700"
              >
                Delete
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
        <p><strong>Assigned To:</strong> {{ viewTask.assigned_to }}</p>
        <p><strong>Created By:</strong> {{ viewTask.created_by }}</p>
        <p><strong>Created At:</strong> {{ formatDate(viewTask.CreatedAt) }}</p>
        <p><strong>Due Date:</strong> {{ formatDate(viewTask.due_date) }}</p>
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

    <div
      v-if="showCreateModal || showEditModal"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center"
    >
      <div class="bg-white rounded-lg p-6 w-96">
        <h2 class="text-xl font-bold mb-4">{{ showEditModal ? 'Edit' : 'Create' }} Task</h2>
        <form @submit.prevent="showEditModal ? updateTask() : createTask()">
          <div class="mb-4">
            <label class="block mb-2" for="title">Title</label>
            <input
              v-model="taskForm.title"
              type="text"
              id="title"
              class="w-full p-2 border border-gray-300 rounded"
              required
            />
          </div>
          <div class="mb-4">
            <label class="block mb-2" for="description">Description</label>
            <input
              v-model="taskForm.description"
              type="text"
              id="description"
              class="w-full p-2 border border-gray-300 rounded"
              required
            />
          </div>
          <div class="mb-4">
            <label class="block mb-2" for="status">Status</label>
            <select
              v-model="taskForm.status"
              id="status"
              class="w-full p-2 border border-gray-300 rounded"
            >
              <option value="Pending">Pending</option>
              <option value="In Progress">In Progress</option>
              <option value="Completed">Completed</option>
            </select>
          </div>
          <div class="mb-4">
            <label class="block mb-2" for="priority">Priority</label>
            <select
              v-model="taskForm.priority"
              id="priority"
              class="w-full p-2 border border-gray-300 rounded"
            >
              <option value="Low">Low</option>
              <option value="Medium">Medium</option>
              <option value="High">High</option>
            </select>
          </div>
          <div class="mb-4">
            <label class="block mb-2" for="assignedTo">Assigned To</label>
            <input
              v-model="taskForm.assigned_to"
              type="email"
              id="assignedTo"
              class="w-full p-2 border border-gray-300 rounded"
            />
          </div>
          <div class="mb-4">
            <label class="block mb-2" for="dueDate">Due Date</label>
            <input
              v-model="taskForm.due_date"
              type="date"
              id="dueDate"
              class="w-full p-2 border border-gray-300 rounded"
            />
          </div>
          <div class="flex justify-between mt-4">
            <button
              @click="cancelAction"
              class="bg-gray-300 text-gray-700 px-4 py-2 rounded"
            >
              Cancel
            </button>
            <button
              type="submit"
              class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
            >
              {{ showEditModal ? 'Update' : 'Create' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      userEmail: localStorage.getItem("email"),
      tasks: [],
      taskForm: {
        ID: null,
        title: "",
        description: "",
        status: "Pending",
        priority: "Low",
        assigned_to: "",
        created_by: "",
        due_date: "",
      },
      showCreateModal: false,
      showEditModal: false,
      showViewModal: false,
      viewTask: null,
    };
  },
  created() {
    this.fetchTasks();
  },
  methods: {
    async fetchTasks() {
      const email = this.userEmail
      try {
        const response = await axios.get(`http://localhost:8080/api/v1/task?email=${email}`);
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
    openCreateModal() {
      this.showCreateModal = true;
    },
    openViewModal(task) {
      this.viewTask = task;
      this.showViewModal = true;
    },
    closeViewModal() {
      this.showViewModal = false;
      this.viewTask = null;
    },
    openEditModal(task) {
      this.taskForm = { ...task };
      this.showEditModal = true;
    },
    cancelAction() {
      if (this.showCreateModal) {
        this.closeCreateModal();
      } else if (this.showEditModal) {
        this.closeEditModal();
      }
    },
    closeCreateModal() {
      this.showCreateModal = false;
      this.resetForm();
    },
    closeEditModal() {
      this.showEditModal = false;
      this.resetForm();
    },
    resetForm() {
      this.taskForm = {
        title: "",
        description: "",
        status: "Pending",
        priority: "Low",
        assigned_to: "",
        due_date: "",
      };
    },
    createTask() {
      const userEmail = localStorage.getItem("email");
      const taskData = {
        title: this.taskForm.title,
        description: this.taskForm.description,
        status: this.taskForm.status,
        priority: this.taskForm.priority,
        assigned_to: this.taskForm.assigned_to,
        created_by: userEmail,
        due_date: new Date(this.taskForm.due_date).toISOString(),
      };

      axios
        .post("http://localhost:8080/api/v1/task", taskData)
        .then(() => {
          alert("Task ถูกสร้างเรียบร้อยแล้ว");
          this.fetchTasks();
          this.closeCreateModal();
        })
        .catch((error) => {
          console.error("Error creating task:", error);
          alert("เกิดข้อผิดพลาดในการสร้าง Task");
        });
    },
    updateTask() {
      const updatedTask = {
        id: this.taskForm.ID,
        title: this.taskForm.title,
        description: this.taskForm.description,
        status: this.taskForm.status,
        assigned_to: this.taskForm.assigned_to,
        created_by: this.taskForm.created_by,
        due_date: new Date(this.taskForm.due_date).toISOString(),
      };

      axios
        .put("http://localhost:8080/api/v1/task", updatedTask)
        .then(() => {
          alert("Task ถูกอัปเดตเรียบร้อยแล้ว");
          this.fetchTasks();
          this.closeEditModal();
        })
        .catch((error) => {
          console.error("Error updating task:", error);
          alert("เกิดข้อผิดพลาดในการอัปเดต Task");
        });
    },
    confirmDelete(ID) {
      const confirmed = window.confirm(`คุณต้องการลบ Task ID ${ID} ใช่หรือไม่?`);
      if (confirmed) {
        this.deleteTask(ID);
      }
    },
    async deleteTask(ID) {
      try {

        await axios.delete("http://localhost:8080/api/v1/task", {
          data: { id: ID },
        });
        
        this.tasks = this.tasks.filter((task) => task.ID !== ID);

        alert(`Task ID ${ID} ถูกลบเรียบร้อยแล้ว`);
      } catch (error) {
        console.error("Error deleting task:", error);
        alert("เกิดข้อผิดพลาดในการลบ Task");
      }
    },
  },
};
</script>


<style scoped>
.table-auto {
  width: 100%;
}
</style>
