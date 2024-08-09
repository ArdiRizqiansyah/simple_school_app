import UserLayout from "@Layout/User";
import { useEffect, useState } from "react";
import { useSelector } from "react-redux";
import Select from 'react-select';
import { axiosInstance } from "../../../../lib/axios";
import { Link, useParams } from "react-router-dom";
import { formatDateToString } from "../../../../helpers/dateHelper";
import toast from "react-hot-toast";
import { z } from "zod";
import { Controller, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import Loading from "../../../../components/loading";

const classroomStudentFormSchema = z.object({
    student_id: z.string().optional().array()
});

export default function Student() {
    const userSelector = useSelector((state) => state.user);
    const params = useParams();
    const [loading, setLoading] = useState(false);

    const [classroomStudents, setClassroomStudents] = useState([]);
    const [students, setStudents] = useState([]);

    const fetchStudents = async () => {
        setLoading(true);

        try {
            const response = await axiosInstance.get(`/classrooms/${params.classroomId}/students`, {
                headers: {
                    Authorization: `Bearer ${userSelector.token}`
                }
            });

            setClassroomStudents(response.data.data.classroom_students);
            setStudents(response.data.data.students ?? []);
        } catch (error) {
            console.error(error);
            toast.error('Gagal mengambil data siswa');
        } finally {
            setLoading(false);
        }
    }

    // create data
    const formClassroomStudent = useForm({
        defaultValues: {
            student_id: ""
        },
        resolver: zodResolver(classroomStudentFormSchema)
    });

    const handleCreateClassroomStudent = async (values) => {
        try {
            await axiosInstance.post(`/classrooms/${params.classroomId}/students`, {
                student_id: values.student_id.map((student) => student.value)
            }, {
                headers: {
                    Authorization: `Bearer ${userSelector.token}`
                }
            });

            toast.success('Siswa berhasil ditambahkan');
            fetchStudents();
            formClassroomStudent.reset();
        } catch (error) {
            console.error(error);
        }
    }
    // end create data

    // delete data
    const handleDeleteClassroomStudent = async (classroomStudentId) => {
        const shouldDelete = window.confirm('Apakah anda yakin ingin menghapus siswa ini?');

        if (!shouldDelete) {
            return;
        }

        try {
            await axiosInstance.delete(`/classrooms/${params.classroomId}/students/${classroomStudentId}`, {
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
    // end delete data

    useEffect(() => {
        fetchStudents()
    }, []);

    return (
        <UserLayout>
            <div className="flex justify-between items-center gap-4 mb-5">
                <div>
                    <h3 className="mb-1">Daftar Siswa</h3>
                    <p className="text-muted">Siswa terdaftar di kelas 10A</p>
                </div>
                <div>
                    <Link to="/user/classroom" className="btn btn-warning">
                        Kembali
                    </Link>
                </div>
            </div>
            <form onSubmit={formClassroomStudent.handleSubmit(handleCreateClassroomStudent)} className="flex gap-4 mb-4">
                <div className="w-full">
                    <Controller 
                        name="student_id"
                        control={formClassroomStudent.control}
                        render={({ field }) => (
                            <Select 
                                {...field}
                                closeMenuOnSelect={false}
                                isMulti
                                options={
                                    students.length > 0 ? (
                                        students.map((student) => ({
                                            value: student.id,
                                            label: student.name
                                        }))
                                        ) : ([{
                                            value: '',
                                            label: 'Pilih Siswa'
                                        }]
                                    )
                                }
                                className="w-full"
                                placeholder="Pilih Siswa"
                            />
                        )}
                    />
                    {/* <Select 
                        name="student_id"
                        closeMenuOnSelect={false}
                        isMulti
                        options={
                            students && (
                                students.map((student) => ({
                                    value: student.id,
                                    label: student.name
                                }))
                            )
                        }
                        className="w-full"
                        placeholder="Pilih Siswa"
                        control={formClassroomStudent.control}
                    /> */}
                    {/* <select
                        className="form-input"
                        {...formClassroomStudent.register('student_id')}
                        multiple
                    >
                        <option value="">Pilih Siswa</option>
                        { students && (
                            students.map((student) => (
                                <option key={student.id} value={student.id}>{student.name}</option>
                            ))
                        )}
                    </select> */}
                    {formClassroomStudent.formState.errors.student_id && <span className="text-red-500">{formClassroomStudent.formState.errors.student_id.message}</span>}
                </div>
                <button 
                    type="submit"
                    className="btn btn-sm btn-primary text-nowrap"
                >
                    Tambah Siswa
                </button>
            </form>
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
                            {loading ? (
                                <tr>
                                    <td colSpan="5" className="text-center">
                                        <Loading />
                                    </td>
                                </tr>
                            ) : (
                                classroomStudents.length ? (
                                    classroomStudents.map((st, index) => (
                                        <tr key={st.id}>
                                            <th>{index+1}</th>
                                            <td>{st.student.name}</td>
                                            <td>{st.student.nis}</td>
                                            <td>{`${st.student.place_birth}, ${formatDateToString(st.student.date_birth)}`}</td>
                                            <td>
                                                <div className="flex gap-2">
                                                    <button 
                                                        type="button"
                                                        className="btn btn-sm btn-danger"
                                                        onClick={() => handleDeleteClassroomStudent(st.id)}
                                                    >
                                                        Hapus
                                                    </button>
                                                </div>
                                            </td>
                                        </tr>
                                    ))
                                ) : (
                                    <tr>
                                        <td colSpan="5" className="text-center">Tidak ada data</td>
                                    </tr>
                                )
                            )}
                        </tbody>
                    </table>
                </div>
            </div>
        </UserLayout>
    );
}