import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";

const classroomFormSchema = z.object({
    name: z
        .string()
        .min(3, "Nama kelas minimal 3 karakter"),
});

export const ClassroomForm = (props) => {
    const {
        onSubmit,
        formRef,
        defaultName,
    } = props;

    const form = useForm({
        defaultValues: {
            name: defaultName || "",
        },
        resolver: zodResolver(classroomFormSchema),
    });

    return (
        <form 
            onSubmit={form.handleSubmit(onSubmit)}
            ref={formRef}
        >
            <div className="mb-5">
                <label htmlFor="name" className="form-label">Nama Kelas</label>
                <input type="text" id="name" className="form-input" {...form.register('name')} />
                {form.formState.errors.name && <span className="text-red-500">{form.formState.errors.name.message}</span>}
            </div>
            <div className="grid">
                <button type="submit" className="btn btn-primary">Simpan</button>
            </div>
        </form>
    );
}