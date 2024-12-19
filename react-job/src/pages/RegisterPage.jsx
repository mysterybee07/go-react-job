import React from 'react'
import { Link } from 'react-router-dom'
import Card from '../Components/Card'


const RegisterPage = () => {
  return (
    <>
      <section className="bg-slate-200">
                <div className="container-xl lg:container m-auto">
                    <div className="grid grid-cols-1 md:grid-cols-2 gap-4 p-4 rounded-lg text-center pt-32">
                        <Card>
                            <h2 className="text-2xl font-bold">Are You Searching for Jobs?</h2>
                            <p className="mt-2 mb-4">
                                Browse jobs and start your career today
                            </p>
                            <Link
                                to="/register/user"
                                className="inline-block bg-black text-white rounded-lg px-4 py-2 hover:bg-gray-700"
                            >
                                Register as Seekers
                            </Link>
                        </Card>
                        <Card bg='bg-slate-300'>
                        <h2 className="text-2xl font-bold">Are You Looking for Experts?</h2>
                            <p className="mt-2 mb-4">
                                Find Experts to serve you from across the world.
                            </p>
                            <Link
                                to="/register/company"
                                className="inline-block bg-indigo-500 text-white rounded-lg px-4 py-2 hover:bg-indigo-600"
                            >
                                Register As Company
                            </Link>
                        </Card>
                    </div>
                </div>
            </section>
    </>
  )
}

export default RegisterPage
