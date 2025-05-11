<template>
    <div class="btn-group dropend" v-for="(button, index) in buttons" :key="index" style="width: 100%">
        <!-- Vérifier si le bouton a un menu déroulant -->
        <button v-if="button.dropdown" type="button" class="dropdown-toggle" data-bs-toggle="dropdown"
            aria-expanded="false" style="width: 100%; height: 5vh; border: 0;">
            {{ button.label }}
        </button>
        <button v-else type="button" class="" style="width: 100%; height: 5vh; border: 0;"
            @click="change(button.label)">
            {{ button.label }}
        </button>

        <!-- Dropdown menu pour les boutons ayant une option de menu -->
        <ul v-if="button.dropdown" class="dropdown-menu">
            <li v-for="(item, index) in button.menuItems" :key="index">
                <!-- Lier le clic de l'élément du menu avec une fonction change -->
                <a class="dropdown-item" href="#" @click="change(`${button.label}_${item}`)">
                    {{ item }}
                </a>
            </li>
        </ul>
    </div>
</template>

<script setup>
    defineProps({
        buttons: {
            type: Array,
            required: true
        }
    })

    // Déclare l'événement 'change' que nous allons émettre
    const emit = defineEmits()

    // Fonction de gestion du clic sur un bouton ou un élément de menu
    const change = (name) => {
        // Émettre l'événement 'change' avec le nom nettoyé de tout espace
        emit('change', name.replace(/\s+/g, '')) // Remplacer tous les espaces
    }
</script>

<style scoped>
    /* Style du bouton et du menu déroulant */
    button:hover {
        background-color: rgb(211, 211, 211);
    }
</style>