import axiosInstance from "../utils/Axios";
import {getTokenBearer} from "../utils/Token";

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
        console.log(error + "Error Signing Up");
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

const getUsersByUsername = async (username) => {
    try {
        const response = await axiosInstance.get(`${URL}/userName/${username}`);
        return response.data;  // Return the user data from the response
    } catch (error) {
        //console.error("Error fetching user by ID:", error);
        throw error;  // Re-throw error so it can be handled by the component
    }
}


const updateUser = async (userId, updatedData) => {
    try {
        const token = getTokenBearer(); // Get the JWT token
        return await axiosInstance.put(`${URL}/update/${userId}`, updatedData, {
            headers: { Authorization: token }
        });
    } catch (error) {
        console.error("Error updating user:", error);
        throw error; // Re-throw for component handling
    }
};

const deleteUser = async (userId) => {
    try {
        const token = getTokenBearer(); // Get the JWT token
        return await axiosInstance.delete(`${URL}/delete/${userId}`, {
            headers: { Authorization: token }
        });
    } catch (error) {
        console.error("Error deleting user:", error);
        throw error; // Re-throw for component handling
    }
};


export default {
    login,
    signUp,
    getUserById,
    getUsersByUsername,
    updateUser,
    deleteUser
}

