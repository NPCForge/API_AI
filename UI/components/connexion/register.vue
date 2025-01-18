<template>
    <!-- Get the API Password -->
    <div v-if="valide != 'Valide'" class="Register shadow d-flex flex-column justify-content-center align-items-center">
        <h3>Key Verification</h3>

        {{ valide }}
        <p>VDCAjPZ8jhDmXfsSufW2oZyU8SFZi48dRhA8zyKUjSRU3T1aBZ7E8FFIjdEM2X1d</p>
        <input
            v-model="password"
            class="form-control input"
            type="text"
            placeholder="API_KEY"
            autocomplete="off"
        >

        <button @click="verifyToken" :disabled="!password.length" class="input btn btn-primary">Register</button>
    </div>
    <!-- Register the first Admin -->
    <div v-if="valide == 'Valide'" class="Register shadow d-flex flex-column justify-content-center align-items-center">
        <h3>First Register</h3>
        <p>{{ message }}</p>
        <input
            v-model="username"
            class="form-control input"
            type="text"
            placeholder="Username"
            autocomplete="off"
        >

        <input
            v-model="password"
            class="form-control input"
            type="password"
            placeholder="Password"
            autocomplete="off"
        >

        <button @click="register" :disabled="!password.length" class="input btn btn-primary">Register</button>
    </div>
</template>

<style scoped>
    .input {
        width: 90%;
        margin-bottom: 10%;
    }

    .Register {
        border-radius: 8px;
        height: 50vh;
        width: 25vw;
        background-color: white;
    }
</style>

<script setup>
    import { firstRegister } from '~/services/AuthServices';

    const config = useRuntimeConfig().public;
    const username = ref("");
    const password = ref("");
    const valide = ref("");
    const message = ref("");
    const emit = defineEmits(['emitCan'])

    const verifyToken = () => {
        if (password.value && config.API_KEY_REGISTER === password.value) {
            valide.value = "Valide";
            password.value = ""
        } else {
            valide.value = "Not Valide";
        }
    };

    const register = async () => {
        const success = await firstRegister(username.value, password.value);
        if (success) {
            message.value = "Utilisateur créé avec succès !";
            emit('emitCan')
        } else {
            message.value = "Erreur lors de la création de l'utilisateur.";
        }
    };
</script>
