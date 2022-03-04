import React from 'react';


import MovieProfile from './components/MovieProfile';
import Home from './components/Home';
import Login from './components/Login';
import SignUp from './components/SignUp';
import UserProfile from './components/UserProfile';
import { Add } from './components/Add';
import {Route,Routes} from 'react-router-dom';
import MovieDetail from './components/MovieDetail';


function App() {
  return (
    <div className="App">
      <Routes>
      <Route path="/" element={<Home/>} />
      <Route path="/MovieProfile" element={<MovieProfile/>} />
      <Route path="/MovieDetail" element={<MovieDetail/>} />
      <Route path="/UserProfile" element={<UserProfile/>} />
      <Route path="/UserProfile/add" element={<Add/>} />
      <Route path="/login" element={<Login/>} />
      <Route path="/signup" element={<SignUp/>} />
      </Routes>
    </div>
      
  );
}

export default App;
