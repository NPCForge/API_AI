<template>
    <div class="d-flex flex-column justify-content-center align-items-center vh-100">
        <h1>Welcome</h1>
        <input type="text" class="input" v-model="identifier" placeholder="Identifier">
        <input type="password" class="input" v-model="password" placeholder="Password">
        <button class="btn btn-primary" @click="connection">Connexion</button>
    </div>
</template>

<script setup>
    import { connect } from "~/services/npcforge.js"
    const identifier = ref("");
    const password = ref("");

    const connection = () => {
        console.log(identifier.value, password.value)
        connect(identifier.value, password.value)
    }
</script>

<style scoped>
    .input {
        width: 20vw;
        background-color: rgba(255, 255, 255, 0);
        height: 5vh;
        margin: 2vh 0;
    }

    button {
        width: 20vw;
        /* background-color: rgba(255, 255, 255, 0); */
        height: 5vh;
        margin: 2vh 0;
    }
</style>