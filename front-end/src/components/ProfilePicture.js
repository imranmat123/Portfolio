import {useState,useEffect} from "react";
import axios from "axios";
import React from "react";

const ProfilePicture = () => {

    const [profilePicture, setProfilePicture] =  useState(null)

    useEffect(() => {
            axios.get("http://localhost:8080/users/profilePicture/1")
                .then(response => {setProfilePicture(response.data)
            })
        },[]
    )

    return profilePicture != null
       ?
           <img src={profilePicture?.data} alt="User profile"></img>
           :
           <img src="/assets/image/userPorfile/stock.jpeg" alt="User profile"></img>

}

export default ProfilePicture