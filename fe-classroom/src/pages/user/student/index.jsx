import UserLayout from "@Layout/User";
import { useEffect, useState } from "react";
import { useSelector } from "react-redux";
import { Link, useSearchParams } from "react-router-dom";
import { axiosInstance } from "../../../lib/axios";
import { formatDateToString } from "../../../helpers/dateHelper";
import toast from "react-hot-toast";
import Pagination from "../../../components/pagination";

export default function Classroom() {
    const userSelector = useSelector((state) => state.user);
    const [searchParams] = useSearchParams();
    const page = searchParams.get('page') || 1;

    const [loading, setLoading] = useState(false);
    const [pagination, setPagination] = useState([]);

    const [students, setStudents] = useState([]);

    const fetchStudents = async () => {
        setLoading(true);

        try {
            const response = await axiosInstance.get('/students', {
                headers: {
                    Authorization: `Bearer ${userSelector.token}`
                },
                params: {
                    page: page
                }
            });

            setStudents(response.data.data || []);
            setPagination(response.data.page);
        } catch (error) {
            console.error(error);
            toast.error('Gagal mengambil data siswa');
        } finally {
            setLoading(false);
        }
    }

    const handleDeleteStudent = async (studentId) => {
        const shouldDelete = window.confirm('Apakah anda yakin ingin menghapus siswa ini?');

        if (!shouldDelete) {
            return;
        }

        try {
            await axiosInstance.delete(`/students/${studentId}`, {
                headers: {
                    Authorization: `Bearer ${userSelector.token}`
                }
            });

            toast.success('Siswa berhasil dihapus');

            fetchStudents();
        } catch (error) {
            console.error(error);
        }
    }

    useEffect(() => {
        fetchStudents()
    }, [page]);

    return (
        <UserLayout>
            <div className="flex justify-between items-center gap-4 mb-5">
                <h3>Daftar Siswa</h3>
                <div>
                    <Link to="/user/student/create" className="btn btn-primary">
                        Tambah Siswa
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
                                <th>NIS</th>
                                <th>TTL</th>
                                <th>Aksi</th>
                            </tr>
                        </thead>
                        <tbody>
                            { loading ? (
                                <tr>
                                    <td colSpan="5" className="text-center">Loading...</td>
                                </tr>
                            ) : (
                                students.map((student, index) => (
                                    <tr key={student.id}>
                                        <th>
                                            {pagination.from + index + 1}
                                        </th>
                                        <td>{student.name}</td>
                                        <td>{student.nis}</td>
                                        {/* format date_birth student */}
                                        <td>{`${student.place_birth}, ${formatDateToString(student.date_birth)}`}</td>
                                        <td>
                                            <div className="flex gap-2">
                                                <Link to={`/user/student/${student.id}/edit`} className="btn btn-sm btn-success">Edit</Link>
                                                <button 
                                                    type="button"
                                                    className="btn btn-sm btn-danger"
                                                    onClick={() => handleDeleteStudent(student.id)}
                                                >
                                                    Hapus
                                                </button>
                                            </div>
                                        </td>
                                    </tr>
                                ))
                            )}
                        </tbody>
                    </table>
                </div>
            </div>

            { students && students.length != 0 &&
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
