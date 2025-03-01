<template>
    <div class="body_index">
        <p @click="TestConnexion">- Afficher les personnes connecté {{ isConnected }}</p>
        <p>- Afficher les utilisateur enregistrer</p>
        <p>- Afficher le model utilisé</p>
        <monitoringConnected class="component"/>
        <monitoringPrompt class="component"/>
    </div>
</template>

<style scoped>
    .body_index {
        background-color: rgb(168, 168, 168);
        height: auto;
        min-height: 100vh;
    }
    .component {
        margin: 1%;
    }
</style>

<script setup>
    import monitoringConnected from '~/components/utils/monitoringConnected.vue';
    import monitoringPrompt from '~/components/utils/monitoringPrompt.vue';
    import { isUp } from '~/services/ApiServices';

    const isConnected = ref(false);

    const TestConnexion = async () => {
        isConnected.value = await isUp()
    }

    definePageMeta({
        middleware: 'auth' // Active le middleware `/protected/_middleware.ts`
    });

    onMounted(() => {
        TestConnexion()
    })
</script>