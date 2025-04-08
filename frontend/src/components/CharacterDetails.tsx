import React from 'react';
import { Character } from '../interfaces/CharacterInterface';
import "../styles/CharacterDetails.css";

const CharacterDetails = ({character, setSelectedChar}: {character: Character, setSelectedChar: React.Dispatch<React.SetStateAction<Character | null>> }) => {
    const handleClick = () => {
        setSelectedChar(null);
    };
    
  return (
    <div id="details-popup" className="fade-in">
      <div className="flex-container">
        <div>
          <div className="charInfo">
            <img src={character.thumbnailPath+"."+character.thumbnailExtension}></img>
            <p>Name: {character.name}</p>
            <p>Description: {character.description ? (
              character.description
            )
            : ("No description.")
            }</p>
          </div>
        </div>
        <div>
          <div className="Comics">
            Test
          </div>
          <div className="Stories">
            Test
          </div>
          <div className="Events">
            Test
          </div>
        </div>
      </div>
      <button type="button" className='closeBtn' onClick={handleClick}>Close</button>
    </div>
  )
}

export default CharacterDetails