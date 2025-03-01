import { useFetch } from '#app';

const isUp = async () => {
    const { data, error } = await useFetch('http://localhost:8000/health');

    if (error.value) {
        console.log('L’API ne répond pas ❌', error.value);
        return false;
    }

    console.log('L’API est en ligne ✅');
    return true;
};

export { isUp };
