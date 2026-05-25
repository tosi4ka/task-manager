import { features } from '@/lib/data/features'

export default function Features() {
	return (
		<section className='px-12 pb-20'>
			<p className='font-mono text-[0.72rem] text-accent uppercase tracking-[0.12em] mb-4'>
				// Key features
			</p>
			<h2 className='text-[clamp(1.8rem,3vw,2.6rem)] font-extrabold tracking-[-0.03em] mb-10'>
				What makes it <em className='text-accent not-italic'>stand out</em>
			</h2>

			<div className='grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-px bg-border_b border border-border_b rounded-[14px] overflow-hidden'>
				{features.map((f, i) => (
					<div
						key={i}
						className='group bg-bg2 p-[2.2rem] relative overflow-hidden transition-colors hover:bg-bg3 flex flex-col'
					>
						<div className='absolute top-0 left-0 right-0 h-px bg-gradient-to-r from-transparent via-accent to-transparent opacity-0 group-hover:opacity-100 transition-opacity' />

						<div className='w-[42px] h-[42px] rounded-[10px] bg-accent/7 border border-accent/15 grid place-items-center mb-5 text-[1.1rem]'>
							{f.icon}
						</div>
						<div className='font-syne font-bold text-[0.95rem] mb-2 tracking-[-0.02em]'>
							{f.name}
						</div>
						<div className='font-mono text-[0.75rem] text-muted leading-[1.7] font-light flex-1'>
							{f.desc}
						</div>
						<div className='w-fit mt-4 font-mono text-[0.67rem] text-accent bg-accent/7 border border-accent/15 px-2.5 py-0.75 rounded-full'>
							{f.tag}
						</div>
					</div>
				))}
			</div>
		</section>
	)
}
