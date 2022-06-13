import React, { useContext, useEffect, useState } from "react";
import "./NavBar.css";

import AuthContext from "../../store/auth-context";
import { Link, useNavigate } from "react-router-dom";

const NavBar = () => {
    const navigate = useNavigate();
    const authContext = useContext(AuthContext);
    const [usernameText, setUsernameText] = useState("Profile");

    const getUser = async () => {
        try {
            const response = await fetch("/api/user", {
                method: "GET",
                headers: {
                    "Authorization": "Bearer " + authContext.token,
                }
            });
            const data = await response.json();
            setUsernameText(data.data.Username);
        } catch (err) {
            console.log(err);
            setUsernameText("Profile");
        }
    }

    useEffect(() => {
        if (authContext.loggedIn) {
            getUser();
        }
    }, [authContext])

    const logoutHandler = () => {
        authContext.logout();
        navigate("/", { replace: true });
    }

    return (
        <div  id="navbar" className="container-fluid d-flex justify-content-between py-2">
            <div className="px-2 mx-2">
                <Link className="custom-navbar-link" to="/">Home</Link>
            </div>
            <div className="d-flex flex-row-reverse">
                {!authContext.loggedIn ? (
                    <div className="px-2 mx-2">
                        <Link className="custom-navbar-link" to="/auth">Login/Signup</Link>
                    </div>
                ) : (
                    <>
                        <div className="px-2 mx-2">
                            <div className="custom-navbar-link" onClick={logoutHandler}>Logout</div>
                        </div>
                        <div className="px-2 mx-2">
                            <Link className="custom-navbar-link" to="/user">{usernameText}</Link>
                        </div>
                    </>
                )}
            </div>
        </div>
    )
}

export default NavBar;