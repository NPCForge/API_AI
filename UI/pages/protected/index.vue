<template>
    <div class="body_index">
        <p>- Afficher les utilisateur enregistrer</p>
        <p>- Afficher le model utilisé</p>
        <monitoringConnected class="component"/>
        <monitoringPrompt class="component"/>

        <div class="alert_box">
            <div v-for="(n, i) in arr" :key="i" data-aos="fade-up">
                <alert :text="n.text" :id="i" @close="removeAlert"/>
            </div>
        </div>
    </div>
</template>

<style scoped>
    .body_index {
        background-color: rgb(168, 168, 168);
        height: auto;
        min-height: 100vh;
    }
    .component {
        margin: 1%;
    }
    .alert_box {
        position: fixed;
        z-index: 8;
        right: 3%;
        top: 0;
        width: 18%;
        height: 100%;
        display: flex;
        flex-direction: column;
        padding-top: 1%;
    }
</style>

<script setup>
    import { onMounted, ref } from 'vue';
    import AOS from 'aos';
    import 'aos/dist/aos.css'; // 💡 Assure-toi d'importer les styles
    import monitoringConnected from '~/components/utils/monitoringConnected.vue';
    import monitoringPrompt from '~/components/utils/monitoringPrompt.vue';
    import { isUp } from '~/services/ApiServices';
    import alert from '~/components/utils/alert.vue';

    const isConnected = ref(false);
    let intervalId = null;

    // ✅ Rendre le tableau réactif
    const arr = ref([
        { text: "Bienvenue" },
    ]);

    const TestConnexion = async () => {
        const res = await isUp();
        if (isConnected.value == res)
            return
        isConnected.value = res
        if (isConnected.value == true)
            arr.value.push({ text: "L'API est allumée ✅." })
        else
            arr.value.push({ text: "L'API est éteinte ❌." })
    };

    // ✅ Supprimer un élément et rafraîchir AOS
    const removeAlert = (id) => {
        arr.value.splice(id, 1);
        console.log(arr.value);
        AOS.refresh(); // 💡 Rafraîchir les animations après suppression
    };

    definePageMeta({
        middleware: 'auth' // Active le middleware `/protected/_middleware.ts`
    });

    onMounted(() => {
        TestConnexion();
        intervalId = setInterval(TestConnexion, 10000);
        AOS.init({
            duration: 1000,
            once: true, // Animation se joue une seule fois
        });
    });

    onBeforeUnmount(() => {
        if (intervalId) {
            clearInterval(intervalId);
        }
    });
</script>
