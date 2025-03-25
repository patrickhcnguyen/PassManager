import React from 'react';
import { useRegisterUser } from '../../Hooks/fetchUser/fetchUser';

const Register = () => {
  const registerMutation = useRegisterUser();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const form = e.currentTarget;
    const formData = new FormData(form);
    
    try {
      await registerMutation.mutateAsync({
        email: formData.get('email') as string,
        username: formData.get('username') as string,
        master_password: formData.get('master_password') as string
      });
      form.reset();
    } catch (error) {
      console.error('Registration failed:', error);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center">
      <form onSubmit={handleSubmit} className="space-y-4">
        <input
          type="text"
          name="username"
          placeholder="Username"
          className="block w-full p-2 border rounded"
        />
        <input
          type="email"
          name="email"
          placeholder="Email"
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
          disabled={registerMutation.isPending}
          className="w-full bg-blue-500 text-white p-2 rounded hover:bg-blue-600 disabled:bg-blue-300"
        >
          {registerMutation.isPending ? 'Registering...' : 'Register'}
        </button>
        
        {registerMutation.isError && (
          <div className="text-red-500">
            {registerMutation.error.message}
          </div>
        )}
      </form>
    </div>
  );
};

export default Register;