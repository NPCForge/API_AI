import { revokeToken } from '../TokenStore';

export default defineEventHandler(async (event) => {
    const token = getHeader(event, 'Authorization')?.replace('Bearer ', '');

    if (!token) {
        throw createError({ statusCode: 400, message: 'Token manquant.' });
    }

    revokeToken(token);

    return { success: true, message: 'Déconnexion réussie.' };
});
