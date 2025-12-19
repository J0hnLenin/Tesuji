<template>
  <div class="game-board-container">
    <div class="board-wrapper">
      <div 
        class="board"
        :style="boardStyle"
        ref="boardRef"
        tabindex="0"
        @keydown="handleKeyDown"
      >
        <div 
          v-for="row in boardSize" 
          :key="row"
          class="row"
        >
          <div 
            v-for="col in boardSize" 
            :key="col"
            class="intersection"
            :class="{
              'has-stone': getStone(row-1, col-1),
              'black': getStone(row-1, col-1) === 1,
              'white': getStone(row-1, col-1) === 2,
              'last-move': isLastMove(row-1, col-1)
            }"
          >
            <div class="stone" v-if="getStone(row-1, col-1)"></div>
          </div>
        </div>
      </div>
    </div>
    
    <div class="board-controls">
      <button 
        @click="previousMove"
        class="btn"
        :disabled="currentMoveIndex === 0"
      >
        ← Предыдущий
      </button>
      <span class="move-counter">
        Ход {{ currentMoveIndex + 1 }} из {{ totalMoves }}
      </span>
      <button 
        @click="nextMove"
        class="btn"
        :disabled="currentMoveIndex === totalMoves - 1"
      >
        Следующий →
      </button>
    </div>
    
    <div class="debug-info" v-if="debug">
      <p>Текущий ход: {{ currentMoveIndex }}</p>
      <p>Всего ходов: {{ totalMoves }}</p>
      <p>Позиция: {{ currentPosition.length }} точек</p>
      <p>Последний ход: {{ lastMoveInfo }}</p>
      <button @click="debug = !debug">Скрыть отладку</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, onMounted, watch } from 'vue'
import type { GameState, Move } from '../types'

interface Props {
  gameData: GameState[]
  boardSize: number
  currentMoveIndex: number
}

const props = defineProps<Props>()
const emit = defineEmits<{
  'update:currentMoveIndex': [index: number]
}>()

const boardRef = ref<HTMLElement>()
const debug = ref(false)

const totalMoves = computed(() => props.gameData.length)
const currentPosition = computed(() => {
  const data = props.gameData[props.currentMoveIndex]?.position.data || []
  console.log(`Ход ${props.currentMoveIndex}: получено ${data.length} позиций, примеры:`, 
    data.slice(0, 5), '...')
  return data
})

const lastMoveInfo = computed(() => {
  const lastMove = props.gameData[props.currentMoveIndex]?.lastMove
  if (!lastMove) return 'нет'
  
  if (lastMove.isPass) return `${lastMove.color}: пас`
  if (lastMove.point) return `${lastMove.color}: (${lastMove.point.x}, ${lastMove.point.y})`
  return `${lastMove.color}: неизвестно`
})

// Используем строку для стилей вместо объекта для избежания TypeScript ошибок
const boardStyle = computed(() => {
  return {
    display: 'grid',
    gridTemplateColumns: `repeat(${props.boardSize}, 1fr)`,
    gridTemplateRows: `repeat(${props.boardSize}, 1fr)`,
    width: '100%',
    aspectRatio: '1 / 1',
    backgroundColor: '#dcb35c',
    border: '2px solid #8b4513',
    position: 'relative' as const
  }
})

function getStone(row: number, col: number): number {
  if (!currentPosition.value || currentPosition.value.length === 0) {
    return 0
  }
  
  const index = col * props.boardSize + row
  if (index >= currentPosition.value.length) {
    console.warn(`Индекс ${index} вне границ массива (${currentPosition.value.length})`)
    return 0
  }
  
  const stoneValue = currentPosition.value[index]
  
  return stoneValue || 0
}

function isLastMove(row: number, col: number): boolean {
  const lastMove = props.gameData[props.currentMoveIndex]?.lastMove
  if (!lastMove || lastMove.isPass || !lastMove.point) return false
  
  // Проверяем координаты последнего хода
  const isLast = lastMove.point.x === col && lastMove.point.y === row
  if (isLast) {
    console.log(`Последний ход на (${col},${row})`)
  }
  return isLast
}

function previousMove() {
  if (props.currentMoveIndex > 0) {
    console.log(`Переход к предыдущему ходу: ${props.currentMoveIndex} -> ${props.currentMoveIndex - 1}`)
    emit('update:currentMoveIndex', props.currentMoveIndex - 1)
  }
}

