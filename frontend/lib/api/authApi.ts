import { AuthResponse, LoginRequest, RegisterRequest } from '@/types'
import { baseApi } from './baseApi'

export const authApi = baseApi.injectEndpoints({
	endpoints: builder => ({
		login: builder.mutation<AuthResponse, LoginRequest>({
			query: credentials => ({
				url: '/auth/login',
				method: 'POST',
				body: credentials,
			}),
		}),
		register: builder.mutation<AuthResponse, RegisterRequest>({
			query: credentials => ({
				url: '/auth/register',
				method: 'POST',
				body: credentials,
			}),
		}),
	}),
})

export const { useLoginMutation, useRegisterMutation } = authApi
