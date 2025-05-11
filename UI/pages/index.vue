<template>
    <div class="d-flex flex-column justify-content-center align-items-center vh-100">
        <h1>Welcome</h1>

        <input type="text" class="input" v-model="identifier" placeholder="Identifier" autocomplete="username" name="username">

        <input type="password" class="input" v-model="password" placeholder="Password" autocomplete="current-password" name="password">

        <button class="btn btn-primary" @click="connection">Connexion</button>
    </div>
</template>


<script setup>
    import { connect } from "~/services/npcforge.js"
    import { ref } from 'vue'
    import { useRouter } from 'vue-router'

    const router = useRouter()

    const identifier = ref("")
    const password = ref("")

    const connection = async () => {
        try {
            const res = await connect(identifier.value, password.value)
            // console.log(res)
            if (res.Status === "Success") {
                router.push('/Secure')
            } else {
                console.error('Connexion échouée')
            }
        } catch (error) {
            console.error('Erreur lors de la connexion:', error)
        }
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