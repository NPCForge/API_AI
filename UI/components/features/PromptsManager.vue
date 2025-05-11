<template>
    <div class="body d-flex flex-column justify-content-center align-items-center">
        <div class="d-flex justify-content-start align-items-center" style="height: 7%; width: 100%;">
            <h1>{{ pageName }}</h1>
        </div>
        <div class="d-flex justify-content-center align-items-center"
            style="height: 100%; width: 100%; overflow-y: scroll;">
            <leftBar :prompts="prompts" @changePrompt="changeCurrentPrompt" @addPrompt="handleAddPrompt"/>

            <div class="d-flex flex-column justify-content-center align-items-center"
                style="height: 100%; width: 100%; padding: 1% 1%; border-left: 1px solid black">
                <!-- header -->
                <div v-if="currentPrompt !== -1 && prompts[currentPrompt]"
                    class="d-flex justify-content-between align-items-center"
                    style="width: 100%; height: 7%; margin-bottom: 1%;">
                    <div style="width: 40%; height: 100%; overflow: hidden;"
                        class="d-flex justify-content-start align-items-end">
                        <h4>{{ prompts[currentPrompt].fileName }}</h4>
                    </div>
                    <div style="width: 60%; height: 100%;" class="action d-flex justify-content-end align-items-center">
                        <Icon v-if="isModified" class="icon"
                            name="material-symbols:save" size="3vh" style="color: black" @click="handleEditPrompt" />
                        <Icon class="icon" name="material-symbols:delete" size="3vh" style="color: black"
                            @click="handleRemovePrompt(prompts[currentPrompt].fileName)" />
                        <Icon class="icon" name="material-symbols:refresh" size="3vh" style="color: black"
                            @click="handleGetPrompts" />
                    </div>
                </div>
                <!-- body -->
                <div v-if="currentPrompt !== -1 && prompts[currentPrompt]"
                    style="width: 100%; height: 100%; overflow-y: scroll;">
                    <EditorContent :editor="editor" style="width: 100%; height: 100%;" />
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
    import { getPrompts, editPrompt, removePrompt } from '~/services/npcforge.js';
    import { useEditor, EditorContent } from '@tiptap/vue-3';
    import StarterKit from '@tiptap/starter-kit';
    import leftBar from '~/components/features/prompt/leftBar.vue'

    const prompts = ref([]);
    const currentPrompt = ref(-1);
    const isModified = ref(false)

    defineProps({
        pageName: String
    });

    const editor = useEditor({
        extensions: [StarterKit],
        content: '',
        onUpdate: ({ editor }) => {
            if (editor.getText() !== prompts.value[currentPrompt.value].content) {
                isModified.value = true;
            } else {
                isModified.value = false;
            }
        },
    });

    onBeforeUnmount(() => {
        editor.value?.destroy()
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
        // Remplacer les retours à la ligne par <br /> sans générer de nouveaux paragraphes
        return text.replace(/\n/g, '<br />');
    };


    const updateEditorContent = (content) => {
        if (editor.value) {
            editor.value.commands.setContent(convertTextToHtml(content));
        }
    };

    const handleAddPrompt = (newPrompt) => {
        prompts.value.push(newPrompt);
    };

    const handleGetPrompts = async () => {
        try {
            prompts.value = await getPrompts();
        } catch (error) {
            console.error('Erreur lors de la récupération des prompts:', error);
        }
    };

    const handleEditPrompt = async () => {
        try {
            if (prompts.value[currentPrompt.value]) {
                // Récupère le texte de l'éditeur
                const editorText = editor.value.getText();
                // Convertir le texte en HTML (en utilisant la nouvelle méthode)
                const htmlContent = convertTextToHtml(editorText);
                // Enregistrer les changements en envoyant le contenu HTML
                await editPrompt(prompts.value[currentPrompt.value].fileName, htmlContent);
                // Recharger les prompts après l'enregistrement
                handleGetPrompts();
            }
        } catch (e) {
            console.error(e);
        }
    };

    const handleRemovePrompt = async (name) => {
        try {
            await removePrompt(name);
            currentPrompt.value = -1
            handleGetPrompts();
        } catch (e) {
            console.error(e);
        }
    };

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