import { useFetch } from 'nuxt/app';

const isSettup = () => {
    const { data, error } = useFetch('/api/get-all-user');

    if (error.value) {
        console.error('Erreur lors de la récupération des utilisateurs :', error.value.message);
        return false;
    }

    return computed(() => data.value && data.value.length > 0);
}

export {
    isSettup
}