export default function Footer() {
	return (
		<footer className='px-12 py-7 border-t border-border_b flex items-center justify-between'>
			<p className='font-mono text-[0.72rem] text-muted'>
				<span className='text-accent'>TaskManager</span> · Full-stack project ·
				Go + Next.js
			</p>
			<div className='flex gap-6'>
				<a
					href='https://github.com/tosi4ka/task-manager'
					target='_blank'
					className='font-mono text-[0.72rem] text-muted transition-colors hover:text-accent'
				>
					GitHub
				</a>
				<a
					href={`${process.env.NEXT_PUBLIC_API_URL}/swagger/index.html`}
					target='_blank'
					className='font-mono text-[0.72rem] text-muted transition-colors hover:text-accent'
				>
					Swagger
				</a>
			</div>
		</footer>
	)
}
