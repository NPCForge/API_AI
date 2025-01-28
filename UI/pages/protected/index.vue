<template>
    <div>
        <h1>Hello world</h1>
    </div>
</template>

<script setup>
    definePageMeta({
        middleware: 'auth' // Active le middleware `/protected/_middleware.ts`
    });
</script>