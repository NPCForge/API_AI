// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	compatibilityDate: '2024-11-01',
	devtools: { enabled: true },
	// modules: ['@nuxtjs/bootstrap-vue'],
	css: [
		'bootstrap/dist/css/bootstrap.min.css',
		'assets/css/global.css'
	],
})
