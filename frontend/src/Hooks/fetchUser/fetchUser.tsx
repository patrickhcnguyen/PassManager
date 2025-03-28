import {
    useMutation,
    useQuery,
} from '@tanstack/react-query';


interface RegisterData {
    email: string;
    username: string;
    master_password: string;
}

interface LoginData {
    login: string;
    master_password: string;
}

// register user
export const useRegisterUser = () => {
    return useMutation({
        mutationFn: async (userData: RegisterData) => {
            const response = await fetch('http://localhost:8080/api/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(userData)
            });

            if (!response.ok) {
                const error = await response.json();
                throw new Error(error.error);
            }

            return response.json();
        }
    });
};

// login user
export const useLoginUser = () => {
    return useMutation({
        mutationFn: async (userData: LoginData) => {
            try {
                const response = await fetch('http://localhost:8080/api/login', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    credentials: 'include',
                    body: JSON.stringify(userData)
                });

                const data = await response.json();
                
                if (!response.ok) {
                    console.log('Login response:', response.status, data);
                    throw new Error(data.error || 'Login failed');
                }

                return data;
            } catch (error) {
                console.error('Login error:', error);
                throw error;
            }
        }
    });
};

// fetch user data
export const useFetchUser = () => {
    return useQuery({
        queryKey: ['user'],
        queryFn: async () => {
            const response = await fetch('http://localhost:8080/api/user', {
                method: 'GET',
                credentials: 'include',
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json',
                }
            });
            
            if (!response.ok) {
                throw new Error('Failed to fetch user data');
            }
            
            return response.json();
        },
        retry: false,
        staleTime: Infinity,
    });
};
