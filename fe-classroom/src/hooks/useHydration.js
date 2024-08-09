import { useEffect, useState } from "react";
import { useDispatch } from "react-redux";
import Cookies from 'js-cookie';
import { axiosInstance } from "../lib/axios";

export const useHydration = () => {
    const dispatch = useDispatch();

    const [loading, setLoading] = useState(true);

    const hydrateAuth = async () => {
        try {
            const user = Cookies.get('user');
            const token = Cookies.get('token');

            if (!user && !token) {
                return ;
            }

            const userResponse = await axiosInstance.get("/users/profile", {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });

            if (user && token) {
                dispatch({
                    type: 'LOGIN',
                    payload: {
                        name: userResponse.data.name,
                        email: userResponse.data.email,
                        token: token,
                    }
                });
            }
        } catch (error) {
            console.log('Error hydrating auth: ', error);
        } finally {
            setLoading(false);
        }
    }

    useEffect(() => {
        hydrateAuth();
    }, []);

    return { loading };
}