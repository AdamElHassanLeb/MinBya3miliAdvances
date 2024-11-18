import axiosInstance from "../utils/Axios";

const URL = '/user'

const login = async (phone_number, password) => {

    try {
        //localStorage.setItem('token', JSON.stringify(token));
        return await axiosInstance.post(URL + '/auth', {phone_number, password});
    }
    catch (error) {
        console.log(error)
    }
}

const signUp = async (first_name, last_name, phone_number, date_of_birth, profession, location, password) => {

    try {
        return await axiosInstance.post(URL + `/create`, {
            first_name: first_name,
            last_name : last_name,
            phone_number : phone_number,
            date_of_birth : date_of_birth,
            profession : profession,
            location : location,
            password : password,
        });
    }catch (error) {
        console.log(error)
    }
}

const getUserById = async (userId) => {
    try {
        const response = await axiosInstance.get(`${URL}/userId/${userId}`);
        return response.data;  // Return the user data from the response
    } catch (error) {
        console.error("Error fetching user by ID:", error);
        throw error;  // Re-throw error so it can be handled by the component
    }
};



export default {
    login,
    signUp,
    getUserById
}

