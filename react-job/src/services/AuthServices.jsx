const API_BASE_URL = '/api/users/login';

export const loginUser = async (login) => {
    const res = await fetch(API_BASE_URL, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      
      body: JSON.stringify(login),
      credentials: 'include',
    });
  
    if (!res.ok) {
      const error = await res.json(); // Assuming the server returns a JSON error message
      throw new Error(error.message || 'Login failed');
    }
  
    return res.json(); // Proceed if the response is successful
  };
  

  export const logoutUser = async () => {
    try {
      const response = await fetch('/api/users/logout', {
        method: 'POST',
        credentials: 'include', // Include cookies in the request
      });
  
      if (response.ok) {
        return true; // Logout successful
      } else {
        throw new Error('Logout failed');
      }
    } catch (error) {
      console.error('Error during logout:', error);
      return false;
    }
  };

  // services/AuthServices.js
export const checkAuth = async () => {
  try {
    const response = await fetch('/api/users/check-auth', {
      credentials: 'include', // Include cookies in the request
    });

    if (response.ok) {
      const data = await response.json();
      return data.isAuthenticated; // Returns true or false
    } else {
      throw new Error('Failed to check authentication');
    }
  } catch (error) {
    console.error('Error checking authentication:', error);
    return false;
  }
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


