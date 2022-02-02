import React, { useState } from 'react';
import Head from 'next/head';
// import '../styles/global.css';
import 'tailwindcss/tailwind.css';
import '../styles/login.module.css';
import LoginForm from '../components/LoginForm';
import Home from './home/home';

import jwt from 'jsonwebtoken';

const Login = () => {
  const [token, setToken] = useState("adfsdfsafa");

  const getToken = (data) => {
    setToken(data);
    console.log("Login data : ", data);
    console.log("Login token ", token)
  }

    return (
      <>
      { token ? 
        <Home token={token} /> :
        <div className="min-h-screen bg-gradient-to-r from-cyan-500 to-blue-500">
          <Head>
            <title>HM System</title>
            <link rel="icon" href="/favicon.ico" />
          </Head>
    
          <main className="flex flex-col justify-center items-center">

            
            <span className="text-7xl bg-clip-text text-white">
              HM System
            </span>
    
            <div className="description mt-10 bg-neutral-100 py-3 rounded">
              <span className="font-thin px-5 py-10">Your management system!</span>
            </div>
    
            <div>
              <LoginForm func={getToken} />
            </div>
    

          </main>

          <footer className="flex justify-center items-end">
            <a
              href="https://vercel.com?utm_source=create-next-app&utm_medium=default-template&utm_campaign=create-next-app"
              target="_blank"
              rel="noopener noreferrer"
            >
              Powered by{' '}
              {/* <img src="/vercel.svg" alt="Vercel" className="logo" /> */}
              HM Teams
            </a>
          </footer>
          
        </div>
      
    }
    </>
    )
}

export default Login;
