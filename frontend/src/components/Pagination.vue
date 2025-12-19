<template>
  <div class="pagination">
    <button 
      @click="goToPage(currentPage - 1)"
      class="pagination-btn"
      :disabled="currentPage === 1"
    >
      ←
    </button>
    
    <div class="page-numbers">
      <button
        v-for="page in visiblePages"
        :key="page"
        @click="goToPage(page)"
        class="page-btn"
        :class="{ active: page === currentPage }"
      >
        {{ page }}
      </button>
      
      <span v-if="showEndEllipsis" class="ellipsis">...</span>
    </div>
    
    <button 
      @click="goToPage(currentPage + 1)"
      class="pagination-btn"
      :disabled="currentPage === totalPages"
    >
      →
    </button>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  currentPage: number
  totalPages: number
  maxVisible?: number
}

const props = withDefaults(defineProps<Props>(), {
  maxVisible: 5
})

const emit = defineEmits<{
  'page-change': [page: number]
}>()

const visiblePages = computed(() => {
  const pages: number[] = []
  let start = Math.max(1, props.currentPage - Math.floor(props.maxVisible / 2))
  let end = start + props.maxVisible - 1
  
  if (end > props.totalPages) {
    end = props.totalPages
    start = Math.max(1, end - props.maxVisible + 1)
  }
  
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  
  return pages
})

const showEndEllipsis = computed(() => {
  return props.totalPages > props.maxVisible && 
         visiblePages.value[visiblePages.value.length - 1] < props.totalPages
})

function goToPage(page: number) {
  if (page >= 1 && page <= props.totalPages) {
    emit('page-change', page)
  }
}
</script>

<style scoped>
.pagination {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  margin-top: 24px;
  flex-wrap: wrap;
}

.pagination-btn,
.page-btn {
  min-width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px solid var(--border-color);
  background: white;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.2s;
}

.pagination-btn:hover:not(:disabled),
.page-btn:hover:not(.active) {
  border-color: var(--secondary-color);
  color: var(--secondary-color);
}

.pagination-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.page-btn.active {
  background-color: var(--secondary-color);
  color: white;
  border-color: var(--secondary-color);
}

.page-numbers {
  display: flex;
  gap: 4px;
}

.ellipsis {
  display: flex;
  align-items: center;
  justify-content: center;
  min-width: 40px;
  height: 40px;
}
</style>