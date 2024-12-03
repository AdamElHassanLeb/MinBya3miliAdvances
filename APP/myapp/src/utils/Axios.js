import axios from 'axios';
import serverAddress from "./ServerAddress";

const axiosInstance = axios.create({
    baseURL: serverAddress() + '/api/v1',
    /* headers: {
       'Content-Type': 'application/json',

     },*/
});


export default axiosInstance
