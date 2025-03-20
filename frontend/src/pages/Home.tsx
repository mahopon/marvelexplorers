import React from "react";
import CharacterCard from "../components/CharacterCard.tsx";

const Home = () => {
    return (
        <div>
            <h1>Welcome to the Marvel Universe!</h1>
            <p>Explore marvel characters and their comics</p>
            <CharacterCard
              name="3-D Man"
              description=""
              resourceURI="http://gateway.marvel.com/v1/public/characters/1011334"
              thumbnailPath="http://i.annihil.us/u/prod/marvel/i/mg/c/e0/535fecbbb9784"
              thumbnailExtension="jpg"
            />
        </div>
    )
}

export default Home;