import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { formatDate } from "../../../helpers/dateHelper";

const studentFormSchema = z.object({
    name: z
        .string()
        .min(3, "Nama siswa minimal 3 karakter"),
    nis: z
        .string()
        .min(3, "NIS minimal 3 karakter"),
    place_birth: z
        .string()
        .min(3, "Tempat lahir minimal 3 karakter"),
    date_birth: z.string().date(),
});

export const StudentForm = (props) => {
    const {
        onSubmit,
        formRef,
        defaultName,
        defaultNis,
        defaultPlaceBirth,
        defaultDateBirth,
    } = props;

    const form = useForm({
        defaultValues: {
            name: defaultName || "",
            nis: defaultNis || "",
            place_birth: defaultPlaceBirth || "",
            date_birth: formatDate(defaultDateBirth) || "",
        },
        resolver: zodResolver(studentFormSchema),
    });

    return(
        <form 
            onSubmit={form.handleSubmit(onSubmit)}
            ref={formRef}
        >
            <div className="mb-5">
                <label htmlFor="name" className="form-label">Nama Siswa</label>
                <input type="text" id="name" className="form-input" {...form.register('name')} />
                {form.formState.errors.name && <span className="text-red-500">{form.formState.errors.name.message}</span>}
            </div>
            <div className="mb-5">
                <label htmlFor="nis" className="form-label">NIS</label>
                <input type="text" id="nis" className="form-input" {...form.register('nis')} />
                {form.formState.errors.nis && <span className="text-red-500">{form.formState.errors.nis.message}</span>}
            </div>
            <div className="mb-5">
                <label htmlFor="place_birth" className="form-label">Tempat Lahir</label>
                <input type="text" id="place_birth" className="form-input" {...form.register('place_birth')} />
                {form.formState.errors.place_birth && <span className="text-red-500">{form.formState.errors.place_birth.message}</span>}
            </div>
            <div className="mb-5">
                <label htmlFor="date_birth" className="form-label">Tanggal Lahir</label>
                <input type="date" id="date_birth" className="form-input" {...form.register('date_birth')} />
                {form.formState.errors.date_birth && <span className="text-red-500">{form.formState.errors.date_birth.message}</span>}
            </div>
            <div className="grid">
                <button type="submit" className="btn btn-primary">Simpan</button>
            </div>
        </form>
    );
}