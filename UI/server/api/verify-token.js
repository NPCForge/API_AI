import { verifyToken } from '../TokenStore';

export default defineEventHandler(async (event) => {
    const token = getHeader(event, 'Authorization')?.replace('Bearer ', '');

    if (!token) {
        throw createError({ statusCode: 401, message: 'Token manquant.' });
    }

    const { valid, userId } = verifyToken(token);

    if (!valid) {
        throw createError({ statusCode: 401, message: 'Token invalide ou expiré.' });
    }

    return { success: true, userId };
});
