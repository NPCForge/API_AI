// middleware/auth.js

export default defineNuxtRouteMiddleware((to, from) => {
    // Vérifier si nous sommes en environnement client
    if (process.client) {
        const token = localStorage.getItem('token')  // Utiliser localStorage uniquement côté client

        if (!token) {
            return navigateTo('/')  // Rediriger vers la page de login si aucun token n'est trouvé
        }
    }
})
