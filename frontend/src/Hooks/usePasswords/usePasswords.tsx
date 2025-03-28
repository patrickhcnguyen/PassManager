import { useMutation, useQuery } from '@tanstack/react-query';

interface CreatePasswordData {
    website_name: string;
    username: string;
    password: string;
}

interface PasswordEntry extends CreatePasswordData {
    ID: number;
}

export const useCreatePassword = () => {
    return useMutation({
        mutationFn: async (passwordData: CreatePasswordData) => {
            const response = await fetch('http://localhost:8080/api/passwords/create', {
                method: 'POST',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(passwordData)
            });

            if (!response.ok) {
                throw new Error('Failed to create password entry');
            }

            return response.json();
        }
    });
};

export const useFetchPasswords = () => {
    return useQuery<PasswordEntry[]>({
        queryKey: ['passwords'],
        queryFn: async () => {
            const response = await fetch('http://localhost:8080/api/passwords', {
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json',
                }
            });

            if (!response.ok) {
                throw new Error('Failed to fetch passwords');
            }

            return response.json();
        }
    });
}; 