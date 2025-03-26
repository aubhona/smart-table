import React, { useState } from 'react';
import './RoomCode.css';

function RoomCode({ onSubmit, error }) {
  const [roomCode, setRoomCode] = useState('');

  const handleChange = (event) => {
    setRoomCode(event.target.value);
  };

  const handleSubmit = () => {
    if (roomCode.length === 4) {
      onSubmit(roomCode);
    } else {
      alert('Неверный код комнаты');
    }
  };

  return (
    <div className="room-code-container">
      <h2>Введите код комнаты</h2>
      <input
        type="text"
        value={roomCode}
        onChange={handleChange}
        placeholder="4 цифры"
      />
      <button onClick={handleSubmit}>Подтвердить</button>
      {error && <p className="room-code-error">{error}</p>}
    </div>
  );
}

export default RoomCode;
