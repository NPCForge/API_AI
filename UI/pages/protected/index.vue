<template>
    <div class="body_index">
        <p>- Afficher les personnes connecté</p>
        <p>- Afficher les utilisateur enregistrer</p>
        <p>- Afficher le model utilisé</p>
        <monitoringConnected />
    </div>
</template>

<style scoped>
    .body_index {
        background-color: rgb(168, 168, 168);
        height: auto;
        min-height: 100vh;
    }
</style>

<script setup>
    import monitoringConnected from '~/components/utils/monitoringConnected.vue';

    definePageMeta({
        middleware: 'auth' // Active le middleware `/protected/_middleware.ts`
    });
</script>