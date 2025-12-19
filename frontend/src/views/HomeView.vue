<template>
  <div class="home-view">
    <div class="hero">
      <div class="container">
        <h1 class="hero-title">Tesuji Go</h1>
        <p class="hero-subtitle">Архив партий для изучения тесудзи</p>
        <router-link to="/games" class="btn">
          Смотреть партии
        </router-link>
      </div>
    </div>
    
    <div class="features container">
      <div class="feature-card">
        <h3>🎮 Интерактивная доска</h3>
        <p>Просматривайте партии ход за ходом с удобным управлением</p>
      </div>
      
      <div class="feature-card">
        <h3>📊 Детальная информация</h3>
        <p>Полные данные о партиях, игроках и правилах</p>
      </div>
      
      <div class="feature-card">
        <h3>📱 Адаптивный дизайн</h3>
        <p>Удобный просмотр на любых устройствах</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
</script>

<style scoped>
.home-view {
  padding-bottom: 40px;
}

.hero {
  background: linear-gradient(135deg, var(--primary-color), var(--secondary-color));
  color: white;
  padding: 60px 0;
  margin-bottom: 40px;
  text-align: center;
}

.hero-title {
  font-size: 2.5rem;
  margin-bottom: 16px;
  font-weight: 700;
}

.hero-subtitle {
  font-size: 1.25rem;
  margin-bottom: 32px;
  opacity: 0.9;
}

.features {
  display: grid;
  gap: 24px;
  grid-template-columns: 1fr;
}

.feature-card {
  background: white;
  padding: 24px;
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0,0,0,0.1);
  text-align: center;
  transition: transform 0.2s;
}

.feature-card:hover {
  transform: translateY(-4px);
}

.feature-card h3 {
  color: var(--primary-color);
  margin-bottom: 12px;
  font-size: 1.25rem;
}

.feature-card p {
  color: #666;
  line-height: 1.6;
}

@media (min-width: 768px) {
  .hero-title {
    font-size: 3rem;
  }
  
  .features {
    grid-template-columns: repeat(3, 1fr);
  }
}
</style>