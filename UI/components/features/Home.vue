<template>
    <div class="body">
        <h3>{{ pageName }}</h3>
        <div class="hb"></div>
        <br>
        <ul>
            <li class="">✅ Connect</li>
            <li class="">✅ Disconnect</li>
            <li class="">✅ Register</li>
            <li class="">✅ RemoveUser (fonctionnel mais a fixer)</li>
            <li class="">✅ Status</li>
            <li class="">✅ CreateEntity</li>
            <li class="orange">❌ RemoveEntity</li>
            <li class="">❌ GetEntities</li>
            <li class="">❌ MakeDecision</li>
            <li class="">❌ NewMessage</li>
        </ul>
    </div>
</template>

<script setup>
    defineProps({
        pageName: String
    })
</script>

<style scoped>
    .body {
        height: 100%;
        width: 100%;
        overflow-y: scroll;
        overflow-x: hidden;
    }
</style>
