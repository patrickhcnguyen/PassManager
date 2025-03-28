import { useState } from 'react';
import { useFetchUser } from '../../Hooks/fetchUser/fetchUser';
import { useCreatePassword, useFetchPasswords } from '../../Hooks/usePasswords/usePasswords';
// show unencrypted password

const Dashboard = () => {
    const { data: userData } = useFetchUser();
    const { data: passwords, refetch: refetchPasswords } = useFetchPasswords();
    const createPasswordMutation = useCreatePassword();
    
    const [showPassword, setShowPassword] = useState<{[key: string]: boolean}>({});

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const form = e.currentTarget;
        const formData = new FormData(form);
        
        try {
            await createPasswordMutation.mutateAsync({
                website_name: formData.get('website_name') as string,
                username: formData.get('username') as string,
                password: formData.get('password') as string
            });
            
            form.reset();
            refetchPasswords(); 
        } catch (error) {
            console.error('Failed to create password entry:', error);
        }
    };

    const togglePasswordVisibility = (id: string) => {
        setShowPassword(prev => ({
            ...prev,
            [id]: !prev[id]
        }));
    };

    return (
        <div className="min-h-screen p-8 space-y-8">
            <div className="max-w-4xl mx-auto space-y-8">
                <div className="bg-white p-6 rounded-lg shadow">
                    <h1 className="text-2xl font-bold mb-4">Welcome to your Password Manager Dashboard!</h1>
                    <p>Username: {userData?.username}</p>
                    <p>Email: {userData?.email}</p>
                </div>

                {/* New Password Form */}
                <div className="bg-white p-6 rounded-lg shadow">
                    <h2 className="text-xl font-semibold mb-4">Add New Password</h2>
                    <form onSubmit={handleSubmit} className="space-y-4">
                        <input
                            type="text"
                            name="website_name"
                            placeholder="Website Name"
                            required
                            className="block w-full p-2 border rounded"
                        />
                        <input
                            type="text"
                            name="username"
                            placeholder="Username/Email for site"
                            required
                            className="block w-full p-2 border rounded"
                        />
                        <input
                            type="password"
                            name="password"
                            placeholder="Password"
                            required
                            className="block w-full p-2 border rounded"
                        />
                        <button 
                            type="submit"
                            disabled={createPasswordMutation.isPending}
                            className="w-full bg-blue-500 text-white p-2 rounded hover:bg-blue-600 disabled:bg-blue-300"
                        >
                            {createPasswordMutation.isPending ? 'Adding...' : 'Add Password'}
                        </button>
                    </form>
                </div>

                {/* Passwords List */}
                <div className="bg-white p-6 rounded-lg shadow">
                    <h2 className="text-xl font-semibold mb-4">Saved Passwords</h2>
                    <div className="space-y-4">
                        {passwords?.map((entry: any) => (
                            <div key={entry.ID} className="border p-4 rounded">
                                <p className="font-semibold">{entry.SiteName}</p>
                                <p>Username: {entry.SiteUsername}</p>
                                <div className="flex items-center gap-2">
                                    <p>
                                        Password: {showPassword[entry.ID] 
                                            ? entry.EncryptedPassword 
                                            : '••••••••'}
                                    </p>
                                    <button
                                        onClick={() => togglePasswordVisibility(entry.ID)}
                                        className="text-sm text-blue-500 hover:text-blue-700"
                                    >
                                        {showPassword[entry.ID] ? 'Hide' : 'Show'}
                                    </button>
                                </div>
                            </div>
                        ))}
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Dashboard; 