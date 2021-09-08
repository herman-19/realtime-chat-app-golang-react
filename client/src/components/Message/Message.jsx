import React from 'react';
import "./Message.scss";

const Message = ({ message }) => {
    //const [message, setMessage] = useState(message);
    let msgObj = JSON.parse(message);

    return (
        <div className="Message">
            {msgObj.body}
        </div>
    )
};

export default Message;
