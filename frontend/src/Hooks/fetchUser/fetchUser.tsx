import {
    useMutation,
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

export const useLoginUser = () => {
    return useMutation({
        mutationFn: async (userData: LoginData) => {
            const response = await fetch('http://localhost:8080/api/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(userData)
            })

            if (!response.ok) {
                const error = await response.json();
                throw new Error(error.error);
            }

            if (response.ok) {
                console.log('login successful');
            }
            const data = await response.json();
            console.log(data);

            return data;
        }
    })
}

// export const useAuthenticatedQuery = () => {
//     const getAuthHeader = () => ({
//         'Authorization': `Bearer ${localStorage.getItem('token')}`,
//         'Content-Type': 'application/json'
//     });
    
//     return { getAuthHeader };
// };

