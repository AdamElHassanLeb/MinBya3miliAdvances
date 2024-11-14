import React from 'react';
import {Link, Route, Router, Routes} from "react-router-dom";
import Login from "./components/Login";

function App() {
    return (
        <div className="App">
            <h1>Hello World</h1>
            <Link to={'/login'}>Login</Link>
            <Routes>
                <Route path="/login" element={<Login/>} />
            </Routes>
        </div>
    )
}

export default App;