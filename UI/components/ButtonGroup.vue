<template>
    <div class="d-flex flex-column w-100 gap-2">
        <div v-for="(button, index) in buttons" :key="index" class="btn-group dropend w-100">
            <!-- Bouton avec dropdown -->
            <button v-if="button.dropdown" type="button" class="btn btn-light text-start dropdown-toggle w-100 py-2"
                data-bs-toggle="dropdown" aria-expanded="false">
                {{ button.label }}
            </button>

            <!-- Bouton simple -->
            <button v-else type="button" class="btn btn-light text-start w-100 py-2" @click="change(button.label)">
                {{ button.label }}
            </button>

            <!-- Menu dropdown -->
            <ul v-if="button.dropdown" class="dropdown-menu shadow-sm">
                <li v-for="(item, itemIndex) in button.menuItems" :key="itemIndex">
                    <a class="dropdown-item" href="#" @click.prevent="change(`${button.label}_${item}`)">
                        {{ item }}
                    </a>
                </li>
            </ul>
        </div>
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
    const emit = defineEmits(['change'])

    // Fonction de gestion du clic sur un bouton ou un élément de menu
    const change = (name) => {
        // Émettre l'événement 'change' avec le nom nettoyé de tout espace
        emit('change', name.replace(/\s+/g, '')) // Remplacer tous les espaces
    }
    </script>

    <style scoped>
    /* Styles supplémentaires si nécessaire */
    .btn-light {
        background-color: transparent;
        border: none;
    }

    .btn-light:hover {
        background-color: rgba(211, 211, 211, 0.5);
    }

    /* Assurer que les boutons ont une hauteur uniforme */
    .btn {
        height: 5vh;
    }
</style>