import UserLayout from "@Layout/User";
import { useEffect, useState } from "react";
import { useSelector } from "react-redux";
import { axiosInstance } from "../../lib/axios";
import { formatDateToString } from "../../helpers/dateHelper";
import toast from "react-hot-toast";
import Loading from "../../components/loading";

export default function Dashboard() {
  const userSelector = useSelector((state) => state.user);
  const [loading, setLoading] = useState(false);

  const [classrooms, setClassrooms] = useState([]);
  const [students, setStudents] = useState([]);
  const [totalClassrooms, setTotalClassrooms] = useState(0);
  const [totalStudents, setTotalStudents] = useState(0);

  const fetchData = async () => {
    setLoading(true);

    try {
      const response = await axiosInstance.get("/users/dashboard", {
        headers: {
          Authorization: `Bearer ${userSelector.token}`,
        },
      });

      setClassrooms(response.data.data.classrooms);
      setStudents(response.data.data.students);
      setTotalClassrooms(response.data.data.total_classrooms);
      setTotalStudents(response.data.data.total_students);
    } catch (error) {
      console.error(error);
      toast.error('Gagal mengambil data dashboard');
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    fetchData();
  }, [userSelector.token]);

  return (
    <UserLayout>
      <div className="grid grid-cols-2 items-center gap-4 mb-5">
        <div className="card">
          <div className="card-body">
            <span className="text-base">Total Kelas</span>
            <h3>
              {totalClassrooms}
            </h3>
          </div>
        </div>
        <div className="card">
          <div className="card-body">
            <span className="text-base">Total Siswa</span>
            <h3>
              {totalStudents}
            </h3>
          </div>
        </div>
      </div>

      <div className="card mb-5">
        <div className="card-header">
          <h5 className="card-title">Daftar Siswa</h5>
        </div>
        <div className="card-body">
          <div className="table-responsive">
            <table className="table">
              <thead>
                <tr>
                  <th scope="col">
                    No
                  </th>
                  <th scope="col">
                    Nama
                  </th>
                  <th scope="col">
                    NIS
                  </th>
                  <th scope="col">
                    TTL
                  </th>
                </tr>
              </thead>
              <tbody>
                {loading ? (
                  <tr>
                    <td colSpan="4" className="text-center">
                      <Loading />
                    </td>
                  </tr>
                ) : (
                  students.length > 0 ? students.map((student, index) => (
                    <tr key={student.id}>
                      <th scope="row">
                        {index + 1}
                      </th>
                      <td>
                        {student.name}
                      </td>
                      <td>
                        {student.nis}
                      </td>
                      <td>
                        {student.place_birth}, {formatDateToString(student.date_birth)}
                      </td>
                    </tr>
                  )) : (
                    <tr>
                      <td colSpan="4" className="text-center">Tidak ada data</td>
                    </tr>
                  )
                )}
              </tbody>
            </table>
          </div>

        </div>
      </div>

      <div className="card mb-5">
        <div className="card-header">
          <h5 className="card-title">Daftar Kelas</h5>
        </div>
        <div className="card-body">
          <div className="table-responsive">
            <table className="table">
              <thead>
                <tr>
                  <th scope="col">
                    No
                  </th>
                  <th scope="col">
                    Nama
                  </th>
                </tr>
              </thead>
              <tbody>
                { loading ? (
                  <tr>
                    <td colSpan="2" className="text-center">
                      <Loading />
                    </td>
                  </tr>
                ) : (
                  classrooms.length > 0 ? classrooms.map((classroom, index) => (
                    <tr key={classroom.id}>
                      <th scope="row">
                        {index + 1}
                      </th>
                      <td>
                        {classroom.name}
                      </td>
                    </tr>
                  )) : (
                    <tr>
                      <td colSpan="2" className="text-center">Tidak ada data</td>
                    </tr>
                  )
                )}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </UserLayout>
  );
}
