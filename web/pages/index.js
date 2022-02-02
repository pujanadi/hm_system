import Login from './Login';
import LoginForm from '../components/LoginForm';

export default function Home({data}) {
  return <Login />;
  // return <LoginForm />
  // return <div>{data}</div>;
}

// export const getServerSideProps = async () => {
//   const res = await fetch('https://api.github.com/repos/developit/preact');
//   const data = await res.json();

//   return {
//     props: { data }
//   }
// }
