<template>
  <div class="p-6 bg-gray-100 min-h-screen">
    <h1 class="text-3xl font-bold text-blue-600 mb-6">Task List</h1>

    <!-- ปุ่มสำหรับสร้าง Task -->
    <div class="mb-4">
      <button
        @click="openCreateModal"
        class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
      >
        + Create Task
      </button>
    </div>

    <!-- ตารางแสดง Task -->
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
                @click="deleteTask(task.ID)"
                class="bg-red-600 text-white px-2 py-1 rounded hover:bg-red-700"
              >
                Delete
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Modal สำหรับดูรายละเอียด Task -->
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

    <!-- Modal สำหรับสร้างหรือแก้ไข Task -->
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
export default {
  data() {
    return {
      tasks: [
        {
          ID: 1,
          title: "Test Task 1",
          description: "Description for Task 1",
          status: "Pending",
          priority: "Medium",
          assigned_to: "a1@a.com",
          created_by: "a2@a.com",
          CreatedAt: "2024-11-23T09:34:05.122976+07:00",
          due_date: "2024-12-01T09:34:05.122976+07:00",
        },
        {
          ID: 2,
          title: "Test Task 2",
          description: "Description for Task 2",
          status: "In Progress",
          priority: "High",
          assigned_to: "a3@a.com",
          created_by: "a4@a.com",
          CreatedAt: "2024-11-23T10:00:00.122976+07:00",
          due_date: "2024-12-03T09:00:00.122976+07:00",
        },
      ],
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
  methods: {
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
      // รีเซ็ตข้อมูลฟอร์ม
      this.taskForm = {
        title: "",
        description: "",
        status: "Pending",
        priority: "Low",
        assigned_to: "",
        due_date: "",
      };
    },
    closeEditModal() {
      this.showEditModal = false;
      // รีเซ็ตข้อมูลฟอร์ม
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
      const newTask = {
        ...this.taskForm,
        ID: this.tasks.length + 1,
        CreatedAt: new Date().toISOString(),
      };
      this.tasks.push(newTask);
      this.closeCreateModal();
    },
    updateTask() {
      const index = this.tasks.findIndex((task) => task.ID === this.taskForm.ID);
      if (index !== -1) {
        this.tasks[index] = { ...this.taskForm };
      }
      this.showEditModal = false;
    },
    deleteTask(ID) {
      this.tasks = this.tasks.filter((task) => task.ID !== ID);
    },
  },
};
</script>

<style scoped>
.table-auto {
  width: 100%;
}
</style>
