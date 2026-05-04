export interface User {
	id: string
	name: string
	email: string
	created_at: string
	updated_at: string
}

export interface Task {
	id: string
	title: string
	description: string
	assigned_by: string
	assigned_to: string
	estimate: number
	status: 'todo' | 'in_progress' | 'done'
	created_at: string
	updated_at: string
	completed_at: string | null
}

export interface AuthResponse {
	access_token: string
	refresh_token: string
	user: User
}

export interface LoginRequest {
	email: string
	password: string
}

export interface RegisterRequest {
	name: string
	email: string
	password: string
}

export interface CreateTaskRequest {
	title: string
	description: string
	assigned_to: string
	assigned_by: string
	estimate: number
}

export interface UpdateTaskRequest {
	title?: string
	description?: string
	assigned_to?: string
	estimate?: number
	status?: 'todo' | 'in_progress' | 'done'
	completed_at?: string | null
}
