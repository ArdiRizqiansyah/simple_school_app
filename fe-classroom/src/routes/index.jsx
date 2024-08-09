import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Login from "../pages/auth/Login";
// import { ProtectedRoute } from "./ProtectedRoute";
import UserDashboard from "../pages/user";
import UserClassroom from "../pages/user/classroom";
import UserClassroomCreate from "../pages/user/classroom/create";
import UserStudent from "../pages/user/student";
import UserStudentCreate from "../pages/user/student/create";
import { ProtectedGuestRoute } from "./ProtectedGuestRoute";
import UserClassroomEdit from "../pages/user/classroom/edit";
import UserStudentEdit from "../pages/user/student/edit";
import UserClassroomStudent from "../pages/user/classroom/student";
import { ProtectedRoute } from "./ProtectedRoute";


const Routes = () => {
    // public routing
    const routesForPublic = [];

    // non authenticated routing
    const routesForNonAuthenticatedOnly = [
        {
            path: '/',
            element: <ProtectedGuestRoute />,
            children: [
                {
                    path: "/",
                    name: "login",
                    element: <Login />,
                }
            ]
        }
    ];

    // authenticated routing
    const routesForAuthenticatedOnly = [
        {
            path: '/user',
            element: <ProtectedRoute />,
            children: [
                {
                    path: "dashboard",
                    name: "user.dashboard",
                    element: <UserDashboard />,
                },
                {
                    path: "classroom",
                    name: "user.classroom",
                    element: <UserClassroom />,
                },
                {
                    path: "classroom/create",
                    name: "user.classroom.create",
                    element: <UserClassroomCreate />,
                },
                {
                    path: "classroom/:classroomId/edit",
                    name: "user.classroom.edit",
                    element: <UserClassroomEdit />,
                },
                {
                    path: "classroom/:classroomId/student",
                    name: "user.classroom.student",
                    element: <UserClassroomStudent />,
                },
                {
                    path: "student",
                    name: "user.student",
                    element: <UserStudent />,
                },
                {
                    path: "student/create",
                    name: "user.student.create",
                    element: <UserStudentCreate />,
                },
                {
                    path: "student/:studentId/edit",
                    name: "user.student.edit",
                    element: <UserStudentEdit />,
                }
            ],
        }
    ];

    // combine all routes
    const router = createBrowserRouter([
        ...routesForPublic,
        ...routesForNonAuthenticatedOnly,
        ...routesForAuthenticatedOnly,
    ]);

    return <RouterProvider router={router} />
}

export default Routes;