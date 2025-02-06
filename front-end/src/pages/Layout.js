import React from "react";
import {Link, Route, Routes} from "react-router-dom";
import Home from "./Home";
import About from "./About";
import Contact from "./Contact";
import ProfilePicture from "../components/ProfilePicture";

function Layout(){

    return(
        <div className="container">
            <div className="navbar">
                <Link to="/">Home</Link>
                <Link to="/about">About</Link>
                <Link to="/contact">Contant</Link>
            </div>
            <div className="profilePicture">
                <ProfilePicture />
            </div>
            <div className="mainSection">
                <Routes>
                    <Route path="/" element={<Home />}></Route>
                    <Route path="/about" element={<About />}></Route>
                    <Route path="/contact" element={<Contact />}></Route>
                </Routes>
            </div>

        </div>
);

}

export default Layout