<template>
  <div class="games-view container">
    <h1 class="page-title">Партии</h1>
    
    <div v-if="loading" class="loading">
      Загрузка...
    </div>
    
    <div v-else-if="error" class="error">
      {{ error }}
    </div>
    
    <div v-else>
      <div class="games-list">
        <div 
          v-for="game in games" 
          :key="game.id"
          class="game-item card"
          @click="goToGame(game.id)"
        >
          <div class="game-item-header">
            <h3 class="game-item-title">{{ game.title }}</h3>
            <span class="game-item-date">{{ formatDate(game.date) }}</span>
          </div>
          
          <div class="game-item-players">
            <div class="player-info">
              <span class="player-name">{{ game.blackPlayer.name }}</span>
              <span class="player-rank">{{ formatRank(game.blackPlayer.rank) }}</span>
            </div>
            <span class="vs">vs</span>
            <div class="player-info">
              <span class="player-name">{{ game.whitePlayer.name }}</span>
              <span class="player-rank">{{ formatRank(game.whitePlayer.rank) }}</span>
            </div>
          </div>
          
          <div class="game-item-meta">
            <span class="meta-item">{{ game.boardSize }}×{{ game.boardSize }}</span>
            <span class="meta-item">{{ formatRules(game.rules) }}</span>
            <span class="meta-item">Коми: {{ game.komi }}</span>
          </div>
        </div>
      </div>
      
      <Pagination
        :current-page="currentPage"
        :total-pages="totalPages"
        @page-change="changePage"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { gameService } from '../api'
import type { GameSummary } from '../types'
import Pagination from '../components/Pagination.vue'

const router = useRouter()
const games = ref<GameSummary[]>([])
const loading = ref(true)
const error = ref('')
const currentPage = ref(1)
const totalPages = ref(1)
const limit = 10

async function loadGames(page: number = 1) {
  loading.value = true
  error.value = ''
  
  try {
    const response = await gameService.getGames(page, limit)
    games.value = response.games
    // Assuming we need to calculate total pages
    totalPages.value = 10
  } catch (err) {
    error.value = 'Ошибка загрузки партий'
    console.error(err)
  } finally {
    loading.value = false
  }
}

function formatDate(dateString: string): string {
  return new Date(dateString).toLocaleDateString('ru-RU')
}

function formatRank(rank: number): string {
  if (rank === 0) return '?'
  if (rank > 0) return `${rank}k`
  return `${Math.abs(rank)}d`
}

function formatRules(rules: string): string {
  const rulesMap: Record<string, string> = {
    'JAPANESE': 'JP',
    'CHINESE': 'CN'
  }
  return rulesMap[rules] || rules
}

function goToGame(id: string) {
  router.push(`/games/${id}`)
}

function changePage(page: number) {
  currentPage.value = page
  loadGames(page)
}

onMounted(() => {
  loadGames()
})
</script>

<style scoped>
.games-view {
  padding: 24px 0;
}

.page-title {
  font-size: 2rem;
  margin-bottom: 24px;
  color: var(--primary-color);
}

.loading,
.error {
  text-align: center;
  padding: 40px;
  font-size: 1.1rem;
}

.error {
  color: #e74c3c;
}

.games-list {
  display: grid;
  gap: 16px;
}

.game-item {
  cursor: pointer;
  transition: transform 0.2s, box-shadow 0.2s;
}

.game-item:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(0,0,0,0.15);
}

.game-item-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 12px;
  flex-wrap: wrap;
  gap: 8px;
}

.game-item-title {
  font-size: 1.2rem;
  color: var(--primary-color);
  margin: 0;
  flex: 1;
  min-width: 200px;
}

.game-item-date {
  color: #666;
  font-size: 0.9rem;
  white-space: nowrap;
}

.game-item-players {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  margin-bottom: 12px;
  padding: 12px;
  background-color: #f8f9fa;
  border-radius: 6px;
  flex-wrap: wrap;
}

.player-info {
  display: flex;
  flex-direction: column;
  align-items: center;
  flex: 1;
  min-width: 120px;
}

.player-name {
  font-weight: 600;
  margin-bottom: 4px;
  text-align: center;
  word-break: break-word;
}

.player-rank {
  font-size: 0.9rem;
  color: #666;
}

.vs {
  font-weight: bold;
  color: #666;
  padding: 0 8px;
}

.game-item-meta {
  display: flex;
  gap: 16px;
  flex-wrap: wrap;
}

.meta-item {
  font-size: 0.9rem;
  color: #666;
  padding: 4px 8px;
  background-color: #e9ecef;
  border-radius: 4px;
}

@media (min-width: 768px) {
  .games-view {
    padding: 32px 0;
  }
  
  .games-list {
    grid-template-columns: repeat(2, 1fr);
    gap: 20px;
  }
}

@media (min-width: 1024px) {
  .games-list {
    grid-template-columns: repeat(3, 1fr);
  }
}
</style>