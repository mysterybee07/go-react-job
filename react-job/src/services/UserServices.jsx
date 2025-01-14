const API_BASE_URL = '/api/users/register';

export const registerUser = async (formData)=>{
    const res = await fetch(API_BASE_URL, {
        method: 'POST',
        body: formData,
    });
    console.log(res);
    return res.json();
}