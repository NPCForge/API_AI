<template>
    <div class="body">
        <h1>{{ pageName }}</h1>
    </div>
</template>

<script setup>
    defineProps({
        pageName: String
    })
</script>

<style scoped>
    .body {
        height: 100%;
        width: 100%;
        overflow-y: scroll;
        overflow-x: hidden;
    }
</style>
