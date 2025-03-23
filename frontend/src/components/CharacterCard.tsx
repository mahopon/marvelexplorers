import React from 'react'

interface Props {
    key: string
    name: string,
    description?: string,
    resourceURI: string,
    thumbnailPath: string,
    thumbnailExtension: string
}

const CharacterCard = (props: Props) => {
  return (
    <div className="charCard fade-in">
        <img className="block" src={props.thumbnailPath+"."+props.thumbnailExtension} alt={"Portrait of "+props.name} loading="eager"/>
        <p className="inline-block">{props.name}</p>
        {/* {props.description ? (
        <p>{props.description}</p> // Renders description if it's defined
      ) : (
        <p>No description available.</p> // Renders this if description is undefined
      )} */}
    </div>
  )
}

export default CharacterCard