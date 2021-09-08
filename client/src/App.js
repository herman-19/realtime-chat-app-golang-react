import './App.css';
import React, { useEffect, useState } from 'react';
import Header from "./components/Header/Header";
import ChatHistory from "./components/ChatHistory/ChatHistory";
import ChatInput from "./components/ChatInput/ChatInput";
import { connect, sendMessage } from "./api/clientAPI";

function App() {
  const [chatHistory, setChatHistory] = useState([]);

  useEffect(() => {
    connect((msg) => {
      setChatHistory(chatHistory => [...chatHistory, msg]);
      console.log(chatHistory);
    });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const send = (e) => {
    if (e.keyCode === 13) {
      sendMessage(e.target.value);
      e.target.value = "";
    }
  };

  return (
    <div className="App">
      <Header />
      <ChatHistory messages={chatHistory} />
      <ChatInput send={send} />
    </div>
  );
}

export default App;
