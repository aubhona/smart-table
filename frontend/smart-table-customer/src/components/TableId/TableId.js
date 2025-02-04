import React, { useState } from 'react';
import './TableId.css';

function TableId() {
  const [tableId, setTableId] = useState('');
  const [error, setError] = useState('');

  const validateTableId = (id) => {
    const isValid = /^\d{6}$/.test(id);
    setError(isValid ? '' : 'Неверный номер стола');
    return isValid;
  };

  const handleChange = (event) => {
    setTableId(event.target.value);
    setError('');
  };

  const handleSubmit = () => {
    if (validateTableId(tableId)) {
      alert(`Проверка успешна! Table ID: ${tableId}`);
      window.location.href = '/room-code';
    }
  };

  return (
    <div className="table-id-container">
      <h2>Введите номер стола</h2>
      <input
        type="text"
        placeholder="печатайте..."
        value={tableId}
        onChange={handleChange}
      />
      <button onClick={handleSubmit}>Подтвердить</button>
      {error && <p className="table-id-error">{error}</p>}
    </div>
  );
}

export default TableId;
