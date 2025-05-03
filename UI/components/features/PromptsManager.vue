<template>
    <div class="body d-flex flex-column justify-content-center align-items-center">
        <div class="d-flex justify-content-start align-items-center" style="height: 7%; width: 100%;">
            <h1>{{ pageName }}</h1>
        </div>
        <div class="d-flex justify-content-center align-items-center" style="height: 100%; width: 100%; overflow-y: scroll;">
            <leftBar :prompts="prompts" @changePrompt="changeCurrentPrompt" />



            <div class="d-flex flex-column justify-content-center align-items-center" style="height: 100%; width: 100%; padding: 1% 1%; border-left: 1px solid black">
                <div v-if="currentPrompt !== -1" class="d-flex justify-content-between align-items-center" style="width: 100%; height: 7%; margin-bottom: 1%;">
                    <div style="width: 40%; height: 100%; overflow: hidden;" class="d-flex justify-content-start align-items-end">
                        <h4>{{ prompts[currentPrompt].fileName }}</h4>
                    </div>
                    <div style="width: 60%; height: 100%;" class="action d-flex justify-content-end align-items-center">
                        <Icon class="icon" name="material-symbols:save" size="3vh" style="color: black" @click="handleEditPrompt"/>
                        <Icon class="icon" name="material-symbols:delete" size="3vh" style="color: black" />
                        <Icon class="icon" name="material-symbols:refresh" size="3vh" style="color: black" @click="handleGetPrompts"/>
                    </div>
                </div>



                <div v-if="currentPrompt !== -1" style="width: 100%; height: 100%; overflow-y: scroll;">
                    <EditorContent :editor="editor" style="width: 100%; height: 100%;"/>
                </div>
                <div v-else>
                    <p>Veuillez sélectionner un prompt.</p>
                </div>
            </div>



        </div>
    </div>
</template>

<script setup>
    import { ref, onMounted, onBeforeUnmount } from 'vue';
    import { getPrompts, editPrompt } from '~/services/npcforge.js';
    import { useEditor, EditorContent } from '@tiptap/vue-3';
    import StarterKit from '@tiptap/starter-kit';
    import leftBar from '~/components/features/prompt/leftBar.vue'

    const prompts = ref([]);
    const currentPrompt = ref(-1);

    defineProps({
        pageName: String
    });

    const editor = useEditor({
        extensions: [StarterKit],
        content: '',
    });

    onBeforeUnmount(() => {
        // Nettoyer l'éditeur avant la destruction du composant
        editor?.destroy();
    });

    const changeCurrentPrompt = (i) => {
        currentPrompt.value = i;
        if (editor.value && prompts.value[currentPrompt.value]) {
            updateEditorContent(prompts.value[currentPrompt.value].content);
        } else {
            console.error("Le contenu du prompt est vide ou l'éditeur n'est pas prêt.");
        }
    };

    const convertTextToHtml = (text) => {
        // Remplacer les retours à la ligne par des balises <p> ou <br>
        return text.replace(/\n/g, '<p></p>');
    };

    const updateEditorContent = (content) => {
        if (editor.value) {
            editor.value.commands.setContent(convertTextToHtml(content));
        }
    };

    const handleGetPrompts = async () => {
        try {
            prompts.value = await getPrompts();
            console.log(prompts.value);
        } catch (error) {
            console.error('Erreur lors de la récupération des prompts:', error);
        }
    };

    const handleEditPrompt = async () => {
        try {
            let res = await(editPrompt(prompts.value[currentPrompt.value].fileName, prompts.value[currentPrompt.value].content));
            console.log(res)
        } catch (e) {
            console.error(e)
        }
    }

    onMounted(() => {
        handleGetPrompts();
    });
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
    cursor: pointer;
}

.icon {
    height: 200px;
}

.icon:hover {
    background-color: rgba(0, 0, 0, 0.612);
    cursor: pointer;
}
</style>