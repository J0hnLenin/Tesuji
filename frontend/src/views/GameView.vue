<template>
  <div class="game-view container">
    <router-link to="/games" class="back-link">
      ← Назад к списку
    </router-link>
    
    <div v-if="loading" class="loading">
      <div class="spinner"></div>
      <p>Загрузка партии...</p>
    </div>
    
    <div v-else-if="error" class="error">
      <div class="error-icon">⚠️</div>
      <p>{{ error }}</p>
      <button @click="retryLoadGame" class="btn">Повторить</button>
    </div>
    
    <div v-else-if="game" class="game-content">
      <GameInfo :summary="game.summary" />
      
      <div class="game-board-section">
        <div v-if="!game.gameData || game.gameData.length === 0" class="no-moves">
          <p>Нет данных о ходах партии</p>
        </div>
        <div v-else>
          <GameBoard
            :game-data="game.gameData"
            :board-size="game.summary.boardSize"
            :current-move-index="currentMoveIndex"
            @update:current-move-index="updateCurrentMoveIndex"
          />
        </div>
      </div>
      
      <div v-if="game.gameData && game.gameData.length > 0" class="move-history">
        <h3>История ходов ({{ game.gameData.length }})</h3>
        <div class="moves-scroll">
          <div 
            v-for="(state, index) in game.gameData" 
            :key="index"
            class="move-item"
            :class="{ 
              active: index === currentMoveIndex,
              'black-move': state.lastMove?.color === 'BLACK',
              'white-move': state.lastMove?.color === 'WHITE',
              'pass-move': state.lastMove?.isPass
            }"
            @click="goToMove(index)"
          >
            <span class="move-number">{{ index + 1 }}</span>
            <span class="move-info">
              {{ getMoveInfo(state.lastMove, game.summary.boardSize) }}
            </span>
            <span v-if="state.lastMove?.isPass" class="pass-label">пас</span>
          </div>
        </div>
      </div>
      
      <div v-else class="no-moves-info">
        <p>Нет информации о ходе партии</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { gameService } from '../api'
import type { Game } from '../types'
import GameInfo from '../components/GameInfo.vue'
import GameBoard from '../components/GameBoard.vue'

const route = useRoute()
const game = ref<Game | null>(null)
const loading = ref(true)
const error = ref('')
const currentMoveIndex = ref(0)

async function loadGame() {
  const gameId = route.params.id as string
  if (!gameId) {
    error.value = 'Не указан ID партии'
    loading.value = false
    return
  }
  
  loading.value = true
  error.value = ''
  game.value = null
  
  try {
    console.log(`Загружаем партию с ID: ${gameId}`)
    const response = await gameService.getGameById(gameId)
    game.value = response.game
    console.log('Партия загружена:', game.value.summary.title)
    console.log('Количество состояний:', game.value.gameData?.length || 0)
    
    if (!game.value.gameData || game.value.gameData.length === 0) {
      console.warn('Партия не содержит данных о ходе')
    }
    
    currentMoveIndex.value = 0
  } catch (err: any) {
    console.error('Ошибка при загрузке партии:', err)
    if (err.response?.status === 404) {
      error.value = 'Партия не найдена'
    } else if (err.response?.status === 500) {
      error.value = 'Ошибка сервера при загрузке партии'
    } else if (err.code === 'ECONNABORTED') {
      error.value = 'Превышено время ожидания ответа от сервера'
    } else if (err.message === 'Network Error') {
      error.value = 'Ошибка сети. Проверьте подключение'
    } else {
      error.value = `Ошибка при загрузке партии: ${err.message || 'Неизвестная ошибка'}`
    }
  } finally {
    loading.value = false
  }
}

function updateCurrentMoveIndex(index: number) {
  if (index >= 0 && game.value?.gameData && index < game.value.gameData.length) {
    currentMoveIndex.value = index
  }
}

function goToMove(index: number) {
  updateCurrentMoveIndex(index)
}

