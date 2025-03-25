import React, { useState } from 'react';
import { useRegisterUser } from '../../Hooks/fetchUser/fetchUser';

const Register = () => {
  const [formData, setFormData] = useState({
    Email: '',
    Username: '',
    MasterPassword: ''  
  });

  const registerMutation = useRegisterUser();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await registerMutation.mutateAsync(formData);
      console.log('Registration successful!');
    } catch (error) {
      console.error('Registration failed:', error);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center">
      <form onSubmit={handleSubmit} className="space-y-4">
        <input
          type="text"
          placeholder="Username"
          value={formData.Username}
          onChange={e => setFormData({...formData, Username: e.target.value})}
          className="block w-full p-2 border rounded"
        />
        <input
          type="email"
          placeholder="Email"
          value={formData.Email}
          onChange={e => setFormData({...formData, Email: e.target.value})}
          className="block w-full p-2 border rounded"
        />
        <input
          type="password"
          placeholder="Master Password"
          value={formData.MasterPassword}
          onChange={e => setFormData({...formData, MasterPassword: e.target.value})}
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