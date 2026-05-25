export const features = [
	{
		icon: '🔐',
		name: 'JWT Auth',
		desc: 'Access + refresh tokens. JWT middleware on all protected routes via Gin.',
		tag: 'internal/auth',
	},
	{
		icon: '🛡️',
		name: 'Rate Limiting',
		desc: '60 requests per minute per IP. Counters stored in Redis.',
		tag: 'Redis · middleware',
	},
	{
		icon: '📋',
		name: 'Tasks CRUD',
		desc: 'Create, update, read and delete tasks. Each task belongs to a user.',
		tag: 'internal/task',
	},
	{
		icon: '🏛️',
		name: 'Clean Architecture',
		desc: 'Strict separation Handler → Service → Repository. Easily testable and scalable.',
		tag: 'sqlx · PostgreSQL',
	},
	{
		icon: '📖',
		name: 'Swagger Docs',
		desc: 'Auto-generated interactive docs. Available at /swagger/index.html.',
		tag: 'docs/ · Swagger UI',
	},
	{
		icon: '🐳',
		name: 'Docker Compose',
		desc: 'One-command deploy. PostgreSQL, Redis and API orchestrated together.',
		tag: 'docker compose up',
	},
]
