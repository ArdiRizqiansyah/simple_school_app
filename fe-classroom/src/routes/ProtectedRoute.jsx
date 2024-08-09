import { useSelector } from "react-redux";
import { Navigate, Outlet } from "react-router-dom";

export const ProtectedRoute = () => {
    const userSelector = useSelector((state) => state.user);

    if (!userSelector.token) {
        return <Navigate to="/" />;
    }

    return <Outlet />;
};