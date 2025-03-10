import { ofetch } from 'ofetch';

const isUp = async () => {
    try {
        const data = await ofetch('/api/isUp', {
            method: 'POST',
        });

        console.log(data.status_code, "");
        if (data.status_code == 200) return true;
        else return false;
    } catch (error) {
        console.log('L’API ne répond pas ❌', error);
        return false;
    }
};

const getConnected = async () => {
    return [{ id: "mathieu" }]
}

export { isUp, getConnected };
