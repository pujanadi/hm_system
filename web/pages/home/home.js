import 'tailwindcss/tailwind.css';
import SideMenu from '../../components/SideMenu';

export default function Home( { token }) {
  return (
    <div className="min-h-screen bg-gradient-to-r from-cyan-500 to-blue-500">
      <SideMenu />
      <div>
        here is your token {token}
      </div>
    </div>
  )
}