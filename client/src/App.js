import './App.css';
import React, { useEffect, useState } from 'react';
import Header from "./components/Header/Header";
import ChatHistory from "./components/ChatHistory/ChatHistory";
import { connect, sendMessage } from "./api/clientAPI";

function App() {
  const [chatHistory, setChatHistory] = useState([]);

  useEffect(() => {
    connect((msg) => {
      console.log("New Message");
      setChatHistory(chatHistory => [...chatHistory, msg]);
      console.log(chatHistory);
    });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const send = () => {
    sendMessage("Client Message!")
  };

  return (
    <div className="App">
      <Header />
      <ChatHistory ch={chatHistory} />
      <button onClick={send}>Send Message</button>
    </div>
  );
}

export default App;
