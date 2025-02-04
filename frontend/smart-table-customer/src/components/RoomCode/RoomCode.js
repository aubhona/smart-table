import React, { useState } from 'react';
import './RoomCode.css';

function RoomCode() {
  const [roomCode, setRoomCode] = useState('');
  const [error, setError] = useState('');

  const validateRoomCode = (id) => {
    const isValid = /^\d{4}$/.test(id);
    setError(isValid ? '' : 'Неверный код комнаты');
    return isValid;
  };

  const handleChange = (event) => {
    setRoomCode(event.target.value);
    setError('');
  };

  const handleSubmit = () => {
    if (validateRoomCode(roomCode)) {
      alert(`Проверка успешна! Room code: ${roomCode}`);
      window.location.href = '/catalog';
    }
  };

  return (
    <div className="room-code-container">
      <h2>Введите код комнаты</h2>
      <input    
        type="text"
        placeholder="печатайте..."
        value={roomCode}
        onChange={handleChange}
      />
      <button onClick={handleSubmit}>Подтвердить</button>
      {error && <p className="room-code-error">{error}</p>}
    </div>
  );
}

export default RoomCode;