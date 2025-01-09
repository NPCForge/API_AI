<template>
    <div class="Connexion shadow d-flex flex-column justify-content-center align-items-center">
        <input class="form-control input" type="text" placeholder="Username" autocomplete="off">
        <input class="form-control input" type="password" placeholder="Password" autocomplete="off">
        <button class="input btn btn-primary">Connexion</button>
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
