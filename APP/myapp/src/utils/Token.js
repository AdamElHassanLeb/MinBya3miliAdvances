const getToken = () => {
    const token = localStorage.getItem('token');
    return token;
};

export const getTokenBearer = () => {
    const token = getToken();
    return `Bearer ${token}`;
}

export default {getTokenBearer}