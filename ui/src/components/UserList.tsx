import type { User } from "../types/user";

interface Props {
  users: User[];
}

export const UserList = ({ users }: Props) => {
  return (
    <div className="card bg-secondary text-light p-4 w-75">
      <h2 className="card-title text-center mb-3">User List</h2>
      <ul className="p-0 m-0">
        {users.map((user) => (
          <li
            key={user.id}
            className="bg-dark text-light py-2 px-3 mb-2 rounded text-left user-list-item"
            style={{ listStyle: "none", cursor: "pointer", transition: "0.2s" }}
          >
            {user.name}
          </li>
        ))}
      </ul>
    </div>
  );
};
