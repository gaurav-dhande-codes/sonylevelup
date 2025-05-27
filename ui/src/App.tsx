import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import playstationLogo from "./assets/playstation.png";
import UserList from "./components/UserList";

const App: React.FC = () => {
  return (
    <Router>
      <div className="container py-4">
        <div className="header d-flex align-items-center gap-3">
          <img
            src={playstationLogo}
            alt="PlayStation Logo"
            style={{ height: "100px", objectFit: "contain" }}
          />
          <div>
            <h1>Sony Level Up Achievement Dashboard</h1>
            <p className="subtitle">Track users' achievement level</p>
          </div>
        </div>

        <div className="container mt-4">
          <Routes>
            <Route path="/" element={<UserList />} />
          </Routes>
        </div>
      </div>
    </Router>
  );
};

export default App;
