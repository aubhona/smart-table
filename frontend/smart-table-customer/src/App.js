import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import TableId from './components/TableId/TableId';
import RoomCode from './components/RoomCode/RoomCode';
//import UsersList from './components/UsersList/UsersList';
import Catalog from './components/Catalog/Catalog';
import Item from './components/Item/Item';
//import Cart from './components/Cart/Cart';
//import Checkout from './components/Checkout/Checkout';
//import Tip from './components/Tip/Tip';
import './App.css';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<TableId />} />
        <Route path="/room-code" element={<RoomCode />} />
        <Route path="/catalog" element={<Catalog />} />
        <Route path="/catalog/item/:id" element={<Item />} />
      </Routes>
    </Router>
  );
}

export default App;
