import {
    useMutation,
} from '@tanstack/react-query';


interface UserData {
    email: string;
    username: string;
    master_password: string;
}

export const useRegisterUser = () => {
    return useMutation({
        mutationFn: async (userData: UserData) => {
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

