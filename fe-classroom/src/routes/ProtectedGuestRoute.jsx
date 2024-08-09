import { useSelector } from "react-redux";
import { Navigate, Outlet } from "react-router-dom"
import Cookies from 'js-cookie';

export const ProtectedGuestRoute = () => {
    const userSelector = useSelector((state) => state.user);
    const token = userSelector.token || Cookies.get('token');

    if (token) {
        return <Navigate to="/user/dashboard" />
    }

    return <Outlet />
}