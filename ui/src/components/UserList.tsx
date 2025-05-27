import React, { useEffect, useState } from "react";
import AchievementModal from "./AchievementModal";
import defaultProfile from "../assets/profilepic/defaultprofile.png";
import { loadProfilePics } from "../utils/loadProfilePics";

const profilePics = loadProfilePics();

interface User {
  id: number;
  name: string;
  email: string;
  profilePic?: string;
}

const UserList: React.FC = () => {
  const [users, setUsers] = useState<User[]>([]);
  const [selectedUserId, setSelectedUserId] = useState<number | null>(null);

  useEffect(() => {
    fetch("http://localhost:5000/users")
      .then((res) => res.json())
      .then((data) => {
        const usersWithPics = data.map((user: User) => ({
          ...user,
          profilePic:
            profilePics[Math.floor(Math.random() * profilePics.length)],
        }));
        setUsers(usersWithPics);
      })
      .catch((err) => console.error("Error fetching users:", err));
  }, []);

  return (
    <>
      <div className="row">
        {users.map((user) => (
          <div className="col-md-4 mb-4" key={user.id}>
            <div
              className="card p-3 h-100"
              style={{ cursor: "pointer" }}
              onClick={() => setSelectedUserId(user.id)}
            >
              <div className="d-flex align-items-center gap-3">
                <img
                  src={user.profilePic || defaultProfile}
                  alt="Profile"
                  className="rounded-circle border border-light shadow"
                  style={{
                    width: "70px",
                    height: "70px",
                    objectFit: "cover",
                    padding: "2px",
                    backgroundColor: "#1e1e1e", // dark background inside circle
                  }}
                />
                <div>
                  <h5 className="mb-1">{user.name}</h5>
                  <p className="mb-0 text-light" style={{ fontSize: "0.9rem" }}>
                    {user.email}
                  </p>
                </div>
              </div>
            </div>
          </div>
        ))}
      </div>

      {selectedUserId !== null && (
        <AchievementModal
          userId={selectedUserId}
          onClose={() => setSelectedUserId(null)}
        />
      )}
    </>
  );
};

export default UserList;
