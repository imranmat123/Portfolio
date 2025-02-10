import React, {useEffect, useState} from "react";
import axios from "axios";
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
    const itemsPerPage = 4
    const pagArray = UserProject.data || []
    const num = numberOfProjects.data || []


    const startIndex = currentPage * itemsPerPage

    const totalPage = Math.ceil(num/itemsPerPage)

    function handleNextPage(){
        console.log("Current page before:", currentPage);
        console.log("Total pages:", totalPage);
        console.log("UserProject:", UserProject);
        console.log("pagArray:", pagArray);
        if(currentPage < totalPage -1){
            setCurrentPage(currentPage+1)
        }
        console.log("Current page after:", currentPage);
    }
    const endIndex = startIndex + itemsPerPage


    const page = pagArray.slice(startIndex, endIndex)


    return numberOfProjects != null
    ?
        <div className="projectsLinks">{
            page.map((UserProject) =>
            <div className="projectNumber" key={UserProject.project_id}>
                <h3>{UserProject.project_name}</h3>
            </div>
            )
        }
            <button className="nextPage" onClick={handleNextPage}>
            </button>
        </div>
        :
        <div className="projectsLinks">
            <p>fuck you buddy</p>
        </div>

}

export default Projects