import React, { useState } from 'react';
import './Message.css';  // Import the CSS file

function Message() {
  const [inputValue, setInputValue] = useState('');
  const [response, setResponse] = useState('');

  const handleChange = (event) => {
    setInputValue(event.target.value);
  };

  const handleClick = () => {
    // Send the input value to the backend
    fetch('http://localhost:8080', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ message: inputValue }),
    })
      .then((response) => response.json())
      .then((data) => {
        // Set the response from the backend
        setResponse(data.message);
      })
      .catch((error) => {
        console.error('Error:', error);
      });
  };

  return (
    <div className="container">
      <h1 className="heading">Hello Welocome to React-Go Full Stack Application</h1>
      <input
        type="text"
        value={inputValue}
        onChange={handleChange}
        className="input"
        placeholder="Enter your message"
      />
      <button onClick={handleClick} className="button">
        Send
      </button>
      <div className="response">{response}</div>
    </div>
  );
}

export default Message
