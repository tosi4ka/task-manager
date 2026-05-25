import { backend, frontend } from '@/lib/data/stack'

export default function StackStrip() {
	return (
		<div className='px-12 py-10 border-t border-border_b flex items-start gap-12'>
			<div className='flex flex-col gap-3'>
				<span className='font-mono text-[0.68rem] text-muted uppercase tracking-[0.1em]'>
					Backend
				</span>
				<div className='flex gap-2 flex-wrap'>
					{backend.map(({ name, highlight }) => (
						<span
							key={name}
							className={`font-mono text-[0.73rem] px-3 py-1.25 rounded-md border bg-bg3 transition-colors hover:border-muted ${
								highlight
									? 'border-accent/30 text-accent'
									: 'border-border_b text-text'
							}`}
						>
							{name}
						</span>
					))}
				</div>
			</div>

			<div className='w-px bg-border_b self-stretch' />

			<div className='flex flex-col gap-3'>
				<span className='font-mono text-[0.68rem] text-muted uppercase tracking-[0.1em]'>
					Frontend <span className='text-gold normal-case'>· in progress</span>
				</span>
				<div className='flex gap-2 flex-wrap'>
					{frontend.map(({ name, highlight }) => (
						<span
							key={name}
							className={`font-mono text-[0.73rem] px-3 py-1.25 rounded-md border bg-bg3 transition-colors hover:border-muted ${
								highlight
									? 'border-[#60A5FA]/30 text-[#60A5FA]'
									: 'border-border_b text-text'
							}`}
						>
							{name}
						</span>
					))}
				</div>
			</div>
		</div>
	)
}
