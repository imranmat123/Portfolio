import React, {useEffect, useState} from "react";
import axios from "axios";
import Slider from "react-slick";
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";
import "../App.css"

const Projects = () => {

    const [numberOfProjects, setNumberOfProjects] = useState([])

    useEffect(() => {
            axios.get("http://localhost:8080/projects/GetProjectsByUserID/1")
                .then(response => {setNumberOfProjects(response.data)})
        },
        [])

    const [UserProject,setUserProject] = useState([])

    useEffect(() =>{
        axios.get("http://localhost:8080/projects/GetAllUserProjects/1")
            .then(response =>{
                setUserProject(response.data)})

    },[])

    const [currentPage, setCurrentPage] = useState(0)
    const pagArray = UserProject.data || []


    const settings = {
        infinite: true,
        slidesToShow: 3,
        slidesToScroll: 1,
        beforeChange: (current, next) => {
            setCurrentPage(next);
        },
        }


    return numberOfProjects != null
    ?
        <div className="projectsLinks">
            <Slider {...settings}>{
                pagArray.map((UserProject) =>
            <div className="projectNumber" key={UserProject.project_id}>
                <h3>{UserProject.project_name}</h3>
            </div>
            )
        }
            </Slider>
        </div>
        :
        <div className="projectsLinks">
            <p>fuck you buddy</p>
        </div>

}

export default Projects