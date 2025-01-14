const API_BASE_URL = '/api/users/register';

export const registerUser = async (formData)=>{
    const res = await fetch(API_BASE_URL, {
        method: 'POST',
        body: formData,
    });
    console.log(res);
    return res.json();
}

export const authUser = async () => {
    try {
      const res = await fetch('/api/users/authorize', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include', // Ensures cookies (including the HttpOnly token) are sent with the request
      });
  
      if (!res.ok) {
        throw new Error('Failed to fetch user data');
      }
  
      const data = await res.json();
      console.log(data); 
      return data;
    } catch (error) {
      console.error('Error fetching user data:', error);
      throw error; 
    }
  };
  