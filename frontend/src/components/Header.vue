<template>
  <header class="header">
    <nav class="nav">
      <div class="container">
        <router-link to="/" class="logo">
          Tesuji Go
        </router-link>
        <div class="nav-links">
          <router-link 
            to="/" 
            class="nav-link"
            :class="{ active: $route.name === 'home' }"
          >
            Главная
          </router-link>
          <router-link 
            to="/games" 
            class="nav-link"
            :class="{ active: $route.name === 'games' }"
          >
            Партии
          </router-link>
        </div>
      </div>
    </nav>
  </header>
</template>

<script setup lang="ts">
</script>

<style scoped>
.header {
  background-color: white;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  position: sticky;
  top: 0;
  z-index: 1000;
}

.nav {
  padding: 12px 0;
}

.nav .container {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.logo {
  font-size: 1.5rem;
  font-weight: bold;
  color: var(--primary-color);
  text-decoration: none;
}

.nav-links {
  display: flex;
  gap: 24px;
}

.nav-link {
  color: var(--text-color);
  text-decoration: none;
  padding: 8px 0;
  position: relative;
}

.nav-link.active {
  color: var(--secondary-color);
}

.nav-link.active::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 2px;
  background-color: var(--secondary-color);
}

@media (min-width: 768px) {
  .nav {
    padding: 16px 0;
  }
}
</style>