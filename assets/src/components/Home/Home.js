import React, { useContext, useEffect, useState } from "react";
import { Link } from "react-router-dom";
import AuthContext from "../../store/auth-context";
import Cookie from "../Cookie/Cookie";

const Home = () => {
    const authContext = useContext(AuthContext);

    const [user, setUser] = useState();

    const getUser = async () => {
        try {
            const response = await fetch("/api/user", {
                method: "GET",
                headers: {
                    "Authorization": "Bearer " + authContext.token,
                }
            });
            const data = await response.json();
            setUser(data.data);
        } catch (err) {
            console.log(err);
            setUser({});
        }
    }

    useEffect(() => {
        if (authContext.loggedIn) {
            getUser();
        }
    }, [])

    return (
        <div className="text-center home-page pt-1 mt-3">
            {/* maybe in the future add ability to play as guest, put for now account is required */}
            {authContext.loggedIn ? (
                <div>
                    <Cookie user={user} />
                </div>
            ) : (
                <>            
                    <h2>Welcome!</h2>
                    <div className="pt-2">
                        <h4>
                            <>To play you must </>
                            <Link className="custom-navbar-link" to="/auth">Login or Signup</Link>
                        </h4>
                    </div>
                </>
            )}
        </div>
    )
}

export default Home;