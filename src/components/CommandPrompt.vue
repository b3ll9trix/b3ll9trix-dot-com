<script setup>
defineProps({
    command: {
        type: String,
        required: true
    }
})
</script>

<template>
    <div class="prompt">
        <p>
            <span class="green">b3ll9trix</span>
            <span class="colon">:</span>
            <span class="path">~</span>
            <span class="dollar">$</span>
            <span> {{ command }}</span>
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

</style>
