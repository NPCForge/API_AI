<template>
    <div class="body">
        <h3>{{ pageName }}</h3>
        
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
