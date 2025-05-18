import { useState, useCallback } from 'react';

const useToast = () => {
  const [toasts, setToasts] = useState([]);

  const addToast = useCallback((message, type = 'success') => {
    const id = Math.random().toString(36).substr(2, 9);
    setToasts((currentToasts) => [...currentToasts, { id, message, type }]);
  }, []);

  const removeToast = useCallback((id) => {
    setToasts((currentToasts) => currentToasts.filter((toast) => toast.id !== id));
  }, []);

  return { toasts, addToast, removeToast };
};

export default useToast; 