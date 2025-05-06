<template>
    <button class="btn btn-secondary envButton" type="button" data-bs-toggle="offcanvas" data-bs-target="#offcanvasRight" aria-controls="offcanvasRight">ENV</button>

    <div class="offcanvas offcanvas-end" tabindex="-1" id="offcanvasRight" aria-labelledby="offcanvasRightLabel">
        <div class="offcanvas-header">
            <h5 class="offcanvas-title" id="offcanvasRightLabel">Environment Variable</h5>
            <button type="button" class="btn-close" data-bs-dismiss="offcanvas" aria-label="Close"></button>
        </div>
        <div class="offcanvas-body">
            <div class="input-group mb-3">
                <span class="input-group-text" id="basic-addon2">Auth Token</span>
                <input type="text" class="form-control" placeholder="Authorization Token" aria-label="Authorization Token" aria-describedby="basic-addon2" v-model="AuthorizationToken" @change="updateEnv('AuthorizationToken', AuthorizationToken)">
            </div>
            <div class="input-group mb-3">
                <span class="input-group-text" id="basic-addon2">API Token</span>
                <input type="text" class="form-control" placeholder="Api Token" aria-label="Api Token" aria-describedby="basic-addon2" v-model="ApiToken" @input="updateEnv('ApiToken', ApiToken)">
            </div>
        </div>
    </div>
</template>

<style scoped>
    .envButton {
        position: fixed;
        top: 1%;
        right: 1%;
    }
</style>

<script setup>
    import { setEnvVariable, getEnvVariable } from '~/services/env';

    const AuthorizationToken = ref("")
    const ApiToken = ref("")

    onMounted(() => {
        try {
            let tmp = getEnvVariable('AuthorizationToken')
            if (tmp) {
                AuthorizationToken.value = tmp
            }
            tmp = getEnvVariable('ApiToken')
            if (tmp) {
                ApiToken.value = tmp
            }
        } catch (e) {
            console.error(e)
        }
    })

    const updateEnv = (store, value) => {
        setEnvVariable(store, value)
    }
</script>
