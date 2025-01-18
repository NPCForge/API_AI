import { getAllUsers } from '../DataBase';

export default defineEventHandler(async () => {
    const user = await getAllUsers();
    console.log("here", user)
    return user;
});
