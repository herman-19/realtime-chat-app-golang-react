import React from 'react';
import "./ChatHistory.scss";

const ChatHistory = ({ messages }) => {
    const msgs = messages.map((msg, idx) => (
        <p className="ChatMsg" key={idx}>{msg.data}</p>
    ));

    return (
        <div className="ChatHistory">
            <h2>Chat History</h2>
            {msgs}
        </div>
    );
};

export default ChatHistory;
