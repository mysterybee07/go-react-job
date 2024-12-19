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
import { addJob, deleteJob, updateJob } from './services/JobServices'
import LoginPage from './pages/LoginPage';
import RegisterPage from './pages/RegisterPage';
import RegisterAsUserPage from './pages/RegisterAsUserPage';
import RegisterAsCompanyPage from './pages/RegisterAsCompanyPage';
import { loginUser } from './services/UserServices'

const App = () => {
  const router = createBrowserRouter(
    createRoutesFromElements(
      <Route path="/" element={<MainLayout />}>
        <Route index element={<HomePage />} />
        <Route path="/jobs" element={<JobsPage />} />
        {/* <Route path='/login' element= {<LoginPage />}/> */}
        <Route path='/register' element= {<RegisterPage />}/>
        <Route path='/register/user' element= {<RegisterAsUserPage />}/>
        <Route path='/register/company' element= {<RegisterAsCompanyPage />}/>
        <Route path="/add-job" element={<AddJobPage addJobSubmit={addJob} />} />
        <Route path="/login" element={<LoginPage loginSubmit={loginUser} />} />

        {/* <Route path='/'/> */}
        <Route
          path="/jobs/:id"
          element={<JobPage deleteJob={deleteJob} />}
          loader={jobLoader}
        />
        <Route
          path="/edit-job/:id"
          element={<EditJobPage updateJobSubmit={updateJob} />}
          loader={jobLoader}
        />
        <Route path="/*" element={<NotFoundPage />} />
      </Route>
    )
  );

  return <RouterProvider router={router} />;
};

export default App;
