const API_BASE_URL = '/api/companies/register';

export const registerCompany = async (formData) => {
  const res = await fetch(API_BASE_URL, {
    method: 'POST',
    body: formData, // Send FormData directly (no headers needed)
  });
//   console.log(res);
  return res.json();
};