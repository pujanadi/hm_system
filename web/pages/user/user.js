import React from 'react';
import SideMenu from '../../components/SideMenu';
import 'tailwindcss/tailwind.css';

function user() {
  return (
    <div className='min-h-screen bg-gradient-to-r from-cyan-500 to-blue-500'>
        <SideMenu />
        <div>You are at user menu</div>
    </div>
  );
}

export default user;
