import { zodResolver } from '@hookform/resolvers/zod';
import DreamerSvg from '@PublicAssets/images/dreamer.svg';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { axiosInstance } from '../../lib/axios';
import { useDispatch } from 'react-redux';
import { jwtDecode } from 'jwt-decode';
import Cookies from 'js-cookie';

const loginFormShcema = z.object({
    email: z
        .string()
        .min(3, "email must be at least 3 characters or more")
        .email('Invalid email address'),
    password: z.string().min(6, "password must be at least 6 characters or more"),
});

export default function Login() {
    
    const form = useForm({
        defaultValues: {
            email: '',
            password: '',
            message: '',
        },
        resolver: zodResolver(loginFormShcema),
        reValidateMode: 'onSubmit',
    })

    const dispatch = useDispatch();

    const handleLogin = async (values) => {
        try {
            const userResponse = await axiosInstance.post("/users/login", {
                email: values.email,
                password: values.password,
            });

            // decode token user
            let decodedToken = jwtDecode(userResponse.data.token);

            dispatch({
                type: 'LOGIN',
                payload: {
                    name: decodedToken.name,
                    email: decodedToken.email,
                    token: userResponse.data.token,
                }
            });

            // simpan data user dan token ke cookie
            Cookies.set('user', JSON.stringify({ name: decodedToken.name, email: decodedToken.email }), { expires: 7 });
            Cookies.set('token', userResponse.data.token, { expires: 7 });

            form.reset();
        } catch (error) {
            console.log(error);

            form.setError('message', {
                type: 'manual',
                message: error.response.data.err_message,
            });
        }
    }

    return (
        <div className="bg-indigo-950">
            <div className="flex justify-center items-center min-h-screen">
                <div className="card !w-1/2">
                    <div className="card-body grid grid-cols-1 lg:grid-cols-2 items-center gap-10">
                        <div className='hidden lg:block'>
                            <img src={DreamerSvg} alt="" />
                        </div>
                        <div>
                            <h3 className='mb-5'>Login</h3>
                            <form 
                                className="space-y-4"
                                onSubmit={form.handleSubmit(handleLogin)}
                            >
                                {
                                    form.formState.errors.message && (
                                        <div className="p-4 mb-4 text-sm text-red-800 rounded-lg bg-red-50" role="alert">
                                            {form.formState.errors.message.message}
                                        </div>
                                    )
                                }

                                <div>
                                    <label htmlFor="email" className="form-label">Email</label>
                                    <input type="email" id="email" {...form.register('email')} className="form-input" />
                                    {form.formState.errors.email && <p className="text-red-500">{form.formState.errors.email.message}</p>}
                                </div>
                                <div>
                                    <label htmlFor="password" className="form-label">Password</label>
                                    <input type="password" id="password" {...form.register('password')} className="form-input" />
                                    {form.formState.errors.password && <p className="text-red-500">{form.formState.errors.password.message}</p>}
                                </div>
                                <div>
                                    <button type="submit" className="btn btn-primary w-full">Login</button>
                                </div>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}
