<template>
    <div class="frontground d-flex flex-column justify-content-center align-items-center">
        <div class="loader"></div>
        <p v-if="Text" style="margin-top: 1%; color: white;">{{ Text }}</p>
    </div>
</template>

<script setup>
    defineProps({
        Text: {
            type: String
        }
    })
</script>

<style scoped>
    .frontground {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background-color: rgba(0, 0, 0, 0.53);
        transition: all 0.5s;
    }
    .loader {
        width: 60px;
        aspect-ratio: 2;
        --_g: no-repeat radial-gradient(circle closest-side,#ffffff 90%,#ffffff00);
        background:
            var(--_g) 0%   50%,
            var(--_g) 50%  50%,
            var(--_g) 100% 50%;
        background-size: calc(100%/3) 50%;
        animation: l3 1s infinite linear;
    }
    @keyframes l3 {
        20%{background-position:0%   0%, 50%  50%,100%  50%}
        40%{background-position:0% 100%, 50%   0%,100%  50%}
        60%{background-position:0%  50%, 50% 100%,100%   0%}
        80%{background-position:0%  50%, 50%  50%,100% 100%}
    }
</style>