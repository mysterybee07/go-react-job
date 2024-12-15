import React from 'react'
import {
  Route,
  createBrowserRouter,
  createRoutesFromElements,
  RouterProvider
} from 'react-router-dom'
import HomePage from './pages/HomePage';
import MainLayout from './Layouts/MainLayout';
import JobsPage from './pages/JobsPage';
import AddJobPage from './pages/AddJobPage';
import NotFoundPage from './pages/NotFoundPage';
import JobPage, { jobLoader } from './pages/JobPage';
import EditJobPage from './pages/EditJobPage';


const App = () => {

  const addJob = async (newJob) => {

    const res = await fetch('/api/jobs',{
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(newJob)
    });
    // console.log(newJob);
    return
  };

  const deleteJob = async (id) => {

    const res = await fetch(`/api/jobs/${id}`, {
      method:'DELETE',      
    })
    // console.log('delete', id);
  };

  const updatedJob = async (job) => {
    const res = await fetch(`/api/jobs/${job.id}`,{
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(job)
    });
    // console.log(newJob);
    return
  }
  
  const router = createBrowserRouter(
    createRoutesFromElements(
      <Route path='/' element={<MainLayout />}>
        <Route index element={<HomePage />} />
        <Route path='/jobs' element={<JobsPage />} />      
        <Route path='/add-job' element={<AddJobPage addJobSubmit={addJob}/>} />      
        <Route path='/jobs/:id' element={<JobPage deleteJob={deleteJob}/>}  loader={jobLoader} />      
        <Route path='/edit-job/:id' element={<EditJobPage updateJobSubmit={updatedJob}/>}  loader={jobLoader} />      
        {/* <Route path='/*' element={<NotFoundPage />} />       */}
        <Route path='/*' element={<NotFoundPage />} />      
      </Route>  
    )
  )

  return <RouterProvider router={router} />;
};

export default App
