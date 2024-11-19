import axiosInstance from "../utils/Axios";
import Token from "../utils/Token";

const URL = '/image';

const uploadListingImage = async (listingId, imageFiles) => {
    try {
        const formData = new FormData();
        imageFiles.forEach((file) => {
            formData.append("images", file); // Append each image with the key 'images'
        });

        const token = Token.getTokenBearer()
        // Send the request with the Authorization header
        const response = await axiosInstance.post(`${URL}/uploadForListing/${listingId}`, formData, {
            headers: {
                'Content-Type': 'multipart/form-data',
                'Authorization': token, // Add the token here
            },
        });
        return response;
    } catch (error) {
        console.error("Error uploading listing images:", error);
    }
};

const uploadProfileImage = async (userId, imageFile) => {
    try {
        const formData = new FormData();
        formData.append("file", imageFile);

        const response = await axiosInstance.post(`${URL}/uploadProfilePicture/${userId}`, formData, {
            headers: {
                'Content-Type': 'multipart/form-data',
            },
        });
        return response;
    } catch (error) {
        console.error("Error uploading profile image:", error);
    }
};

const getImageById = async (imageId) => {
    try {
        const response = await axiosInstance.get(`${URL}/imageId/${imageId}`);
        return response;
    } catch (error) {
        console.error("Error fetching image by ID:", error);
    }
};

const getImagesByListingId = async (listingId) => {
    try {
        const response = await axiosInstance.get(`${URL}/listing/${listingId}`);

        //console.log(response.data.images[0])

        return response.data.images
    } catch (error) {
        console.error("Error fetching images for listing:", error);
        return [];
    }
};


const getImagesByUserId = async (userId) => {
    try {
        const token = Token.getTokenBearer()
        const response = await axiosInstance.get(`${URL}/user/${userId}`, {
            headers: {
                'Authorization': token, // Add the token here with Bearer prefix
            },
        });
        if (response.data && response.data.images) {
            return response.data.images.map(image => image.base64Image);
        }
        return [];
    } catch (error) {
        console.error("Error fetching images for user:", error);
        return [];
    }
};


const getImagesByUserProfile = async (userId) => {
    try {
        const response = await axiosInstance.get(`${URL}/profile/${userId}`);
        return response;
    } catch (error) {
        console.error("Error fetching profile images:", error);
    }
};

const deleteImage = async (imageId) => {
    try {
        const response = await axiosInstance.delete(`${URL}/${imageId}`);
        return response;
    } catch (error) {
        console.error("Error deleting image:", error);
    }
};

const updateImage = async (imageId, showOnProfile) => {
    try {
        const response = await axiosInstance.put(`${URL}/${imageId}/${showOnProfile}`);
        return response;
    } catch (error) {
        console.error("Error updating image visibility:", error);
    }
};

export default {
    uploadListingImage,
    uploadProfileImage,
    getImageById,
    getImagesByListingId,
    getImagesByUserId,
    getImagesByUserProfile,
    deleteImage,
    updateImage,
};
