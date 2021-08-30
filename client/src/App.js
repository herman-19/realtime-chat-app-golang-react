import './App.css';
import React, { useEffect } from 'react';
import { connect, sendMessage } from "./api/clientAPI";

function App() {
  useEffect(() => {
    connect();
  }, []);

  const send = () => {
    sendMessage("Client Message!")
  };

  return (
    <div className="App">
      <button onClick={send}>Send Message</button>
    </div>
  );
}

export default App;
