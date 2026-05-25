import Endpoints from '@/components/landing/Endpoints'
import Features from '@/components/landing/Features'
import Footer from '@/components/landing/Footer'
import Hero from '@/components/landing/Hero'
import Navbar from '@/components/landing/Navbar'
import StackStrip from '@/components/landing/StackStrip'
import Stats from '@/components/landing/Stats'

export default function Home() {
	return (
		<>
			<Navbar />
			<Hero />
			<Stats />
			<Endpoints />
			<Features />
			<StackStrip />
			<Footer />
		</>
	)
}
