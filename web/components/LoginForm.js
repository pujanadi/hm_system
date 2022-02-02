import React, { useState, SyntheticEvent } from 'react';
import { Tab } from '@headlessui/react';
import 'tailwindcss/tailwind.css';
import Link from 'next/link';
import jwt from 'jsonwebtoken';
import Home from '../pages/home/home';

const LoginForm = ( props ) => {

  const [username, setUsername] = useState();
  const [password, setPassword] = useState();
  const [message, setMessage] = useState();

  const handleLogin = async (event) => {
    event.preventDefault();

    const response = await fetch ( 'http://localhost:8080/auth',
      {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          // name: event.target.name.value
          Username: username,
          Password: password
        }),
      }
    ).then((t) => t.json());

    if(response.token){
      props.func(response.token);
      const json = await jwt.decode(response.token);
      console.log(json);
      
    } else {
      setMessage('Login failed');
    }
    
  }

  return (
    <div className="p-5 bg-white my-10 rounded font-mono">
      <div className='text-center text-red-500'>
        <span>{message}</span>
      </div>

      <form onSubmit={handleLogin} className='from-orange-50'>
        <div>
          <div>
            <span>Username</span>
          </div>
          <div className="pb-5">
            <input type="text" name="username" className="rounded" onChange={(e) => setUsername(e.target.value)} autoComplete='username' required />
          </div>

          <div>
            <span>Password</span>
          </div>
          <div className="pb-5">
            <input type="password" name="password" className="rounded" onChange={(e) => setPassword(e.target.value)} autoComplete='current-password' required />
          </div>

          <div>
            <input type="submit" value="Login" className="color-white container bg-gradient-to-r from-cyan-500 to-blue-500 py-2 px-5 rounded" />
          </div>
        </div>
          
      </form>
    </div>
  );
}

export default LoginForm;

