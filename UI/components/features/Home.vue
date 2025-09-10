<template>
    <div class="w-100 h-100 p-4">
        <!-- Header -->
        <h1 class="h2 fw-semibold text-dark mb-2">{{ pageName }}</h1>
        <p class="text-muted mb-3">Tableau de bord des fonctionnalités API</p>
        <hr class="mb-4">

        <!-- Status Cards -->
        <div class="row g-3 mb-4">
            <div class="col-12 col-md-4">
                <div class="card border">
                    <div class="card-body">
                        <div class="d-flex align-items-center">
                            <span class="h2 fw-bold text-dark me-2 mb-0">{{ completedCount }}</span>
                            <span class="text-muted">Fonctionnalités complètes</span>
                        </div>
                    </div>
                </div>
            </div>

            <div class="col-12 col-md-4">
                <div class="card border">
                    <div class="card-body">
                        <div class="d-flex align-items-center">
                            <span class="h2 fw-bold text-dark me-2 mb-0">{{ pendingCount }}</span>
                            <span class="text-muted">En développement</span>
                        </div>
                    </div>
                </div>
            </div>

            <div class="col-12 col-md-4">
                <div class="card border">
                    <div class="card-body">
                        <div class="d-flex align-items-center">
                            <span class="h2 fw-bold text-dark me-2 mb-0">{{ progressPercentage }}%</span>
                            <span class="text-muted">Progression</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <!-- Features List -->
        <div class="card border">
            <div class="card-header bg-white border-bottom">
                <h5 class="fw-medium text-dark mb-0">État des fonctionnalités</h5>
            </div>
            <div class="card-body">
                <ul class="list-unstyled mb-0">
                    <li v-for="(feature, index) in features" :key="feature.name"
                        class="d-flex align-items-center py-2"
                        :class="{ 'border-bottom': index < features.length - 1 }">
                        <span class="me-3" :class="{ 'text-warning': feature.status === 'pending' }">
                            {{ feature.status === 'completed' ? '✅' : '❌' }}
                        </span>
                        <span class="text-dark">{{ feature.name }}</span>
                        <span v-if="feature.note" class="ms-2 small text-muted">({{ feature.note }})</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</template>

<script setup>
import { computed, ref } from 'vue'

defineProps({
    pageName: {
        type: String,
        default: 'Home'
    }
})

// Liste des fonctionnalités avec leur statut
const features = ref([
    { name: 'Connect', status: 'completed' },
    { name: 'Disconnect', status: 'completed' },
    { name: 'Register', status: 'completed' },
    { name: 'RemoveUser', status: 'completed', note: 'fonctionnel mais a fixer' },
    { name: 'Status', status: 'completed' },
    { name: 'CreateEntity', status: 'completed' },
    { name: 'RemoveEntity', status: 'completed' },
    { name: 'GetEntities', status: 'completed' },
    { name: 'MakeDecision', status: 'pending' },
    { name: 'NewMessage', status: 'pending' }
])

// Calculs des statistiques
const completedCount = computed(() =>
    features.value.filter(f => f.status === 'completed').length
)

const pendingCount = computed(() =>
    features.value.filter(f => f.status === 'pending').length
)

const totalCount = computed(() => features.value.length)

const progressPercentage = computed(() =>
    Math.round((completedCount.value / totalCount.value) * 100)
)
</script>