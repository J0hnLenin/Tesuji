<template>
  <div class="game-info card">
    <h2 class="game-title">{{ summary.title }}</h2>
    
    <div class="game-meta">
      <div class="meta-item">
        <span class="meta-label">Дата:</span>
        <span class="meta-value">{{ formatDate(summary.date) }}</span>
      </div>
      <div class="meta-item">
        <span class="meta-label">Доска:</span>
        <span class="meta-value">{{ summary.boardSize }}×{{ summary.boardSize }}</span>
      </div>
      <div class="meta-item">
        <span class="meta-label">Коми:</span>
        <span class="meta-value">{{ summary.komi }}</span>
      </div>
      <div class="meta-item">
        <span class="meta-label">Правила:</span>
        <span class="meta-value">{{ formatRules(summary.rules) }}</span>
      </div>
    </div>
    
    <div class="players">
      <div class="player black-player">
        <h3>Чёрные</h3>
        <div class="player-name">{{ summary.blackPlayer.name }}</div>
        <div class="player-rank">{{ formatRank(summary.blackPlayer.rank) }}</div>
      </div>
      
      <div class="vs">vs</div>
      
      <div class="player white-player">
        <h3>Белые</h3>
        <div class="player-name">{{ summary.whitePlayer.name }}</div>
        <div class="player-rank">{{ formatRank(summary.whitePlayer.rank) }}</div>
      </div>
    </div>
    
    <div class="game-status">
      <span :class="['status-badge', summary.isFinished ? 'finished' : 'in-progress']">
        {{ summary.isFinished ? 'Завершена' : 'В процессе' }}
      </span>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { GameSummary } from '../types'

interface Props {
  summary: GameSummary
}

defineProps<Props>()

function formatDate(dateString: string): string {
  return new Date(dateString).toLocaleDateString('ru-RU')
}

function formatRules(rules: string): string {
  const rulesMap: Record<string, string> = {
    'JAPANESE': 'Японские',
    'CHINESE': 'Китайские'
  }
  return rulesMap[rules] || rules
}

function formatRank(rank: number): string {
  if (rank === 0) return 'Неизвестно'
  if (rank > 0) return `${rank} кю`
  return `${Math.abs(rank)} дан`
}
</script>

<style scoped>
.game-info {
  width: 100%;
}

.game-title {
  font-size: 1.5rem;
  margin-bottom: 16px;
  color: var(--primary-color);
  word-break: break-word;
}

.game-meta {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 12px;
  margin-bottom: 20px;
}

.meta-item {
  display: flex;
  flex-direction: column;
}

.meta-label {
  font-size: 0.9rem;
  color: #666;
  margin-bottom: 4px;
}

.meta-value {
  font-weight: 500;
}

.players {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20px;
  padding: 16px;
  background-color: #f8f9fa;
  border-radius: 8px;
}

.player {
  flex: 1;
  text-align: center;
}

.player h3 {
  font-size: 1rem;
  margin-bottom: 8px;
  color: var(--primary-color);
}

.player-name {
  font-weight: 600;
  margin-bottom: 4px;
}

.player-rank {
  font-size: 0.9rem;
  color: #666;
}

.vs {
  padding: 0 16px;
  font-weight: bold;
  color: #666;
}

.black-player h3 {
  color: var(--stone-black);
}

.white-player h3 {
  color: var(--stone-white);
  -webkit-text-stroke: 0.5px #000;
  text-stroke: 0.5px #000;
}

.game-status {
  text-align: center;
}

.status-badge {
  display: inline-block;
  padding: 6px 12px;
  border-radius: 20px;
  font-size: 0.9rem;
  font-weight: 500;
}

.status-badge.finished {
  background-color: #e8f5e9;
  color: #2e7d32;
}

.status-badge.in-progress {
  background-color: #fff3e0;
  color: #ef6c00;
}

@media (min-width: 768px) {
  .game-meta {
    grid-template-columns: repeat(4, 1fr);
  }
}
</style>