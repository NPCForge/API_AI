const tokenStore = new Map();

// Ajouter un token
export const setToken = (token, userId) => {
    const expiresAt = Date.now() + 3600 * 1000; // Expire dans 1 heure
    tokenStore.set(token, { userId, expiresAt });
};

// Vérifier un token
export const verifyToken = (token) => {
    const data = tokenStore.get(token);
    if (!data) return { valid: false };

    // Vérifie l'expiration
    if (data.expiresAt < Date.now()) {
        tokenStore.delete(token); // Supprime les tokens expirés
        return { valid: false };
    }

    return { valid: true, userId: data.userId };
};

// Révoquer un token
export const revokeToken = (token) => {
    tokenStore.delete(token);
};
