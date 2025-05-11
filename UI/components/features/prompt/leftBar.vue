<template>
    <div class="file-container">
        <div class="header">
            <Icon class="icon" name="material-symbols:add" size="2vh" style="color: black" @click="handleCreate"/>
        </div>
        <div class="file-list" v-for="(n, i) in prompts" :key="i" @click="changeCurrentPrompt(i)">
            {{ n.fileName }}
        </div>
    </div>
</template>

<script setup>
    import { ref } from 'vue';
    import { createPrompt } from '~/services/npcforge.js';

    // Définir les props venant du parent (prompts)
    const props = defineProps({
        prompts: {
            type: Array,
            required: true
        }
    });

    // Événements personnalisés pour le changement de prompt
    const emits = defineEmits(['changePrompt', 'addPrompt']);

    // Fonction pour gérer le changement de prompt
    const changeCurrentPrompt = (i) => {
        emits('changePrompt', i);
    };

    // Fonction pour la création d'un prompt
    const handleCreate = async () => {
        let fileName = prompt("Please enter the file name:", "myfile");
        if (!fileName) return; // Si le nom du fichier est vide, on arrête la fonction.

        // Assurer que le nom du fichier a une extension ".txt"
        fileName = fileName.replace(/\.txt$/, "") + ".txt";

        // Si le nom du fichier est invalide (ex. contient des caractères spéciaux ou est vide)
        if (!fileName.match(/^[\w\-\.]+$/)) {
            alert("Le nom du fichier n'est pas valide.");
            return;
        }

        console.log("Nom du fichier créé :", fileName);

        try {
            // Appeler la fonction pour créer le prompt avec le nom de fichier
            await createPrompt(fileName);

            // Émettre l'événement 'addPrompt' pour ajouter le nouveau prompt au parent
            emits('addPrompt', { fileName });

            // Notifier le parent pour changer le prompt
            const newIndex = props.prompts.length; // Utilise l'index de la liste mise à jour
            changeCurrentPrompt(newIndex); // Change le prompt sélectionné après l'ajout
        } catch (error) {
            console.error("Erreur lors de la création du prompt:", error);
        }
    };
</script>

<style scoped>
    /* Styles déjà existants */
    .file-container {
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        align-items: center;
        height: 100%;
        width: 17%;
        padding: 1%;
    }

    .header {
        display: flex;
        justify-content: flex-end;
        align-items: center;
        width: 100%;
        height: 5%;
    }

    .file-list {
        display: flex;
        justify-content: center;
        align-items: center;
        width: 100%;
        height: 30px;
        background-color: rgb(234, 234, 234);
        border: 1px solid black;
        transition: all 0.2s;
        cursor: pointer;
    }

    .file-list:hover {
        background-color: rgb(198, 198, 198);
    }

    .icon {
        cursor: pointer;
    }
</style>
