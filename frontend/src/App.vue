<template>
  <div id="app">
    <Header></Header>
    <main>
      <router-view />
    </main>
  </div>
</template>

<script setup lang="ts">
import type Header from './components/Header.vue'
</script>

<style>
#app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

main {
  flex: 1;
}
</style>