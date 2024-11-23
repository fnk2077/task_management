<template>
    <div>
      <button
        @click="$emit('view', task)"
        class="bg-blue-500 text-white px-2 py-1 rounded hover:bg-blue-600 mr-2"
      >
        View
      </button>
      <button
        @click="$emit('edit', task)"
        class="bg-yellow-500 text-white px-2 py-1 rounded hover:bg-yellow-600 mr-2"
      >
        Edit
      </button>
      <button
        @click="$emit('delete', task.ID)"
        class="bg-red-600 text-white px-2 py-1 rounded hover:bg-red-700"
      >
        Delete
      </button>
    </div>
</template>
  
<script>
  export default {
    props: {
      task: Object,
    },
  };
</script>
  