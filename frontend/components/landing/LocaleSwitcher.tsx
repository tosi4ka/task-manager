'use client'

import { usePathname, useRouter } from '@/i18n/navigation'
import { useLocale } from 'next-intl'

interface LocaleSwitcherProps {
	className?: string
}

export default function LocaleSwitcher({ className }: LocaleSwitcherProps) {
	const locale = useLocale()
	const router = useRouter()
	const pathname = usePathname()

	const toggle = () => {
		const nextLocale = locale === 'fr' ? 'en' : 'fr'
		router.replace(pathname, { locale: nextLocale })
	}

	const activeCls = 'text-text'
	const inactiveCls = 'text-muted cursor-pointer hover:text-text'

	return (
		<button onClick={toggle} className={className}>
			<span className={locale === 'fr' ? activeCls : inactiveCls}>FR</span>

			<span className={locale === 'en' ? activeCls : inactiveCls}>EN</span>
		</button>
	)
}
