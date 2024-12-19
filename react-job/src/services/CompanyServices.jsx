const API_BASE_URL = '/api/companies/register'


export const registerCompany = async(newCompany) =>{
    const res = await fetch(API_BASE_URL, {
        method:'POST',
        headers:{
            'Content-Type':'application/json',
        },
        body: JSON.stringify(newCompany),
    });
    return res.json();
}