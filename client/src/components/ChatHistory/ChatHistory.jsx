import React from 'react';
import "./ChatHistory.scss";

const ChatHistory = ({ ch }) => {
    const messages = ch.map((msg, idx) => (
        <p className="ChatMsg" key={idx}>{msg.data}</p>
    ));

    return (
        <div className="ChatHistory">
            <h2>Chat History</h2>
            {messages}
        </div>
    );
};

export default ChatHistory;
