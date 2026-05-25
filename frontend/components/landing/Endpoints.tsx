export default function Endpoints() {
	const endpoints = [
		{ method: 'GET', path: '/health', desc: 'Health check', auth: false },
		{
			method: 'POST',
			path: '/auth/register',
			desc: 'Register new user',
			auth: false,
		},
		{
			method: 'POST',
			path: '/auth/login',
			desc: 'Login · returns JWT + RT',
			auth: false,
		},
		{
			method: 'POST',
			path: '/task/createTask',
			desc: 'Create a new task',
			auth: true,
		},
		{
			method: 'PATCH',
			path: '/task/updateTask/:id',
			desc: 'Update existing task',
			auth: true,
		},
		{
			method: 'GET',
			path: '/task/tasksList/:id',
			desc: 'List tasks by user',
			auth: true,
		},
		{
			method: 'GET',
			path: '/task/getById/:id',
			desc: 'Get task by ID',
			auth: true,
		},
		{
			method: 'DELETE',
			path: '/task/deleteTask/:id',
			desc: 'Delete a task',
			auth: true,
		},
	]

	const methodClass: Record<string, string> = {
		GET: 'bg-accent/12 text-accent',
		POST: 'bg-[#60A5FA]/12 text-[#60A5FA]',
		PATCH: 'bg-gold/12 text-gold',
		DELETE: 'bg-accent2/12 text-accent2',
	}

	return (
		<section className='px-12 py-20'>
			<p className='font-mono text-[0.72rem] text-accent uppercase tracking-[0.12em] mb-4'>
				// API Endpoints
			</p>
			<h2 className='text-[clamp(1.8rem,3vw,2.6rem)] font-extrabold tracking-[-0.03em] mb-10 font-syne'>
				Available <em className='text-accent not-italic'>routes</em>
			</h2>

			<div className='border border-border_b rounded-[14px] overflow-hidden bg-bg2'>
				<table className='w-full border-collapse'>
					<thead>
						<tr className='bg-bg3 border-b border-border_b'>
							<th className='px-[1.4rem] py-[0.9rem] font-mono text-[0.7rem] text-muted uppercase tracking-[0.1em] text-left'>
								Method
							</th>
							<th className='px-[1.4rem] py-[0.9rem] font-mono text-[0.7rem] text-muted uppercase tracking-[0.1em] text-left'>
								Endpoint
							</th>
							<th className='px-[1.4rem] py-[0.9rem] font-mono text-[0.7rem] text-muted uppercase tracking-[0.1em] text-left'>
								Description
							</th>
							<th className='px-[1.4rem] py-[0.9rem] font-mono text-[0.7rem] text-muted uppercase tracking-[0.1em] text-left'>
								Auth
							</th>
						</tr>
					</thead>
					<tbody>
						{endpoints.map((ep, i) => (
							<tr
								key={i}
								className='border-b border-border_b last:border-0 hover:bg-bg3 transition-colors'
							>
								<td className='px-[1.4rem] py-[0.95rem]'>
									<span
										className={`font-mono text-[0.68rem] font-bold px-2 py-0.5 rounded-[5px] ${methodClass[ep.method]}`}
									>
										{ep.method}
									</span>
								</td>
								<td className='px-[1.4rem] py-[0.95rem] font-mono text-[0.8rem] text-text'>
									{ep.path}
								</td>
								<td className='px-[1.4rem] py-[0.95rem] font-mono text-[0.8rem] text-text'>
									{ep.desc}
								</td>
								<td className='px-[1.4rem] py-[0.95rem]'>
									<span
										className={`font-mono text-[0.68rem] ${ep.auth ? 'text-accent' : 'text-muted'}`}
									>
										{ep.auth ? '✓ JWT' : '— public'}
									</span>
								</td>
							</tr>
						))}
					</tbody>
				</table>
			</div>
		</section>
	)
}
