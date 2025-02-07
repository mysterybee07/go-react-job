import React from 'react'
import { useState, useEffect } from 'react';
import JobListing from './JobListing';
import Spinner from './Spinner';

const JobListings = ({ isHome = false }) => {
    // const jobListing = isHome ? jobs.slice(0,3) : jobs;
    const [jobs, setJobs] = useState([]);
    const [loading, setLoading] = useState(true);
    useEffect(() => {
        const apiUrl = isHome ? '/api/jobs?limit=3':'/api/jobs';

        const fetchJobs = async () => {
            try {
                const res = await fetch(apiUrl);
                const data = await res.json();
                console.log(data);
                setJobs(data.jobs);
            } catch (error) {
                console.log("Error fetching data", error);
            } finally {
                setLoading(false);
            }
        }
        fetchJobs();

    }, []);


    return (
        <>
            <section className="bg-slate-200 px-4 py-10">
                <div className="container-xl lg:container m-auto">
                    <h2 className="text-3xl font-bold text-indigo-500 mb-6 text-center">
                        {isHome ? 'Recent Jobs' : 'Browse Jobs'}
                    </h2>

                    {/* <!-- Job Listing 1 --> */}
                    {loading ? (
                        <Spinner loading={loading} />
                    ) : (
                        <div className="grid grid-cols-1 md:grid-cols-3 gap-6 ">
                        {/* <div className="flex flex-row jus"> */}
                            {jobs.map((job) => (
                                <JobListing key={job.id} job={job} />
                            ))}
                        </div>
                    )}

                </div>

            </section>
        </>
    )
}

export default JobListings
