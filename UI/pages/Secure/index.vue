<template>
    <div class="d-flex justify-content-start align-items-start vh-100">
        <div class="d-flex flex-column justify-content-start align-items-start vw-15 vh-100">
            <ButtonGroup :buttons="buttons" @change="changePage"/>
        </div>
        <div class="d-flex justify-content-start align-items-start vw-100 vh-100 red">
            <Settings
                v-if="currentPage === 'Settings'"
                :pageName="'Settings'"
            />
            <EntitiesManager
                v-if="currentPage === 'EntitiesManager'"
                :pageName="'Entities Manager'"
            />
            <PromptsManager
                v-if="currentPage === 'PromptsManager'"
                :pageName="'Prompts Manager'"
            />
            <UsersManager
                v-if="currentPage === 'UsersManager'"
                :pageName="'Users Manager'"
            />
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import ButtonGroup from '~/components/ButtonGroup.vue'

import Settings from "~/components/features/Settings.vue"
import EntitiesManager from "~/components/features/EntitiesManager.vue"
import PromptsManager from "~/components/features/PromptsManager.vue"
import UsersManager from "~/components/features/UsersManager.vue"

const currentPage = ref("")

const changePage = (page) => {
    currentPage.value = page
}

const buttons = ref([
    {
        label: 'Disconnect',
        dropdown: false,
    },
    {
        label: 'Simulate',
        dropdown: true,
        menuItems: ['Connect', 'Register', 'Disconnect', 'RemoveUser', 'CreateEntity', 'RemoveEntity', 'GetEntities', 'NewMessage', 'MakeDecision'],
    },
    {
        label: 'Entities Manager',
        dropdown: false,
    },
    {
        label: 'Users Manager',
        dropdown: false,
    },
    {
        label: 'Prompts Manager',
        dropdown: false,
    },
    {
        label: 'Settings',
        dropdown: false,
    },
])

</script>