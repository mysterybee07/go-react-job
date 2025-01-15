// components/Navbar.js
import React from 'react';
import { NavLink, useNavigate } from 'react-router-dom';
import { useAuth } from '../Contexts/AuthContext';
import logo from '../assets/images/logo.png';
import { logoutUser } from '../services/AuthServices';

const Navbar = () => {
  const navigate = useNavigate();
  const { isAuthenticated, logout } = useAuth();

  const handleLogout = async () => {
    const success = await logoutUser(); // Call your logout service
    if (success) {
      logout(); // Update the global authentication state
      navigate('/login');
    } else {
      alert('Logout failed. Please try again.');
    }
  };

  const linkClass = ({ isActive }) =>
    isActive
      ? 'bg-black text-white hover:bg-gray-900 hover:text-white rounded-md px-3 py-2'
      : 'text-white hover:bg-gray-900 hover:text-white rounded-md px-3 py-2';

  return (
    <nav className="bg-slate-700 border-b border-violet-200">
      <div className="mx-auto max-w-7xl px-2 sm:px-6 lg:px-8">
        <div className="flex h-20 items-center justify-between">
          <div className="flex flex-1 items-center justify-center md:items-stretch md:justify-start">
            <NavLink className="flex flex-shrink-0 items-center mr-4" to="/">
              <img className="h-10 w-auto" src={logo} alt="React Jobs" />
              <span className="hidden md:block text-white text-2xl font-bold ml-2">
                Kaam Deu
              </span>
            </NavLink>
            <div className="md:ml-auto">
              <div className="flex space-x-2">
                <NavLink to="/" className={linkClass}>
                  Home
                </NavLink>
                {!isAuthenticated && (
                  <>
                    <NavLink to="/login" className={linkClass}>
                      Sign In
                    </NavLink>
                    <NavLink to="/register" className={linkClass}>
                      Register
                    </NavLink>
                  </>
                )}
                {isAuthenticated && (
                  <button onClick={handleLogout} className={linkClass({ isActive: false })}>
                    Logout
                  </button>
                )}
              </div>
            </div>
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;