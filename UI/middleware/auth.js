export default defineNuxtRouteMiddleware(async (to, from) => {
    if (process.client) {
        // Récupérer le token depuis le localStorage
        const token = localStorage.getItem('authToken');

        // Vérifie si le token existe
        if (!token) {
            console.warn('Utilisateur non authentifié, redirection vers la page de connexion.');
            return navigateTo('/'); // Redirige vers la page de connexion
        }

        try {
            // Vérifie le token en appelant l'API
            const response = await $fetch('/api/verify-token', {
                method: 'GET',
                headers: {
                    Authorization: `Bearer ${token}`, // Envoi du token dans les en-têtes
                },
            });

            if (!response.success) {
                console.warn('Token invalide, redirection vers la page de connexion.');
                localStorage.removeItem('authToken'); // Supprime le token
                return navigateTo('/'); // Redirige vers la page de connexion
            }

            // Si le token est valide, l'utilisateur est autorisé à continuer
            console.log('Token valide, accès autorisé.');
        } catch (error) {
            console.error('Erreur lors de la vérification du token :', error);
            localStorage.removeItem('authToken'); // Supprime le token
            return navigateTo('/'); // Redirige vers la page de connexion
        }
    }
});
