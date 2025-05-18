import React from 'react';
import './LoadingScreen.css';

const LoadingScreen = ({ message }) => {
  return (
    <div className="loading-screen">
      <div className="loading-container">
        <h2 className="loading-title">{message}</h2>
        <div className="loading-spinner"></div>
      </div>
    </div>
  );
};

export default LoadingScreen; 