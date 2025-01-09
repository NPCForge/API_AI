// https://nuxt.com/docs/api/configuration/nuxt-config
import * as dotenv from 'dotenv'
import { resolve } from 'path'

dotenv.config({ path: resolve(__dirname, '../.env.local') })

export default defineNuxtConfig({
    compatibilityDate: '2024-11-01',
    devtools: { enabled: true },
    css: [
        '~/assets/css/custom.css'
    ],
    plugins: [
        {
            src: '~/plugins/bootstrap.client.ts',
            mode: 'client'
        }
    ],
    runtimeConfig: {
        // Charger automatiquement toutes les variables NUXT_ comme publiques
        public: Object.fromEntries(
            Object.entries(process.env)
        ),
    }
})
