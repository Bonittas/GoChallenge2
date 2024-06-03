import React from 'react';
import './App.css';
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Message from './components/Message';

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Message />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
