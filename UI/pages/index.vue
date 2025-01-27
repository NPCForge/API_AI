<template>
    <div class="body d-flex justify-content-center align-items-center">
        <div v-if="loading">Loading...</div>
        <Connexion v-else-if="can" />
        <Register v-else @emitCan="changeCan" />
    </div>
</template>

<style scoped>
    .body {
        background-color: rgb(229, 229, 229);
        width: 100vw;
        height: 100vh;
        position: fixed;
    }
</style>

<script setup>
    import { isSettup } from '~/services/AuthServices';
    import Connexion from '~/components/connexion/connexion.vue';
    import Register from '~/components/connexion/register.vue';

    const can = ref(false);
    const loading = ref(true); // Add a loading state

    const checkSetup = async () => {
        console.log("checkSetup start")
        try {
            const result = await isSettup();
            can.value = result.value; // Set the value of `can`
        } catch (error) {
            console.error("Error checking setup:", error);
            can.value = false; // Default fallback
        } finally {
            loading.value = false; // Stop loading after the check
        }
    };

    // Perform the setup check on component mount
    checkSetup();

    const changeCan = () => {
        can.value = !can.value;
    };
</script>
