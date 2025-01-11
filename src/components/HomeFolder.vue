<script setup>
import {onMounted, onUnmounted, ref} from 'vue'
import CommandPrompt from './CommandPrompt.vue';
const cmds = ref(['resh ./welcome.sh'])
function enter(event) {
  if (event.key =="Enter") {
    cmds.value.push("cmdbox")
  }

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
      <span v-for="(cmd, index) in cmds" :key="cmd.key">
        <CommandPrompt v-model="cmds[index]" />
      </span>
  </div>
</template>

<style scoped>

.console{
  display: grid;
  place-items: left;
}
</style>
