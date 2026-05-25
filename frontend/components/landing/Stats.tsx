import { useTranslations } from 'next-intl'

export default function Stats() {
	const t = useTranslations('stats')

	return (
		<>
			<div className='bg-bg2 border-y border-border_b px-12 py-8 flex items-center gap-0'>
				<div className='flex items-center'>
					<div className='px-[1.4rem] py-[0.6rem] border border-accent/40 bg-accent/5 rounded-lg font-mono text-sm text-accent'>
						Handler
					</div>
					<span className='font-mono text-muted px-3'>→</span>
					<div className='px-[1.4rem] py-[0.6rem] border border-border_b bg-bg3 rounded-lg font-mono text-sm text-text'>
						Service
					</div>
					<span className='font-mono text-muted px-3'>→</span>
					<div className='px-[1.4rem] py-[0.6rem] border border-border_b bg-bg3 rounded-lg font-mono text-sm text-text'>
						Repository
					</div>
					<span className='font-mono text-muted px-3'>→</span>
					<div className='px-[1.4rem] py-[0.6rem] border border-border_b bg-bg3 rounded-lg font-mono text-sm text-text'>
						Database
					</div>
				</div>
				<span className='ml-auto font-mono text-[0.7rem] text-muted uppercase tracking-[0.08em]'>
					Clean Architecture · Handler → Service → Repository → DB
				</span>
			</div>

			<div className='flex border-y border-border_b bg-bg2'>
				<div className='flex-1 px-8 py-[1.6rem] border-r border-border_b'>
					<div className='text-[2rem] font-extrabold tracking-[-0.04em]'>
						<span className='text-accent'>60</span>/min
					</div>
					<div className='mt-1 font-mono text-[0.72rem] text-muted uppercase tracking-[0.06em]'>
						{t('rateLimit')}
					</div>
				</div>

				<div className='flex-1 px-8 py-[1.6rem] border-r border-border_b'>
					<div className='text-[2rem] font-extrabold tracking-[-0.04em]'>
						<span className='text-accent'>8</span>
					</div>
					<div className='mt-1 font-mono text-[0.72rem] text-muted uppercase tracking-[0.06em]'>
						{t('endpoints')}
					</div>
				</div>

				<div className='flex-1 px-8 py-[1.6rem] border-r border-border_b'>
					<div className='text-[2rem] font-extrabold tracking-[-0.04em]'>
						JWT<span className='text-accent'>+</span>RT
					</div>
					<div className='mt-1 font-mono text-[0.72rem] text-muted uppercase tracking-[0.06em]'>
						{t('auth')}
					</div>
				</div>

				<div className='flex-1 px-8 py-[1.6rem]'>
					<div className='text-[2rem] font-extrabold tracking-[-0.04em]'>
						<span className='text-accent'>:8081</span>
					</div>
					<div className='mt-1 font-mono text-[0.72rem] text-muted uppercase tracking-[0.06em]'>
						{t('port')}
					</div>
				</div>
			</div>
		</>
	)
}
