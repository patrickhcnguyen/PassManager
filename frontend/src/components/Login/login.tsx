import React from 'react';
import { useLoginUser } from '../../Hooks/fetchUser/fetchUser';
import { useNavigate } from 'react-router-dom';

const Login = () => {
    const loginMutation = useLoginUser();
    const navigate = useNavigate();

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const form = e.currentTarget;
        const formData = new FormData(form);
        
        try {
            const result = await loginMutation.mutateAsync({
                login: formData.get('login') as string,
                master_password: formData.get('master_password') as string
            });
            
            if (result.token) {
                alert('Successfully logged in');
                form.reset();
                navigate('/dashboard');
            }
        } catch (error) {
            alert('Login failed. Please check your credentials.');
        }
    }

    return (
        <div className="min-h-screen flex items-center justify-center">
            <form onSubmit={handleSubmit} className="space-y-4">
                <input
                    type="text"
                    name="login"
                    placeholder="Email or Username"
                    className="block w-full p-2 border rounded"
                />
                <input
                    type="password"
                    name="master_password"
                    placeholder="Master Password"
                    className="block w-full p-2 border rounded"
                />
                <button 
                    type="submit"
                    disabled={loginMutation.isPending}
                    className="w-full bg-blue-500 text-white p-2 rounded hover:bg-blue-600 disabled:bg-blue-300"
                >
                    {loginMutation.isPending ? 'Logging in...' : 'Login'}
                </button>
            </form>
        </div>
    );
};

export default Login;
