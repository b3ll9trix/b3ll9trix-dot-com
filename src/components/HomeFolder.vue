<script setup>
import {onMounted, onUnmounted, ref} from 'vue'
import CommandPrompt from './CommandPrompt.vue';
const cmds = ref(['resh ./welcom.sh'])
function enter() {
  cmds.value.push('')
}
onMounted(() => {
  window.addEventListener('keyup', enter )
})

onUnmounted(() => {
  window.removeEventListener('keyup', enter)
})
</script>

<template>
  <div class="console">
      <span v-for="cmd in cmds" :key="cmd.id">
        <CommandPrompt :command="cmd" />
      </span>

  </div>
</template>

<style scoped>

.console{
  display: grid;
  place-items: left;
}
</style>
