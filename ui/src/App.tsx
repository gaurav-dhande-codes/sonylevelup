import { useEffect, useState } from "react";
import { UserList } from "./components/UserList";
import type { User } from "./types/user";

const App = () => {
  const [users, setUsers] = useState<User[]>([]);

  useEffect(() => {
    const mockUsers: User[] = [
      { id: 1, name: "Aiko Tanaka" },
      { id: 2, name: "Liam Chen" },
      { id: 3, name: "Sofia Hernandez" },
      { id: 4, name: "Aiko Tanaka" },
      { id: 5, name: "Liam Chen" },
      { id: 6, name: "Sofia Hernandez" },
    ];
    setTimeout(() => setUsers(mockUsers), 500);
  }, []);

  return (
    <div className="bg-dark text-light min-vh-100">
      <div className="container d-flex justify-content-center pt-5">
        <UserList users={users} />
      </div>
    </div>
  );
};

export default App;
