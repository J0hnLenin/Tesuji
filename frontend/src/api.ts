import axios from 'axios'
import type { GamesListResponse, GameDetailResponse, GameSummary, Game } from './types'

const api = axios.create({
  baseURL: '/api/v1',
  headers: {
    'Content-Type': 'application/json'
  },
  timeout: 10000 // 10 секунд таймаут
})

// Интерсептор для обработки ошибок
api.interceptors.response.use(
    (  response: any) => response,
    (  error: { response: { status: any }; message: any }) => {
    console.error('API Error:', error.response?.status, error.message)
    return Promise.reject(error)
  }
)

export const gameService = {
  // Получение списка партий (только summary)
  async getGames(page: number = 1, limit: number = 10): Promise<GamesListResponse> {
    try {
      const response = await api.get<GamesListResponse>('/games', {
        params: { page, limit }
      })
      return response.data
    } catch (error) {
      console.error('Error fetching games list:', error)
      throw error
    }
  },

  // Получение полной информации о конкретной партии
  async getGameById(id: string): Promise<GameDetailResponse> {
    try {
      const response = await api.get<GameDetailResponse>(`/games/${id}`)
      return response.data
    } catch (error) {
      console.error(`Error fetching game ${id}:`, error)
      throw error
    }
  },

  // Получение summary партии (если нужно отдельно)
  async getGameSummary(id: string): Promise<GameSummary> {
    try {
      const response = await api.get<GameSummary>(`/games/${id}/summary`)
      return response.data
    } catch (error) {
      console.error(`Error fetching game summary ${id}:`, error)
      throw error
    }
  }
}