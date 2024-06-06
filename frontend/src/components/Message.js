import React, { useState } from 'react';
import { PingServiceClient } from '../proto/service_grpc_web_pb';
import { PingRequest } from '../proto/service_pb';
import './Message.css';

const backendService = new PingServiceClient('http://localhost:8080');

const sendPingRequestWithRetry = (message, retries, callback) => {
  const attemptRequest = (attempt) => {
    const request = new PingRequest();
    request.setMessage(message);

    backendService.ping(request, {}, (err, response) => {
      if (err) {
        console.error('Error:', err.message);
        if (attempt < retries) {
          setTimeout(() => attemptRequest(attempt + 1), 1000); // Retry after 1 second
        } else {
          callback(err, null);
        }
      } else {
        callback(null, response.getMessage());
      }
    });
  };

  attemptRequest(0);
};

function Message() {
  const [inputMessage, setInputMessage] = useState('');
  const [responseMessage, setResponseMessage] = useState('');

  const handleClick = () => {
    sendPingRequestWithRetry(inputMessage, 3, (err, response) => {
      if (err) {
        setResponseMessage('Error: ' + err.message);
      } else {
        setResponseMessage(response);
      }
    });
  };

  return (
    <>     
    
    <div className="message-container">
        <div className='title'> <h1>React Go Mini Project</h1></div>
      <input 
        type="text" 
        value={inputMessage} 
        onChange={(e) => setInputMessage(e.target.value)} 
        placeholder="Enter your message" 
      />
      <button onClick={handleClick}>Send Ping</button>
      <p>{responseMessage}</p>
    </div>
    </>
  );
}

export default Message;
