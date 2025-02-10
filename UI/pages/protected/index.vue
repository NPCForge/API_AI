<template>
    <div>
        <p>- Afficher les personnes connecté</p>
        <p>- Afficher les utilisateur enregistrer</p>
        <p>- Afficher le model utilisé</p>
        <monitoringConnected />
    </div>
</template>

<script setup>
    import monitoringConnected from '~/components/utils/monitoringConnected.vue';

    definePageMeta({
        middleware: 'auth' // Active le middleware `/protected/_middleware.ts`
    });
</script>