import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import NavBar from "./components/NavBar/NavBar";
import AuthPage from "./pages/AuthPage";
import HomePage from "./pages/HomePage";

function App() {
    return (
        <Router>
            <NavBar />
            <Routes>
                <Route exact path="/" element={<HomePage />} />
                <Route exact path="/auth" element={<AuthPage />} />
            </Routes>
        </Router>
    );
}

export default App;
