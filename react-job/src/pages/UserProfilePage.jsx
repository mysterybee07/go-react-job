import React, { useEffect, useState } from 'react';

const UserProfilePage = ({ authorizedUser }) => {
  const [userData, setUserData] = useState(null);

  useEffect(() => {
    const fetchUserData = async () => {
      const data = await authorizedUser();
      setUserData(data);
    };
    fetchUserData();
  }, [authorizedUser]);

  if (!userData) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <h1>User Profile</h1>
      <p>Name: {userData.name}</p>
      <p>Email: {userData.email}</p>
      <p>Message: {userData.message}</p>
      
      {/* Display Profile Image */}
      {userData.image && <img src={`http://localhost:8080/${userData.image}`} alt="Profile" width={256}/>}
      
      {/* Display Resume */}
      {userData.resume && (
        <a href={`http://localhost:8080/${userData.resume}`} target="_blank" rel="noopener noreferrer">
          View Resume
        </a>
      )}
    </div>
  );
};

export default UserProfilePage;
