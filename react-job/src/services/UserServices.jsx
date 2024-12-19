const API_BASE_URL = '/api/users/login';


export const loginUser = async (login) => {
    const res = await fetch(API_BASE_URL, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(login),
    });
  
    if (!res.ok) {
      const error = await res.json(); // Assuming the server returns a JSON error message
      throw new Error(error.message || 'Login failed');
    }
  
    return res.json(); // Proceed if the response is successful
  };
  

  // export const loginUser = async (login) => {
//   try {
//     const res = await fetch(API_BASE_URL, {
//       method: 'POST',
//       headers: {
//         'Content-Type': 'application/json',
//       },
//       body: JSON.stringify(login),
//     });

//     // Check if the response status is OK (200-299)
//     if (!res.ok) {
//       throw new Error(`Login failed: ${res.statusText}`);
//     }

//     const data = await res.json();

//     // Return the response data (you can modify this based on your API's response)
//     return data;
//   } catch (error) {
//     // Handle errors (network issues, invalid response, etc.)
//     console.error('Login error:', error);
//     throw error;  // Re-throw the error if you want to handle it elsewhere
//   }
// };