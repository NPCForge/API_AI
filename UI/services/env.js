// services/env.js

export const setEnvVariable = (store, value) => {
    localStorage.setItem(store, value);
}

export const getEnvVariable = (store) => {
    return localStorage.getItem(store);
}
