import React from 'react'

interface Props {
    name: string,
    description?: string,
    resourceURI: string,
    thumbnailPath: string,
    thumbnailExtension: string
}

const CharacterCard = (props: Props) => {
  return (
    <div className="charCard">
        <img src={props.thumbnailPath+"."+props.thumbnailExtension} height="150px" width="150px" alt={"Portrait of "+props.name}/>
        <p>{props.name}</p>
        {props.description ? (
        <p>{props.description}</p> // Renders description if it's defined
      ) : (
        <p>No description available.</p> // Renders this if description is undefined
      )}
    </div>
  )
}

export default CharacterCard