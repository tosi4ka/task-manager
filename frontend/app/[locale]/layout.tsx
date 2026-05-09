import { ReduxProvider } from '@/lib/store/provider'
import type { Metadata } from 'next'
import { NextIntlClientProvider } from 'next-intl'
import { getMessages } from 'next-intl/server'
import { Geist, Geist_Mono, Syne } from 'next/font/google'
import '../globals.css'

const geistSans = Geist({
	variable: '--font-geist-sans',
	subsets: ['latin'],
})

const geistMono = Geist_Mono({
	variable: '--font-geist-mono',
	subsets: ['latin'],
})

const syne = Syne({
	variable: '--font-syne-family',
	subsets: ['latin'],
	weight: ['400', '600', '700', '800'],
})

export const metadata: Metadata = {
	title: 'Task Manager',
	description: 'Task management application',
}

export default async function RootLayout({
	children,
	params,
}: Readonly<{
	children: React.ReactNode
	params: Promise<{ locale: string }>
}>) {
	const { locale } = await params
	const messages = await getMessages()

	return (
		<html
			lang={locale}
			className={`${geistSans.variable} ${geistMono.variable} ${syne.variable} h-full antialiased`}
		>
			<body className='min-h-full flex flex-col'>
				<NextIntlClientProvider locale={locale} messages={messages}>
					<ReduxProvider>{children}</ReduxProvider>
				</NextIntlClientProvider>
			</body>
		</html>
	)
}
