<template>
    <div v-if="isVisible" class="modal-overlay">
      <div class="modal-container">
        <h2 class="text-xl font-bold">{{ modalTitle }}</h2>
  
        <!-- ถ้า modal เป็นการดู (view) ให้แสดงข้อมูล task -->
        <div v-if="modalAction === 'View' && viewTask">
          <p><strong>Title:</strong> {{ viewTask.title }}</p>
          <p><strong>Description:</strong> {{ viewTask.description }}</p>
          <p><strong>Status:</strong> {{ viewTask.status }}</p>
          <p><strong>Priority:</strong> {{ viewTask.priority }}</p>
          <p><strong>Assigned To:</strong> {{ viewTask.assigned_to }}</p>
          <p><strong>Created By:</strong> {{ viewTask.created_by }}</p>
          <p><strong>Created At:</strong> {{ formatDate(viewTask.CreatedAt) }}</p>
          <p><strong>Due Date:</strong> {{ formatDate(viewTask.due_date) }}</p>
        </div>
  
        <!-- ฟอร์มสำหรับ Create หรือ Edit -->
        <div v-else>
          <form @submit.prevent="handleSubmit">
            <label for="title">Title</label>
            <input v-model="taskForm.title" id="title" type="text" required />
  
            <label for="description">Description</label>
            <textarea v-model="taskForm.description" id="description" required></textarea>
  
            <label for="status">Status</label>
            <select v-model="taskForm.status">
              <option value="Pending">Pending</option>
              <option value="In Progress">In Progress</option>
              <option value="Completed">Completed</option>
            </select>
  
            <label for="priority">Priority</label>
            <select v-model="taskForm.priority">
              <option value="Low">Low</option>
              <option value="Medium">Medium</option>
              <option value="High">High</option>
            </select>
  
            <label for="assigned_to">Assigned To</label>
            <input v-model="taskForm.assigned_to" id="assigned_to" type="email" required />
  
            <label for="due_date">Due Date</label>
            <input v-model="taskForm.due_date" id="due_date" type="date" required />
  
            <button type="submit">{{ modalAction }} Task</button>
          </form>
        </div>
  
        <button @click="closeModal" class="close-button">Close</button>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      isVisible: Boolean,
      modalTitle: String,
      modalAction: String,
      taskForm: Object,
      handleSubmit: Function,
      closeModal: Function,
      viewTask: Object,  // รับข้อมูล viewTask
      formatDate: Function,
    },
  };
  </script>
  
  <style scoped>
  .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
  }
  
  .modal-container {
    background: white;
    padding: 20px;
    border-radius: 8px;
    width: 400px;
    max-width: 90%;
  }
  
  .close-button {
    background-color: red;
    color: white;
    padding: 10px;
    border-radius: 5px;
  }
  </style>
  