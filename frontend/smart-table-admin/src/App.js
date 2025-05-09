import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./pages/Home";
import RegistrationForm from "./pages/RegistrationForm";
import LoginForm from "./pages/LoginForm";
import RestaurantsList from "./pages/RestaurantsList.js"
import PlacesAndDishes from "./pages/PlacesAndDishes.js"
import PlaceDetail from "./pages/PlaceDetail.js";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/register" element={<RegistrationForm />} />
        <Route path="/login" element={<LoginForm />} />
        <Route path="/restaurants" element={<RestaurantsList />} />
        <Route path="/restaurants/:restaurant_uuid/places-dishes" element={<PlacesAndDishes />} />
        <Route path="/restaurants/:restaurant_uuid/:place_uuid" element={<PlaceDetail />} />
      </Routes>
    </Router>
  );
}

export default App;