function getMoveInfo(lastMove: any, boardSize: number): string {
  if (!lastMove) return 'Начало игры'
  if (lastMove.isPass) return `${lastMove.color === 'BLACK' ? 'Ч' : 'Б'}`
  
  if (lastMove.point) {
    // Координаты в Го: A-T (пропускаем I), 1-19
    let column = lastMove.point.x
    let row = boardSize - lastMove.point.y
    
    // Преобразуем номер колонки в букву
    let columnLetter = ''
    if (column < 8) {
      columnLetter = String.fromCharCode(65 + column) // A-H
    } else {
      columnLetter = String.fromCharCode(65 + column + 1) // Пропускаем I: A-T
    }
    
    return `${lastMove.color === 'BLACK' ? 'Ч' : 'Б'}: ${columnLetter}${row}`
  }
  
  return 'Ход'
}

function retryLoadGame() {
  loadGame()
}

// Следим за изменением ID в URL
watch(() => route.params.id, (newId, oldId) => {
  if (newId !== oldId) {
    loadGame()
  }
})

// Загружаем при монтировании
onMounted(() => {
  loadGame()
})
</script>

<style scoped>
.game-view {
  padding: 24px 0;
  min-height: 100vh;
}

.back-link {
  display: inline-flex;
  align-items: center;
  margin-bottom: 24px;
  color: var(--secondary-color);
  text-decoration: none;
  font-weight: 500;
  padding: 8px 12px;
  border-radius: 4px;
  transition: background-color 0.2s;
}

.back-link:hover {
  background-color: #f0f7ff;
  text-decoration: none;
}

.loading {
  text-align: center;
  padding: 60px 20px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
}

.spinner {
  width: 50px;
  height: 50px;
  border: 4px solid #f3f3f3;
  border-top: 4px solid var(--secondary-color);
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.error {
  text-align: center;
  padding: 60px 20px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
}

.error-icon {
  font-size: 48px;
}

.error p {
  font-size: 1.1rem;
  color: #e74c3c;
  margin: 0;
}

.game-content {
  display: flex;
  flex-direction: column;
  gap: 32px;
}

.game-board-section {
  display: flex;
  justify-content: center;
}

.no-moves {
  text-align: center;
  padding: 40px;
  background: #f8f9fa;
  border-radius: 8px;
  color: #666;
}

.move-history {
  background: white;
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
}

.move-history h3 {
  margin-bottom: 16px;
  color: var(--primary-color);
  font-size: 1.3rem;
  font-weight: 600;
}

.moves-scroll {
  max-height: 300px;
  overflow-y: auto;
  padding: 4px;
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(140px, 1fr));
  gap: 10px;
}

.move-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 12px;
  border: 2px solid transparent;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s ease;
  background: #f8f9fa;
}

.move-item:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(0,0,0,0.1);
  background: #e9ecef;
}

.move-item.active {
  background: var(--secondary-color);
  color: white;
  border-color: var(--secondary-color);
}

.move-item.black-move.active {
  background: #2c3e50;
}

.move-item.white-move.active {
  background: #7f8c8d;
}

.move-number {
  font-weight: 700;
  min-width: 28px;
  font-size: 0.9rem;
}

.move-info {
  font-size: 0.9rem;
  flex: 1;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.pass-label {
  font-size: 0.8rem;
  color: #666;
  background: #e9ecef;
  padding: 2px 6px;
  border-radius: 4px;
  margin-left: auto;
}

.no-moves-info {
  text-align: center;
  padding: 20px;
  background: #f8f9fa;
  border-radius: 8px;
  color: #666;
}

.btn {
  padding: 10px 20px;
  background-color: var(--secondary-color);
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 500;
  transition: background-color 0.2s;
}

.btn:hover {
  background-color: #2980b9;
}

/* Стили для скроллбара */
.moves-scroll::-webkit-scrollbar {
  width: 8px;
}

.moves-scroll::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 4px;
}

.moves-scroll::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 4px;
}

.moves-scroll::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}

@media (max-width: 768px) {
  .moves-scroll {
    grid-template-columns: repeat(auto-fill, minmax(120px, 1fr));
    gap: 8px;
  }
  
  .move-item {
    padding: 8px 10px;
  }
}

@media (max-width: 480px) {
  .moves-scroll {
    grid-template-columns: 1fr;
  }
  
  .game-view {
    padding: 16px 0;
  }
}
</style>