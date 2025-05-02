<template>
    <div class="d-flex justify-content-start align-items-start vh-100">
        <div class="d-flex flex-column justify-content-start align-items-start vw-15 vh-100 blue">
            <ButtonGroup :buttons="buttons" />
        </div>
        <div class="d-flex justify-content-start align-items-start vw-100 vh-100"></div>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import ButtonGroup from '~/components/ButtonGroup.vue'

const buttons = ref([
    {
        label: 'Disconnect',
        dropdown: false,
        onClick: () => { console.log('Button 1 clicked') }  // Fonction associée à Button 1
    },
    {
        label: 'Simulate',
        dropdown: true,
        menuItems: ['Action', 'Another action', 'Something else'],
        onClick: () => { console.log('Button 2 clicked') }  // Fonction associée à Button 2
    },
    {
        label: 'Entities manager',
        dropdown: false,
        onClick: () => { console.log('Button 1 clicked') }  // Fonction associée à Button 1
    },
    {
        label: 'Users manager',
        dropdown: false,
        onClick: () => { console.log('Button 1 clicked') }  // Fonction associée à Button 1
    },
    {
        label: 'Prompt manager',
        dropdown: false,
        onClick: () => { console.log('Button 1 clicked') }  // Fonction associée à Button 1
    },
    {
        label: 'Settings',
        dropdown: false,
        onClick: () => { console.log('Button 1 clicked') }  // Fonction associée à Button 1
    },
])

</script>