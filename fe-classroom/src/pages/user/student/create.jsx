import UserLayout from "@Layout/User";
import { useSelector } from "react-redux";
import { Link } from "react-router-dom";
import { axiosInstance } from "../../../lib/axios";
import toast from "react-hot-toast";
import { StudentForm } from "../../../components/navbar/forms/StudentForm";
import { useRef } from "react";


export default function StudentCreate() {
    const userSelector = useSelector((state) => state.user);
    const formRef = useRef();

    const handleCreateStudent = async (values) => {
        try {
            await axiosInstance.post("/students", {
                name: values.name,
                nis: values.nis,
                place_birth: values.place_birth,
                date_birth: values.date_birth,
            }, {
                headers:{
                    Authorization: `Bearer ${userSelector.token}`,
                },
            });

            // reset form
            formRef.current.reset();

            toast.success('Siswa berhasil ditambahkan');
        } catch (error) {
            console.log(error);
        }
    }

    return (
        <UserLayout>
            <div className="flex justify-between items-center gap-4 mb-5">
                <h3>Tambah Siswa</h3>
                <div>
                    <Link to="/user/student" className="btn btn-warning">
                        Kembali
                    </Link>
                </div>
            </div>
            <div className="card">
                <div className="card-body">
                    <StudentForm
                        onSubmit={handleCreateStudent}
                        formRef={formRef}
                    />
                </div>
            </div>
        </UserLayout>
    )
}
