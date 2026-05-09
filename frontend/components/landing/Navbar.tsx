'use client'

import { useTranslations } from 'next-intl'
import Link from 'next/link'
import CheckIcon from '../icons/CheckIcon'

export default function Navbar() {
	const t = useTranslations('nav')

	return (
		<nav className='fixed top-0 left-0 right-0 z-50 bg-bg border-b border-border_b shadow-sm'>
			<div className=' mx-auto px-12 flex items-center justify-between h-[68px]'>
				<div className='flex gap-[10px] items-center font-syne font-extrabold text-[1.1rem] tracking-[-0.5px]'>
					<div className='w-[30px] h-[30px] bg-accent rounded-[7px] grid place-items-center text-black'>
						<CheckIcon className='w-4 h-4' />
					</div>
					<span className=' text-text'>Task</span>
					<span className=' text-accent'>Manager</span>
				</div>
				<div className='hidden md:flex gap-8 text-sm text-muted font-syne font-semibold text-[0.82rem] uppercase tracking-[0.04em]'>
					<Link href='#'>{t('api')}</Link>
					<a
						href={`${process.env.NEXT_PUBLIC_API_URL}/swagger/index.html`}
						target='_blank'
					>
						{t('swagger')}
					</a>
					<Link href='#'>{t('architecture')}</Link>
				</div>
				<div className='flex items-center gap-3'>
					<button className='bg-transparent border border-border_b text-text py-[0.45rem] px-[1.2rem] rounded-[6px] font-syne text-[0.82rem] font-semibold cursor-pointer transition-colors duration-200 hover:border-accent hover:text-accent w-[180px] '>
						{t('login')}
					</button>
					<button className='bg-accent border border-border_b text-bg py-[0.45rem] px-[1.2rem] rounded-[6px] font-syne text-[0.82rem] font-semibold cursor-pointer transition-colors duration-200 hover:border-accent hover:bg-transparent hover:text-accent w-[180px]'>
						{t('register')}
					</button>
				</div>
			</div>
		</nav>
	)
}
