import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import OrderFlow from './components/Tester/OrderFlow'
import TableId from './components/TableId/TableId';
import RoomCode from './components/RoomCode/RoomCode';
import UsersList from './components/UsersList/UsersList'
//import UsersList from './components/UsersList/UsersList';
import Catalog from './components/Catalog/Catalog';
import Item from './components/Item/Item';
import Cart from './components/Cart/Cart';
//import Checkout from './components/Checkout/Checkout';
//import Tip from './components/Tip/Tip';
import './App.css';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<OrderFlow />} />
        <Route path="/table-id" element={<TableId />} />
        <Route path="/room-code" element={<RoomCode />} />
        <Route path="/catalog/users-list" element={<UsersList />} />
        <Route path="/catalog" element={<Catalog />} />
        <Route path="/catalog/item/:id" element={<Item />} />
        <Route path="/cart" element={<Cart />} />
      </Routes>
    </Router>
  );
}

export default App;
