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


export default {
    login,
    signUp,
}

