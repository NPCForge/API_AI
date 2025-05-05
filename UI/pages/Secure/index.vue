<template>
    <div class="d-flex justify-content-start align-items-start vh-100">
        <div class="d-flex flex-column justify-content-start align-items-start vw-15 vh-100 leftBar">
            <ButtonGroup :buttons="buttons" @change="changePage"/>
        </div>
        <div class="d-flex justify-content-start align-items-start vw-100 vh-100" style="padding: 0.5%;">
            <Home
                v-if="currentPage === 'Home'"
                :pageName="'Home'"
            />
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
            <Simulate
                v-if="currentPage.split('_')[0] === 'Simulate'"
                :pageName="'Simulate'"
                :current="currentPage.split('_')[1]"
            />
        </div>
    </div>
</template>

<script setup>
    import { useRouter } from 'vue-router'
    import { ref } from 'vue'
    import ButtonGroup from '~/components/ButtonGroup.vue'

    import Home from "~/components/features/Home.vue"
    import Settings from "~/components/features/Settings.vue"
    import EntitiesManager from "~/components/features/EntitiesManager.vue"
    import PromptsManager from "~/components/features/PromptsManager.vue"
    import UsersManager from "~/components/features/UsersManager.vue"
    import Simulate from "~/components/features/Simulate.vue"
    import { disconnect } from '~/services/npcforge'

    const currentPage = ref("Home")
    const router = useRouter()

    const changePage = (page) => {
        if (page === "Disconnect") {
            disconnect()
            router.push('/')
            return
        }
        currentPage.value = page
    }

    const buttons = ref([
        {
            label: 'Home',
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
        {
            label: 'Disconnect',
            dropdown: false,
        },
    ])

    definePageMeta({
        middleware: 'auth',
    });
</script>

<style scoped>
    .leftBar {
        background-color: rgb(242, 242, 242);
        border-right: 3px solid rgb(207, 207, 207);
    }
</style>