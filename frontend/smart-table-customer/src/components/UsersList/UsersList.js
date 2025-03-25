import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import "./UsersList.css";

function UsersList() {
  const navigate = useNavigate();
  const [roomCode] = useState("");
  const [users] = useState([
    { id: 1, name: "User 1" },
    { id: 2, name: "User 2" },
    { id: 3, name: "User 3" },
    { id: 4, name: "User 4" },
    { id: 5, name: "User 5" },
    { id: 6, name: "User 6" },
    { id: 7, name: "User 7" },
    { id: 8, name: "User 8" },
    { id: 9, name: "User 9" },
    { id: 10, name: "User 10" },
    { id: 11, name: "User 11" },
    { id: 12, name: "User 12" },

  ]);

  const goBack = () => {
    navigate("/catalog");
  };

  return (
    <div className="users-container">
      <div className="users-header">
        <div className="room-code">1234</div>
      </div>

      <h2>Список пользователей</h2>

      <div className="users-list">
        
        {users.map((user) => (
          <div key={user.id} className="user-item">
            <div className="user-picture">Фотка</div>
            <div className="user-name">{user.name}</div>
          </div>
        ))}
      </div>

      <div className="bottom-section">
        <button className="go-back-button" onClick={goBack} >
          Назад
        </button>
      </div>
    </div>
  );
}

export default UsersList;
