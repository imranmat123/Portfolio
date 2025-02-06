import React, {useEffect, useState} from "react";
import axios from "axios";

const AboutMe = () =>{
    const[aboutMe, setAboutMe] = useState(null)


    useEffect(() => {
        function fetchAboutMe(){
            axios.get("http://localhost:8080/users/aboutMe/1")
                .then(function(response){setAboutMe(response.data)})
        }
        fetchAboutMe()
    },[])


    return aboutMe != null
    ?
        <a>{aboutMe?.data}</a>
        :
        <a>"we haven't been able to pull from the database"</a>
}

export default AboutMe