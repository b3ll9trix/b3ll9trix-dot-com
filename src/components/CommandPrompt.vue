<script setup>
// defineProps({
//     command: {
//         type: String,
//         required: true
//     }
// })
const cmdbox = defineModel()
</script>

<template>
    <div class="prompt">
        <p>
            <span class="green">b3ll9trix</span>
            <span class="colon">:</span>
            <span class="path">~</span>
            <span class="dollar">$</span>
            <span> <input class="cmd" v-model="cmdbox"></span>
            <!-- <span class="cursor">┃</span> -->
        </p>


    </div>
</template>

<style scoped>
p {
    font-family: monospace;
}
.prompt {
    display: inline-block;
}
.cmd {
  white-space: nowrap;
  overflow-x: hidden;
  text-align: left;
  color:rgba(235, 235, 235, 0.64);
  min-width: 100vh;
  vertical-align: middle;
  font:inherit;
  font-size: 15px;
  outline: none;
  border: none;
  padding:3px;
  height: 3vh;
  min-height: 100%;
  background-color: #181818;
  resize: none;
}

</style>
