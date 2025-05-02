<template>
    <div class="btn-group " v-for="(button, index) in buttons" :key="index" style="width: 100%">
        <!-- Vérifier si le bouton a un menu déroulant -->
        <button v-if="button.dropdown" type="button" class="dropdown-toggle btn btn-primary" data-bs-toggle="dropdown"
            aria-expanded="false" style="width: 100%; height: 5vh; border: 0;" @click="button.onClick">
            <!-- Lier la fonction onClick -->
            {{ button.label }}
        </button>
        <button v-else type="button" class="btn btn-primary" style="width: 100%; height: 5vh; border: 0;"
            @click="button.onClick"> <!-- Lier la fonction onClick -->
            {{ button.label }}
        </button>

        <!-- Dropdown menu pour les boutons ayant une option de menu -->
        <ul v-if="button.dropdown" class="dropdown-menu">
            <li v-for="(item, index) in button.menuItems" :key="index">
                <a class="dropdown-item" href="#">{{ item }}</a>
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
</script>

<style scoped>
/* Style du bouton et du menu déroulant */
</style>