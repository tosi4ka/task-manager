import { CreateTaskRequest, Task, UpdateTaskRequest } from '@/types'
import { baseApi } from './baseApi'

export const taskApi = baseApi.injectEndpoints({
	endpoints: builder => ({
		createTask: builder.mutation<Task, CreateTaskRequest>({
			query: credentials => ({
				url: '/task/createTask',
				method: 'POST',
				body: credentials,
			}),
		}),
		updateTask: builder.mutation<Task, { id: string } & UpdateTaskRequest>({
			query: ({ id, ...body }) => ({
				url: `/task/updateTask/${id}`,
				method: 'PATCH',
				body,
			}),
		}),
		deleteTask: builder.mutation<void, string>({
			query: id => ({
				url: `/task/deleteTask/${id}`,
				method: 'DELETE',
			}),
		}),
		listTasks: builder.query<Task[], string>({
			query: id => ({
				url: `/task/tasksList/${id}`,
				method: 'GET',
			}),
		}),
		getById: builder.query<Task, string>({
			query: id => ({
				url: `/task/getById/${id}`,
				method: 'GET',
			}),
		}),
	}),
})

export const {
	useCreateTaskMutation,
	useUpdateTaskMutation,
	useDeleteTaskMutation,
	useListTasksQuery,
	useGetByIdQuery,
} = taskApi
