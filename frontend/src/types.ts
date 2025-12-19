export type Color = 'BLACK' | 'WHITE'

export type Rules = 'JAPANESE' | 'CHINESE'

export interface Point {
  x: number
  y: number
}

export interface Move {
  point?: Point
  isPass: boolean
  color: Color
}

export interface Player {
  name: string
  rank: number
}

export interface Position {
  data: number[]
}

export interface GameState {
  lastMove?: Move
  position: Position
}

export interface GameSummary {
  id: string
  boardSize: number
  title: string
  komi: number
  date: string
  isFinished: boolean
  rules: Rules
  blackPlayer: Player
  whitePlayer: Player
}

export interface Game {
  summary: GameSummary
  gameData: GameState[]
}

export interface GamesListResponse {
  games: GameSummary[]
}

export interface GameDetailResponse {
  game: Game
}

export interface ApiResponse<T> {
  data: T
  status: number
  message?: string
}