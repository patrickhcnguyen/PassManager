import React from 'react';

const Register = () => {
  return (
    <div>
        <form>
            <input type="text" placeholder="Username" />
            <input type="email" placeholder="Email" />
            <input type="password" placeholder="Password" />
            <button type="submit">Register</button>
        </form>
    </div>
  );
};

export default Register;