<script setup>
import { Text } from 'vue';
import CommandPrompt from './CommandPrompt.vue';
</script>

<template>
  <div class="console">
      <CommandPrompt command="tree"/>
      <!-- <CommandPrompt command="" /> -->
  </div>
</template>

<style scoped>

.console{
  display: grid;
  place-items: left;
}
</style>