function nextMove() {
  if (props.currentMoveIndex < totalMoves.value - 1) {
    console.log(`Переход к следующему ходу: ${props.currentMoveIndex} -> ${props.currentMoveIndex + 1}`)
    emit('update:currentMoveIndex', props.currentMoveIndex + 1)
  }
}

function handleKeyDown(event: KeyboardEvent) {
  switch(event.key) {
    case 'ArrowLeft':
      event.preventDefault()
      previousMove()
      break
    case 'ArrowRight':
      event.preventDefault()
      nextMove()
      break
    case 'd':
      if (event.ctrlKey || event.metaKey) {
        event.preventDefault()
        debug.value = !debug.value
      }
      break
  }
}

// Отладочный вывод при изменении текущего хода
watch(() => props.currentMoveIndex, (newIndex, oldIndex) => {
  console.log(`Ход изменен: ${oldIndex} -> ${newIndex}`)
  console.log(`Данные позиции:`, props.gameData[newIndex]?.position.data?.slice(0, 10))
})

onMounted(() => {
  boardRef.value?.focus()
  console.log('GameBoard mounted')
  console.log('Размер доски:', props.boardSize)
  console.log('Всего ходов:', totalMoves.value)
  console.log('Текущий ход:', props.currentMoveIndex)
  
  // Отладочный вывод первых состояний
  if (props.gameData && props.gameData.length > 0) {
    for (let i = 0; i < Math.min(3, props.gameData.length); i++) {
      console.log(`Состояние ${i}:`, {
        lastMove: props.gameData[i]?.lastMove,
        positionLength: props.gameData[i]?.position.data?.length,
        first10: props.gameData[i]?.position.data?.slice(0, 10)
      })
    }
  }
})
</script>

<style scoped>
.game-board-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
}

.board-wrapper {
  width: 175%;
  display: flex;
  left: 50%;
  justify-content: center;
  overflow: hidden;
}

.board {
  cursor: pointer;
  outline: none;
}

.board:focus {
  outline: 2px solid var(--secondary-color);
  outline-offset: 2px;
}

.row {
  display: contents;
}

.intersection {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
}

.intersection::before,
.intersection::after {
  content: '';
  position: absolute;
  background-color: var(--grid-color);
}

.intersection::before {
  width: 100%;
  height: 1px;
  top: 50%;
  left: 0;
}

.intersection::after {
  width: 1px;
  height: 100%;
  top: 0;
  left: 50%;
}

.stone {
  width: 85%;
  height: 85%;
  border-radius: 50%;
  z-index: 2;
  position: relative;
}

.has-stone.black .stone {
  background-color: var(--stone-black);
  box-shadow: inset 0 0 5px rgba(255,255,255,0.1);
}

.has-stone.white .stone {
  background-color: var(--stone-white);
  box-shadow: inset 0 0 5px rgba(0,0,0,0.1);
}

.last-move::after {
  content: '';
  position: absolute;
  width: 20%;
  height: 20%;
  background-color: #ff0000;
  border-radius: 50%;
  z-index: 3;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border: 2px solid white;
  box-shadow: 0 0 3px rgba(0,0,0,0.5);
}

.board-controls {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 20px;
  width: 100%;
  flex-wrap: wrap;
}

.move-counter {
  font-weight: 500;
  min-width: 120px;
  text-align: center;
  font-size: 1.1rem;
}

.btn {
  padding: 10px 20px;
  min-width: 140px;
  background-color: var(--secondary-color);
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 16px;
  font-weight: 500;
  transition: all 0.2s;
}

.btn:hover:not(:disabled) {
  background-color: #2980b9;
  transform: translateY(-1px);
}

.btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
  transform: none;
}

.debug-info {
  margin-top: 20px;
  padding: 15px;
  background-color: #f8f9fa;
  border: 1px solid #dee2e6;
  border-radius: 8px;
  font-family: monospace;
  font-size: 14px;
  max-width: 500px;
  width: 100%;
}

.debug-info p {
  margin: 5px 0;
}

.debug-info button {
  margin-top: 10px;
  padding: 5px 10px;
  background-color: #6c757d;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

@media (max-width: 480px) {
  .board-controls {
    flex-direction: column;
    gap: 10px;
  }
  
  .btn {
    width: 100%;
  }
  
  .board {
    max-width: 95vw;
  }
}
</style>