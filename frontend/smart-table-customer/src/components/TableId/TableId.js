import React, { useState } from 'react';
import './TableId.css';

function TableId({ onSubmit, error }) {
  const [tableId, setTableId] = useState('');

  const handleChange = (event) => {
    setTableId(event.target.value);
  };

  const handleSubmit = () => {
    if (tableId.length === 6) {
      onSubmit(tableId);
    } else {
      alert('Неверный ID стола');
    }
  };

  return (
    <div className="table-id-container"> 
      <h2>Введите ID стола</h2>
      <input
        type="text"
        value={tableId}
        onChange={handleChange}
        placeholder="ID стола"
      />
      <button onClick={handleSubmit}>Подтвердить</button>
      {error && <p className="table-id-error">{error}</p>}
    </div>
  );
}

export default TableId;

