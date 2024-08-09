import UserLayout from "@Layout/User";
import { useEffect, useState } from "react";
import { Link, useSearchParams } from "react-router-dom";
import { axiosInstance } from "../../../lib/axios";
import { useSelector } from "react-redux";
import toast from "react-hot-toast";
import Loading from "../../../components/loading";
import Pagination from "../../../components/pagination";

export default function Classroom() {
    const userSelector = useSelector((state) => state.user);
    const [searchParams] = useSearchParams();
    const page = searchParams.get('page') || 1;

    const [loading, setLoading] = useState(false);
    const [pagination, setPagination] = useState([]);

    const [classrooms, setClassrooms] = useState([]);

    const fetchClassrooms = async () => {
        setLoading(true);

        try {
            const response = await axiosInstance.get("/classrooms", {
                headers: {
                    Authorization: `Bearer ${userSelector.token}`,
                },
                params: {
                    page: page,
                },
            });

            setClassrooms(response.data.data || []);
            setPagination(response.data.page);
        } catch (error) {
            console.error(error);
            toast.error('Gagal mengambil data kelas');
        } finally {
            setLoading(false);
        }
    }

    const handleDeleteClassroom = async (classroomId) => {
        const shouldDelete = window.confirm('Apakah anda yakin ingin menghapus kelas ini?');

        if (!shouldDelete) {
            return;
        }

        try {
            await axiosInstance.delete(`/classrooms/${classroomId}`, {
                headers: {
                    Authorization: `Bearer ${userSelector.token}`,
                },
            });

            toast.success('Kelas berhasil dihapus');
            fetchClassrooms();
        } catch (error) {
            console.error(error);
            toast.error('Gagal menghapus kelas');
        }
    }

    useEffect(() => {
        fetchClassrooms();
    }, [userSelector.token, page]);

    return (
        <UserLayout>
            <div className="flex justify-between items-center gap-4 mb-5">
                <h3>Daftar Kelas</h3>
                <div>
                    <Link to="/user/classroom/create" className="btn btn-primary">
                        Tambah Kelas
                    </Link>
                </div>
            </div>
            <div className="card">
                <div className="table-responsive">
                    <table className="table">
                        <thead>
                            <tr>
                                <th>No</th>
                                <th>Nama</th>
                                <th>Total Siswa</th>
                                <th>Aksi</th>
                            </tr>
                        </thead>
                        <tbody>
                            {loading ? (
                                <tr>
                                    <td colSpan="4">
                                        <Loading />
                                    </td>
                                </tr>
                            ) : (
                                classrooms.length === 0 ? (
                                    <tr>
                                        <td colSpan="4" className="text-center">Tidak ada data</td>
                                    </tr>
                                ) : (
                                    classrooms.map((classroom, index) => (
                                        <tr key={classroom.id}>
                                            <th>
                                                {pagination.from + index + 1}
                                            </th>
                                            <td>{classroom.name}</td>
                                            <td>10</td>
                                            <td>
                                                <div className="flex gap-2">
                                                    <Link 
                                                        to={`/user/classroom/${classroom.id}/student`} 
                                                        className="btn btn-sm btn-secondary"
                                                    >
                                                        Siswa
                                                    </Link>
                                                    <Link 
                                                        to={`/user/classroom/${classroom.id}/edit`} 
                                                        className="btn btn-sm btn-success"
                                                    >
                                                        Edit
                                                    </Link>
                                                    <button
                                                        type="button"
                                                        className="btn btn-sm btn-danger"
                                                        onClick={() => handleDeleteClassroom(classroom.id)}
                                                    >
                                                        Hapus
                                                    </button>
                                                </div>
                                            </td>
                                        </tr>
                                    ))
                                )
                            )}
                        </tbody>
                    </table>
                </div>
            </div>

            { classrooms && classrooms.length != 0 &&
                <div className="my-5">
                    <Pagination
                        currentPage={pagination.current_page}
                        totalPages={pagination.total_page}
                    />
                </div>
            }
        </UserLayout>
    );
}
