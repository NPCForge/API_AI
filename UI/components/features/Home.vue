<template>
    <div class="body">
        <h3>{{ pageName }}</h3>
        <div class="hb"></div>
        <br>
        <ul>
            <li>✅ Connect</li>
            <li>✅ Disconnect</li>
            <li>✅ Register</li>
            <li>🟠 RemoveUser (fonctionnel mais a fixer)</li>
            <li>❌ Status</li>
            <li>❌ CreateEntity</li>
            <li>❌ RemoveEntity</li>
            <li>❌ GetEntities</li>
            <li>❌ MakeDecision</li>
            <li>❌ NewMessage</li>
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
