<template>
    <div style="height: 100%; width: 100%;">
        <h3>Simulation route "Remove Entity"</h3>
        <div class="hb"></div>
        <div class="d-flex justify-content-center align-items-start rounded"
            style="margin: 1%; height: 30vh; position: relative; overflow: hidden;">
            <Loader v-if="isLoading" :Text="LoadingMessage" />
            <div style="width: 100%; padding: 1%;" class="d-flex justify-content-start align-items-start">
                <div class="d-flex flex-column justify-content-start align-items-start">
                    <label for="Action">Action</label>
                    <input type="text" placeholder="Action" disabled value="RemoveEntity" name="Action">
                    <label for="AuthorizationToken">Authorization Token</label>
                    <input type="text" placeholder="Authorization Token" name="AuthorizationToken" v-model="AuthorizationToken">
                    <label for="Checksum">Checksum</label>
                    <input type="text" placeholder="Checksum" name="Checksum" v-model="Checksum">
                </div>
                <!-- <div class="mx-3 d-flex flex-column justify-content-start align-items-start"></div> -->
            </div>
            <div style="width: 75%; padding: 1%;" class="d-flex justify-content-end align-items-start">
                <button class="btn btn-primary mx-1" @click="RunHttp">Run HTTP Request</button>
                <button class="btn btn-warning" @click="RunWs">Run WebSocket Request</button>
            </div>
        </div>

        <div v-if="res !== ''" style="width: 100%; padding: 1%; border: 1px solid black;" class="rounded">
            <p>Request result:</p>
            <pre><code class="language-json">{{ formattedJson }}</code></pre>
        </div>

    </div>
</template>

<style scoped>
    code {
        border-radius: 5px;
        padding: 12px;
        font-size: 16px;
        color: #333;
        font-family: "Courier New", Courier, monospace;
        white-space: pre-wrap;
        word-wrap: break-word;
        overflow-x: auto;
        margin-top: 10px;
    }

    pre {
        margin: 0;
        padding: 0;
    }

    p {
        font-weight: bold;
        margin-bottom: 5px;
    }
</style>

<script setup>
    import { onMounted, ref, computed } from 'vue'
    import hljs from 'highlight.js'
    import 'highlight.js/styles/github.css'
    import Loader from '~/components/Loader.vue'
    import { RemoveEntity } from '~/services/npcforge'
    import { getEnvVariable } from '~/services/env'

    const LoadingMessage = ref("Loading")
    const isLoading = ref(false)
    const res = ref("")

    const AuthorizationToken = ref("")
    const Checksum = ref("")

    onMounted(() => {
        try {
            hljs.highlightAll();
            const tmp = getEnvVariable('AuthorizationToken')
            if (tmp)
                AuthorizationToken.value = tmp
        } catch (e) {
            console.error(e)
        }
    })

    const RunHttp = async () => {
        try {
            isLoading.value = true
            LoadingMessage.value = "Running Http Request"
            res.value = await RemoveEntity(Checksum.value, AuthorizationToken.value)
        } catch (e) {
            console.error(e)
        } finally {
            isLoading.value = false
        }
    }

    const RunWs = () => {
        isLoading.value = true
        LoadingMessage.value = "Running WebSocket Request"
    }

    const formattedJson = computed(() => {
        try {
            return JSON.stringify(JSON.parse(res.value), null, 4);
        } catch (e) {
            return res.value;
        }
    })
</script>
