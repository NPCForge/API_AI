<template>
    <div style="height: 100%; width: 100%;">
        <Connect v-if="current === 'Connect'" />
        <CreateEntity v-if="current === 'CreateEntity'" />
        <Disconnect v-if="current === 'Disconnect'" />
        <GetEntities v-if="current === 'GetEntities'" />
        <MakeDecision v-if="current === 'MakeDecision'" />
        <NewMessage v-if="current === 'NewMessage'" />
        <Register v-if="current === 'Register'" />
        <RemoveEntity v-if="current === 'RemoveEntity'" />
        <RemoveUser v-if="current === 'RemoveUser'" />
        <Status v-if="current === 'Status'" />
    </div>
</template>

<script setup>
    import Connect from './simulate/Connect.vue';
    import CreateEntity from './simulate/CreateEntity.vue';
    import Disconnect from './simulate/Disconnect.vue';
    import GetEntities from './simulate/GetEntities.vue';
    import MakeDecision from './simulate/MakeDecision.vue';
    import NewMessage from './simulate/NewMessage.vue';
    import Register from './simulate/Register.vue';
    import RemoveEntity from './simulate/RemoveEntity.vue';
    import RemoveUser from './simulate/RemoveUser.vue';
    import Status from './simulate/Status.vue';

    defineProps({
        current: {
            type: String,
            required: true
        }
    })
</script>