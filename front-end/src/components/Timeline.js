import React, {useEffect} from "react";
import TimelineItem from "./TimeLineItem";
import "../App.css"
import TimeLineItem from "./TimeLineItem";

const Timeline = () => {


    const timelineBoxes = [
        {id:1, title: "number one", discription: "lol im num 1"},
        {id:2, title: "number two", discription: "lol im num 2"},
        {id:3, title: "number tree", discription: "lol im num 3"},
        {id:4, title: "number four", discription: "lol im num 4"}]

    useEffect(() =>{
        const timeItem = document.querySelectorAll(".timeline-Item")

        const observer = new IntersectionObserver(entries =>{
            entries.forEach(entry => {
                console.log(entry)
                entry.target.classList.toggle("show", entry.isIntersecting)
            })
        },{threshold: 0.1})

        timeItem.forEach(item =>{
            observer.observe(item)
    })


})

    return <div className="timeline-box">
        <div className="timeline-line">
            {
                timelineBoxes.map((items, index)=>(
                    <TimelineItem anything={items} key={items.id} title={items.title} discription={items.discription}
                                  side={index %2 == 0? "left": "right"}/>
                ))
            }
        </div>

    </div>
}

export default Timeline