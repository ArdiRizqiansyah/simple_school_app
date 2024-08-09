import UserLayout from "@Layout/User";
import { Link, useParams } from "react-router-dom";
import { axiosInstance } from "../../../lib/axios";
import toast from "react-hot-toast";
import { useSelector } from "react-redux";
import { ClassroomForm } from "../../../components/navbar/forms/ClassroomForm";
import { useEffect, useState } from "react";


export default function ClassroomEdit() {
    const userSelector = useSelector((state) => state.user);

    const [classroom, setClassroom] = useState({
        name: "",
    });
    const params = useParams();

    const fetchClassroom = async () => {
        try {
            const response = await axiosInstance.get(`/classrooms/${params.classroomId}`, {
                headers: {
                    Authorization: `Bearer ${userSelector.token}`,
                },
            });

            setClassroom(response.data);
        } catch (error) {
            console.error(error);
            toast.error('Gagal mengubah data kelas');
        }
    }

    useEffect(() => {
        fetchClassroom();
    }, []);

    const handleEditClassroom = async (values) => {
        try {
            await axiosInstance.put(`/classrooms/${params.classroomId}`, {
                name: values.name,
            }, {
                headers:{
                    Authorization: `Bearer ${userSelector.token}`,
                },
            });

            toast.success('Kelas berhasil diubah');
        } catch (error) {
            console.log(error);
        }
    }

    return (
        <UserLayout>
            <div className="flex justify-between items-center gap-4 mb-5">
                <h3>Edit Kelas</h3>
                <div>
                    <Link to="/user/classroom" className="btn btn-warning">
                        Kembali
                    </Link>
                </div>
            </div>
            <div className="card">
                <div className="card-body">
                    {classroom.id ?(
                        <ClassroomForm 
                            onSubmit={handleEditClassroom}
                            defaultName={classroom.name}
                        />
                    ) : null }
                </div>
            </div>
        </UserLayout>
    )
}
