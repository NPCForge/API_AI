// middleware/auth.js

import { status } from '~/services/npcforge.js';

export default defineNuxtRouteMiddleware(async (to, from) => {
    if (process.client) {
        const token = localStorage.getItem('token'); // Utiliser localStorage uniquement côté client

        if (!token) {
            return navigateTo('/'); // Rediriger vers la page de login si aucun token n'est trouvé
        }

        const isAuthenticated = await status(); // Appeler la fonction status

        if (isAuthenticated.Status === "Failed") {
            return navigateTo('/'); // Si la connexion est invalide, rediriger vers la page de login
        }
    }
});
