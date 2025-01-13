
const API_BASE_URL = '/api/jobs';

export const addJob = async (newJob) => {
  const res = await fetch(API_BASE_URL, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(newJob),
  });
  console.log(res)
  return res.json(); // Optional: Return response if needed
};

export const deleteJob = async (id) => {
  const res = await fetch(`${API_BASE_URL}/${id}`, {
    method: 'DELETE',
  });
  return res.json(); // Optional: Return response if needed
};

export const updateJob = async (job) => {
  const res = await fetch(`${API_BASE_URL}/${job.id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(job),
  });
  return res.json(); // Optional: Return response if needed
};

