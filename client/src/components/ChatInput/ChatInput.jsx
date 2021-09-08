import React from 'react';
import "./ChatInput.scss";

const ChatInput = ({ send }) => {
    return (
        <div className="ChatInput">
            <input onKeyDown={send} />
        </div>
    )
};

export default ChatInput;
