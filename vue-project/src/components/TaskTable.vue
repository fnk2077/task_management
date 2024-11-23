<template>
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
            <task-actions :task="task" @view="openViewModal" @edit="openEditModal" @delete="deleteTask" />
          </td>
        </tr>
      </tbody>
    </table>
</template>
  
<script>
  import TaskActions from './TaskActions.vue';
  
  export default {
    components: {
      TaskActions,
    },
    props: {
      tasks: Array,
      formatDate: Function,
      openViewModal: Function,
      openEditModal: Function,
      deleteTask: Function,
    }
  };
</script>
  