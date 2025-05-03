<template>
    <div class="file-container">
        <div class="header">
            <Icon class="icon" name="material-symbols:add" size="2vh" style="color: black" />
        </div>
        <div class="file-list" v-for="(n, i) in prompts" :key="i" @click="changeCurrentPrompt(i)">
            {{ n.fileName }}
        </div>
    </div>
</template>


<script setup>
    defineProps({
        prompts: Array
    });

    const emits = defineEmits(['changePrompt'])

    const changeCurrentPrompt = (i) => {
        emits('changePrompt', i)
    };
</script>

<style scoped>
.file-container {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: center;
    height: 100%;
    width: 17%;
    padding: 1%;
}

.header {
    display: flex;
    justify-content: flex-end;
    align-items: center;
    width: 100%;
    height: 5%;
}

.file-list {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    height: 30px;
    /* Ajuste la hauteur de la ligne si nécessaire */
    background-color: rgb(234, 234, 234);
    border: 1px solid black;
    transition: all 0.2s;
    cursor: pointer;
}

.file-list:hover {
    background-color: rgb(198, 198, 198);
}

.icon {
    cursor: pointer;
}
</style>
