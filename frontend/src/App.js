import React, { useState } from 'react';
import './App.css';
import { EchoMessageRequest, MessageEchoServiceClient } from './proto/MessageEchoServiceClientPb'; // Update the import path

const client = new MessageEchoServiceClient('http://localhost:8080');

function App() {
  const [message, setMessage] = useState('');

  const handleClick = () => {
    const request = new EchoMessageRequest();
    request.setMessage('Hello'); // Set your message here

    client.echoMessage(request, {}, (err, response) => {
      if (err) {
        console.error('Error:', err);
        setMessage('Error occurred. Please try again later.');
      } else {
        setMessage(response.getMessage());
      }
    });
  };

  return (
    <div className="App">
      <input type="text" />
      <button onClick={handleClick}>Ping</button>
      <p>Response: {message}</p>
    </div>
  );
}

export default App;
