<template>
    <div class="Connexion shadow d-flex flex-column justify-content-center align-items-center">
        <h3 class="mb-3">Connexion</h3>
        <p>{{ valide }}</p>
        <input v-model="username" class="form-control input" type="text" placeholder="Username" autocomplete="off">
        <input v-model="password" class="form-control input" type="password" placeholder="Password" autocomplete="off">
        <button @click="connect" class="input btn btn-primary">Connexion</button>
    </div>
</template>

<style scoped>
    .input {
        width: 90%;
        margin-bottom: 10%;
    }

    .Connexion {
        border-radius: 8px;
        height: 50vh;
        width: 25vw;
        background-color: white;
    }
</style>

<script setup>
    import { login } from '~/services/AuthServices';
    const username = ref("");
    const password = ref("");
    const valide = ref("");

    const connect = () => {
        if (username.value != "" || password.value != "") {
            if (login(username.value, password.value))
                valide.value = "valide"
            else
                valide.value = "not valide"
        }
    }
</script>
