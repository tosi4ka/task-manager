import createMiddleware from 'next-intl/middleware'

const intlMiddleware = createMiddleware({
	locales: ['fr', 'en'],
	defaultLocale: 'fr',
})

export default intlMiddleware

export const config = {
	matcher: ['/((?!api|_next|_vercel|.*\\..*).*)'],
}
