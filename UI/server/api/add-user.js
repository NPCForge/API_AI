import { addUser } from '../DataBase';

export default defineEventHandler(async (event) => {
    const body = await readBody(event);
    const { name, password, access } = body;

    if (!name || !password) {
        throw createError({ statusCode: 400, message: 'Name and password are required' });
    }

    const user = await addUser(name, password, access || 0);
    return user;
});
