import React from "react";
import AboutMe from "../components/AboutMe";

function About(){
    return (
            <div className="HomeContainer">
                <div className="titleOfPage">
                    <h2>Hello from the About page</h2>
                </div>
                <div className="aboutInfo">
                    <AboutMe/>
                </div>

            </div>);

}

export default About