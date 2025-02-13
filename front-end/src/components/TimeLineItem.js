import React from "react";
import "../App.css"


const TimelineItem = ({anything, side}) => {
    return(
        <div className="timeline-Item">
            <div className="timeline-dot">
            </div>
                <div className={`timeline-content-${side}`}>
                    <h2>{anything.title}</h2>
                    <a>{anything.discription}</a>
                </div>
        </div>

    )
}


export default TimelineItem