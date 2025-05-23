import React, { useState } from 'react';
import './RoomCode.css';

function RoomCode({ onSubmit, error }) {
  const [roomCode, setRoomCode] = useState('');
  const [localError, setLocalError] = useState('');

  const handleChange = (event) => {
    setRoomCode(event.target.value);
    setLocalError('');
  };

  const handleSubmit = () => {
    if (!roomCode) {
      setLocalError('Введите код комнаты');
      return;
    }
    onSubmit(roomCode);
  };

  return (
    <div className="room-code-container">
      <h2>Введите код комнаты</h2>
      <input
        type="text"
        value={roomCode}
        onChange={handleChange}
        placeholder="Код"
        autoFocus
      />
      <button onClick={handleSubmit}>Подтвердить</button>
      {(error || localError) && (
        <p className="room-code-error">{error || localError}</p>
      )}
    </div>
  );
}

export default RoomCode;
