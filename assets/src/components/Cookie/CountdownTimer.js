import React from "react";
import { useCountdown } from "../../hooks/useCountdown";

const ExpiredNotice = () => {
    return (
        <div>
            Claim Cookie
        </div>
    );
};

const ShowCounter = ({ days, hours, minutes, seconds }) => {
    if (isNaN(days) || isNaN(hours) || isNaN(minutes) || isNaN(seconds)) {
        return (
            <div>
                Loading...
            </div>
        )
    }
    return (
        <div>
            {days > 9 ? seconds : "0" + days}:{hours > 9 ? hours : "0" + hours}:{minutes > 9 ? minutes : "0" + minutes}:{seconds > 9 ? seconds : "0" + seconds}
        </div>
    );
};

const CountdownTimer = ({ targetDate }) => {
    const [days, hours, minutes, seconds] = useCountdown(targetDate);

    if (days + hours + minutes + seconds <= 0) {
        return <ExpiredNotice />;
    } else {
        return (
            <ShowCounter
                days={days}
                hours={hours}
                minutes={minutes}
                seconds={seconds}
            />
        );
    }
};

export default CountdownTimer;