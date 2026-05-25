export const endpoints = [
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
