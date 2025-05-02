// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	compatibilityDate: '2024-11-01',
	devtools: { enabled: true },
	// modules: ['@nuxtjs/bootstrap-vue'],
	css: [
		'bootstrap/dist/css/bootstrap.min.css',
		'assets/css/global.css'
	],
	server: {
		port: 4000,  // Le port sur lequel Nuxt.js sera lancé
		host: '0.0.0.0',  // Pour permettre l'accès à partir de n'importe quelle interface
	},
})
