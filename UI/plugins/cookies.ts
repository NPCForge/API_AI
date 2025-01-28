export default defineNuxtPlugin(() => {
    return {
        provide: {
            cookies: {
                // Récupérer un cookie par son nom
                get(name: string) {
                    if (process.server) {
                        const cookies = useRequestHeaders(['cookie']);
                        const cookie = cookies?.cookie?.split(';').find(c => c.trim().startsWith(`${name}=`));
                        return cookie ? decodeURIComponent(cookie.split('=')[1]) : null;
                    } else if (process.client) {
                        const cookie = document.cookie.split(';').find(c => c.trim().startsWith(`${name}=`));
                        return cookie ? decodeURIComponent(cookie.split('=')[1]) : null;
                    }
                    return null;
                },

                // Définir un cookie
                set(name: string, value: string, options: { path?: string; maxAge?: number } = {}) {
                    if (process.client) {
                        let cookie = `${encodeURIComponent(name)}=${encodeURIComponent(value)}`;
                        if (options.path) cookie += `; path=${options.path}`;
                        if (options.maxAge) cookie += `; max-age=${options.maxAge}`;
                        document.cookie = cookie;
                    }
                },

                // Supprimer un cookie
                remove(name: string, options: { path?: string } = {}) {
                    this.set(name, '', { path: options.path || '/', maxAge: 0 });
                }
            }
        }
    };
});
