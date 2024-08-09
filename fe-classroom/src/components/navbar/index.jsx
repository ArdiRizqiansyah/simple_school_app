import { useDispatch, useSelector } from "react-redux";
import { Link } from "react-router-dom";

export default function Navbar() {
    const userSelector = useSelector((state) => state.user);
    const dispatch = useDispatch();

    const handleLogout = () => {
        dispatch({
            type: "LOGOUT",
        });
    }

    return (
        <>
            <nav className="bg-white border-gray-200">
                <div className="flex flex-wrap justify-between items-center mx-auto max-w-screen-xl p-4">
                    <a
                        href="https://flowbite.com"
                        className="flex items-center space-x-3 rtl:space-x-reverse"
                    >
                        <img
                            src="https://flowbite.com/docs/images/logo.svg"
                            className="h-8"
                            alt="Flowbite Logo"
                        />
                        <span className="self-center text-2xl font-semibold whitespace-nowrap">
                            School App
                        </span>
                    </a>
                    <div className="flex items-center space-x-6 rtl:space-x-reverse">
                        <img className="w-8 h-8 rounded-full" src="/public/assets/images/user.png" alt="user photo" />
                        <span className="text-sm">Hello, {userSelector.name}</span>
                    </div>
                </div>
            </nav>
            <nav className="bg-gray-50">
                <div className="max-w-screen-xl px-4 py-3 mx-auto">
                    <div className="flex items-center justify-between">
                        <ul className="flex flex-row font-medium mt-0 space-x-8 rtl:space-x-reverse text-sm">
                            <li>
                                <Link
                                    to="/user/dashboard"
                                    className="text-gray-900 hover:underline"
                                    aria-current="page"
                                >
                                    Dashboard
                                </Link>
                            </li>
                            <li>
                                <Link
                                    to="/user/classroom"
                                    className="text-gray-900 hover:underline"
                                >
                                    Kelas
                                </Link>
                            </li>
                            <li>
                                <Link
                                    to="/user/student"
                                    className="text-gray-900 hover:underline"
                                >
                                    Siswa
                                </Link>
                            </li>
                        </ul>
                        <ul className="flex flex-row font-medium mt-0 space-x-8 rtl:space-x-reverse text-sm">
                            <li>
                                <button
                                    onClick={handleLogout}
                                    className="text-gray-900 hover:underline"
                                >
                                    Logout
                                </button>
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>
        </>
    )
}
