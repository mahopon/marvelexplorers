import React from 'react';
import {Character} from "../interfaces/CharacterInterface.tsx";

interface CharacterProps {
  character: Character;
  onClick: (char: Character) => void;
}
// Dispatch is type of function that is used to modify state
// SetStateAction is the action passed to the dispatch to set the new value
const CharacterCard: React.FC<CharacterProps> = ({ character, onClick }: { character: Character, onClick: (char: Character) => void}) => {
  const handleClick = () => {
    onClick(character); // Pass the character object to the onClick handler
    console.log(character);
  };
  return (
    <div onClick={handleClick} className="charCard fade-in">
        <img className="block" src={character.thumbnailPath+"."+character.thumbnailExtension} alt={"Portrait of "+character.name} loading="eager"/>
        <p className="inline-block">{character.name}</p>
    </div>
  )
}

export default CharacterCard