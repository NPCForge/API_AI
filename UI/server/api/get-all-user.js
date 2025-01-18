import { getAllUsers } from '../DataBase';

export default defineEventHandler(async () => {
    const user = await getAllUsers();
    return user;
});
