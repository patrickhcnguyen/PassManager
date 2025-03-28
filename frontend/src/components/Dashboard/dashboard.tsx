import { useFetchUser } from '../../Hooks/fetchUser/fetchUser';

const Dashboard = () => {
    const { data: userData } = useFetchUser();

    return (
        <div className="min-h-screen flex flex-col items-center justify-center space-y-4">
            {/* make a call to the backend to decode jwt token and grab user data */}
            <h1 className="text-2xl font-bold">Welcome to your Password Manager Dashboard!</h1>
            <p>Username: {userData?.username}</p>
            <p>Email: {userData?.email}</p>
        </div>
    );
};

export default Dashboard; 