import React, { useEffect, useState } from 'react';
import { toast } from 'react-toastify';

const UserProfilePage = ({ authorizedUser }) => {
  const [userData, setUserData] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const data = await authorizedUser();
        setUserData(data);
      } catch (error) {
        console.error('Error fetching user data:', error);
        setError('Failed to fetch user data. Please try again.');
        toast.error('Failed to fetch user data. Please try again.');
      } finally {
        setIsLoading(false);
      }
    };

    fetchUserData();
  }, [authorizedUser]);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>{error}</div>;
  }

  if (!userData) {
    return <div>No user data found.</div>;
  }

  return (
    <div>
      <h1>User Profile</h1>
      <p>Name: {userData.name}</p>
      <p>Email: {userData.email}</p>
      <p>Message: {userData.message}</p>

      {/* Display Profile Image */}
      {userData.image && (
        <img
          src={`http://localhost:8080/${userData.image}`}
          alt="Profile"
          width={256}
        />
      )}

      {/* Display Resume */}
      {userData.resume && (
        <a
          href={`http://localhost:8080/${userData.resume}`}
          target="_blank"
          rel="noopener noreferrer"
        >
          View Resume
        </a>
      )}
    </div>
  );
};

export default UserProfilePage;