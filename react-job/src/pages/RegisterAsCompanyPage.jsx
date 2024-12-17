import React, { useState } from 'react'
import { toast } from 'react-toastify';

const RegisterAsCompanyPage = ({registerNewCompany}) => {
    
const[name, setName ] = useState();
const[contactEmail, setContactEmail] = useState();
const[contactPhone, setContactPhone] = useState();
const[address, setAddress] = useState();
const[imageUrl, setImageUrl] = useState();
const[description, setDescription] = useState();

const submitForm=(e)=> {
    e.preventDefault();

    const newCompany = {
        name,
        contactEmail,
        contactPhone,
        address, 
        imageUrl,
        description
    }

    registerNewCompany(newCompany);
    toast.success('Company registered successfully');
    return Navigate('/login')
}

  return (
    <div>
      
    </div>
  )
}

export default RegisterAsCompanyPage
