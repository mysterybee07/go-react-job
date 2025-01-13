import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom';
import { toast } from 'react-toastify';

const RegisterAsUserPage = ({ registerNewUser }) => {

  const navigate = useNavigate();

  const [name, setName] = useState('');
  const [contactEmail, setContactEmail] = useState('');
  const [contactPhone, setContactPhone] = useState('');
  const [address, setAddress] = useState('');
  const [imageUrl, setImageUrl] = useState('');
  const [resume, setResume] = useState('');
  const [password, setPassword] = useState('');


const submitForm = (e) => {
  e.preventDefault();

  const newUser = {
    name,
    contact_email: contactEmail,
    contact_phone: contactPhone,
    address,
    imageUrl,
    resume,
    password,
  };


  // Call the passed in function to register the user
  registerNewUser(newUser);

  // show success nortification
  toast.success("User registered successfully");

  // Navigate to login page
  navigate('/login');
};

  return (
    <div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
      <div className="sm:mx-auto sm:w-full sm:max-w-sm">
        <img
          alt="Your Company"
          src="https://tailwindui.com/plus/img/logos/mark.svg?color=indigo&shade=600"
          className="mx-auto h-10 w-auto"
        />
        <h2 className="mt-10 text-center text-2xl font-bold tracking-tight text-gray-900">
          Register
        </h2>
      </div>

      <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
        <form onSubmit={submitForm} className="space-y-6">
          <div>
            <label htmlFor="image_url" className="block text-sm font-medium text-gray-900">
              Profile Image
            </label>
            <div className="mt-2 flex justify-center">
              {/* Circular Preview (Clickable) */}
              <label
                htmlFor="image_url"
                className="relative h-32 w-32 rounded-full border-2 border-gray-300 overflow-hidden cursor-pointer"
              >
                {imageUrl ? (
                  <img
                    src={imageUrl}
                    alt="Profile Preview"
                    className="h-full w-full object-cover"
                  />
                ) : (
                  <div className="flex h-full w-full items-center justify-center bg-gray-100 text-gray-500">
                    Click to Upload
                  </div>
                )}
                {/* Hidden File Input */}
                <input
                  id="image_url"
                  name="image_url"
                  type="file"
                  accept="image/*"
                  onChange={(e) => {
                    const file = e.target.files[0];
                    if (file) {
                      const reader = new FileReader();
                      reader.onload = () => setImageUrl(reader.result);
                      reader.readAsDataURL(file);
                    }
                  }}
                  className="hidden"
                />
              </label>
            </div>
          </div>

          {/* Name */}
          <div>
            <label htmlFor="name" className="block text-sm font-medium text-gray-900">
              Full Name
            </label>
            <div className="mt-2">
              <input
                id="name"
                name="name"
                type="text"
                required
                value={name}
                onChange={(e) => setName(e.target.value)}
                className="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline outline-1 outline-gray-300 placeholder-gray-400 focus:outline-indigo-600"
              />
            </div>
          </div>

          {/* Contact Email */}
          <div>
            <label htmlFor="contact_email" className="block text-sm font-medium text-gray-900">
              Email Address
            </label>
            <div className="mt-2">
              <input
                id="contact_email"
                name="contact_email"
                type="email"
                required
                value={contactEmail}
                onChange={(e) => setContactEmail(e.target.value)}
                className="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline outline-1 outline-gray-300 placeholder-gray-400 focus:outline-indigo-600"
              />
            </div>
          </div>

          {/* Contact Phone */}
          <div>
            <label htmlFor="contact_phone" className="block text-sm font-medium text-gray-900">
              Phone Number
            </label>
            <div className="mt-2">
              <input
                id="contact_phone"
                name="contact_phone"
                type="text"
                required
                value={contactPhone}
                onChange={(e) => setContactPhone(e.target.value)}
                className="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline outline-1 outline-gray-300 placeholder-gray-400 focus:outline-indigo-600"
              />
            </div>
          </div>

          {/* Address */}
          <div>
            <label htmlFor="address" className="block text-sm font-medium text-gray-900">
              Address
            </label>
            <div className="mt-2">
              <input
                id="address"
                name="address"
                type="text"
                required
                value={address}
                onChange={(e) => setAddress(e.target.value)}
                className="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline outline-1 outline-gray-300 placeholder-gray-400 focus:outline-indigo-600"
              />
            </div>
          </div>


          {/* Submit Button */}
          <div>
            <button
              type="submit"
              className="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus:outline-indigo-600"
            >
              Register
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}

export default RegisterAsUserPage
