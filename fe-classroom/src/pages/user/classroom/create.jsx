import UserLayout from "@Layout/User";
import { Link } from "react-router-dom";
import { axiosInstance } from "../../../lib/axios";
import toast from "react-hot-toast";
import { useSelector } from "react-redux";
import { ClassroomForm } from "../../../components/navbar/forms/ClassroomForm";
import { useRef } from "react";


export default function ClassroomCreate() {
    const userSelector = useSelector((state) => state.user);
    const formRef = useRef();

    const handleCreateClassroom = async (values) => {
        try {
            await axiosInstance.post("/classrooms", {
                name: values.name,
            }, {
                headers:{
                    Authorization: `Bearer ${userSelector.token}`,
                },
            });

            // reset form
            formRef.current.reset();

            toast.success('Kelas berhasil ditambahkan');
        } catch (error) {
            console.log(error);
            toast.error('Gagal menambahkan kelas');
        }
    }

    return (
        <UserLayout>
            <div className="flex justify-between items-center gap-4 mb-5">
                <h3>Tambah Kelas</h3>
                <div>
                    <Link to="/user/classroom" className="btn btn-warning">
                        Kembali
                    </Link>
                </div>
            </div>
            <div className="card">
                <div className="card-body">
                    <ClassroomForm 
                        onSubmit={handleCreateClassroom}
                        formRef={formRef}
                    />
                </div>
            </div>
        </UserLayout>
    )
}
