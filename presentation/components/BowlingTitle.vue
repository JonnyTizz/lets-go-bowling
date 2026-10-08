<script setup lang="ts">
import goLogo from '../assets/go-logo.svg'
</script>

<template>
  <h1 class="bowling-title">
    Let's <img :src="goLogo" alt="Go" class="go-logo"> Bowling
  </h1>
</template>

<style scoped>
.bowling-title {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.35em;
  font-size: 3.2rem;
  margin-bottom: 0;
}
.go-logo {
  height: 0.95em;
  /* the logo's speed lines make it look off-centre; nudge it back */
  margin-left: -0.15em;
}
</style>
