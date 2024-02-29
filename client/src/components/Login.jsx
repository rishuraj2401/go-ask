// import React from 'react';
import { GoogleAuthProvider, signInWithPopup } from 'firebase/auth';
import { auth } from '../firebase';
import { useContext } from 'react';
import { DataContext } from './DataContext';
const Login = () => {
    const {
        user, 
        setUser
    }= useContext(DataContext)
    const loginHandler = async () => {
        try {
          const provider = new GoogleAuthProvider();
          const { user } = await signInWithPopup(auth, provider);
          setUser(user)
          console.log({
            name: user.displayName,
            email: user.email,
            _id: user.uid,
          });}
          catch{
              console.log("error");
          }}

    return (
        <div>
            <button  type='button' onClick={loginHandler}>button</button>
            
        </div>
    );
}

export default Login;
