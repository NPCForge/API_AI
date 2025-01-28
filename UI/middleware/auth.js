export default defineNuxtRouteMiddleware(async (to, from) => {
    const cookies = useNuxtApp().$cookies;
    const token = cookies.get('authToken'); // Récupérer le token depuis les cookies

    if (!token) {
        console.warn('Utilisateur non authentifié.');
        return navigateTo('/');
    }

    try {
        const response = await $fetch('/api/verify-token', {
            method: 'GET',
            headers: {
                Authorization: `Bearer ${token}`,
            },
        });

        if (!response.success) {
            cookies.remove('authToken'); // Supprimer le token en cas d'échec
            return navigateTo('/');
        }
    } catch (error) {
        console.error('Erreur lors de la vérification du token :', error);
        cookies.remove('authToken'); // Supprimer le token en cas d'erreur
        return navigateTo('/');
    }
});
