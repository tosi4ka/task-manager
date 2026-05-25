import { useTranslations } from 'next-intl'

export default function Hero() {
	const t = useTranslations('hero')

	return (
		<section className='relative min-h-screen flex flex-col justify-center px-12 pt-40 pb-24 overflow-hidden'>
			<div
				className='absolute inset-0 pointer-events-none'
				style={{
					backgroundImage: `
            linear-gradient(rgba(255,255,255,0.013) 1px, transparent 1px),
            linear-gradient(90deg, rgba(255,255,255,0.013) 1px, transparent 1px)
        `,
					backgroundSize: '60px 60px',
				}}
			/>
			<div className='absolute rounded-full blur-[120px] opacity-[0.16] pointer-events-none w-[600px] h-[600px] bg-accent -top-24 -right-24' />
			<div className='absolute rounded-full blur-[120px] opacity-[0.16] pointer-events-none w-[400px] h-[400px] bg-[#4060FF] bottom-0 -left-20' />
			<div className='absolute rounded-full blur-[120px] opacity-[0.16] pointer-events-none w-[250px] h-[250px] bg-accent2 top-1/2 right-1/4' />

			<div className='inline-flex items-center gap-2 bg-accent/8 border border-accent/20 rounded-full px-3.5 py-1.5 mb-8 w-fit'>
				<div className='w-1.5 h-1.5 rounded-full bg-accent animate-pulse' />
				<span className='font-mono text-xs text-accent tracking-wide'>
					{t('badge')}
				</span>
			</div>

			<h1 className='text-[clamp(2.8rem,6.5vw,6rem)] font-extrabold leading-none tracking-[-0.04em] max-w-[760px] font-syne'>
				{t('title1')}
				<br />
				<span className='text-accent'>{t('title2')}</span>
				<br />
				<span className='text-muted'>{t('title3')}</span>
			</h1>

			<p className='mt-[1.6rem] max-w-[500px] font-mono text-[0.88rem] text-muted font-light leading-[1.75] whitespace-pre-line'>
				{t('sub')}
			</p>

			<div className='flex items-center gap-4 mt-10'>
				<button className='bg-accent text-bg font-syne font-bold text-[0.95rem] px-8 py-3 rounded-lg transition-all hover:opacity-[0.88] hover:bg-transparent hover:text-accent hover:border-accent border'>
					{t('start')}
				</button>
				<a
					href={`${process.env.NEXT_PUBLIC_API_URL}/swagger/index.html`}
					target='_blank'
					className='font-mono text-[0.78rem] text-gold bg-gold/7 border border-gold/20 px-[1.4rem] py-3 rounded-lg transition-colors hover:bg-gold/12'
				>
					{t('swagger')}
				</a>
			</div>

			<div className='hidden xl:block absolute right-12 top-1/2 -translate-y-1/2 w-[460px] bg-bg2 border border-border_b rounded-[14px] overflow-hidden shadow-[0_40px_120px_rgba(0,0,0,0.6)]'>
				<div className='bg-bg3 px-4 py-3 flex items-center gap-2 border-b border-border_b'>
					<div className='w-[11px] h-[11px] rounded-full bg-[#FF5F57]' />
					<div className='w-[11px] h-[11px] rounded-full bg-[#FFBD2E]' />
					<div className='w-[11px] h-[11px] rounded-full bg-[#28CA41]' />
					<span className='ml-auto font-mono text-[0.7rem] text-muted'>
						main.go — task-manager/backend
					</span>
				</div>

				<div className='p-5 font-mono text-[0.77rem] leading-[1.95]'>
					<div>
						<span className='text-[#BB86FC]'>package</span> main
					</div>
					<div>&nbsp;</div>
					<div>
						<span className='text-[#BB86FC]'>import</span> (
					</div>
					<div>
						&nbsp;&nbsp;
						<span className='text-gold'>"github.com/gin-gonic/gin"</span>
					</div>
					<div>
						&nbsp;&nbsp;
						<span className='text-gold'>"task-manager/internal/server"</span>
					</div>
					<div>
						&nbsp;&nbsp;
						<span className='text-gold'>"task-manager/internal/config"</span>
					</div>
					<div>)</div>
					<div>&nbsp;</div>
					<div>
						<span className='text-[#BB86FC]'>func</span>{' '}
						<span className='text-accent'>main</span>() {'{'}
					</div>
					<div>
						&nbsp;&nbsp;cfg := config.<span className='text-accent'>Load</span>
						()
					</div>
					<div>
						&nbsp;&nbsp;r := gin.<span className='text-accent'>New</span>()
					</div>
					<div>&nbsp;</div>
					<div>
						&nbsp;&nbsp;<span className='text-[#4A5568]'>// Auth routes</span>
					</div>
					<div>
						&nbsp;&nbsp;r.<span className='text-accent'>POST</span>(
						<span className='text-gold'>"/auth/register"</span>, ...)
					</div>
					<div>
						&nbsp;&nbsp;r.<span className='text-accent'>POST</span>(
						<span className='text-gold'>"/auth/login"</span>, ...)
					</div>
					<div>&nbsp;</div>
					<div>
						&nbsp;&nbsp;
						<span className='text-[#4A5568]'>// Tasks — JWT required</span>
					</div>
					<div>
						&nbsp;&nbsp;tasks := r.<span className='text-accent'>Group</span>(
						<span className='text-gold'>"/task"</span>)
					</div>
					<div>
						&nbsp;&nbsp;tasks.<span className='text-accent'>Use</span>
						(middleware.<span className='text-accent'>JWT</span>())
					</div>
					<div>
						&nbsp;&nbsp;tasks.<span className='text-accent'>POST</span>(
						<span className='text-gold'>"/createTask"</span>, ...)
					</div>
					<div>{'}'}</div>
				</div>
			</div>
		</section>
	)
}
