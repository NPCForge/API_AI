<template>
    <div class="container shadow">
        <div class="header">
            <UButton
                icon="material-symbols:cancel"
                size="2xs"
                color="black"
                variant="ghost"
                :trailing="false"
                @click="close"
            />
        </div>
        <div class="content">
            <p>{{ props.text }}</p>
        </div>
    </div>
</template>

<script setup>
    const emit = defineEmits(['close']);

    const props = defineProps({
        text: String,
        id: Number
    });

    const close = () => {
        emit("close", props.id);
    };
</script>

<style scoped>
    .container {
        height: 4rem;
        width: 100%;
        background-color: rgb(255, 255, 255);
        margin-bottom: 5px;
        padding: 0;
        overflow: hidden;
        border-radius: 2%;
    }
    .header {
        /* background-color: pink; */
        height: 40%;
        width: 100%;
        display: flex;
        flex-direction: row-reverse;
        padding: 1%;
    }
    .content {
        /* background-color: rgb(137, 62, 75); */
        height: 60%;
        width: 100%;
        text-align: center;
        overflow: hidden;
    }
</style>
