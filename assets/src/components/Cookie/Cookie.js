import React, { useContext, useEffect, useState } from "react";
import "./Cookie.css";
import AuthContext from "../../store/auth-context";
import CountdownTimer from "./CountdownTimer";

const Cookie = (props) => {
    const authContext = useContext(AuthContext);
    const [cookieCount, setCookieCount] = useState(props.user?.Cookies);
    const [targetDate, setTargetDate] = useState(Date.parse(props.user?.LastClaimed) + (1000 * 60 * 30))

    const tryAddCookie = async () => {
        try {
            const response = await fetch("/api/cookies", {
                method: "PUT",
                headers: {
                    Authorization: "Bearer " + authContext.token,
                },
            });
            const data = await response.json();
            if (response.status === 200) {
                // succesfully claimed cookies
                const newTotal = data.data.Cookies;
                // TODO: create an alert or something showing how many cookies you got
                // const addedCookies = newTotal - cookieCount;
                // console.log(addedCookies); 
                const lastClaimed = data.data.LastClaimed;
                setCookieCount(newTotal);
                setTargetDate(Date.parse(lastClaimed) + (1000 * 60 * 30));
            } else if (response.status === 202) {
                // cooldown not done
                // TODO: maybe disable the button when theres a cooldown instead of allowing the call to get here?
                console.log("cooldown not done");
            }
        } catch (err) {
            // TODO: handle errors?
            console.log(err);
        }
    };

    useEffect(() => {
        setTargetDate(Date.parse(props.user?.LastClaimed) + (1000 * 60 * 30))
        setCookieCount(props.user?.Cookies);
    }, [props])

    console.log(Date.parse(props.user?.LastClaimed) + (1000 * 60 * 30));

    return (
        <div className="cookie">
            <button onClick={tryAddCookie} className="cookie-button">
                <div className="cookie-text">
                    <CountdownTimer targetDate={targetDate} />
                </div>
                <div className="cookie-count">
                    Total Cookies: {cookieCount}
                </div>
            </button>
        </div>
    );
};

export default Cookie;
