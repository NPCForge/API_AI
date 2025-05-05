<template>
    <div style="height: 100%; width: 100%;">
        <h3>Simulation route "Connect"</h3>
        <div class="hb"></div>
        <div class="d-flex justify-content-center align-items-start rounded"
            style="margin: 1%; height: 30vh; position: relative; overflow: hidden;">
            <Loader v-if="isLoading" :Text="LoadingMessage" />
            <div style="width: 100%; padding: 1%;" class="d-flex flex-column justify-content-start align-items-start">
                <label for="Action">Action</label>
                <input type="text" placeholder="Action" disabled value="Connect" name="Action">
                <label for="Identifier">Identifier</label>
                <input type="text" placeholder="Identifier" name="Identifier" v-model="identifier">
                <label for="Password">Password</label>
                <input type="text" placeholder="Password" name="Password" v-model="password">
            </div>
            <div style="width: 100%; padding: 1%;" class="d-flex justify-content-end align-items-start">
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
    import { connect } from '~/services/npcforge'

    const LoadingMessage = ref("Loading")
    const isLoading = ref(false)

    const identifier = ref("User42")
    const password = ref("Password")
    const res = ref("")

    onMounted(() => {
        // Optionally highlight syntax for code blocks if you use them
        hljs.highlightAll();
    })

    const RunHttp = async () => {
        try {
            isLoading.value = true
            LoadingMessage.value = "Running Http Request"
            res.value = await connect(identifier.value, password.value, true)
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

    // Format JSON with indentation of 4 spaces
    const formattedJson = computed(() => {
        try {
            return JSON.stringify(JSON.parse(res.value), null, 4);
        } catch (e) {
            return res.value;
        }
    })
</script>
