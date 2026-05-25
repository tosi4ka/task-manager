import { createNavigation } from 'next-intl/navigation'

export const { useRouter, usePathname } = createNavigation({
	locales: ['fr', 'en'],
})
