import React, { useEffect } from 'react'
import { useState } from 'react'
import { useNavigate } from 'react-router-dom';
import { toast } from 'react-toastify';

const AddJobPage = ({ addJobSubmit }) => {
    const [title, setTitle] = useState('');
    const [type, setType] = useState('Full-Time');
    const [location, setLocation] = useState('');
    const [description, setDescription] = useState('');
    const [salary, setSalary] = useState('Under $50K');
    // const [company_id, setCompany_id] = useState('Under $50K');
    // const [companyName, setCompanyName] = useState('');
    // const [companyDescription, setCompanyDescription] = useState('');
    // const [contactEmail, setContactEmail] = useState('');
    // const [contactPhone, setContactPhone] = useState('');

    const [companies, setCompanies] = useState([]); // State for fetched companies
    const [selectedCompany, setSelectedCompany] = useState(""); // State for the selected company
    const [loading, setLoading] = useState(true); // State to handle loading

    const navigate = useNavigate();

    const submitForm = (e) => {
        e.preventDefault();

        if (!selectedCompany) {
            toast.error('Please select a company.');
            return;
        }

        const newJob = {
            title,
            type,
            location,
            description,
            salary,
            company_id: Number(selectedCompany),
        };

        addJobSubmit(newJob);
        toast.success('Job Added Successfully');
        return navigate('/jobs');
    }

    // Fetch companies on component mount
    useEffect(() => {
        const apiUrl = "/api/companies";
        const fetchCompanies = async () => {
            try {
                const res = await fetch(apiUrl);
                const data = await res.json();
                console.log(data);
                setCompanies(data.companies);
            } catch (error) {
                console.error("Error fetching companies:", error);
            } finally {
                setLoading(false);
            }
        };

        fetchCompanies();
    }, []);

    return (
        <>
            <section className="bg-indigo-50">
                <div className="container m-auto max-w-2xl py-24">
                    <div
                        className="bg-white px-6 py-8 mb-4 shadow-md rounded-md border m-4 md:m-0"
                    >
                        <form onSubmit={submitForm}>
                            <h2 className="text-3xl text-center font-semibold mb-6">Add Job</h2>

                            <div className="mb-4">
                                <label htmlFor="type" className="block text-gray-700 font-bold mb-2"
                                >Job Type</label
                                >
                                <select
                                    id="type"
                                    name="type"
                                    className="border rounded w-full py-2 px-3"
                                    required
                                    value={type}
                                    onChange={(e) => setType(e.target.value)}
                                >
                                    <option value="Full-Time">Full-Time</option>
                                    <option value="Part-Time">Part-Time</option>
                                    <option value="Remote">Remote</option>
                                    <option value="Internship">Internship</option>
                                </select>
                            </div>

                            <div className="mb-4">
                                <label className="block text-gray-700 font-bold mb-2"
                                >Job Listing Name</label
                                >
                                <input
                                    type="text"
                                    id="title"
                                    name="title"
                                    className="border rounded w-full py-2 px-3 mb-2"
                                    placeholder="eg. Beautiful Apartment In Miami"
                                    required
                                    value={title}
                                    onChange={(e) => setTitle(e.target.value)}
                                />
                            </div>
                            <div className="mb-4">
                                <label
                                    htmlFor="description"
                                    className="block text-gray-700 font-bold mb-2"
                                >Description</label
                                >
                                <textarea
                                    id="description"
                                    name="description"
                                    className="border rounded w-full py-2 px-3"
                                    rows="4"
                                    placeholder="Add any job duties, expectations, requirements, etc"
                                    value={description}
                                    onChange={(e) => setDescription(e.target.value)}
                                ></textarea>
                            </div>

                            <div className="mb-4">
                                <label htmlFor="type" className="block text-gray-700 font-bold mb-2"
                                >Salary</label
                                >
                                <select
                                    id="salary"
                                    name="salary"
                                    className="border rounded w-full py-2 px-3"
                                    required
                                    value={salary}
                                    onChange={(e) => setSalary(e.target.value)}
                                >
                                    <option value="Under $50K">Under $50K</option>
                                    <option value="$50K - 60K">$50K - $60K</option>
                                    <option value="$60K - 70K">$60K - $70K</option>
                                    <option value="$70K - 80K">$70K - $80K</option>
                                    <option value="$80K - 90K">$80K - $90K</option>
                                    <option value="$90K - 100K">$90K - $100K</option>
                                    <option value="$100K - 125K">$100K - $125K</option>
                                    <option value="$125K - 150K">$125K - $150K</option>
                                    <option value="$150K - 175K">$150K - $175K</option>
                                    <option value="$175K - 200K">$175K - $200K</option>
                                    <option value="Over $200K">Over $200K</option>
                                </select>
                            </div>

                            <div className='mb-4'>
                                <label className='block text-gray-700 font-bold mb-2'>
                                    Location
                                </label>
                                <input
                                    type='text'
                                    id='location'
                                    name='location'
                                    className='border rounded w-full py-2 px-3 mb-2'
                                    placeholder='Company Location'
                                    required
                                    value={location}
                                    onChange={(e) => setLocation(e.target.value)}
                                />
                            </div>

                            <div>
                                <h3 className="text-2xl mb-5">Company Info</h3>

                                {loading ? (
                                    <p>Loading companies...</p>
                                ) : (
                                    <div className="mb-4">
                                        <label
                                            htmlFor="company"
                                            className="block text-gray-700 font-bold mb-2"
                                        >
                                            Select Company
                                        </label>
                                        <select
                                            id="company"
                                            name="company_id"
                                            className="border rounded w-full py-2 px-3"
                                            value={selectedCompany}
                                            onChange={(e) => setSelectedCompany(e.target.value)}
                                        >
                                            <option value="" disabled>
                                                -- Select a Company --
                                            </option>
                                            {companies.map((company) => (
                                                <option key={company.ID} value={company.ID}>
                                                    {company.name}
                                                </option>
                                            ))}
                                        </select>
                                    </div>
                                )}
                            </div>

                            <div>
                                <button
                                    className="bg-indigo-500 hover:bg-indigo-600 text-white font-bold py-2 px-4 rounded-full w-full focus:outline-none focus:shadow-outline"
                                    type="submit"
                                >
                                    Add Job
                                </button>
                            </div>
                        </form>
                    </div>
                </div>
            </section>
        </>
    )
}

export default AddJobPage
