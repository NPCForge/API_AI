<template>
    <div class="monitoring-connected shadow-lg rounded">
        <div class="head shadow d-flex align-self-center justify-content-between">
            <p>Monitoring - Connected</p>
            <UButton
                icon="material-symbols:refresh"
                size="xs"
                color="primary"
                variant="solid"
                label="Refresh"
                :trailing="false"
                @click="request"
            />
        </div>
        <div class="body">
            <div class="fiche d-flex align-self-center justify-content-start">
                <p>Id:</p>
            </div>
        </div>
    </div>
</template>

<style scoped>
    .monitoring-connected {
        height: 300px;
        width: 500px;
        background-color: white;
        display: block;
    }

    .head {
        width: 100%;
        height: 10%;
        padding: 1%;
    }

    .body {
        width: 100%;
        height: 90%;
        overflow-y: auto;
        padding-top: 1%;
    }

    .fiche {
        background-color: blue;
        width: 100%;
        height: 10%;
        margin-bottom: 1%;
        padding-left: 1%;
        transition: opacity 0.3s;
    }

    .fiche:hover {
        opacity: 0.7;
        cursor: pointer;
    }
</style>

<script setup>
    const request = () => {
        console.log("Fetching population data...");
        // Implémenter la requête vers getPopulation ici
    };
</script>
