<template>
    <div class="body d-flex flex-column justify-content-center align-items-center">
        <div class="d-flex justify-content-start align-items-center" style="height: 7%; width: 100%;"><h1>{{ pageName }}</h1></div>
        <div class="d-flex justify-content-center align-items-center" style="height: 100%; width: 100%;">
            <div class="d-flex flex-column justify-content-start align-items-center" style="height: 100%; width: 17%; padding: 1% 1%;">
                <div class="file d-flex justify-content-center align-items-center" v-for="(n, i) in prompts" :key="i" @click="changeCurrentPrompt(i)">
                    {{ n.fileName }}
                </div>
            </div>
            <div class="d-flex justify-content-center align-items-center" style="height: 100%; width: 100%; padding: 1% 1%; border-left: 1px solid black">
                {{ prompts[currentPrompt].content }}
            </div>
        </div>
    </div>
</template>

<script setup>
    import { getPrompts } from '~/services/npcforge.js'  // Assure-toi que getPrompts est correctement importée

    defineProps({
        pageName: String
    })

    const prompts = ref([]);
    const currentPrompt = ref(-1)

    const changeCurrentPrompt = (i) => {
        currentPrompt.value = i
    }

    // Fonction pour gérer l'appel à getPrompts lors du clic
    const handleGetPrompts = async () => {
        prompts.value = await getPrompts();  // Attendre la réponse de getPrompts
        console.log(prompts.value);  // Afficher les prompts récupérés
    }

    onMounted(() => { handleGetPrompts() })
</script>

<style scoped>
    .body {
        height: 100%;
        width: 100%;
        overflow-y: scroll;
        overflow-x: hidden;
    }

    .file {
        width: 100%;
        height: 30px;
        background-color: rgb(234, 234, 234);
        border: 1px solid black;
        transition: all 0.2s;
    }

    .file:hover {
        background-color: rgb(198, 198, 198);
        cursor:pointer;
    }
</style>