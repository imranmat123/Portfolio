import React from 'react';
import Layout from "./pages/Layout";
import './App.css';
import {BrowserRouter} from "react-router-dom";
import AboutMe from "./components/AboutMe";
import GitHubLogo from "./components/GitHubLogo";


function App() {
  return (

    <div className="mainBody">

        <BrowserRouter>
            <Layout/>
        </BrowserRouter>
        <div className="projects">
            <div className="titleOfProjects">
                <h2>My Projects</h2>
            </div>
            <div className="projetsMainBox">
                <div className="githubLogo">
                    <GitHubLogo/>
                </div>
                <div className="projectsLinks">

                </div>

            </div>

        </div>
    </div>
  );
}

export default App;
