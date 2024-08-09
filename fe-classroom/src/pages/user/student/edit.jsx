import UserLayout from "@Layout/User";
import { useEffect, useState } from "react";
import { useSelector } from "react-redux";
import { Link, useParams } from "react-router-dom";
import { axiosInstance } from "../../../lib/axios";
import toast from "react-hot-toast";
import { StudentForm } from "../../../components/navbar/forms/StudentForm";

export default function StudentEdit() {
    const userSelector = useSelector((state) => state.user);

    const [student, setStudent] = useState({
        name: "",
        nis: "",
        place_birth: "",
        date_birth: "",
    });
    const params = useParams();

    const fetchStudent = async () => {
        try {
            const response = await axiosInstance.get(`/students/${params.studentId}`, {
                headers: {
                    Authorization: `Bearer ${userSelector.token}`,
                },
            });

            console.log(response.data);

            setStudent(response.data);
        } catch (error) {
            console.error(error);
        }
    }

    useEffect(() => {
        fetchStudent();
    }, []);

    const handleEditStudent = async (values) => {
        try {
            await axiosInstance.put(`/students/${params.studentId}`, {
                name: values.name,
                nis: values.nis,
                place_birth: values.place_birth,
                date_birth: values.date_birth,
            }, {
                headers:{
                    Authorization: `Bearer ${userSelector.token}`,
                },
            });

            toast.success('Siswa berhasil diubah');
        } catch (error) {
            console.log(error);
        }
    }

    return (
        <UserLayout>
            <div className="flex justify-between items-center gap-4 mb-5">
                <h3>Edit Siswa</h3>
                <div>
                    <Link to="/user/student" className="btn btn-warning">
                        Kembali
                    </Link>
                </div>
            </div>
            <div className="card">
                <div className="card-body">
                    {student.id ?(
                        <StudentForm 
                            onSubmit={handleEditStudent}
                            defaultName={student.name}
                            defaultNis={student.nis}
                            defaultPlaceBirth={student.place_birth}
                            defaultDateBirth={student.date_birth}
                        />
                    ) : (
                        <div>Loading...</div>
                    )}
                </div>
            </div>
        </UserLayout>
    )
}