export default defineNuxtRouteMiddleware(async (to, from) => {
    const token = localStorage.getItem('authToken');

    if (!token) {
        console.warn('Accès refusé : utilisateur non authentifié.');
        return navigateTo('/'); // Redirige vers la page de connexion
    }

    try {
        const response = await $fetch('/api/verify-token', {
            method: 'GET',
            headers: {
                Authorization: `Bearer ${token}`,
            },
        });

        if (!response.success) {
            localStorage.removeItem('authToken');
            return navigateTo('/');
        }
    } catch (error) {
        console.error('Erreur lors de la vérification du token :', error);
        localStorage.removeItem('authToken');
        return navigateTo('/');
    }
});
