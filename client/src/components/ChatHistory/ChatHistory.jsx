import React from 'react';
import Message from "../Message/Message";
import "./ChatHistory.scss";

const ChatHistory = ({ messages }) => {
    const msgs = messages.map((msg, idx) => (
        <Message key={idx} message={msg.data} />
    ));

    return (
        <div className="ChatHistory">
            <h2>Chat History</h2>
            {msgs}
        </div>
    );
};

export default ChatHistory;
